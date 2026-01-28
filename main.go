package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var swap int = 0

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		} else {
			if swap == 1 {
				fmt.Println("Swapped from jp to global. you can now launch global uma.")
			} else {
				fmt.Println("Swapped from global to jp. you can now launch jp uma.")
			}
		}
		fmt.Println("\nPress Enter to exit...")
		fmt.Scanln()
	}()
	appdata := os.Getenv("APPDATA")

	if appdata == "" {
		panic("appdata not found!")
	}

	targetDir := filepath.Join(filepath.Dir(appdata), "LocalLow", "Cygames")
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		panic(err)
	}

	if _, err := os.Stat(filepath.Join(targetDir, "Umamusumejp")); !os.IsNotExist(err) {
		if err := os.Rename((filepath.Join(targetDir, "Umamusume")), filepath.Join(targetDir, "Umamusumegb")); err != nil {
			panic("rename failed on renaming global folder to 'Umamusumegb'")
		}
		if err := os.Rename((filepath.Join(targetDir, "Umamusumejp")), filepath.Join(targetDir, "Umamusume")); err != nil {
			panic("rename failed on renaming jp folder to 'Umamusume'")
		}

	} else if _, err := os.Stat(filepath.Join(targetDir, "Umamusumegb")); !os.IsNotExist(err) {
		if err := os.Rename((filepath.Join(targetDir, "Umamusume")), filepath.Join(targetDir, "Umamusumejp")); err != nil {
			panic("rename failed on renaming jp folder to 'Umamusumejp'")
		}
		if err := os.Rename((filepath.Join(targetDir, "Umamusumegb")), filepath.Join(targetDir, "Umamusume")); err != nil {
			panic("rename failed on renaming global folder to 'Umamusume'")
		}
		swap = 1
	} else {
		panic("no Umamusumejp or Umamusumegb folder found! Crashing..")
	}

}
