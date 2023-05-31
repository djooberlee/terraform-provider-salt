# Salt Terraform Provider

![alpha](https://img.shields.io/badge/stability%3F-beta-yellow.svg) [![Tests](https://github.com/djooberlee/terraform-provider-salt/actions/workflows/test.yml/badge.svg)](https://github.com/djooberlee/terraform-provider-salt/actions/workflows/test.yml)

A Terraform provider serving as an interop layer for an Terraform [roster
module](https://docs.saltstack.com/en/latest/topics/ssh/roster.html) that is (not upstream yet](https://github.com/saltstack/salt/pull/48873).

This provider is derived from and inspired by [terraform-provider-ansible](https://github.com/nbering/terraform-provider-ansible).
Read the [introductory blog post](http://nicholasbering.ca/tools/2018/01/08/introducing-terraform-provider-ansible/) for an explanation of the design
motivations behind the original ansible provider.

## Getting started

The provide is available for auto-installation from the [Terraform Registry](https://registry.terraform.io/providers/djooberlee/salt/latest).

In your `main.tf` file, specify the version you want to use:


```hcl
terraform {
  required_providers {
    salt = {
      source = "djooberlee/salt"
      version = "0.0.2"
    }
  }
}
```

And now run terraform init:

```
$ terraform init
```

### Binding Salt hosts to Terraform resources

For example, to bind a Salt host to a virtual machine created with the libvirt provider:

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
    host = "${libvirt_domain.domain.network_interface.0.addresses.0}"
}
```

### Setting up Salt

The goal is to create a self-contained folder where you will store both the terraform file describing the infrastructure and the Salt states to configure them.

```
.
├── etc
│   └── salt
│       ├── master
│       └── pki
│           └── master
│               └── ssh
│                   ├── salt-ssh.rsa
│                   └── salt-ssh.rsa.pub
├── main.tf
├── Saltfile
└── srv
    ├── pillar
    │   ├── terraform.sls
    │   └── top.sls
    └── salt
        ├── master
        │   └── init.sls
        ├── minion
        │   └── init.sls
        ├── minion-ssh
        │   └── init.sls
        └── top.sls
```

* As Salt will create several files once you run it, make sure your `.gitignore` is good enough to avoid checking in generated files:

```
terraform.tfstate*
*.qcow2
var
.terraform
```

* `Saltfile` should point salt to the local folder configuration:

```yaml
salt-ssh:
  config_dir: etc/salt
  max_procs: 30
  wipe_ssh: True
```

* `etc/salt/master` should let `salt-ssh` know that the states and pillar are also stored in the same folder, and should enable the terraform roster.

```yaml
root_dir: .
file_roots:
  base:
    - srv/salt
pillar_roots:
  base:
    - srv/pillar
roster: terraform
```

*NOTE*: The roster module may not [be upstream yet](https://github.com/saltstack/salt/pull/48873).

### Giving `salt-ssh` access to terraform resources via ssh

Salt by default uses the keys in `etc/salt/pki/master`. You can pre-generate those with `ssh-keygen`.

For this, you can use something like cloud-init to pre-configure your Terraform resources to pre-authorize the salt-ssh key. With [terraform-provider-libvirt](https://github.com/djooberlee/terraform-provider-libvirt) you can achieve this by using a cloud-init resource:

```hcl
resource "libvirt_cloudinit" "common_init" {
  name = "test-init.iso"
  user_data = <<EOF
#cloud-config
disable_root: 0
ssh_pwauth:   1
users:
  - name: root
    ssh-authorized-keys:
      - ${file("etc/salt/pki/master/ssh/salt-ssh.rsa.pub")}
EOF
}
```

And then referencing this resource from each virtual machine:

```hcl
  cloudinit = "${libvirt_cloudinit.common_init.id}"
```

For AWS resources, you can pass the cloud-init configuration using `user_data` ([Documentation](https://www.terraform.io/docs/providers/template/d/cloudinit_config.html)).

### Passing information from Terraform to Salt via Pillar

Sometimes you need to use infrastructure data in the Salt states. For example, the amount of resources of certain type or the ip address of some resource. For this you can put it into the [pillar](https://docs.saltstack.com/en/latest/topics/tutorials/pillar.html).

We would like to add some pillar integration at the resource level later. For now you can use `local_file` resources to write pillar sls files:

```hcl
resource "local_file" "pillar_database_cluster" {
  filename = "${path.module}/srv/pillar/terraform_database_cluster.sls"
  content = <<EOF
terraform:
  database_master_ip: ${salt_host.master.host}
EOF
}
```

Then include this pillar in the virtual machines that should receive it by editing `srv/pillar/top.sls`:

```yaml
base:
  'dbslave*':
    - terraform_database_cluster
```

As this pillar file is generated, make sure you include it in `.gitignore`:

```
terraform.tfstate*
*.qcow2
var
.terraform
srv/pillar/terraform.sls
```

See [more advanced examples](examples/).

### Testing that everything works

If everything is in place, you can start managing the resources with Salt:

```console
salt-ssh '*' test.ping
vm0:
    True
vm1:
    True
vm2:
    True
```

You can also run `salt-ssh '*' pillar.items` to check the machines receive the right pillar data, and `salt-ssh '*' state.apply` to apply the state.

## Manual installation

You can also manually download the provider from the [releases section](https://github.com/djooberlee/terraform-provider-salt/releases) on Github. To install it, refer to the [Terraform documentation](https://www.terraform.io/docs/cli/config/config-file.html#provider-installation).

You will now find the binary at `$GOPATH/bin/terraform-provider-salt`.

## Building from source

Ensure you have the latest version of Go installed on your system, terraform usually
takes advantage of features available only inside of the latest stable release.

```console
go get github.com/djooberlee/terraform-provider-salt
cd $GOPATH/src/github.com/djooberlee/terraform-provider-salt
make
```

## Authors

* Duncan Mac-Vicar P. <duncan@mac-vicar.eu>

See also the list of [contributors](https://github.com/djooberlee/terraform-provider-salt/graphs/contributors) who participated in this project.

This provider is derived/forked from [terraform-provider-ansible](https://github.com/nbering/terraform-provider-ansible).

## License

Contributions specific to this project are made available under the
[Mozilla Public License](./LICENSE).

Code under the `vendor/` directory is copyright of the various package owners,
and made available under their own license considerations.

