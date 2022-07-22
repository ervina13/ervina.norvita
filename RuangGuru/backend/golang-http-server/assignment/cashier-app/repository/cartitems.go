package repository

import (
	"strconv"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/db"
)

type CartItemRepository struct {
	db db.DB
}

const cartItemDbName = "cart_items"

var cartItemColums = []string{"category", "product_name", "price", "quantity"}

func NewCartItemRepository(db db.DB) CartItemRepository {
	return CartItemRepository{db}
}

func (u *CartItemRepository) LoadOrCreate() ([]CartItem, error) {
	records, err := u.db.Load(cartItemDbName)
	if err != nil {
		records = [][]string{cartItemColums}
		if err := u.db.Save(cartItemDbName, records); err != nil {
			return nil, err
		}
	}

	result := make([]CartItem, 0)
	for i := 1; i < len(records); i++ {
		price, err := strconv.Atoi(records[i][2])
		if err != nil {
			return nil, err
		}

		qty, err := strconv.Atoi(records[i][3])
		if err != nil {
			return nil, err
		}

		cartItem := CartItem{
			Category:    records[i][0],
			ProductName: records[i][1],
			Price:       price,
			Quantity:    qty,
		}
		result = append(result, cartItem)
	}

	return result, nil
}

func (u *CartItemRepository) Save(cartItems []CartItem) error {
	records := [][]string{cartItemColums}
	for i := 0; i < len(cartItems); i++ {
		records = append(records, []string{
			cartItems[i].Category,
			cartItems[i].ProductName,
			strconv.Itoa(cartItems[i].Price),
			strconv.Itoa(cartItems[i].Quantity),
		})
	}
	return u.db.Save(cartItemDbName, records)
}

func (u *CartItemRepository) SelectAll() ([]CartItem, error) {
	result, _ := u.LoadOrCreate()
	return result, nil
	//return u.LoadOrCreate() // TODO: replace this
}

func (u *CartItemRepository) Add(product Product) error {
	result, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

	newRecords := [][]string{cartItemColums}
	var isAddQty bool

	for i := 0; i < len(result); i++ {
		if result[i].ProductName == product.ProductName {
			result[i].Quantity += 1
			isAddQty = true
		}

		newRecords = append(newRecords, []string{
			result[i].Category,
			result[i].ProductName,
			strconv.Itoa(result[i].Price),
			strconv.Itoa(result[i].Quantity),
		})
	}

	if isAddQty == false {
		newRecords = append(newRecords, []string{
			product.Category,
			product.ProductName,
			strconv.Itoa(product.Price),
			"1",
		})
	}

	err = u.db.Save(cartItemDbName, newRecords)
	if err != nil {
		return err
	}
	return nil

}

func (u *CartItemRepository) ResetCartItems() error {
	records := [][]string{cartItemColums}
	err := u.db.Save(cartItemDbName, records)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *CartItemRepository) TotalPrice() (int, error) {
	//load cart items
	result, err := u.LoadOrCreate()
	if err != nil {
		return 0, err
	}

	//loop cartitems
	var total int
	for i := 0; i < len(result); i++ {
		total = total + result[i].Price*result[i].Quantity
	}
	return total, nil // TODO: replace this
}

func searchCartItem(cartItems []CartItem, product Product) (int, bool) {
	for i, item := range cartItems {
		if item.Category == product.Category && item.ProductName == product.ProductName && item.Price == product.Price {
			return i, true
		}
	}
	return -1, false
}
