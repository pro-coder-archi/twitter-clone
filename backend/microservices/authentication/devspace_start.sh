#!/bin/bash
set +e  #* continue on errors

echo ☁️ installing sudo
set -ex && apk --no-cache add sudo

echo ☁️ installing nodemon
sudo npm install -g nodemon

echo 🧊 you are now in the dev container.

# include project's bin/ folder in PATH
export PATH="./bin:$PATH"

# navigating to the authentication microservice directory
cd ./backend/microservices/authentication

# open shell
bash --norc
