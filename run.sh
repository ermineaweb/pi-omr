#!/bin/bash

if [[ $# -eq 0 ]] ; then
    echo "filename needed in arguments"
    exit 0
fi

(cd xml-parser-go && ./build.sh)
(cd pibuzzer && ./build.sh)

mkdir -p xml
NAME="${1%.*}"

docker run -it --rm -v $(pwd)/output:/output -v $(pwd)/input/"$1":/input/"$1" toprock/audiveris 

unzip output/"$NAME"/"$NAME".mxl -d ./xml
mv xml/"$NAME".xml xml/sheet.xml
rm xml/META-INF -R

docker run -it --rm -v $(pwd)/xml/sheet.xml:/app/sheet.xml xml-parser-go:latest > sheet_string

docker run -it --rm -v $(pwd)/sheet_string:/app/sheet_string pibuzzer:latest