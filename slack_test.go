package golib

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCaller(t *testing.T) {
	t.Run("TEST getCaller", func(t *testing.T) {
		b := strings.Contains(getCaller(), "runtime")
		assert.Equal(t, true, b)
	})
}

func TestSendNotification(t *testing.T) {
	t.Run("TEST SendNotification", func(*testing.T) {
		os.Setenv("SLACK_NOTIFIER", "true")
		title := "test"
		body := "test"
		ctx := "test"
		err := errors.New("test")
		SendNotification(title, body, ctx, err)
	})
}
