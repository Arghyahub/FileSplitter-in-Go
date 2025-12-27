Here’s a professional README for your Go file splitter project:

---

# File Splitter in Go

A simple command-line tool written in **Go** to split large files into smaller parts. Useful for breaking files into manageable chunks for storage, transfer, or processing.

## Features

* Splits any file into a specified number of parts.
* Automatically calculates part sizes to evenly distribute data.
* Works with binary files like videos, images, or documents.
* Lightweight and fast due to buffered read/write operations.

## Requirements

* Go 1.20+ installed
* Compatible with Linux, macOS, and Windows

## Installation

1. Clone this repository:

```bash
git clone <repository_url>
cd <repository_directory>
```

2. Build the project:

```bash
go build -o splitter main.go
```

## Usage

```bash
./splitter <FILENAME> <NUMBER_OF_PARTS>
```

* `<FILENAME>` – Path to the file you want to split
* `<NUMBER_OF_PARTS>` – Total number of parts to split into

**Example:**

```bash
./splitter video.mp4 4
```

This will generate:

```
video.mp4-part1
video.mp4-part2
video.mp4-part3
video.mp4-part4
```

## How It Works

* Opens the original file and calculates the size of each part.
* Reads the file in **blocks of 1000 bytes**.
* Writes each block sequentially into the new part files.

## Notes

* Ensure you have sufficient disk space for the split files.
* File permissions for new parts are set to `0777` (read/write/execute for all).

## License

MIT License – free to use, modify, and distribute.

---

If you want, I can also **add a section explaining how to merge the split files back** like `cat file-part* > originalfile` to make it complete. Do you want me to do that?
