---
layout: "salt"
page_title: "Provider: salt"
sidebar_current: "docs-salt-index"
description: |-
  The Salt provider is a logical provider to expose terraform resources to salt-ssh using the roster.
---

# Salt provider

It requires a not yet upstreamed terraform roster module.

## Example usage

Given a simple salt-ssh tree with a `Saltfile`

```yaml
salt-ssh:
  config_dir: etc/salt
  max_procs: 30
  wipe_ssh: True
```

and `etc/salt/master`

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

In the same folder as your `Saltfile`, create terraform file with resources like cloud instances, virtual machines, etc. For every single of those that you want to manage with Salt, create a `salt_host` resource:

```hcl
resource "salt_host" "dbminion" {
  salt_id = "dbserver"
  host = "${libvirt_domain.vm-db.network_interface.0.addresses.0}"
  user = "root"
  passwd = "linux"
}
```

You can use the `count` attribute to create multiple roster entries with a single definition. You can see some [examples](https://github.com/dmacvicar/terraform-provider-salt/tree/master/examples) in the git repository.


```console
terraform apply
...
Apply complete! Resources: 14 added, 0 changed, 0 destroyed.
```

Once the resources are created, running `salt-ssh` in the directory (assuming salt already includes the companion terraform roster), will automatically use the available resources exposed as `salt_host` as part of the roster.

```console
salt-ssh '*' test.ping
db0:
    True
db1:
    True
web0:
    True
web1:
    True
```
