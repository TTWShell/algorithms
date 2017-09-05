#!/bin/bash

FILE_PATH="/tmp/algorithms-coverage.out"

echo $@

go test -coverprofile=${FILE_PATH} -v --cover $@

PWD_SIZE=$(( ${#PWD} + 3 ))
sed -n '1!p' ${FILE_PATH} | cut -b ${PWD_SIZE}- > ${FILE_PATH}.tmp
echo "mode: set" > ${FILE_PATH}.ok
sed -n '1!p' ${FILE_PATH}.tmp | awk '$0="\.\/"$0' >> ${FILE_PATH}.ok

go tool cover -html=${FILE_PATH}.ok
