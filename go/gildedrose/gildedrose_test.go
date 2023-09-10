package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Foo(t *testing.T) {
	var items = []*gildedrose.Item{
		{gildedrose.AGED_BRIE, 0, 0},
	}

	gildedrose.UpdateItems(items)

	if items[0].Name != gildedrose.AGED_BRIE {
		t.Errorf("Name: Expected %s but got %s ", gildedrose.AGED_BRIE, items[0].Name)
	}
}

func TestAgedBrieSellin0Quality0UpdatesToQuality2AndSellinIsDecreasedByOne(t *testing.T) {
	var items = []*gildedrose.Item{
		{gildedrose.AGED_BRIE, 0, 0},
	}

	gildedrose.UpdateItems(items)

	if items[0].Name != gildedrose.AGED_BRIE {
		t.Errorf("Name: Expected %s but got %s ", gildedrose.AGED_BRIE, items[0].Name)
	}

	if items[0].Quality != 2 {
		t.Errorf("Name: Expected %d but got %d for Quality", 2, items[0].Quality)

	}

	if items[0].SellIn != -1 {
		t.Errorf("Name: Expected %d but got %d for Sellin ", -1, items[0].SellIn)

	}
}

func TestAgedBrieSellin10Quality50UpdatesToQuality50AndSellinIsDecreasedByOne(t *testing.T) {
	var items = []*gildedrose.Item{
		{gildedrose.AGED_BRIE, 10, 50},
	}

	gildedrose.UpdateItems(items)

	if items[0].Name != gildedrose.AGED_BRIE {
		t.Errorf("Name: Expected %s but got %s ", gildedrose.AGED_BRIE, items[0].Name)
	}

	if items[0].Quality != 50 {
		t.Errorf("Name: Expected %d but got %d for Quality", 50, items[0].Quality)

	}

	if items[0].SellIn != 9 {
		t.Errorf("Name: Expected %d but got %d for Sellin ", 9, items[0].SellIn)

	}
}

func TestAgedBrieSellin100Quality40UpdatesToQuality42AndSellinIsDecreasedByOne(t *testing.T) {
	var items = []*gildedrose.Item{
		{gildedrose.AGED_BRIE, 100, 40},
	}

	gildedrose.UpdateItems(items)

	if items[0].Name != gildedrose.AGED_BRIE {
		t.Errorf("Name: Expected %s but got %s ", gildedrose.BACKSTAGE, items[0].Name)
	}

	if items[0].Quality != 41 {
		t.Errorf("Name: Expected %d but got %d for Quality", 41, items[0].Quality)

	}

	if items[0].SellIn != 99 {
		t.Errorf("Name: Expected %d but got %d for Sellin ", 99, items[0].SellIn)

	}
}

func TestSulfurasSellin100Quality40UpdatesToQuality42AndSellinDoesNotDecrease(t *testing.T) {
	var items = []*gildedrose.Item{
		{gildedrose.SULFURAS, 100, 40},
	}

	gildedrose.UpdateItems(items)

	if items[0].Name != gildedrose.SULFURAS {
		t.Errorf("Name: Expected %s but got %s ", gildedrose.SULFURAS, items[0].Name)
	}

	if items[0].Quality != 40 {
		t.Errorf("Name: Expected %d but got %d for Quality", 40, items[0].Quality)

	}

	if items[0].SellIn != 100 {
		t.Errorf("Name: Expected %d but got %d for Sellin ", 100, items[0].SellIn)

	}
}

func TestSulfurasSellin100Quality51UpdatesToQuality51AndSellinDoesNotDecrease(t *testing.T) {
	var items = []*gildedrose.Item{
		{gildedrose.SULFURAS, 100, 51},
	}

	gildedrose.UpdateItems(items)

	if items[0].Name != gildedrose.SULFURAS {
		t.Errorf("Name: Expected %s but got %s ", gildedrose.SULFURAS, items[0].Name)
	}

	if items[0].Quality != 51 {
		t.Errorf("Name: Expected %d but got %d for Quality", 51, items[0].Quality)

	}

	if items[0].SellIn != 100 {
		t.Errorf("Name: Expected %d but got %d for Sellin ", 100, items[0].SellIn)

	}
}

func TestBackstageSellin100Quality51UpdatesToQuality51AndSellinDecreases(t *testing.T) {
	var items = []*gildedrose.Item{
		{gildedrose.BACKSTAGE, 100, 51},
	}

	gildedrose.UpdateItems(items)

	if items[0].Name != gildedrose.BACKSTAGE {
		t.Errorf("Name: Expected %s but got %s ", gildedrose.BACKSTAGE, items[0].Name)
	}

	if items[0].Quality != 51 {
		t.Errorf("Name: Expected %d but got %d for Quality", 51, items[0].Quality)

	}

	if items[0].SellIn != 99 {
		t.Errorf("Name: Expected %d but got %d for Sellin ", 99, items[0].SellIn)

	}
}

func TestBackstageSellin100Quality40UpdatesToQuality42AndSellinDecreases(t *testing.T) {
	var items = []*gildedrose.Item{
		{gildedrose.BACKSTAGE, 100, 40},
	}

	gildedrose.UpdateItems(items)

	if items[0].Name != gildedrose.BACKSTAGE {
		t.Errorf("Name: Expected %s but got %s ", gildedrose.BACKSTAGE, items[0].Name)
	}

	if items[0].Quality != 41 {
		t.Errorf("Name: Expected %d but got %d for Quality", 41, items[0].Quality)

	}

	if items[0].SellIn != 99 {
		t.Errorf("Name: Expected %d but got %d for Sellin ", 99, items[0].SellIn)

	}
}
