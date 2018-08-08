# Salt Terraform Provider

![alpha](https://img.shields.io/badge/stability%3F-beta-yellow.svg) [![Build Status](https://travis-ci.org/dmacvicar/terraform-provider-salt.svg?branch=master)](https://travis-ci.org/dmacvicar/terraform-provider-salt) [![Coverage Status](https://coveralls.io/repos/github/dmacvicar/terraform-provider-salt/badge.svg?branch=master)](https://coveralls.io/github/dmacvicar/terraform-provider-salt?branch=master)

A Terraform provider serving as an interop layer for an Terraform [roster
module](https://docs.saltstack.com/en/latest/topics/ssh/roster.html) that is (not upstream yet](https://github.com/saltstack/salt/pull/48873).

This provider is derived from and inspired by [terraform-provider-ansible](https://github.com/nbering/terraform-provider-ansible).
Read the [introductory blog post](http://nicholasbering.ca/tools/2018/01/08/introducing-terraform-provider-ansible/) for an explanation of the design
motivations behind the original ansible provider.

## Table of Content
- [Downloading](#Downloading)
- [Installing](#Installing)
- [Quickstart](#using-the-provider)
- [Building from source](#building-from-source)

## Downloading

Builds for openSUSE, CentOS, Ubuntu, Fedora are created with openSUSE's [OBS](https://build.opensuse.org). The build definitions are available for both the [stable](https://build.opensuse.org/package/show/home:dmacvicar:terraform-provider-salt:stable/terraform-provider-salt) and [master](https://build.opensuse.org/project/show/home:dmacvicar:terraform-provider-salt) branches.

## Using released builds

* Head to the [releases section](https://github.com/dmacvicar/terraform-provider-salt/releases) and download the latest stable release build for your distribution.

## Using unstable builds

* Head to the [download area of the OBS project](https://download.opensuse.org/repositories/home:/dmacvicar:/terraform-provider-salt/) and download the build for your distribution.

## Building from source

This project uses [glide](https://github.com/Masterminds/glide) to vendor all its
dependencies.

You do not have to interact with `glide` since the vendored packages are **already included in the repo**.

Ensure you have the latest version of Go installed on your system, terraform usually
takes advantage of features available only inside of the latest stable release.

```console
go get github.com/dmacvicar/terraform-provider-salt
cd $GOPATH/src/github.com/dmacvicar/terraform-provider-salt
go install
```

You will now find the binary at `$GOPATH/bin/terraform-provider-salt`.

# Installing

[Copied from the Terraform documentation](https://www.terraform.io/docs/plugins/basics.html):
> To install a plugin, put the binary somewhere on your filesystem, then configure Terraform to be able to find it. The configuration where plugins are defined is ~/.terraformrc for Unix-like systems and %APPDATA%/terraform.rc for Windows.

## Using the provider

## Terraform Configuration Example

```hcl
resource "libvirt_domain" "domain" {
  name = "domain-${count.index}"
  memory = 1024
  disk {
       volume_id = "${element(libvirt_volume.volume.*.id, count.index)}"
  }

  network_interface {
    network_name = "default"
    hostname = "minion${count.index}"
    wait_for_lease = 1
  }
  cloudinit = "${libvirt_cloudinit.init.id}"
  count = 2
}

resource "salt_host" "example" {
    host = "${libvirt_domain.domain.network_interface.0.addresses.0}
}
```

See [more advanced examples](examples/).

## Authors

* Duncan Mac-Vicar P. <dmacvicar@suse.de>

See also the list of [contributors](https://github.com/dmacvicar/terraform-provider-salt/graphs/contributors) who participated in this project.

This provider is derived/forked from [terraform-provider-ansible](https://github.com/nbering/terraform-provider-ansible).

## License

Contributions specific to this project are made available under the
[Mozilla Public License](./LICENSE).

Code under the `vendor/` directory is copyright of the various package owners,
and made available under their own license considerations.

