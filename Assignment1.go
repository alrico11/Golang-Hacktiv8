package main

import "fmt"

type alasanku struct {
	alasan string
}

func main() {
	var arrNama = [3]string{"Alrico Rizki", "Ardian Pandu", "Wimas Rizqi"}
	var arrAlamat = [3]string{"Purwoyoso", "GrahaPadma", "Tugu"}
	var arrPekerjaan = [3]string{"Mahasiswa", "Mahasiswa", "Mahasiswa"}
	s1 := alasanku{alasan: "Ingin Menjadi Ahli Pemrograman GoLang"}
	s2 := alasanku{alasan: "Ingin Menjadi Programmer"}
	s3 := alasanku{alasan: "Ingin Menjadi Backend Pemrograman GoLang"}
	for i := 0; i < 4; i++ {
		if i == 0 {
			fmt.Println("Nama 	    : ", arrNama[i])
			fmt.Println("Alamat	: ", arrAlamat[i])
			fmt.Println("Pekerjaan : ", arrPekerjaan[i])
			fmt.Println(s1.alasan, "\n")
		} else if i == 1 {
			fmt.Println("Nama 	    : ", arrNama[i])
			fmt.Println("Alamat	: ", arrAlamat[i])
			fmt.Println("Pekerjaan : ", arrPekerjaan[i])
			fmt.Println(s2.alasan, "\n")
		} else if i == 2 {
			fmt.Println("Nama 	    : ", arrNama[i])
			fmt.Println("Alamat	: ", arrAlamat[i])
			fmt.Println("Pekerjaan : ", arrPekerjaan[i])
			fmt.Println(s3.alasan, "\n")
		}

	}

}
