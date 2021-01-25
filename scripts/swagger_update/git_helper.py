from git import Repo
import shutil
import os


last_git_message = ""

def show_git_progress(op_code, cur_count, max_count, message):
    global last_git_message
    if message != "" and message != last_git_message:
        print(message)
        last_git_message = message


def clone_swagger_specs(target_folder):
    if os.path.exists(target_folder):
        print("Deleting " + target_folder + "...")
        shutil.rmtree(target_folder)

    os.mkdir(target_folder)
    with open(target_folder + "/.gitignore", "w") as f:
        f.write("*")


    print("Cloning specs...")
    repo = Repo()
    repo.clone_from(
        "https://github.com/azure/azure-rest-api-specs",
        target_folder + "/azure-rest-api-specs",
        progress=show_git_progress,
        multi_options=["--depth=1"],
    )
    print("Cloned")

