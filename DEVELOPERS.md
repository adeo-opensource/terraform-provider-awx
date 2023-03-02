# How to develop

## Build locally the provider

### Needed

* [GoReleaser](https://goreleaser.com/install/)
* [pre-commit](https://pre-commit.com/#install)
* Python 3

### Prerequisites

Install pre-commit

```sh
pre-commit install
```

### Build the provider

```sh
goreleaser build --snapshot --rm-dist
```

### Copy Provider

Copy the provider to user's `~/.terraform.d` folder.
> Important: if building the provider in an operating system other than Linux x86_64.
> Adjust the paths below replacing `linux_amd64` with the corresponding platform code.  
> E.g.: `darwin_amd64` for macOS.

```sh
mkdir -p ~/.terraform.d/plugins/github.com/adeo-opensource/awx/0.1/linux_amd64/
find ./dist/terraform-provider-awx_linux_amd64_v1/* -name 'terraform-provider-awx*' -print0 | xargs -0 -I {} mv {} ~/.terraform.d/plugins/github.com/adeo-opensource/awx/0.1/linux_amd64/
```

### Run tests and ensure they're all passing

```sh
go test ./test -count=1
```

### Use it locally

```terraform
terraform {
  required_providers {
    awx = {
      source  = "github.com/adeo-opensource/awx"
      version = "0.1"
    }
  }
}
```

