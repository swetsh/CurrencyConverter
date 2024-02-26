package main

import (
	pb "CurrencyConverterService/converter"
	"CurrencyConverterService/pkg/config"
	"CurrencyConverterService/pkg/models"
	"context"
	"log"
	"net"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCurrencyConverterServer
}

func (*server) ConvertCurrency(ctx context.Context, req *pb.ConversionRequest) (*pb.ConversionResponse, error) {

	var fromCurrency models.Currency
	var toCurrency models.Currency

	if err := config.GetDB().Where("currency = ?", req.FromCurrency).Find(&fromCurrency).Error; err != nil {
		return nil, err
	}

	if err := config.GetDB().Where("currency = ?", req.ToCurrency).Find(&toCurrency).Error; err != nil {
		return nil, err
	}

	return &pb.ConversionResponse{
		ConvertedAmount: (req.Amount / toCurrency.ExchangeRate) * fromCurrency.ExchangeRate,
	}, nil
}

func main() {

	config.DatabaseConnection()

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &server{}

	pb.RegisterCurrencyConverterServer(serverRegistrar, service)

	log.Println("Server started...")

	err = serverRegistrar.Serve(lis)

	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
