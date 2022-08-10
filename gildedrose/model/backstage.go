package model

type Backstage struct {
	*Item
}

func (b *Backstage) Update() {
	decreaseSellIn(b.Item)

	if isExpired(b.Item) {
		b.Quality = 0
		return
	}

	if b.SellIn < 5 {
		increaseQuality(b.Item, 3)
		return
	}

	if b.SellIn < 10 {
		increaseQuality(b.Item, 2)
		return
	}

	increaseQuality(b.Item, 1)

}
