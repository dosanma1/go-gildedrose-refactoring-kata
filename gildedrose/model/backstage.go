package model

type Backstage struct{}

func (b *Backstage) Update(item *Item) {
	item.decreaseSellIn()

	if item.isExpired() {
		item.Quality = 0
		return
	}

	if item.SellIn < 5 {
		item.increaseQuality(3)
		return
	}

	if item.SellIn < 10 {
		item.increaseQuality(2)
		return
	}

	item.increaseQuality(1)

}
