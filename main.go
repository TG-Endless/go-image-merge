package main

import (
	"flag"
	"fmt"
	"goimagemerge/goimagemerge"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func help() {
	os.Exit(2)
}

func main() {
	//fromDir := filepath.ToSlash(os.Args[1])
	//toDir := filepath.ToSlash(os.Args[2])
	f := flag.String("f", "/from/dir", "Src dir containing images to merge")
	t := flag.String("t", "/to/dir", "Dest dir to save merged.jpg")
	x := flag.Int("x", 0, "horizon count")
	y := flag.Int("y", 0, "vertical count")
	flag.Parse()

	fmt.Printf("Merging........")

	fromDir := filepath.ToSlash(*f)
	toDir := filepath.ToSlash(*t)
	rows := *y
	cols := *x

	var grids = []*goimagemerge.Grid{}
	files, err := ioutil.ReadDir(fromDir)
	if err != nil {
		help()
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		f := strings.ToLower(file.Name())
		if strings.HasSuffix(f, ".png") || strings.HasSuffix(f, ".jpg") || strings.HasSuffix(f, ".jpeg") {
			grids = append(grids, &goimagemerge.Grid{
				ImageFilePath: fromDir + "/" + file.Name(),
			})
		}
	}

	if rows <= 0 {
		rows = len(grids)
	}

	if cols <= 0 {
		cols = 1
	}

	rgba, err := goimagemerge.New(grids, cols, rows).Merge()
	if err != nil {
		help()
	}

	file, err := os.Create(toDir + "/merged.jpg")
	if err != nil {
		help()
	}

	err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 100})
	if err != nil {
		help()
	}
}
