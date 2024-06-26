package keeper

import (

    "encoding/binary"
    
     "quotechain/x/quotechain/types"
 
     "cosmossdk.io/store/prefix"

     "github.com/cosmos/cosmos-sdk/runtime"

     sdk "github.com/cosmos/cosmos-sdk/types"

)




func (k Keeper) AppendQuote(ctx sdk.Context, quote types.Quote) uint64 {
    
    num := k.GetQuoteNum(ctx)
    quote.Id = num
    k.SetQuote(ctx, quote)
    k.SetQuoteNum(ctx, num+1) 
    return num
}

func (k Keeper) GetQuoteNum(ctx sdk.Context) uint64 {
    storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    store := prefix.NewStore(storeAdapter, []byte{})
    byteKey := types.KeyPrefix(types.QuoteNumKey)
    bz := store.Get(byteKey)
    if bz == nil {
        return 0
    }
    return binary.BigEndian.Uint64(bz)
}
func GetQuoteIDBytes(id uint64) []byte {
     bz := make([]byte, 8)
     binary.BigEndian.PutUint64(bz, id)
     return bz
     }

func (k Keeper) SetQuoteNum(ctx sdk.Context, num uint64) {
    storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))    
    store := prefix.NewStore(storeAdapter, []byte{})
    byteKey := types.KeyPrefix(types.QuoteNumKey)
    bz := make([]byte, 8)
    binary.BigEndian.PutUint64(bz, num)
    store.Set(byteKey, bz)
}

func (k Keeper) SetQuote(ctx sdk.Context, post types.Quote) {
    storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.QuoteNumKey))
    b := k.cdc.MustMarshal(&post)
    store.Set(GetQuoteIDBytes(post.Id), b)
    }




  




 




              
     
