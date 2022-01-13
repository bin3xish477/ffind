# ffind

`ffind` is a tool to find files of interest on a compromised host during a penetration test

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
go build -ldflags '-w -s' *.go
```

### Usage

```console
ffind -p /
ffind
```

### Enable ANSI Color Escape Sequence in PowerShell

By default, PowerShell is not set to escape ANSI color so the output of `ffind` will look weird. To fix this, open a PowerShell prompt and run the following command to enable
support:

```powershell
Set-ItemProperty HKCU:\Console VirtualTerminalLevel -Type DWORD 1
```

Then, open up a new PowerShell prompt, and you should be good to go:

![image](https://user-images.githubusercontent.com/44281620/148887074-80709884-7dc8-490a-9d5a-dba6e5b7778b.png)
