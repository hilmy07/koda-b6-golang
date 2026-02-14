package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("Yoshinoya Menu")
	fmt.Println("1. Yakiniku Beef Bowl\t\t\t\t\tIDR 61.000")
	fmt.Println("2. Beef Wakame Udon Soup\t\t\t\tIDR 79.000")
	fmt.Println("3. Es Kopcen Kopi Susu\t\t\t\t\tIDR 38000")
	fmt.Println("4. Blackpepper Beef Bowl + Onsen Egg\t\t\tIDR 78000")
	fmt.Println("5. Red Hot Chilli Tori Don \t\t\t\tIDR 47.000")
	fmt.Println("6. Ebi Shrimp Bowl\t\t\t\t\tIDR 55.000")
	fmt.Println("7. Garlic Chicken Teriyaki Set - Red Hot Chilli\t\tIDR 41000")
	fmt.Println("8. Garlic Chicken Teriyaki\t\t\t\tIDR 31.000")
	fmt.Println("9. Chicken Karage\t\t\t\t\tIDR 32.000")
	fmt.Println("10. Taro Milk 1 Liter\t\t\t\t\tIDR 71.000")
}

func cart(){
}

func history(){
}

func main() {

	for true {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				fmt.Println("Mengakhiri program...")
			}
    	}()

		fmt.Println("Welcome to Yoshinoya Restaurant!")
		fmt.Println("================================")
		fmt.Println("1. Lihat Menu")
		fmt.Println("2. Lihat Keranjang")
		fmt.Println("3. Riwayat Pesanan")
		fmt.Println("4. Keluar")

		var choice int
		fmt.Print("Masukkan pilihan (1-4): ")
		fmt.Scan(&choice)

		switch choice {
		case 1: 
			showMenu()
		case 2:
			cart()
		case 3:
			history()
		case 4: 
			os.Exit(0)
		default:
			panic("Pilihan tidak valid")
		}
	}
	
	// os.Exit(0)
}