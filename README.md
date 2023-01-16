# 🦊 dgb.meter.notifications

A playground for the Go language

# Pre-requisites

- [Go](https://go.dev/)

# Config

Environment Variables:

- METER_READINGS_SMTP_USERNAME
- METER_READINGS_SMTP_PASSWORD
- METER_READINGS_SMTP_HOST
- METER_READINGS_SMTP_PORT
- METER_READINGS_SMTP_FROM
- METER_READINGS_RECIPIENTS
- METER_READINGS_WEBSITE
- METER_READINGS_SUBJECT

# Usage

The included Dockerfile can be used to run the application

Under linux or Wsl create a symlink to the dockerfile in the root directory using

    ln -s ./build/Dockerfile Dockerfile

Build the image from the root directory

    docker build -t <your tag> .

Run the image

    docker run -it -d --rm -p 8000:8000 <your tag>
