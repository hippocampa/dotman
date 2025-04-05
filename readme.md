# Dotman: My Personal Dotfiles Manager

```
  _____        _
 |  __ \      | |
 | |  | | ___ | |_ _ __ ___   __ _ _ __
 | |  | |/ _ \| __| '_ ` _ \ / _` | '_ \
 | |__| | (_) | |_| | | | | | (_| | | | |
 |_____/ \___/ \__|_| |_| |_|\__,_|_| |_|


```

It was originally named _dotshadow_, since I wanted to try some different approaches to manage dotfiles… but those ideas didn’t really work out.

But hey, here it is — my own take on a dotfile manager, now called Dotman.

> Dotman is highly inspired by
> [GNU Stow](https://www.gnu.org/software/stow/).
> But it works in reverse — and maybe even better (I hope).

## What's working so far

This is an early release, so Dotman is still fresh and evolving. Things might break, and it definitely needs more testing and refinement.

### Implemented

- `init` - initialize dotman stores
- `track <file>` - copy a file into the store and replace it with symlink
- `restore <file|all>` - restore a broken/missing link

### Coming Soon

- `sync` - track file hashes and detect changes
- `status` - show which files are dirty, missing or broken
- `untrack` - undo tracking (optionally for restoring file)

## How to use

Dotman is designed to be simple and intuitive — no reverse folder magic or clunky layout rules like GNU Stow. Just track files in-place and let Dotman do the rest.

1. Init

   ```bash
   dotman init
   ```

   Creates a .dotfiles/ store in your home directory.

2. Track a file

   ```bash
   dotman track ~/Desktop/test.md
   ```

   What it does:

   1. Copies the file to .dotfiles/Desktop/test.md
   2. Replaces the original with a symlink
   3. Stores a mapping in .dotfiles/index.json

   So now:

   ~/Desktop/test.md → ~/.dotfiles/Desktop/test.md

3. Restore a file

   Let’s say you deleted a file (but still have it in your dotman store):

   ```bash
   dotman restore ~/Desktop/test.md
   ```

   Or restore everything at once:

   ```bash
   dotman restore --all
   ```

   This will re-create the symlinks using whatever is stored in `.dotfiles/index.json.`

## Installation

Download the latest binary from the [Releases](https://github.com/hippocampa/dotman/releases) page.

```bash
chmod +x dotman
mv dotman /usr/local/bin
```

## Philosophy

Dotman **does not** reverse your folder structure.

You track files from where they are (usually in `$HOME`), and Dotman keeps a copy of them in `.dotfiles/`. It handles linking, indexing, and (soon) syncing — while staying out of your way.

No weirdness. No reverse symlinks. No need to move all your dotfiles into one mega repo manually.

## Contributing

This project is still evolving, and any feedback, ideas, or issues are very welcome.

Feel free to open an issue or just yell at me nicely.

## Credits

- [Golang](https://go.dev/)
- [Urfavecli](https://github.com/urfave/cli)

Made with love and a lot of `ln -s` by [me](https://github.com/hippocampa).
