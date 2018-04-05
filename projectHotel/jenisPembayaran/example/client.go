package main

import (
	"context"
	"time"

	cli "projectHotel/hotel/jenisPembayaran/endpoint"
	svc "projectHotel/hotel/jenisPembayaran/server"
	opt "projectHotel/hotel/util/grpc"

	util "projectHotel/hotel/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCJenisPembayaranClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add JenisPembayaran
	//client.AddJenisPembayaranService(context.Background(), svc.JenisPembayaran{IdJenisPembayaran: "JP001", MetodePembayaran: "Cash"})

	//List JenisPembayaran
	// jens, _ := client.ReadJenisPembayaranService(context.Background())
	// fmt.Println("semua jenisPembayaran:",jens)

	//Update JenisPembayaran
	client.UpdateJenisPembayaranService(context.Background(), svc.JenisPembayaran{MetodePembayaran: "Debit", Status: "1", IdJenisPembayaran: "JP002"})

	//Get JenisPembayaran By Metode
	// jenMetode, _ := client.ReadJenisPembayaranByMetodeService(context.Background(), "Depok")
	// fmt.Println("daftar jenisPembayaran berdasarkan nama:", jenMetode)
}
