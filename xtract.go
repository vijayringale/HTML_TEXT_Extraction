package htmlwordextract

import (
	"bytes"
	"html"
	"io"
	"net/http"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

var unline = -1
var trimFunc = unicode.IsSpace

type TrimFunc func(r rune) bool


func SetTrimFunc(f TrimFunc){
	trimFunc = f
}

func Page(url string)(string,error){
	return PageLim(url , unline)
}

func PageLim(url string,lim int)(string , error){
	rs ,err := http.Get(url)
	if err!= nil {
		return "Error : -" ,err
	}
	defer rs.Body.Close()
	return xtract(rs.Body , lim ),err
}


func Value(htmlVal string) string {
	return ValueLim(htmlVal , unline)
}

func ValueLim(htmlVal string, lim int) string{
	b:= bytes.NewReader([]byte(htmlVal))
	return xtract(b,lim)
}


func xtract(r io.Reader , lim int) string {
	z:= html.NewTokenizer(r)
	rs:=bytes.NewBufferString("")

	for {
		t := z.Next()
		if t == html.ErrorToken {
			return rs.String()

		}
		if t== html.TextToken{
			if trimFunc != nil {
				if rs.Len()>0{
					rs.Write([]byte(""))
				}
				rs.Write(bytes.TrimFunc(z.Text(),trimFunc))
			}else{
				rs.Write(z.Text())
			}
			if lim != unline {
				 v:= strings.Fields(rs.String())
				 wc := len(v)
				 if wc> lim {
					return strings.Join(v[:min(wc,lim)]," ")
				}
			}

		}
	}
}
func min(x,y int) int {
	if x>y {
		return y
	}
	return x
}