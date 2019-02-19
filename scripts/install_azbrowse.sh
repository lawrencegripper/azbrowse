#!/bin/bash

# Inspired by great work by Alex Ellis openfaas cli installer script -> https://github.com/openfaas/faas-cli

version=$(curl -sI https://github.com/lawrencegripper/azbrowse/releases/latest | grep Location | awk -F"/" '{ printf "%s", $NF }' | tr -d '\r')
if [ ! $version ]; then
    echo "Failed please install manually"
    exit 1
fi

hasCli() {
    has=$(which azbrowse)

    if [ "$?" = "0" ]; then
        echo
        echo "You already have azbrowse installed. It should self update so this script is calling it a day"
        echo
        exit 1
    fi

    hasCurl=$(which curl)
    if [ "$?" = "1" ]; then
        echo "You need curl to use this script."
        exit 1
    fi
}

getPackage() {
    uname=$(uname)
    userid=$(id -u)

    if [[ $uname = *Darwin* ]]; then
      suffix="-darwin-amd64"
    fi 

    if [[ $uname = *"Linux"* ]]; then
      suffix="-linux-amd64"
    fi 

    echo $suffix

    if [ -z "$suffix" ]; then
        echo "Could detect version, please install manually"
        exit 1
    fi

    targetFile="/tmp/azbrowse$suffix"
    
    if [ "$userid" != "0" ]; then
        targetFile="$(pwd)/azbrowse$suffix"
    fi

    if [ -e $targetFile ]; then
        rm $targetFile
    fi

    url=https://github.com/lawrencegripper/azbrowse/releases/download/$version/azbrowse$suffix
    echo "Downloading package $url as $targetFile"

    curl --fail -sSL $url --output $targetFile

    if [ "$?" != "0" ]; then
        echo "Failed to find release :( URL: $url "
        exit 1
    fi 

    chmod +x $targetFile

    echo "Download complete."
       
    if [ "$userid" != "0" ]; then
        
        echo
        echo "=========================================================" 
        echo "==    As the script was run as a non-root user the     =="
        echo "==    following commands may need to be run manually   =="
        echo "========================================================="
        echo
        echo "  sudo cp azbrowse$suffix /usr/local/bin/azbrowse"
        echo "  sudo ln -sf /usr/local/bin/azbrowse /usr/local/bin/azbrowse"
        echo
    else
        echo
        echo "Running as root - Attempting to move azbrowse to /usr/local/bin"
        mv $targetFile /usr/local/bin/azbrowse
    
        if [ "$?" = "0" ]; then
            echo "azbrowse installed to /usr/local/bin"
        fi
        if [ -e $targetFile ]; then
            rm $targetFile
        fi

        echo "Installed verion:"
        echo
        azbrowse --version
        
        echo "You may need to reload your terminal to pick up the new var"
    fi
    
}

hasCli
getPackage