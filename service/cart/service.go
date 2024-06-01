package cart

import (
	"fmt"

	"github.com/pandfun/ecom/types"
)

func getCartItemIDs(items []types.CartItem) ([]int, error) {
	productIDs := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalud quantitiy for the product %d", item.ProductID)
		}

		productIDs[i] = item.ProductID
	}

	return productIDs, nil
}

func (h *Handler) createOrder(ps []types.Product, items []types.CartItem, userID int) (int, float64, error) {

	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[product.ID] = product
	}

	// check if products are in stock
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, err
	}

	// calculate total price
	totalPrice := calculateTotalPrice(items, productMap)

	// update the new product quantity in db
	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity

		err := h.productStore.UpdateProduct(product)
		if err != nil {
			return 0, 0, err
		}
	}

	// create order
	orderID, err := h.orderStore.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "some address",
	})
	if err != nil {
		return 0, 0, err
	}

	// create order item
	for _, item := range items {
		h.orderStore.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
	}

	return orderID, totalPrice, nil
}

func checkIfCartIsInStock(cartItems []types.CartItem, products map[int]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not available, please refresh your cart", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not avaible in the quantitiy requested", product.Name)
		}
	}

	return nil
}

func calculateTotalPrice(cartItems []types.CartItem, products map[int]types.Product) float64 {
	var total float64

	for _, item := range cartItems {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}

	return total
}
