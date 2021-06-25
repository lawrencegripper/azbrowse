import file_helper
import shutil
import os
import re
import yaml
import json


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
#          |       |- readme.md  (this lists api versions and the files in each version) 
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
#          |       |- resource-type1 (e.g. `Microsoft.Compute`)
#          |       |    |- common
#          |       |    |   |- *.json (want these)
#          |       |    |- stable (NB - may preview if no stable)
#          |       |    |    |- 2018-10-01
#          |       |    |        |- *.json   (want these)
#          |       |- api-set.json  (based on content in readme.md but easier for subsequent parsing)
#          |   |-resource-manager (we're only interested in the contents of this folder)
#          |       |- resource-type1 (e.g. `Microsoft.Compute`)
#          |       |    |- common
#          |       |    |   |- *.json (want these)
#          |       |    |- stable (NB - may preview if no stable)
#          |       |    |    |- 2018-10-01
#          |       |    |        |- *.json   (want these)
#          |       |- api-set.json  (based on content in readme.md but easier for subsequent parsing)
#           ...


class ApiSet:
    def __init__(self, resource_provider_name, base_folder, api_version):
        self.resource_provider_name = resource_provider_name
        self.base_folder = base_folder
        self.api_version = api_version
    def get_resource_provider_name(self):
        return self.resource_provider_name

    def get_base_folder(self):
        return self.base_folder

    def get_api_version(self):
        return self.api_version

class ApiVersion:
    def __init__(self, name, input_files, addition_input_file_paths):
        self.name = name
        self.input_files = input_files
        self.addition_input_file_paths = addition_input_file_paths

    def get_name(self):
        return self.name

    def get_input_files(self):
        return self.input_files + self.addition_input_file_paths

    def to_json(self):
        return json.dumps(self.__dict__, ensure_ascii=False, sort_keys=True)

tag_regex = re.compile("openapi-type: [a-z\\-]+\ntag: ([a-z\\-0-9]*)")
tag_from_header_regex = re.compile("### Tag: (package-[0-9]{4}-[0-9]{2}.*)")
def get_api_version_tag(resource_provider_name, readme_contents, overrides):
    override = overrides.get(resource_provider_name)
    if override != None:
        return override

    match = tag_regex.search(readme_contents)
    if match == None:
        return None

    tag = match.group(1)

    if 'preview' not in tag:
        return tag
    
    print('default tag was preview, falling back to latest stable tag')
    match = tag_from_header_regex.search(readme_contents)
    if match == None:
        return None
    
    return match.group(1)

code_block_end_regex = re.compile("^[\\s]*```[\\s]*$", flags=re.MULTILINE)
def find_api_version(resource_provider_name, readme_contents, version_tag, input_file_additions):

    # Regex to match:   ```yaml $(tag) == 'the-version-tag`
    # Also match:       ```yaml $(tag) == 'the-version-tag` || $(tag) == 'some-other-tag'
    # But don't match   ```yaml $(tag) == 'the-version-tag' && $(another-condition)
    start_match = re.search(
        "^```[\\s]*yaml [^&^\\n]*\\$\\(tag\\) == '" + version_tag + "'[^&^\\n]*$",
        readme_contents,
        flags=re.MULTILINE,
    )
    if start_match == None:
        return None

    end_match = code_block_end_regex.search(readme_contents, start_match.end())
    yaml_contents = readme_contents[start_match.end() : end_match.start()]

    yaml_data = yaml.load(yaml_contents, Loader=yaml.BaseLoader)
    input_files = []
    if yaml_data != None:
        input_files = [file.replace("\\", "/") for file in yaml_data["input-file"]]
    
    additional_input_file_paths = get_additional_files_for_version(input_file_additions, resource_provider_name, version_tag)
    api_version = ApiVersion(
        version_tag, 
        input_files, 
        additional_input_file_paths
    )
    
    return api_version


def get_api_version_from_readme(resource_provider_name, readme_path, version_overrides, input_file_additions):
    if not os.path.isfile(readme_path):
        return None
    print("==> Opening: " + readme_path)
    with open(readme_path, "r", encoding="utf8") as stream:
        contents = stream.read()

    version_tag = get_api_version_tag(resource_provider_name, contents, version_overrides)
    if version_tag == None:
        print("==> no version tag found in readme: " + readme_path)
        return None

    api_version = find_api_version(resource_provider_name, contents, version_tag, input_file_additions)
    return api_version

def copy_api_sets_to_swagger_specs(api_sets, source_folder, target_folder):
    for api_set in api_sets:
        print("\nCopying " + api_set.get_resource_provider_name())
        api_version = api_set.get_api_version()

        resource_provider_source = (
            source_folder
            + "/" + api_set.get_base_folder()
        )
        resource_provider_target = (
            target_folder
            + "/" + api_set.get_base_folder()
        )

        # The core of this method is to copy the files defined by the api version
        # There are some places where additional files that aren't listed in the definition tend to live
        # For now we're handling the separate cases (e.g. `common`)
        # We _could_ load the specs and scan for linked files and build out the list that way
        # Doing that would remove the need for these additional checks
        # as well as fixing the problem with definitions referenced back in other folders as with comsmos-db etc

        # Look for `common` folder under the `resource-manager` folder
        file_helper.copy_child_folder_if_exists(
            resource_provider_source,
            resource_provider_target,
            "/common",
        )

        # Look for `common` folders under the resource type folder
        resource_type_folders = set(
            [x[0 : x.index("/")] for x in api_version.get_input_files()]
        )
        for resource_type_folder in resource_type_folders:
            file_helper.copy_child_folder_if_exists(
                resource_provider_source,
                resource_provider_target,
                resource_type_folder + "/common",
            )

        # Look for `entityTypes` or `definitions` folders under api versions
        api_version_folders = set(
            [x[0 : x.rfind("/")] for x in api_version.get_input_files()]
        )
        for api_version_folder in api_version_folders:
            file_helper.copy_child_folder_if_exists(
                resource_provider_source,
                resource_provider_target,
                api_version_folder + "/entityTypes",
            )
            file_helper.copy_child_folder_if_exists(
                resource_provider_source,
                resource_provider_target,
                api_version_folder + "/definitions",
            )

            # Hack: Handle the case where a package version folder has files which aren't in the README docs
            file_helper.copy_child_folder_if_exists(
                resource_provider_source,
                resource_provider_target,
                api_version_folder,
                ignore='examples'
            )

            # find 'common.json' or ... 'Common.json'
            if os.path.exists(resource_provider_source + "/" + api_version_folder + "/common.json"):
                file_helper.copy_file_ensure_paths(resource_provider_source, resource_provider_target, api_version_folder + "/common.json")
            elif os.path.exists(resource_provider_source + "/" + api_version_folder + "/Common.json"):
                file_helper.copy_file_ensure_paths(resource_provider_source, resource_provider_target, api_version_folder + "/Common.json")

        # Copy the files defined in the api version
        for file in api_version.get_input_files():
            file_helper.copy_file_ensure_paths(resource_provider_source, resource_provider_target, file)

        # Write api-set.json file per folder with contents to load for swagger-codegen
        api_set_filename = resource_provider_target + "/api-set.json"
        print("Writing " + api_set_filename)
        with open(api_set_filename, "w") as f:
            f.write(api_version.to_json())

def get_api_set_for_folder(spec_folder, api_folder, resource_provider_name, version_overrides, input_file_additions):
    api_version = get_api_version_from_readme(resource_provider_name, api_folder + "/readme.md", version_overrides, input_file_additions)
    if api_version == None:
        return None
    spec_relative_folder = api_folder[len(spec_folder) + 1 :]
    
    api_set = ApiSet(
        resource_provider_name,
        spec_relative_folder, 
        api_version
    )
    return api_set

def get_additional_files_for_version(input_file_additions, resource_provider_name, api_version):
    files_to_add = []
    try:
        files_to_add = input_file_additions[resource_provider_name][api_version]
    except KeyError:
        pass
    
    return files_to_add


def get_api_sets(spec_folder, version_overrides, input_file_additions):
    rp_folders = sorted([f.path for f in os.scandir(spec_folder) if f.is_dir()])
    api_sets = []
    for folder in rp_folders:
        resource_provider_name = folder.split("/")[-1]
        if version_overrides.get(resource_provider_name) == "":
            print("Resource provider " + resource_provider_name + " is skipped in config")
            continue

        got_api_set = False
        
        for api_type_folder in ["resource-manager", "data-plane"]:
            qualified_api_type_folder = folder + "/" + api_type_folder
            if not os.path.exists(qualified_api_type_folder):
                continue
            
            api_set = get_api_set_for_folder(spec_folder, qualified_api_type_folder, resource_provider_name, version_overrides, input_file_additions)
            if api_set != None:
                api_sets.append(api_set)
                got_api_set = True
            else:
                # didn't find readme.md under (e.g.) search/data-plane/
                # look for search/data-plane/*/readme.md
                print("\n*************************************************************************************")
                print(qualified_api_type_folder)
                # sub_folders = [f.path for f in os.scandir(qualified_api_type_folder) if f.is_dir() and os.path.exists(qualified_api_type_folder + "/" + f.path + "/readme.md")]
                sub_folders = [f.path for f in os.scandir(qualified_api_type_folder) if f.is_dir()]
                for sub_folder in sub_folders:
                    print(sub_folder)
                    api_set = get_api_set_for_folder(spec_folder, sub_folder, resource_provider_name, version_overrides, input_file_additions)
                    if api_set != None:
                        print("got api_set")
                        api_sets.append(api_set)
                        got_api_set = True


        if not got_api_set:
            print("***No api version found, ignoring: " + folder)

    return api_sets
