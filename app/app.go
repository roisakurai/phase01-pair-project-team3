package app

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"p1ppteam3/config"
	"p1ppteam3/repositories"
	"strings"
)

func RunApp() {
	db := config.ConnectDB()
	defer db.Close()

	config.Migrate(db)

	for {
		success, name := login(db)
		if success {
			mainMenu(db, name)
		}
	}
}

var reader = bufio.NewReader(os.Stdin)

func input(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// Login

func login(db *sql.DB) (bool, string) {
	fmt.Println("====================================================")
	fmt.Println("Selamat Datang! Silahkan Login untuk melanjutkan")
	fmt.Println("====================================================")

	email := input("Masukkan Email : ")
	password := input("Masukkan Password : ")

	var storedPassword string
	var name string

	err := db.QueryRow(
		"SELECT name, password FROM operators WHERE email = ? AND status = 'active'",
		email,
	).Scan(&name, &storedPassword)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Email tidak ditemukan!\n")
			return false, ""
		}
		fmt.Println("Error:", err)
		return false, ""
	}

	if password != storedPassword {
		fmt.Println("Password salah!\n")
		return false, ""
	}

	fmt.Println("Login berhasil!")
	return true, name
}

// Operator

func createOperator() {
	fmt.Println("Silahkan isi form berikut untuk menambah operator!")
	fmt.Println("Nama : ")
	fmt.Println("Email : ")
	fmt.Println("Password : ")
	fmt.Println("No. HP : ")
	fmt.Println("Tanggal Join (DD/MM/YYYY) : ")

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

	fmt.Println("Isi ID untuk update data Operator")

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

	fmt.Println("Isi ID untuk menghapus operator")

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

// Product
// func createProduct() {
// 	fmt.Println("Silahkan isi form berikut untuk menambah Produk!")
// 	fmt.Println("SKU : ")
// 	fmt.Println("Nama Produk : ")
// 	fmt.Println("Type (1. finished, 2. raw, 3. semi-finished) : ")
// 	fmt.Println("Unit : ")
// 	fmt.Println("Standard Cost : ")

// 	// setelah itu cek yang kosong
// 	// if ada yang kosong then
// 	fmt.Println("Input ada yang kosong, silahkan isi ulang!")
// 	// loop ke create product

// 	// kalo berhasil
// 	fmt.Println("Data Produk berhasil ditambahkan\n")

// 	// masukin ke table products
// 	// isi table products_id (auto increment)

// 	// kembali ke menu operator

// }

func createProduct(db *sql.DB) {
	fmt.Println("=== Tambah Produk ===")

	sku := input("SKU: ")
	name := input("Nama Produk: ")
	typeInput := input("Type (finished/raw/semi-finished): ")
	unit := input("Unit: ")
	costInput := input("Standard Cost: ")

	// validasi basic
	if sku == "" || name == "" || typeInput == "" || unit == "" || costInput == "" {
		fmt.Println("Semua field wajib diisi!\n")
		return
	}

	// convert cost ke float
	var cost float64
	_, err := fmt.Sscanf(costInput, "%f", &cost)
	if err != nil {
		fmt.Println("Standard cost harus angka!\n")
		return
	}

	err = repositories.CreateProduct(db, sku, name, typeInput, unit, cost)
	if err != nil {
		fmt.Println("Gagal menambahkan produk:", err)
		return
	}

	fmt.Println("Produk berhasil ditambahkan!\n")
}

// func listProduct() {
// 	// show table product_id
// 	// show table sku
// 	// show table name
// 	// show table type
// 	// show table unit
// 	// show table standard_cost
// 	// show table created_at
// 	// show table updated_at

// 	fmt.Println("1. Back")
// }

func listProduct(db *sql.DB) {
	fmt.Println("\n=== List Product ===")

	products, err := repositories.ListProducts(db)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(products) == 0 {
		fmt.Println("Tidak ada data produk\n")
		return
	}

	fmt.Printf("%-5s %-10s %-25s %-15s %-10s %-15s %-20s %-20s\n",
		"ID", "SKU", "Name", "Type", "Unit", "Cost", "Created At", "Updated At")

	for _, p := range products {
		fmt.Printf("%-5d %-10s %-25s %-15s %-10s %-15.2f %-20s %-20s\n",
			p.ID, p.SKU, p.Name, p.Type, p.Unit, p.Cost, p.CreatedAt, p.UpdatedAt)
	}
}

// func updateProduct() {
// 	// show listproduct

// 	fmt.Println("Isi ID untuk update data Produk")

// 	fmt.Println("Operator ID : 1") // tidak bisa diedit
// 	// setelah dipilih maka akan muncul data, user bisa melakukan edit
// 	fmt.Println("SKU : FG-001")
// 	fmt.Println("Produk : Roti Tawar Gandum ")
// 	fmt.Println("Type : finished")
// 	fmt.Println("Unit : pcs")
// 	fmt.Println("Standard Cost : 12000.00")

// 	fmt.Print("1. Selesai Edit")
// 	fmt.Print("2. Back")

// 	// validasi sebelum edit
// 	fmt.Print("Apakah anda yakin dengan data yang sudah diedit? y/n")

// 	// if y
// 	fmt.Println("Data Operator dengan ID 2 telah diubah")

// 	// update ke table product
// 	// isi table created_at dengan CURRENT_TIMESTAMP

// 	// kembali ke menu product

// 	// if n show listproduct lagi

// }

func updateProduct(db *sql.DB) {
	fmt.Println("=== Update Produk ===")

	// input ID
	idInput := input("Masukkan Product ID: ")
	var id int
	_, err := fmt.Sscanf(idInput, "%d", &id)
	if err != nil {
		fmt.Println("ID harus angka!\n")
		return
	}

	// ambil data lama (optional tapi bagus UX)
	var sku, name, productType, unit string
	var cost float64

	err = db.QueryRow(`
		SELECT sku, name, type, unit, standard_cost
		FROM products
		WHERE product_id = ?
	`, id).Scan(&sku, &name, &productType, &unit, &cost)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Product tidak ditemukan!\n")
			return
		}
		fmt.Println("Error:", err)
		return
	}

	// tampilkan data lama
	fmt.Println("\nData saat ini:")
	fmt.Println("SKU:", sku)
	fmt.Println("Nama:", name)
	fmt.Println("Type:", productType)
	fmt.Println("Unit:", unit)
	fmt.Println("Cost:", cost)

	// input baru (boleh kosong → pakai lama)
	newSKU := input("SKU baru (kosongkan jika tidak diubah): ")
	newName := input("Nama baru: ")
	newType := input("Type baru: ")
	newUnit := input("Unit baru: ")
	newCostInput := input("Cost baru: ")

	if newSKU != "" {
		sku = newSKU
	}
	if newName != "" {
		name = newName
	}
	if newType != "" {
		productType = newType
	}
	if newUnit != "" {
		unit = newUnit
	}
	if newCostInput != "" {
		_, err := fmt.Sscanf(newCostInput, "%f", &cost)
		if err != nil {
			fmt.Println("Cost harus angka!\n")
			return
		}
	}

	// konfirmasi
	confirm := input("Apakah anda sudah yakin dengan updatenya? (y/n): ")
	if confirm != "y" {
		fmt.Println("Update dibatalkan\n")
		return
	}

	// update ke DB
	err = repositories.UpdateProduct(db, id, sku, name, productType, unit, cost)
	if err != nil {
		fmt.Println("Gagal update:", err)
		return
	}

	fmt.Println("Produk berhasil diupdate!\n")
}

// func deleteProduct() {
// 	// show listproduct

// 	fmt.Println("Isi ID untuk menghapus Produk, ")

// 	// validasi sebelum hapus
// 	fmt.Print("Apakah anda yakin ingin mengahpus? y/n")

// 	// if y
// 	fmt.Println("Produk dengan ID 3 berhasil dihapus")

// 	// if n show listproduk lagi

// 	// kalo id yang diisi tidak ada di list/ga sesuai
// 	fmt.Println("ID tidak ditemukan!")

// 	fmt.Println("1. Hapus lagi")
// 	fmt.Println("2. Back")

// 	// if hapus lagi show listproduk lagi (loop)

// }

func deleteProduct(db *sql.DB) {
	fmt.Println("=== Hapus Produk ===")

	for {
		// tampilkan list dulu (opsional tapi bagus UX)
		listProduct(db)

		idInput := input("Masukkan Product ID yang ingin dihapus: ")

		var id int
		_, err := fmt.Sscanf(idInput, "%d", &id)
		if err != nil {
			fmt.Println("ID harus angka!\n")
			continue
		}

		confirm := input("Apakah anda yakin ingin menghapus? (y/n): ")
		if confirm != "y" {
			fmt.Println("Penghapusan dibatalkan\n")
			return
		}

		rows, err := repositories.DeleteProduct(db, id)
		if err != nil {
			fmt.Println("Gagal menghapus:", err)
			return
		}

		if rows == 0 {
			fmt.Println("ID tidak ditemukan!\n")
		} else {
			fmt.Printf("Produk dengan ID %d berhasil dihapus\n\n", id)
		}

		// pilihan setelah hapus
		fmt.Println("1. Hapus lagi")
		fmt.Println("2. Back")

		choice := input("Pilih: ")
		if choice == "2" {
			return
		}
	}
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

	fmt.Println("Isi ID untuk update data Mesin")

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

	fmt.Println("Isi ID untuk menghapus operator")

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
	fmt.Println("Kode Order : ")
	fmt.Println("Quantity Plan : ")
	fmt.Println("Quantity Actual : ")
	fmt.Println("Status : ")
	fmt.Println("Start Date (DD/MM/YYYY) : ")
	fmt.Println("End Date (DD/MM/YYYY) : ")

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

	fmt.Println("Isi ID untuk update data Production Orders")

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

	fmt.Println("Isi ID untuk menghapus operator")

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
	fmt.Println("Transaction Type : ")
	fmt.Println("Quantity : ")
	fmt.Println("Reference Type : ")
	fmt.Println("Reference ID : ")
	fmt.Println("Status : ")
	fmt.Println("Batch Number : ")
	fmt.Println("Transaction Date (DD/MM/YYYY) : ")

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

	fmt.Println("Isi ID untuk update data Production Orders")

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

	fmt.Println("Isi ID untuk menghapus operator")

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

func mainMenu(db *sql.DB, userName string) {
	for {
		fmt.Println("\n------------------")
		fmt.Println("Main Menu")
		fmt.Println("------------------")
		fmt.Printf("Hai %s! Silahkan Pilih menu dengan memasukkan angka\n", userName)
		fmt.Println("1. Operators")
		fmt.Println("2. Machines")
		fmt.Println("3. Products")
		fmt.Println("4. Production Orders")
		fmt.Println("5. Inventory")
		fmt.Println("6. Reports")
		fmt.Println("7. Exit")

		choice := input("Pilih : ")

		switch choice {
		case "1":
			operatorMenu(db)
		case "2":
			machineMenu()
		case "3":
			productMenu(db)
		case "4":
			orderMenu()
		case "5":
			inventoryMenu()
		case "6":
			reportInventory()
		case "7":
			fmt.Println("Selamat Tinggal!\n")
			return
		default:
			fmt.Println("Input anda tidak valid!")
		}
	}
}

func operatorMenu(db *sql.DB) {
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
			repositories.CreateOperator(db)
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

func productMenu(db *sql.DB) {
	for {
		fmt.Println("\n------------------")
		fmt.Println("Product Menu")
		fmt.Println("------------------")
		fmt.Println("1. Add Product")
		fmt.Println("2. List Product")
		fmt.Println("3. Update Product")
		fmt.Println("4. Delete Product")
		fmt.Println("5. Back")

		c := input("Pilih : ")

		switch c {
		case "1":
			createProduct(db)
		case "2":
			listProduct(db)
		case "3":
			updateProduct(db)
		case "4":
			deleteProduct(db)
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
		fmt.Println("Production Orders Menu")
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
