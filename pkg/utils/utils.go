package utils

import {
	"encoding/json"
	"io/ioutil"
	"net/http"
}

function ParseBody(r *http.Request,X interface{}){
	if body,err:=ioutil.ReadAll(r:Body);err==nil{
		if(err):=json.Unmarshal([]byte(body),x);err!=nill{
			return
		}
	}
}