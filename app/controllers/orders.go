package controllers

import (
  "github.com/revel/revel"
  "github.com/kazuooooo/amazon_order_bot/app/amazon"
  "io/ioutil"
  "encoding/json"
  "github.com/kazuooooo/amazon_order_bot/app/notifier"
  "fmt"
)

type Orders struct {
  *revel.Controller
}

type OrderInfo struct {
  Id string `json:"amazon_id"`
}

func (c Orders) Index() revel.Result {
  bytes, err := ioutil.ReadAll(c.Request.Body)
  if err != nil {
    s := err.Error()
    fmt.Printf("type: %T; value: %q\n", s, s)
  }
  fmt.Printf("%s", string(bytes))
  // fmt.Printf("%s", string(err))
  var orderInfo OrderInfo
  json.Unmarshal(bytes, &orderInfo)
  revel.INFO.Println(orderInfo.Id)
  notifier.Slack{}.PostMessage(orderInfo.Id + "を注文します。")
  amazon.WebDriver{}.Order(orderInfo.Id)
  notifier.Slack{}.PostMessage(orderInfo.Id + "を注文しました。")
  return c.Render()
}
