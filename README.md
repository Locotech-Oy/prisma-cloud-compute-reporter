# Prisma Cloud Compute Reporter

This is a CLI tool for parsing the JSON output of a Prisma Cloud Compute scan report (twistcli) and converting it to other formats.

***NOTE: This tool is provided as is and without any official support. Issues will be handled on a best effort basis. Pull requests for bugfixes and new features are welcome, if they make sense to us we will attempt to incorporate them.***

## Install

Pre-built binaries are available from the releases page of the Github repository. These are not signed. Just download the raw file that matches the architecture to your system, make it executable and then run. Example one-liner for macOS that will download and extract the latest executable:

```curl
curl -s https://api.github.com/repos/Locotech-Oy/prisma-cloud-compute-reporter/releases/latest \
| grep "browser_download_url.*darwin_amd64" \
| cut -d : -f 2,3 \
| tr -d \" \
| xargs curl -L \
| tar -xz
```

after this, make the downloaded file executable:

```text
chmod +x ./prisma-cloud-compute-reporter
```

If you prefer to build your own binary, clone this repo and run go build

```text
go build -o <my-binary-name> main.go
```

Or you can use go install to install binary to global $GOPATH

```text
go install github.com/Locotech-Oy/prisma-cloud-compute-reporter@latest
```

## Usage

Running the tool with no arguments or with the -h flag will show a simple command help

```
./prisma-cloud-compute-reporter -h
```

Parsing image scan result and including both compliance and vulnerabilities into the same junit report

```text
./prisma-cloud-compute-reporter image parse -o pcc-junit-report.xml pcc_pipeline-scan_sample.json
```

Use the following flags to adjust output:

| Flag                  | Description                       |
|---                    |---                                |
| --compliance          | Include compliance results        |
| --vulnerabilities     | Include vulnerabilities results   |

If neither ```--compliance``` or ```--vulnerabilities``` is provided all reports are merged into one file.

## Development

### Tests

Use standard golang test commands from the root of the project to run unit tests.

Run all tests:

```text
go test ./...
```

Run tests including coverage and open result in browser:

```text
go test -coverprofile=c.out -coverpkg=./... ./... && go tool cover -html=c.out -o coverage.html
open coverage.html
```
