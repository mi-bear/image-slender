package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
	"github.com/mi-bear/image-slender/slender"
	"golang.org/x/sync/errgroup"
)

var Config = struct {
	Images []string
}{}

func main() {
	if err := execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
	}
}

func execute() error {
	configor.Load(&Config, "images.yaml")
	images := Config.Images

	eg := &errgroup.Group{}

	for _, image := range images {
		image := image
		eg.Go(func() error {
			slenderImage := slender.SlenderImage(image)
			return slenderImage.Make()
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}
