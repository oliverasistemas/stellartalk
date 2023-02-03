CHAIN_0=stellartalk-chain-0
CHAIN_NAME=stellartalk

rm -rf ~/."$CHAIN_NAME"/

echo "Init configuration"
stellartalkd init stellarTalkNode --chain-id "$CHAIN_0"

echo "add validator keys to test backend"
stellartalkd keys add validator_0 --keyring-backend test
VALIDATOR_0_ADDRESS=$(stellartalkd keys show validator_0 -a --keyring-backend test)

echo "add a genesis account to genesis.json"
stellartalkd add-genesis-account "$VALIDATOR_0_ADDRESS" 100000000000000stake

echo "Generate a genesis tx carrying a self delegation"
stellartalkd gentx validator_0 100000000000stake --chain-id "$CHAIN_0" --keyring-backend test

echo "collect genesis txs and output a genesis.json file"
stellartalkd collect-gentxs