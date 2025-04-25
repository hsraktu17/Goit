# Goit – A Git-like Version Control System in Go 🌀

Goit is a simplified version control system written in **Golang**, designed to mimic how Git works under the hood using **Merkle Trees**, **content-addressable storage**, and core concepts like **blobs**, **trees**, and **commits**.

> 🚧 This is a learning-focused project built to understand Git internals from scratch.

---

## 🚀 Features

- `goit init` – Initialize a Goit repository
- `goit hash-object <file>` – Store a file as a blob and return its hash
- `goit add <file>` – Stage a file (planned)
- `goit write-tree` – Create a Merkle tree from added files (planned)
- `goit commit -m "msg"` – Create a commit with snapshot (planned)
- `goit cat-file <hash>` – Read back stored object (planned)

---

## 🧱 Repository Structure

