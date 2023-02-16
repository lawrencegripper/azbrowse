package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/blang/semver"
	"github.com/lawrencegripper/azbrowse/internal/pkg/storage"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const versionCacheKey = "versionCacheKey"
const autoUpdateKey = "autoUpdateKey"

func updateLastCheckedTime() {
	// The value of the cache key doesn't matter. Key used to hold and identify if ttl has expired
	err := storage.PutCacheForTTL(versionCacheKey, "set")
	if err != nil {
		log.Panic(err)
	}
}

func confirmAndSelfUpdate() {
	skipUpdateEnv := os.Getenv("AZBROWSE_SKIP_UPDATE")
	if skipUpdateEnv != "" {
		log.Println("AZBROWSE_SKIP_UPDATE set so update check skipped")
		return
	}

	isCacheValid, _, _ := storage.GetCacheWithTTL(versionCacheKey, time.Hour*6)
	// Allow users to force an update by setting env
	forceUpdate := os.Getenv("AZBROWSE_FORCE_UPDATE")
	if forceUpdate == "" && isCacheValid {
		log.Print("Skipping update check as already run in last 6 hours. Set AZBROWSE_FORCE_UPDATE=true to force update check")
		return
	}

	log.Print("Checking for updates")

	latest, found, err := selfupdate.DetectLatest("lawrencegripper/azbrowse")
	if err != nil {
		log.Println("Error occurred while detecting version:", err)
		return
	}

	updateLastCheckedTime()

	v, err := semver.Parse(strings.Replace(version, "v", "", 1))
	if err != nil {
		log.Println("Failed to parse version number")
		log.Println(err.Error())
	}
	if !found || latest.Version.LTE(v) {
		log.Println("Current version is the latest")
		return
	}

	// Disable auto update if we're running from Snap as the filesystem is readonly so we can't update ourselves
	isSnap := os.Getenv("SNAP_NAME")
	if isSnap == "azbrowse" {
		fmt.Print("\n\n UPDATE AVAILABLE \n \n Release notes: "+latest.ReleaseNotes+" \n You installed via snap - upgrade to: ", latest.Version, " by running 'sudo snap refresh azbrowse'. Press any key to continue.")
		_, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Panicf("Invalid input: '%v'", err)
			return
		}
		return
	}

	autoUpdateRaw, err := storage.GetCache(autoUpdateKey)
	if err != nil {
		panic(fmt.Errorf("Failed getting autoupdate key: %w", err))
	}

	if autoUpdateRaw != "true" {
		fmt.Print("\n\n UPDATE AVAILABLE \n \n Release notes: "+latest.ReleaseNotes+" \n Do you want to update to: ", latest.Version, "? 'y'/'n' or 'auto' (for always automatically): ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil || (input != "y\n" && input != "n\n" && input != "y\r\n" && input != "n\r\n" && input != "auto\n" && input != "auto\r\n") {
			log.Panicf("Invalid input: '%s'\n", input)
			return
		}
		if input == "n\n" || input == "n\r\n" {
			return
		}

		if input == "auto\n" || input == "auto\r\n" {
			fmt.Println("Auto update set for future updates")
			err := storage.PutCache(autoUpdateKey, "true")
			if err != nil {
				panic(fmt.Errorf("Failed setting autoupdate key: %w", err))
			}
		}
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
