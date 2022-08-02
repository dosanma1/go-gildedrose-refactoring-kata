package gildedrose_test

import (
	"testing"

	"github.com/dosanma1/go-gildedrose-refactoring-kata/gildedrose"
	"github.com/dosanma1/go-gildedrose-refactoring-kata/gildedrose/model"
	"github.com/stretchr/testify/assert"
)

func Test_StandardItemQualityDecreasesNormalBeforExpiry(t *testing.T) {
	var items = []*model.Item{
		{Name: "Elixir of the Mongoose", SellIn: 10, Quality: 20},
	}

	for i := 0; i < 5; i++ {
		gildedrose.UpdateQuality(items)
	}

	assert.Equal(t, 5, items[0].SellIn)
	assert.Equal(t, 15, items[0].Quality)
}

func Test_StandardItemQualityDecreasesFasterAfterExpiry(t *testing.T) {
	var items = []*model.Item{
		{Name: "Elixir of the Mongoose", SellIn: 1, Quality: 10},
	}

	for i := 0; i < 3; i++ {
		gildedrose.UpdateQuality(items)
	}

	assert.Equal(t, -2, items[0].SellIn)
	assert.Equal(t, 5, items[0].Quality)
}

func Test_ItemQualityDoesNotGoBelowZero(t *testing.T) {
	var items = []*model.Item{
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 1, Quality: 5},
		{Name: "Elixir of the Mongoose", SellIn: 1, Quality: 5},
	}

	for i := 0; i < 3; i++ {
		gildedrose.UpdateQuality(items)
	}

	for _, item := range items {
		assert.Equal(t, 0, item.Quality)
	}
}

func Test_AgedBrieQualityIncreaseNormalBeforeExpiry(t *testing.T) {
	var items = []*model.Item{
		{Name: "Aged Brie", SellIn: 10, Quality: 10},
	}

	for i := 0; i < 5; i++ {
		gildedrose.UpdateQuality(items)
	}

	assert.Equal(t, 5, items[0].SellIn)
	assert.Equal(t, 15, items[0].Quality)
}

func Test_AgedBrieQualityIncreaseFasterAfterExpiry(t *testing.T) {
	var items = []*model.Item{
		{Name: "Aged Brie", SellIn: 1, Quality: 5},
	}

	for i := 0; i < 3; i++ {
		gildedrose.UpdateQuality(items)
	}

	assert.Equal(t, -2, items[0].SellIn)
	assert.Equal(t, 10, items[0].Quality)
}

func Test_ItemQualityDoesNotGoAboveFifty(t *testing.T) {
	var items = []*model.Item{
		{Name: "Aged Brie", SellIn: 5, Quality: 45},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 49},
	}

	for i := 0; i < 10; i++ {
		gildedrose.UpdateQuality(items)
	}

	for _, item := range items {
		assert.Equal(t, 50, item.Quality)
	}
}

func Test_SulfurasDoesNotAgeOrChangeInQuality(t *testing.T) {
	var items = []*model.Item{
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 5, Quality: 80},
	}

	for i := 0; i < 10; i++ {
		gildedrose.UpdateQuality(items)
	}

	assert.Equal(t, 5, items[0].SellIn)
	assert.Equal(t, 80, items[0].Quality)
}

func Test_BackstagePassQualityIncreaseNormalBeforeTenDays(t *testing.T) {
	var items = []*model.Item{
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20},
	}

	for i := 0; i < 5; i++ {
		gildedrose.UpdateQuality(items)
	}

	assert.Equal(t, 10, items[0].SellIn)
	assert.Equal(t, 25, items[0].Quality)
}

func Test_BackstagePassQualityIncreaseDoubleBetweenTenAndFiveDays(t *testing.T) {
	var items = []*model.Item{
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20},
	}

	for i := 0; i < 10; i++ {
		gildedrose.UpdateQuality(items)
	}

	assert.Equal(t, 5, items[0].SellIn)
	assert.Equal(t, 35, items[0].Quality)
}

func Test_BackstagePassQualityIncreaseTripeBetweenFiveDaysAndExpiry(t *testing.T) {
	var items = []*model.Item{
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 20},
	}

	for i := 0; i < 10; i++ {
		gildedrose.UpdateQuality(items)
	}

	assert.Equal(t, 0, items[0].SellIn)
	assert.Equal(t, 45, items[0].Quality)
}

func Test_BackstagePassQualityIsZeroAfterExpiry(t *testing.T) {
	var items = []*model.Item{
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 20},
	}

	for i := 0; i < 11; i++ {
		gildedrose.UpdateQuality(items)
	}

	assert.Equal(t, -1, items[0].SellIn)
	assert.Equal(t, 0, items[0].Quality)
}

func Test_ConjuredItemQualityDecreasesAtDoubleBeforeExpiry(t *testing.T) {
	var items = []*model.Item{
		{Name: "Conjured Mana Cake", SellIn: 10, Quality: 20},
	}

	for i := 0; i < 5; i++ {
		gildedrose.UpdateQuality(items)
	}

	assert.Equal(t, 5, items[0].SellIn)
	assert.Equal(t, 10, items[0].Quality)
}
