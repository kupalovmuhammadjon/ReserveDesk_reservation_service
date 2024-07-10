package main

import (
	"fmt"
	"log"
	"net"
	"reservation_service/config"
	pbrev "reservation_service/genproto/reservations"

	// pbm "reservation_service/genproto/menu"
	// pbr "reservation_service/genproto/restaurant"
	"reservation_service/pkg/logger"
	"reservation_service/service"
	"reservation_service/storage/postgres"

	"google.golang.org/grpc"
)

func main() {

	cfg := config.Load()

	logger, err := logger.New("debug", "develop", "app.log")
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := postgres.ConnectDB()
	if err != nil {
		logger.Fatal(err.Error())
		return
	}

	fmt.Println(logger)
	fmt.Println("wsdfghn")

	listener, err := net.Listen("tcp", cfg.RESERVATION_SERVICE_PORT)
	if err != nil {
		logger.Fatal("error while making tcp connection")
		return
	}

	server := grpc.NewServer()

	pbrev.RegisterReservationServiceServer(server, service.NewReservationService(db, cfg))
	// pbr.RegisterRestaurantServer(server, service.NewRestaurantService(db))
	// pbm.RegisterMenuServiceServer(server, service.NewMenuService(db))

	err = server.Serve(listener)
	if err != nil{
		return
	}
}
