windows:
	GOOS=windows GOARCH=amd64 go build -o ./ffind.exe ./main.go
windows-release:
	GOOS=windows GOARCH=amd64 go build -o ./ffind.exe ./main.go
	tar -czvf ffind-windows-amd64.tar.gz ./ffind.exe ./README.md
	sha256sum ffind-windows-amd64.tar.gz > ffind-windows-amd64.tar.gz.sha256sum
	GOOS=windows GOARCH=386 go build -o ./ffind.exe ./main.go
	tar -czvf ffind-windows-386.tar.gz ./ffind.exe ./README.md
	sha256sum ffind-windows-386.tar.gz > ffind-windows-386.tar.gz.sha256sum

darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./ffind ./main.go
darwin-release:
	GOOS=darwin GOARCH=amd64 go build -o ./ffind ./main.go
	tar -czvf ffind-darwin-amd64.tar.gz ./ffind ./README.md
	sha256sum ffind-darwin-amd64.tar.gz > ffind-darwin-amd64.tar.gz.sha256sum

linux:
	GOOS=linux GOARCH=amd64 go build -o ./ffind ./main.go
linux-release:
	GOOS=linux GOARCH=amd64 go build -o ./ffind ./main.go
	tar -czvf ffind-linux-amd64.tar.gz ./ffind ./README.md
	sha256sum ffind-linux-amd64.tar.gz > ffind-linux-amd64.tar.gz.sha256sum
	GOOS=linux GOARCH=386 go build -o ./ffind ./main.go
	tar -czvf ffind-linux-386.tar.gz ./ffind ./README.md
	sha256sum ffind-linux-386.tar.gz > ffind-linux-386.tar.gz.sha256sum

releases: linux-release windows-release darwin-release

clean:
	@if [ -f ffind ]; then rm ffind; fi
	@if [ -f ffind.exe ]; then rm ffind.exe; fi
