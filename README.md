# 🦊 dgb.meter.readings.bot

A playground for the Go language

# Pre-requisites

- [Go](https://go.dev/)

# Config

Environment Variables:



# Usage

The included Dockerfile can be used to run the application

Under linux or Wsl create a symlink to the dockerfile in the root directory using

    ln -s ./build/Dockerfile Dockerfile

Build the image from the root directory

    docker build -t <your tag> .

Run the image

    docker run -it -d --rm -p 8000:8000 <your tag>
