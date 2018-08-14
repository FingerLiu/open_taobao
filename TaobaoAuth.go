// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package open_taobao

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"net/url"
)

func GetUrlForAuth(appKey, redirectUri, state string) (*url.URL, error) {
	if appKey == "" {
		return nil, errors.New("[GetUrlForAuth] AppKey is null")
	}
	if redirectUri == "" {
		return nil, errors.New("[GetUrlForAuth] RedirectUri is null")
	}
	tbDomain := viper.GetString("TAOBAO_DOMAIN")
	if len(tbDomain) == 0 {
		tbDomain = "tbsandbox.com"
	}
	u, err := url.Parse(fmt.Sprintf("https://oauth.%s/authorize", tbDomain))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("client_id", appKey)
	q.Set("response_type", "code")
	q.Set("redirect_uri", redirectUri)
	q.Set("state", state)
	q.Set("view", "web")
	u.RawQuery = q.Encode()

	return u, nil
}

type TokenGetResponse struct {
	AccessToken       string `json:"access_token"`
	TokenType         string `json:"token_type"`
	ExpiresIn         int    `json:"expires_in"`
	RefreshToken      string `json:"refresh_token"`
	ReExpiresIn       int    `json:"re_expires_in"`
	R1ExpiresIn       int    `json:"r1_expires_in"`
	R2ExpiresIn       int    `json:"r2_expires_in"`
	W1ExpiresIn       int    `json:"w1_expires_in"`
	W2ExpiresIn       int    `json:"w2_expires_in"`
	TaobaoUserNick    string `json:"taobao_user_nick"`
	TaobaoUserId      string `json:"taobao_user_id"`
	SubTaobaoUserId   string `json:"sub_taobao_user_id"`
	SubTaobaoUserNick string `json:"sub_taobao_user_nick"`
}

type TokenGetRequest struct {
	TaobaoRequest
}

func (t *TokenGetRequest) SetAppKey(val string) {
	t.SetValue("client_id", val)
}

func (t *TokenGetRequest) SetAppSecret(val string) {
	t.SetValue("client_secret", val)
}

func (t *TokenGetRequest) SetCode(val string) {
	t.SetValue("code", val)
}

func (t *TokenGetRequest) SetRedirectUri(val string) {
	t.SetValue("redirect_uri", val)
}

func (t *TokenGetRequest) SetState(val string) {
	t.SetValue("state", val)
}

func (t *TokenGetRequest) GetResponse() (*TokenGetResponse, []byte, error) {
	if t.GetValue("client_id") == "" {
		return nil, nil, errors.New("[TokenGetRequest] AppKey is null")
	}
	if t.GetValue("client_secret") == "" {
		return nil, nil, errors.New("[TokenGetRequest] AppSecret is null")
	}
	if t.GetValue("code") == "" {
		return nil, nil, errors.New("[TokenGetRequest] Code is null")
	}
	if t.GetValue("redirect_uri") == "" {
		return nil, nil, errors.New("[TokenGetRequest] RedirectUri is null")
	}

	tbDomain := viper.GetString("TAOBAO_DOMAIN")
	if len(tbDomain) == 0 {
		tbDomain = "tbsandbox.com"
	}
	t.SetReqUrl(fmt.Sprintf("https://oauth.%s/token", tbDomain))

	t.SetValue("grant_type", "authorization_code")
	t.SetValue("view", "web")

	resp := new(TokenGetResponse)
	data, err := executeRequest(t, resp, InsecureSkipVerify, DisableCompression)
	return resp, data, err
}

type TokenRefreshRequest struct {
	TaobaoRequest
}

func (t *TokenRefreshRequest) SetRefreshToken(val string) {
	t.SetValue("refresh_token", val)
}

func (t *TokenRefreshRequest) GetResponse() (*TokenGetResponse, []byte, error) {
	if t.GetValue("client_id") == "" {
		return nil, nil, errors.New("[TokenGetRequest] AppKey is null")
	}
	if t.GetValue("client_secret") == "" {
		return nil, nil, errors.New("[TokenGetRequest] AppSecret is null")
	}
	if t.GetValue("code") == "" {
		return nil, nil, errors.New("[TokenGetRequest] Code is null")
	}
	if t.GetValue("redirect_uri") == "" {
		return nil, nil, errors.New("[TokenGetRequest] RedirectUri is null")
	}
	if t.GetValue("refresh_token") == "" {
		return nil, nil, errors.New("[TokenGetRequest] RefreshToken is null")
	}
	tbDomain := viper.GetString("TAOBAO_DOMAIN")
	if len(tbDomain) == 0 {
		tbDomain = "tbsandbox.com"
	}
	t.SetReqUrl(fmt.Sprintf("https://oauth.%s/token", tbDomain))

	t.SetValue("grant_type", "refresh_token")
	t.SetValue("view", "web")

	resp := new(TokenGetResponse)
	data, err := executeRequest(t, resp, InsecureSkipVerify, DisableCompression)
	return resp, data, err
}
