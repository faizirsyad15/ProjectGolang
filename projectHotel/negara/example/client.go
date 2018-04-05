package main

import (
	"context"
	"time"

	cli "projectHotel/hotel/negara/endpoint"
	svc "projectHotel/hotel/negara/server"
	opt "projectHotel/hotel/util/grpc"

	util "projectHotel/hotel/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCNegaraClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Negara
	//client.AddNegaraService(context.Background(), svc.Negara{IdNegara: "NG003", NamaNegara: "Australia"})

	//List Negara
	// negs, _ := client.ReadNegaraService(context.Background())
	// fmt.Println("semua negara:",negs)

	//Update Negara
	client.UpdateNegaraService(context.Background(), svc.Negara{NamaNegara: "Swiss", Status: "1", IdNegara: "NG003"})

	//Get Negara By Nama
	// negNama, _ := client.ReadNegaraByNamaService(context.Background(), "Depok")
	// fmt.Println("daftar negara berdasarkan nama:", negNama)
}
