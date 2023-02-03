package types

const (
	// ModuleName defines the module name
	ModuleName = "stellartalk"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_stellartalk"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ChatKey      = "Chat/value/"
	ChatCountKey = "Chat/count/"
)
