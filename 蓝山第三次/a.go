package main

import "fmt"

// RAM 内存结构体
type RAM struct {
	name  string
	price float64
	stock int
}

type InventoryManager interface {
	CheckStock() int
	UpdateStock(int)
	PrintInventory()
}

// ElectronicProduct 电子产品结构体
type ElectronicProduct struct {
	RAM
	brand string
	model string
}

type ElectronicInventoryManager interface {
	InventoryManager
	PrintBrandModel()
}

func (r *RAM) CheckStock() int {
	return r.stock
}
func (r *RAM) UpdateStock(quantity int) {
	r.stock += quantity
}

func (r *RAM) PrintInventory() {
	fmt.Printf("name:%s\nprice:%.2f\nstock:%d\n", r.name, r.price, r.stock)
}

func (e *ElectronicProduct) PrintBrandModel() {
	fmt.Printf("Product: %s\nBrand: %s\nModel: %s\n", e.name, e.brand, e.model)
}
func main() {
	GuangWei := RAM{
		name:  "GuangWei",
		price: 256.5,
		stock: 99,
	}

	HuaWei := ElectronicProduct{
		RAM: RAM{
			name:  "HuaWei Mate 60",
			price: 6999.99,
			stock: 999,
		},
		brand: "Mate 60",
		model: "1",
	}

	fmt.Println("GuangWei Inventory:")
	GuangWei.PrintInventory()
	fmt.Println("Previous Stock:", GuangWei.CheckStock())
	GuangWei.UpdateStock(1)
	fmt.Printf("Updated Stock: %d", GuangWei.CheckStock())

	fmt.Println("\nElectronic Product Inventory:")
	HuaWei.PrintInventory()
	fmt.Println("Previous Stock:", HuaWei.CheckStock())
	HuaWei.UpdateStock(-99)
	fmt.Printf("Updated Stock: %d", HuaWei.CheckStock())
	fmt.Println("\nBrand and Model Information:")
	HuaWei.PrintBrandModel()
}
