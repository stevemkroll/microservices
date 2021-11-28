package main

import (
	"gateway/pkg"
	"log"
)

func main() {
	s := pkg.NewService()
	if err := s.Run(); err != nil {
		log.Panic(err)
	}
}
