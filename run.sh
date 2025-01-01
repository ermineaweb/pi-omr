#!/bin/bash

if [[ $# -eq 0 ]] ; then
    echo "Filename needed in arguments"
    exit 0
fi

if [ ! -f ./input/"$1" ]; then
    echo "File input/$1 not found"
    exit 0
fi

mkdir -p xml
NAME="${1%.*}"

docker run -it --rm -v $(pwd)/output:/output -v $(pwd)/input/"$1":/input/"$1" ermineaweb/audiveris 

unzip output/"$NAME"/"$NAME".mxl -d ./xml
mv xml/"$NAME".xml xml/sheet.xml
rm xml/META-INF -R

docker run -it --rm -v $(pwd)/xml/sheet.xml:/app/sheet.xml ermineaweb/xmlparser > sheet_string

docker run -it --rm -v $(pwd)/sheet_string:/app/sheet_string ermineaweb/pibuzzer