package remote

import (
	"bufio"
	"net"
)

// BlockingRead reads from a connection, but blocks until we have data in the connection
func BlockingRead(conn net.Conn) *[]byte {
	byteChan := make(chan []byte)
	b := make([]byte, 4096)
	reader := bufio.NewReader(conn)
	go func() {
		for {
			n, _ := reader.ReadBytes('\n')
			if len(n) > 0 {
				byteChan <- n
				break
			}
		}
	}()
	for {
		select {
		case b = <- byteChan:
			return &b
		default:
			continue
		}
	}
}

