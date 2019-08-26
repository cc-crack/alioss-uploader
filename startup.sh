#!/bin/bash

docker run --name alioss-uploader -v $(pwd)/config.json:/etc/alioss-uploader/config.json -p 9002:9002  -d  pk8995/alioss-uploader