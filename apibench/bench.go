package apibench

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup //定义一个同步等待的组
var wgPost sync.WaitGroup

func bench()  {

	baseHost := "http://localhost:7101/api"
	//baseHost = "http://120.55.168.35:7101/api"
	baseHost = "https://fungameh.top/api"

	envRouter := baseHost + "/wechat/apply/env"
	submitRouter := baseHost + "/wechat/apply/submit"

	//test get

	for i:= 0;i<100;i++{
		wg.Add(1) //添加一个计数
		go testGet(envRouter)
	}
	wg.Wait() //阻塞直到所有任务完成

	//test post



	for i:= 10;i<40;i++{
		index := i
		wgPost.Add(1) //添加一个计数
		data := make(map[string]interface{})
		data["open_id"] =  "0123456789012345678901234567"
		data["name"] = "单元测试1"
		data["street"] =  1
		data["reg_time"] = "2019"
		data["employee"] = 10
		data["employee_insure"] = 10
		data["reg_address"]= "注册地址"
		data["last_three_sales"] = []float32{2,3,4}
		data["last_three_assets"] = []float32{2,3,3}
		data["tech_state"] = 1
		data["patent_num"] = 1
		data["is_accident"] = 0
		data["contact_name"] = "联系人"
		data["contact_phone"] = "18940183332"

		data["name"] = fmt.Sprintf("单元测试1-%d",index)
		data["open_id"] = fmt.Sprintf("01234567890123456789012345%d",index)
		go testPost(submitRouter,data)
	}
	wgPost.Wait() //阻塞直到所有任务完成
	fmt.Println("over")

}

func testGet(url string)  {
	startTimetamp := time.Now()
	fmt.Println("request get -->",url)
	resp, err := http.Get(url)


	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	elapsed := time.Since(startTimetamp)
	fmt.Println("request get <--",string(body))
	fmt.Println("该函数执行完成耗时：", elapsed)
	wg.Done()
}

func testPost(url string,payload interface{})  {
	bytesData, _ := json.Marshal(payload)
	startTimetamp := time.Now()

	fmt.Println("request post -->",url)
	fmt.Println("request post request-->",string(bytesData))
	resp, err := http.Post(url,"application/json", bytes.NewReader(bytesData))


	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	elapsed := time.Since(startTimetamp)
	fmt.Println("request post <--",string(body))
	fmt.Println("该函数执行完成耗时：", elapsed)
	wgPost.Done()
}