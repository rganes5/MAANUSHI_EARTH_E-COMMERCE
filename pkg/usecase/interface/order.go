package interfaces

import (
	"context"

	"github.com/rganes5/maanushi_earth_e-commerce/pkg/utils"
)

type OrderUseCase interface {
	PlaceNewOrder(ctx context.Context, addressId uint, paymentId uint, userId uint) error
	CancelOrder(ctx context.Context, userId uint, orderDetailsId uint) error
	UpdateStatus(ctx context.Context, orderDetailsId uint, statusId uint) error
	ReturnOrder(ctx context.Context, orderDetailsId uint, statusId uint) error
	ListOrders(ctx context.Context, id uint, pagination utils.Pagination) ([]utils.ResponseOrders, error)
	AdminListOrders(ctx context.Context, pagination utils.Pagination) ([]utils.ResponseOrdersAdmin, error)
	ListOrderDetails(ctx context.Context, orderId uint, pagination utils.Pagination) ([]utils.ResponseOrderDetails, error)
}
