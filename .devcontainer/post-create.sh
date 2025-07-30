#!/bin/bash
set -e


if [ ! -d ".venv" ]; then
  echo "Creating Python virtual environment..."
  python3 -m venv .venv
else
  echo "Virtual environment already exists."
fi
source .venv/bin/activate

pip3 install black rope

echo "!! Run 'source .venv/bin/activate' to activate the virtual environment."


bundler install