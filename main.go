package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/amendgit/kit"
)

var (
	soruceDir = kit.SourceDir()
	hintsDir  = path.Join(soruceDir, "hints")
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	opt, args := os.Args[1], os.Args[2:]
	switch opt {
	case "+", "edit":
		editHint(args)
	case "help":
		showHelp(args)
	default:
		showHint(os.Args[1:])
	}
}

func showHelp(args []string) {
	fmt.Print("help message of hint")
}

func showHint(args []string) {
	name := args[0]
	if !kit.IsPathExist(hintsDir) {
		os.MkdirAll(hintsDir, 0700)
	}
	hintPath := path.Join(hintsDir, name+".md")
	bs, _ := ioutil.ReadFile(hintPath)
	fmt.Print(string(bs))
}

func editHint(args []string) {
	name := args[0]
	hintPath := path.Join(hintsDir, name+".md")
	fmt.Printf("hintPath %v", hintPath)
	exec.Command("code", hintPath).Run()
}
