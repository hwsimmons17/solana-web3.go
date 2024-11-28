# solana-web3.go
The point of this project is to create a high quality Solana client in Golang similar to the client developed by [Solana Labs](https://github.com/solana-labs/solana-web3.js).

## Goals of the project
* Create APIs you are used to using in JS world in Golang
* Better help the developer understand what is going on under the hood if they wish
* Make it simple for developers to create beautiful Solana clients on the backend in Golang

## Why have I undertaken this?
When I joined the Solana community in 2021, dApps were all the rage. Most development philosophies were either you would be building Solana programs in Rust or you would be building JS clients in the browser. The support for Solana clients supported that with really the only two good clients having been written in Rust and in JS.
In founding Teleport we were unique in that we were building for mobile and using backends (written in Go). The Solana clients available to us at the time were OK for what we needed them to be, but we ended up re-writing the Golang client from scratch because it didn't support the use-cases we needed. Now in late 2024, I see the landscape of Solana and I am nothing but excited for the opportunities. I think we were narrow-sighted in the dApps evnisioned back in 2021. I see a world in which people will be using Solana in a multitude of ways, most of which are yet to be envisioned. Because of this, I want to give back to the Solana ecosystem I love so much.

## Why Golang?
Golang to me is the perfect language for the backend. It strikes a great balance between performance and readability. I believe Solana needs to find it's way into centralized backends.