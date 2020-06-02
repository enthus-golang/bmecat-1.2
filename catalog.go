package bmecat

import "encoding/xml"

type TNewCatalog struct {
	XMLName xml.Name `xml:"T_NEW_CATALOG"`

	CatalogGroupSystem *CatalogGroupSystem `xml:"CATALOG_GROUP_SYSTEM"`
	Articles           []Article           `xml:"ARTICLE"`
}

type Article struct {
	XMLName xml.Name `xml:"ARTICLE"`

	SupplierArticleID   string              `xml:"SUPPLIER_AID" validate:"required,max=32"`
	ArticleDetails      ArticleDetails      `xml:"ARTICLE_DETAILS"`
	ArticleOrderDetails ArticleOrderDetails `xml:"ARTICLE_ORDER_DETAILS"`
	ArticlePriceDetails ArticlePriceDetails `xml:"ARTICLE_PRICE_DETAILS"`
	MimeInfo            *MIMEInfo           `xml:"MIME_INFO"`
}

type ArticleDetails struct {
	XMLName xml.Name `xml:"ARTICLE_DETAILS"`

	DescriptionShort             string `xml:"DESCRIPTION_SHORT" validate:"required,max=80"`
	DescriptionLong              string `xml:"DESCRIPTION_LONG,omitempty" validate:"max=64000"`
	EAN                          string `xml:"EAN,omitempty" validate:"max=14"`
	SupplierAlternativeArticleID string `xml:"SUPPLIER_ALT_AID,omitempty" validate:"max=50"`
	//BuyerArticleID               string `xml:"BUYER_AID,omitempty" validate:"max=50"`
	ManufacturerArticleID       string `xml:"MANUFACTURER_AID,omitempty" validate:"max=50"`
	ManufacturerName            string `xml:"MANUFACTURER_NAME,omitempty" validate:"max=50"`
	ManufacturerTypeDescription string `xml:"MANUFACTURER_TYPE_DESCR,omitempty" validate:"max=50"`
	ERPGroupBuyer               string `xml:"ERP_GROUP_BUYER,omitempty" validate:"max=10"`
	ERPGroupSupplier            string `xml:"ERP_GROUP_SUPPLIER,omitempty" validate:"max=10"`
	DeliveryTime                int    `xml:"DELIVERY_TIME,omitempty"`
	//SpecialTreatmentClass        string   `xml:"SPECIAL_TREATMENT_CLASS,omitempty" validate:"max=20"`
	Keyword      string `xml:"KEYWORD,omitempty" validate:"max=50"`
	Remarks      string `xml:"REMARKS,omitempty" validate:"max=64000"`
	Segment      string `xml:"SEGMENT,omitempty" validate:"max=100"`
	ArticleOrder int    `xml:"ARTICLE_ORDER,omitempty"`
	//ArticleStatus                []string `xml:"ARTICLE_STATUS,omitempty" validate:"max=250"`
}

type ArticleOrderDetails struct {
	XMLName xml.Name `xml:"ARTICLE_ORDER_DETAILS"`

	OrderUnit                     string  `xml:"ORDER_UNIT" validate:"max=3"`
	ContentUnit                   string  `xml:"CONTENT_UNIT,omitempty" validate:"max=3"`
	NumberContentUnitPerOrderUnit float64 `xml:"NO_CU_PER_OU,omitempty"`
	QuantityMinimum               int     `xml:"QUANTITY_MIN,omitempty"`
	QuantityInterval              int     `xml:"QUANTITY_INTERVAL,omitempty"`
}

type ArticlePriceDetails struct {
	XMLName xml.Name `xml:"ARTICLE_PRICE_DETAILS"`

	Dates         []Datetime     `xml:"DATETIME" validate:"max=2"` // for valid_start_date & valid_end_date
	DailyPrice    bool           `xml:"DAILY_PRICE,omitempty"`
	ArticlePrices []ArticlePrice `xml:"ARTICLE_PRICE" validate:"required,min=1"`
}

type ArticlePrice struct {
	XMLName   xml.Name `xml:"ARTICLE_PRICE"`
	PriceType string   `xml:"PRICE_TYPE,attr" validate:"required,max=20"`

	PriceAmount   float64 `xml:"PRICE_AMOUNT" validate:"required"`
	PriceCurrency string  `xml:"PRICE_CURRENCY,omitempty" validate:"max=3"`
	Tax           float64 `xml:"TAX,omitempty"`
	PriceFactor   float64 `xml:"PRICE_FACTOR,omitempty"`
	LowerBound    float64 `xml:"LOWER_BOUND,omitempty"`
	Territory     string  `xml:"TERRITORY,omitempty"`
}

type MIMEInfo struct {
	XMLName xml.Name `xml:"MIME_INFO"`

	MIME []MIME `xml:"MIME"`
}

type MIME struct {
	XMLName xml.Name `xml:"MIME"`

	Type            string `xml:"MIME_TYPE,omitempty" validate:"max=20"`
	Source          string `xml:"MIME_SOURCE" validate:"max=250"`
	Description     string `xml:"MIME_DESCR,omitempty" validate:"max=250"`
	AlternativeText string `xml:"MIME_ALT,omitempty" validate:"max=50"`
	Purpose         string `xml:"MIME_PURPOSE,omitempty" validate:"max=20"`
	Order           int    `xml:"MIME_ORDER,omitempty"`
}

type CatalogGroupSystem struct {
	XMLName xml.Name `xml:"CATALOG_GROUP_SYSTEM"`

	GroupSystemID          string             `xml:"GROUP_SYSTEM_ID,omitempty" validate:"max=50"`
	GroupSystemName        string             `xml:"GROUP_SYSTEM_NAME,omitempty" validate:"max=50"`
	CatalogStructure       []CatalogStructure `xml:"CATALOG_STRUCTURE"`
	GroupSystemDescription string             `xml:"GROUP_SYSTEM_DESCRIPTION,omitempty" validate:"max=250"`
}

type CatalogStructure struct {
	XMLName xml.Name `xml:"CATALOG_STRUCTURE"`
	Type    string   `xml:"type,attr" validate:"len=4"`

	GroupID          string    `xml:"GROUP_ID" validate:"max=50"`
	GroupName        string    `xml:"GROUP_NAME" validate:"max=50"`
	GroupDescription string    `xml:"GROUP_DESCRIPTION,omitempty" validate:"max=250"`
	ParentID         string    `xml:"PARENT_ID" validate:"max=50"`
	GroupOrder       int       `xml:"GROUP_ORDER,omitempty"`
	MIMEInfo         *MIMEInfo `xml:"MIME_INFO,omitempty"`
	Keyword          string    `xml:"KEYWORD" validate:"max=50"`
}
