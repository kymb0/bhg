package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

//			if input.Contains("\\\\") != 1 {
func main() {
	prompt := promptui.Prompt{
		Label: "Where should we send tha l00tz? (UNC path of your SMB server EG \\\\10.13.3.7\\\\GUEST): ",
		Validate: func(input string) error {
			if strings.Contains(input, "\\\\") == false {
				return errors.New("Are you SURE that's a UNC path???")
			}
			return nil
		},
	}

	keyword, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Search for %q\n", keyword)
}
