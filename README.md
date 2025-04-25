# Goit – A Git-like Version Control System in Go 🔀

Goit is a simplified version control system written in **Golang**, designed to mimic how Git works under the hood using **Merkle Trees**, **content-addressable storage**, and core concepts like **blobs**, **trees**, and **commits**.

> 🧠 Git, but written by you. From scratch. In Go.

---

## 🚀 Features

- `goit init` – Initialize a new repository
- `goit hash-object <file>` – Store a file as a blob and return its SHA-1 hash
- `goit add <file>` – Stage files for commit (coming soon)
- `goit write-tree` – Create a Merkle tree from staged files (coming soon)
- `goit commit -m "msg"` – Save a commit snapshot (coming soon)
- `goit cat-file <hash>` – Inspect stored objects (coming soon)

---

## 🧱 Repository Structure

Once you run `goit init`, it creates the following structure:

```
.goit/
├── objects/         # All blob/tree/commit objects stored by SHA-1 hash
│   └── ab/cd1234... # Split into folders by first 2 chars
├── refs/            # Branch references (e.g., refs/main)
│   └── main         # Points to latest commit hash on 'main'
└── HEAD             # Points to the current ref (e.g., refs/main)
```

Each object stored inside `.goit/objects/` is:
- Prefixed with a header like `blob 1234\0`
- Then followed by raw file content or structured data
- Saved using its SHA-1 hash as the file name (Git-style Merkle tree)

---

## 🦪 Example Usage

```bash
# Initialize a Goit repository
go run cmd/main.go init

# Create a test file
echo "hello world" > hello.txt

# Hash and store the object
go run cmd/main.go hash-object hello.txt
# Output: <sha1 hash>
```

---

## 💠 Tech Stack

- Language: [Golang](https://go.dev/)
- Structure: Layered CLI/core/models/storage architecture
- Hashing: SHA-1 (like Git)
- Filesystem-based object database

---

## 📚 Concepts Covered

- How Git stores files as blobs
- Using Merkle Trees for directory structure
- SHA-based content-addressable storage
- Commits as pointers to trees and parent commits
- Rebuilding Git from the ground up

---

## 🧰 Project Structure

```
goit/
├── cmd/              # CLI command entry point (main.go)
├── core/             # Logic for init, hash-object, add, etc.
├── models/           # Structs for blob/tree/commit (coming soon)
├── storage/          # File reading/writing helpers
├── go.mod
└── README.md
```

---

## 🧠 Future Plans

- `goit add`, `goit commit`, `goit log`, `goit status`
- Branching and checkout
- Remote sync and push/pull (maybe 😉)
- Web-based UI (Goithub)

---

## 🤝 Contributing

This is a learning-focused project, so contributions are welcome in the form of:
- Clean and simple PRs
- Issues, discussions, and suggestions
- Feature extensions or alternate implementations

---

## 📖 License

This project is licensed under the MIT License.

---

## ✨ Inspired By

- [Pro Git Book](https://git-scm.com/book/en/v2)
- [Git Internals](https://github.com/pluralsight/git-internals-pdf)
- Git source code (C)

