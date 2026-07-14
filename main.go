package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/tree"
)

// App info
var (
	mainTitle     = "ghostshell"
	treeTitle     = "○ Version"
	latestVersion = "v0.1.1"
)

// Main colors
var (
	white = lipgloss.Color("#FFFFFF")
	green = lipgloss.Color("#0de572")
	mauve = lipgloss.Color("#7b0dea")
)

// Style definitions
var headerStyle = lipgloss.NewStyle().
	Bold(true).
	Align(lipgloss.Center).
	Foreground((white)).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground((mauve))

var treeRootStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground((mauve))

var treeSubRootStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground((green))

func expandPath(path string) string {

	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Error getting home directory: %v\n", err)
		}
		return filepath.Join(home, path[1:])
	}
	return path
}

func loadVersion(versionFile string) {

	if versionFile == "" {
		return
	}

	versionFile = expandPath(versionFile)
	content, err := os.ReadFile(versionFile)
	if err != nil {
		log.Printf("Warning: Could not read version file %s: %v\n", versionFile, err)
		return
	}
	latestVersion = strings.TrimSpace(string(content))
}

func makeTree(version string) {

	t := tree.Root(treeRootStyle.Render(treeTitle)).
		Child(
			tree.New().
				Root(treeSubRootStyle.Render(version)),
		)
	lipgloss.Println(t)
}

func main() {

	// Define flags
	titleFlag := flag.String("title", mainTitle, "Main title to display")
	versionFlag := flag.String("version", latestVersion, "Version to display")
	versionFileFlag := flag.String("version-file", "", "Path to file containing version (overrides --version)")
	treeTitleFlag := flag.String("tree-title", treeTitle, "Tree root title")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "ghostshell - A simple and elegant greeter for your CLI apps\n\n")
		fmt.Fprintf(os.Stderr, "Usage: ghostshell [options]\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Update from flags
	mainTitle = *titleFlag
	treeTitle = *treeTitleFlag

	// Clear screen
	fmt.Printf("\033[2J\033[H")

	// Load version from file if provided, otherwise use flag value
	if *versionFileFlag != "" {
		loadVersion(*versionFileFlag)
	} else {
		latestVersion = *versionFlag
	}

	// Render output
	lipgloss.Println(headerStyle.Render("", mainTitle, ""))
	makeTree(latestVersion)
}
