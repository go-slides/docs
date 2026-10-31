package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func run(t *testing.T, dir string, args ...string) {
	t.Helper()
	c := exec.Command("git", append([]string{"-C", dir, "-c", "user.name=t", "-c", "user.email=t@t", "-c", "commit.gpgsign=false"}, args...)...)
	if out, err := c.CombinedOutput(); err != nil {
		t.Fatalf("git %v: %v\n%s", args, err, out)
	}
}

// site builds a repository root whose themes/hextra is a clone with the tags
// v1, v2 and v3 on three commits, checked out at checkout.
func site(t *testing.T, pin, checkout string) string {
	t.Helper()
	root := t.TempDir()
	sub := filepath.Join(root, "themes", "hextra")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	run(t, sub, "init", "-q")
	for _, v := range []string{"v1", "v2", "v3"} {
		if err := os.WriteFile(filepath.Join(sub, "f"), []byte(v), 0o644); err != nil {
			t.Fatal(err)
		}
		run(t, sub, "add", "f")
		run(t, sub, "commit", "-q", "-m", v)
		run(t, sub, "tag", v)
	}
	run(t, sub, "checkout", "-q", checkout)
	if pin != "" {
		if err := os.WriteFile(filepath.Join(root, ".hextra-version"), []byte(pin+"\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

func TestCheck(t *testing.T) {
	cases := []struct {
		name, pin, at string
		wantErr       string
	}{
		{"at the pin", "v2", "v2", ""},
		{"after the pin", "v2", "v3", ""},
		{"BEFORE the pin: the stale-branch downgrade", "v2", "v1", "OLDER"},
		{"unknown tag", "v9", "v3", "not in"},
		{"two words in the pin", "v2 v3", "v3", "one tag"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			root := site(t, c.pin, c.at)
			msg, err := Check(root, ".hextra-version", "themes/hextra")
			if c.wantErr == "" {
				if err != nil || !strings.Contains(msg, "not before") {
					t.Fatalf("got %q, %v", msg, err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), c.wantErr) {
				t.Fatalf("want an error with %q, got %q, %v", c.wantErr, msg, err)
			}
		})
	}
}

func TestCheckNoPinFile(t *testing.T) {
	root := site(t, "", "v3")
	if _, err := Check(root, ".hextra-version", "themes/hextra"); err == nil {
		t.Fatal("a missing pin file must fail")
	}
}

func TestCheckNoSubmodule(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, ".hextra-version"), []byte("v1\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Check(root, ".hextra-version", "themes/hextra"); err == nil {
		t.Fatal("a missing submodule must fail")
	}
}

func TestRootCmd(t *testing.T) {
	root := site(t, "v2", "v1")
	cmd := RootCmd()
	var out, errb bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&errb)
	cmd.SetArgs([]string{"--root", root})
	if err := cmd.Execute(); err == nil || !strings.Contains(errb.String(), "OLDER") {
		t.Fatalf("err %v, stderr %q", err, errb.String())
	}
	root = site(t, "v2", "v3")
	cmd = RootCmd()
	out.Reset()
	cmd.SetOut(&out)
	cmd.SetArgs([]string{"--root", root})
	if err := cmd.Execute(); err != nil || !strings.Contains(out.String(), "not before") {
		t.Fatalf("err %v, stdout %q", err, out.String())
	}
}

func TestMainExits(t *testing.T) {
	var code int
	osExit = func(c int) { code = c }
	defer func() { osExit = os.Exit }()
	old := os.Args
	defer func() { os.Args = old }()
	os.Args = []string{"hextrapin", "--root", t.TempDir()}
	main()
	if code != 1 {
		t.Fatalf("exit code %d", code)
	}
}
