VALIDATOR_0_ADDRESS=$(stellartalkd keys show validator_0 -a --keyring-backend test)
BOB_ADDRESS=$(stellartalkd keys show bob -a --keyring-backend test)
ALICE_ADDRESS=$(stellartalkd keys show alice -a --keyring-backend test)
MESSAGE="The Great Wall of China is not visible from space with the naked eye, despite a common myth to the contrary."

echo "Creating first chat from $BOB_ADDRESS (Bob) to $ALICE_ADDRESS (Alice)"
stellartalkd tx bank send  $VALIDATOR_0_ADDRESS $BOB_ADDRESS 2000stake --keyring-backend test
stellartalkd tx bank send  $VALIDATOR_0_ADDRESS $ALICE_ADDRESS 2000stake --keyring-backend test

stellartalkd tx stellartalk create-chat "$MESSAGE"  "$VALIDATOR_0_ADDRESS" --from "$VALIDATOR_0_ADDRESS" --keyring-backend test

#echo "List chats"
#sleep 10
#stellartalkd query stellartalk list-chat
