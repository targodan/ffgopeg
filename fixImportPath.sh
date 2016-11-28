#!/bin/bash

branch=
if [[ "$1" != "" ]]; then
    branch="$1"
else
    branch=`git branch | grep '*' | cut -d' ' -f2`
fi

from=
to=
if [[ "$branch" == "master" || "$branch" == "release"* ]]; then
    from='"github\.com\/targodan\/ffgopeg\/([^"]+)"'
    to='"gopkg\.in\/targodan\/ffgopeg.v1\/\1"'
else
    from='"gopkg\.in\/targodan\/ffgopeg.v1\/([^"]+)"'
    to='"github\.com\/targodan\/ffgopeg\/\1"'
fi

find . -name '*.go' -type f -exec sed -i -E 's/'"$from"'/'"$to"'/g' {} \;
