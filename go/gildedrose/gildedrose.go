package gildedrose

const SULFURAS = "Sulfuras, Hand of Ragnaros"
const BACKSTAGE = "Backstage passes to a TAFKAL80ETC concert"
const AGED_BRIE = "Aged Brie"

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		if items[i].Name != AGED_BRIE && items[i].Name != BACKSTAGE {
			if items[i].Quality > 0 {
				if items[i].Name != SULFURAS {
					items[i].Quality = items[i].Quality - 1
				}
			}
		} else {
			if items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
				if items[i].Name == BACKSTAGE {
					if items[i].SellIn < 11 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
					if items[i].SellIn < 6 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
				}
			}
		}

		if items[i].Name != SULFURAS {
			items[i].SellIn = items[i].SellIn - 1
		}

		if items[i].SellIn < 0 {
			if items[i].Name != AGED_BRIE {
				if items[i].Name != BACKSTAGE {
					if items[i].Quality > 0 {
						if items[i].Name != SULFURAS {
							items[i].Quality -= 1
						}
					}
				} else {
					items[i].Quality -= items[i].Quality
				}
			} else {
				if items[i].Quality < 50 {
					items[i].Quality += 1
				}
			}
		}
	}
}
