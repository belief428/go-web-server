package web

import (
	"testing"
)

func TestNewWeb(t *testing.T) {
	NewWeb()(nil).Start(nil)
}
