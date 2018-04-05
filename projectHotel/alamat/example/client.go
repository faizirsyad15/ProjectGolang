package main

import (
	"context"
	"time"

	cli "projectHotel/hotel/alamat/endpoint"
	svc "projectHotel/hotel/alamat/server"
	opt "projectHotel/hotel/util/grpc"

	util "projectHotel/hotel/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCAlamatClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Alamat
	client.AddAlamatService(context.Background(), svc.Alamat{IdAlamat: "AL003", AlamatRumah: "Jl. Raya Cilebut", RtRw: "01/02", NoRumah: "26", IdKota: "KT003", IdNegara: "NG001"})

	//Get Alamat By No Rumah
	// norNoRumah, _ := client.ReadAlamatByNoRumahService(context.Background(), "55")
	// fmt.Println("daftar alamat berdasarkan no rumah:", norNoRumah)

	//List Alamat
	// amts, _ := client.ReadAlamatService(context.Background())
	// fmt.Println("semua alamat:",amts)

	//Update Alamat
	//client.UpdateAlamatService(context.Background(), svc.Alamat{AlamatRumah: "Jl. Raya Margonda", RtRw: "02/08", NoRumah: "45", IdKota: "KT001", IdNegara: "NG001", Status: "1", IdAlamat: "AL002"})
}
