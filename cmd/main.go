package main

import (
	"os"

	"github.com/Lucasnhso/codepix-go/application/grpc"
	"github.com/Lucasnhso/codepix-go/infraestructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}