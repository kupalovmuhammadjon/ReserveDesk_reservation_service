package postgres

import (
	"database/sql"
	pb "reservation_service/genproto/restaurant"
	"time"

	_ "github.com/lib/pq"
)


type RestaurantRepo struct {
	Db *sql.DB
}

func NewRestaurantRepo(db *sql.DB) RestaurantRepo {
	return RestaurantRepo{Db: db}
}

func (r *RestaurantRepo) CreateRestauranT(req *pb.RestaurantInfo) error {
	_, err := r.Db.Exec("insert into restaurants(name, address, total_avb_seats, phone_number, description, created_at) values($1, $2, $3, $4, $5, $6)",
		req.Name, req.Address, req.TotalAvbSeats, req.PhoneNumber, req.Description, time.Now())
	
	if err != nil {
		return err
	}
	return nil
}

func (r *RestaurantRepo) UpdateRestauranT(id *pb.Id) error {
	_, err := r.Db.Exec("update restaurants set ")
}