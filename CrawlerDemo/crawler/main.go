package main

import (
	"flag"
	"log"
	"path"

	"github.com/pantskun/commonutils/pathutils"
	"github.com/pantskun/golearn/CrawlerDemo/crawlerutil"
	"github.com/pantskun/golearn/CrawlerDemo/etcd"
	"github.com/pantskun/golearn/CrawlerDemo/htmlutil"
	"github.com/pantskun/golearn/CrawlerDemo/xcrawler"
	"github.com/pantskun/golearn/CrawlerDemo/xlogutil"
)

func main() {
	var (
		u               string
		depth           int
		useMultiprocess bool
		logPath         string
	)

	flag.StringVar(&u, "url", "https://www.ssetech.com.cn", "url")
	flag.IntVar(&depth, "depth", 1, "crawler depth")
	flag.BoolVar(&useMultiprocess, "useMultiprocess", false, "use multiprocess")
	flag.StringVar(&logPath, "log", "", "set the log output file")
	flag.Parse()

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	if logPath != "" {
		logfile, err := htmlutil.CreateFile(path.Join(pathutils.GetModulePath("CrawlerDemo"), "logs", logPath))
		if err != nil {
			xlogutil.Error("log file create faild")
		} else {
			log.SetOutput(logfile)
		}
	}

	var etcdInteractor etcd.Interactor

	if useMultiprocess {
		i, err := etcd.NewInteractor()
		if err != nil {
			xlogutil.Error(err)
			return
		}

		etcdInteractor = i
	} else {
		i, err := etcd.NewInteractorWithEmbed()
		if err != nil {
			xlogutil.Error(err)
			return
		}

		etcdInteractor = i
	}

	handleHTML := func(element xcrawler.HTMLElement) {
		crawlerutil.HandleElementWithURL(element, etcdInteractor)
	}

	handleResp := func(resp xcrawler.Response) {
		crawlerutil.HandleDownloadResp(resp, etcdInteractor)
	}

	c := xcrawler.NewCrawler(depth)

	c.AddHTMLHandler(handleHTML, crawlerutil.FilterElementWithURL)
	c.AddResponseHandler(handleResp)

	c.Visit(u)

	xlogutil.Warning("finished")
}
