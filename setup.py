from subprocess import check_call

from setuptools import setup
from setuptools.command.develop import develop
from setuptools.command.install import install


class PreDevelopCommand(develop):
    def run(self):
        check_call(["go", "build", "-o", "bin/cotangent", "./cmd/cotangent"])
        develop.run(self)


class PreInstallCommand(install):
    def run(self):
        check_call(["go", "build", "-o", "bin/cotangent", "./cmd/cotangent"])
        install.run(self)


import setuptools

with open("README.md", "r", encoding="utf-8") as fh:
    long_description = fh.read()

setuptools.setup(
    name="cotangent",
    version="0.0.1",
    author="Ken Shibata",
    author_email="kenxshibata@gmail.com",
    description="A simple Markdown renderer",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/colourdelete/cotangent",
    project_urls={
        "Bug Tracker": "https://github.com/colourdelete/cotangent/issues",
    },
    classifiers=[
        "Programming Language :: Python :: 3",
    ],
    package_dir={"": "src"},
    packages=setuptools.find_packages(where="src"),
    python_requires=">=3.6",
)
