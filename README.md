# ASCII Art Justify

A command-line tool that renders text as ASCII art with support for multiple alignment types and banner styles. The output adapts dynamically to your terminal size.

## Usage

```bash
go run . [OPTION] [STRING] [BANNER]
```

### Arguments

| Argument | Description | Required |
|----------|-------------|----------|
| `--align=<type>` | Alignment type (see below) | No (defaults to left) |
| `STRING` | Text to render as ASCII art | Yes |
| `BANNER` | Banner style (see below) | No (defaults to standard) |

### Alignment Types

| Flag | Description |
|------|-------------|
| `--align=left` | Aligns text to the left |
| `--align=right` | Aligns text to the right |
| `--align=center` | Centers text horizontally |
| `--align=justify` | Distributes spaces evenly between words |

### Banner Styles

| Banner | Description |
|--------|-------------|
| `standard` | Classic ASCII art style |
| `shadow` | Shadow effect style |
| `thinkertoy` | Toy-like block style |

## Examples

```bash
# Default (left aligned, standard banner)
go run . "hello"

# Right aligned with shadow banner
go run . --align=right "hello world" shadow

# Centered with thinkertoy banner
go run . --align=center "hello" thinkertoy

# Justified across terminal width
go run . --align=justify "how are you doing" standard

# Multiple lines using \n
go run . --align=center "hello\nworld"
```

## Example Output

```bash
$ go run . --align=right "hi" standard

                                                                               _|   _
                                                                               _|  (_)
                                                                            _|_|_|  _
                                                                               _|  | |
                                                                               _|  | |
                                                                              _|_|  |_|
```

## Project Structure

```
ascii-justify/
├── main.go          # Argument parsing and program entry point
├── banner.go        # loadBanner — reads banner files from disk
├── render.go        # renderChar, RenderWord, renderWords
├── alignment.go     # Alignment — handles all 4 alignment types
├── utils.go         # getTerminalWidth — queries terminal size
└── banner/
    ├── standard.txt
    ├── shadow.txt
    └── thinkertoy.txt
```

## How It Works

### Character Rendering
Each banner file stores ASCII art for all printable characters (starting from space, ASCII 32). Every character occupies exactly 8 lines, separated by a blank line. To find character `c`:

```
startLine = (c - 32) * 9 + 1
```

### Alignment
The program queries the terminal width dynamically using `stty size`. Each alignment type pads the rendered text differently:

- **Left** — no padding, renders as-is
- **Right** — adds padding to the left
- **Center** — splits padding evenly on both sides
- **Justify** — distributes spaces between words across the full terminal width, with leftover spaces added one per gap from the left

### Newlines
Use `\n` in your string to render multiple lines:
```bash
go run . "hello\nworld"
```

## Error Handling

Invalid flag format:
```
Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard
```

## Requirements

- Go 1.18 or higher
- Unix-based terminal (Linux/macOS)

## Author

Built as part of the ascii-art series of projects at Zone01.
