package models

type ClientReq struct {
	Event string   `json:"event"`
	Data  []string `json:"data"`
}

type DatabaseResponse struct {
	Orders   string
	Sales    string
	Products string
}

func (d *DatabaseResponse) ToString() string {
	return d.Orders + " " + d.Sales + " " + d.Products
}
