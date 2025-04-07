all: clean build

.PHONY: clean
clean:
	rm -rf build

.PHONY: build
build:
	go build -ldflags="-w -s" -o build/ .

dev:
	go run .

.PHONY: test-manual
test-manual: clean build
	echo "=====[TEST HELP]====="
	./build/daylight --help
	echo "=====[TEST SHORT]====="
	./build/daylight --short
	echo "=====[TEST SHORT, XMAS DAY]====="
	./build/daylight --short --date="2025-12-25"
	echo "=====[TEST SHORT, POLAR DAY]====="
	./build/daylight --short --date="2025-12-25" --latitude="-89.925" --longitude "1.427"
	echo "=====[TEST POLAR DAY]====="
	./build/daylight --latitude -90 --longitude 0 --date="2025-01-02"
	echo "=====[TEST POLAR NIGHT]====="
	./build/daylight --latitude="82.4" --longitude -14.3 --date="2025-01-02"
	echo "=====[TEST CAPE TOWN]====="
	./build/daylight --date="2025-02-03" --latitude="-33.92" --longitude 18.42 --timezone "Africa/Johannesburg"
	echo "=====[TEST NOCOLOR]====="
	NO_COLOR=true ./build/daylight
	echo "=====[TEST HERE/NOW]====="
	./build/daylight

.PHONY: test
test:
	go test ./internal
