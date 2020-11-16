package person

// Person represents a sample person who can buy things
// at our cafe
type Person struct {
	Name  string
	Money int
}

// New creates a new customer
func New(name string, money int) *Person {
	return &Person{
		Name:  name,
		Money: money,
	}
}

// Buy returns true if someone can buy multiple items.
// It also removes the total price from the person's wallet if Buy
// can be successful
func (p *Person) Buy(items ...interface{ GetPrice() int }) bool {
	total := 0
	for _, item := range items {
		total += item.GetPrice()
	}
	if total > p.Money {
		return false
	}
	p.Money = buy(total, p.Money)
	return true
}

// this function is invisible to code outside of the package
// unexported function - encapsulation
func buy(price, wallet int) int {
	return wallet - price
}
