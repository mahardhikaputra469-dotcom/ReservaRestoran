package main
import "fmt"


const MAX = 100


type meja struct{
    Nomor int
    Kapasitas int
    Tersedia bool
}


type pelanggan struct{
    ID int
    Nama string
}


type reservasi struct{
    IDReservasi int
    pelangganID int
    nomorID int
    hari string
    jam string
}


var dataMeja [MAX]meja
var dataPelanggan [MAX]pelanggan
var dataReservasi [MAX]Reservasi


var nMeja int
var nPelanggan int
var nReservasi int


func tambahMeja(nomor int, kapasitas int, status string) {
    dataMeja[nMeja].Nomor = nomor
    dataMeja[nMeja].Kapasitas = kapasitas
    dataMeja[nMeja].Status = status
    nMeja++
}


func tampilMeja() {
    var i int
    fmt.Println("\n==== DATA MEJA ===")
    for i = 0; i < nMeja; i++ {
        fmt.Println(dataMeja[i].Nomor, dataMeja[i].Kapasitas, dataMeja[i].Status)
    }
}


func sequentialSearchMeja(nomor int) int {
    var i int
    var idx int = -1
    for i = 0; i < nMeja; i++ {
        if dataMeja[i].Nomor == nomor{
            idx = i
        }
    }
    return idx
}


func binarySearchMeja(nomor int) int {
    var left, right, mid int
    var idx int = -1


    left = 0
    right = nMeja - 1
    for left <= right {
        mid = (left + right) /2


        if dataMeja[mid].Nomor == nomor {
            idx = mid
            left = right + 1
        }else if dataMeja[mid].Nomor < nomor {
            left = mid + 1
        }else {
            right = mid - 1
        }
    }
    return idx
}


func editMeja(nomor int, kapasitasBaru int) {
    var idx int = sequentialSearchMeja(nomor)
    if idx != -1 {
        dataMeja[idx].Kapasitas = kapasitasBaru
        fmt.Println("Data berhasil diubah")
    }else {
        fmt.Println("Data tidak ditemukan")
    }
}


func main() {
}