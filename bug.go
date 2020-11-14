package main

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
)

func GetHtml(url string)io.Reader{
	resp,err:=http.Get(url)
	if err!=nil{
		log.Fatal("get err",err)
	}
	return resp.Body
}

//func query
func main(){
	_,f1 := os.Create("./page1.html")
	r1:=GetHtml("https://blog.lenconda.top/")
	/*r2:=GetHtml("https://blog.lenconda.top/page/2/")
	r3:=GetHtml("https://blog.lenconda.top/page/3/")
	r4:=GetHtml("https://blog.lenconda.top/page/4/")
	r5:=GetHtml("https://blog.lenconda.top/page/5/")
	r6:=GetHtml("https://blog.lenconda.top/page/6/")
	r7:=GetHtml("https://blog.lenconda.top/page/7/")*/
	dom1,err1:=goquery.NewDocumentFromReader(r1)
	if err1!=nil{
		log.Fatalln(err1)
	}
	/*dom2,err2:=goquery.NewDocumentFromReader(r2)
	if err2!=nil{
		log.Fatalln(err2)
	}
	dom3,err3:=goquery.NewDocumentFromReader(r3)
	if err3!=nil{
		log.Fatalln(err3)
	}
	dom4,err4:=goquery.NewDocumentFromReader(r4)
	if err4!=nil{
		log.Fatalln(err4)
	}
	dom5,err5:=goquery.NewDocumentFromReader(r5)
	if err5!=nil{
		log.Fatalln(err5)
	}
	dom6,err6:=goquery.NewDocumentFromReader(r6)
	if err6!=nil{
		log.Fatalln(err6)
	}
	dom7,err7:=goquery.NewDocumentFromReader(r7)
	if err7!=nil{
		log.Fatalln(err7)
	}*/
	dom1.Find("h2").Each(func(i int, selection *goquery.Selection) {
		_,val1:=dom1.Find("a").Html()
		_,err1:= f1.WriteString(val1)
		if err1 != nil{
			log.Println(err1)
		}
		_,val2:=dom1.Find("span.post-meta").Html()
		_,err2:= f1.WriteString(val2)
		if err2 != nil{
			log.Println(err2)
		}
		_,val3:=dom1.Find("p.post-excerpt").Html()
		_,err3:= f1.WriteString(val3)
		if err3 != nil{
			log.Println(err3)
		}
	})
}