#!/bin/bash
set +e  #* continue on errors

echo "👍 you are now in the dev container."

# include project's bin/ folder in PATH
export PATH="./bin:$PATH"

# open shell
bash --norc
