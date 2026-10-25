// Command i18n checks that the documentation's translations pair up with the
// English pages, and gives translated headings the English anchors.
//
// Every page of content/en/ must exist, under the same name, in each other
// language; no page may exist only in a translation. A translation must keep
// what is not prose exactly as the English page has it:
//
//   - the headings: same number, same levels, in the same order, each with the
//     English heading's anchor as an explicit {#id}, so that links into a page
//     (a relref with a fragment, a bookmark) reach the same place in every
//     language;
//   - fenced code blocks, character for character;
//   - inline code spans, as a multiset (whitespace runs count as one space);
//   - Hugo shortcodes, in order;
//   - Markdown link targets, as a multiset;
//   - a title and a description in the front matter.
//
// Usage, from the documentation repository's root:
//
//	go run ./tools/i18n            # check; exit status 1 if a page is missing in a language
//	go run ./tools/i18n -strict    # ... and if a translation differs where it must not
//	go run ./tools/i18n -fix       # add the missing {#id} to translated headings, then check
//	go run ./tools/i18n -built public   # after hugo: the built pages' heading ids
//
// The languages are those of hugo.yaml (languages:), so that declaring one
// there without content fails the build here first.
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strings"
	"unicode"
)

func main() {
	root := flag.String("root", ".", "the documentation repository's root")
	fix := flag.Bool("fix", false, "add the English anchors to translated headings")
	strict := flag.Bool("strict", false, "fail on every difference, not only on a missing page")
	built := flag.String("built", "", "a built site (hugo's public/): check that every page's headings have the English page's ids")
	flag.Parse()
	if *built != "" {
		langs, err := Languages(filepath.Join(*root, "hugo.yaml"))
		if err != nil {
			fmt.Fprintln(os.Stderr, "i18n:", err)
			os.Exit(1)
		}
		problems, n, err := CheckBuilt(*built, langs)
		if err != nil {
			fmt.Fprintln(os.Stderr, "i18n:", err)
			os.Exit(1)
		}
		for _, p := range problems {
			if *strict {
				fmt.Println(p)
			} else {
				fmt.Println("::warning::" + p)
			}
		}
		if len(problems) > 0 && *strict {
			fmt.Printf("i18n: %d problem(s) in the built site\n", len(problems))
			os.Exit(1)
		}
		fmt.Printf("i18n: %d translated pages of the built site compared with the English headings' ids (%d warning(s))\n", n, len(problems))
		return
	}
	langs, err := Languages(filepath.Join(*root, "hugo.yaml"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "i18n:", err)
		os.Exit(1)
	}
	problems, err := Run(*root, langs, *fix)
	if err != nil {
		fmt.Fprintln(os.Stderr, "i18n:", err)
		os.Exit(1)
	}
	// A page missing in a language, or only in a translation, fails. A
	// translation that differs from its English page where it must not is
	// reported as a warning (an annotation in GitHub Actions), unless -strict:
	// an English edit then does not turn every pull request red until the
	// translations follow, and a broken link into a page is caught by the link
	// check anyway.
	failures := 0
	for _, p := range problems {
		if *strict || Missing(p) {
			fmt.Println(p)
			failures++
		} else {
			fmt.Println("::warning::" + p)
		}
	}
	if failures > 0 {
		fmt.Printf("i18n: %d problem(s)\n", failures)
		os.Exit(1)
	}
	fmt.Printf("i18n: %d languages, every page of content/en has its translation (%d warning(s))\n", len(langs), len(problems))
}

// Missing reports whether a problem is a page that exists in only one of the
// two languages.
func Missing(problem string) bool {
	return strings.Contains(problem, " is missing (") || strings.Contains(problem, " has no English page")
}

var langRE = regexp.MustCompile(`(?m)^  ([a-z]{2}(?:-[A-Za-z]+)?):\s*$`)

// Languages returns the codes of hugo.yaml's languages, other than en, in
// file order: the two-space-indented keys under "languages:".
func Languages(hugoYAML string) ([]string, error) {
	b, err := os.ReadFile(hugoYAML)
	if err != nil {
		return nil, err
	}
	text := string(b)
	i := strings.Index(text, "\nlanguages:\n")
	if i < 0 {
		return nil, fmt.Errorf("%s declares no languages:", hugoYAML)
	}
	text = text[i+len("\nlanguages:\n"):]
	// the block ends at the next unindented line
	if j := regexp.MustCompile(`(?m)^\S`).FindStringIndex(text); j != nil {
		text = text[:j[0]]
	}
	var out []string
	for _, m := range langRE.FindAllStringSubmatch(text, -1) {
		if m[1] != "en" {
			out = append(out, m[1])
		}
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("%s declares no language besides en", hugoYAML)
	}
	return out, nil
}

// Run checks every language, and with fix rewrites translated headings.
func Run(root string, langs []string, fix bool) ([]string, error) {
	en := filepath.Join(root, "content", "en")
	pages, err := pagesOf(en)
	if err != nil {
		return nil, err
	}
	var problems []string
	for _, lang := range langs {
		dir := filepath.Join(root, "content", lang)
		have, err := pagesOf(dir)
		if err != nil {
			problems = append(problems, fmt.Sprintf("%s: %v", lang, err))
			continue
		}
		for _, p := range pages {
			if !slices.Contains(have, p) {
				problems = append(problems, fmt.Sprintf("content/%s/%s is missing (content/en/%[2]s exists)", lang, p))
			}
		}
		for _, p := range have {
			if !slices.Contains(pages, p) {
				problems = append(problems, fmt.Sprintf("content/%s/%s has no English page", lang, p))
			}
		}
		for _, p := range pages {
			if !slices.Contains(have, p) {
				continue
			}
			src, err := os.ReadFile(filepath.Join(en, p))
			if err != nil {
				return nil, err
			}
			tr, err := os.ReadFile(filepath.Join(dir, p))
			if err != nil {
				return nil, err
			}
			if fix {
				fixed, err := AddAnchors(string(src), string(tr))
				if err == nil && fixed != string(tr) {
					if err := os.WriteFile(filepath.Join(dir, p), []byte(fixed), 0o644); err != nil {
						return nil, err
					}
					tr = []byte(fixed)
				}
			}
			for _, msg := range Compare(string(src), string(tr)) {
				problems = append(problems, fmt.Sprintf("content/%s/%s: %s", lang, p, msg))
			}
		}
	}
	return problems, nil
}

var headingIDRE = regexp.MustCompile(`<h[1-6][^>]*?\sid=(?:"([^"]+)"|([^\s>"]+))`)

// headingIDs lists the ids of the headings of a built page, in order.
func headingIDs(html []byte) []string {
	var out []string
	for _, m := range headingIDRE.FindAllSubmatch(html, -1) {
		id := string(m[1])
		if id == "" {
			id = string(m[2])
		}
		out = append(out, id)
	}
	return out
}

// CheckBuilt compares, in a built site, the ids of each translated page's
// headings with those of the English page at the same path (the languages
// live in public/<lang>/..., English at the root). It returns the problems and
// how many translated pages it compared.
func CheckBuilt(public string, langs []string) ([]string, int, error) {
	var problems []string
	n := 0
	err := filepath.WalkDir(public, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(public, p)
		first := strings.SplitN(filepath.ToSlash(rel), "/", 2)[0]
		if d.IsDir() && slices.Contains(langs, first) {
			return filepath.SkipDir
		}
		if d.IsDir() || d.Name() != "index.html" {
			return nil
		}
		en, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		want := headingIDs(en)
		for _, lang := range langs {
			tp := filepath.Join(public, lang, rel)
			tr, err := os.ReadFile(tp)
			if err != nil {
				continue // no translation of this generated page (a redirect, a tag page)
			}
			n++
			got := headingIDs(tr)
			if !slices.Equal(want, got) {
				problems = append(problems, fmt.Sprintf("%s/%s: heading ids %q, the English page's are %q", lang, filepath.ToSlash(rel), first4(got), first4(want)))
			}
		}
		return nil
	})
	return problems, n, err
}

func first4(s []string) []string {
	if len(s) > 4 {
		return append(slices.Clone(s[:4]), "…")
	}
	return s
}

// pagesOf lists the Markdown files under dir, slash-separated and sorted.
func pagesOf(dir string) ([]string, error) {
	var out []string
	err := filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(p, ".md") {
			rel, _ := filepath.Rel(dir, p)
			out = append(out, filepath.ToSlash(rel))
		}
		return nil
	})
	sort.Strings(out)
	return out, err
}

// Page is what a Markdown page is made of, as far as this check cares.
type Page struct {
	Front      string
	Headings   []Heading
	Fences     []string
	Spans      []string
	Shortcodes []string
	Targets    []string
}

// Heading is an ATX heading: its level, its text as written and its explicit
// id, if it has one.
type Heading struct {
	Level int
	Text  string
	ID    string
	Line  int // index in the page's lines
}

var (
	headingRE   = regexp.MustCompile(`^(#{1,6})\s+(.*?)\s*$`)
	idRE        = regexp.MustCompile(`\s*\{#([A-Za-z0-9_:.-]+)\}\s*$`)
	fenceRE     = regexp.MustCompile("^(\\s*)(`{3,}|~{3,})(.*)$")
	shortcodeRE = regexp.MustCompile(`\{\{[<%].*?[>%]\}\}`)
	targetRE    = regexp.MustCompile(`\]\(([^)\s]+)(?:\s+"[^"]*")?\)`)
)

// Parse reads a page.
func Parse(src string) Page {
	var p Page
	lines := strings.Split(src, "\n")
	i := 0
	if len(lines) > 0 && strings.TrimRight(lines[0], " \r") == "---" {
		for j := 1; j < len(lines); j++ {
			if strings.TrimRight(lines[j], " \r") == "---" {
				p.Front = strings.Join(lines[1:j], "\n")
				i = j + 1
				break
			}
		}
	}
	var prose []string
	for ; i < len(lines); i++ {
		line := lines[i]
		if m := fenceRE.FindStringSubmatch(line); m != nil {
			indent, fence, info := m[1], m[2], strings.TrimSpace(m[3])
			var body []string
			j := i + 1
			for ; j < len(lines); j++ {
				if c := fenceRE.FindStringSubmatch(lines[j]); c != nil && c[2][0] == fence[0] && len(c[2]) >= len(fence) && strings.TrimSpace(c[3]) == "" {
					break
				}
				body = append(body, strings.TrimPrefix(lines[j], indent))
			}
			p.Fences = append(p.Fences, info+"\n"+strings.Join(body, "\n"))
			i = j
			continue
		}
		if m := headingRE.FindStringSubmatch(line); m != nil {
			h := Heading{Level: len(m[1]), Text: m[2], Line: i}
			if id := idRE.FindStringSubmatch(h.Text); id != nil {
				h.ID = id[1]
				h.Text = idRE.ReplaceAllString(h.Text, "")
			}
			p.Headings = append(p.Headings, h)
		}
		prose = append(prose, line)
	}
	text := strings.Join(prose, "\n")
	p.Shortcodes = shortcodeRE.FindAllString(text, -1)
	for _, m := range targetRE.FindAllStringSubmatch(text, -1) {
		p.Targets = append(p.Targets, m[1])
	}
	p.Spans = codeSpans(shortcodeRE.ReplaceAllString(text, ""))
	return p
}

// codeSpans returns the inline code spans of text: a run of n backticks opens
// one and the next run of exactly n closes it. Whitespace runs inside a span
// count as one space.
func codeSpans(text string) []string {
	var out []string
	for i := 0; i < len(text); {
		if text[i] != '`' {
			i++
			continue
		}
		j := i
		for j < len(text) && text[j] == '`' {
			j++
		}
		n := j - i
		k := j
		closed := false
		for k < len(text) {
			if text[k] != '`' {
				k++
				continue
			}
			e := k
			for e < len(text) && text[e] == '`' {
				e++
			}
			if e-k == n {
				out = append(out, strings.Join(strings.Fields(text[j:k]), " "))
				i = e
				closed = true
				break
			}
			k = e
		}
		if !closed {
			i = j
		}
	}
	return out
}

// Slug is Hugo's default anchor for a heading (goldmark's "github" style):
// the text without its Markdown, lowercased, punctuation dropped, spaces and
// hyphens kept as hyphens. Duplicate slugs on a page are numbered by the caller.
func Slug(heading string) string {
	t := plain(heading)
	var b strings.Builder
	for _, r := range strings.ToLower(t) {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-':
			b.WriteRune(r)
		case unicode.IsSpace(r):
			b.WriteByte('-')
		}
	}
	return b.String()
}

var (
	linkTextRE = regexp.MustCompile(`\[([^\]]*)\]\([^)]*\)`)
	htmlTagRE  = regexp.MustCompile(`<[^>]+>`)
)

// plain removes the Markdown of an inline text: code backticks, emphasis,
// links (keeping their text) and HTML tags.
func plain(s string) string {
	s = shortcodeRE.ReplaceAllString(s, "")
	s = linkTextRE.ReplaceAllString(s, "$1")
	s = htmlTagRE.ReplaceAllString(s, "")
	return strings.NewReplacer("`", "", "**", "", "__", "", "*", "").Replace(s)
}

// IDs are the anchors of a page's headings: the explicit one, or the slug,
// with -1, -2 ... on a repeat, as Hugo numbers them.
func IDs(hs []Heading) []string {
	seen := map[string]int{}
	out := make([]string, len(hs))
	for i, h := range hs {
		id := h.ID
		if id == "" {
			id = Slug(h.Text)
		}
		base := id
		if n := seen[base]; n > 0 {
			id = fmt.Sprintf("%s-%d", base, n)
		}
		seen[base]++
		out[i] = id
	}
	return out
}

// Compare lists how a translation differs from its English page, where it
// must not.
func Compare(en, tr string) []string {
	a, b := Parse(en), Parse(tr)
	var out []string
	for _, key := range []string{"title", "description"} {
		if !regexp.MustCompile(`(?m)^` + key + `:\s*\S`).MatchString(b.Front) {
			out = append(out, "the front matter has no "+key)
		}
	}
	if len(a.Headings) != len(b.Headings) {
		out = append(out, fmt.Sprintf("%d headings, the English page has %d", len(b.Headings), len(a.Headings)))
	} else {
		ids := IDs(a.Headings)
		for i, h := range b.Headings {
			if h.Level != a.Headings[i].Level {
				out = append(out, fmt.Sprintf("heading %q is level %d, the English one is level %d", h.Text, h.Level, a.Headings[i].Level))
			}
			if h.ID != ids[i] {
				out = append(out, fmt.Sprintf("heading %q needs {#%s} (has %q)", h.Text, ids[i], h.ID))
			}
		}
	}
	if !slices.Equal(a.Fences, b.Fences) {
		out = append(out, diff("fenced code blocks", a.Fences, b.Fences))
	}
	if x, y := sorted(a.Spans), sorted(b.Spans); !slices.Equal(x, y) {
		out = append(out, diff("inline code spans", x, y))
	}
	if !slices.Equal(a.Shortcodes, b.Shortcodes) {
		out = append(out, diff("shortcodes", a.Shortcodes, b.Shortcodes))
	}
	if x, y := sorted(a.Targets), sorted(b.Targets); !slices.Equal(x, y) {
		out = append(out, diff("link targets", x, y))
	}
	return out
}

func sorted(s []string) []string {
	c := slices.Clone(s)
	sort.Strings(c)
	return c
}

// diff describes how got differs from want as multisets: what is missing and
// what is extra, a few of each.
func diff(what string, want, got []string) string {
	count := func(s []string) map[string]int {
		m := map[string]int{}
		for _, x := range s {
			m[x]++
		}
		return m
	}
	w, g := count(want), count(got)
	var missing, extra []string
	for k, n := range w {
		if g[k] < n {
			missing = append(missing, short(k))
		}
	}
	for k, n := range g {
		if w[k] < n {
			extra = append(extra, short(k))
		}
	}
	sort.Strings(missing)
	sort.Strings(extra)
	if len(missing) == 0 && len(extra) == 0 {
		return what + " are in another order than in the English page"
	}
	return fmt.Sprintf("%s differ: missing %q, extra %q", what, first(missing, 3), first(extra, 3))
}

func first(s []string, n int) []string {
	if len(s) > n {
		return s[:n]
	}
	return s
}

func short(s string) string {
	s = strings.ReplaceAll(s, "\n", "⏎")
	if r := []rune(s); len(r) > 60 {
		return string(r[:60]) + "…"
	}
	return s
}

// AddAnchors gives each heading of a translation the anchor of the English
// heading at the same position, when it has no explicit one. It changes
// nothing, and reports an error, when the two pages' headings do not pair up.
func AddAnchors(en, tr string) (string, error) {
	a, b := Parse(en), Parse(tr)
	if len(a.Headings) != len(b.Headings) {
		return tr, fmt.Errorf("%d headings against %d", len(b.Headings), len(a.Headings))
	}
	ids := IDs(a.Headings)
	lines := strings.Split(tr, "\n")
	for i, h := range b.Headings {
		if h.ID == "" {
			lines[h.Line] = strings.TrimRight(lines[h.Line], " \r") + " {#" + ids[i] + "}"
		}
	}
	return strings.Join(lines, "\n"), nil
}
