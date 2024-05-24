

# Overview

EpicChain-Contract contains all EpicChain related contracts written for the [EpicChain-Go](https://github.com/nspcc-dev/epicchain-go) compiler. These contracts are deployed both on the mainchain and the sidechain.

**Mainchain contracts:**

- epicchain
- processing

**Sidechain contracts:**

- alphabet
- audit
- balance
- container
- epicchainid
- netmap
- nns
- proxy
- reputation

# Getting Started 

## Prerequisites

To compile smart contracts, you need:

- [epicchain-go](https://github.com/epicchainlabs/epicchain-go) >= 0.104.0

## Compilation

To build and compile smart contracts, run the `make all` command. Compiled contracts `*_contract.nef` and manifest `config.json` files are placed in the corresponding directories. Generated RPC binding files `rpcbinding.go` are placed in the corresponding `rpc` directories.

```bash
$ make all
/home/user/go/bin/cli contract compile -i alphabet -c alphabet/config.yml -m alphabet/config.json -o alphabet/alphabet_contract.nef --bindings alphabet/bindings_config.yml
mkdir -p rpc/alphabet
/home/user/go/bin/cli contract generate-rpcwrapper -o rpc/alphabet/rpcbinding.go -m alphabet/config.json --config alphabet/bindings_config.yml
...
```

You can specify the path to the `epicchain-go` binary with the `epicchainGO` environment variable:

```bash
$ epicchainGO=/home/user/epicchain-go/bin/epicchain-go make all
```

Remove compiled files with the `make clean` command.

## Building Debian Package

To build a Debian package containing compiled contracts, run the `make debpackage` command. The package will install compiled contracts `*_contract.nef` and manifest `config.json` with corresponding directories to `/var/lib/epicchain/contract` for further usage. It will download and build `epicchain-go` if needed.

To clean package-related files, use `make debclean`.

# Testing

Smart contract tests reside in the `tests/` directory. To execute the test suite after applying changes, simply run `make test`.

```bash
$ make test
ok      github.com/nspcc-dev/epicchain-contract/tests       0.462s
```

# License

Contracts are licensed under the GPLv3 license, bindings and other integration code are provided under the Apache 2.0 license - see [LICENSE.md](LICENSE.md) for details.
