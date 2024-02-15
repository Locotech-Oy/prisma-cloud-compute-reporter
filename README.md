# Prisma Cloud Compute Reporter

This is a CLI tool for parsing the JSON output of a Prisma Cloud Compute scan report (twistcli) and converting it to other formats.

***NOTE: This tool is provided as is and without any official support. Issues will be handled on a best effort basis. Pull requests for bugfixes and new features are welcome, we will try to include them as quickly as possible. Please see the [contribution](CONTRIBUTING.md) guidelines.***

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

### Docker image

The tool has been packaged into a docker image as well, available at <https://hub.docker.com/r/locotech/pcc-reporter>. The docker image is kept up to date with the main branch. The available commands and options are the same as for the pre-built binaries, however if you wish to use the latest version as based on the latest commits to the main branch, use the docker image.

For usage instructions, simply run

```text
docker run --rm locotech/pcc-reporter:latest
```

## Usage

Running the tool with no arguments or with the -h flag will show a simple command help

```bash
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

## JUnit file format gotchas

Due to the way Prisma Cloud produces it's scan reports, the JUnit report may not be quite intuitive, especially if tracking passed/failed rates and trends. Prisma Cloud only reports on vulnerabilities and compliance issues it detects, so the junit report we produce will always contain only failing or skipped tests. Thus a situation where Prisma Cloud detects zero issues, meaning everything is ok, would produce a junit report with 0% passed tests since the report would not contain any test cases at all. The following rules apply to the junit report:

1. If Prisma Cloud has been set up to only report, but not fail a scan when an issue is detected, the junit report will report this as a skipped test.
2. If Prisma Cloud has been set up to fail scans when an issue is detected, the junit report will report the issue as a failed test.

## CI/CD integrations

### Azure pipeline

Below is a example snippet for downloading the latest release of this tool and running it in an Azure pipeline. This assumes Twistcli has been run to produce a JSON formatted scan report that can be accessed in the pipeline workspace. It parses the report into two different outputs, in order to separate the junit tests into vulnerabilities and compliance results and publish them separately for better overview.

```yaml

...

    - task: DownloadGitHubRelease@0
      displayName: Download PCC reporter
      inputs:
        connection: '<some-github-connection>'
        userRepository: 'Locotech-Oy/prisma-cloud-compute-reporter'
        defaultVersionType: 'latest'
        itemPattern: '*linux_amd64.tar.gz'
        downloadPath: '$(System.ArtifactsDirectory)'

    - task: CmdLine@2
      displayName: 'PCC report parsing'
      inputs:
        failOnStderr: false
        script: |

          PCC_ARCHIVE=$(find $(System.ArtifactsDirectory) -maxdepth 1 -type f -iname "*.tar.gz" | head -1)
          echo $PCC_ARCHIVE
          mkdir -p ./pcc-reporter
          tar -xvzf $PCC_ARCHIVE -C ./pcc-reporter
          ./pcc-reporter/prisma-cloud-compute-reporter image parse --vulnerability -o ./twistcli-reporting/pcc-junit-report_vulnerabilities.xml ./twistcli-reporting/pc-scan-report.json
          ./pcc-reporter/prisma-cloud-compute-reporter image parse --compliance -o ./twistcli-reporting/pcc-junit-report_compliance.xml ./twistcli-reporting/pc-scan-report.json



    - task: PublishTestResults@2
      inputs:
        testResultsFormat: 'JUnit'
        testResultsFiles: '**/twistcli-reporting/pcc-junit-report_compliance.xml'
        testRunTitle: 'PCC compliance scan results'
    - task: PublishTestResults@2
      inputs:
        testResultsFormat: 'JUnit'
        testResultsFiles: '**/twistcli-reporting/pcc-junit-report_vulnerabilities.xml'
        testRunTitle: 'PCC vulnerability scan results'

...

```

## Development

### Tests

Use standard golang test commands from the root of the project to run unit tests.

Run all tests:

```bash
go test ./...
```

Run tests including coverage and open result in browser:

```bash
go test -coverprofile=c.out -coverpkg=./... ./... && go tool cover -html=c.out -o coverage.html
open coverage.html
```

### Build docker image locally

To build the docker image locally, run

```bash
docker build -t locotech/pcc-reporter:local .
```
