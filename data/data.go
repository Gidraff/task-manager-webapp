package data

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	// Avoid using db driver directly
	_ "github.com/lib/pq"
)

// Db a handle to the db and rep
// a pool of db connection
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=taskman dbname=taskman password=taskman sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

/*
Encrypt plainText, and returns a
cipher of type string
*/
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
