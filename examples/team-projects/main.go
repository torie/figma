package main

import (
	"flag"
	"log"
	"os"

	"github.com/torie/figma"
)

func main() {
	at := flag.String("access-token", "", "personal access token from Figma")
	id := flag.String("id", "", "id to team")
	flag.Parse()

	if *at == "" || *id == "" {
		flag.Usage()
		os.Exit(-1)
	}

	c := figma.New(*at)

	projects, err := c.TeamProjects(*id)
	if err != nil {
		log.Println(err)
	}

	for _, project := range projects {
		log.Println(project)
	}
}
