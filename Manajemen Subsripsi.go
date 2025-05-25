package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MaxSubscriptions = 100
	MaxTransactions  = 100
)

// Tipe Bentukan untuk Subscription dan Transaction
type Subscription struct {
	ID        int
	Name      string
	Cost      float64
	StartDate string // format YYYY-MM-DD
	Category  string
}

type Transaction struct {
	ID          int
	Description string
	Amount      float64
	Date        string // format YYYY-MM-DD
	Type        string // "income" atau "expense"
}

// Variabel global: array utama dan counter
var subscriptions [MaxSubscriptions]Subscription
var subCount int

var transactions [MaxTransactions]Transaction
var txCount int

var reader = bufio.NewReader(os.Stdin)

func main() {
	running := true
	for running {
		showMenu()
		choice := readInt("Pilih menu: ")
		switch choice {
		case 1:
			AddSubscription(&subscriptions, &subCount)
		case 2:
			ListSubscriptions(&subscriptions, subCount)
		case 3:
			EditSubscription(&subscriptions, subCount)
		case 4:
			DeleteSubscription(&subscriptions, &subCount)
		case 5:
			AddTransaction(&transactions, &txCount)
		case 6:
			ListTransactions(&transactions, txCount)
		case 7:
			EditTransaction(&transactions, txCount)
		case 8:
			DeleteTransaction(&transactions, &txCount)
		case 9:
			running = false
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
	fmt.Println("Terima kasih telah menggunakan aplikasi.")
}

func showMenu() {
	fmt.Println("\n=== Aplikasi Manajemen Subskripsi & Keuangan Pribadi ===")
	fmt.Println("1. Tambah Subscription")
	fmt.Println("2. Lihat Subscription")
	fmt.Println("3. Edit Subscription")
	fmt.Println("4. Hapus Subscription")
	fmt.Println("5. Tambah Transaction")
	fmt.Println("6. Lihat Transaction")
	fmt.Println("7. Edit Transaction")
	fmt.Println("8. Hapus Transaction")
	fmt.Println("9. Keluar")
}

// ----------------------------------------
// Subscriptions
// ----------------------------------------
func AddSubscription(arr *[MaxSubscriptions]Subscription, count *int) {
	if *count >= MaxSubscriptions {
		fmt.Println("Kapasitas maksimum subscription tercapai.")
		return
	}
	idx := *count
	sub := &arr[idx]
	sub.ID = idx + 1
	sub.Name = readString("Nama layanan: ")
	sub.Cost = readFloat("Biaya per periode: ")
	sub.StartDate = readString("Tanggal mulai (YYYY-MM-DD): ")
	sub.Category = readString("Kategori: ")
	*count++
	fmt.Println("Subscription berhasil ditambahkan.")
}

func ListSubscriptions(arr *[MaxSubscriptions]Subscription, count int) {
	if count == 0 {
		fmt.Println("Belum ada subscription.")
		return
	}
	field := readString("Urutkan berdasarkan (name/cost/date/category): ")
	order := readString("Urutan (asc/desc): ")
	SelectionSortSubs(arr, count, field, order)
	fmt.Println("\nDaftar Subscription:")
	for i := 0; i < count; i++ {
		sub := arr[i]
		fmt.Printf("%d. %s | %.2f | %s | %s\n", sub.ID, sub.Name, sub.Cost, sub.StartDate, sub.Category)
	}
}

func EditSubscription(arr *[MaxSubscriptions]Subscription, count int) {
	key := readString("Masukkan nama layanan yang ingin diedit: ")
	idx := SequentialSearchSubs(arr, count, key)
	if idx == -1 {
		fmt.Println("Subscription tidak ditemukan.")
		return
	}
	sub := &arr[idx]
	sub.Name = readString("Nama layanan: ")
	sub.Cost = readFloat("Biaya per periode: ")
	sub.StartDate = readString("Tanggal mulai (YYYY-MM-DD): ")
	sub.Category = readString("Kategori: ")
	fmt.Println("Subscription berhasil diupdate.")
}

func DeleteSubscription(arr *[MaxSubscriptions]Subscription, count *int) {
	key := readString("Masukkan nama layanan yang ingin dihapus: ")
	idx := SequentialSearchSubs(arr, *count, key)
	if idx == -1 {
		fmt.Println("Subscription tidak ditemukan.")
		return
	}
	for i := idx; i < *count-1; i++ {
		arr[i] = arr[i+1]
	}
	*count--
	fmt.Println("Subscription berhasil dihapus.")
}

// ----------------------------------------
// Transactions
// ----------------------------------------
func AddTransaction(arr *[MaxTransactions]Transaction, count *int) {
	if *count >= MaxTransactions {
		fmt.Println("Kapasitas maksimum transaksi tercapai.")
		return
	}
	idx := *count
	tx := &arr[idx]
	tx.ID = idx + 1
	tx.Description = readString("Deskripsi transaksi: ")
	tx.Amount = readFloat("Jumlah (positif income, negatif expense): ")
	tx.Date = readString("Tanggal (YYYY-MM-DD): ")
	tx.Type = readString("Tipe (income/expense): ")
	*count++
	fmt.Println("Transaksi berhasil ditambahkan.")
}

func ListTransactions(arr *[MaxTransactions]Transaction, count int) {
	if count == 0 {
		fmt.Println("Belum ada transaksi.")
		return
	}
	field := readString("Urutkan berdasarkan (description/amount/date/type): ")
	order := readString("Urutan (asc/desc): ")
	SelectionSortTx(arr, count, field, order)
	fmt.Println("\nDaftar Transaksi:")
	for i := 0; i < count; i++ {
		tx := arr[i]
		fmt.Printf("%d. %s | %.2f | %s | %s\n", tx.ID, tx.Description, tx.Amount, tx.Date, tx.Type)
	}
}

func EditTransaction(arr *[MaxTransactions]Transaction, count int) {
	key := readString("Masukkan deskripsi transaksi yang ingin diedit: ")
	idx := SequentialSearchTx(arr, count, key)
	if idx == -1 {
		fmt.Println("Transaksi tidak ditemukan.")
		return
	}
	tx := &arr[idx]
	tx.Description = readString("Deskripsi transaksi: ")
	tx.Amount = readFloat("Jumlah: ")
	tx.Date = readString("Tanggal (YYYY-MM-DD): ")
	tx.Type = readString("Tipe (income/expense): ")
	fmt.Println("Transaksi berhasil diupdate.")
}

func DeleteTransaction(arr *[MaxTransactions]Transaction, count *int) {
	key := readString("Masukkan deskripsi transaksi yang ingin dihapus: ")
	idx := SequentialSearchTx(arr, *count, key)
	if idx == -1 {
		fmt.Println("Transaksi tidak ditemukan.")
		return
	}
	for i := idx; i < *count-1; i++ {
		arr[i] = arr[i+1]
	}
	*count--
	fmt.Println("Transaksi berhasil dihapus.")
}

// ----------------------------------------
// Search Functions
// ----------------------------------------
func SequentialSearchSubs(arr *[MaxSubscriptions]Subscription, count int, key string) int {
	for i := 0; i < count; i++ {
		if strings.EqualFold(arr[i].Name, key) {
			return i
		}
	}
	return -1
}

func SequentialSearchTx(arr *[MaxTransactions]Transaction, count int, key string) int {
	for i := 0; i < count; i++ {
		if strings.EqualFold(arr[i].Description, key) {
			return i
		}
	}
	return -1
}

func BinarySearchSubs(arr *[MaxSubscriptions]Subscription, count int, key string) int {
	low, high := 0, count-1
	for low <= high {
		mid := (low + high) / 2
		if strings.EqualFold(arr[mid].Name, key) {
			return mid
		} else if strings.ToLower(arr[mid].Name) < strings.ToLower(key) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func BinarySearchTx(arr *[MaxTransactions]Transaction, count int, key string) int {
	low, high := 0, count-1
	for low <= high {
		mid := (low + high) / 2
		if strings.EqualFold(arr[mid].Description, key) {
			return mid
		} else if strings.ToLower(arr[mid].Description) < strings.ToLower(key) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// ----------------------------------------
// Sorting Functions
// ----------------------------------------
func SelectionSortSubs(arr *[MaxSubscriptions]Subscription, count int, field, order string) {
	for i := 0; i < count-1; i++ {
		idx := i
		for j := i + 1; j < count; j++ {
			if compareSubs(arr[j], arr[idx], field, order) {
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

func SelectionSortTx(arr *[MaxTransactions]Transaction, count int, field, order string) {
	for i := 0; i < count-1; i++ {
		idx := i
		for j := i + 1; j < count; j++ {
			if compareTx(arr[j], arr[idx], field, order) {
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

func InsertionSortSubs(arr *[MaxSubscriptions]Subscription, count int, field, order string) {
	for i := 1; i < count; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && compareSubs(key, arr[j], field, order) {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func InsertionSortTx(arr *[MaxTransactions]Transaction, count int, field, order string) {
	for i := 1; i < count; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && compareTx(key, arr[j], field, order) {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func compareSubs(a, b Subscription, field, order string) bool {
	var less bool
	switch field {
	case "name":
		less = strings.ToLower(a.Name) < strings.ToLower(b.Name)
	case "cost":
		less = a.Cost < b.Cost
	case "date":
		less = a.StartDate < b.StartDate
	case "category":
		less = strings.ToLower(a.Category) < strings.ToLower(b.Category)
	default:
		less = a.ID < b.ID
	}
	if order == "asc" {
		return less
	}
	return !less
}

func compareTx(a, b Transaction, field, order string) bool {
	var less bool
	switch field {
	case "description":
		less = strings.ToLower(a.Description) < strings.ToLower(b.Description)
	case "amount":
		less = a.Amount < b.Amount
	case "date":
		less = a.Date < b.Date
	case "type":
		less = strings.ToLower(a.Type) < strings.ToLower(b.Type)
	default:
		less = a.ID < b.ID
	}
	if order == "asc" {
		return less
	}
	return !less
}

// ----------------------------------------
// Helper Input Functions
// ----------------------------------------
func readString(prompt string) string {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			return input
		}
		fmt.Println("Input tidak boleh kosong.")
	}
}

func readInt(prompt string) int {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		n, err := strconv.Atoi(input)
		if err == nil {
			return n
		}
		fmt.Println("Masukkan angka yang valid.")
	}
}

func readFloat(prompt string) float64 {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		f, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return f
		}
		fmt.Println("Masukkan angka desimal yang valid.")
	}
}
