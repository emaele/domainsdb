package domainsdb

// Result contains all the information about the query requested
type Result struct {
	Total   int      `json:"total"`
	Time    int      `json:"time"`
	Domains []Domain `json:"domains"`
}

// Domain contains the information about one domain from the search
type Domain struct {
	Domain     string      `json:"domain"`
	Suffix     string      `json:"suffix,omitempty"`
	ExpiryDate interface{} `json:"expiry_date"`
	CreateDate interface{} `json:"create_date"`
	UpdateDate string      `json:"update_date"`
	IsDead     bool        `json:"isDead"`
	Resolvable bool        `json:"resolvable,omitempty"`
	A          []string    `json:"A,omitempty"`
	TXT        []string    `json:"TXT,omitempty"`
	Country    string      `json:"country,omitempty"`
	NS         []string    `json:"NS,omitempty"`
	CNAME      interface{} `json:"CNAME,omitempty"`
	MX         []struct {
		Exchange string `json:"exchange"`
		Priority int    `json:"priority"`
	} `json:"MX,omitempty"`
}
