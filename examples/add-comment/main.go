package main

import (
	"flag"
	"log"
	"os"

	"github.com/torie/figma"
)

func main() {
	at := flag.String("access-token", "", "personal access token from Figma")
	key := flag.String("key", "", "key to Figma file")
	msg := flag.String("message", "", "the text to add as comment")
	x := flag.Float64("x", 0, "the X position where the comment should be added")
	y := flag.Float64("y", 0, "the Y position where the comment should be added")
	flag.Parse()

	if *at == "" || *key == "" {
		flag.Usage()
		os.Exit(-1)
	}

	c := figma.New(*at)

	comment, err := c.AddComment(*key, *msg, figma.Vector{*x, *y})
	if err != nil {
		log.Println(err)
	}

	log.Println("Added comment:", comment)
}
