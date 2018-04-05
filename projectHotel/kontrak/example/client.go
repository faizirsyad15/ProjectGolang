package main

import (
	"context"
	"time"

	cli "projectHotel/hotel/kontrak/endpoint"
	svc "projectHotel/hotel/kontrak/server"
	opt "projectHotel/hotel/util/grpc"

	util "projectHotel/hotel/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCKontrakClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Kontrak
	//client.AddKontrakService(context.Background(), svc.Kontrak{IdKontrak: "KT002", TanggalMulai: "2014-02-18", TanggalSelesai: "0000-00-00", Keterangan: "Karyawan Tetap"})

	//Get Kontrak By Selesai
	// ktkSelesai, _ := client.ReadKontrakBySelesaiService(context.Background(), "083819805020")
	// fmt.Println("daftar kontrak berdasarkan telepon:", ktkSelesai)

	//List Kontrak
	// ktks, _ := client.ReadKontrakService(context.Background())
	// fmt.Println("semua kontrak:",ktks)

	//Update Kontrak
	client.UpdateKontrakService(context.Background(), svc.Kontrak{TanggalMulai: "2017-03-03", TanggalSelesai: "2019-03-03", Keterangan: "Karyawan Kontrak", Status: "1", IdKontrak: "KT001"})

	//Get Customer By Mulai
	// ktkMulai, _ := client.ReadKontrakByMulaiService(context.Background(), "Faiz Irsyad")
	// fmt.Println("daftar kontrak berdasarkan nama:", ktkMulai)
}
