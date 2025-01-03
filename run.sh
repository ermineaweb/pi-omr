#!/bin/bash

if [[ $# -eq 0 ]] ; then
    echo "Filename needed in arguments"
    exit 0
fi

if [ ! -f ./input/"$1" ]; then
    echo "File input/$1 not found"
    exit 0
fi

(cd xml-parser-go && ./build.sh)

rm -R result

NAME="${1%.*}"

docker run -it --rm -v $(pwd)/output:/output -v $(pwd)/input/"$1":/input/"$1" ermineaweb/audiveris

unzip output/"$NAME"/"$NAME".mxl -d ./result
rm result/META-INF -R

docker run -it --rm -v $(pwd)/result/"$NAME".xml:/app/"$NAME".xml xml-parser-go "$NAME".xml > result/"$NAME".rs