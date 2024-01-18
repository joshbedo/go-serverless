package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/joho/godotenv"
	"github.com/joshbedo/serverless-go/Config"
	"github.com/joshbedo/serverless-go/Routers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ginLambda *ginadapter.GinLambda
var err error

func init() {
	fmt.Printf("Gin cold start\n")

	// load env vars
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to db
	dsn := fmt.Sprintf("%s&parseTime=True", os.Getenv("DSN"))

	Config.DB, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})

	if err != nil {
		panic("failed to connect database")
	}

	// setup routers
	r := Routers.SetupRouters()

	// if prod use ginLambda adapter otherwise just r.Run()
	if os.Getenv("stage") == "prod" {
		ginLambda = ginadapter.New(r)
	} else {
		r.Run(":3000")
	}
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
