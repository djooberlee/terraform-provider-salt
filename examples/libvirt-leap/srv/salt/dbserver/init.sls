postgresql-server:
  pkg.installed: []

postgresql:
  service.running:
    - enable: True
    - require:
        - pkg: postgresql-server

