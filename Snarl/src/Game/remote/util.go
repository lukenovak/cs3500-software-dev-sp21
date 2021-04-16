package remote

import "net"

func BlockingRead(conn net.Conn) *[]byte {
	byteChan := make(chan []byte)
	b := make([]byte, 4096)
	go func() {
		for {
			n, _ := conn.Read(b)
			if n > 0 {
				byteChan <- b
				break
			}
		}
	}()
	for {
		select {
		case <- byteChan:
			return &b
		default:
			continue
		}
	}
}

