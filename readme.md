# Stellartalk chat blockchain

Stellartalk is decentralized blockchain that allows addresses to send messages between each other. It is built using Cosmos SDK and tendermint consensus engine. Initial bootstrap made with [Ignite CLI](https://ignite.com/cli).

## Install binaries
Install stellartalk binaries in your $GOPATH/bin folder.
```
make install
```

## Initialize local node
This command will clean up the configuration folder, create a new one add validators to test keyring and create genesis file and transactions.
```
make init
```

## Start local node
```
make start
```
You should see something similar to this in your console

```shell
8:51AM INF starting node with ABCI Tendermint in-process
8:51AM INF service start impl=multiAppConn module=proxy msg={}
8:51AM INF service start connection=query impl=localClient module=abci-client msg={}
8:51AM INF service start connection=snapshot impl=localClient module=abci-client msg={}
8:51AM INF service start connection=mempool impl=localClient module=abci-client msg={}
8:51AM INF service start connection=consensus impl=localClient module=abci-client msg={}
8:51AM INF service start impl=EventBus module=events msg={}
8:51AM INF service start impl=PubSub module=pubsub msg={}
8:51AM INF service start impl=IndexerService module=txindex msg={}
8:51AM INF ABCI Handshake App Info hash= height=0 module=consensus protocol-version=0 software-version=
8:51AM INF ABCI Replay Blocks appHeight=0 module=consensus stateHeight=0 storeHeight=0
8:51AM INF initializing blockchain state from genesis.json
8:51AM INF asserting crisis invariants inv=1/12 module=x/crisis name=group/Group-TotalWeight
...........
8:51AM INF Version info block=11 p2p=8 tendermint_version=0.34.23
8:51AM INF This node is a validator addr=FF64609E695010834C305B3E5D9DF6B8973A3D4E module=consensus pubKey=vK7pj+4UbQgy9hSGWPGDvoFFl380Ffu6mXvjiYc/Eqw=
8:51AM INF P2P Node ID ID=1f969171c689af4585735f9bbfabb444cac6a7dc file=/Users/ao/.stellartalk/config/node_key.json module=p2p
```

## Sending chats
Now that your local Stellartalk node is running you can send your first message to other person using the following command
```shell
  stellartalkd tx stellartalk create-chat [content] [recipient] [flags]
```
Example:

```shell
stellartalkd tx stellartalk create-chat "First chat" cosmos10mamzxu2krqdewrjl7s5c3qzjwjvtcyavqlhh2  --from cosmos10mamzxu2krqdewrjl7s5c3qzjwjvtcyavqlhh2  --keyring-backend test
```
