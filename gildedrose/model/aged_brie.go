package model

type AgedBrie struct {
	*Item
}

func (a *AgedBrie) Update() {
	decreaseSellIn(a.Item)

	if isExpired(a.Item) {
		increaseQuality(a.Item, 2)
		return
	}

	increaseQuality(a.Item, 1)

}
