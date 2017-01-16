package alipay

type (
	Pay struct {
		Method           string
		BizContent       *Biz
	}

	Biz struct {
		OutTradeNo       string        `json:"out_trade_no"`
		TimeoutExpress   string        `json:"timeout_express"`
		TotalAmount      string        `json:"total_amount"`
		SellId           string        `json:"seller_id"`                      //todo 0
		Subject          string        `json:"subject"`
		Body             string        `json:"body"`
		ProductCode      string        `json:"product_code"`
	}
)
