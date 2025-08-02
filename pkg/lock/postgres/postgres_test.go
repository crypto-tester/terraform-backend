//go:build integration || postgres
// +build integration postgres

package postgres

import (
	"testing"

	"github.com/spf13/viper"

	"github.com/crypto-tester/terraform-backend/pkg/lock/util"
)

func init() {
	viper.AutomaticEnv()
}

func TestLock(t *testing.T) {
	l, err := NewLock("locks")
	if err != nil {
		t.Error(err)
	}

	util.LockTest(t, l)
}
