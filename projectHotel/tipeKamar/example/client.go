package main

import (
	"context"
	"time"

	cli "projectHotel/hotel/tipeKamar/endpoint"
	svc "projectHotel/hotel/tipeKamar/server"
	opt "projectHotel/hotel/util/grpc"

	util "projectHotel/hotel/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCTipeKamarClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add TipeKamar
	//client.AddTipeKamarService(context.Background(), svc.TipeKamar{IdTipeKamar: "TK002", NamaTipeKamar: "Suite Room", HargaKamar: 700000})

	//Get TipeKamar By Telepon
	// tipTelepon, _ := client.ReadTipeKamarByHargaService(context.Background(), "083819805020")
	// fmt.Println("daftar tipeKamar berdasarkan telepon:", tipTelepon)

	//List TipeKamar
	// tips, _ := client.ReadTipeKamarService(context.Background())
	// fmt.Println("semua tipeKamar:",tips)

	//Update TipeKamar
	client.UpdateTipeKamarService(context.Background(), svc.TipeKamar{NamaTipeKamar: "Standard Room", HargaKamar: 500000, Status: "1", IdTipeKamar: "TK002"})

	//Get Customer By Nama
	// tipNama, _ := client.ReadTipeKamarByNamaService(context.Background(), "Faiz Irsyad")
	// fmt.Println("daftar tipeKamar berdasarkan nama:", tipNama)
}
