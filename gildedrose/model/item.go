package model

const (
	MAX_VALUE = 50
	MIN_VALUE = 0
)

type Item struct {
	Name            string
	SellIn, Quality int
}

func (i *Item) increaseQuality(quantity int) {
	i.Quality += quantity

	if i.Quality >= MAX_VALUE {
		i.Quality = MAX_VALUE
	}
}

func (i *Item) decreaseQuality(quantity int) {
	i.Quality -= quantity

	if i.Quality <= MIN_VALUE {
		i.Quality = MIN_VALUE
	}
}

func (i *Item) decreaseSellIn() {
	i.SellIn--
}

func (i *Item) isExpired() bool {
	return i.SellIn < 0
}
