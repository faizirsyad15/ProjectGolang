package main

import (
	"context"
	"time"

	cli "projectHotel/hotel/jabatan/endpoint"
	svc "projectHotel/hotel/jabatan/server"
	opt "projectHotel/hotel/util/grpc"

	util "projectHotel/hotel/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCJabatanClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Jabatan
	//client.AddJabatanService(context.Background(), svc.Jabatan{IdJabatan: "JB002", NamaJabatan: "Resepsionist"})

	//List Jabatan
	// jabs, _ := client.ReadJabatanService(context.Background())
	// fmt.Println("semua jabatan:",jabs)

	//Update Jabatan
	client.UpdateJabatanService(context.Background(), svc.Jabatan{NamaJabatan: "Chef", Status: "1", IdJabatan: "JB002"})

	//Get Jabatan By Nama
	// jabNama, _ := client.ReadJabatanByNamaService(context.Background(), "Depok")
	// fmt.Println("daftar jabatan berdasarkan nama:", jabNama)
}
