# terraform-provider-awx

[![All Contributors](https://img.shields.io/github/all-contributors/adeo-opensource/terraform-provider-awx?style=flat&label=Contributors&color=informational)](#contributors)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=adeo-opensource_terraform-provider-awx&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=adeo-opensource_terraform-provider-awx)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=adeo-opensource_terraform-provider-awx&metric=coverage)](https://sonarcloud.io/summary/new_code?id=adeo-opensource_terraform-provider-awx)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=adeo-opensource_terraform-provider-awx&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=adeo-opensource_terraform-provider-awx)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=adeo-opensource_terraform-provider-awx&metric=bugs)](https://sonarcloud.io/summary/new_code?id=adeo-opensource_terraform-provider-awx)

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

```terraform
# It is strongly recommended to configure this provider through the
# environment variables described above.
provider "awx" {}
```

### Basic auth usage

```terraform
provider "awx" {
  hostname = "https://my-awx"
  username = "test"
  password = "changeme" # pragma: allowlist secret
}
```

### Token usage

```terraform
provider "awx" {
  hostname = "https://my-awx"
  token = "test"
}
```

## How to contribute? 

[Learn about how to contribute](CONTRIBUTING.md)

## Where to ask Questions?

Questions can be asked in form of issues in this repository:
[Open an issue][open-issue]

## Changelog 

[Learn about the latest improvements](CHANGELOG.md)

## License

Project is under Apache 2.0 license. See [License](LICENSE) file for more information.

## Contributors ✨

Thanks goes to these wonderful people ([emoji key][all-contributors-emoji-url]):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors][all-contributors-url] specification.  
Contributions of any kind welcome!


[all-contributors-url]: https://github.com/all-contributors/all-contributors
[all-contributors-emoji-url]: https://allcontributors.org/docs/en/emoji-key
[open-issue]: https://github.com/adeo-opensource/terraform-provider-awx/issues/new/choose
