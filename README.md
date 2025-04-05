A modern CLI utility belt for developers, built with Go.

DevGo is a streamlined command-line toolkit that provides fast, offline utilities for common developer tasks. Instead of searching the web for one-off tools or writing custom scripts, DevGo offers a unified interface to productivity-enhancing utilities.
Key Features

Zero Web Dependency: All utilities work completely offline
Clipboard Integration: Results are automatically copied to your clipboard for immediate use
Consistent Interface: Standardized command structure and help documentation
Developer-Centric: Built by developers, for developers, solving real-world workflow needs

Installation

## From Source
bashCopy# Clone the repository
git clone https://github.com/jorgerodrigues/devgo.git

Build the binary
cd devgo
go build -o devgo main.go

### Optional: Move to a directory in your PATH
mv devgo /usr/local/bin/

# Command Structure
All DevGo commands follow a consistent pattern:
Copydevgo <command> [subcommand] [flags]
Help is available throughout:
bashCopy# General help
devgo --help

## Command-specific help
devgo uuid --help

# Available Utilities
## UUID Generation
Generate cryptographically secure UUID v4s directly to your clipboard.
bashCopy# Generate a UUID v4 and copy to clipboard
devgo uuid

## Image to Base64 Conversion
Convert image files to Base64 encoded strings for embedding in HTML, CSS, or JSON.
bashCopy# Convert an image to Base64 and copy to clipboard
devgo imgToBase64 path/to/image.png

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

License
This project is licensed under the MIT License
