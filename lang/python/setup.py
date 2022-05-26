from setuptools import setup, find_packages
from os import path

version = "0.0.3"

this_directory = path.abspath(path.dirname(__file__))
with open(path.join(this_directory, "README.md"), encoding="utf-8") as f:
    long_description = f.read()

setup(
    name="pyemo",
    packages=find_packages(),
    version=version,
    description="Emoji based semantic scoped debuging ",
    long_description=long_description,
    long_description_content_type="text/markdown",
    author="Teal Finance",
    author_email="Teal.Finance@pm.me",
    url="https://github.com/teal-finance/emo",
    keywords=["debuging", "emoji"],
    classifiers=[
        "Development Status :: 3 - Alpha",
        "Intended Audience :: Developers",
        "License :: OSI Approved :: MIT License",
        "Programming Language :: Python :: 3.8",
    ],
    zip_safe=False,
)
