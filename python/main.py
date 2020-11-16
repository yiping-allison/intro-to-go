from cafe.cafe import Cafe
from person.person import Person


def main():
	print("Welcome to Python's main function!")

	print("Starting our sample cafe...")
	cafe = Cafe()

	print("Fred's Turn")
	person = Person("Fred", 100)
	person.Buy(cafe.Drinks[0].getPrice(), cafe.Drinks[2].getPrice(), cafe.CafeFood[2].getPrice())
	print(person.getMoney())

	print("Jess' Turn")
	person = Person("Jess", 0)
	person.Buy(cafe.CafeFood[0].getPrice())
	print(person.getMoney())


if __name__ == "__main__":
	main()
