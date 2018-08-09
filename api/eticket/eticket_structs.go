// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package eticket

const VersionNo = "20170816"

/* 回复参数 */
type SendMaCallbackResp struct {
	AttributeMap `json:"attribute_map"`
}

/* 回复结果 */
type ConsumeMaCallbackResp struct {
	AttributeMap `json:"attribute_map"`
}

/* 回复参数 */
type SendFailCallbackResp struct {
	AttributeMap `json:"attribute_map"`
}

type IsvMa struct {
	Code 	string 	`json:"code"`
	Num	 	int 	`json:"num"`
}
