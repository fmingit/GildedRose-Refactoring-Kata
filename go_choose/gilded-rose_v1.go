package main

import "fmt"

type Item struct {
	name            string
	sellIn, quality int
}

var (
	quality_lower = 0
	quality_upper = 50
)

func qualityNoLimit(items []*Item) {
	// Sulfuras: 不会降低品质`Quality`
	if items[i].name != "Sulfuras, Hand of Ragnaros" {
		// sellIn: 传奇物品"Sulfuras"永不到期,其他-1
		items[i].sellIn = items[i].sellIn - 1

		// quality:
		// 1) Aged Brie: `Quality`会随着时间推移而提高
		if items[i].name == "Aged Brie" {
			if items[i].quality >= quality_upper {
				return
			}
			items[i].quality = items[i].quality + 1
			return
		}

		// 2) Backstage passes: sellIn<11时
		if items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].sellIn >= 11 {
				if items[i].quality <= quality_lower {
					return
				}
				items[i].sellIn = items[i].sellIn - 1
				return
			}
			if items[i].quality > quality_upper {
				return
			}
			if items[i].sellIn < 0 {
				items[i].quality = items[i].quality - items[i].quality
				return
			}
			if items[i].sellIn < 6 {
				items[i].quality = items[i].quality + 3
				return
			}
			if items[i].sellIn < 11 {
				items[i].quality = items[i].quality + 2
				return
			}
		}

		// 3) 其他: sellIn>0则-1; sellIn<0则-2

	}

}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
	}

}
