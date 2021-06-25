import shutil
import os
from shutil import copytree, ignore_patterns

def get_folder_for_file(file):
    return file[0 : file.rfind("/")]


def copy_file_ensure_paths(source_base, target_base, file):
    source_file = source_base + "/" + file
    target_file = target_base + "/" + file
    print("--> " + target_file)
    target_folder = get_folder_for_file(target_file)
    os.makedirs(target_folder, exist_ok=True)
    shutil.copy(source_file, target_file)

def copy_child_folder_if_exists(source_base, target_base, relative_path, ignore=""):
    source_path = source_base + "/" + relative_path
    if os.path.exists(source_path):
        target_path = target_base + "/" + relative_path
        print("---> " + relative_path)
        copytree(source_path, target_path, dirs_exist_ok=True, ignore=ignore_patterns(ignore))
