# ffind

`ffind` is a tool to find configuration files or files with sensitive information based on file extensions.

![image](https://user-images.githubusercontent.com/44281620/148688336-e0aea33a-8f4a-4fa4-adc3-3229531ed0f6.png)

### Installation

#### With Go

```
go install -v github.com/bin3xish477/ffind@latest
```

#### With Git

```
git clone https://github.com/bin3xish477/ffind.git
cd ffind
go build *.go
sudo install ffind /usr/local/bin
```

### Usage

```console
ffind -p /
ffind
```
