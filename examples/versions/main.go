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
	flag.Parse()

	if *at == "" || *key == "" {
		flag.Usage()
		os.Exit(-1)
	}

	c := figma.New(*at)

	versions, err := c.FileVersions(*key)
	if err != nil {
		log.Println(err)
	}

	for _, version := range versions {
		log.Println(version)
	}
}
