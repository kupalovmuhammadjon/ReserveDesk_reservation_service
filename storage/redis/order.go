package redis

import (
	"context"
	"encoding/json"
	pb "reservation_service/genproto/order"
	"time"

	"github.com/redis/go-redis/v9"
)

type OrderRepo struct {
	Db *redis.Client
}

func (o *OrderRepo) CreateOrder(order *pb.Order) error {
	menuItems, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}

	orderData := map[string]any{
		"reservation_id": order.ReservationId,
		"items":          string(menuItems),
		"quantity":       order.Quantity,
		"created_at":     time.Now().Unix(),
	}

	_, err = o.Db.HSet(context.Background(), order.ReservationId, orderData).Result()

	return err
}

func (o *OrderRepo) GetOrderByReservationId(id string) (*pb.OrderInfo, error) {

	orderByte, err := o.Db.Get(context.Background(), id).Bytes()
	if err != nil {
		return nil, err
	}

	order := pb.OrderInfo{}

	err = json.Unmarshal(orderByte, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
