// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package eticket

const VersionNo = "20180806"

type Send struct {
	Token     	string `json:"token"`
	OuterId   	string `json:"outer_id"`
	BizType   	int `json:"biz_type"`
	IsvMaList 	struct {
		IsvMa []*IsvMa `json:"isv_ma"`
	}           `json:"isv_ma_list"`
}

type ReSend struct {
	Token     	string `json:"token"`
	OuterId   	string `json:"outer_id"`
	BizType   	int `json:"biz_type"`
	IsvMaList 	struct {
		IsvMa []*IsvMa `json:"isv_ma"`
	}           `json:"isv_ma_list"`
}

type IsvMa struct {
	Code string `json:"code"`
	Num int `json:"num"`
}

type IsvMaDetail struct {
	Code string 	`json:"code"`
	Num int 		`json:"num"`
	QrImage string	`json:"qr_image"`
}

type SendFail struct {
	Token     	string `json:"token"`
	OuterId   	string `json:"outer_id"`
	BizType   	int `json:"biz_type"`
	SubCode   	string `json:"sub_code"`
	SubMsg   	string `json:"sub_msg"`
}

type Consume struct {
	Token     	string  `json:"token"`
	OuterId   	string	`json:"outer_id"`
	BizType   	int 	`json:"biz_type"`
	Code   		string 	`json:"code"`
	ConsumeNum  int 	`json:"consume_num"`
	IsvMaList 	struct {
		IsvMaDetail []*IsvMa `json:"isv_ma"`
	}			`json:"isv_ma_list"`
	PosId		string 	`json:"pos_id"`
	SerialNum	string	`json:"serial_num"`
}

