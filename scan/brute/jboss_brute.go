package brute

import (
	"fmt"
	"github.com/yhy0/Jie/pkg/protocols/httpx"
	"github.com/yhy0/logging"
)

func JbossBrute(url string) (username string, password string) {
	if req, err := httpx.RequestBasic("asdasdascsacacs", "asdasdascsacacs", url+"/jmx-console/", "GET", "", false, nil); err == nil {
		if req.StatusCode == 401 {
			for uspa := range jbossuserpass {
				if req2, err2 := httpx.RequestBasic(jbossuserpass[uspa].Username, jbossuserpass[uspa].Password, url+"/jmx-console/", "GET", "", false, nil); err2 == nil {
					if req2.StatusCode == 200 || req2.StatusCode == 403 {
						logging.Logger.Infof(fmt.Sprintf("Found vuln Jboss password|%s:%s|%s\n", jbossuserpass[uspa].Username, jbossuserpass[uspa].Password, url))
						return jbossuserpass[uspa].Username, jbossuserpass[uspa].Password
					}
				}
			}
		}
	}
	return "", ""
}
