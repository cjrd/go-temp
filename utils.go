/*
 utils provides functions utility functions for the temp package.
*/

package temp

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// reseed the random generator
func reseed() uint32 {
	return uint32(time.Now().UnixNano() + int64(os.Getpid()))
}

var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var lenLetters = len(letters)

// Returns a random short string of a random letter + 9 numbers + random letter
// useful for e.g. filenames, web ids, etc
func randShortString() string {
	r := reseed()
	r = r*1664525 + 1013904223 // constants from Numerical Recipes
	return string(letters[rand.Intn(lenLetters)]) + strconv.Itoa(int(1e9+r%1e9)) +
		string(letters[rand.Intn(lenLetters)])
}
