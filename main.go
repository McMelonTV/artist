package main

import "artist/internal/db"

func main() {
	err := db.Init()
	if err != nil {
		panic(err)
	}

	select {}
}
