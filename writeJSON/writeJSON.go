package writeJSON

import (
    "fmt"
    //"log"
    //"net/http"
    //"os"
    //"strings"
    //"github.com/PuerkitoBio/goquery"
    //"encoding/json"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "time"
    //"io/ioutil"
    //"github.com/gocolly/colly"
)

func DbStore(productDetails string,productName string){
    
    timestamp:=time.Now().Format("01-02-2006 15:04:05")
    
    database,err:=sql.Open("sqlite3","./ProductDetails.db")

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
    statement.Exec(productName,productDetails,timestamp)

    rows,_:=database.Query("SELECT PRODUCT_NAME,PRODUCT_DETAILS_JSON,TIMESTAMP FROM WEBSCRAPER")

    var PRODUCT_NAME string
    var PRODUCT_DETAILS_JSON string
    var dt string   
    for rows.Next(){
        rows.Scan(&PRODUCT_NAME,&PRODUCT_DETAILS_JSON,&dt)
        fmt.Println("PRODUCT_NAME:"+PRODUCT_NAME)
        fmt.Println("PRODUCT_DETAILS_JSON:"+PRODUCT_DETAILS_JSON)
        fmt.Println("TIMESTAMP:"+dt+"\n")
}
    /*file,err:=json.MarshalIndent(data,""," ")
    
    if err!=nil{
        log.Println("Unable to create JSON file")
        return
    }*/
    //_=ioutil.WriteFile("/home/ishu/Desktop/AmazonProject/AmazonProductDetails.json",[]byte(data),0644)
    //log.Println(data)
}
