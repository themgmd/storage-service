# Storage Service

Its simple storage-service for your VPS/VDS server

## Supported File types (to upload):
```
IMAGE: png, jpg, webp, svg, jpeg, bmp
VIDEO: mp4, wav, mov
AUDIO: mp3
DOCS: doc, docx, pptx, xlsx, csv
TEXT: txt
```

# Usage

## Installation
To use the service, you need install:
1. [Docker](https://docs.docker.com/engine/install/)
2. [Docker compose](https://docs.docker.com/compose/install/)

Clone this repo to your machine
```shell
cd DIR_PATH
git clone https://github.com/onemgvv/storage-service.git

cd storage-service
docker compose up
```

The application will start on port 5029