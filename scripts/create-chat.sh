echo "add validator keys to test backend"
VALIDATOR_0_ADDRESS=$(stellartalkd keys show validator_0 -a --keyring-backend test)

echo "Creating first chat from $VALIDATOR_0_ADDRESS"
stellartalkd tx stellartalk create-chat "First chat"  $VALIDATOR_0_ADDRESS --from $VALIDATOR_0_ADDRESS --keyring-backend test

echo "List chats"
stellartalkd query stellartalk list-chat
