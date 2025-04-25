package core

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitRepo() error {
	if _, err := os.Stat(".goit"); os.IsNotExist(err) {
		fmt.Println("Reinitializing existing Goit repository...")
		err := os.RemoveAll(".goit")
		if err != nil {
			return fmt.Errorf("failed to remove existing .goit: %v", err)
		}
	}

	dirs := []string{
		".goit",
		filepath.Join(".goit", "object"),
		filepath.Join(".goit", "refs"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}

	headPath := filepath.Join(".goit", "HEAD")
	if err := os.WriteFile(headPath, []byte("ref: refs/main\n"), 0644); err != nil {
		return err
	}

	mainRefPath := filepath.Join(".goit", "refs", "main")
	if err := os.WriteFile(mainRefPath, []byte(""), 0644); err != nil {
		return err
	}
	return nil
}
