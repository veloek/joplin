package main

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
)

const templateFile string = "note.tmpl"

type note struct {
	ID         uuid.UUID
	NotebookID uuid.UUID
	Title      string
	Timestamp  time.Time
}

func (n *note) generate() (string, error) {
	funcMap := template.FuncMap{
		"time": func(t time.Time) string {
			return t.UTC().Format(time.RFC3339)
		},
		"uuid": func(id uuid.UUID) string {
			return strings.Replace(id.String(), "-", "", -1)
		},
	}

	dir, err := getExeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(dir, templateFile)

	t, err := template.New("").Funcs(funcMap).ParseFiles(path)
	if err != nil {
		return "", err
	}

	output := new(strings.Builder)
	err = t.ExecuteTemplate(output, templateFile, n)
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func getExeDir() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}

	ex, err = filepath.EvalSymlinks(ex)
	if err != nil {
		return "", err
	}

	return filepath.Dir(ex), nil
}
