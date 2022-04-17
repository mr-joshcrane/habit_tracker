package main

import (
	"fmt"
	"os"

	"habit/stores/networkstore"
	"habit"
)

func main() {
	defaultPath := "localhost:8080"
	s, err := networkstore.Open(defaultPath)
	// user, ok := os.LookupEnv("USER")
	// if !ok {
	// 	user = "USER"
	// }
	// s, err := pbfilestore.Open(fmt.Sprintf("habit_%s", user))
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	habit.RunCLI(s)
}
