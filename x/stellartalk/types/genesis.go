package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ChatList: []Chat{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in chat
	chatIdMap := make(map[uint64]bool)
	chatCount := gs.GetChatCount()
	for _, elem := range gs.ChatList {
		if _, ok := chatIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for chat")
		}
		if elem.Id >= chatCount {
			return fmt.Errorf("chat id should be lower or equal than the last id")
		}
		chatIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
