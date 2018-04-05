package main

import (
	"context"
	"time"

	cli "projectHotel/hotel/kamar/endpoint"
	svc "projectHotel/hotel/kamar/server"
	opt "projectHotel/hotel/util/grpc"

	util "projectHotel/hotel/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCKamarClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Kamar
	client.AddKamarService(context.Background(), svc.Kamar{IdKamar: "KM001", IdTipeKamar: "TK002", IdMenu: "MN001"})

	//Get Kamar By Telepon
	// tipTelepon, _ := client.ReadKamarByHargaService(context.Background(), "083819805020")
	// fmt.Println("daftar kamar berdasarkan telepon:", tipTelepon)

	//List Kamar
	// tips, _ := client.ReadKamarService(context.Background())
	// fmt.Println("semua kamar:",tips)

	//Update Kamar
	//client.UpdateKamarService(context.Background(), svc.Kamar{IdTipeKamar: "TK001", IdMenu: "MN001", Status: "1", IdKamar: "TK002"})

	//Get Customer By Nama
	// tipNama, _ := client.ReadKamarByNamaService(context.Background(), "Faiz Irsyad")
	// fmt.Println("daftar kamar berdasarkan nama:", tipNama)
}
