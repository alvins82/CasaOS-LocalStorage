package v2

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEnsureDefaultDirectories(t *testing.T) {
	root := t.TempDir()

	if err := EnsureDefaultDirectories(root); err != nil {
		t.Fatal(err)
	}

	expectedDirectories := []string{
		"AppData",
		"Documents",
		"Downloads",
		"Gallery",
		filepath.Join("Media", "Movies"),
		filepath.Join("Media", "TV Shows"),
		filepath.Join("Media", "Music"),
	}
	for _, relativePath := range expectedDirectories {
		info, err := os.Stat(filepath.Join(root, relativePath))
		if err != nil {
			t.Fatalf("expected %s to be created: %v", relativePath, err)
		}
		if !info.IsDir() {
			t.Fatalf("expected %s to be a directory", relativePath)
		}
	}
}

func TestEnsureDefaultDirectoriesPreservesExistingContent(t *testing.T) {
	root := t.TempDir()
	existingFile := filepath.Join(root, "Documents", "readme.txt")
	if err := os.MkdirAll(filepath.Dir(existingFile), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(existingFile, []byte("keep me"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := EnsureDefaultDirectories(root); err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(existingFile)
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "keep me" {
		t.Fatalf("existing file was changed: %q", content)
	}
}

func TestEnsureDefaultDirectoriesRejectsFilePath(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "Gallery"), []byte("not a directory"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := EnsureDefaultDirectories(root); err == nil {
		t.Fatal("expected an error when a default directory path is a file")
	}
}
