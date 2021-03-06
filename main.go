package main
//rules.txt存放的是除了域名后的规则，例如规则 "xx.com" AND "passw"则 rules.txt写的是 AND "passw"

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func getExecutePath() string {  //获取路径
	dir, err := filepath.Abs("./")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
	return dir
}




func makeRules() {
	curPath := getExecutePath()
	file_bytes1, err := ioutil.ReadFile(curPath + "/domains.txt")
	if err != nil {
		panic(err)
	}
	domain_list := strings.Split(string(file_bytes1), "\n")
	//fmt.Println(domain)


	file_bytes2, err := ioutil.ReadFile(curPath + "/rules.txt")
	if err != nil {
		panic(err)
	}
	rule_list := strings.Split(string(file_bytes2), "\n")


	for _, v1 := range domain_list {
		//fmt.Println(v1)
		for _,v2 := range rule_list {
			//strings.TrimSpace(v1) //去除两边多余的符号
			rule := "\"" + strings.TrimSpace(v1) + "\" " + v2  // 拼接域名和规则
			fmt.Println(rule)
			

		}

	}


}


func main() {
	makeRules()
}