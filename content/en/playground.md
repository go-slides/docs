---
title: "Playground"
weight: 15
description: "Write a deck in your browser and watch the slides follow: go-slides compiled to WebAssembly, with completion, help and examples."
tags: [playground, browser]
---

**[Open the playground →](https://go-slides.github.io/playground/)**

The playground is go-slides itself, v0.20.0, compiled to WebAssembly and
running in your browser. Type a deck on the left; the slides follow on the
right. Nothing is sent to a server: not the deck, not a share link.

The playground's own interface (ribbon, help, messages) speaks the language of
the page: open `/fr/playground/` for French, and so on for the twelve languages
of this site. The language selector of the page changes it and your choice is
kept in this browser; a link that names a language always wins. The example
decks stay in English. Ukrainian, Russian, Chinese, Japanese and Korean are
drawn with a subset of Noto Sans (0.06 to 2.3 MiB) that the page downloads for
those languages only.

## What it has

- **Live preview** of every slide, rendered off the page's thread (a Web
  Worker), so typing never waits for it. Click a slide to move the editor
  to its line; move the caret and the preview scrolls to its slide.
- **Slideshow**: the presentation itself, with the keyboard, the
  on-screen controller and the transitions.
- **Presenter view**: **Present** (or `p` in the slideshow) opens it in a
  second window: the current slide, the next step, the speaker notes (the
  deck's `<!-- comments -->`), a clock and a timer. The arrow keys, Space,
  Home and End work in either window, and each follows the other, also
  when you edit the deck. If the window does not open, allow pop-ups for
  the site. The overview view (`o`) is in the downloaded page only.
- **Completion** as you type: directive keys (spot `_` ones in comments),
  the values of `kind`, `theme`, `size`, `paginate`, and template functions
  and data fields inside `{{ }}`.
- **Help** for what is under the caret: a directive, a kind (what it sets),
  a template function. The help pane also links to these pages. Since v0.16.0
  the directive help follows the interface language, in 11 languages.
- **Problems** on their line: template, YAML and HCL errors are highlighted
  in the editor and listed under it, and the preview keeps the last deck
  that rendered.
- **Examples**: a welcome deck, templates that generate slides, slide kinds,
  an HCL front matter, themes and spot directives, and math, images and
  backgrounds (`![bg]`, a split `![bg left:40%]`, filters).
- **Files**: the panel left of the editor is the deck's project, laid out
  the way the [command]({{< relref "cli" >}}) and the
  [VS Code extension]({{< relref "vscode" >}}) read it: `deck.md`,
  `slides.yaml`, `themes/`, `kinds/`, `media/`.
  - Click a file to edit it. The editor colours it by its language
    (Markdown, YAML, HCL, JSON, CSS) and marks its problems.
  - Right-click (or Shift+F10) for New File, New Folder, Rename (F2),
    Delete (asks first) and New theme / New kind. Drag a file onto a folder
    to move it. The keyboard follows the WAI-ARIA tree pattern, type-ahead
    included.
  - **Project ▸ New theme** makes `themes/my-gaia.css` (or from default,
    uncover): `@import` of the built-in theme and the rules you add. It
    joins the Theme list. **New kind** and **New data file** start from
    templates too. Each is listed in `slides.yaml`, so `slides build deck.md`
    finds it.
  - **Project ▸ Download project (.zip)** saves it all. Unzipped,
    `slides build deck.md` builds it. **Open project** takes a `.zip` or a
    folder back. An archive that leaves its folder (`..`, absolute paths),
    holds more than 256 files or 8 MiB in a file, or inflates more than 100
    times is refused.
  - Everything is kept in this browser with the draft. A project holds up
    to 64 MiB, and the line under the files says how much is used.
- **Media**: drop or paste an image, a video or a sound on the editor and it
  goes into `media/`, and the deck names it: `![](media/photo.png)`. Drop
  files on a folder of the panel to put them there. The deck stays small and
  readable, and nothing is uploaded. The preview's sandbox loads them
  without being opened up (the page hands it the bytes). **Save as ▸ HTML
  page** puts them in the page. `https:` images work too, but the browser
  fetches them from their server, which then sees your address. The
  preview sends no referrer.
- **Resizable panes**: the file panel, the editor, the slide thumbnails, the
  preview and the help are separated by bars you can drag (the pointer
  becomes a resize cursor over one). The pane beside the bar takes the room,
  the others keep their size. Double-click a bar, or press Enter on it, to
  pull the pane shut and bring it back. From the keyboard, F6 moves between
  the bars, the arrow keys move one by 16 px, Home and End take the pane
  before it to its smallest and largest, and Escape returns to the editor.
  The View tab still switches panes off. The sizes you set are kept in this
  browser, with the draft; the default sizes (190, 168 and 280 px, the editor
  and the preview sharing the rest 6 to 5) are never stored.
- **Format**: the Format tab edits the Markdown for you, and each press is
  one step of Undo.
  - **Text**: Bold, Italic, Strikethrough and Inline code work on the selection
    or the word at the caret and toggle (Ctrl+B, Ctrl+I, Ctrl+Shift+X and
    Ctrl+E; Cmd on a Mac). A button shows pressed when the caret is inside
    its marks. Markdown has no underline: Underline (Ctrl+U) wraps the text in
    `<u>…</u>`, a tag Marp's default HTML allow-list keeps, so it works with
    raw HTML off.
  - **Paragraph**: Heading 1 to 3 (Ctrl+Alt+1..3), bullets (Ctrl+Shift+8),
    numbering (Ctrl+Shift+7) and a quote work on the selected lines and toggle.
    The horizontal rule is the `<hr>` tag: `---`, `***` and `___` would each
    start a new slide in Marp.
  - **Table**: a grid picks the columns and rows. At the caret in a table you
    add or remove a row or a column and align the column (`:--`, `:-:`,
    `--:`); the table is rewritten aligned after each edit.
  - **Slide text** left, centre or right writes a `<style scoped>` block into
    the current slide. Marp has no alignment syntax, and a `style` attribute
    does not pass the default HTML allow-list, but the scoped block does, so it
    works with raw HTML off and shows in the preview, the HTML page and the PDF.
  - **Link** (Ctrl+K) writes `[text](url)` with the url selected.
- **Highlighting** of the front matter's keys at any depth, in YAML (also
  in `{a: 1}` mappings) and in HCL.
- **Theme** picker: default, gaia, uncover, the project's themes
  (`themes/*.css`) and the themes the deck defines (`/* @theme name */` in a
  `<style>`). It writes `theme:` into the front
  matter, YAML or HCL, so a shared link keeps it, and it follows the front
  matter when you edit it.
- The preview **keeps its place** as you type: only the slides that
  changed are replaced, and a change of theme reopens the preview on the
  same slide.
- **Share link**: the deck travels in the link itself (in its `#` part,
  compressed), so opening the link opens the deck. Links carry up to
  64 KiB compressed, and the deck only: when the project holds more
  (themes, kinds, media), Share says so and offers the `.zip`.
- **Save as**:
  - the HTML page, the same as `slides deck.md` writes;
  - a **PDF** (vector, its text selectable, with the outline);
  - an editable **PowerPoint** or **OpenDocument** deck.

  They are made in your browser, by the same code as `slides --pdf`,
  `--pptx` and `--odp`: the 100-slide benchmark deck's PDF is
  byte-identical to the command's. The first export downloads the export
  engine (about 5 MiB, once; your browser keeps it until the next update),
  with a progress line, and compiles it as it arrives. Then it downloads only
  the fonts your slides draw. On a 3 MB/s link the first PDF of the welcome
  deck is ready in under 3 seconds. The 100-slide deck then takes about
  3 seconds. An
  `https:` image is included only if its server allows it (CORS); `data:`
  images always are.
- **Your draft is kept**:
  - The deck is saved in this browser as you type (IndexedDB, for this site
    only) and comes back when you reload. Nothing is sent anywhere. The
    project's text files (themes, kinds, data, `slides.yaml`) are kept the
    same way, even when you close the tab right after typing; media are
    stored as they are dropped.
  - A shared link opens its deck and keeps your draft: **Draft ▸ Restore my
    draft** brings it back.
  - An example, **Draft ▸ Reset to the example** (which asks first) and
    another tab never erase it either: **Draft ▸ Restore replaced** swaps it
    back.
  - **Draft ▸ Forget my drafts** erases everything stored.

Every export renders the project the preview shows: its `slides.yaml`,
themes, kinds, data files and media, the media inside the document.

Images as files (slides `--images`) are not made in the browser: use the
[command]({{< relref "cli" >}}) or the
[VS Code extension]({{< relref "vscode" >}}).

## Theme editor {#theme-editor}

**Home ▸ Look ▸ Theme editor** opens a pane that edits one of the project's
themes with controls, and the preview follows. The theme is kept as a package
folder, `theme-src/<name>/`; the sheet the preview and the exports read,
`themes/<name>.css`, is compiled from it, so what you see is what you export. A
theme package you import (**Import…** in the pane, **Add media…**, or a drop)
keeps its files there, so you can edit it.

- **Palette**: the background, text, heading, accent and link colours. Pick a
  colour or type it (`#1a2b3c`, `rgb()`, `hsl()` or a name).
- **Typography**: the body, heading and code fonts, with their sizes and
  weights. The fonts are the generic and common ones, and the fonts you add
  (**Add font…**: `.woff2`, `.woff`, `.ttf` or `.otf`).
- **Logo**: an image (`.svg`, `.png`), where it goes (a corner, or the whole
  background) and its size. **Page number and footer**: where the page number
  goes, and the sizes and alignment.
- **Details**: what `theme.hcl` says: version, description, authors, licence,
  homepage, the base theme (`extends`) and the go-slides version it needs.
- **Contrast (WCAG)**: for each text on its background, the ratio and its
  level, `AA` or `AAA`, or "Fail", live. Headings count as large text (3:1
  passes AA; normal text needs 4.5:1).

**Edit CSS** shows the theme's sheet in the editor. The controls write only
inside a block fenced by two `gsta:visual` comments, and read it back when you
edit its values by hand. What you write outside the block is never touched. If
you add something inside the block that the controls cannot show, they stop
writing and the pane says so; **Restore the block** rewrites it from the values
it holds and nothing else.

**Import…** opens a `.gsta.xz` or a loose `.css`. A package is checked as
`slides theme verify` checks it, with the same limits: a package that is not
safe (a path that leaves it, a link, a script in an SVG, a remote `url()`, a
checksum that does not match) is refused with its `theme.*` code, in Problems,
and nothing changes. In a loose sheet an `@import "gaia";` becomes the base
theme (`extends`), because a package names its base there. A font you add has
the licence `LicenseRef-unspecified` until you say it under Details.

**Export .gsta.xz** builds the package in your browser and reads it back as
`slides theme verify` does, so the file passes it; **Export CSS** downloads the
compiled sheet. Nothing is uploaded. Every edit is a step of **Undo** (edits of
one control in a row are one step) and keeps what it replaced as a version;
**Save version** keeps the theme and its version, by name, in the Versions
pane.

## Undo, Redo and versions

**Undo** and **Redo** are on the Home tab and on Ctrl+Z, Ctrl+Shift+Z or
Ctrl+Y (Cmd on a Mac). They take back the last edit of the file shown (each
file keeps its own history), and then the last operation on the whole
project: an example picked, Reset, a .zip or a folder opened, a draft
restored, a file or folder deleted or renamed, a version restored. Redo brings
back everything as it was when Undo was pressed.

**Versions** keep the whole project in this browser, with nothing sent
anywhere. One is kept a little after you stop typing (at most every 2
minutes), before each operation on the whole project, and at each export and
share link. **Save version…** keeps one under a name. **Home ▸ Versions**
opens the Versions pane:

- **Preview** shows a version in the preview.
- **Compare with now** shows each file's changes line by line, inline or side
  by side; F7 and Shift+F7 move between changes.
- **Restore** puts a version back; Undo takes the restore back.
- **Delete** removes a version.

What is kept: the 20 newest automatic versions, then one an hour for a day and
one a day for a month. Named versions stay until you delete them (50 at most).
All versions together take at most 96 MiB, and files that did not change are
stored once. When the browser's storage is full, the oldest automatic versions
go first.

## GitHub {#github}

**Home ▸ History ▸ GitHub** opens a pane that loads a deck's folder from a
GitHub repository and proposes your changes back as a pull request. Fill in:

- **Repository (owner/name)**;
- **Branch (empty: the default)**;
- **Folder (empty: the root)**: the folder must hold a `deck.md`.

Above these fields the pane shows the **git frame**: a small panel in which you
choose the forge, type the **token** (masked), tick **Remember on this device**,
and press **Connect** (or **Sign in with GitLab**) or **Disconnect**. The frame
is a page of another site, `go-slides-usercontent.github.io`, embedded in the
playground: **the playground page never sees your token**.

The token is a fine-grained personal access token
([create one](https://github.com/settings/personal-access-tokens/new)) limited
to the repositories you choose, with the permissions Contents read/write and
Pull requests read/write (Metadata read is implied). Then:

- **Load** replaces the project with the folder from GitHub. Undo takes it
  back, and the project as it was is kept as a version "From GitHub".
- **Propose changes…** asks for a title, creates a new branch named
  `slides/<title>-<date-time>` and opens a pull request. It never writes to an
  existing branch or to the default branch. **Open the pull request** takes you
  to it.
- **Check for changes** compares the newer version on GitHub with yours in the
  Compare view. **Take theirs** loads it.

No proxy is used for GitHub: the frame talks to it directly.

The token is typed in the frame and stays on its site. The playground page
cannot read it (the frame is another origin, with storage of its own), no
command the page can send returns it, and the page's Content Security Policy
allows it no connection to GitHub or GitLab at all. The frame sends the token
only to `api.github.com` (or `gitlab.com`). It is never saved in the draft, the
versions, the share links or the exports, and it never reaches the slides'
preview or the render worker. **Remember on this device** stores it in the
frame's own localStorage, on `go-slides-usercontent.github.io`, not in the
playground's storage, until you press **Disconnect** (or Forget my drafts).

**If you used an earlier version:** a token it kept in the playground page's own
storage is removed the first time you open this version, and the pane tells you.
Type it again in the frame (and tick Remember if you want it kept).

One risk remains: a script injected into the playground page could still send
the frame's commands (read the loaded folder, propose a branch and a pull
request) as you, but it cannot read the token. Prefer a token limited to the
repositories you need. For teams, the self-hosted server mode (coming) keeps
tokens on the server.

### GitLab.com {#gitlab}

The pane has a forge list: **GitHub** and **GitLab.com**. For GitLab, give
either a personal access token with the scope `api`, or press **Sign in with
GitLab** (OAuth authorization code with PKCE, no secret). Sign-in opens a small
window of the frame's site and returns to it; the playground page you are
working in is not navigated away. Its tokens stay in the frame's memory only
(never remembered on the device, even with **Remember on this device**), refresh
themselves before they expire, and **Disconnect** revokes them at GitLab. A
sign-in not started by the frame, or a reused link, is refused. Sign-in is
available once the site's GitLab application is registered; until then, use a
personal access token.

Requests go only to `api.github.com` and `gitlab.com`, and only from the frame:
the playground page makes none, and sends no referrer.

Why `api`? GitLab's
[scope table](https://docs.gitlab.com/security/tokens/access_token_scopes/)
says `write_repository` "Does not support API authentication": it is for git
over HTTP only, which GitLab serves without CORS, so a browser cannot use it.
Proposals use the REST API, which needs `api`, "complete read and write access
to the API". A stolen `api` token can do anything its user can do through the
API, on every project that user reaches. Prefer a project or group access
token limited to the deck repositories, or a dedicated GitLab account that is
only a member of the deck projects.

## Security

The deck is shown in an iframe sandboxed without the page's origin: it can
reach neither the playground nor its storage, submit no form and navigate
nothing. Its page carries a Content Security Policy that runs only its own
scripts. Raw HTML in the deck follows Marp's allowlist, as `slides` does by
default. A render that runs for more than 5 seconds (a template that loops)
is stopped.

The presenter window keeps that sandbox. It is opened by the playground
itself and shows the deck in an iframe sandboxed the same way. The two
iframes cannot share storage, so the playground relays the slide number
between them in messages. No frame that runs the deck gets the
playground's origin.

The project's media reach the sandboxed frames as bytes in a message, and
each frame makes its own `blob:` URLs of them, in its own origin. A `blob:`
URL made by the playground does not load there, and the frames' policy
gains only `blob:` for media and fonts.

The playground can be served from your own machine, behind a sign-in:
see [Self-hosting]({{< relref "self-hosting" >}}).

## Size and speed

The module is 8.8 MiB as GitHub Pages serves it (gzip; Pages does not
serve brotli). It is downloaded once with a progress bar, then cached by the
browser until the next deploy. On an Apple M4 Max with Chrome 155, measured
against the live page, the editor is ready 2.1 s after the page opens, 1.5 s
of which is the download, and the first slides appear at 2.4 s. A change to
the 100-slide benchmark deck reaches the preview in about 30 ms: go-slides
renders only the slides that changed, and only those are replaced in the
preview.
