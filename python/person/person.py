class Person:
	def __init__(self, name: str, money: int):
		self.Name = name
		self.Money = money

	def Buy(self, *args) -> bool:
		total = 0
		for val in args:
			total += val
		if total > self.Money:
			return False
		self.Money -= total
		return True

	def getMoney(self) -> int:
		return self.Money

	def getName(self) -> str:
		return self.Name
