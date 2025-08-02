//go:build integration || redis
// +build integration redis

package redis

import (
	"testing"

	"github.com/google/uuid"
	"github.com/spf13/viper"

	"github.com/crypto-tester/terraform-backend/pkg/lock/util"
	"github.com/crypto-tester/terraform-backend/pkg/terraform"
)

func init() {
	viper.AutomaticEnv()
}

func TestLock(t *testing.T) {
	l := NewLock()

	util.LockTest(t, l)
}

func TestGetLock(t *testing.T) {
	l := NewLock()

	expectedLock := uuid.New().String()

	s := &terraform.State{
		ID:      terraform.GetStateID("test", "test"),
		Project: "test",
		Name:    "test",
		Lock:    terraform.LockInfo{ID: expectedLock},
	}

	{
		err := l.setLock(s)
		if err != nil {
			t.Error(err)
		}
	}

	// retrieve it again
	{
		lock, err := l.getLock(s)
		if err != nil {
			t.Error(err)
		}

		if lock.ID != expectedLock {
			t.Errorf("lock mismatch: %s != %s", lock.ID, expectedLock)
		}
	}

	// delete lock
	{
		err := l.deleteLock(s)
		if err != nil {
			t.Error(err)
		}
	}
}
