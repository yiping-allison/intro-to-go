class Cafe:
	def __init__(self):
		self.Name = "Happy Coffee"
		self.Drinks = [
			Drink("Americano", 3),
			Drink("Cappuccino", 5),
			Drink("Latte", 5)
		]
		self.CafeFood = [
			Food("chocolate chip cookies", 5, 78),
			Food("blueberry muffin", 5, 467),
			Food("fruit cup", 7, 124),
		]


class Drink:
	def __init__(self, name: str, price: int):
		self.Name = name 
		self.Price = price

	def getName(self) -> str:
		return self.Name

	def getPrice(self) -> int:
		return self.Price

	
class Food:
	def __init__(self, name: str, price: int, calories: int):
		self.Name = name
		self.Price = price
		self.Calories = calories

	def getName(self) -> str:
		return self.Name

	def getPrice(self) -> int:
		return self.Price

	def getCalories(self) -> int:
		return self.Calories
