package keeper

import (
	"blog/x/blog/types"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	// get post count
	count := k.GetPostCount(ctx)
	post.Id = count
	// get adapter
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))
	appendValue := k.cdc.MustMarshal(&post)         // 序列化
	store.Set(GetPostIDBytes(post.Id), appendValue) // 通过id设置post
	k.SetPostCount(ctx, count+1)
	return count
}

func (k Keeper) GetPost(ctx sdk.Context, id uint64) (val types.Post, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx)) // 获取store
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))
	b := store.Get(GetPostIDBytes(id)) // 通过id获取post
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val) // 反序列化
	return val, true
}

func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx)) // 获取store
	store := prefix.NewStore(storeAdapter, []byte{})                        // id 传入空字节数组
	bytesKey := types.KeyPrefix(types.PostCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.AppendUint64(bz, count)
	store.Set(bytesKey, bz)
}

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx)) // 获取store
	store := prefix.NewStore(storeAdapter, []byte{})                        // id 传入空字节数组
	byteKey := types.KeyPrefix(types.PostCountKey)
	bz := store.Get(byteKey) // 通过key获取value
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetPostIDBytes(id uint64) []byte {
	bz := make([]byte, 8)              // 8 bytes
	binary.BigEndian.PutUint64(bz, id) // 将id转换为字节
	return bz
}
