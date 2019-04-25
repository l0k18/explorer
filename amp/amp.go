package amp

import "html/template"

type AMP struct {
	BlockHeight   template.HTML `json:"blockheight"`
	BlockHash     template.HTML `json:"blockhash"`
	Tx            template.HTML `json:"tx"`
	Addr          template.HTML `json:"addr"`
	NextBlockHash template.HTML `json:"nbh"`
	PrevBlockHash template.HTML `json:"pbh"`
}

func AMPS() AMP {
	amp := AMP{
		BlockHeight:   template.HTML(`<a href="/block/{{height}}">{{height}}</a>`),
		BlockHash:     template.HTML(`<a href="/blockhash/{{hash}}">{{hash}}</a>`),
		Tx:            template.HTML(`<a href="/tx/{{.}}">{{.}}</a>`),
		Addr:          template.HTML(`<a href="/addr/{{.}}">{{.}}</a>`),
		NextBlockHash: template.HTML(`<a href="/blockhash/{{nextblockhash}}">{{nextblockhash}}</a>`),
		PrevBlockHash: template.HTML(`<a href="/blockhash/{{previousblockhash}}">{{previousblockhash}}</a>`),
	}
	return amp
}
