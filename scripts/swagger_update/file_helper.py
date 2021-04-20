import shutil
import os
from distutils.dir_util import copy_tree

def get_folder_for_file(file):
    return file[0 : file.rfind("/")]


def copy_file_ensure_paths(source_base, target_base, file):
    source_file = source_base + "/" + file
    target_file = target_base + "/" + file
    print("--> " + target_file)
    target_folder = get_folder_for_file(target_file)
    os.makedirs(target_folder, exist_ok=True)
    shutil.copy(source_file, target_file)

def copy_child_folder_if_exists(source_base, target_base, relative_path):
    source_path = source_base + "/" + relative_path
    if os.path.exists(source_path):
        target_path = target_base + "/" + relative_path
        print("---> " + relative_path)
        copy_tree(source_path, target_path)
