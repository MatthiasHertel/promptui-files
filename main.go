package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/manifoldco/promptui"
)

func getFiles() []string {
	files, err := ioutil.ReadDir("./test")
	if err != nil {
		log.Fatal(err)
	}

	var result []string
	for _, f := range files {
		result = append(result, f.Name())
	}
	return result
}

func main() {
	prompt := promptui.Select{
		Label: "Select Files",
		Items: getFiles(),
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}
