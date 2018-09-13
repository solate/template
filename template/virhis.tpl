package basic

import (
	"encoding/json"
	"gitlab.dev.okapp.cc/med-his-group/medhis-platform-service/virhis"
	"io/ioutil"
	"net/url"
)

// %s
func (s *VirHIS) %s(req *virhis.%s, reply *virhis.%s) (err error) {

	data := &url.Values{}
	%s
	res, err := HttpGet(s.URL+%s, data)
	if err != nil {
		logger.Error(s.URL+%s, " remote err: ", err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error("read body err: ", err)
		return
	}

	reply = new(virhis.%s)
	err = json.Unmarshal(body, reply)
	if err != nil {
		logger.Error("unmarshal err: ", err)
		return
	}

	return
}