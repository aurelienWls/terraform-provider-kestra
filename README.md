# Terraform `Kestra` Provider

> Terraform Provider for Kestra resources.

## Using the provider

This Terraform Provider is available to install automatically via `terraform init`. It is recommended to setup the following Terraform configuration to pin the major version:

```hcl
terraform {
  required_providers =  {
    source = "kestra-io/kestra"
    version = "~> X.Y" # where X.Y is the current major version and minor version
  }
}
```

Additional documentation, including available resources and their arguments/attributes can be found on the [Terraform documentation website](https://registry.terraform.io/providers/kestra-io/kestra/latest/docs).

## Developing the Provider
If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and need a Kestra cluster.

```sh
$ make testacc
```

### Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
- [Go](https://golang.org/doc/install) >= 1.16

### Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command: 
```sh
$ go install
```

### Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.
