package controllers

import (
  "github.com/revel/revel"
  // "amazon_order_bot/app/amazon"
  "io/ioutil"
  "encoding/json"
)

type Orders struct {
  *revel.Controller
}

type Order struct {
  Id string `json:"id"`
}

func (c Orders) Index() revel.Result {
  bytes, _ := ioutil.ReadAll(c.Request.Body)
  var order Order
  json.Unmarshal(bytes, &order)
  revel.INFO.Println(order.AmazonId)

  // amazon.WebDriver{}.Order(order.Id)
  return c.Render()
}
