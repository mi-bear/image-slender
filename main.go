package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
	"github.com/mi-bear/image-slender/slender"
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

	for _, image := range images {
		image := image
		slenderImage := slender.SlenderImage(image)
		if err := slenderImage.Make(); err != nil {
			return err
		}
	}
	return nil
}
