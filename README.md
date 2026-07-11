## ghostshell

A very simple, yet elegant greeter for your terminal. Built with Go and lipgloss.

### Installation

#### Using curl
Dowload the latest release and the `sha256sum`:
```bash
curl -L -O "https://github.com/g5ostXa/ghostshell/releases/download/v0.1.0/ghostshell-v0.1.0-linux-amd64"
curl -L -O "https://github.com/g5ostXa/ghostshell/releases/download/v0.1.0/sha256sum.txt"
```
Verify your download with `sha256sum`:
```bash
sha256sum -c sha256sum.txt
```
You should get something like this:
```
ghostshell-v0.1.0-linux-amd64: OK
```
Make sure the binary is executable:
```bash
chmod +x ghostshell-v0.1.0-linux-amd64
```
Move the binary in your `$GOBIN` as `ghostshell`:
```bash
mv ./ghostshell-v0.1.0-linux-amd64 "$GOBIN/ghostshell"
```
#### Using git
Clone the the latest git:

```bash
git clone https://github.com/g5ostXa/ghostshell.git
```

Build the binary:
```bash
cd /path/to/ghostshell
make install
```

The binary will be installed to `$GOBIN` (or `$HOME/go/bin` by default).

### Quick Start

Run the greeter:

```bash
ghostshell
```

### Usage

```bash
ghostshell [options]
```

#### Options

- `--title string` - Main title to display (default: "ghostshell")
- `--version string` - Version to display (default: "0.0-1")
- `--version-file string` - Path to file containing version (overrides --version). Supports `~` for home directory.
- `--tree-title string` - Tree root title (default: "○ Version")
- `-h, --help` - Show help message

#### Examples

```bash
# Custom title
ghostshell --title "MyApp"

# Custom version
ghostshell --version "1.2.3"

# Load version from file
ghostshell --version-file ~/.myapp/version.txt

# Custom title and version file
ghostshell --title "MyApp" --version-file ~/.myapp/version.txt

# Show help
ghostshell -h
```

### Customization

#### Colors and Styling

To customize colors and styling, edit the definitions in `main.go`:

```go
var (
    white = lipgloss.Color("#FFFFFF")
    green = lipgloss.Color("#0de572")
    mauve = lipgloss.Color("#7b0dea")
)
```

Modify the style definitions to change appearance:

```go
var headerStyle = lipgloss.NewStyle().
    Bold(true).
    Align(lipgloss.Center).
    Foreground((white)).
    BorderStyle(lipgloss.RoundedBorder()).
    BorderForeground((mauve))
```

#### Version File Example

Create a version file at `~/.myapp/version.txt`:

```bash
mkdir -p ~/.myapp
echo "1.0.0" > ~/.myapp/version.txt
```

Then use it:

```bash
ghostshell --title "MyApp" --version-file ~/.myapp/version.txt
```

### Development

Build locally:

```bash
make build
```

Run after building:

```bash
make run
```

Clean build artifacts:

```bash
make clean
```
