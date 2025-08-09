# LDAP Sync Service

A lightweight server that keeps your LDAP group hierarchy in sync.
Define a simple group tree in YAML, and `ldap-sync` will apply it to your LDAP server automatically.

## Features

- [ ] Create groups in LDAP from a YAML definition
- [ ] Support for LDAPS (secure LDAP)
- [ ] Remove unmanaged groups that were added without `ldap-sync`
- [ ] Automatically populate group descriptions from YAML

## Use Cases

### Maintain an LDAP group hierarchy without worrying about software compatibility

Some software, like Jira, supports nested LDAP groups to manage permissions.
You can use FreeIPA to manage both groups and their hierarchy.

However, other software cannot handle nested groups. In those cases, you might resort to **flattening** your groups and
adding them directly to each user. While this works, you still need a way to manage the hierarchy externally.

That’s where `ldap-sync` comes in. It lets you maintain a proper hierarchy in your configuration while applying a
flattened structure to LDAP where needed.

For example, you can use it with [lldap](https://github.com/lldap/lldap), which does not support nested groups, and
still manage your hierarchy cleanly with `ldap-sync`.

