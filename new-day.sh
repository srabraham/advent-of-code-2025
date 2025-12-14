#!/usr/bin/env bash

day="day${1}"

script_dir=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd "${script_dir}"

cp -r day00 ${day}
cd ${day}
mv day00.go ${day}.go
mv day00_test.go ${day}_test.go

sed -i '' "s/day00/${day}/g" *


