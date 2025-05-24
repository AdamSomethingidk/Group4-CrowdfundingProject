package main

import "fmt"

type Project struct {
	Creator        string
	ProjectName    string
	Category       string
	TotalDonator   int
	TargetDonation float64
	TotalDonation  float64
}

var mainArr = [10]Project{}

func menu() {
	fmt.Println("=============================================")
	fmt.Println("             CROWDFUNDING SYSTEM             ")
	fmt.Println("=============================================")
	fmt.Println("1. Buat proyek baru")
	fmt.Println("2. Ubah proyek")
	fmt.Println("3. Hapus proyek")
	fmt.Println("4. Berikan donasi ke proyek")
	fmt.Println("5. Cari proyek berdasarkan nama (sequential)")
	fmt.Println("6. Cari proyek berdasarkan nama (binary)")
	fmt.Println("7. Cari proyek berdasarkan kategori (sequential)")
	fmt.Println("8. Urutkan proyek berdasarkan total dana terkumpul (Selection)")
	fmt.Println("9. Urutkan proyek berdasarkan total dana terkumpul (Insertion)")
	fmt.Println("10. Urutkan proyek berdasarkan jumlah donatur (Selection)")
	fmt.Println("11. Urutkan proyek berdasarkan jumlah donatur (Insertion)")
	fmt.Println("12. Tampilkan proyek yang telah memenuhi target pendanaan")
	fmt.Println("0. Keluar")
	fmt.Println("=============================================")
	fmt.Print("Masukan Pilihan: ")
}

func main() {
	var n, currentSize int
	n = 100
	test()
	for n != 0 {
		menu()
		fmt.Scan(&n)
		if n == 1 {
			createProject(&currentSize)
		}
		if n == 2 {
			ChangeProject()
		}
		if n == 3 {
			deleteProject(&currentSize)
		}
		if n == 6 {
			findProjectNameBin()
		}
		if n == 7 {
			findProjectCategorySeq()
		}
		if n == 9 {
			sortTotalDonationInsert()
		}
	}

}

//Hanya untuk testing
func test() {
	//index 0
	mainArr[0].Creator = "Z"
	mainArr[0].ProjectName = "Zproject"
	mainArr[0].Category = "Z"
	//index 1
	mainArr[1].Creator = "A"
	mainArr[1].ProjectName = "Aproject"
	mainArr[1].Category = "A"
	//index 2
	mainArr[2].Creator = "C"
	mainArr[2].ProjectName = "Cproject"
	mainArr[2].Category = "C"
	mainArr[2].TotalDonator = 12
	mainArr[2].TargetDonation = 25000
	mainArr[2].TotalDonation = 26000
	//index 3
	mainArr[3].Creator = "B"
	mainArr[3].ProjectName = "Bproject"
	mainArr[3].Category = "B"
	mainArr[3].TotalDonator = 8
	mainArr[3].TargetDonation = 15000
	mainArr[3].TotalDonation = 10000
	//index 4
	mainArr[4].Creator = "E"
	mainArr[4].ProjectName = "Eproject"
	mainArr[4].Category = "A"
}

//Buat proyek baru
func createProject(currentSize *int) {
	var newProject Project
	var nameTaken bool
	nameTaken = false
	newProject.TotalDonation = 0
	newProject.TotalDonator = 0
	fmt.Print("Masukan nama anda (tidak boleh kosong): ")
	fmt.Scan(&newProject.Creator)
	fmt.Print("Masukan nama proyek {tidak boleh kosong}: ")
	fmt.Scan(&newProject.ProjectName)
	fmt.Print("Masukan kategori proyek (tidak boleh kosong): ")
	fmt.Scan(&newProject.Category)
	fmt.Print("Masukan target pendanaan (tidak boleh kosong): ")
	fmt.Scan(&newProject.TargetDonation)
	for i := 0; i < len(mainArr); i++ {
		if mainArr[i].ProjectName == newProject.ProjectName {
			nameTaken = true
		}
	}
	if nameTaken {
		fmt.Print("Nama proyek telah terpakai")
	} else {
		mainArr[*currentSize] = newProject
		*currentSize = *currentSize + 1
	}
}

//Ubah proyek
func ChangeProject() {
	var name, projectToChange string
	var pickedProject Project
	var found bool
	found = false
	fmt.Print("Masukan nama anda: ")
	fmt.Scan(&name)
	for i := 0; i < len(mainArr); i++ {
		if mainArr[i].Creator == name {
			fmt.Println("--------------------")
			fmt.Println("Nama proyek: ", mainArr[i].ProjectName)

		}
	}
	fmt.Print("Pilih proyek yang akan diubah (gunakan nama): ")
	fmt.Scan(&projectToChange)
	for j := 0; j < len(mainArr); j++ {
		if mainArr[j].ProjectName == projectToChange && mainArr[j].Creator == name {
			found = true
			fmt.Print("Ubah nama: ")
			fmt.Scan(&pickedProject.ProjectName)
			fmt.Print("Ubah kategori: ")
			fmt.Scan(&pickedProject.Category)
			fmt.Print("Ubah target pendanaan: ")
			fmt.Scan(&pickedProject.TargetDonation)
			mainArr[j].ProjectName = pickedProject.ProjectName
			mainArr[j].Category = pickedProject.Category
			mainArr[j].TargetDonation = pickedProject.TargetDonation
			break
		}
	}
	if !found {
		fmt.Print("Proyek tidak ditemukan atau bukan proyek anda.")
	}
}

//Hapus proyek
func deleteProject(currentSize *int) {
	var name, verify, projectToDelete string
	fmt.Print("Masukan nama anda: ")
	fmt.Scan(&name)
	for i := 0; i < len(mainArr); i++ {
		if mainArr[i].Creator == name {
			fmt.Println("----------------------")
			fmt.Println("Nama proyek: ", mainArr[i].ProjectName)
		}
	}
	fmt.Print("Pilih proyek yang akan dihapus (gunakan nama): ")
	fmt.Scan(&projectToDelete)
	for j := 0; j < len(mainArr); j++ {
		if mainArr[j].ProjectName == projectToDelete && mainArr[j].Creator == name {
			fmt.Print("Proyek ", mainArr[j], " akan dihapus, apakah anda yakin? Y/N")
			fmt.Scan(&verify)
			if verify == "y" || verify == "Y" {
				for k := j; k < *currentSize-1; k++ {
					mainArr[k] = mainArr[k+1]
				}
				fmt.Print("Proyek telah dihapus.")
				*currentSize = *currentSize - 1
			} else {
				fmt.Print("Penghapusan proyek dibatalkan.")
			}
		}
	}
}

//Cari proyek berdasarkan nama (binary)
func findProjectNameBin() {
	var tempArr = [10]Project{}
	var temp Project
	var j, low, high, mid, index int
	var pickedName string
	tempArr = mainArr
	index = -1
	//sort array
	for i := 1; i < len(tempArr); i++ {
		temp = tempArr[i]
		j = i - 1
		for j >= 0 && tempArr[j].ProjectName > temp.ProjectName {
			tempArr[j+1] = tempArr[j]
			j = j - 1
		}
		tempArr[j+1] = temp
	}
	fmt.Println(tempArr)
	fmt.Print("Pilih nama yang ingin dicari: ")
	fmt.Scan(&pickedName)
	low = 0
	high = len(tempArr)
	for high >= low {
		mid = (low + high) / 2
		if tempArr[mid].ProjectName == pickedName {
			index = mid
		} else if tempArr[mid].ProjectName > pickedName {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if index == -1 {
		fmt.Print("Proyek dengan kategori ", pickedName, " tidak ditemukan.")
	} else {
		fmt.Println(tempArr[index])
	}
}

//Cari proyek berdasarkan kategori (sequential)
func findProjectCategorySeq() {
	var pickedCategory string
	fmt.Print("Masukan kategori proyek yang ingin dicari: ")
	fmt.Scan(&pickedCategory)
	for i := 0; i < len(mainArr); i++ {
		if mainArr[i].Category == pickedCategory {
			fmt.Println("Nama pemilik proyek: ", mainArr[i].Creator)
			fmt.Println("Nama proyek: ", mainArr[i].ProjectName)
			fmt.Println("Kategori proyek: ", mainArr[i].Category)
			fmt.Println("Jumlah donator: ", mainArr[i].TotalDonator)
			fmt.Println("Target dana: ", mainArr[i].TargetDonation)
			fmt.Println("Dana yang terkumpul: ", mainArr[i].TotalDonation)
		}
	}
}

//Urutkan proyek berdasarkan total dana terkumpul (Insertion)
func sortTotalDonationInsert() {
	var tempArr = [10]Project{}
	var temp Project
	var j int
	tempArr = mainArr
	fmt.Print("Sebelum sorting: ")
	fmt.Println(tempArr)
	for i := 1; i <= len(tempArr); i++ {
		temp = tempArr[i]
		j = i - 1
		for j >= 0 && tempArr[j].TotalDonation > temp.TotalDonation {
			tempArr[j+1] = tempArr[j]
			j = j - 1
		}
		tempArr[j+1] = temp
	}
	fmt.Print("Setelah sorting: ")
	fmt.Println(tempArr)
}

