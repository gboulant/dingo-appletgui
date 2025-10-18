all:

test:
	@make -C demos/guiapp build

clean:
	@make -C demos/guiapp $@
	@rm -f output.*

cov:
	@go test -coverprofile=output.cov
	@go tool cover -func=output.cov

doc:
	@go tool doc -all
	@go tool doc -C demos/guiapp -cmd -all

gomod.update:
	@rm -f go.mod go.sum
	@go mod init dynsound
	@go mod tidy

gomod.dev:
	@go mod edit -replace github.com/gboulant/dingo-applet=../applet
	@go mod edit -replace github.com/gboulant/dingo-stdrw=../stdrw

gomod.std:
	@go mod edit -dropreplace github.com/gboulant/dingo-applet
	@go mod edit -dropreplace github.com/gboulant/dingo-stdrw
