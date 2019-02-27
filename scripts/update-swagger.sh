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
# The output to create is
#     swagger-specs
#       |- top-level
#           |- common-types
#           |- service1
#           |    |- resource-type1
#           |    |    |- common 
#           |    |    |    |- *.json
#           |    |    |- latest 
#           |    |         |- placeholder (need to preserve the relative layout of the versioned files and common folder) 
#           |    |              |- *.json
#           |    |- resource-type2
#           |         |- latest 
#           |              |- placeholder 
#           |                   |- *.json
#           |- service2
#           |    |- resource-type1
#           |    |    |- latest 
#           |    |         |- placeholder 
#           |    |              |- *.json
#           ...


# Get top-level 'service' folders
serviceFolders=$(ls -d $ApiRepo/specification/*/)

for serviceFolder in $serviceFolders
do
    serviceName=$(basename $serviceFolder)
    echo "$serviceName"

    # Get resource-type folders 
    { 
        swaggerFolders=$(ls -d ${serviceFolder}resource-manager/*/)
    } || {
        swaggerFolders=""
    }
    for swaggerFolder in $swaggerFolders
    do
        resourceType=$(basename $swaggerFolder)
        echo "    $resourceType"

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
        for specFolder in $specFolders
        do
            latestSpecFolder=$specFolder
        done

        # if we found a latest version then start copying
        if [[ -n "$latestSpecFolder" ]];
        then
            latestSpec=$(basename $latestSpecFolder)

            # Check if we have a common folder to copy
            if [[ -d $swaggerFolder/common ]];
            then
                mkdir swagger-specs/top-level/$serviceName/$resourceType --parents
                cp $swaggerFolder/common swagger-specs/top-level/$serviceName/$resourceType -r
            fi
            # Check if we have a common folder to copy
            if [[ -d ${serviceFolder}resource-manager/common ]];
            then
                mkdir swagger-specs/top-level/$serviceName --parents
                cp ${serviceFolder}resource-manager/common swagger-specs/top-level/$serviceName -r
            fi

            mkdir swagger-specs/top-level/$serviceName/$resourceType/$specBranch/$latestSpec --parents
            cp ${latestSpecFolder}* swagger-specs/top-level/$serviceName/$resourceType/$specBranch/$latestSpec/ -r
        fi
    done
    echo ""
done

# Temporary fixup for Microsoft.Web
# The latest version is 2018-02-01, except for Certificates.json which has 2018-11-01
# Since the 2018-11-01 Certificates has no real difference (and has references to the 2018-02-01 folder)
# manually copy back the 2018-02-01 files
echo "WARNING: Forcing 2018-02-01 version of web/Microsoft.Web as a temporary fix"
rm swagger-specs/top-level/web/Microsoft.Web/stable/2018-11-01 -r # this will fail if there is a newer versino than 2018-11-01, which is a good sanity check
mkdir swagger-specs/top-level/web/Microsoft.Web/stable/2018-02-01 --parents
cp $ApiRepo/specification/web/resource-manager/Microsoft.Web/stable/2018-02-01/*.json swagger-specs/top-level/web/Microsoft.Web/stable/2018-02-01

# copy the top-level `common-types` folder
cp $ApiRepo/specification/common-types swagger-specs -r
