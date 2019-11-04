package main

type Item struct {
	name            string
	sellIn, quality int
}

var (
	quality_lower = 0
	quality_upper = 50
)

func qualityNoLimit(item *Item) (sellIn, quality int) {
	// Sulfuras: 不会降低品质`Quality`,永不到期
	if item.name != "Sulfuras, Hand of Ragnaros" {
		item.sellIn = item.sellIn - 1

		// quality:
		// 1) Aged Brie: `Quality`会随着时间推移而提高
		if item.name == "Aged Brie" {
			if item.sellIn < 0 {
				item.quality = item.quality + 2
				return item.sellIn, item.quality
			}
			item.quality = item.quality + 1
			return item.sellIn, item.quality
		}

		// 2) Backstage passes: sellIn<11时
		if item.name == "Backstage passes to a TAFKAL80ETC concert" {
			if item.sellIn < 0 {
				item.quality = item.quality - item.quality
				return item.sellIn, item.quality
			}
			if item.sellIn < 6 {
				item.quality = item.quality + 3
				return item.sellIn, item.quality
			}
			if item.sellIn < 10 {
				item.quality = item.quality + 2
				return item.sellIn, item.quality
			}
			item.quality = item.quality + 1
			return item.sellIn, item.quality
		}

		// 3) 其他: sellIn>0则-1; sellIn<0则-2
		if item.sellIn < 0 {
			item.quality = item.quality - 2
			return item.sellIn, item.quality
		}
		item.quality = item.quality - 1

	}
	return item.sellIn, item.quality
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		items[i].sellIn, items[i].quality = qualityNoLimit(items[i])
		if items[i].quality < quality_lower {
			items[i].quality = quality_lower
		}
		if items[i].quality > quality_upper && items[i].name != "Sulfuras, Hand of Ragnaros" {
			items[i].quality = quality_upper
		}
	}
}
