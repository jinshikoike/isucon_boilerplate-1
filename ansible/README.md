# Ansible 

# setup
- `ansible-galaxy install -r requirements.yml`

## add your personal settings to `group_vars/all.yml`
e.g.

- group_vars/secret_personal_var.yml

```
# Example
---
user:
  name: jinshi
  pub_key: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDH+HHxNMMiGXIB4hwE7s8u3hdUiW37ON/XApnjiIStgPP6uA0+jhvvoAXqvV7Taxkp4sXXsbSsrPvIaEHoIS6m9HwjrUebvLsWhHohzFh/GG2GUAwYf36RURRwNX11g+9iBQtFmmBCELSRXjbIG48rQcmJS/CleHWKkmBJRhX1T48pys9HKTGOECJS63h4nfNQ5tSaKL7DQERimyOxMztg57mZsT2Gt/pvEjYL6/r1qEYUC2D758bdDGUGPSc2T1AYMAsmAOvvXlyPftd4g49qecqVJe8Bl23UGc5oKQLV5quV1S2rMTH9XtYhp0/QhmnagBLJ5zYPjHFjeUdM1b3H ansible-generated on localhost.localdomain"
  git_repos:
    - repo: "https://github.com/jinshikoike/dotfile.git"
      dest: "/.dotfiles"
    - repo: "https://github.com/jinshikoike/dotfile.git"

  # Each Below Item is appended to bash_profile.
  bash_profiles: 
    - export TES_TES="testestes"
    - export TES1_TES1="tes1tes1tes1"

  run_commands:
    - echo $TES_TES
```

## add your personal files you want to upload
if you put file in `files/secrets` directory, all the files in that directory will be uploaded `/home/[LOGIN_USER_NAME]/secrets` on remote host.
- note: `files/secrets` directory is gitignored. so this directory can not be seen by another user.

# With Vagrant

`vagrant up`

## provision 

`vagrant provision`

## debug in vagrant
vagrant provision is Too slow for some reason.
So you should to debug playbook on the vm on the vagrnat.

- `vagrant ssh`
- `cd path/to/playbook/directory`
- `ansible-playbook -i hosts/development [playbook name].yml --connection=local`

## site.yml
this playbook call all other playbooks(common,nginx,redis,mysql... etc)

## other playbooks (nginx,redis,mysql..)
I don't write explanation. please see playbook yourself.

# provision hosts which is not local machine
- Edit `hosts/[env name you want to provision]` to add some host.
- Execute `ansible-playbook -i hosts [playbook name].yml`

