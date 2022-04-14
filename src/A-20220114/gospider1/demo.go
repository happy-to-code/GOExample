package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"time"
)

var chromeDefaltOption = [...]chromedp.ExecAllocatorOption{
	chromedp.NoFirstRun,
	chromedp.NoDefaultBrowserCheck,
	// After Puppeteer's default behavior.
	chromedp.Flag("disable-background-networking", true),
	chromedp.Flag("enable-features", "NetworkService,NetworkServiceInProcess"),
	chromedp.Flag("disable-background-timer-throttling", true),
	chromedp.Flag("disable-backgrounding-occluded-windows", true),
	chromedp.Flag("disable-breakpad", true),
	chromedp.Flag("disable-client-side-phishing-detection", true),
	chromedp.Flag("disable-default-apps", true),
	chromedp.Flag("disable-dev-shm-usage", true),
	chromedp.Flag("disable-extensions", true),
	chromedp.Flag("disable-features", "site-per-process,Translate,BlinkGenPropertyTrees"),
	chromedp.Flag("disable-hang-monitor", true),
	chromedp.Flag("disable-ipc-flooding-protection", true),
	chromedp.Flag("disable-popup-blocking", true),
	chromedp.Flag("disable-prompt-on-repost", true),
	chromedp.Flag("disable-renderer-backgrounding", true),
	chromedp.Flag("disable-sync", true),
	chromedp.Flag("force-color-profile", "srgb"),
	chromedp.Flag("metrics-recording-only", true),
	chromedp.Flag("safebrowsing-disable-auto-update", true),
	chromedp.Flag("enable-automation", true),
	chromedp.Flag("password-store", "basic"),
	chromedp.Flag("use-mock-keychain", true),
}

const storePath = "E:/20.06.16Project/GOExample/src/A-20220114/gospider1/"

func main() {
	addressList := []string{
		"0x00000000219ab540356cbb839cbe05303d7705fa",
		// "0x00000000219ab540356cbb839cbe05303d7705fb",
	}
	for _, addr := range addressList {
		fmt.Println("addr:", addr)
		allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), chromeDefaltOption[:]...)
		defer cancel()

		ctx, cancel := chromedp.NewContext(allocCtx)
		defer cancel()
		ch := addNewTabListener(ctx)

		err := chromedp.Run(ctx,
			chromedp.Sleep(time.Second*1),
			chromedp.Navigate("https://dune.xyz/labels"),
			chromedp.WaitVisible("#search", chromedp.ByID),
			// chromedp.SetValue("#search", addr, chromedp.ByID),
			chromedp.SendKeys("#search", addr, chromedp.ByID),
			chromedp.Click("form.search_form__vdpkL button", chromedp.BySearch),
		// chromedp.WaitVisible("table.table_table__fuS_N > tbody", chromedp.BySearch),
		)
		if err != nil {
			log.Println("Run err : %v\n", err)
		}
		time.Sleep(time.Second * 3)
		newCtx, cancel := chromedp.NewContext(ctx, chromedp.WithTargetID(<-ch))
		defer cancel()

		var html string
		err = chromedp.Run(newCtx,
			chromedp.OuterHTML("#__next > div > main > div > section > div.cell_cell__EYJT6.cell_scrollable__KigLT > table", &html, chromedp.ByQuery),
		)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("html:", html)
		fmt.Println("===================================================================================================")
		dealAndStoreMsg(html, addr)
	}
}

func dealAndStoreMsg(html string, addr string) {
	if strings.TrimSpace(html) != "" {
		err, parsedStr := parseDuneHtml(html)
		if err != nil {
			fmt.Println("dealAndStoreMsg3")
			log.Error("解析html数据失败:%v", err)
			return
		}
		if strings.TrimSpace(parsedStr) != "" {
			nanosecond := time.Now().Nanosecond()
			sprintf := fmt.Sprintf("_%d", nanosecond)
			saveFile(storePath+addr+sprintf+"_dune.txt", parsedStr)
		}
	}
}

// 解析dune网站的HTML
func parseDuneHtml(htmlStr string) (error, string) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlStr))
	if err != nil {
		log.Error(err)
		return err, ""
	}

	var bf bytes.Buffer
	dom.Find("td").Each(func(i int, selection *goquery.Selection) {
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
	return nil, result
}

// 保存文件
func saveFile(fileName, file string) error {
	var f *os.File
	var err error

	if checkFileIsExist(fileName) { // 如果文件存在
		f, err = os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0666) // 打开文件
		log.Warningf("[%s]文件存在", fileName)
	} else {
		f, err = os.Create(fileName) // 创建文件
	}
	defer f.Close()

	if err != nil {
		log.Errorf("创建文件失败,err:%v", err)
		return err
	}

	w := bufio.NewWriter(f) // 创建新的 Writer 对象
	_, err = w.WriteString(file)
	if err != nil {
		log.Errorf("文件写入失败,err:%v", err)
		return err
	}
	w.Flush()

	return nil
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

/**
 * 注册新tab标签的监听服务
 */
func addNewTabListener(ctx context.Context) <-chan target.ID {
	mux := http.NewServeMux()
	ts := httptest.NewServer(mux)
	defer ts.Close()

	return chromedp.WaitNewTarget(ctx, func(info *target.Info) bool {
		return info.URL != ""
	})
}
