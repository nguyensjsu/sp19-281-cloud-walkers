package main

import (
	//"fmt"
	"net/url"
	"strconv"
	"strings"
)

type parseInfo struct{
	pathItems []string
	queryParams url.Values
}

func toInt(s string) (int){
	i, _ := strconv.Atoi(s)

	if(i < 0){
		i = 100
	}
	return i
}

func depth(pp parseInfo)(int){
	d, ex := pp.queryParams["depth"]

	if(ex == true){
		return toInt(d[0]);
	}
	return 0;
}

func parseUrl(url *url.URL) (parseInfo){

	var ret parseInfo

	if(url.Path != ""){
		ret.pathItems = strings.Split(url.Path, "/")
	}

	ret.queryParams =  url.Query()

	return ret
}

func getPath(pp parseInfo)(string){
	if(len(pp.pathItems) > 0){
		return pp.pathItems[len(pp.pathItems)-1];
	}

	return ""
}

func queryVals(values url.Values , key string)[]string{
	return values[key]
}

func paramCount(values url.Values, id string)(int){
	return len(values[id])
}
