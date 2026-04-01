package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func input(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// Login

func login() bool {

	fmt.Println("====================================================")
	fmt.Println("Selamat Datang! Silahkan Login untuk melanjutkan")
	fmt.Println("====================================================")
	email := input("Masukkan Email : ")
	password := input("Masukkan Password : ")

	if email == "admin@gmail.com" && password == "admin" {
		fmt.Println("Login berhasil!")
		return true
	}

	fmt.Println("Email atau password salah!\n")
	return false
}

// Operator

func createOperator() {
	fmt.Println("Silahkan isi form berikut untuk menambah operator!")
	fmt.Println("Nama 						: ")
	fmt.Println("Email 						: ")
	fmt.Println("Password 					: ")
	fmt.Println("No. HP						: ")
	fmt.Println("Tanggal Join (DD/MM/YYYY)	: ")

	// setelah itu cek yang kosong
	// if ada yang kosong then
	fmt.Println("Input ada yang kosong, silahkan isi ulang!")
	// loop ke create operator

	// kalo berhasil
	fmt.Println("Operator berhasil ditambahkan\n")

	// masukin ke db operator
	// isi table operator_id (auto increment)

	// kembali ke menu operator

}

func listOperators() {
	// show table operator_id
	// show table name
	// show table email
	// show table password
	// show table phone
	// show table date_join
	// show table status

	fmt.Println("1. Back")
}

func updateOperator() {
	// show listoperators

	fmt.Println("Isi ID untuk update data Operator, contoh : 1")

	fmt.Println("Operator ID : 2") // tidak bisa diedit
	// setelah dipilih maka akan muncul data, user bisa melakukan edit
	fmt.Println("Nama : roi")
	fmt.Println("Email : roi@gmail.com ")
	fmt.Println("Password : roi123")
	fmt.Println("No. HP	: 08123456789")
	fmt.Println("Tanggal Join (DD/MM/YYYY) : 2/04/2026")
	fmt.Println("Status : Active")

	fmt.Print("1. Selesai Edit")
	fmt.Print("2. Back")

	// validasi sebelum edit
	fmt.Print("Apakah anda yakin dengan data yang sudah diedit? y/n")

	// if y
	fmt.Println("Data Operator dengan ID 2 telah diubah")

	// update ke db operator
	// kembali ke menu operator

	// if n show listoperators lagi

}

func deleteOperator() {
	// show listoperators

	fmt.Println("Isi ID untuk menghapus operator, contoh : 1")

	// validasi sebelum hapus
	fmt.Print("Apakah anda yakin ingin mengahpus? y/n")

	// if y
	fmt.Println("Operator dengan ID 3 berhasil dihapus")

	// if n show listoperators lagi

	// kalo id yang diisi tidak ada di list/ga sesuai
	fmt.Println("ID tidak ditemukan!")

	fmt.Println("1. Hapus lagi")
	fmt.Println("2. Back")

	// if hapus lagi show listoperators lagi (loop)

}

// Machine

func createMachine() {
	fmt.Println("Silahkan isi form berikut untuk menambah mesin!")
	fmt.Println("Kode Mesin : ")
	fmt.Println("Nama Mesin : ")
	fmt.Println("Status Mesin (Active/Inactive)	: ")

	// setelah itu cek yang kosong
	// if ada yang kosong then
	fmt.Println("Input ada yang kosong, silahkan isi ulang!")
	// loop ke create operator

	// kalo berhasil
	fmt.Println("Data Mesin berhasil ditambahkan\n")

	// create ke db mesin
	// isi table machine_id (auto increment)
	// isi table created_at dengan waktu saat ini

	// kembali ke menu operator

}

func listMachines() {
	// show table machine_id
	// show table machine_code
	// show table name
	// show table status
	// show table created_at

	fmt.Println("1. Back")
}

func updateMachines() {
	// show listMachines

	fmt.Println("Isi ID untuk update data Mesin, contoh : 1")

	fmt.Print("Machine ID : 3") // tidak bisa diedit
	// setelah dipilih maka akan muncul data, user bisa melakukan edit
	fmt.Println("Kode Mesin : 213")
	fmt.Println("Nama Mesin : Molding ")
	fmt.Println("Status : Working")

	fmt.Print("1. Selesai Edit")
	fmt.Print("2. Back")

	// validasi sebelum edit
	fmt.Print("Apakah anda yakin dengan data yang sudah diedit? y/n")

	// if y
	fmt.Println("Data Mesin dengan ID 3 telah diubah")

	// update ke db machines
	// kembali ke menu operator

	// if n show listoMachines lagi

}

func deleteMachines() {
	// show listMachines

	fmt.Println("Isi ID untuk menghapus operator, contoh : 1")

	// validasi sebelum hapus
	fmt.Print("Apakah anda yakin ingin mengahpus? y/n")

	// if y
	fmt.Println("Mesin dengan ID 5 berhasil dihapus")

	// if n show listMachines lagi

	// kalo id yang diisi tidak ada di list/ga sesuai
	fmt.Println("ID tidak ditemukan!")

	fmt.Println("1. Hapus lagi")
	fmt.Println("2. Back")

	// if hapus lagi show listMachines lagi (loop)

}

// Order

func createOrder() {
	fmt.Println("Silahkan isi form berikut untuk menambah Production Order!")
	fmt.Println("Kode Order 				: ")
	fmt.Println("Quantity Plan 				: ")
	fmt.Println("Quantity Actual			: ")
	fmt.Println("Status 					: ")
	fmt.Println("Start Date (DD/MM/YYYY)	: ")
	fmt.Println("End Date (DD/MM/YYYY)		: ")

	// setelah itu cek yang kosong
	// if ada yang kosong then
	fmt.Println("Input ada yang kosong, silahkan isi ulang!")
	// loop ke create order

	// kalo berhasil
	fmt.Println("Data Production Order berhasil ditambahkan\n")

	// create ke db Production Order
	// isi tabel production_order_id (auto increment)

	// kembali ke menu Production Order
}

func listOrders() {
	// show table production_order_id
	// show table order_code
	// show table product_id (FK DB Products)
	// show table qty_plan
	// show table qty_actual
	// show table status
	// show table start_date
	// show table end_date
	// show table created_at

	fmt.Println("1. Back")
}

func updateOrders() {
	// show listOrders

	fmt.Println("Isi ID untuk update data Production Orders, contoh : 1")

	fmt.Print("Production Orders ID : 4") // tidak bisa diedit
	// setelah dipilih maka akan muncul data, user bisa melakukan edit
	fmt.Println("Kode Order : 67")
	fmt.Println("Qty Plan : 100 Pcs ")
	fmt.Println("Qty Actual : 100 Pcs ")
	fmt.Println("Status : Finished")

	fmt.Print("1. Selesai Edit")
	fmt.Print("2. Back")

	// validasi sebelum edit
	fmt.Print("Apakah anda yakin dengan data yang sudah diedit? y/n")

	// if y
	fmt.Println("Data Production Orders dengan ID 3 telah diubah")

	// update ke db production_orders
	// kembali ke menu production orders

	// if n show listoMachines lagi

}

func deleteOrders() {
	// show listOrders

	fmt.Println("Isi ID untuk menghapus operator, contoh : 1")

	// validasi sebelum hapus
	fmt.Print("Apakah anda yakin ingin mengahpus? y/n")

	// if y
	fmt.Println("Mesin dengan ID 9 berhasil dihapus")

	// if n show listOrders lagi

	// kalo id yang diisi tidak ada di list/ga sesuai
	fmt.Println("ID tidak ditemukan!")

	fmt.Println("1. Hapus lagi")
	fmt.Println("2. Back")

	// if hapus lagi show listOrders lagi (loop)
}

// Inventory

func createInventory() {
	fmt.Println("Silahkan isi form berikut untuk menambah Inventory Transactions!")
	fmt.Println("Transaction Type 				: ")
	fmt.Println("Quantity 						: ")
	fmt.Println("Reference Type					: ")
	fmt.Println("Reference ID					: ")
	fmt.Println("Status 						: ")
	fmt.Println("Batch Number					: ")
	fmt.Println("Transaction Date (DD/MM/YYYY)	: ")

	// setelah itu cek yang kosong
	// if ada yang kosong then
	fmt.Println("Input ada yang kosong, silahkan isi ulang!")
	// loop ke create inventory

	// kalo berhasil
	fmt.Println("Data Inventory Transaction berhasil ditambahkan\n")

	// create ke db Inventory Transaction
	// isi tabel inventory_transaction_id (auto increment)

	// kembali ke menu Inventory Transaction
}

func listInventory() {
	// show table inventory_transaction_id
	// show table order_code
	// show table product_id (FK DB Products)
	// show table qty_plan
	// show table qty_actual
	// show table status
	// show table start_date
	// show table end_date
	// show table created_at

	fmt.Println("1. Back")
}

func updateInventory() {
	// show listInventory

	fmt.Println("Isi ID untuk update data Production Orders, contoh : 1")

	fmt.Print("Production Orders ID : 7") // tidak bisa diedit
	// setelah dipilih maka akan muncul data, user bisa melakukan edit
	fmt.Println("Transaction Type : Material")
	fmt.Println("Quantity : 50 Pcs ")
	fmt.Println("Reference Type : Raw")
	fmt.Println("Reference ID : 11 ")
	fmt.Println("Batch Number : 671")
	fmt.Println("Transaction Date : 2/04/2026")

	fmt.Print("1. Selesai Edit")
	fmt.Print("2. Back")

	// validasi sebelum edit
	fmt.Print("Apakah anda yakin dengan data yang sudah diedit? y/n")

	// if y
	fmt.Println("Data Inventory Transactions dengan ID 7 telah diubah")

	// update ke db inventory_transactions
	// kembali ke menu Inventory Transactions

	// if n show listInventory lagi
}

func deleteInventory() {
	// show listInventory

	fmt.Println("Isi ID untuk menghapus operator, contoh : 1")

	// validasi sebelum hapus
	fmt.Print("Apakah anda yakin ingin mengahpus? y/n")

	// if y
	fmt.Println("Mesin dengan ID 9 berhasil dihapus")

	// if n show listInventory lagi

	// kalo id yang diisi tidak ada di list/ga sesuai
	fmt.Println("ID tidak ditemukan!")

	fmt.Println("1. Hapus lagi")
	fmt.Println("2. Back")

	// if hapus lagi show listInventory lagi (loop)
}

// Report

func reportInventory() {

}

// Main Menu

func mainMenu() {
	for {
		fmt.Println("\n------------------")
		fmt.Println("Main Menu")
		fmt.Println("------------------")
		fmt.Println("Silahkan Pilih menu dengan memasukkan angka")
		fmt.Println("1. Operators")
		fmt.Println("2. Machines")
		fmt.Println("3. Production Orders")
		fmt.Println("4. Inventory")
		fmt.Println("5. Reports")
		fmt.Println("6. Exit")

		choice := input("Pilih : ")

		switch choice {
		case "1":
			operatorMenu()
		case "2":
			machineMenu()
		case "3":
			orderMenu()
		case "4":
			inventoryMenu()
		case "5":
			reportInventory()
		case "6":
			fmt.Println("Selamat Tinggal!\n")
			return
		default:
			fmt.Println("Input anda tidak valid!")
		}
	}
}

func operatorMenu() {
	for {
		fmt.Println("\n------------------")
		fmt.Println("Operator Menu")
		fmt.Println("------------------")
		fmt.Println("1. Create Account")
		fmt.Println("2. List Operator Account")
		fmt.Println("3. Update Operator")
		fmt.Println("4. Delete Account")
		fmt.Println("5. Back")

		c := input("Pilih : ")

		switch c {
		case "1":
			createOperator()
		case "2":
			listOperators()
		case "3":
			updateOperator()
		case "4":
			deleteOperator()
		case "5":
			return
		default:
			fmt.Println("Input anda tidak valid!")
		}
	}
}

func machineMenu() {
	for {
		fmt.Println("\n------------------")
		fmt.Println("Machine Menu")
		fmt.Println("------------------")
		fmt.Println("1. Add Machine")
		fmt.Println("2. List Machine")
		fmt.Println("3. Update Machine")
		fmt.Println("4. Delete Machine")
		fmt.Println("5. Back")

		c := input("Pilih : ")

		switch c {
		case "1":
			createMachine()
		case "2":
			listMachines()
		case "3":
			updateMachines()
		case "4":
			deleteMachines()
		case "5":
			return
		default:
			fmt.Println("Input anda tidak valid!")
		}
	}
}

func orderMenu() {
	for {
		fmt.Println("\n------------------")
		fmt.Println("Order Menu")
		fmt.Println("------------------")
		fmt.Println("1. Add Production Orders")
		fmt.Println("2. List Production Orders")
		fmt.Println("3. Update Production Orders")
		fmt.Println("4. Delete Production Orders")
		fmt.Println("5. Back")

		c := input("Pilih : ")

		switch c {
		case "1":
			createOrder()
		case "2":
			listOrders()
		case "3":
			updateOrders()
		case "4":
			deleteOrders()
		case "5":
			return
		default:
			fmt.Println("Input anda tidak valid!")
		}
	}
}

func inventoryMenu() {
	for {
		fmt.Println("\n------------------")
		fmt.Println("Inventory Transaction Menu")
		fmt.Println("------------------")
		fmt.Println("1. Create Inventory Transaction")
		fmt.Println("2. List Inventory Transaction")
		fmt.Println("3. Update Inventory Transaction")
		fmt.Println("4. Delete Inventory Transaction")
		fmt.Println("5. Back")

		c := input("Pilih : ")

		switch c {
		case "1":
			createInventory()
		case "2":
			listInventory()
		case "3":
			updateInventory()
		case "4":
			deleteInventory()
		case "5":
			return
		default:
			fmt.Println("Input anda tidak valid!")
		}
	}
}

func RunApp() {
	for {
		if login() {
			mainMenu()
		}
	}
}
