package main

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestSlug(t *testing.T) {
	for in, want := range map[string]string{
		"Images, PDF, PPTX and ODP":            "images-pdf-pptx-and-odp",
		"`slides.yaml`":                        "slidesyaml",
		"What it does today":                   "what-it-does-today",
		"A [link](https://x.test) in **bold**": "a-link-in-bold",
		"Vector PDF":                           "vector-pdf",
		"HCL: expressions & more":              "hcl-expressions--more",
	} {
		if got := Slug(in); got != want {
			t.Errorf("Slug(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestIDsNumberRepeats(t *testing.T) {
	hs := []Heading{{Text: "Notes"}, {Text: "Notes"}, {Text: "Other", ID: "custom"}, {Text: "Notes"}}
	if got, want := IDs(hs), []string{"notes", "notes-1", "custom", "notes-2"}; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

const english = "---\ntitle: Hello\ndescription: A page.\nweight: 1\n---\n\n# Hello\n\nRun `slides build` or ``a `b` c``,\nsee [docs]({{< relref \"cli#flags\" >}}) and [x](https://x.test).\n\n## Flags\n\n```sh\nslides --pdf # a comment\n```\n\n- in a list:\n\n  ```go\n  fmt.Println(\"hi\")\n  ```\n"

func translated(mod func(string) string) string {
	tr := "---\ntitle: Bonjour\ndescription: Une page.\nweight: 1\n---\n\n# Bonjour {#hello}\n\nLancez `slides build` ou ``a `b` c``,\nvoyez [la doc]({{< relref \"cli#flags\" >}}) et [x](https://x.test).\n\n## Options {#flags}\n\n```sh\nslides --pdf # a comment\n```\n\n- dans une liste :\n\n  ```go\n  fmt.Println(\"hi\")\n  ```\n"
	if mod != nil {
		tr = mod(tr)
	}
	return tr
}

func TestCompareAcceptsAFaithfulTranslation(t *testing.T) {
	if got := Compare(english, translated(nil)); len(got) != 0 {
		t.Errorf("unexpected problems: %v", got)
	}
}

func TestCompareFindsWhatChanged(t *testing.T) {
	for name, c := range map[string]struct {
		mod  func(string) string
		want string
	}{
		"code span":    {func(s string) string { return strings.Replace(s, "`slides build`", "`slides construire`", 1) }, "inline code spans differ"},
		"fence":        {func(s string) string { return strings.Replace(s, "# a comment", "# un commentaire", 1) }, "fenced code blocks differ"},
		"indented":     {func(s string) string { return strings.Replace(s, `fmt.Println("hi")`, `fmt.Println("salut")`, 1) }, "fenced code blocks differ"},
		"shortcode":    {func(s string) string { return strings.Replace(s, "cli#flags", "cli#options", 1) }, "shortcodes differ"},
		"target":       {func(s string) string { return strings.Replace(s, "https://x.test", "https://y.test", 1) }, "link targets differ"},
		"heading":      {func(s string) string { return strings.Replace(s, "\n## Options {#flags}\n", "\n", 1) }, "1 headings, the English page has 2"},
		"level":        {func(s string) string { return strings.Replace(s, "## Options", "### Options", 1) }, "is level 3"},
		"anchor":       {func(s string) string { return strings.Replace(s, "{#flags}", "", 1) }, "needs {#flags}"},
		"wrong anchor": {func(s string) string { return strings.Replace(s, "{#flags}", "{#options}", 1) }, "needs {#flags}"},
		"description":  {func(s string) string { return strings.Replace(s, "description: Une page.\n", "", 1) }, "no description"},
	} {
		t.Run(name, func(t *testing.T) {
			got := strings.Join(Compare(english, translated(c.mod)), "\n")
			if !strings.Contains(got, c.want) {
				t.Errorf("want %q in:\n%s", c.want, got)
			}
		})
	}
}

func TestSpansSurviveRewrapping(t *testing.T) {
	a := Parse("x `one\ntwo` y")
	b := Parse("x `one two` y")
	if !reflect.DeepEqual(a.Spans, b.Spans) {
		t.Errorf("%v != %v", a.Spans, b.Spans)
	}
}

func TestAddAnchors(t *testing.T) {
	bare := strings.ReplaceAll(translated(nil), " {#hello}", "")
	bare = strings.ReplaceAll(bare, " {#flags}", "")
	got, err := AddAnchors(english, bare)
	if err != nil {
		t.Fatal(err)
	}
	if got != translated(nil) {
		t.Errorf("got:\n%s", got)
	}
	if _, err := AddAnchors(english, "# one heading only\n"); err == nil {
		t.Error("headings that do not pair up were anchored")
	}
}

func tree(t *testing.T, files map[string]string) string {
	t.Helper()
	root := t.TempDir()
	for p, s := range files {
		p = filepath.Join(root, filepath.FromSlash(p))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(s), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

const hugo = "title: x\nlanguages:\n  en:\n    label: English\n  fr:\n    label: Français\n  zh:\n    label: 简体中文\n    locale: zh-Hans\nenableRobotsTXT: true\n"

func TestLanguages(t *testing.T) {
	root := tree(t, map[string]string{"hugo.yaml": hugo})
	got, err := Languages(filepath.Join(root, "hugo.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, []string{"fr", "zh"}) {
		t.Errorf("got %v", got)
	}
	root = tree(t, map[string]string{"hugo.yaml": "title: x\n"})
	if _, err := Languages(filepath.Join(root, "hugo.yaml")); err == nil {
		t.Error("a site without languages was accepted")
	}
}

func TestRunFailsOnAMissingPageAndAnOrphan(t *testing.T) {
	root := tree(t, map[string]string{
		"content/en/a.md":      english,
		"content/en/sub/b.md":  english,
		"content/fr/a.md":      translated(nil),
		"content/fr/orphan.md": translated(nil),
		"content/zh/a.md":      translated(nil),
		"content/zh/sub/b.md":  translated(nil),
	})
	problems, err := Run(root, []string{"fr", "zh"}, false)
	if err != nil {
		t.Fatal(err)
	}
	all := strings.Join(problems, "\n")
	for _, want := range []string{"content/fr/sub/b.md is missing", "content/fr/orphan.md has no English page"} {
		if !strings.Contains(all, want) {
			t.Errorf("missing %q in:\n%s", want, all)
		}
	}
	if strings.Contains(all, "content/zh") {
		t.Errorf("zh is complete, yet:\n%s", all)
	}
}

func TestRunFixAddsAnchors(t *testing.T) {
	bare := strings.ReplaceAll(strings.ReplaceAll(translated(nil), " {#hello}", ""), " {#flags}", "")
	root := tree(t, map[string]string{"content/en/a.md": english, "content/fr/a.md": bare})
	if problems, _ := Run(root, []string{"fr"}, false); len(problems) == 0 {
		t.Fatal("a translation without anchors passed")
	}
	if problems, err := Run(root, []string{"fr"}, true); err != nil || len(problems) != 0 {
		t.Fatalf("after -fix: %v %v", problems, err)
	}
}

func TestCheckBuilt(t *testing.T) {
	page := func(ids ...string) string {
		var b strings.Builder
		for _, id := range ids {
			b.WriteString("<h2 id=" + id + ">x</h2><h3 class=y id=\"" + id + "-2\">y</h3>")
		}
		return b.String()
	}
	root := tree(t, map[string]string{
		"index.html":        page("a"),
		"cli/index.html":    page("flags", "pdf"),
		"fr/index.html":     page("a"),
		"fr/cli/index.html": page("options", "pdf"),
	})
	problems, n, err := CheckBuilt(root, []string{"fr"})
	if err != nil || n != 2 {
		t.Fatalf("n=%d err=%v", n, err)
	}
	if len(problems) != 1 || !strings.Contains(problems[0], "fr/cli/index.html") {
		t.Errorf("problems: %v", problems)
	}
}
