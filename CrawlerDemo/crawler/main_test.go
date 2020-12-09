package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
	"golang.org/x/net/html"
)

func TestNoCookie(t *testing.T) {
	c := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			t.Log("req: ", req)
			t.Log("via num: ", len(via))
			t.Log("via: ", via[0])
			return http.ErrUseLastResponse
		},
	}

	u, _ := url.Parse("https://beian.miit.gov.cn/")
	req := &http.Request{
		Method: "GET",
		URL:    u,
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Println("error:", err)
		return
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	// log.Println(resp.Status)
	// log.Println(string(body))
	// log.Println(resp.Header.Get("Set-Cookie"))

	nodes := GetScriptNodes(string(body))
}

func TestMain(t *testing.T) {
	c := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			t.Log("req: ", req)
			t.Log("via num: ", len(via))
			t.Log("via: ", via[0])
			return http.ErrUseLastResponse
		},
	}

	u, _ := url.Parse("https://beian.miit.gov.cn/")
	req := &http.Request{
		Method: "GET",
		URL:    u,
	}

	cookie1 := &http.Cookie{
		Name:   "__jsl_clearance",
		Value:  "1607477369.705|0|58zrDTjXSxEcoaex%2FmMQRxB6K5M%3D",
		MaxAge: 3600,
		Path:   "/",
	}

	cookie2 := &http.Cookie{
		Name:   "__jsluid_s",
		Value:  "b22d8bb35f3a95528dc62bcb9c4df624",
		MaxAge: 3600,
		Path:   "/",
	}

	req.Header = make(http.Header)
	req.AddCookie(cookie1)
	req.AddCookie(cookie2)
	req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36 Edg/87.0.664.57"}

	resp, err := c.Do(req)
	if err != nil {
		log.Println("error:", err)
		return
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(resp.Status)
	log.Println(string(body))
}

func TestJS(t *testing.T) {
	vm := otto.New()
	vm.Run(`
	cookie=('_')+('_')+('j')+('s')+('l')+('_')+('c')+('l')+('e')+('a')+('r')+('a')+('n')+('c')+('e')+('_')+('s')+('=')+((+true)+'')+([2]*(3)+'')+(~~''+'')+(2+5+'')+(-~(3)+'')+(-~[6]+'')+(1+6+'')+(~~''+'')+(~~''+'')+(-~[6]+'')+('.')+(0+1+0+1+'')+(1+8+'')+(3+3+'')+('|')+('-')+(-~{}+'')+('|')+('l')+('L')+('w')+('W')+('b')+('N')+('Q')+('F')+('L')+('p')+(5+'')+('H')+('z')+('t')+('z')+(7+'')+('d')+('x')+(5+'')+('%')+(+!+[]*2+'')+('F')+('e')+('G')+(+!+[]+'')+('K')+('K')+('I')+('A')+('%')+((1+[2]>>2)+'')+('D')+(';')+('m')+('a')+('x')+('-')+('a')+('g')+('e')+('=')+(-~[2]+'')+(6+'')+(~~''+'')+(~~''+'')+(';')+('p')+('a')+('t')+('h')+('=')+('/');location.href=location.pathname+location.search
	`)

	value, err := vm.Get("cookie")
	if err != nil {
		t.Log(err)
		return
	}

	t.Log(value.ToString())
}

func GetScriptNodes(body string) []*html.Node {
	root, err := html.Parse(strings.NewReader(body))
	if err != nil {
		log.Println("error:", err)
		return nil
	}

	doc := goquery.NewDocumentFromNode(root)

	s := doc.Find("Script")
	if s == nil {
		return nil
	}

	return s.Nodes
}

func GetCookieFromScriptNode(node *html.Node) string {
	contentNode := node.FirstChild

	if contentNode == nil {
		return ""
	}

	content := contentNode.Data

	a
}
