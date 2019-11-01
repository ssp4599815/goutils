package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

/*
常见函数：
func (t Result) Exists() bool // 判断某值是否存在
func (t Result) Value() interface{}
func (t Result) Int() int64
func (t Result) Uint() uint64
func (t Result) Float() float64
func (t Result) String() string
func (t Result) Bool() bool
func (t Result) Time() time.Time
func (t Result) Array() []gjson.Result
func (t Result) Map() map[string]gjson.Result
func (t Result) Get(path string) Result
func (t Result) ForEach(iterator func(key, value Result) bool) // 可传闭包函数
func (t Result) Less(token Result, caseSensitive bool) bool
*/

func main() {
	pwd, _ := os.Getwd()
	jsonBytes, _ := ioutil.ReadFile(path.Join(pwd, "json/gjson/demo.json"))
	json := string(jsonBytes)

	// 判断 json 是否合法
	if !gjson.Valid(json) {
		log.Fatal("invalid json ")
	}
	// 获取json 中的 age
	fmt.Println(gjson.Get(json, "age").Int())

	// 获取last name
	fmt.Println(gjson.Get(json, "name.last").String())

	// 获取 childern 数组
	childrens := gjson.Get(json, "children").Array()
	fmt.Println(childrens)

	// 获取 第二个孩子
	fmt.Println(gjson.Get(json, "children.1").String())
	fmt.Println(gjson.Get(json, "children|1").String())
	// 使用通配符获取第三个孩子
	fmt.Println(gjson.Get(json, "child*.2").String())

	// 使用过滤器 或者 自定义过滤器 |@fun:arg
	// 反转数组函数
	fmt.Println(gjson.Get(json, "children|@reverse").Array())

	// 自定义函数 - 更改为大写
	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)
		}
		return json
	})
	fmt.Println(gjson.Get(json, "children|@case:upper").Array())

	// 直接解析为 map
	jsonMap := gjson.Parse(json).Map()
	fmt.Println(jsonMap)
}
