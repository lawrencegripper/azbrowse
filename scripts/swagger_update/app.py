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
        # azureactivedirectory 2020-07-01-preview references files from 2020-03-01-preview which are not listed in the input files
        "azureactivedirectory": "package-2020-03-01-preview",
        # ./azsadmin seems very broken, lots of references to files cross versions
        "azsadmin": "",
        # recovery service and recoveryservicebackup list the same templateurls with different api version
        # this causes indeterminate behaviour when generating swagger apisets
        "recoveryservices": "",
        "recoveryservicesbackup": "",
        "recoveryservicessiterecovery": "",
        "automation": "package-2015-10",
        "applicationinsights": "package-2020-02-12",
        # Pin as has missing/invalid Microsoft.Databricks/preview/2022-04-01-preview/databricks.json file
        "databricks": "package-2021-04-01-preview",
        "billing":"",
        # ai/dataplane readme references package-2023-11-06-beta but the files do not exist
        "ai": "package-2024-05-01-preview",
        # 2024/11/25 17:05:43 Error expanding Swagger: open /workspaces/azbrowse/swagger-specs/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/operations.json: no such file or directory
        "policyinsights": ""
    }

    # This allows you to augment the included files for each README.MD for a specific tag
    # this is useful when files which are needed are incorrectly left out of the 'input-file'
    resource_provider_input_file_additions = {
        "compute": {
            "package-2024-03-01": ["./Microsoft.Compute/ComputeRP/stable/2023-09-01/virtualMachineScaleSet.json"]
        }
        # Example:
        # "cosmos-db" : {
        #     "package-2020-04" : ["./Microsoft.DocumentDB/stable/2019-08-01/cosmos-db.json"]
        # },
    }

    print(
        "\n****************************************************************************"
    )
    print("Cloning azure-rest-api-sets repo")
    git_helper.clone_or_update_swagger_specs("swagger-temp")

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
        resource_provider_version_overrides,
        resource_provider_input_file_additions
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

