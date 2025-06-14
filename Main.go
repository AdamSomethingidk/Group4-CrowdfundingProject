package main

import (
	"fmt"
	"strings"
)

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
	fmt.Println("8. Tampilkan total donasi masuk pada proyek")
	fmt.Println("9. Urutkan proyek berdasarkan total dana terkumpul (Selection)")
	fmt.Println("10. Urutkan proyek berdasarkan total dana terkumpul (Insertion)")
	fmt.Println("11. Urutkan proyek berdasarkan jumlah donatur (Selection)")
	fmt.Println("12. Urutkan proyek berdasarkan jumlah donatur (Insertion)")
	fmt.Println("13. Tampilkan proyek yang telah memenuhi target pendanaan")
	fmt.Println("0. Keluar")
	fmt.Println("=============================================")
	fmt.Print("Masukan Pilihan: ")
}

func main() {
	var n, currentSize int
	n = 100
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
		if n == 4 {
			donateToProject()
		}
		if n == 5 {
			findProjectByNameSeq()
		}
		if n == 6 {
			findProjectNameBin()
		}
		if n == 7 {
			findProjectCategorySeq()
		}
		if n == 8 {
			showTotalDonations()
		}
		if n == 9 {
			sortTotalDonationSelection()
		}
		if n == 10 {
			sortTotalDonationInsert()
		}
		if n == 11 {
			sortTotalDonatorSelection() 
 		}
		if n == 12 {
			sortTotalDonatorInsert()
		}
		if n == 13 {
			showCompletedProjects()
		}
	}
}

// Buat proyek baru
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

// Ubah proyek
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

// Hapus proyek
func deleteProject(currentSize *int) {
	var name, verify, projectToDelete string
	fmt.Print("Masukan nama anda: ")
	fmt.Scan(&name)
	var found bool
	found = false

	// Display projects owned by the user
	fmt.Println("Proyek Anda:")
	for i := 0; i < *currentSize; i++ {
		if mainArr[i].Creator == name {
			fmt.Println("----------------------")
			fmt.Println("Nama proyek: ", mainArr[i].ProjectName)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada proyek yang ditemukan untuk nama ini.")
		return
	}

	fmt.Print("Pilih proyek yang akan dihapus (gunakan nama): ")
	fmt.Scan(&projectToDelete)

	// Find and delete the project
	for j := 0; j < *currentSize; j++ {
		if mainArr[j].ProjectName == projectToDelete && mainArr[j].Creator == name {
			fmt.Print("Proyek ", mainArr[j].ProjectName, " akan dihapus, apakah anda yakin? Y/N: ")
			fmt.Scan(&verify)
			if verify == "y" || verify == "Y" {
				// Shift elements to the left to remove the project
				for k := j; k < *currentSize-1; k++ {
					mainArr[k] = mainArr[k+1]
				}
				// Clear the last element
				mainArr[*currentSize-1] = Project{}
				fmt.Println("Proyek telah dihapus.")
				*currentSize = *currentSize - 1
			} else {
				fmt.Println("Penghapusan proyek dibatalkan.")
			}
			return
		}
	}

	fmt.Println("Proyek tidak ditemukan atau bukan proyek anda.")
}

// Berikan donasi ke proyek
func donateToProject() {
	var projectName string
	var donationAmount float64
	var found bool
	found = false
	fmt.Print("Masukkan nama proyek yang ingin Anda donasi: ")
	fmt.Scan(&projectName)
	fmt.Print("Masukkan jumlah donasi: ")
	fmt.Scan(&donationAmount)
	for i := 0; i < len(mainArr); i++ {
		if mainArr[i].ProjectName == projectName {
			mainArr[i].TotalDonation += donationAmount
			mainArr[i].TotalDonator += 1
			fmt.Printf("Terima kasih atas donasi Anda sebesar %.2f untuk proyek %s\n", donationAmount, projectName)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Proyek tidak ditemukan.")
	}
}

// Tampilkan total donasi masuk pada proyek
func showTotalDonations() {
	fmt.Println("=============================================")
	fmt.Println("          TOTAL DONASI MASUK PER PROYEK     ")
	fmt.Println("=============================================")
	for i := 0; i < len(mainArr); i++ {
		if mainArr[i].ProjectName != "" { // Check if the project exists
			fmt.Printf("Nama Proyek: %s, Total Donasi: %.2f\n", mainArr[i].ProjectName, mainArr[i].TotalDonation)
		}
	}
	fmt.Println("=============================================")
}

// Cari proyek berdasarkan nama (binary)
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
	high = len(tempArr) - 1
	for high >= low {
		mid = (low + high) / 2
		if tempArr[mid].ProjectName == pickedName {
			index = mid
			break
		} else if tempArr[mid].ProjectName > pickedName {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if index == -1 {
		fmt.Print("Proyek dengan nama ", pickedName, " tidak ditemukan.")
	} else {
		fmt.Println(tempArr[index])
	}
}

// Cari proyek berdasarkan nama (sequential)
func findProjectByNameSeq() {
	var pickedName string
	fmt.Print("Masukkan nama proyek yang ingin dicari: ")
	fmt.Scan(&pickedName)

	var found bool
	found = false
	for i := 0; i < len(mainArr); i++ {
		if mainArr[i].ProjectName == pickedName {
			fmt.Println("Proyek ditemukan:")
			fmt.Println("Nama Pemilik Proyek: ", mainArr[i].Creator)
			fmt.Println("Nama Proyek: ", mainArr[i].ProjectName)
			fmt.Println("Kategori Proyek: ", mainArr[i].Category)
			fmt.Println("Jumlah Donatur: ", mainArr[i].TotalDonator)
			fmt.Println("Target Dana: ", mainArr[i].TargetDonation)
			fmt.Println("Dana yang Terkumpul: ", mainArr[i].TotalDonation)
			found = true
			break
		}
	}
	
	if !found {
		fmt.Println("Proyek dengan nama", pickedName, "tidak ditemukan.")
	}
}
 
 
// Cari proyek berdasarkan kategori (sequential)
  func findProjectCategorySeq() {
	var pickedCategory string
	fmt.Print("Masukkan kategori proyek yang ingin dicari: ")
	fmt.Scan(&pickedCategory)
	// Normalize the input category to lowercase and trim spaces
	pickedCategory = strings.ToLower(strings.TrimSpace(pickedCategory))

	var found bool
	found = false

	for i := 0; i < len(mainArr); i++ {
		// Normalize the project category to lowercase and trim spaces for comparison
		if strings.ToLower(strings.TrimSpace(mainArr[i].Category)) == pickedCategory {
			if !found {
				fmt.Println("=============================================")
				fmt.Println("Proyek yang ditemukan dalam kategori:", pickedCategory)
				fmt.Println("=============================================")
				found = true
			}
			fmt.Println("Nama Pemilik Proyek: ", mainArr[i].Creator)
			fmt.Println("Nama Proyek: ", mainArr[i].ProjectName)
			fmt.Println("Kategori Proyek: ", mainArr[i].Category)
			fmt.Println("Jumlah Donatur: ", mainArr[i].TotalDonator)
			fmt.Println("Target Dana: ", mainArr[i].TargetDonation)
			fmt.Println("Dana yang Terkumpul: ", mainArr[i].TotalDonation)
			fmt.Println("---------------------------------------------")
		}
	}

	if !found {
		fmt.Println("Tidak ada proyek ditemukan untuk kategori:", pickedCategory)
	}
}

// Urutkan proyek berdasarkan total dana terkumpul (Selection)
func sortTotalDonationSelection() {
	// Create a copy of the mainArr to sort
	tempArr := make([]Project, len(mainArr))
	copy(tempArr, mainArr[:]) // Copy the mainArr to tempArr

	fmt.Println("Sebelum sorting:")
	for _, project := range tempArr {
		if project.ProjectName != "" { // Only print existing projects
			fmt.Printf("Nama Proyek: %s, Total Donasi: %.2f\n", project.ProjectName, project.TotalDonation)
		}
	}

	// Selection Sort based on TotalDonation
	for i := 0; i < len(tempArr)-1; i++ {
		// Find the index of the minimum element in the unsorted part
		minIndex := i
		for j := i + 1; j < len(tempArr); j++ {
			if tempArr[j].TotalDonation < tempArr[minIndex].TotalDonation {
				minIndex = j
			}
		}
		// Swap the found minimum element with the first element
		tempArr[i], tempArr[minIndex] = tempArr[minIndex], tempArr[i]
	}

	fmt.Println("Setelah sorting:")
	for _, project := range tempArr {
		if project.ProjectName != "" { // Only print existing projects
			fmt.Printf("Nama Proyek: %s, Total Donasi: %.2f\n", project.ProjectName, project.TotalDonation)
		}
	}
}


// Urutkan proyek berdasarkan total dana terkumpul (Insertion)
func sortTotalDonationInsert() {
	var tempArr = [10]Project{}
	var temp Project
	var j int
	tempArr = mainArr
	fmt.Print("Sebelum sorting: ")
	fmt.Println(tempArr)
	for i := 1; i < len(tempArr); i++ {
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

// Urutkan proyek berdasarkan jumlah donatur (Selection)
func sortTotalDonatorSelection() {
	// Create a copy of the mainArr to sort
	tempArr := make([]Project, len(mainArr))
	copy(tempArr, mainArr[:]) // Copy the mainArr to tempArr

	fmt.Println("Sebelum sorting:")
	for _, project := range tempArr {
		if project.ProjectName != "" { // Only print existing projects
			fmt.Printf("Nama Proyek: %s, Jumlah Donatur: %d\n", project.ProjectName, project.TotalDonator)
		}
	}

	// Selection Sort based on TotalDonator
	for i := 0; i < len(tempArr)-1; i++ {
		// Find the index of the maximum element in the unsorted part
		maxIndex := i
		for j := i + 1; j < len(tempArr); j++ {
			if tempArr[j].TotalDonator > tempArr[maxIndex].TotalDonator {
				maxIndex = j
			}
		}
		// Swap the found maximum element with the first element
		tempArr[i], tempArr[maxIndex] = tempArr[maxIndex], tempArr[i]
	}

	fmt.Println("Setelah sorting:")
	for _, project := range tempArr {
		if project.ProjectName != "" { // Only print existing projects
			fmt.Printf("Nama Proyek: %s, Jumlah Donatur: %d\n", project.ProjectName, project.TotalDonator)
		}
	}
}

// Urutkan proyek berdasarkan jumlah donatur (Insertion)
func sortTotalDonatorInsert() {
	// Create a copy of the mainArr to sort
	tempArr := make([]Project, len(mainArr))
	copy(tempArr, mainArr[:]) // Copy the mainArr to tempArr

	fmt.Println("Sebelum sorting:")
	for _, project := range tempArr {
		if project.ProjectName != "" { // Only print existing projects
			fmt.Printf("Nama Proyek: %s, Jumlah Donatur: %d\n", project.ProjectName, project.TotalDonator)
		}
	}

	// Insertion Sort based on TotalDonator
	for i := 1; i < len(tempArr); i++ {
		key := tempArr[i]
		j := i - 1

		// Move elements of tempArr[0..i-1], that are greater than key.TotalDonator,
		// to one position ahead of their current position
		for j >= 0 && tempArr[j].TotalDonator > key.TotalDonator {
			tempArr[j+1] = tempArr[j]
			j = j - 1
		}
		tempArr[j+1] = key
	}

	fmt.Println("Setelah sorting:")
	for _, project := range tempArr {
		if project.ProjectName != "" { // Only print existing projects
			fmt.Printf("Nama Proyek: %s, Jumlah Donatur: %d\n", project.ProjectName, project.TotalDonator)
		}
	}
}


// Tampilkan proyek yang telah memenuhi target pendanaan
func showCompletedProjects() {
	for i := 0; i < len(mainArr); i++ {
		if mainArr[i].TotalDonation >= mainArr[i].TargetDonation {
			fmt.Println("Nama proyek: ", mainArr[i].ProjectName)
			fmt.Println("Kategori proyek: ", mainArr[i].Category)
			fmt.Println("Target dana: ", mainArr[i].TargetDonation)
			fmt.Println("Dana yang terkumpul: ", mainArr[i].TotalDonation)
			fmt.Println("Jumlah donatur: ", mainArr[i].TotalDonator)
		}
	}
}
