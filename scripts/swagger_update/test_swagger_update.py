import pytest
from swagger_update import *


def test_get_api_version_from_readme_with_invalid_file_path_returns_none():
    result = get_api_version_from_readme(
        "./scripts/swagger_update/test_data/does_not_exist.md"
    )
    assert result == None, "Should return None for file that doesn't exist"


def test_get_api_version_from_readme_with_simple_tags_returns_correct_fileset():
    api_version = get_api_version_from_readme(
        "./scripts/swagger_update/test_data/file_with_simple_tags.md"
    )

    assert api_version != None, "Expected api_version"

    assert api_version.get_name() == "package-2019-06-preview"
    input_files = api_version.get_input_files()
    assert len(input_files) == 3
    assert (
        input_files[0]
        == "Microsoft.ContainerRegistry/stable/2019-05-01/containerregistry.json"
    )
    assert (
        input_files[1]
        == "Microsoft.ContainerRegistry/preview/2019-06-01-preview/containerregistry_build.json"
    )
    assert (
        input_files[2]
        == "Microsoft.ContainerRegistry/preview/2019-05-01-preview/containerregistry_scopemap.json"
    )


def test_get_api_version_from_readme_with_multiple_tags_on_a_line_returns_correct_fileset():
    api_version = get_api_version_from_readme(
        "./scripts/swagger_update/test_data/file_with_multiple_tags_per_line.md"
    )

    assert api_version != None, "Expected api_version"

    assert api_version.get_name() == "package-2019-08"
    input_files = api_version.get_input_files()
    assert len(input_files) == 2
    assert (
        input_files[0]
        == "Microsoft.CertificateRegistration/stable/2019-08-01/AppServiceCertificateOrders.json"
    )
    assert (
        input_files[1]
        == "Microsoft.CertificateRegistration/stable/2019-08-01/CertificateRegistrationProvider.json"
    )
