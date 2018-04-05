package main

import (
	"context"
	"time"

	cli "projectHotel/hotel/kota/endpoint"
	svc "projectHotel/hotel/kota/server"
	opt "projectHotel/hotel/util/grpc"

	util "projectHotel/hotel/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCKotaClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Kota
	//client.AddKotaService(context.Background(), svc.Kota{IdKota: "KT002", NamaKota: "Depok"})

	//List Kota
	// kots, _ := client.ReadKotaService(context.Background())
	// fmt.Println("semua kota:",kots)

	//Update Kota
	client.UpdateKotaService(context.Background(), svc.Kota{NamaKota: "Kutai", Status: "1", IdKota: "KT002"})

	//Get Kota By Nama
	// kotNama, _ := client.ReadKotaByNamaService(context.Background(), "Depok")
	// fmt.Println("daftar kota berdasarkan nama:", kotNama)
}
