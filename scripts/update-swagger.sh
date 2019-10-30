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

    apiTypes=("resource-manager" "data-plane")
    for apiType in "${apiTypes[@]}"
    do
        # Get resource-type folders 
        { 
            swaggerFolders=""
            if [[ -d "${serviceFolder}$apiType" ]] 
            then
                swaggerFolders=$(ls -d ${serviceFolder}$apiType/*/)
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
            set +e # ls commands below may error - temporarily continue on errors
            specFolders=""
            specBranch=""
            specFolders=$(ls -d ${swaggerFolder}stable/*/ 2>/dev/null)
            echo $specFolders
            if [[ -n "$specFolders" ]]; 
            then
                specBranch="stable"
            fi
            if [[ -z "$specBranch" ]]; 
            then 
                specFolders=$(ls -d ${swaggerFolder}preview/*/ 2>/dev/null)
                specBranch="preview"
            fi
            latestSpecFolder=""
            # specFolder: specification/web/resource-manager/Microsoft.Web/stable/2000-01-01
            for specFolder in $specFolders
            do
                latestSpecFolder=$specFolder
            done
            set -e

            # if we found a latest version then start copying
            if [[ -n "$latestSpecFolder" ]];
            then
                latestSpec=$(basename $latestSpecFolder)

                # Check if we have a common folder to copy at the serviceFolder level
                if [[ -d ${serviceFolder}$apiType/common ]];
                then
                    mkdir swagger-specs/$serviceName/$apiType --parents
                    cp ${serviceFolder}$apiType/common swagger-specs/$serviceName/$apiType -r
                fi
                # Check if we have a common folder to copy at the swaggerFolder level
                if [[ -d ${serviceFolder}$apiType/$resourceType/common ]];
                then
                    mkdir swagger-specs/$serviceName/$apiType/$resourceType --parents
                    cp $swaggerFolder/common swagger-specs/$serviceName/$apiType/$resourceType -r
                fi

                # Copy the spec folder
                mkdir swagger-specs/$serviceName/$apiType/$resourceType/$specBranch/$latestSpec --parents
                cp ${latestSpecFolder}* swagger-specs/$serviceName/$apiType/$resourceType/$specBranch/$latestSpec/ -r
            fi
        done
    done
    echo ""
done

# Temporary fixup for Microsoft.Web
# The latest version is 2019-08-01, but in testing this wasn't yet rolled out
# Since the 2018-11-01 Certificates has no real difference (and has references to the 2018-02-01 folder), just use 2018-02-1
# manually copy back the 2018-02-01 files
echo "WARNING: Forcing 2018-02-01 version of web/Microsoft.Web etc as a temporary fix"
rm swagger-specs/web/resource-manager/Microsoft.Web/stable/2019-08-01 -r # this will fail if there is a newer version than 2018-11-01, which is a good sanity check in case we haven't remembered to revert by then
mkdir swagger-specs/web/resource-manager/Microsoft.Web/stable/2018-02-01 --parents
cp $ApiRepo/specification/web/resource-manager/Microsoft.Web/stable/2018-02-01/*.json swagger-specs/web/resource-manager/Microsoft.Web/stable/2018-02-01

rm swagger-specs/web/resource-manager/Microsoft.CertificateRegistration/stable/2019-08-01 -r # this will fail if there is a newer version than 2018-11-01, which is a good sanity check in case we haven't remembered to revert by then
mkdir swagger-specs/web/resource-manager/Microsoft.CertificateRegistration/stable/2018-02-01 --parents
cp $ApiRepo/specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2018-02-01/*.json swagger-specs/web/resource-manager/Microsoft.CertificateRegistration/stable/2018-02-01

rm swagger-specs/web/resource-manager/Microsoft.DomainRegistration/stable/2019-08-01 -r # this will fail if there is a newer version than 2018-11-01, which is a good sanity check in case we haven't remembered to revert by then
mkdir swagger-specs/web/resource-manager/Microsoft.DomainRegistration/stable/2018-02-01 --parents
cp $ApiRepo/specification/web/resource-manager/Microsoft.DomainRegistration/stable/2018-02-01/*.json swagger-specs/web/resource-manager/Microsoft.DomainRegistration/stable/2018-02-01

# copy the top-level `common-types` folder
cp $ApiRepo/specification/common-types swagger-specs -r
