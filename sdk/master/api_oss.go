package master

import (
	"encoding/json"
	"net/http"

	"github.com/chubaofs/chubaofs/proto"
	"github.com/chubaofs/chubaofs/util/oss"
)

type OSSAPI struct {
	mc *MasterClient
}

func (api *OSSAPI) GetAKInfo(accesskey string) (akPolicy *oss.AKPolicy, err error) {
	var request = newAPIRequest(http.MethodGet, proto.OSSGetAKInfo)
	request.addParam("ak", accesskey)
	var data []byte
	if data, err = api.mc.serveRequest(request); err != nil {
		return
	}
	akPolicy = &oss.AKPolicy{}
	if err = json.Unmarshal(data, akPolicy); err != nil {
		return
	}
	return
}

func (api *OSSAPI) AddPolicy(accesskey string, policy *oss.UserPolicy) (akPolicy *oss.AKPolicy, err error) {
	var body []byte
	if body, err = json.Marshal(policy); err != nil {
		return
	}
	var request = newAPIRequest(http.MethodPost, proto.OSSAddPolicy)
	request.addParam("ak", accesskey)
	request.addBody(body)
	var data []byte
	if data, err = api.mc.serveRequest(request); err != nil {
		return
	}
	akPolicy = &oss.AKPolicy{}
	if err = json.Unmarshal(data, akPolicy); err != nil {
		return
	}
	return
}

func (api *OSSAPI) DeletePolicy(accesskey string, policy *oss.UserPolicy) (akPolicy *oss.AKPolicy, err error) {
	var body []byte
	if body, err = json.Marshal(policy); err != nil {
		return
	}
	var request = newAPIRequest(http.MethodPost, proto.OSSDeletePolicy)
	request.addParam("ak", accesskey)
	request.addBody(body)
	var data []byte
	if data, err = api.mc.serveRequest(request); err != nil {
		return
	}
	akPolicy = &oss.AKPolicy{}
	if err = json.Unmarshal(data, akPolicy); err != nil {
		return
	}
	return
}

func (api *OSSAPI) DeleteVolPolicy(vol string) (err error) {
	var request = newAPIRequest(http.MethodPost, proto.OSSDeleteVolPolicy)
	request.addParam("name", vol)
	if _, err = api.mc.serveRequest(request); err != nil {
		return
	}
	return
}