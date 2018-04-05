package main

import (
	"context"
	"time"

	cli "projectHotel/hotel/departemen/endpoint"
	svc "projectHotel/hotel/departemen/server"
	opt "projectHotel/hotel/util/grpc"

	util "projectHotel/hotel/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCDepartemenClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Departemen
	//client.AddDepartemenService(context.Background(), svc.Departemen{IdDepartemen: "DP002", NamaDepartemen: "Food Dept"})

	//List Departemen
	// deps, _ := client.ReadDepartemenService(context.Background())
	// fmt.Println("semua departemen:",deps)

	//Update Departemen
	client.UpdateDepartemenService(context.Background(), svc.Departemen{NamaDepartemen: "Housekeeping Dept", Status: "1", IdDepartemen: "DP002"})

	//Get Departemen By Nama
	// depNama, _ := client.ReadDepartemenByNamaService(context.Background(), "Depok")
	// fmt.Println("daftar departemen berdasarkan nama:", depNama)
}
