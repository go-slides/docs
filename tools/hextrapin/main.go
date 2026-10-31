// Command hextrapin fails when the hextra submodule (themes/hextra) is older
// than the version recorded in .hextra-version.
//
// A branch cut before another one bumped the theme carries the old submodule
// pointer, and a squash merge of it silently moves the theme BACK: docs #57
// undid #56 ten minutes after it, and the language menu got its flags again.
// The pin file is the version the site is known to need; raising it is a
// deliberate edit, so a downgrade cannot slip through a rebase.
//
// Usage, from the repository's root:
//
//	go run ./tools/hextrapin
package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// osExit is replaced by tests.
var osExit = os.Exit

func main() {
	if err := RootCmd().Execute(); err != nil {
		osExit(1)
	}
}

// RootCmd is the command.
func RootCmd() *cobra.Command {
	var root, pin, sub string
	cmd := &cobra.Command{
		Use:           "hextrapin",
		Short:         "fail when the hextra submodule is older than the pinned version",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			msg, err := Check(root, pin, sub)
			if err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), "hextrapin:", err)
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), msg)
			return nil
		},
	}
	cmd.Flags().StringVar(&root, "root", ".", "the repository's root")
	cmd.Flags().StringVar(&pin, "pin", ".hextra-version", "the pin file, relative to the root")
	cmd.Flags().StringVar(&sub, "submodule", "themes/hextra", "the theme's submodule, relative to the root")
	return cmd
}

func git(dir string, args ...string) (string, error) {
	c := exec.Command("git", append([]string{"-C", dir}, args...)...)
	var out, errb bytes.Buffer
	c.Stdout, c.Stderr = &out, &errb
	c.Stdin = nil
	err := c.Run()
	if err != nil {
		return strings.TrimSpace(out.String()), fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(errb.String()))
	}
	return strings.TrimSpace(out.String()), nil
}

// Check returns a message when the submodule is at the pinned version or
// after it, and an error when it is before it, or when the pin cannot be read.
func Check(root, pinFile, subDir string) (string, error) {
	b, err := os.ReadFile(filepath.Join(root, pinFile))
	if err != nil {
		return "", err
	}
	tag := strings.TrimSpace(string(b))
	if tag == "" || strings.ContainsAny(tag, " \t\n") {
		return "", fmt.Errorf("%s must hold one tag, e.g. v0.24.1; it holds %q", pinFile, string(b))
	}
	sub := filepath.Join(root, subDir)
	head, err := git(sub, "rev-parse", "HEAD")
	if err != nil {
		return "", fmt.Errorf("the submodule %s is not checked out: %w", subDir, err)
	}
	want, err := git(sub, "rev-parse", "--verify", "--quiet", tag+"^{commit}")
	if err != nil {
		// A shallow checkout may not hold the tag: ask the remote once.
		if _, ferr := git(sub, "fetch", "--quiet", "--tags", "origin"); ferr == nil {
			want, err = git(sub, "rev-parse", "--verify", "--quiet", tag+"^{commit}")
		}
		if err != nil {
			return "", fmt.Errorf("the pinned tag %s is not in %s: %w", tag, subDir, err)
		}
	}
	_, err = git(sub, "merge-base", "--is-ancestor", want, head)
	if err == nil {
		return fmt.Sprintf("hextrapin: %s is at %s, not before the pinned %s (%s)", subDir, short(head), tag, short(want)), nil
	}
	var ee *exec.ExitError
	if errors.As(unwrapAll(err), &ee) && ee.ExitCode() == 1 {
		return "", fmt.Errorf("%s is at %s, which is OLDER than the pinned %s (%s): a stale branch moved the theme back; update the submodule to %s or later, or raise the pin deliberately", subDir, short(head), tag, short(want), tag)
	}
	return "", err
}

func unwrapAll(err error) error {
	for {
		u := errors.Unwrap(err)
		if u == nil {
			return err
		}
		err = u
	}
}

func short(s string) string {
	if len(s) > 8 {
		return s[:8]
	}
	return s
}
