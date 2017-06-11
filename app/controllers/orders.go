package controllers

import (
  "github.com/revel/revel"
  "amazon_order_bot/app/amazon"
  "io/ioutil"
  "encoding/json"
)

type Orders struct {
  *revel.Controller
}

type OrderInfo struct {
  Id string `json:"amazon_id"`
}

func (c Orders) Index() revel.Result {
  bytes, _ := ioutil.ReadAll(c.Request.Body)
  var orderInfo OrderInfo
  json.Unmarshal(bytes, &orderInfo)
  revel.INFO.Println(orderInfo.Id)
  amazon.WebDriver{}.Order(orderInfo.Id)
  return c.Render()
}
