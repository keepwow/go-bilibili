# Go-Bilibili

Go-Bilibili is a program written in Go that interacts with the Bilibili platform.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have installed the latest version of Go.
- You have a `<Windows/Linux/Mac>` machine.
- **You have installed `ffmpeg` on your machine. This is crucial for the program to function correctly.**

## Installing Go-Bilibili

To install Go-Bilibili, follow these steps within the current direcotry:

```bash
go build go-bilibili.go         # build the source code
mv go-bilibili /usr/local/bin   # move the output to any path in the PATH, usually /usr/local/bin
```

## Using Go-Bilibili

To use Go-Bilibili, follow these steps:

```bash
Usage: go-bilibili -convert=yes [-inputDir=/path/to/input] [-outputDir=/path/to/output] [--dry-run]"
```

Or you can just run from the source code:

```bash
go run go-bilibili.go -convert=yes [-inputDir=/path/to/input] [-outputDir=/path/to/output] [--dry-run]"
```

### Examples

```bash
% go-bilibili
Usage: go run main.go -convert=yes [-inputDir=/path/to/input] [-outputDir=/path/to/output] [--dry-run]

% go-bilibili -convert=yes -outputDir="$HOME/Videos"
Do you want to keep the original directories? (yes/no): yes
Successfully created output file: /Users/keepwow/Videos/【全英】你还在边听听力边看文本？！/【全英】你还在边听听力边看文本？！.mp4

% go-bilibili -convert=yes -outputDir="$HOME/Videos"
Do you want to keep the original directories? (yes/no): no
Output file already exists. Do you want to keep the original files? (yes/no): no
Removing original files...
```

## Contributing to Go-Bilibili

To contribute to Go-Bilibili, follow these steps:

1. Fork this repository.
2. Create a branch: `git checkout -b '<branch_name>'`.
3. Make your changes and commit them: `git commit -m '<commit_message>'`
4. Push to the original branch: `git push origin '<project_name>/<location>'`
5. Create the pull request.

Alternatively, see the GitHub documentation on [creating a pull request](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request).

## Contact

If you want to contact me, you can reach me at [GitHub: keepwow](https://github.com/keepwow).

## License

This project uses the following license: [MIT](https://opensource.org/licenses/MIT).
