package main

import (
	"context"
	"fmt"
	"time"

	cli "projectHotel/hotel/karyawan/endpoint"
	//svc "projectHotel/hotel/karyawan/server"
	opt "projectHotel/hotel/util/grpc"

	util "projectHotel/hotel/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCKaryawanClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Karyawan
	//client.AddKaryawanService(context.Background(), svc.Karyawan{IdKaryawan: "KR005", NamaKaryawan: "Delon", Alamat: "Jakarta Barat", NoTelepon: "087221342633", Keterangan: "Baik", IdJabatan: "JB002", IdDepartemen: "DP001", IdGolongan: "GL002", IdKontrak: "KT002"})

	//Get Karyawan By Telepon
	// karTelepon, _ := client.ReadKaryawanByTeleponService(context.Background(), "085661342111")
	// fmt.Println("daftar karyawan berdasarkan telepon:", karTelepon)

	//List Karyawan
	// kars, _ := client.ReadKaryawanService(context.Background())
	// fmt.Println("semua karyawan:", kars)

	//Update Karyawan
	//client.UpdateKaryawanService(context.Background(), svc.Karyawan{NamaKaryawan: "Reva", Alamat: "Jakarta Utara", NoTelepon: "083881342633", Keterangan: "Cukup Baik", IdJabatan: "JB003", IdDepartemen: "DP001", IdGolongan: "GL002", IdKontrak: "KT001", Status: "1", IdKaryawan: "KR004"})

	//Get Customer By Nama
	// karNama, _ := client.ReadKaryawanByNamaService(context.Background(), "Faiz Irsyad")
	// fmt.Println("daftar karyawan berdasarkan nama:", karNama)

	//Get Customer By Keterangan Like
	karKeterangan, _ := client.ReadKaryawanByKeteranganService(context.Background(), "B%")
	fmt.Println("daftar karyawan berdasar huruf depan B:", karKeterangan)
}
