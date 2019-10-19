#!/bin/bash
set -e

# This script updates the local copy of the azure-rest-api-specs in this repo
# Only the latest specs for each service are held in this repo
# Keeping the specs local gives a record of which specs are used in a build
# It also ensures repeatable build of a given commit in this repo
# as the build doesn't pull the latest specs


# Get the latest azure swagger specs
# Put them inside a folder with a .gitignore to avoid adding them to this repo in full
rm -rf swagger-temp
mkdir swagger-temp
echo "*" > swagger-temp/.gitignore
git clone https://github.com/azure/azure-rest-api-specs swagger-temp/azure-rest-api-specs --depth=1
ApiRepo="swagger-temp/azure-rest-api-specs"

# Reset the swagger-specs folder in this repo
rm -rf ./swagger-specs
mkdir ./swagger-specs

#
# The `specification` folder in the azure-rest-api-specs repo contains the folder hierarchy for the swagger specs
#
#     specification
#          |-service1 (e.g. `cdn` or `compute`)
#          |   |-common
#          |   |-quickstart-templates
#          |   |-data-plane
#          |   |-resource-manager (we're only interested in the contents of this folder)
#          |       |- resource-type1 (e.g. `Microsoft.Compute`)
#          |       |    |- common
#          |       |    |   |- *.json (want these)
#          |       |    |- preview
#          |       |    |    |- 2016-04-20-preview
#          |       |    |        |- *.json
#          |       |    |- stable
#          |       |    |    |- 2015-06-15
#          |       |    |        |- *.json
#          |       |    |    |- 2017-12-01
#          |       |    |        |- *.json
#          |       |    |        |- examples
#          |       |    |    |- 2018-10-01
#          |       |    |        |- *.json   (want these)
#          |       |    |        |- examples
#          |       |- misc files (e.g. readme) 
#          ...
#
#
# For each top level folder (service name) iterate the resource type folders under resource-manager
# For each resource type find the latest stable release (or the latest preview if no stable is available)
#   and then take the json files in that directory (ignoring subfolders such as examples)
#
#
# The output to create is
#  swagger-specs
#          |-service1 (e.g. `cdn` or `compute`)
#          |   |-common   (want these)
#          |   |-quickstart-templates
#          |   |-data-plane
#          |   |-resource-manager (we're only interested in the contents of this folder)
#          |       |- resource-type1 (e.g. `Microsoft.Compute`)
#          |       |    |- common
#          |       |    |   |- *.json (want these)
#          |       |    |- stable (NB - may preview if no stable)
#          |       |    |    |- 2018-10-01
#          |       |    |        |- *.json   (want these)
#          |       |- misc files (e.g. readme) 
#           ...


# Get top-level 'service' folders
serviceFolders=$(ls -d $ApiRepo/specification/*/)

# serviceFolder: e.g. specification/web
for serviceFolder in $serviceFolders 
do
    serviceName=$(basename $serviceFolder)
    echo "$serviceName - $serviceFolder"

    # Get resource-type folders 
    { 
        swaggerFolders=""
        if [[ -d "${serviceFolder}resource-manager" ]] 
        then
            swaggerFolders=$(ls -d ${serviceFolder}resource-manager/*/)
        fi
    } || {
        swaggerFolders=""
    }
    # swaggerFolder: specification/web/resource-manager/Microsoft.Web
    for swaggerFolder in $swaggerFolders
    do
        resourceType=$(basename $swaggerFolder)
        echo "    $resourceType - $swaggerFolder"

        # Get latest version folder
        {        
            specFolders=$(ls -d ${swaggerFolder}stable/*/ 2>/dev/null)
            specBranch="stable"
        } || {
            specFolders=$(ls -d ${swaggerFolder}preview/*/ 2>/dev/null)
            specBranch="preview"
        } || {
            specFolders=""
            specBranch=""
        }
        latestSpecFolder=""
        # specFolder: specification/web/resource-manager/Microsoft.Web/stable/2000-01-01
        for specFolder in $specFolders
        do
            latestSpecFolder=$specFolder
        done

        # if we found a latest version then start copying
        if [[ -n "$latestSpecFolder" ]];
        then
            latestSpec=$(basename $latestSpecFolder)

            # Check if we have a common folder to copy at the serviceFolder level
            if [[ -d ${serviceFolder}resource-manager/common ]];
            then
                mkdir swagger-specs/$serviceName/resource-manager --parents
                cp ${serviceFolder}resource-manager/common swagger-specs/$serviceName/resource-manager -r
            fi
            # Check if we have a common folder to copy at the swaggerFolder level
            if [[ -d ${serviceFolder}resource-manager/$resourceType/common ]];
            then
                mkdir swagger-specs/$serviceName/resource-manager/$resourceType --parents
                cp $swaggerFolder/common swagger-specs/$serviceName/resource-manager/$resourceType -r
            fi

            # Copy the spec folder
            mkdir swagger-specs/$serviceName/resource-manager/$resourceType/$specBranch/$latestSpec --parents
            cp ${latestSpecFolder}* swagger-specs/$serviceName/resource-manager/$resourceType/$specBranch/$latestSpec/ -r
        fi
    done
    echo ""
done

# copy the top-level `common-types` folder
cp $ApiRepo/specification/common-types swagger-specs -r
