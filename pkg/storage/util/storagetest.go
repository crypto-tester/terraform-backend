package util

import (
	"errors"
	"testing"

	"github.com/spf13/viper"

	"github.com/crypto-tester/terraform-backend/pkg/storage"
	"github.com/crypto-tester/terraform-backend/pkg/terraform"
)

func init() {
	viper.AutomaticEnv()
}

func StorageTest(t *testing.T, s storage.Storage) {
	state := &terraform.State{
		ID:      terraform.GetStateID("test", "test"),
		Project: "test",
		Name:    "test",
		Data:    []byte("test"),
	}

	nonExisting, err := s.GetState(state.ID)
	if nonExisting != nil || !errors.Is(err, storage.ErrStateNotFound) {
		t.Error("non existing state should return ErrStateNotFound")
	}

	if err := s.SaveState(state); err != nil {
		t.Error(err)
	}

	savedState, err := s.GetState(state.ID)
	if err != nil {
		t.Error(err)
	}

	if string(state.Data) != string(savedState.Data) {
		t.Errorf("state data does not match")
	}

	state.Data = []byte("test2")

	if err := s.SaveState(state); err != nil {
		t.Error(err)
	}

	savedState, err = s.GetState(state.ID)
	if err != nil {
		t.Error(err)
	}

	if string(state.Data) != string(savedState.Data) {
		t.Errorf("state data does not match")
	}

	err = s.DeleteState(state.ID)
	if err != nil {
		t.Error(err)
	}
}
