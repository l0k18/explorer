package amp

import "html/template"

type AMP struct {
	BlockHeight    template.HTML `json:"blockheight"`
	BlockHash      template.HTML `json:"blockhash"`
	Tx             template.HTML `json:"tx"`
	Addr           template.HTML `json:"addr"`
	NextBlockHash  template.HTML `json:"nbh"`
	PrevBlockHash  template.HTML `json:"pbh"`
	CurrentBlockTx template.HTML `json:"cbt"`
	Date           template.HTML `json:"date"`
}

func AMPS() AMP {
	amp := AMP{
		BlockHeight:    template.HTML(`<button on="tap:AMP.setState({id: '{{height}}', tp: 'block'})">{{height}}</button>`),
		BlockHash:      template.HTML(`<button on="tap:AMP.setState({id: '{{hash}}', tp: 'block'})">{{hash}}</button>`),
		Tx:             template.HTML(`<button on="tap:AMP.setState({id: '{{.}}', tp: 'tx'})">{{.}}</button>`),
		Addr:           template.HTML(`<button on="tap:AMP.setState({id: '{{.}}', tp: 'addr'})">{{.}}</button>`),
		NextBlockHash:  template.HTML(`<button on="tap:AMP.setState({id: '{{nextblockhash}}', tp: 'block'})">{{nextblockhash}}</button>`),
		PrevBlockHash:  template.HTML(`<button on="tap:AMP.setState({id: '{{previousblockhash}}', tp: 'block'})">{{previousblockhash}}</button>`),
		CurrentBlockTx: template.HTML(`<button on="tap:AMP.setState({id: '{{currentblocktx}}', tp: 'block'})">{{currentblocktx}}</button>`),
		Date:           template.HTML(`<template type="amp-mustache"><small>{{dayName}} {{day}} {{monthName}} {{year}}</small></template>`),
	}
	return amp
}
