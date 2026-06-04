# Ansible deployment for Projeto Korp V4

This directory contains an automated deployment workflow for the Projeto Korp stack.
The playbook is intended for local execution on Fedora hosts.

## Contents

- `ansible.cfg` - Ansible configuration for local execution
- `inventory.ini` - inventory targeting `localhost`
- `group_vars/all.yml` - common deployment variables
- `playbook.yml` - steps to install Docker, create network, deploy the stack, and validate it

## How it works

The playbook performs these actions:
1. Updates system packages on Fedora
2. Installs Docker CE, Docker Compose plugin, and required packages
3. Ensures Docker is running and the user is part of the `docker` group
4. Creates the Docker network `korp-network`
5. Validates required project files exist
6. Builds and starts the stack with `docker compose up -d --build`
7. Waits for ports `80`, `9090`, and `3000`
8. Validates the application endpoint

## Run the playbook

From the repository root:

```bash
ansible-playbook ansible/playbook.yml
```

This uses the root `ansible.cfg`, so you do not need to pass `ANSIBLE_CONFIG`, `-i`, `--connection`, or `--ask-become-pass` manually.

## Notes

- This playbook is written for Fedora and uses `dnf`.
- It expects to run on the same host where the repository is available.
- The Ansible run may require `sudo` privileges to install Docker and start services.
