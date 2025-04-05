# DevGo - A modern CLI utility belt for developers, built with Go.

DevGo is a streamlined command-line toolkit that provides fast, offline utilities for common developer tasks..

Installation

## From Source

Clone the repository:
```bash
git clone https://github.com/jorgerodrigues/devgo.git
```

Build the binary
```bash
cd devgo
go build -o devgo main.go
```

### Optional: Move to a directory in your PATH

```bash
mv devgo /usr/local/bin/
```

# Command Structure

All DevGo commands follow a consistent pattern:

```bash
devgo <command> [subcommand] [flags]
```

Help is available throughout:
General help
```bash
devgo --help
```

## Command-specific help

```bash
devgo uuid --help
```

# Available Utilities

## UUID Generation

Generate cryptographically secure UUID v4s directly to your clipboard.
Generate a UUID v4 and copy to clipboard
```bash
devgo uuid
```

## Image to Base64 Conversion

Convert image files to Base64 encoded strings for embedding in HTML, CSS, or JSON.
Convert an image to Base64 and copy to clipboard
```bash
devgo imgToBase64 path/to/image.png
```

## More coming soon...

# Contributing

Contributions are welcome! If you have ideas for new utilities or improvements:

Check the issues to see if your idea is already being discussed
Fork the repository
Create a feature branch (git checkout -b feature/amazing-utility)
Commit your changes (git commit -m 'Add some amazing utility')
Push to the branch (git push origin feature/amazing-utility)
Open a Pull Request

Please include tests and documentation with your submissions.

# License
This project is licensed under the MIT License
