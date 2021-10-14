.PHONY: i2i-client clean

os=$(shell uname -s | tr '[:upper:]' '[:lower:]')
arch="amd64"

i2i-client:
	@GOOS=${os} GOARCH=${arch} CGO_ENABLED=0 go build ./cmd/i2i-client

install: i2i-client
	mv ./i2i-client ~/go/bin/
	~/go/bin/i2i-client cfg  completion zsh > ~/.oh-my-zsh/completions/_i2i-client

clean:
	@rm -f i2i-client

