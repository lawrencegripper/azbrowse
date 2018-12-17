package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/blang/semver"
	"github.com/lawrencegripper/azbrowse/version"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

func confirmAndSelfUpdate() {
	latest, found, err := selfupdate.DetectLatest("lawrencegripper/azbrowse")
	if err != nil {
		log.Println("Error occurred while detecting version:", err)
		return
	}

	v, err := semver.Parse(version.BuildDataVersion)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if !found || latest.Version.LTE(v) {
		log.Println("Current version is the latest")
		return
	}

	fmt.Print("\n\n UPDATE AVAILABLE \n \n Release notes: "+latest.ReleaseNotes+" \n Do you want to update to: ", latest.Version, "? (y/n): ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil || (input != "y\n" && input != "n\n" && input != "y\r\n" && input != "n\r\n") {
		log.Printf("Invalid input: '%s'\n", input)
		return
	}
	if input == "n\n" || input == "n\r\n" {
		return
	}

	exe, err := os.Executable()
	if err != nil {
		log.Println("Could not locate executable path")
		return
	}
	if err := selfupdate.UpdateTo(latest.AssetURL, exe); err != nil {
		log.Println("Error occurred while updating binary:", err)
		return
	}
	log.Println("Successfully updated to version", latest.Version)
}
