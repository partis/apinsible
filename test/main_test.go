package main_test

import (
  "testing"
  "github.com/partis/apinsible"
  "net"
)

func Test(t *testing.T) {
  conn, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", 80))
	if err != nil {
		t.Error(err.Error)
	}
	conn.Close()
	t.Pass
}
