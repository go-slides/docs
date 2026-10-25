package main

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"encoding/base64"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestMeta(t *testing.T) {
	for _, c := range []struct{ name, deck, title, desc string }{
		{"yaml", "---\nslides: true\ntitle: The default theme\ndescription: Clean and light.\n---\n\n# x\n", "The default theme", "Clean and light."},
		{"quoted", "---\ntitle: \"A: b\"\ndescription: 'It''s fine'\n---\n", "A: b", "It''s fine"},
		{"hcl", "---hcl\ntitle    = \"HCL front matter\"\ndescription = \"In HCL2.\"\n---\n", "HCL front matter", "In HCL2."},
		{"none", "# no front matter\ntitle: not this\n", "", ""},
		{"body is not the front matter", "---\ntheme: x\n---\ntitle: in the body\n", "", ""},
	} {
		t.Run(c.name, func(t *testing.T) {
			title, desc := Meta(c.deck)
			if title != c.title || desc != c.desc {
				t.Errorf("got %q, %q; want %q, %q", title, desc, c.title, c.desc)
			}
		})
	}
}

// decode is the playground's DecodeShare, to prove that Share is what it reads.
func decode(t *testing.T, frag string) string {
	t.Helper()
	data, ok := strings.CutPrefix(frag, "deck=")
	if !ok {
		t.Fatalf("no deck= prefix: %.20s", frag)
	}
	raw, err := base64.RawURLEncoding.DecodeString(data)
	if err != nil {
		t.Fatal(err)
	}
	b, err := io.ReadAll(flate.NewReader(bytes.NewReader(raw)))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func TestShareRoundTrip(t *testing.T) {
	deck := "---\ntitle: é\n---\n\n# Hello $E=mc^2$ 😄\n"
	frag, err := Share(deck)
	if err != nil {
		t.Fatal(err)
	}
	if got := decode(t, frag); got != deck {
		t.Errorf("round trip: %q", got)
	}
	if strings.ContainsAny(strings.TrimPrefix(frag, "deck="), "+/=# ") {
		t.Errorf("not URL-safe: %q", frag)
	}
}

func TestShareRefusesWhatCannotTravel(t *testing.T) {
	// Random-looking text does not deflate: it is the case a link cannot carry.
	var sb strings.Builder
	x := uint32(1)
	for sb.Len() < 200<<10 {
		x = x*1664525 + 1013904223
		sb.WriteByte(byte('a' + (x>>24)%26))
	}
	if _, err := Share(sb.String()); err == nil {
		t.Error("a deck too large for a link was shared")
	}
	if _, err := Share(strings.Repeat("x", maxSource+1)); err == nil {
		t.Error("a deck past the playground's bound was shared")
	}
}

func fixture(t *testing.T) string {
	t.Helper()
	src := t.TempDir()
	write := func(p, s string) {
		p = filepath.Join(src, "examples", filepath.FromSlash(p))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(s), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("one/deck.md", "---\ntitle: One\ndescription: A single file.\n---\n\n# One\n")
	write("two/deck.md", "---\ntitle: Two\ndescription: Two files.\n---\n\n![bg](media/a.svg)\n")
	write("two/media/a.svg", "<svg xmlns='http://www.w3.org/2000/svg'/>")
	return src
}

func TestOthers(t *testing.T) {
	src := fixture(t)
	got, err := Others(filepath.Join(src, "examples", "two"))
	if err != nil {
		t.Fatal(err)
	}
	if want := []string{"media/a.svg"}; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	got, _ = Others(filepath.Join(src, "examples", "one"))
	if len(got) != 0 {
		t.Errorf("a single file has no others: %v", got)
	}
}

func TestZipIsStableAndImportable(t *testing.T) {
	dir := filepath.Join(fixture(t), "examples", "two")
	a, err := Zip(dir)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := Zip(dir)
	if !bytes.Equal(a, b) {
		t.Error("two archives of the same folder differ")
	}
	r, err := zip.NewReader(bytes.NewReader(a), int64(len(a)))
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	for _, f := range r.File {
		names = append(names, f.Name)
	}
	if want := []string{"deck.md", "media/a.svg"}; !reflect.DeepEqual(names, want) {
		t.Errorf("entries %v, want %v", names, want)
	}
}

// TestMain lets this test binary stand in for the slides command: a script
// would not be portable, so TestBuild runs the binary itself, with an
// environment variable that makes it act as `slides build ... -o OUT DECK`.
func TestMain(m *testing.M) {
	if os.Getenv("GALLERY_FAKE_SLIDES") == "1" {
		out := ""
		for i, a := range os.Args {
			if a == "-o" {
				out = os.Args[i+1]
			}
		}
		_ = os.WriteFile(out, []byte("\xff\xd8jpeg"), 0o644)
		os.Stderr.WriteString("slides: deck => " + out + " (3 slides)\n")
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func TestBuild(t *testing.T) {
	src := fixture(t)
	out := t.TempDir()
	self, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	t.Setenv("GALLERY_FAKE_SLIDES", "1") // the child process inherits it
	entries, err := Build(src, out, self, "https://example.test/playground/", "https://example.test/repo")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 2 {
		t.Fatalf("%d entries", len(entries))
	}
	one, two := entries[0], entries[1]
	if one.Name != "one" || two.Name != "two" {
		t.Fatalf("order: %s, %s", one.Name, two.Name)
	}
	if !strings.HasPrefix(one.Playground, "https://example.test/playground/#deck=") || one.Zip != "" {
		t.Errorf("a single file is shared by link: %+v", one)
	}
	if two.Playground != "" || two.Zip != "gallery/two.zip" {
		t.Errorf("two files go as a zip: %+v", two)
	}
	if one.Slides != 3 || one.Title != "One" || one.Repo != "https://example.test/repo/one" {
		t.Errorf("entry: %+v", one)
	}
	for _, p := range []string{"data/gallery.json", "static/gallery/one.jpg", "static/gallery/one.md", "static/gallery/two.zip"} {
		if _, err := os.Stat(filepath.Join(out, p)); err != nil {
			t.Error(err)
		}
	}
	if _, err := os.Stat(filepath.Join(out, "static/gallery/one.zip")); err == nil {
		t.Error("a single file needs no zip")
	}
}

func TestSortDecks(t *testing.T) {
	d := func(n string) string { return filepath.Join("examples", n, "deck.md") }
	got := []string{d("zeta"), d("talk"), d("alpha"), d("theme-default"), d("kinds")}
	sortDecks(got)
	want := []string{d("theme-default"), d("kinds"), d("talk"), d("alpha"), d("zeta")}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
}

func TestBuildRefusesADeckWithoutMetadata(t *testing.T) {
	src := t.TempDir()
	p := filepath.Join(src, "examples", "x", "deck.md")
	_ = os.MkdirAll(filepath.Dir(p), 0o755)
	_ = os.WriteFile(p, []byte("# no front matter\n"), 0o644)
	if _, err := Build(src, t.TempDir(), "/nonexistent", "", ""); err == nil || !strings.Contains(err.Error(), "title and a description") {
		t.Errorf("got %v", err)
	}
}
