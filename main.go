package main

import (
	"fmt"
	"os"
	"strconv"
	"weekly/service"
)

func showMenu() {

	menus, err := service.GetMenu()
	if err != nil {
		fmt.Println("Gagal ambil menu:", err)
		return
	}

	fmt.Println("\n===== MENU =====")
	for _, m := range menus {
		fmt.Printf("%-3d %-40s IDR %8s\n",
			m.ID, m.Name, m.Price)
	}

	var pilih int
	fmt.Print("\nPilih ID menu (0 kembali): ")
	fmt.Scan(&pilih)

	if pilih == 0 {
		return
	}

	for _, m := range menus {
		if m.ID == pilih {
			service.AddToCart(m)
			fmt.Println("✅ Masuk cart")
			return
		}
	}

	fmt.Println("Menu tidak ditemukan")
}

func cart() {

	cart := service.GetCart()

	if len(cart) == 0 {
		fmt.Println("Cart kosong")
		return
	}

	total := 0

	fmt.Println("\n===== CART =====")
	for i, item := range cart {

		priceInt, _ := strconv.Atoi(item.Price)
		total += priceInt

		fmt.Printf("%d. %-40s IDR %s\n",
			i+1, item.Name, item.Price)
	}

	fmt.Println("Total:", total)

	fmt.Println("\n1. Checkout")
	fmt.Println("2. Clear Cart")
	fmt.Println("0. Kembali")

	var pilih int
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		checkout()
	case 2:
		service.ClearCart()
		fmt.Println("Cart dikosongkan")
	}
}

func checkout() {

	cart := service.GetCart()

	if len(cart) == 0 {
		fmt.Println("Cart kosong")
		return
	}

	service.AddHistory(cart)
	service.ClearCart()

	fmt.Println("✅ Checkout berhasil")
}

func history() {

	histories := service.GetHistory()

	if len(histories) == 0 {
		fmt.Println("Belum ada history")
		return
	}

	fmt.Println("\n===== HISTORY =====")

	for i, order := range histories {

		fmt.Printf("\nOrder #%d\n", i+1)

		total := 0

		for _, item := range order {

			priceInt, _ := strconv.Atoi(item.Price)
			total += priceInt

			fmt.Printf("- %-35s IDR %s\n",
				item.Name, item.Price)
		}

		fmt.Println("Total:", total)
	}
}

func main() {

	for {

		fmt.Println("\n1. Menu")
		fmt.Println("2. Cart")
		fmt.Println("3. History")
		fmt.Println("4. Exit")

		var pilih int
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			showMenu()
		case 2:
			cart()
		case 3:
			history()
		case 4:
			os.Exit(0)
		}
	}
}
