package conn

import (
	"log"

	"github.com/fhs/gompd/mpd"
)

// Conn is a var to expose the MPD connection
var Conn *mpd.Client

// Connect to local mpd client
func Connect() {
	conn, err := mpd.Dial("tcp", "localhost:6600")
	if err != nil {
		log.Fatalln(err)
	}
	Conn = conn
}
