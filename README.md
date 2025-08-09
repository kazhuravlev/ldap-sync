# LDAP Sync service

Thin server that keep LDAP group hierarchy in sync. Write a simple group tree in Yaml and `ldap-sync` will apply this to
LDAP server.

## Features

- [ ] Create groups from Yaml to LDAP server
- [ ] Support for LDAPS
- [ ] Remove groups that was added without `ldap-sync`
- [ ] Auto populate group description from Yaml

## Use cases

### You want to have group hierarchy in LDAP and no worry about all other software supports it

Some software like Jira can use nested groups to manage permissions. Ypu can use FreeIPA to manage groups and its
hierarchy.

But some other software is unable to use group hierarchy and in that case you should create a special flow to add some
permissions from ldap to this software.

In order to solve this - you can use flat groups and add each group to user directly. But someone should manage
hierarchy from the outside. This project created to solve this issue. You can use https://github.com/lldap/lldap that
not support nested groups and manage hierarchy with `ldap-sync`.
