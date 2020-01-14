#!/bin/bash
# ./test.sh ./leetcode/array
set -e

FILE_PATH="/tmp/algorithms-coverage.out"

echo $@

go test -coverprofile=${FILE_PATH} -v --cover $@

PWD_SIZE=$(( ${#PWD} + 2 ))
echo ${PWD_SIZE}
sed -n '1!p' ${FILE_PATH} | cut -b ${PWD_SIZE}- > ${FILE_PATH}.tmp
echo "mode: set" > ${FILE_PATH}.ok
sed -n '1!p' ${FILE_PATH}.tmp | awk '$0="\.\/"$0' >> ${FILE_PATH}.ok
echo ${FILE_PATH}

go tool cover -html=${FILE_PATH}.ok
