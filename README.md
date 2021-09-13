# go-rln

Wrappers for [kilic/rln](https://github.com/kilic/rln) along with an implementation for rate-limiting using RLN inspired
by the [Waku v2 RLN Relay](https://rfc.vac.dev/spec/17/) built by [Status](https://status.im).

Further research can be found here:
 - https://forum.vac.dev/t/vac-3-zk/97
 - https://github.com/vacp2p/research/tree/master/rln-research
 - https://ethresear.ch/t/semaphore-rln-rate-limiting-nullifier-for-spam-prevention-in-anonymous-p2p-setting/5009

The goal of this is to create a rate-limiter for blockchains where block production is cheap.
