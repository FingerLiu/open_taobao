// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package eticket

import (
	"github.com/FingerLiu/open_taobao"
	"encoding/json"
)

/* 码商重发电子凭证回调接口 */
type EticketMerchantMaResendRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 业务类型 */
func (r *EticketMerchantMaResendRequest) SetBizType(value string) {
	r.SetValue("biz_type", value)
}

/* 待重发的码列表 */
func (r *EticketMerchantMaResendRequest) SetIsvMaList(value []IsvMa) {
	j, _ := json.Marshal(value)
	r.SetValue("isv_ma_list", string(j))
}

/* 业务id（订单号） */
func (r *EticketMerchantMaResendRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 需要跟发码通知获取到的参数一致 */
func (r *EticketMerchantMaResendRequest) SetToken(value string) {
	r.SetValue("token", value)
}

func (r *EticketMerchantMaResendRequest) GetResponse(accessToken string) (*EticketMerchantMaResendResponse, []byte, error) {
	var resp EticketMerchantMaResendResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.eticket.merchant.ma.resend", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type EticketMerchantMaResendResponse struct {
	RespBody *SendMaCallbackResp `json:"resp_body"`
	RetCode  string              `json:"ret_code"`
	RetMsg   string              `json:"ret_msg"`
}

type EticketMerchantMaResendResponseResult struct {
	Response *EticketMerchantMaResendResponse `json:"eticket_merchant_ma_resend_response"`
}

/* 电子凭证核销接口 */
type EticketMerchantMaConsumeRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 业务类型 */
func (r *EticketMerchantMaConsumeRequest) SetBizType(value string) {
	r.SetValue("biz_type", value)
}

/* 需要被核销的码 */
func (r *EticketMerchantMaConsumeRequest) SetCode(value string) {
	r.SetValue("code", value)
}

/* 核销份数 */
func (r *EticketMerchantMaConsumeRequest) SetConsumeNum(value string) {
	r.SetValue("consume_num", value)
}

/* 核销后换码的码列表 */
func (r *EticketMerchantMaConsumeRequest) SetIsvMaList(value []IsvMa) {
	j, _ := json.Marshal(value)
	r.SetValue("isv_ma_list", string(j))
}

/* 业务id（订单号） */
func (r *EticketMerchantMaConsumeRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 机具编号 */
func (r *EticketMerchantMaConsumeRequest) SetPosId(value string) {
	r.SetValue("pos_id", value)
}

/* 核销序列号，需要保证唯一 */
func (r *EticketMerchantMaConsumeRequest) SetSerialNum(value string) {
	r.SetValue("serial_num", value)
}

/* 需要跟发码通知获取到的参数一致 */
func (r *EticketMerchantMaConsumeRequest) SetToken(value string) {
	r.SetValue("token", value)
}

func (r *EticketMerchantMaConsumeRequest) GetResponse(accessToken string) (*EticketMerchantMaConsumeResponse, []byte, error) {
	var resp EticketMerchantMaConsumeResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.eticket.merchant.ma.consume", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type EticketMerchantMaConsumeResponse struct {
	RespBody *ConsumeMaCallbackResp `json:"resp_body"`
	RetCode  string                 `json:"ret_code"`
	RetMsg   string                 `json:"ret_msg"`
}

type EticketMerchantMaConsumeResponseResult struct {
	Response *EticketMerchantMaConsumeResponse `json:"eticket_merchant_ma_consume_response"`
}

/* 码商发码成功回调接口 */
type EticketMerchantMaSendRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 业务类型 */
func (r *EticketMerchantMaSendRequest) SetBizType(value string) {
	r.SetValue("biz_type", value)
}

/* 需要发送的码列表 */
func (r *EticketMerchantMaSendRequest) SetIsvMaList(value []IsvMa) {
	j, _ := json.Marshal(value)
	r.SetValue("isv_ma_list", string(j))
}

/* 业务id（订单号） */
func (r *EticketMerchantMaSendRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 需要跟发码通知获取到的参数一致 */
func (r *EticketMerchantMaSendRequest) SetToken(value string) {
	r.SetValue("token", value)
}

func (r *EticketMerchantMaSendRequest) GetResponse(accessToken string) (*EticketMerchantMaSendResponse, []byte, error) {
	var resp EticketMerchantMaSendResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.eticket.merchant.ma.send", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type EticketMerchantMaSendResponse struct {
	RespBody *SendMaCallbackResp `json:"resp_body"`
	RetCode  string              `json:"ret_code"`
	RetMsg   string              `json:"ret_msg"`
}

type EticketMerchantMaSendResponseResult struct {
	Response *EticketMerchantMaSendResponse `json:"eticket_merchant_ma_send_response"`
}

/* 针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证 */
type EticketMerchantMaFailsendRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 业务类型 */
func (r *EticketMerchantMaFailsendRequest) SetBizType(value string) {
	r.SetValue("biz_type", value)
}

/* 业务id（订单号） */
func (r *EticketMerchantMaFailsendRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 错误原因码 */
func (r *EticketMerchantMaFailsendRequest) SetSubCode(value string) {
	r.SetValue("sub_code", value)
}

/* 错误码描述 */
func (r *EticketMerchantMaFailsendRequest) SetSubMsg(value string) {
	r.SetValue("sub_msg", value)
}

/* 需要与发码通知获取的值一致 */
func (r *EticketMerchantMaFailsendRequest) SetToken(value string) {
	r.SetValue("token", value)
}

func (r *EticketMerchantMaFailsendRequest) GetResponse(accessToken string) (*EticketMerchantMaFailsendResponse, []byte, error) {
	var resp EticketMerchantMaFailsendResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.eticket.merchant.ma.failsend", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type EticketMerchantMaFailsendResponse struct {
	RespBody *SendFailCallbackResp `json:"resp_body"`
	RetCode  string                `json:"ret_code"`
	RetMsg   string                `json:"ret_msg"`
}

type EticketMerchantMaFailsendResponseResult struct {
	Response *EticketMerchantMaFailsendResponse `json:"eticket_merchant_ma_failsend_response"`
}
