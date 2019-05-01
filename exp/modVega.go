package exp

type VizVegaChart struct {
	Width   int      `json:"width"`
	Height  int      `json:"height"`
	Padding Padding  `json:"padding"`
	Data    []Data   `json:"data"`
	Scales  []Scales `json:"scales"`
	Axes    []Axes   `json:"axes"`
	Marks   []Marks  `json:"marks"`
}

type Padding struct {
	Top    int `json:"top"`
	Left   int `json:"left"`
	Bottom int `json:"bottom"`
	Right  int `json:"right"`
}
type Data struct {
	Name   string   `json:"name"`
	Values []Values `json:"values"`
}
type Values struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Scales struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Range  string `json:"range"`
	Domain Domain `json:"domain"`
	Nice   bool   `json:"nice,omitempty"`
}

type Domain struct {
	Data  string `json:"data"`
	Field string `json:"field"`
}
type Axes struct {
	Type  string `json:"type"`
	Scale string `json:"scale"`
}

type X struct {
	Scale string `json:"scale"`
	Field string `json:"field"`
}

type Width struct {
	Scale  string `json:"scale"`
	Band   bool   `json:"band"`
	Offset int    `json:"offset"`
}
type Y struct {
	Scale string `json:"scale"`
	Field string `json:"field"`
}
type Y2 struct {
	Scale string `json:"scale"`
	Value int    `json:"value"`
}
type Enter struct {
	X     X     `json:"x"`
	Width Width `json:"width"`
	Y     Y     `json:"y"`
	Y2    Y2    `json:"y2"`
}
type Fill struct {
	Value string `json:"value"`
}
type Update struct {
	Fill Fill `json:"fill"`
}
type Hover struct {
	Fill struct {
		Value string `json:"value"`
	} `json:"fill"`
}
type Properties struct {
	Enter  Enter  `json:"enter"`
	Update Update `json:"update"`
	Hover  Hover  `json:"hover"`
}
type From struct {
	Data string `json:"data"`
}
type Marks struct {
	Type       string     `json:"type"`
	From       From       `json:"from"`
	Properties Properties `json:"properties"`
}
