import os
from subprocess import Popen
from cookiecutter.main import cookiecutter

PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)
app_name = '{{ cookiecutter.module_name }}'

def init_protoc():
    """
    Initialises protoc on the new project folder
    """
    path = os.path.join('api',app_name,'v1')
    PROTOC_COMMANDS = [
        ["kratos", "proto","client",path],
    ]

    for command in PROTOC_COMMANDS:
        w = Popen(command, cwd=os.path.join(PROJECT_DIRECTORY))
        w.wait()

def init_wire():
    """
    Initialises wire on the new project folder
    """
    WIRE_COMMANDS = [
        ["wire", "."],
    ]

    for command in WIRE_COMMANDS:
        w = Popen(command, cwd=os.path.join(PROJECT_DIRECTORY,'cmd','server'))
        w.wait()

init_protoc()
init_wire()