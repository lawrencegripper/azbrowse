import api_set
import file_helper
import git_helper
import shutil
import os

# TODO
# - restructure!

if __name__ == "__main__":
    # resource_provider_version_overrides is keyed on RP name with the value being the tag to force
    resource_provider_version_overrides = {
        "cosmos-db": "package-2019-08-preview",  # the 2019-12 version includes 2019-08-preview files that reference files not in the 2019-12 list!
        # frontdoor 2020-01 references 2019-11-01/network.json which is not listed in the input files
        # frontdoor 2019-11 references 2019-05-01/network.json which is not listed in the input files
        # frontdoor 2019-10 references 2019-10-01/network.json which is not listed in the input files
        # frontdoor 2019-05 references 2019-03-01/network.json which is not listed in the input files
        # frontdoor 2019-04 references 2019-03-01/network.json which is not listed in the input files
        "frontdoor" : "",
        # azureactivedirectory 2020-07-01-preview references files from 2020-03-01-preview which are not listed in the input files
        "azureactivedirectory": "package-2020-03-01-preview",
        # ./azsadmin seems very broken, lots of references to files cross versions
        "azsadmin": "",
        # storage package-2019-06 references privatelinks.json which is not listed in the input files
        "storage": "package-2019-04"
    }

    print(
        "\n****************************************************************************"
    )
    print("Cloning azure-rest-api-sets repo")
    git_helper.clone_swagger_specs("swagger-temp")

    print(
        "\n****************************************************************************"
    )
    print("Deleting ")
    if os.path.exists("swagger-specs"):
        print("Deleting swagger-specs...")
        shutil.rmtree("swagger-specs")

 
    print(
        "\n****************************************************************************"
    )
    print("Discovering api-sets:")
    api_sets = api_set.get_api_sets(
        "./swagger-temp/azure-rest-api-specs/specification", 
        resource_provider_version_overrides
    )

    print(
        "\n****************************************************************************"
    )
    print("Copying api-set files:")
    api_set.copy_api_sets_to_swagger_specs(
        api_sets,
        "./swagger-temp/azure-rest-api-specs/specification",
        "./swagger-specs",
    )
    shutil.copytree(
        "./swagger-temp/azure-rest-api-specs/specification/common-types",
        "./swagger-specs/common-types",
    )

