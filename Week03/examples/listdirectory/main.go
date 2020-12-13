package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("..", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("err1", err)
			return err
		}
		fmt.Println(path, info.Size())
		return nil
	})
}
