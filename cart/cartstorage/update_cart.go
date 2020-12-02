package cartstorage

import (
	"context"
	"fmt"
	"fooddlv/cart/cartmodel"
	"fooddlv/common"
)

/**
Create a cart, return ID if can create a item, and 0 if error
*/
func (store *cartMysql) UpdateCart(ctx context.Context, cartUpdateData *cartmodel.CartUpdate) (int, error) {
	// init db
	db := store.db.Begin()
	// create data to db
	fmt.Println("update the cart item", cartUpdateData)
	if err := db.Table(cartmodel.Cart{}.TableName()).Updates(&cartUpdateData).Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	return cartUpdateData.ID, nil
}