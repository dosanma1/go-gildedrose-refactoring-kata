package gildedrose_test

import (
	"testing"

	"github.com/dosanma1/go-gildedrose-refactoring-kata/gildedrose"
	"github.com/dosanma1/go-gildedrose-refactoring-kata/gildedrose/model"
	"github.com/stretchr/testify/assert"
)

func Test_UpdateQuality(t *testing.T) {
	tt := []struct {
		name    string
		items   []*model.Item
		days    int
		sellIn  int
		quality int
	}{
		{
			name:    "standard item quality decreases normal before expiry",
			items:   []*model.Item{{Name: "Elixir of the Mongoose", SellIn: 10, Quality: 20}},
			days:    5,
			sellIn:  5,
			quality: 15,
		},
		{
			name:    "standard item quality decreases faster after expiry",
			items:   []*model.Item{{Name: "Elixir of the Mongoose", SellIn: 1, Quality: 10}},
			days:    3,
			sellIn:  -2,
			quality: 5,
		},
		{
			name: "item quality does not go below zero",
			items: []*model.Item{
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 1, Quality: 5},
				{Name: "Elixir of the Mongoose", SellIn: 1, Quality: 5},
			},
			days:    3,
			sellIn:  -2,
			quality: 0,
		},
		{
			name:    "aged brie quality increases normal before expiry",
			items:   []*model.Item{{Name: "Aged Brie", SellIn: 10, Quality: 10}},
			days:    5,
			sellIn:  5,
			quality: 15,
		},
		{
			name:    "aged brie quality increases faster after expiry",
			items:   []*model.Item{{Name: "Aged Brie", SellIn: 1, Quality: 5}},
			days:    3,
			sellIn:  -2,
			quality: 10,
		},
		{
			name: "item quality does not go above fifty",
			items: []*model.Item{
				{Name: "Aged Brie", SellIn: 10, Quality: 45},
				{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 49},
			},
			days:    10,
			sellIn:  0,
			quality: 50,
		},
		{
			name:    "sulfuras does not age or change in quality",
			items:   []*model.Item{{Name: "Sulfuras, Hand of Ragnaros", SellIn: 5, Quality: 80}},
			days:    10,
			sellIn:  5,
			quality: 80,
		},
		{
			name:    "backstage pass quality increase normal before ten days",
			items:   []*model.Item{{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20}},
			days:    5,
			sellIn:  10,
			quality: 25,
		},
		{
			name:    "backstage pass quality increase double between ten and five days",
			items:   []*model.Item{{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20}},
			days:    10,
			sellIn:  5,
			quality: 35,
		},
		{
			name:    "backstage pass quality increase double between ten and five days",
			items:   []*model.Item{{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 20}},
			days:    10,
			sellIn:  0,
			quality: 45,
		},
		{
			name:    "backstage pass quality increase triple between five days and expiry",
			items:   []*model.Item{{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 20}},
			days:    10,
			sellIn:  0,
			quality: 45,
		},
		{
			name:    "backstage pass quality is zero after expiry",
			items:   []*model.Item{{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 20}},
			days:    11,
			sellIn:  -1,
			quality: 0,
		},
		{
			name:    "conjured item quality decreases at double before expiry",
			items:   []*model.Item{{Name: "Conjured Mana Cake", SellIn: 10, Quality: 20}},
			days:    5,
			sellIn:  5,
			quality: 10,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i < tc.days; i++ {
				gildedrose.UpdateQuality(tc.items)
			}

			for _, item := range tc.items {
				assert.Equal(t, tc.sellIn, item.SellIn)
				assert.Equal(t, tc.quality, item.Quality)
			}
		})
	}
}
