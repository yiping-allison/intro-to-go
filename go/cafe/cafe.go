package cafe

// Cafe represents the metadata of a cafe
type Cafe struct {
	Name     string
	Drinks   []Drink
	CafeFood []Food
}

// Drink represents a drink you can get at the cafe
type Drink struct {
	Name  string
	Price int
}

func (d Drink) GetPrice() int {
	return d.Price
}

// Food represent a food item you can get at the cafe
type Food struct {
	Name     string
	Price    int
	Calories int
}

func (f Food) GetPrice() int {
	return f.Price
}

func (f Food) GetCalories() int {
	return f.Calories
}

// Init initializes our sample cafe
func (c *Cafe) Init() {
	c.Name = "Happy Coffee"
	c.Drinks = []Drink{
		{"Americano", 3},
		{"Cappuccino", 5},
		{"Latte", 5},
	}
	c.CafeFood = []Food{
		{"chocolate chip cookies", 5, 78},
		{"blueberry muffin", 5, 467},
		{"fruit cup", 7, 124},
	}
}
