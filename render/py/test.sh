#!/bin/sh
# -*- coding: utf-8 -*-

set -e

curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | python
export PATH="$PATH:$HOME/.poetry/bin"
poetry install
poetry run tox
exec poetry run coverage report -m
