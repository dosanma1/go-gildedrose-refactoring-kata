package model

const (
	maxValue = 50
	minValue = 0
)

type Item struct {
	Name            string
	SellIn, Quality int
}

func increaseQuality(item *Item, quantity int) {
	item.Quality += quantity

	if item.Quality >= maxValue {
		item.Quality = maxValue
	}
}

func decreaseQuality(item *Item, quantity int) {
	item.Quality -= quantity

	if item.Quality <= minValue {
		item.Quality = minValue
	}
}

func decreaseSellIn(item *Item) {
	item.SellIn--
}

func isExpired(item *Item) bool {
	return item.SellIn < 0
}
