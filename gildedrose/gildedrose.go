package gildedrose

import (
	"strings"

	"github.com/dosanma1/go-gildedrose-refactoring-kata/gildedrose/model"
)

const (
	agedBrie  = "Aged Brie"
	sulfuras  = "Sulfuras, Hand of Ragnaros"
	backstage = "Backstage passes to a TAFKAL80ETC concert"
	conjured  = "Conjured"
)

func UpdateQuality(items []*model.Item) {
	for _, item := range items {
		uItem := getDegradableItem(item)
		uItem.Update()
	}

}

func getDegradableItem(item *model.Item) model.DegradableItem {
	switch item.Name {
	case sulfuras:
		return &model.Sulfuras{Item: item}
	case agedBrie:
		return &model.AgedBrie{Item: item}
	case backstage:
		return &model.Backstage{Item: item}
	default:
		if strings.HasPrefix(item.Name, conjured) {
			return &model.Conjured{Item: item}
		} else {
			return &model.StandardItem{Item: item}
		}
	}
}
