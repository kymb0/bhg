package main

import (
    "fmt"
    "os"
	"string"
    
    "github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Prompt{
		Label: "Where should we send tha l00tz? (UNC path of your SMB server EG \\\\10.13.3.7\\\\GUEST): ",
		Validate: func(input string) error {
			if input.Contains("\\\\") != 1 {
				return errors.New("Search term must have at least 3 characters")
			}
			return nil
		},
	}
    }
    
    keyword, err := prompt.Run()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    
    fmt.Printf("Search for %q\n", keyword)
}
