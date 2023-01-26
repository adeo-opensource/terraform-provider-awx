# terraform-provider-awx

[![All Contributors](https://img.shields.io/github/all-contributors/adeo-opensource/terraform-provider-awx?style=flat&label=Contributors&color=informational)](#contributors)

## Description

The AWX provider allow Terraform to read from, write to, and configure AWX.

## Roadmap

[Resources managed by the provider](ROADMAP.md)

## AWX authentication configuration options

There are environment variables that allow you to authenticate yourself:

* AWX_HOSTNAME
* AWX_USERNAME
* AWX_PASSWORD
* AWX_TOKEN

## Example Usage

```
# It is strongly recommended to configure this provider through the
# environment variables described above.
provider "awx" {}
```

### Basic auth usage

```
provider "awx" {
  hostname = "https://my-awx"
  username = "test"
  password = "changeme"
}
```

### Token usage

```
provider "awx" {
  hostname = "https://my-awx"
  token = "test"
}
```

## How to contribute? 

[Learn about how to contribute](CONTRIBUTING.md)

## Changelog 

[Learn about the latest improvements](CHANGELOG.md)

## License

Project is under Apache 2.0 license. See [License](LICENSE) file for more information.

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
