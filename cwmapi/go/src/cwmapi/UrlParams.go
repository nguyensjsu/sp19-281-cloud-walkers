package main

import (
	"net/url"
	"strconv"
	"strings"
)

type parseInfo struct{
	pathItems []string
	queryParams map[string] string
}

func toInt(s string) (int){
	i, _ := strconv.Atoi(s)
	return i
}

func depth(pp parseInfo)(int){
	return toInt(pp.queryParams["depth"])
}

func parseUrl(url *url.URL) (parseInfo){

	var ret parseInfo

	if(url.Path != ""){
		ret.pathItems = strings.Split(url.Path, "/")
	}

	ret.queryParams = make(map[string] string)

	if(url.RawQuery != "") {
		rq := strings.Split(url.RawQuery, "&>")

		for i := 0; i < len(rq); i++ {
			cq := strings.Split(rq[i], "=")
			ret.queryParams[cq[0]] = cq[1]
		}
	}

	return ret
}

func getPath(pp parseInfo)(string){
	if(len(pp.pathItems) > 0){
		return pp.pathItems[len(pp.pathItems)-1];
	}

	return ""
}
