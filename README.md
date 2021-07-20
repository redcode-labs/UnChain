<h1 align="center"> UnChain </h1> <br>
<p align="center">
  <a>
    <img src="img.png" width="500" height="600">
  </a>
</p>

<p align="center">
  A tool to find redirection chains in multiple URLs
</p>

## Introduction
UnChain automates process of finding and following `30X` redirects by extracting "Location" header of HTTP responses.

### Building
To build UnChain simple run:
```
go build -o unchain ./cmd/main.go
```

## Usage

```
usage: unchain [-h|--help] -u|--url "<value>"

arguments:

  -h  --help  Print help information
  -u  --url   File containing urls or a single url
```
