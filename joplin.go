package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

func main() {
	if len(os.Args) < 3 {
		binary := filepath.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "usage: %s <title> <notebook-id>\n", binary)
		os.Exit(1)
	}

	nbID, err := uuid.Parse(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid notebook ID")
		os.Exit(1)
	}

	n := note{
		ID:         uuid.New(),
		NotebookID: nbID,
		Title:      os.Args[1],
		Timestamp:  time.Now(),
	}

	content, err := n.generate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error generating note: %s", err)
		os.Exit(1)
	}

	noteID := strings.Replace(n.ID.String(), "-", "", -1)
	noteFile := noteID + ".md"
	err = os.WriteFile(noteFile, []byte(content), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error saving note: %s", err)
		os.Exit(1)
	}

	fmt.Println(noteFile)
}
