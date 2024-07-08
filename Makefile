CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

exp:
	export DBURL='postgres://postgres:0509@localhost:5432/reja?sslmode=disable'

mig-up:
	migrate -database 'postgres://postgres:0509@localhost:5432/reja?sslmode=disable' -path migrations up
mig-down:
	migrate -database 'postgres://postgres:0509@localhost:5432/reja?sslmode=disable' -path migrations down



mig-create:
	migrate create -ext sql -dir migrations -seq create_courier

mig-insert:
	migrate create -ext sql -dir db/migrations -seq insert_table


# mig-force:
#  	migrate -database 'postgres://postgres:0509@localhost:5432/reja?sslmode=disable' -path migration force 1

mig-foce:
	migrate -database 'postgres://postgres:0509@localhost:5432/reja?sslmode=disable' -path migrations force 1