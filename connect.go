package twitchbot

import (
	"log"
	"net"
	"time"
)

// User a struct that handles the connection.
type User struct {
	name  string
	token string
}

// New method for struct connect initialization.
func New(userName, token string) *User {
	return &User{name: userName, token: token}
}

// Connect method for irc server connection.
func Connect(u User, url string) {
	conn, err := net.DialTimeout("tcp", url, 2*time.Second)
	if err != nil {
		// error handling.
		log.Fatal("Dial Error: ", err)
	}

	var size int
	var b []byte
	b = make([]byte, 1024)
	_, err = conn.Write([]byte("PASS " + u.token + "\r\n"))
	if err != nil {
		log.Fatal("Write Error: ", err)
	}
	_, err = conn.Write([]byte("NICK " + u.name + "\r\n"))
	if err != nil {
		log.Fatal("Write Error: ", err)
	}
	size, err = conn.Read(b)
	if err != nil {
		log.Fatal("Read Error: ", err)
	}
	log.Fatalf("Received %d bytes: %s\n", size, string(b[:size]))

	// return conn
}
