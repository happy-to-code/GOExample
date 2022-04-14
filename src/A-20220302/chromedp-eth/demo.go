package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	// options := []chromedp.ExecAllocatorOption{
	// 	chromedp.Flag("headless", true), // debug使用
	// 	chromedp.Flag("blink-settings", "imagesEnabled=false"),
	// 	chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	// }
	// 初始化参数，先传一个空的数据
	// options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background())
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	ch := addNewTabListener(ctx)

	// #content > div.container.space-bottom-2 > div > div > div.row.mb-3
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://etherscan.io/labelcloud"),
		chromedp.WaitVisible("#content > div.container.space-bottom-2 > div > div > div.row.mb-3", chromedp.BySearch),
		// chromedp.Navigate("https://etherscan.io"),
		// chromedp.WaitVisible("#txtSearchInput", chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}

	newCtx, cancel := chromedp.NewContext(ctx, chromedp.WithTargetID(<-ch))
	defer cancel()

	var res string
	err = chromedp.Run(newCtx,
		chromedp.OuterHTML(`div[class="row mb-3"]`, &res, chromedp.BySearch),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 输出了新打开标签页的html
	fmt.Println("--->", res)
	fmt.Println("==================================================")
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
