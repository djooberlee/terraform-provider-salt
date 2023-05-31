---
layout: "salt"
page_title: "Salt: salt_host"
sidebar_current: "docs-salt-host"
description: |-
  Manages a host exposed to the Salt SSH roster
---

# salt\_host

Manages a resource exposed as a roster item to [salt-ssh](https://docs.saltstack.com/en/latest/topics/ssh/) via the [roster](https://docs.saltstack.com/en/latest/topics/ssh/roster.html).

## Example Usage

Single host:

```hcl
resource "salt_host" "example" {
  salt_id = "minion0"
  host = "${libvirt_volume.domain.foo.network_interface.0.addresses.0}"
  user = "root"
}
```

Multiple instances using `count`:

```hcl
resource "salt_host" "example" {
  salt_id = "minion${count.index}"
  host = "${element(libvirt_volume.domain.*.id, count.index)}.network_interface.0.addresses.0 ${libvirt_domain.domain.network.ip}
  user = "root"
  count = 2
}
```

## Argument Reference

The provider supports the same attributes as the [roster](https://docs.saltstack.com/en/latest/topics/ssh/roster.html).

The following arguments are supported:

* `salt_id` - (Required) The id to reference the target system (minion) with
* `host` - (Required) The IP address or DNS name of the remote host
* `user` - (Required) The user to log in as
* `passwd` - The password to log in with
* `port` - The target system's ssh port number
* `sudo` - Boolean to run command via sudo
* `sudo_user` - Str: Set this to execute Salt as a sudo user other than root.
   This user must be in the same system group as the remote user
   that is used to login and is specified above. Alternatively,
   the user must be a super-user.
* `tty` - Set this option to true if sudo is also set to
   True and requiretty is also set on the target system
* `priv` - File path to ssh private key, defaults to salt-ssh.rsa
   The priv can also be set to agent-forwarding to not specify
   a key, but use ssh agent forwarding
* `timeout` - Number of seconds to wait for response when establishing
   an SSH connection
* `minion_opts` - Dictionary of minion opts
* `thin_dir` - The target system's storage directory for Salt
   components. Defaults to /tmp/salt-<hash>.
* `cmd_umask` - umask to enforce for the salt-call command. Should be in
   octal (so for 0o077 in YAML you would do 0077, or 63)
