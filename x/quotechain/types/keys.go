package types

const (
	// ModuleName defines the module name
	ModuleName = "quotechain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_quotechain"

        QuoteNumKey = "Quote/value/"
)

var (
	ParamsKey = []byte("p_quotechain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
