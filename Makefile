# cmd for debug cat -e -t -v Makefile

# Install all dependents.
install:
	go get -t -v ./...

# Test all code.
test-all:
	go test -v -cover ./...

# test one or multi, example: make test "FILES=leetcode/array/trap*.go"
test:
	go test -count=1 -v -cover ${FILES}

# show cover in browser, example: make cover "FILES=leetcode/array/trap*.go"
cover:
	./test.sh ${FILES}
