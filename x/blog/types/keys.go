package types

const (
	// ModuleName defines the module name
	ModuleName = "blog"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blog"

	// PostKey is a prefix for post store
	PostKey = "Post/value/"

	// this key will be used to keep track of the ID of the lastest post
	PostCountKey = "Post/count/"
)

var (
	ParamsKey = []byte("p_blog")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
