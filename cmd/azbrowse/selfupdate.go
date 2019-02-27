package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

func confirmAndSelfUpdate() {
	latest, found, err := selfupdate.DetectLatest("lawrencegripper/azbrowse")
	if err != nil {
		log.Println("Error occurred while detecting version:", err)
		return
	}

	v, err := semver.Parse(version)
	if err != nil {
		log.Panicln(err.Error())
	}
	if !found || latest.Version.LTE(v) {
		log.Println("Current version is the latest")
		return
	}

	fmt.Print("\n\n UPDATE AVAILABLE \n \n Release notes: "+latest.ReleaseNotes+" \n Do you want to update to: ", latest.Version, "? (y/n): ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil || (input != "y\n" && input != "n\n" && input != "y\r\n" && input != "n\r\n") {
		log.Panicf("Invalid input: '%s'\n", input)
		return
	}
	if input == "n\n" || input == "n\r\n" {
		return
	}

	exe, err := os.Executable()
	if err != nil {
		log.Panicln("Could not locate executable path")
	}
	if err := selfupdate.UpdateTo(latest.AssetURL, exe); err != nil {
		log.Panicln("Error occurred while updating binary:", err)
	}
	log.Println("Successfully updated to version", latest.Version)
}
