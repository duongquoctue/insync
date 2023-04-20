#!/usr/bin/env bash

# chmod 755 if needed
FILE=$1

PROTO_DEF_DIR="proto"
PROTO_GEN_DIR="pkg"

if [[ "$FILE" == "" ]]; then
    FILES="$(find $PROTO_DEF_DIR -type f -name "*.proto")"
else
    PATTERN="$FILE.proto"
    FILES="$(find $PROTO_DEF_DIR -type f -name "$PATTERN")"
fi

if [[ "$FILES" == "" ]]; then
    echo "Empty files! Exiting now..."
    exit
fi


protoc $FILES \
    --go_out="$PROTO_GEN_DIR" \
    --go_opt=paths=source_relative \
    --go-grpc_out="$PROTO_GEN_DIR" \
    --go-grpc_opt=paths=source_relative \
