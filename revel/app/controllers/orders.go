package controllers

import (
  "github.com/revel/revel"
  // "amazon_order_bot/revel/app/amazon"
  "io/ioutil"
  "encoding/json"
  // "amazon_order_bot/revel/app/notifier"
)

type Orders struct {
  *revel.Controller
}

type OrderInfo struct {
  Id string `json:"id"`
}

func (c Orders) Index() revel.Result {
  var f interface{}
  body, _ := ioutil.ReadAll(c.Request.Body)
  json.Unmarshal(body, &f)
  ret := f.(map[string]interface{})
  revel.INFO.Println(ret["id"].(string))
  // bytes, _ := ioutil.ReadAll(c.Request.Body)
  // var orderInfo OrderInfo
  // json.Unmarshal(bytes, &orderInfo)
  // revel.INFO.Println(orderInfo)
  // revel.INFO.Println(orderInfo.Id)
  // notifier.Slack{}.PostMessage(orderInfo.Id + "を注文します。")
  // amazon.WebDriver{}.Order(orderInfo.Id)
  // notifier.Slack{}.PostMessage(orderInfo.Id + "を注文しました。")
  return c.Render()
}
