package responses

type HeaderPartner struct {
	D struct {
	SalesOrder                  string `json:"SalesOrder"`
	PartnerFunction             string `json:"PartnerFunction"`
	Customer                    string `json:"Customer"`
	Supplier                    string `json:"Supplier"`
	} `json:"d"`
}
