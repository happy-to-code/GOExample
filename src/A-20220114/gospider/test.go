package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"strings"
	"time"
)

func main1() {
	url := "http://news.iciba.com/"
	selector := "body > div > div.banner > div.swiper-container-place > div > div.swiper-slide.swiper-slide-0.swiper-slide-visible.swiper-slide-active > a.item.item-big > div.item-bottom"
	sel := "document.querySelector(\"body\")"

	content, err := GetHttpHtmlContent(url, selector, sel)
	if err != nil {
		panic(err)
	}
	data, err := GetSpecialData(content, ".chinese")
	if err != nil {
		panic(err)
	}
	fmt.Println("==>", data)

}

// GetHttpHtmlContent 获取网站上爬取的数据
func GetHttpHtmlContent(url string, selector string, sel interface{}) (string, error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	// 初始化参数，先传一个空的数据
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)

	// create context
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	// 执行一个空task, 用提前创建Chrome实例
	chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	// 创建一个上下文，超时时间为40s
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 40*time.Second)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.OuterHTML(sel, &htmlContent, chromedp.ByJSPath),
	)
	if err != nil {
		log.Println("Run err : %v\n", err)
		return "", err
	}
	log.Println(htmlContent)

	return htmlContent, nil
}

// var = `<div class="cell_cell__EYJT6 cell_scrollable__KigLT"><table role="table" class="table_table__fuS_N"><thead><tr role="row"><th colspan="1" role=
// "columnheader" aria-sort="none"><button type="button" title="Toggle SortBy" style="cursor: pointer;">Label</button></th><th colspan="1" role="columnheader" aria-sort="n
// one"><button type="button" title="Toggle SortBy" style="cursor: pointer;">Type</button></th><th colspan="1" role="columnheader" aria-sort="none"><button type="button" t
// itle="Toggle SortBy" style="cursor: pointer;">Creator</button></th><th colspan="1" role="columnheader" aria-sort="none"><button type="button" title="Toggle SortBy" styl
// e="cursor: pointer;">Source</button></th><th colspan="1" role="columnheader" aria-sort="none"><button type="button" title="Toggle SortBy" style="cursor: pointer;">Updat
// ed</button></th></tr></thead><tbody role="rowgroup"><tr role="row"><td role="cell">depositcontract</td><td role="cell">contract_name</td><td role="cell">dune</td><td ro
// le="cell">ethereum_mainnet_contracts</td><td role="cell"><time datetime="Tue, 23 Mar 2021 14:45:14 GMT" title="Tue, 23 Mar 2021 14:45:14 GMT">11 months ago</time></td><
// /tr><tr role="row"><td role="cell">eth2</td><td role="cell">project</td><td role="cell">dune</td><td role="cell">ethereum_mainnet_contracts</td><td role="cell"><time da
// tetime="Fri, 05 Mar 2021 09:40:08 GMT" title="Fri, 05 Mar 2021 09:40:08 GMT">12 months ago</time></td></tr><tr role="row"><td role="cell">uniswap user</td><td role="cel
// l">dapp usage</td><td role="cell">hagaetc</td><td role="cell">ethereum_mainnet_dex_traders</td><td role="cell"><time datetime="Mon, 07 Dec 2020 10:06:37 GMT" title="Mon
// , 07 Dec 2020 10:06:37 GMT">1 year ago</time></td></tr><tr role="row"><td role="cell">dex trader</td><td role="cell">activity</td><td role="cell">hagaetc</td><td role="
// cell">ethereum_mainnet_activity_type</td><td role="cell"><time datetime="Mon, 07 Dec 2020 10:06:37 GMT" title="Mon, 07 Dec 2020 10:06:37 GMT">1 year ago</time></td></tr
// ></tbody></table><ul class="table_footer__aB65O table_footer-defaults__4f8i4"></ul></div>`

// GetSpecialData 得到具体的数据
func GetSpecialData(htmlContent string, selector string) ([]string, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Println(err)
		return []string{}, err
	}

	var str []string
	dom.Find(selector).Each(func(i int, selection *goquery.Selection) {
		str = append(str, selection.Text())
		// str = selection.Text()
	})
	return str, nil
}

func main() {
	var s = `<tbody role="rowgroup"><tr role="row"><td role="cell">depositcontract</td><td role="cell">contract_name</td><td role="cell">dune</td><td role="cell">ethereum_mainn
et_contracts</td><td role="cell"><time datetime="Tue, 23 Mar 2021 14:45:14 GMT" title="Tue, 23 Mar 2021 14:45:14 GMT">11 months ago</time></td></tr><tr role="row"><td r
ole="cell">eth2</td><td role="cell">project</td><td role="cell">dune</td><td role="cell">ethereum_mainnet_contracts</td><td role="cell"><time datetime="Fri, 05 Mar 2021
 09:40:08 GMT" title="Fri, 05 Mar 2021 09:40:08 GMT">12 months ago</time></td></tr><tr role="row"><td role="cell">uniswap user</td><td role="cell">dapp usage</td><td ro
le="cell">hagaetc</td><td role="cell">ethereum_mainnet_dex_traders</td><td role="cell"><time datetime="Mon, 07 Dec 2020 10:06:37 GMT" title="Mon, 07 Dec 2020 10:06:37 G
MT">1 year ago</time></td></tr><tr role="row"><td role="cell">dex trader</td><td role="cell">activity</td><td role="cell">hagaetc</td><td role="cell">ethereum_mainnet_a
ctivity_type</td><td role="cell"><time datetime="Mon, 07 Dec 2020 10:06:37 GMT" title="Mon, 07 Dec 2020 10:06:37 GMT">1 year ago</time></td></tr></tbody>`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(s))
	if err != nil {
		log.Println(err)
		panic(err)
	}

	var bf bytes.Buffer
	dom.Find("td").Each(func(i int, selection *goquery.Selection) {
		fmt.Println("==>", selection.Text())
		bf.WriteString(strings.TrimSpace(selection.Text()))
		bf.WriteString(";")

		if (i+1)%5 == 0 {
			attr, exists := selection.Find("time").Attr("title")
			if exists {
				attr = strings.ReplaceAll(attr, "\n", "")
				bf.WriteString(strings.TrimSpace(attr))
				bf.WriteString("##")
			}
		}
	})

	result := bf.String()
	if strings.HasSuffix(result, "##") {
		result = result[:strings.LastIndex(result, "##")]
	}
	fmt.Println("==============>", result)

	var f *os.File
	// 	保存文件
	fileName := "E:/20.06.16Project/GOExample/src/A-20220114/gospider/"
	fileName = fileName + "test.txt"
	if checkFileIsExist(fileName) { // 如果文件存在
		f, err = os.OpenFile(fileName, os.O_APPEND, 0666) // 打开文件
		fmt.Println("文件存在")
	} else {
		f, err = os.Create(fileName) // 创建文件
		fmt.Println("文件不存在")
	}
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f) // 创建新的 Writer 对象
	n, err := w.WriteString(result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("写入 %d 个字节n", n)

	w.Flush()
	f.Close()

}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
