package main

import (
	"errors"
	"fmt"

	"github.com/jpreese/slowimage"
)

func main() {
	i := slowimage.Image{
		Filename: "whatever",
	}

	result, err := SlowDownload(i)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

type Downloader interface {
	Download() string
}

func SlowDownload(d Downloader) (string, error) {
	filename := d.Download()
	if filename == "" {
		return "", errors.New("blank filename")
	}

	return filename, nil
}
