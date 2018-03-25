package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

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

	f, err := c.File(*key)
	if err != nil {
		log.Println(err)
	}

	docs := f.Nodes()
	log.Printf("Got %d documents", len(docs))

	fdocs := frames(docs)
	log.Printf("Got %d frames", len(fdocs))

	images, err := c.Images(*key, 2, figma.ImageFormatPNG, ids(fdocs)...)
	if err != nil {
		log.Println(err)
	}

	log.Printf("Downloading %d images\n", len(images))
	os.MkdirAll(*key, os.ModePerm)
	for _, img := range images {
		rc, err := download(img)
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(rc)
		if err != nil {
			log.Fatal(err)
		}
		path := filepath.Join(*key, fmt.Sprintf("%s.png", img.NodeID))
		if err := ioutil.WriteFile(path, data, 0666); err != nil {
			log.Fatal(err)
		}
	}
}

func ids(docs []figma.Node) []string {
	var res []string
	for i := range docs {
		res = append(res, docs[i].ID)
	}
	return res
}

func frames(docs []figma.Node) []figma.Node {
	var res []figma.Node
	for i := range docs {
		if docs[i].Type == figma.NodeTypeFrame {
			res = append(res, docs[i])
		}
	}
	return res
}

func download(i figma.Image) (io.ReadCloser, error) {
	resp, err := http.Get(i.URL)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
