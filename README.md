# Prisma Cloud Compute Reporter

This is a CLI tool for parsing the JSON output of a Prisma Cloud Compute scan report (twistcli) and converting it to other formats.

***NOTE: This tool is provided as is and without any official support. Issues will be handled on a best effort basis. Pull requests for bugfixes and new features are welcome, if they make sense to us we will attempt to incorporate them.***

## Install

Pre-built binaries are available in the ```bin/``` folder for major platforms. These are not signed. Just download the raw file that matches the architecture to your system, make it executable and then run. Example one-liner for macOS:

```curl
curl -o pcc-reporter https://raw.github.com/xyz
```

If you prefer to build your own binary, just clone this repo and run

```text
go build -o <my-binary-name> main.go
```

Or you can use go install

```text
go install https://xyz@latest
```
