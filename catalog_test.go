package bmecat_test

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/mclgmbh/golang-pkg/bmecat"
)

func TestCatalog(t *testing.T) {
	catalog := bmecat.TNewCatalog{
		Articles: []bmecat.Article{
			{
				SupplierArticleID: "123123",
				ArticleDetails: bmecat.ArticleDetails{
					DescriptionShort: "some text",
					DescriptionLong:  "some longer text",
					EAN:              "NOPENOPENOPE",
				},
				ArticleOrderDetails: bmecat.ArticleOrderDetails{
					OrderUnit:   "PA",
					ContentUnit: "C62",
				},
				ArticlePriceDetails: bmecat.ArticlePriceDetails{
					ArticlePrices: []bmecat.ArticlePrice{
						{
							PriceType:     "net_customer",
							PriceAmount:   200.01,
							PriceCurrency: "USD",
							Tax:           0.21,
							LowerBound:    1,
						},
						{
							PriceType:     "net_customer_exp",
							PriceAmount:   199.99,
							PriceCurrency: "EUR",
							Tax:           0.19,
							LowerBound:    1,
						},
					},
				},
				MimeInfo: &bmecat.MIMEInfo{
					MIME: []bmecat.MIME{
						{
							Type:            "image/jpg",
							Source:          "asd.jpg",
							Description:     "sssssssssssss",
							AlternativeText: "picture of asd",
							Purpose:         "normal",
						},
						{
							Type:            "image/jpg",
							Source:          "asd2.jpg",
							Description:     "sssssssssssss",
							AlternativeText: "picture of asd2",
							Purpose:         "detail",
						},
					},
				},
			},
		},
	}

	b, err := xml.MarshalIndent(catalog, "", "  ")
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	assert.Equal(t, `<T_NEW_CATALOG>
  <ARTICLE>
    <SUPPLIER_AID>123123</SUPPLIER_AID>
    <ARTICLE_DETAILS>
      <DESCRIPTION_SHORT>some text</DESCRIPTION_SHORT>
      <DESCRIPTION_LONG>some longer text</DESCRIPTION_LONG>
      <EAN>NOPENOPENOPE</EAN>
    </ARTICLE_DETAILS>
    <ARTICLE_ORDER_DETAILS>
      <ORDER_UNIT>PA</ORDER_UNIT>
      <CONTENT_UNIT>C62</CONTENT_UNIT>
    </ARTICLE_ORDER_DETAILS>
    <ARTICLE_PRICE_DETAILS>
      <ARTICLE_PRICE PRICE_TYPE="net_customer">
        <PRICE_AMOUNT>200.01</PRICE_AMOUNT>
        <PRICE_CURRENCY>USD</PRICE_CURRENCY>
        <TAX>0.21</TAX>
        <LOWER_BOUND>1</LOWER_BOUND>
      </ARTICLE_PRICE>
      <ARTICLE_PRICE PRICE_TYPE="net_customer_exp">
        <PRICE_AMOUNT>199.99</PRICE_AMOUNT>
        <PRICE_CURRENCY>EUR</PRICE_CURRENCY>
        <TAX>0.19</TAX>
        <LOWER_BOUND>1</LOWER_BOUND>
      </ARTICLE_PRICE>
    </ARTICLE_PRICE_DETAILS>
    <MIME_INFO>
      <MIME>
        <MIME_TYPE>image/jpg</MIME_TYPE>
        <MIME_SOURCE>asd.jpg</MIME_SOURCE>
        <MIME_DESCR>sssssssssssss</MIME_DESCR>
        <MIME_ALT>picture of asd</MIME_ALT>
        <MIME_PURPOSE>normal</MIME_PURPOSE>
      </MIME>
      <MIME>
        <MIME_TYPE>image/jpg</MIME_TYPE>
        <MIME_SOURCE>asd2.jpg</MIME_SOURCE>
        <MIME_DESCR>sssssssssssss</MIME_DESCR>
        <MIME_ALT>picture of asd2</MIME_ALT>
        <MIME_PURPOSE>detail</MIME_PURPOSE>
      </MIME>
    </MIME_INFO>
  </ARTICLE>
</T_NEW_CATALOG>`, string(b))
}
