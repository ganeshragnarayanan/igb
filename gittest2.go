package main

import (
	"fmt"
	"os/exec"
)

func main() {
    statusText, err := exec.Command("git", "status", "-s").Output()
	if err != nil {
	     fmt.Println("hello world1")
		
	}
	if len(statusText) > 0 {
		_, err = exec.Command("git", "add", ".").Output()
		if err != nil {
		   fmt.Println("hello world2")
		}

		commitMsg := "Sweet commit:\n" + string(statusText)
		_, err = exec.Command("git", "commit", "-a", "-m", commitMsg).Output()
		if err != nil {
			fmt.Println("hello world3")
		}

		_, err = exec.Command("git", "push").Output()
		if err != nil {
		    fmt.Println("hello world4")
			
		}

		
	} else {
		
	}
}