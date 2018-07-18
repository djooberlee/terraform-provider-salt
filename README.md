# Salt Terraform Provider

A Terraform provider serving as an interop layer for an Terraform [roster
module][1] that does not exist yet.

This provider is derived and inspired by [terraform-provider-ansible][7].
Read the [introductory blog post][3] for an explanation of the design
motivations behind the original ansible provider.

## Installation

Installation can be accomplished in two different ways:

1. Installing a pre-compiled release (recommended)
2. Compiling from source

### Compiling From Source

> Note: Terraform requires Go 1.9 or later to successfully compile.

If you'd like to take advantage of features not yet available in a pre-compiled
release, you can compile `terraform-provider-salt` from source.

In order to compile, you will need to have Go installed on your workstation.
Official instructions on how to install Go can be found [here][5].

Alternatively, you can use [gimme][6] as a quick and easy way to install Go:

```shell
$ sudo wget -O /usr/local/bin/gimme https://raw.githubusercontent.com/travis-ci/gimme/master/gimme
$ sudo chmod +x /usr/local/bin/gimme
$ gimme 1.10
# copy the output to your `.bashrc` and source `.bashrc`.
```

Once you have a working Go installation, you can compile
`terraform-provider-salt` by doing the following:

```shell
$ go get github.com/nbering/terraform-provider-ansible
$ cd $GOPATH/src/github.com/nbering/terraform-provider-ansible
$ make
```

You should now have a `terraform-provider-salt` binary located at
`$GOPATH/bin/terraform-provider-salt`. Copy this binary to a designated
directory as described in Terraform's [plugin installation instructions][2]

## Terraform Configuration Example

```
resource "ansible_salt" "example" {
    hostname = "example.com"
}
```

## License

Contributions specific to this project are made available under the
[Mozilla Public License](./LICENSE).

Code under the `vendor/` directory is copyright of the various package owners,
and made available under their own license considerations.

[1]: https://docs.saltstack.com/en/latest/topics/ssh/roster.html
[2]: https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin
[3]: http://nicholasbering.ca/tools/2018/01/08/introducing-terraform-provider-ansible/
[4]: https://github.com/nbering/terraform-provider-ansible/releases
[5]: https://golang.org/doc/install
[6]: https://github.com/travis-ci/gimme
[7]: https://github.com/nbering/terraform-provider-ansible
