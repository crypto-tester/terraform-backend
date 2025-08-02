package filesystem

import (
	"errors"
	"fmt"
	"os"

	"github.com/crypto-tester/terraform-backend/pkg/storage"
	"github.com/crypto-tester/terraform-backend/pkg/terraform"
)

const Name = "fs"

type FileSystemStorage struct {
	directory string
}

func NewFileSystemStorage(directory string) (*FileSystemStorage, error) {
	err := os.MkdirAll(directory, 0700)
	if err != nil {
		return nil, fmt.Errorf("failed to create directory %s: %v", directory, err)
	}

	return &FileSystemStorage{
		directory: directory,
	}, nil
}

func (f *FileSystemStorage) GetName() string {
	return Name
}

func (f *FileSystemStorage) SaveState(s *terraform.State) error {
	return os.WriteFile(f.getFileName(s.ID), s.Data, 0600)
}

func (f *FileSystemStorage) GetState(id string) (*terraform.State, error) {
	if _, err := os.Stat(f.getFileName(id)); errors.Is(err, os.ErrNotExist) {
		return nil, storage.ErrStateNotFound
	}

	d, err := os.ReadFile(f.getFileName(id))
	if err != nil {
		return nil, err
	}

	return &terraform.State{
		Data: d,
	}, nil
}

func (f *FileSystemStorage) DeleteState(id string) error {
	return os.Remove(f.getFileName(id))
}

func (f *FileSystemStorage) getFileName(id string) string {
	return fmt.Sprintf("%s/%s.tfstate", f.directory, id)
}

func (f *FileSystemStorage) CountStoredObjects() (int, error) {
	d, err := os.Open(f.directory)
	if err != nil {
		return 0, err
	}
	defer d.Close()

	list, err := d.Readdirnames(-1)
	if err != nil {
		return 0, err
	}

	return len(list), nil
}
