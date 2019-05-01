package exp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.parallelcoin.io/marcetin/explorer/jdb"
	"github.com/gorilla/mux"
)

func GetChartData(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	var chartData []Addr
	jdb.JDB.Read("data/"+v["coin"]+"/"+v["cat"], v["data"], &chartData)
	var chartValues []Values
	for _, c := range chartData {
		chartValues = append(chartValues, Values{
			Y: int(c.Value),
		})
	}
	VVC := VizVegaChart{
		Width:  640,
		Height: 480,
		Padding: Padding{
			Top:    24,
			Left:   48,
			Bottom: 24,
			Right:  0,
		},
		Data: []Data{
			Data{
				Name:   "table",
				Values: chartValues,
			},
		},
		Scales: []Scales{
			Scales{
				Name:  "x",
				Type:  "ordinal",
				Range: "width",
				Domain: Domain{
					Data:  "table",
					Field: "x",
				},
			}, Scales{
				Name:  "y",
				Type:  "linear",
				Range: "height",
				Domain: Domain{
					Data:  "table",
					Field: "y",
				},
				Nice: true,
			},
		},
		Axes: []Axes{
			Axes{
				Type:  "x",
				Scale: "x",
			}, Axes{
				Type:  "y",
				Scale: "y",
			},
		},
		Marks: []Marks{
			Marks{
				Type: "rect",
				From: From{
					Data: "table",
				},
				Properties: Properties{
					Enter: Enter{
						X: X{
							Scale: "x",
							Field: "x",
						},
						Width: Width{
							Scale:  "x",
							Band:   true,
							Offset: -1,
						},
						Y: Y{
							Scale: "y",
							Field: "y",
						},
						Y2: Y2{
							Scale: "y",
							Value: 0,
						},
					},
					Update: Update{
						Fill: Fill{
							Value: "steelblue",
						},
					},
					Hover: Hover{
						Fill: Fill{
							Value: "red",
						},
					},
				},
			},
		},
	}
	out, err := json.Marshal(VVC)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
