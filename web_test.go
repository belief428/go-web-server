package go_web_server

import "testing"

func TestNewWeb(t *testing.T) {
	NewWeb()(nil).Start(nil)
}
