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