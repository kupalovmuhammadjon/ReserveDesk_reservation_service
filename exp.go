package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Order struct {
	ReservationId string   `protobuf:"bytes,2,opt,name=reservation_id,json=reservationId,proto3" json:"reservation_id,omitempty"`
	MenuItemsId   []string `protobuf:"bytes,3,rep,name=menu_items_id,json=menuItemsId,proto3" json:"menu_items_id,omitempty"`
	Quantity      int32
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	order := Order{"123", []string{"gf", "gfd"}, 2}

	menuItems, err := json.Marshal(order.MenuItemsId)
	if err != nil {
		panic(err)
	}
	orderData := map[string]any{
		"reservation_id": order.ReservationId,
		"menu_items":     string(menuItems),
		"quantity":       order.Quantity,
		"created_at":     time.Now().Unix(),
	}

	_, err = rdb.HSet(context.Background(), order.ReservationId, orderData).Result()
	if err != nil {
		panic(err)
	}

	res := rdb.HGetAll(context.Background(), order.ReservationId).Val()
	fmt.Println(res)
}
