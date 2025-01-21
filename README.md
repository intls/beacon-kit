</br>

<div align="center">
  <a href="https://github.com/berachain/beacon-kit">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="https://res.cloudinary.com/duv0g402y/image/upload/v1718034312/BeaconKitBanner.png">
      <img alt="beacon-kit-banner" src="https://res.cloudinary.com/duv0g402y/image/upload/v1718034312/BeaconKitBanner.png" width="auto" height="auto">
    </picture>
  </a>
</div>
<h2 >
  A modular framework for building EVM consensus clients ⛵️✨
</h2>

<div>

[![CI status](https://github.com/berachain/beacon-kit/workflows/pipeline/badge.svg)](https://github.com/berachain/beacon-kit/actions/workflows/pipeline.yml)
[![CodeCov](https://codecov.io/gh/berachain/beacon-kit/graph/badge.svg?token=0l5iJ3ZbzV)](https://codecov.io/gh/berachain/beacon-kit)
[![Telegram Chat](https://img.shields.io/endpoint?color=neon&logo=telegram&label=chat&url=https%3A%2F%2Ftg.sumanjay.workers.dev%2Fbeacon_kit)](https://t.me/beacon_kit)
[![X Follow](https://img.shields.io/twitter/follow/berachain)](https://x.com/berachain)
[![Discord](https://img.shields.io/discord/924442927399313448?label=discord)](https://discord.gg/berachain)

</div>

## What is BeaconKit?

[BeaconKit](https://docs.berachain.com/learn/what-is-beaconkit) is a modular framework for building EVM based consensus clients.
The framework offers the most user-friendly way to build and operate an EVM blockchain, while ensuring a functionally identical execution environment to that of the Ethereum Mainnet.

## Supported Execution Clients

Through utilizing the [Ethereum Engine API](https://github.com/ethereum/execution-apis/blob/main/src/engine)
BeaconKit is able to support all 6 major Ethereum execution clients:

- [**Geth**](https://geth.ethereum.org/): Official Go implementation of the Ethereum protocol.
- [**Erigon**](https://erigon.tech/): More performant, feature-rich client forked from `go-ethereum`.
- [**Nethermind**](https://www.nethermind.io/): .NET based client with full support for Ethereum protocols.
- [**Besu**](https://www.lfdecentralizedtrust.org/projects/besu): Enterprise-grade client, Apache 2.0 licensed, written in Java.
- [**Reth**](https://reth.rs/): Rust-based client focusing on performance and reliability.
- [**Ethereumjs**](https://ethereumjs.readthedocs.io/en/latest/#): Javascript based client managed by the Ethereum Foundation.

## Running a Local Development Network

**Prerequisites:**

- [Docker](https://docs.docker.com/engine/install/)
- [Golang 1.23.0+](https://go.dev/doc/install)
- [Foundry](https://book.getfoundry.sh/)

Start by opening two terminals side-by-side:

**Terminal 1:**

```bash
# Start the sample BeaconKit Consensus Client:
make start
```

**Terminal 2:**

```bash
# Start an Ethereum Execution Client:
make start-reth # or start-geth start-besu start-erigon start-nethermind start-ethereumjs
```

The account with
`private-key=0xfffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306`
corresponding with `address=0x20f33ce90a13a4b5e7697e3544c3083b8f8a51d4` is
preloaded with the native EVM token.

## Multinode Local Devnet

Please refer to the [Kurtosis README](https://github.com/berachain/beacon-kit/blob/main/kurtosis/README.md) for more information on how to run a multinode local devnet.
