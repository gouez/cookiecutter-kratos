import os
import re
import sh
import pytest

PATTERN = '{{(\s?cookiecutter)[.](.*?)}}'
RE_OBJ = re.compile(PATTERN)

@pytest.fixture
def context():
    return {
        "project_name": "kratos",
        "app_name":"kratos",
        "version": "0.0.1"
}