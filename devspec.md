# 📄 Developer Brief: `dotman` — A Modern Dotfiles Manager

---

## 📌 Project Goals

**Primary Goal:**  
Build a dotfiles manager that tracks and manages configuration files directly from the user's `$HOME` directory — with **normal-direction symlinks**, **state awareness**, and optional future versioning.

---

## ✨ Core Philosophy

- ✅ Track dotfiles **in-place** (from `$HOME`)
- ✅ Use **forward symlinks** (`$HOME/.zshrc` → `.dotman/files/zsh/.zshrc`)
- ✅ Maintain a **clear internal state** for each file
- ❌ Avoid reverse symlinks (like GNU Stow)
- ⚙️ Focus on **ergonomic CLI usage**, not Git-like workflows

---

## 📁 Core Folder Layout (Internal Store)

```text
.dotman/
├── files/                # Stored copies of tracked files
│   └── zsh/.zshrc
├── index.json            # Tracks managed files and metadata
└── config.json           # Optional global config
```

---

## 🔧 Core Commands (CLI Interface)

| Command                 | Behavior                                               |
| ----------------------- | ------------------------------------------------------ |
| `dotman init`           | Create `.dotman/`, initialize structure                |
| `dotman track <file>`   | Copy file to `.dotman/files/...`, replace with symlink |
| `dotman untrack <file>` | Remove symlink, restore original from backup           |
| `dotman status`         | Show tracked files, missing links, or changes          |
| `dotman restore <file>` | Replace symlink with saved file                        |
| `dotman list`           | Show all tracked files and their status                |
| `dotman clean`          | Remove broken links or files no longer in use          |

---

## 🛠 Core Behavior by Command

### `init`

- Create `.dotman/`
- Set up subfolders: `files/`, `index.json`
- Optionally create `config.json`

---

### `track`

1. Ensure the file exists
2. Compute relative path (e.g., `~/.zshrc` → `zsh/.zshrc`)
3. Copy file into `.dotman/files/...`
4. Create symlink: `$HOME/.zshrc` → `.dotman/files/zsh/.zshrc`
5. Record in `index.json`:
   ```json
   {
     "source": "/home/user/.zshrc",
     "stored": ".dotman/files/zsh/.zshrc",
     "status": "linked"
   }
   ```

---

### `untrack`

1. Remove symlink from `$HOME/.zshrc`
2. Replace it with the original file from `.dotman/files/...` (optional restore)
3. Remove entry from `index.json`

---

### `status`

- For each entry in `index.json`:
  - Check if symlink exists and is valid
  - Check if file has been modified (optional hash compare)
  - Display:
    - ✅ linked
    - ⚠️ changed
    - ❌ missing
    - 🔗 not a symlink

---

### `restore`

- Copy stored file back to `$HOME`
- Replace symlink if broken
- Optionally prompt before overwriting

---

### `clean`

- Remove any tracked files that are no longer in use
- Identify and warn about broken symlinks

---

## 📚 Data Structures (Examples)

### `index.json`

```json
{
  "tracked": [
    {
      "source": "/home/user/.zshrc",
      "stored": ".dotman/files/zsh/.zshrc",
      "status": "linked",
      "lastTrackedHash": "abc123"
    }
  ]
}
```

### `config.json` (optional)

```json
{
  "defaultStorePath": "/home/user/.dotman/files",
  "exclude": [".cache", ".local"]
}
```

---

## 🛤️ Build Roadmap (Phase-Based)

### ✅ Phase 1: Core Tracking

- [ ] `init` to set up the store
- [ ] `track` to copy file + symlink + update index
- [ ] `untrack` to undo
- [ ] `list` to display all tracked files
- [ ] Basic `index.json` read/write

---

### ✅ Phase 2: State Awareness

- [ ] `status` command
- [ ] Check for broken symlinks
- [ ] Hash comparison for modified files
- [ ] `restore` for recovery

---

### 🧪 Phase 3: UX Polish

- [ ] Add `--verbose`, `--dry-run`, `--force` flags
- [ ] Optional colorized CLI output
- [ ] Optional TUI (e.g. using `bubbletea` or similar)

---

### 🚀 Phase 4: (Optional Features)

- [ ] Git integration (`dotman push/pull`)
- [ ] Snapshots (like Git commits, but simpler)
- [ ] Remote sync (SFTP, rsync, etc.)
- [ ] Templates or profiles (e.g., work vs. personal)

---

## 🧠 Key Design Decisions

- Use `$HOME` as source of truth
- Avoid copying everything — track only what you need
- Default to symlinking, but allow copying as fallback
- Prefer simplicity over flexibility (you’re not building Git)

---

## 🧑‍💻 Dev Notes

- Language: Go
- Config format: JSON (can evolve to TOML/YAML if needed)
- Dependencies: minimal — avoid big libraries unless needed
