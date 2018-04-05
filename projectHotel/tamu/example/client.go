package main

import (
	"context"
	"time"

	cli "projectHotel/hotel/tamu/endpoint"
	svc "projectHotel/hotel/tamu/server"
	opt "projectHotel/hotel/util/grpc"

	util "projectHotel/hotel/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCTamuClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Tamu
	//client.AddTamuService(context.Background(), svc.Tamu{IdTamu: "TM002", NamaTamu: "Putri", NoTelepon: "087771222111", JenisKelamin: "Perempuan", IdAlamat: "AL002"})

	//Get Tamu By Telepon
	// tamTelepon, _ := client.ReadTamuByTeleponService(context.Background(), "083819805020")
	// fmt.Println("daftar tamu berdasarkan telepon:", tamTelepon)

	//List Tamu
	// tams, _ := client.ReadTamuService(context.Background())
	// fmt.Println("semua tamu:",tams)

	//Update Tamu
	client.UpdateTamuService(context.Background(), svc.Tamu{NamaTamu: "Putri Hamdah", NoTelepon: "087771212133", JenisKelamin: "Perempuan", IdAlamat: "AL002", Status: "1", IdTamu: "TM002"})

	//Get Customer By Nama
	// tamNama, _ := client.ReadTamuByNamaService(context.Background(), "Faiz Irsyad")
	// fmt.Println("daftar tamu berdasarkan nama:", tamNama)
}
