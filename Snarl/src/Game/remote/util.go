package remote

import (
	"bufio"
)

// BlockingRead reads from a connection, but blocks until we have data in the connection
func BlockingRead(r *bufio.Reader) *[]byte {
	byteChan := make(chan []byte)
	b := make([]byte, 4096)
	go func() {
		for {
			n, _ := r.ReadBytes('\n')
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

