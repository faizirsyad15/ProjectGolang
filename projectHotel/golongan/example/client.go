package main

import (
	"context"
	"time"

	cli "projectHotel/hotel/golongan/endpoint"
	svc "projectHotel/hotel/golongan/server"
	opt "projectHotel/hotel/util/grpc"

	util "projectHotel/hotel/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCGolonganClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Golongan
	//client.AddGolonganService(context.Background(), svc.Golongan{IdGolongan: "GL002", NamaGolongan: "Golongan B"})

	//List Golongan
	// gols, _ := client.ReadGolonganService(context.Background())
	// fmt.Println("semua golongan:",gols)

	//Update Golongan
	client.UpdateGolonganService(context.Background(), svc.Golongan{NamaGolongan: "Golongan A", Status: "1", IdGolongan: "GL001"})

	//Get Golongan By Nama
	// golNama, _ := client.ReadGolonganByNamaService(context.Background(), "Depok")
	// fmt.Println("daftar golongan berdasarkan nama:", golNama)
}
