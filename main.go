package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	// this needs to be package level so it won't be dropped during optimisation
	hash []byte

	maxDuration    = flag.Duration("max", 100*time.Millisecond, "a maximum duration that should be spent on hashing")
	passwordLength = flag.Int("length", 30, "the length of the randomly generated password")
)

func main() {
	flag.Parse()

	// generate a random password
	passwd := make([]byte, *passwordLength)
	rand.Read(passwd)

	log.Printf("Generated password: %s", base64.StdEncoding.EncodeToString(passwd))

	for i := bcrypt.DefaultCost; i < bcrypt.MaxCost; i++ {
		var err error

		start := time.Now()
		if hash, err = bcrypt.GenerateFromPassword(passwd, i); err != nil {
			panic(err)
		}
		d := time.Now().Sub(start)

		log.Printf("Computed hash with cost factor %d in %s", i, d)

		if d >= *maxDuration {
			break
		}
	}
}
