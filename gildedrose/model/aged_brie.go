package model

type AgedBrie struct{}

func (a *AgedBrie) Update(item *Item) {
	item.decreaseSellIn()

	if item.isExpired() {
		item.increaseQuality(2)
	} else {
		item.increaseQuality(1)
	}
}
