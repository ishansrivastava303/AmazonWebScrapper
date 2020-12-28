package main

import (
    "fmt"
    //b "AmazonProject/DBStore"
    "log"
    "net/http"
    //"os"
    "strings"
    "time"
    "github.com/PuerkitoBio/goquery"
    "encoding/json"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    //_ "github.com/go-sql-driver/mysql"
    //"io/ioutil"
    //"github.com/gocolly/colly"
)

/*const(
    sqlite3Str="sqlite3"
    memStr=":memory:"
)*/

type Fact struct{
    Name string `json:"name"`
    ImageURL string `json:"imageURL"`
    Description []string `json:"description"`
    Price []string `json:"price"`
    TotalReviews string `json:"totalReviews"`
}


type Product struct{
    Url string `json:"Url"`
    ProductDetails Fact `json:"ProductDetails"` 
}

var name,imageURL,totalReviews string
var description [] string
var price [] string

func main() {
    
    url:="https://www.amazon.com/PlayStation-4-Pro-1TB-Console/dp/B01LOP8EZC/?th=1"    
    response, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    
    document, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error loading HTTP response body. ", err)
    }

    document.Find("#imgTagWrapperId").Each(func(index int, parent *goquery.Selection){
    parent.Find("img").Each(func(index int, element *goquery.Selection) {
        imgALT, existsALT := element.Attr("alt")
        imgURL, existsURL := element.Attr("data-a-dynamic-image")
        split:=strings.Split(imgURL,",")
        indexOfJPG:=strings.Index(split[0],".jpg")
        if existsALT && existsURL  {
            name=imgALT
            imageURL=split[0][2:(indexOfJPG+4)]
        }
    })})
    
    var reviewsFlag=1
    document.Find("#acrCustomerReviewText").Each(func(index int, element *goquery.Selection) {
        reviews:=element.Text()
        if reviewsFlag==1{
            totalReviews=reviews
            reviewsFlag=0
        }
        
    })
    
    priceFlag:=0
    
    document.Find("#priceblock_ourprice").Each(func(index int, element *goquery.Selection) {
        priceFlag=1
        //price:=element.Text()
        price=append(price,element.Text())
    })
    
    if priceFlag==0{
      document.Find("#edition_0_price").Each(func(index int, element *goquery.Selection) {

          str:=strings.ReplaceAll(element.Text(),"\n","")
          price=append(price,strings.Trim(str," "))
        
    })
        document.Find("#edition_1_price").Each(func(index int, element *goquery.Selection) {
        
       str:=strings.ReplaceAll(element.Text(),"\n","")        
            price=append(price,strings.Trim(str," "))
    }) 
        //fmt.Println(prices)
    }
    
      document.Find("#feature-bullets").Each(func(index int, feature *goquery.Selection) {
      
          feature.Find("li").Each(func(index int, list *goquery.Selection) {
              str:=strings.ReplaceAll(list.Text(),"\n","")
              description=append(description,strings.Trim(str," "))
    })    
    })
    
    
    fact:=Fact{
        Name:name,
        ImageURL:imageURL,
        Description:description,
        Price:price,
        TotalReviews:totalReviews,
    }
    
    
    allFacts:=make([]Fact,0)
    allFacts=append(allFacts,fact)
    
    
    
    product:=Product{
        Url:url,
        ProductDetails:fact,
    }
    productDetails:=make([]Product,0)
    productDetails=append(productDetails,product)
    //enc:=json.NewEncoder(os.Stdout)
    //enc.SetIndent(""," ")
    //enc.Encode(allFacts)
    s,_:=json.MarshalIndent(product,""," ")
    //log.Println(string(b))
    //b.DbStore(string(s),fact.Name)
    
    timestamp:=time.Now().Format("01-02-2006 15:04:05")
    
    database,err:=sql.Open("sqlite3","./DBStore/ProductDetails.db")

    if err!=nil{
        fmt.Println("not able to connect to db server")
    }    
    
    defer database.Close()

    fmt.Println("Connected successfully")
    
    
    statement,_:=database.Prepare("CREATE TABLE IF NOT EXISTS WEBSCRAPER(PRODUCT_NAME TEXT PRIMARY KEY,PRODUCT_DETAILS_JSON TEXT,TIMESTAMP TEXT)")       
    
    if err!=nil{
        panic(err.Error())
    }    
    statement.Exec()    
    
    fmt.Println("WEBSCRAPER TABLE CREATED")
    statement,err=database.Prepare("INSERT INTO WEBSCRAPER(PRODUCT_NAME,PRODUCT_DETAILS_JSON,TIMESTAMP) VALUES(?,?,?)")
    
    if err!=nil{
        panic(err.Error())
    }    
    statement.Exec(fact.Name,string(s),timestamp)

    
    
}


    

