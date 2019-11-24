package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/adrienkohlbecker/picturelint/validators"
	"github.com/logrusorgru/aurora"

	"github.com/adrienkohlbecker/picturelint/cases"
	"github.com/adrienkohlbecker/picturelint/picture"
)

func main() {

	files := make(chan string, 0)
	wg := sync.WaitGroup{}

	go func() {

		for f := range files {
			visitFile(f)
			wg.Done()
		}

	}()

	err := filepath.Walk(os.Args[1], func(p string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", p, err)
			return err
		}

		if !info.IsDir() {

			if info.Name() == ".DS_Store" || info.Name() == "Thumbs.db" {
				return nil
			}
			if strings.HasPrefix(info.Name(), ".") {
				return nil
			}

			ext := strings.ToLower(path.Ext(p))
			if ext == ".xmp" {
				return nil
			}

			wg.Add(1)
			files <- p

		}

		return nil

	})

	if err != nil {
		fmt.Println("error walking the path")
		os.Exit(1)
	}

	close(files)
	wg.Wait()

}

func visitFile(path string) error {

	p, err := picture.Load(path)
	if err != nil {
		fmt.Printf("Error loading %s: %s\n", path, err)
		os.Exit(1)
	}

	res := cases.Run(p)
	failed := false

	fileOutput := ""

	fileOutput += fmt.Sprintf("%s:\n", aurora.Gray(10, p.Path))

	for _, kase := range res {
		if kase.Status == validators.StatusFailed {
			failed = true
			fileOutput += fmt.Sprintf("  %s\n", aurora.Red(fmt.Sprintf("✗ %s", kase.Legend)))
		}
	}

	if failed {
		fmt.Print(fileOutput)
		return fmt.Errorf("failed file: %s", path)
	}

	return nil

}
