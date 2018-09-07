// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package open_taobao

import (
	"errors"
	"time"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type TaobaoMethodRequest struct {
	TaobaoRequest
}

func (t *TaobaoMethodRequest) GetResponse(accessToken, apiMethodName string, resp interface{}) ([]byte, error) {
	tbDomain := viper.GetString("TAOBAO_DOMAIN")
	if len(tbDomain) == 0 {
		tbDomain = "tbsandbox.com"
	}
	if accessToken == "" {
		return nil, errors.New("[" + apiMethodName + "] AccessToken is null")
	}
	url := fmt.Sprintf("https://eco.%s/router/rest", tbDomain)
	t.SetReqUrl(url)

	t.SetValue("method", apiMethodName)
	t.SetValue("format", "json")
	t.SetValue("v", "2.0")
	if taobaoConfig.isTop {
		t.SetValue("session", accessToken)
		t.SetValue("timestamp", time.Now().Format("2006-01-02 15:04:05"))
		t.SetValue("app_key", taobaoConfig.appKey)
		t.SetValue("sign_method", "md5")
		reqUrl := fmt.Sprintf("http://gw.api.%s/router/rest", tbDomain)
		t.SetReqUrl(reqUrl)

		if taobaoConfig.requestUrl != "" {
			t.SetReqUrl(taobaoConfig.requestUrl)
		}
	} else {
		t.SetValue("access_token", accessToken)
	}
	log.Debug().Msgf("tmall reqUrl %s", t.GetReqUrl())
	log.Debug().Msgf("tmall values %s", t.GetValues())

	var debug string
	debug = t.GetValue("debug")
	if debug == "true" || debug == "True" {
		log.Info().Msg("fake tmall, not call tmall due to debug == true")
		var res []byte
		str := "{\"code\": \"isv.success-all\", \"msg\":\"成功\", \"debug\":\"true\"}"
		res = []byte(str)
		return res, nil
	}

	return executeRequest(t, resp, InsecureSkipVerify, DisableCompression)
}
