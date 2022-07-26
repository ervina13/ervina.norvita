package repository

type TransactionRepository struct {
	cartItemRepository CartItemRepository
	salesRepository    SalesRepository
}

func NewTransactionRepository(cartItemRepository CartItemRepository, salesRepository SalesRepository) TransactionRepository {
	return TransactionRepository{cartItemRepository, salesRepository}
}

func (u *TransactionRepository) Pay(cartItems []CartItem, amount int) (int, error) {
	//return 0, nil // TODO: replace this
	totalPrice, err := u.cartItemRepository.TotalPrice()

	if err != nil {
		return 0, err
	}
	return amount - totalPrice, nil
}
