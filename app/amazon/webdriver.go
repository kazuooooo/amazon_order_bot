package amazon

import (
  "os"
  "github.com/revel/revel"
  "github.com/sclevine/agouti"
)

var EMAIL    = os.Getenv("AMAZON_ORDER_EMAIL")
var PASSWORD = os.Getenv("AMAZON_ORDER_PASSWORD")
var ENVIRONMENT = os.Getenv("AMAZON_ORDER_ENVIRONMENT")


type WebDriver struct {

}

func (c WebDriver) Hello() {
  print("hello webdriver")
}


func (c WebDriver) Order(productId string) string {
    // Declare Web Driver
    agoutiDriver := agouti.PhantomJS()

    agoutiDriver.Start()
    defer agoutiDriver.Stop() // defer 延ばす、延期する
    page, _ := agoutiDriver.NewPage()

    // Login
    // go to amazon top page
    if err := page.Navigate("https://www.amazon.com/"); err != nil {
      log.Fatalf("Navigate Error:%v", err)
    }

    // click sign in
    if err := page.Find("#nav-link-accountList").Click(); err != nil {
      log.Fatalf("Find Error:%v", err)
    }
    // input form
    login(page)

    // submit
    page.Find("#signInSubmit").Click()

    // Purchase
    // go to product page
    revel.INFO.Println("https://www.amazon.co.jp/dp/" + productId)
    page.Navigate("https://www.amazon.co.jp/dp/" + productId)
            page.Screenshot("b")

    // add to cart
    page.Find("#add-to-cart-button").Click()
    // proceed to payment
    page.Find("#hlb-ptc-btn-native").Click()

    // ReLogin TODO: make DRY
    // input form
    login(page)
    // submit
    page.Find("#signInSubmit").Click()

    // purchase
    // page.FindByName("placeYourOrder1").Click()
    page.Screenshot("Ordered.png")
    return "hoge"
}

func login(page *agouti.Page) {
  page.Find("#ap_email").Fill(EMAIL)
  page.Find("#ap_password").Fill(PASSWORD)
}
