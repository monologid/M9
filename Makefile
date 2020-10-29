default: build-darwin-amd64 build-linux-amd64 build-windows-amd64

dev:
	air -c air.toml

build-darwin-amd64:
	env GOOS=darwin GOARCH=amd64 go build -o ./tmp/m9-darwin-amd64 . && cd ./tmp && zip -r m9-darwin-amd64.zip m9-darwin-amd64

build-dragonfly-amd64:
	env GOOS=dragonfly GOARCH=amd64 go build -o ./tmp/m9-dragonfly-amd64 . && cd ./tmp && zip -r m9-dragonfly-amd64.zip m9-dragonfly-amd64

build-freebsd-386:
	env GOOS=freebsd GOARCH=386 go build -o ./tmp/m9-freebsd-386 . && cd ./tmp && zip -r m9-freebsd-386.zip m9-freebsd-386

build-freebsd-amd64:
	env GOOS=freebsd GOARCH=amd64 go build -o ./tmp/m9-freebsd-amd64 . && cd ./tmp && zip -r m9-freebsd-amd64.zip m9-freebsd-amd64

build-freebsd-arm:
	env GOOS=freebsd GOARCH=arm go build -o ./tmp/m9-freebsd-arm . && cd ./tmp && zip -r m9-freebsd-arm.zip m9-freebsd-arm

build-linux-386:
	env GOOS=linux GOARCH=386 go build -o ./tmp/m9-linux-386 . && cd ./tmp && zip -r m9-linux-386.zip m9-linux-386

build-linux-amd64:
	env GOOS=linux GOARCH=amd64 go build -o ./tmp/m9-linux-amd64 . && cd ./tmp && zip -r m9-linux-amd64.zip m9-linux-amd64

build-linux-arm:
	env GOOS=linux GOARCH=arm go build -o ./tmp/m9-linux-arm . && cd ./tmp && zip -r m9-linux-arm.zip m9-linux-arm

build-linux-arm64:
	env GOOS=linux GOARCH=arm64 go build -o ./tmp/m9-linux-arm64 . && cd ./tmp && zip -r m9-linux-arm64.zip m9-linux-arm64

build-linux-ppc64:
	env GOOS=linux GOARCH=ppc64 go build -o ./tmp/m9-linux-ppc64 . && cd ./tmp && zip -r m9-linux-ppc64.zip m9-linux-ppc64

build-linux-ppc64le:
	env GOOS=linux GOARCH=ppc64le go build -o ./tmp/m9-linux-ppc64le . && cd ./tmp && zip -r m9-linux-ppc64le.zip m9-linux-ppc64le

build-linux-mips:
	env GOOS=linux GOARCH=mips go build -o ./tmp/m9-linux-mips . && cd ./tmp && zip -r m9-linux-mips.zip m9-linux-mips

build-linux-mipsle:
	env GOOS=linux GOARCH=mipsle go build -o ./tmp/m9-linux-mipsle . && cd ./tmp && zip -r m9-linux-mipsle.zip m9-linux-mipsle

build-linux-mips64:
	env GOOS=linux GOARCH=mips64 go build -o ./tmp/m9-linux-mips64 . && cd ./tmp && zip -r m9-linux-mips64.zip m9-linux-mips64

build-linux-mips64le:
	env GOOS=linux GOARCH=mips64le go build -o ./tmp/m9-linux-mips64le . && cd ./tmp && zip -r m9-linux-mips64le.zip m9-linux-mips64le

build-netbsd-386:
	env GOOS=netbsd GOARCH=386 go build -o ./tmp/m9-netbsd-386 . && cd ./tmp && zip -r m9-netbsd-386.zip m9-netbsd-386

build-netbsd-amd64:
	env GOOS=netbsd GOARCH=amd64 go build -o ./tmp/m9-netbsd-amd64 . && cd ./tmp && zip -r m9-netbsd-amd64.zip m9-netbsd-amd64

build-netbsd-arm:
	env GOOS=netbsd GOARCH=arm go build -o ./tmp/m9-netbsd-arm . && cd ./tmp && zip -r m9-netbsd-arm.zip m9-netbsd-arm

build-openbsd-386:
	env GOOS=openbsd GOARCH=386 go build -o ./tmp/m9-openbsd-386 . && cd ./tmp && zip -r m9-openbsd-386.zip m9-openbsd-386

build-openbsd-amd64:
	env GOOS=openbsd GOARCH=amd64 go build -o ./tmp/m9-openbsd-amd64 . && cd ./tmp && zip -r m9-openbsd-amd64.zip m9-openbsd-amd64

build-openbsd-arm:
	env GOOS=openbsd GOARCH=arm go build -o ./tmp/m9-openbsd-arm . && cd ./tmp && zip -r m9-openbsd-arm.zip m9-openbsd-arm

build-solaris-amd64:
	env GOOS=solaris GOARCH=amd64 go build -o ./tmp/m9-solaris-amd64 . && cd ./tmp && zip -r m9-solaris-amd64.zip m9-solaris-amd64

build-windows-386:
	env GOOS=windows GOARCH=386 go build -o ./tmp/m9-windows-386 . && cd ./tmp && zip -r m9-windows-386.zip m9-windows-386

build-windows-amd64:
	env GOOS=windows GOARCH=amd64 go build -o ./tmp/m9-windows-amd64 . && cd ./tmp && zip -r m9-windows-amd64.zip m9-windows-amd64