package types

const (
	// ModuleName defines the module name
	ModuleName = "pepe"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_pepe"
)

var (
	ParamsKey = []byte("p_pepe")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
