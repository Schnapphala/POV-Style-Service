package DTO

type Style struct {
	StyleId int `json:"styleId"`
	StyleNumber string `json:"styleNumber"`
	BrandName string `json:"brandName"`
	StyleState int8 `json:"styleState"`
	CanonicalStyleLink string `json:"canonicalStyleLink"`
	StyleName string `json:"styleName"`
	SizeChartCode string `json:"sizeChartCode"` /*wird benötigt ???*/
	FitName string `json:"fitName"`
	IsLeather int8 `json:"isLeather"`
	IsNewToday int8 `json:"isNewToday"`
	IsSpecialAward int8 `json:"isSpecialAward"`
	ShowBonnesAffaireStamp bool `json:"showBonnesAffaireStamp"`
	IsNewSeason int8 `json:"isNewSeason"`
	IsKeyStyle int8 `json:"isKeyStyle"`
	IsBuyThreePayTwo int8 `json:"isBuyThreePayTwo"`
	TwoFit int8 `json:"twoFit"`
	ClothingThermalLevel int8 `json:"clothingThermalLevel"`
	Colors []Color `json:"colors"`
}

type Color struct {
	ColorNumber string `json:"colorNumber"`
	ColorName string `json:"colorName"`
	BaseColor string `json:"baseColor"`
	IsLightColor bool `json:"isLightColor"`
	IsOneColor bool `json:"isOneColor"`
	ColorPos int `json:"colorPos"`
	MarkdownPromo MarkdownPromo `json:"markdownPromo"`
	RetailIntervalToDisplay string `json:"retailIntervalToDisplay"`//???
	ColorThumb string `json:"colorThumb"`
	SpvLink string `json:"spvLink"`
	ReductionPercentage int `json:"reductionPercentage"`
	LowestPriceBasic string `json:"lowestPriceBasic"`
	LowestPriceNew string `json:"lowestPriceNew"`
	BiggerPriceExists bool `json:"biggerPriceExists"`
	PriceConfigDisplay PriceConfigDisplay `json:"priceConfigDisplay"`
	Sizes []Size `json:"sizes"`
}

type MarkdownPromo struct {

}

type Image struct {
	MainImage ImagePath `json:"mainImage"`
	SecondImage ImagePath `json:"secondImage"`
}

type ImagePath struct {
	Path string `json:"path"`
}

type PriceConfigDisplay struct {
	BasicPriceDisply PriceDisplay
	RetailPriceDisplay PriceDisplay
}

type PriceDisplay struct {
	Show bool `json:"show"`
	Colorreduced bool `json:"colorreduced"`
	Strikethrough bool `json:"strikethrough"`
}

type Size struct {
	SizeValue string `json:"SizeValue"`
	IsAvailable int8 `json:"IsAvailable"`
	PriceBasic string `json:"PriceBasic"`
	PriceNew string `json:"PriceNew"`
}
