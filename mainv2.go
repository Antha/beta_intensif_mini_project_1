package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyLibrary struct{
	book_kode string
	title string
	writer string
	publisher string
	page_number int
	year int
}

var ListOfBook [] MyLibrary 
var firstInit = false

func main() {
	ListOfBook = append(ListOfBook, MyLibrary{
		book_kode:    "1001",
		title:        "Conan Movie",
		writer:       "Aoyama Gosho",
		publisher:    "Toei Animation",
		page_number:  99,
		year:         2019,
	})

	Menu()
}

func Menu()  {
	// pilihanMenu := 0
	var pilihanMenu int
	fmt.Println("=================================")
	fmt.Println("Sistem Manajemen Perpustakaan")
	fmt.Println("=================================")
	fmt.Println("Silahkan Pilih : ")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Lihat Daftar Buku")
	fmt.Println("3. Edit Buku")
	fmt.Println("4. Hapus Buku")
	fmt.Println("5. Keluar")
	fmt.Println("=================================")
	fmt.Print("Masukan Pilihan : ")
	_, err := fmt.Scanln(&pilihanMenu)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	fmt.Println("")

	switch pilihanMenu {
	case 1:
		AddBook()
	case 2:
		ViewBook()
	case 3:
		EditBook()
	case 4:
		DeleteBook()
	case 5:
		os.Exit(0)
	}

	Menu()
}

func DeleteBook() {
	//panic("unimplemented")
	inputanUser := bufio.NewReader(os.Stdin)

	urutanDeleted := 0
	isExist := false
	book_kode := ""


	fmt.Print("Silahkan Masukan Kode Buku Yang Mau Dihapus : (ketik e untuk ke menu )")
	book_kode, _ = inputanUser.ReadString('\n')
	book_kode = strings.TrimSpace(book_kode)

	if(book_kode == "e"){
		Menu();
	}

	for urutan, book := range ListOfBook {
		//println(`Urutan :`,urutan,`\n`)
		if (book.book_kode == book_kode){
			urutanDeleted = urutan
			isExist = true
		}
	}

	if(isExist == false){
		println("Maaf Kode Buku Tidak Ada\n")
		DeleteBook();
	}

	if urutanDeleted == 0 && len(ListOfBook) == 1 {
        // Only one item in the list
        ListOfBook = nil
    } else {
        ListOfBook = append(
            ListOfBook[:urutanDeleted],
            ListOfBook[urutanDeleted+1:]...,
        )
    }

	fmt.Println("Buku Berhasil Dihapus!")
	fmt.Println("")
}

func EditBook(){
	inputanUser := bufio.NewReader(os.Stdin)

	urutanEdited := 0
	isExist := false
	book_kode := ""
	title := ""
	writer := ""
	publisher := ""
	page_number := 0
	year := 0

	fmt.Print("Silahkan Masukan Kode Buku Yang Mau Diedit : (ketik e untuk ke menu) ")
	book_kode, _ = inputanUser.ReadString('\n')
	book_kode = strings.TrimSpace(book_kode)

	if(book_kode == "e"){
		Menu();
	}

	for urutan, book := range ListOfBook {
		if (book.book_kode == book_kode){
			urutanEdited = urutan
			isExist = true
		}
	}

	if(isExist == false){
		println("Maaf Kode Buku Tidak Ada\n")
		EditBook();
	}

	fmt.Print("Silahkan Masukan Judul : ")
	title, _ = inputanUser.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Silahkan Masukan Penulis : ")
	writer, _ = inputanUser.ReadString('\n')
	writer = strings.TrimSpace(writer)

	fmt.Print("Silahkan Masukan Publisher : ")
	publisher, _ = inputanUser.ReadString('\n')
	publisher = strings.TrimSpace(publisher)

	fmt.Print("Silahkan Masukan Jumlah Halaman : ")
	pageInput, _ := inputanUser.ReadString('\n')
	pageInput = strings.TrimSpace(pageInput)
	page_number, _ = strconv.Atoi(pageInput)

	fmt.Print("Silahkan Masukan Tahun Terbit : ")
	yearInput, _ := inputanUser.ReadString('\n')
	yearInput = strings.TrimSpace(yearInput)
	year, _ = strconv.Atoi(yearInput)

	ListOfBook[urutanEdited] = MyLibrary{
		book_kode:    book_kode,
		title:        title,
		writer:       writer,
		publisher:    publisher,
		page_number:  page_number,
		year:         year,
	}

	fmt.Println("")
}

func ViewBook() {
	//panic("unimplemented")
	fmt.Println("=================================")
	fmt.Println("Daftar Buku")
	fmt.Println("=================================")
	if(len(ListOfBook) == 0){
		fmt.Println("Buku Lagi Kosong");
	}else{
		for urutan, book := range ListOfBook {
			fmt.Printf("%d. Kode Buku : %s, Judul : %s, Penulis : %s, Penerbit : %s, Jumlah Hal. : %d, Tahun Terbit : %d,  \n",
				urutan+1,
				book.book_kode,
				book.title,
				book.writer,
				book.publisher,
				book.page_number,
				book.year,
			)
		}
	}
	

	fmt.Println("")
}

func AddBook() {
	inputanUser := bufio.NewReader(os.Stdin)

	book_kode := ""
	title := ""
	writer := ""
	publisher := ""
	page_number := 0
	year := 0

	fmt.Print("Silahkan Masukan Kode Buku : ")
	book_kode, _ = inputanUser.ReadString('\n')
	book_kode = strings.TrimSpace(book_kode)

	fmt.Print("Silahkan Masukan Judul : ")
	title, _ = inputanUser.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Silahkan Masukan Penulis : ")
	writer, _ = inputanUser.ReadString('\n')
	writer = strings.TrimSpace(writer)

	fmt.Print("Silahkan Masukan Publisher : ")
	publisher, _ = inputanUser.ReadString('\n')
	publisher = strings.TrimSpace(publisher)

	fmt.Print("Silahkan Masukan Jumlah Halaman : ")
	pageInput, _ := inputanUser.ReadString('\n')
	pageInput = strings.TrimSpace(pageInput)
	page_number, _ = strconv.Atoi(pageInput)

	fmt.Print("Silahkan Masukan Tahun Terbit : ")
	yearInput, _ := inputanUser.ReadString('\n')
	yearInput = strings.TrimSpace(yearInput)
	year, _ = strconv.Atoi(yearInput)

	ListOfBook = append(ListOfBook, MyLibrary{
		book_kode:    book_kode,
		title:        title,
		writer:       writer,
		publisher:    publisher,
		page_number:  page_number,
		year:         year,
	})

	fmt.Println("Berhasil Menambah Buku!")
	fmt.Println("")
}

