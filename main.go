package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/kdisneur/obsidiantomela/export"
)

func main() {
	fset := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fset.Usage = func() {
		fmt.Fprintf(fset.Output(), "%s <obsidian-folder> <mela-folder>", fset.Name())
		fmt.Fprintln(fset.Output())
		fmt.Fprintln(fset.Output(), "\tobsidian-folder")
		fmt.Fprintln(fset.Output(), "\t\tpath to where all Obsidian recipes are stored")
		fmt.Fprintln(fset.Output(), "\t\tnote: the script will only take recipes at the root level")
		fmt.Fprintln(fset.Output())
		fmt.Fprintln(fset.Output(), "\tmela-folder")
		fmt.Fprintln(fset.Output(), "\t\tbase folder where a new Mela folder will be created to hold all recipes")
		fmt.Fprintln(fset.Output(), "\t\tnote: folder format is mela_<timestamp>")
	}
	_ = fset.Parse(os.Args[1:])

	if fset.NArg() != 2 {
		fset.Usage()
		os.Exit(1)
	}

	err := run(fset.Arg(0), fset.Arg(1))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run(obsidianFolder, melaBaseFolder string) error {
	melaFolder := path.Join(melaBaseFolder, fmt.Sprintf("mela_%d", time.Now().Unix()))
	err := os.MkdirAll(melaFolder, 0755)
	if err != nil {
		return fmt.Errorf("failed to create Mela export folder %q: %v", melaFolder, err)
	}

	exporter, err := export.LoadObsidianFolder(obsidianFolder)
	if err != nil {
		return err
	}

	err = exporter.Export(melaFolder)
	if err != nil {
		return err
	}

	fmt.Printf("recipes exported to %q\n", melaFolder)

	return nil
}
