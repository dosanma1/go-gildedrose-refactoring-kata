package model

type StandardItem struct{}

func (s *StandardItem) Update(item *Item) {
	item.decreaseSellIn()

	if item.isExpired() {
		item.decreaseQuality(2)
	} else {
		item.decreaseQuality(1)
	}
}
