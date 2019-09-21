#!/bin/bash

XC_ARCH=${XC_ARCH:-386 amd64}
XC_OS=${XC_OS:-linux darwin windows}
DIR=$(cd $(dirname ${0}) && pwd)
cd ${DIR}

rm -rf pkg/
gox \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -output "pkg/{{.OS}}_{{.Arch}}/{{.Dir}}"

mkdir -p ./pkg/dist

for SUBDIR in $(find ./pkg -mindepth 1 -maxdepth 1 -type d); do
    PLATFORM=$(basename ${SUBDIR})
    if [ $PLATFORM == "dist" ]; then
        continue
    fi
    ARCHIVE_NAME=${VERSION}-${PLATFORM}
    cd $SUBDIR
    zip ${DIR}/pkg/dist/${NAME}-${VERSION}-${PLATFORM}.zip ./*
    cd $DIR
done

for ARCHIVE in ./pkg/dist/*; do
    ARCHIVE_NAME=$(basename ${ARCHIVE})
    echo Uploading: ${ARCHIVE_NAME}
done
