// Command gallery prepares the documentation's gallery of example decks.
//
// It reads the decks of go-slides/slides' examples/ directory, renders the
// first slide of each with go-slides itself (the slides command, built from the
// same checkout), and writes what the Hugo page needs:
//
//	data/gallery.json            one entry per deck, in the order of its name
//	static/gallery/NAME.jpg      the first slide
//	static/gallery/NAME.md       the deck's source, shown on the page
//	static/gallery/NAME.zip      the deck's folder, when it has files besides deck.md
//
// A deck that is one file also gets an "Open in the playground" link: the
// playground's share format, "#deck=" and the deck, deflated and base64url
// encoded (playground/share.go in go-slides/go-slides.github.io). A deck with
// a media/, kinds/ or parts/ folder cannot travel in a link; it is offered as a
// .zip, which the playground's file panel imports.
//
// Usage, from the root of the documentation repository:
//
//	go run ./tools/gallery -src ../slides
package main

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Limits of the playground's share link (playground/share.go there).
const (
	maxSource = 1 << 20
	maxShare  = 64 << 10
)

// Entry is one deck of the gallery, as data/gallery.json holds it.
type Entry struct {
	Name        string   `json:"name"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Slides      int      `json:"slides"`
	Source      string   `json:"source"`     // the deck's text, relative to the site root
	Thumbnail   string   `json:"thumbnail"`  // relative to static/
	Playground  string   `json:"playground"` // the link, when the deck is one file that fits
	Zip         string   `json:"zip"`        // relative to static/, when the deck has other files
	Files       []string `json:"files"`      // the files beside deck.md
	Repo        string   `json:"repo"`       // the deck's folder on GitHub
}

func main() {
	var (
		src        = flag.String("src", "", "a checkout of go-slides/slides (its examples/ directory and cmd/slides)")
		out        = flag.String("out", ".", "the documentation repository's root")
		bin        = flag.String("slides", "", "the slides command; by default one is built from -src")
		playground = flag.String("playground", "https://go-slides.github.io/playground/", "the playground's address")
		repo       = flag.String("repo", "https://github.com/go-slides/slides/tree/main/examples", "where the examples live")
	)
	flag.Parse()
	if *src == "" {
		fatal(errors.New("-src is required"))
	}
	if *bin == "" {
		tmp, err := os.MkdirTemp("", "gallery")
		if err != nil {
			fatal(err)
		}
		defer os.RemoveAll(tmp)
		*bin = filepath.Join(tmp, "slides")
		build := exec.Command("go", "build", "-o", *bin, "./cmd/slides")
		build.Dir = *src
		build.Env = append(os.Environ(), "GOWORK=off")
		if b, err := build.CombinedOutput(); err != nil {
			fatal(fmt.Errorf("building the slides command: %w\n%s", err, b))
		}
	}
	entries, err := Build(*src, *out, *bin, *playground, *repo)
	if err != nil {
		fatal(err)
	}
	fmt.Printf("gallery: %d decks\n", len(entries))
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "gallery:", err)
	os.Exit(1)
}

// Build writes the gallery's files under out, from the decks of src/examples.
func Build(src, out, slidesBin, playground, repo string) ([]Entry, error) {
	decks, err := filepath.Glob(filepath.Join(src, "examples", "*", "deck.md"))
	if err != nil {
		return nil, err
	}
	if len(decks) == 0 {
		return nil, fmt.Errorf("no deck in %s", filepath.Join(src, "examples"))
	}
	sortDecks(decks)
	static := filepath.Join(out, "static", "gallery")
	if err := os.RemoveAll(static); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(static, 0o755); err != nil {
		return nil, err
	}
	var entries []Entry
	for _, deck := range decks {
		dir := filepath.Dir(deck)
		name := filepath.Base(dir)
		text, err := os.ReadFile(deck)
		if err != nil {
			return nil, err
		}
		title, desc := Meta(string(text))
		if title == "" || desc == "" {
			return nil, fmt.Errorf("%s: the front matter needs a title and a description", deck)
		}
		files, err := Others(dir)
		if err != nil {
			return nil, err
		}
		e := Entry{
			Name: name, Title: title, Description: desc, Files: files,
			Source:    "static/gallery/" + name + ".md",
			Thumbnail: "gallery/" + name + ".jpg",
			Repo:      repo + "/" + name,
		}
		if err := os.WriteFile(filepath.Join(out, e.Source), text, 0o644); err != nil {
			return nil, err
		}
		n, err := render(slidesBin, deck, filepath.Join(static, name+".jpg"))
		if err != nil {
			return nil, fmt.Errorf("%s: %w", name, err)
		}
		e.Slides = n
		if len(files) == 0 {
			if frag, err := Share(string(text)); err == nil {
				e.Playground = strings.TrimSuffix(playground, "/") + "/#" + frag
			}
		}
		if len(files) > 0 || e.Playground == "" {
			z, err := Zip(dir)
			if err != nil {
				return nil, err
			}
			if err := os.WriteFile(filepath.Join(static, name+".zip"), z, 0o644); err != nil {
				return nil, err
			}
			e.Zip = "gallery/" + name + ".zip"
		}
		entries = append(entries, e)
	}
	b, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(filepath.Join(out, "data"), 0o755); err != nil {
		return nil, err
	}
	return entries, os.WriteFile(filepath.Join(out, "data", "gallery.json"), append(b, '\n'), 0o644)
}

// order is the gallery's order: the themes, then one feature at a time, then
// the complete talk. A deck it does not name follows, by name.
var order = []string{
	"theme-default", "theme-gaia", "theme-uncover", "kinds", "templates", "hcl",
	"math", "images", "code", "emoji-and-fit", "presenting", "talk",
}

func sortDecks(decks []string) {
	rank := func(deck string) int {
		name := filepath.Base(filepath.Dir(deck))
		for i, n := range order {
			if n == name {
				return i
			}
		}
		return len(order)
	}
	sort.SliceStable(decks, func(i, j int) bool {
		if ri, rj := rank(decks[i]), rank(decks[j]); ri != rj {
			return ri < rj
		}
		return decks[i] < decks[j]
	})
}

var slidesRE = regexp.MustCompile(`\((\d+) slides?\)`)

// render writes the first slide of deck as a JPEG, with the slides command,
// and returns the deck's number of slides.
func render(bin, deck, jpg string) (int, error) {
	// --image reports one slide: the count comes from a build of the page.
	page := jpg + ".html"
	defer os.Remove(page)
	var count string
	for _, args := range [][]string{
		{"build", "--allow-local-files", "-o", page, deck},
		{"build", "--allow-local-files", "--image", "jpeg", "--jpeg-quality", "80", "-o", jpg, deck},
	} {
		cmd := exec.Command(bin, args...)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		if err := cmd.Run(); err != nil {
			return 0, fmt.Errorf("slides %s: %w\n%s", args[0], err, stderr.String())
		}
		if count == "" {
			count = stderr.String()
		}
	}
	m := slidesRE.FindStringSubmatch(count)
	if m == nil {
		return 0, fmt.Errorf("slides did not say how many slides: %q", count)
	}
	return strconv.Atoi(m[1])
}

// Meta returns a deck's title and description, from its front matter, in YAML
// ("title: x") or HCL ("title = "x"").
func Meta(deck string) (title, description string) {
	front := frontMatter(deck)
	return field(front, "title"), field(front, "description")
}

func frontMatter(deck string) string {
	lines := strings.Split(deck, "\n")
	if len(lines) == 0 || !strings.HasPrefix(lines[0], "---") {
		return ""
	}
	for i := 1; i < len(lines); i++ {
		if strings.TrimRight(lines[i], " \r") == "---" {
			return strings.Join(lines[1:i], "\n")
		}
	}
	return ""
}

func field(front, key string) string {
	re := regexp.MustCompile(`(?m)^` + key + `\s*[:=]\s*(.+?)\s*$`)
	m := re.FindStringSubmatch(front)
	if m == nil {
		return ""
	}
	v := m[1]
	if u, err := strconv.Unquote(v); err == nil {
		return u
	}
	return strings.Trim(v, `'"`)
}

// Others lists the files of a deck's folder besides deck.md, slash-separated
// and sorted.
func Others(dir string) ([]string, error) {
	var out []string
	err := filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(dir, p)
		if d.IsDir() || rel == "deck.md" || strings.HasPrefix(d.Name(), ".") {
			return nil
		}
		out = append(out, filepath.ToSlash(rel))
		return nil
	})
	sort.Strings(out)
	return out, err
}

// Share returns the playground's fragment for a deck, without the "#".
func Share(deck string) (string, error) {
	if len(deck) > maxSource {
		return "", errors.New("the deck is larger than the playground takes")
	}
	var buf bytes.Buffer
	w, _ := flate.NewWriter(&buf, flate.BestCompression)
	_, _ = w.Write([]byte(deck))
	_ = w.Close()
	s := "deck=" + base64.RawURLEncoding.EncodeToString(buf.Bytes())
	if len(s) > maxShare {
		return "", errors.New("the deck is too large to travel in a link")
	}
	return s, nil
}

// Zip archives a deck's folder, deck.md at its top and the rest where it is,
// the files sorted and their times fixed, so that the same decks give the same
// bytes.
func Zip(dir string) ([]byte, error) {
	files, err := Others(dir)
	if err != nil {
		return nil, err
	}
	names := append([]string{"deck.md"}, files...)
	var buf bytes.Buffer
	w := zip.NewWriter(&buf)
	when := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, n := range names {
		b, err := os.ReadFile(filepath.Join(dir, filepath.FromSlash(n)))
		if err != nil {
			return nil, err
		}
		h := &zip.FileHeader{Name: n, Method: zip.Deflate, Modified: when}
		h.SetMode(0o644)
		f, err := w.CreateHeader(h)
		if err != nil {
			return nil, err
		}
		if _, err := f.Write(b); err != nil {
			return nil, err
		}
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
