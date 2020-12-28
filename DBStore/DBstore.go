package main

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
    //"time"
    //"io/ioutil"
    //"github.com/gocolly/colly"
)

func main(){
    DbStore()
}

func DbStore(){
    
   
    
    database,err:=sql.Open("sqlite3","./ProductDetails.db")

    if err!=nil{
        fmt.Println("not able to connect to db server")
    }    
    
    defer database.Close()

    fmt.Println("Connected successfully")
    
    
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
    
}
