provider "libvirt" {
  uri = "qemu:///system"
}

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

resource "libvirt_volume" "opensuse_leap" {
  name   = "leap.iso"
  source = "openSUSE-Leap-42.3-JeOS-for-OpenStack-Cloud.x86_64.qcow2"
}

resource "libvirt_volume" "volume-web" {
  name           = "volume-web-${count.index}"
  base_volume_id = "${libvirt_volume.opensuse_leap.id}"
  count          = 2
}

resource "libvirt_volume" "volume-db" {
  name           = "volume-db-${count.index}"
  base_volume_id = "${libvirt_volume.opensuse_leap.id}"
  count          = 2
}

resource "libvirt_domain" "vm-web" {
  name   = "web-${count.index}"
  memory = 512

  disk {
    volume_id = "${element(libvirt_volume.volume-web.*.id, count.index)}"
  }

  network_interface {
    network_name   = "default"
    wait_for_lease = true
  }

  cloudinit = "${libvirt_cloudinit.common_init.id}"

  count = 2
}

resource "libvirt_domain" "vm-db" {
  name   = "db-${count.index}"
  memory = 512

  disk {
    volume_id = "${element(libvirt_volume.volume-db.*.id, count.index)}"
  }

  network_interface {
    network_name   = "default"
    wait_for_lease = true
  }

  cloudinit = "${libvirt_cloudinit.common_init.id}"

  count = 2
}

resource "salt_host" "webminion" {
  salt_id = "web${count.index}"
  host    = "${element(libvirt_domain.vm-web.*.network_interface.0.addresses[count.index], 0)}"
  user    = "root"
  count   = 2
  passwd  = ""
}

resource "salt_host" "dbminion" {
  salt_id = "db${count.index}"
  host    = "${element(libvirt_domain.vm-db.*.network_interface.0.addresses[count.index], 0)}"
  user    = "root"
  count   = 2
  passwd  = ""
}
