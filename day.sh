#!/bin/bash

set -e

DAY="$1"

if [ -z "$DAY" ]; then
    echo "Usage: $0 <day>"
    exit 1
fi

mkdir $DAY
cp  -r 'template/.' $DAY
echo "Good luck!"
