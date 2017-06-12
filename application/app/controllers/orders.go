package controllers

import (
  "github.com/revel/revel"
  "amazon_order_bot/applicationapp/amazon"
  "io/ioutil"
  "encoding/json"
  "amazon_order_bot/app/notifier"
)

type Orders struct {
  *revel.Controller
}

type OrderInfo struct {
  Id string `json:"amazon_id"`
}

func (c Orders) Index() revel.Result {
  var jsonData map[string]interface{}
  c.Params.BindJSON(&jsonData)
  revel.INFO.Println(jsonData["amazon_id"])
  // bytes, _ := ioutil.ReadAll(c.Request.Body)
  // var orderInfo OrderInfo
  // json.Unmarshal(bytes, &orderInfo)
  // revel.INFO.Println(orderInfo.Id)
  // notifier.Slack{}.PostMessage(orderInfo.Id + "を注文します。")
  // amazon.WebDriver{}.Order(orderInfo.Id)
  // notifier.Slack{}.PostMessage(orderInfo.Id + "を注文しました。")
  return c.Render()
}
