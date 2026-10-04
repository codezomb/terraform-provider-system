# Terraform Provider for (Operating) System

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/neuspaces/terraform-provider-system)](https://github.com/neuspaces/terraform-provider-system/releases)
[![main](https://github.com/neuspaces/terraform-provider-system/actions/workflows/main.yml/badge.svg)](https://github.com/neuspaces/terraform-provider-system/actions/workflows/main.yml)
[![GitHub Discussions](https://img.shields.io/github/discussions/neuspaces/terraform-provider-system)](https://github.com/neuspaces/terraform-provider-system/discussions)
[![GitHub License](https://img.shields.io/github/license/neuspaces/terraform-provider-system)](https://github.com/neuspaces/terraform-provider-system/blob/main/LICENSE)
[![Terraform Registry](https://img.shields.io/badge/terraform-registry-5c4ee5.svg)](https://registry.terraform.io/providers/neuspaces/system/latest)

Releases: [registry.terraform.io](https://registry.terraform.io/providers/neuspaces/system/latest)

Documentation: [registry.terraform.io](https://registry.terraform.io/providers/neuspaces/system/latest/docs)

Discuss: [github.com/discussions](https://github.com/neuspaces/terraform-provider-system/discussions)

> **Fork.** This is `codezomb/terraform-provider-system`, a maintained fork of
> [neuspaces/terraform-provider-system](https://github.com/neuspaces/terraform-provider-system) at v0.5.0.
> It adds: `host_key` works with Ed25519 keys, an `overwrite` attribute on `system_file` and
> `system_folder` to adopt existing paths, and Go 1.26 with current dependencies. The badges and registry links
> below point at upstream; this fork is not published to a registry.

The Terraform Provider for (Linux Operating) System allows managing files, directories, users, groups, packages, and services on remote servers on operating system level agent-less via SSH.

> Even in a cloud-native heaven ☁️, there will still be use cases for pets 🐈

## Highlights

* Manage files, directories, users, groups, packages, and services on remote servers
* Connect to and authenticate with remote servers via SSH
* No agent on remote server required
* Seamless integration with Terraform providers of [all major IaaS cloud providers](examples)
* Support for Debian, Alpine, and Fedora Linux confirmed via acceptance test suite

## Quick Starts

- [Using the provider](https://registry.terraform.io/providers/neuspaces/system/latest/docs)
- [Examples](examples)

## Use Cases

The provider aims to allow configuring individual remote servers according to *mutable infrastructure* approach. You might find your use case in the following non-exhaustive list:

* Individual servers or virtual machines which are not recreated when configuration changes
* Share or distribute server or virtual machine configuration as using Terraform modules

The provider is not suitable for *immutable infrastructure* approaches such as fleets of homogeneous virtual machines. In this case, you may consider a more suitable configuration mechanism.

## User documentation

Refer to the comprehensive [user documentation of the provider in the Terraform Registry](https://registry.terraform.io/providers/neuspaces/system/latest/docs).

## Frequently Asked Questions

Responses to the most frequently asked questions can be found in the [FAQ](https://registry.terraform.io/providers/neuspaces/system/latest/docs/guides/faq).

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12+
