package exp

type Node struct {
	Coin     string `json:"coin"`
	IP       string `json:"ip"`
	Port     int64  `json:"port"`
	LastSeen string `json:"lastseen"`
}

type NodeMap struct {
	Coin          string  `json:"coin" form:"coin"`
	IP            string  `json:"ip" form:"ip"`
	Country_short string  `"country_short" form:"country_short"`
	Country_long  string  `json:"country_long" form:"country_long"`
	Region        string  `json:"region" form:"region"`
	City          string  `json:"city" form:"city"`
	Latitude      float32 `json:"latitude" form:"latitude"`
	Longitude     float32 `json:"longitude" form:"longitude"`
	Zipcode       string  `json:"zipcode" form:"zipcode"`
	Timezone      string  `json:"timezone" form:"timezone"`
}
