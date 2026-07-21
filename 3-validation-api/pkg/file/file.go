package file

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileDeps struct {
	path string
}

type FileWrite struct {
	Hash  string
	Email string
}

func createFile(path string) {
	_, err := os.Create(path)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func ReadFile(deps FileDeps) ([]byte, error) {
	b, err := os.ReadFile(deps.path)

	if err != nil {
		createFile(deps.path)
		return nil, err
	}

	return b, nil
}

func WriteFile(path string, email string, hash string) error {
	item := FileWrite{
		Email: email,
		Hash:  hash,
	}
	var items []FileWrite
	data, err := os.ReadFile(path)
	if err == nil {
		err = json.Unmarshal(data, &items)
		if err != nil {
			return fmt.Errorf("unmarshal error: %w", err)
		}
	}

	items = append(items, item)

	data, err = json.MarshalIndent(items, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("write file error: %w", err)
	}

	return nil
}
