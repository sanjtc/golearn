package main

import (
	"net"
	"testing"

	"gotest.tools/assert"
)

func TestProcessInterrupt(t *testing.T) {
	listen, err := net.Listen("tcp", "localhost:2333")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		_ = processInterrupt("localhost", "2333")
	}()

	conn, err := listen.Accept()
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 20)
	size, err := conn.Read(buf)
	conn.Close()

	if err != nil {
		t.Fatal(err)
	}

	msg := string(buf[:size])

	assert.Equal(t, msg, "interrupt")
}

// func TestSSHTask(t *testing.T) {
// 	go newSSHServer(t)

// 	_ = runSrc("localhost", "2333", "test", "test", 4, "https://www.ssetech.com.cn/")
// }

// func newSSHServer(t *testing.T) {
// 	listen, err := net.Listen("tcp", "localhost:2333")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	conn, err := listen.Accept()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	hostKeyBytes := []byte(`-----BEGIN OPENSSH PRIVATE KEY-----
// b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
// QyNTUxOQAAACCbwxBo/3QT+gE3R2U0m71gJvCeLY5wYzaaDBXd6J59HQAAAJDpU9P06VPT
// 9AAAAAtzc2gtZWQyNTUxOQAAACCbwxBo/3QT+gE3R2U0m71gJvCeLY5wYzaaDBXd6J59HQ
// AAAEDJR51JvnXwYB6ZDMIHqtE1ke12AfQ/T0Fc5OZ5FOmiRpvDEGj/dBP6ATdHZTSbvWAm
// 8J4tjnBjNpoMFd3onn0dAAAACXJvb3RAa2FsaQECAwQ=
// -----END OPENSSH PRIVATE KEY-----`)

// 	hostKey, err := ssh.ParsePrivateKey(hostKeyBytes)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	serverConfig := &ssh.ServerConfig{}
// 	serverConfig.NoClientAuth = true
// 	serverConfig.AddHostKey(hostKey)

// 	sshConn, chans, reqs, err := ssh.NewServerConn(conn, serverConfig)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Check remote address
// 	log.Printf("new connection from %s (%s)", sshConn.RemoteAddr(), sshConn.ClientVersion())

// 	go handleChannels(chans)
// 	go handleRequests(reqs)
// }

// func handleRequests(reqs <-chan *ssh.Request) {
// 	for req := range reqs {
// 		log.Printf("recieved out-of-band request: %+v", req)
// 	}
// }

// func handleChannels(chans <-chan ssh.NewChannel) {
// 	for newChannel := range chans {
// 		log.Println(newChannel.ChannelType())

// 		channel, requests, err := newChannel.Accept()
// 		if err != nil {
// 			log.Printf("could not accept channel (%s)", err)
// 			continue
// 		}

// 		log.Println("channel data:", channel)

// 		for req := range requests {
// 			switch req.Type {
// 			case "exec":
// 				log.Println("exec:", req.Type)
// 			case "shell":
// 				log.Println("shell:", req.Type)

// 				_, _ = channel.Write([]byte("test"))
// 				channel.Close()
// 			default:
// 				log.Println("default:", req.Type)
// 			}

// 			err := req.Reply(true, nil)
// 			if err != nil {
// 				log.Println(err)
// 			}
// 		}

// 		channel.Close()
// 	}
// }
