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

func main() {
	names := []string{"周杰伦", "张学友"}
	for _, name := range names {
		allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), chromeDefaltOption[:]...)
		defer cancel()

		ctx, cancel := chromedp.NewContext(allocCtx)
		defer cancel()
		ch := addNewTabListener(ctx)

		err := chromedp.Run(ctx,
			chromedp.Navigate("https://music.migu.cn/v3/music/player/audio"),
			chromedp.SetValue("#search_ipt", name, chromedp.ByID),
			chromedp.Click(`i[class="iconfont cf-nav-sousuo"]`, chromedp.BySearch),
		)
		if err != nil {
			log.Fatal(err)
		}

		newCtx, cancel := chromedp.NewContext(ctx, chromedp.WithTargetID(<-ch))
		defer cancel()

		var res string
		err = chromedp.Run(newCtx,
			chromedp.OuterHTML(`div[class="songlist-body"]`, &res, chromedp.BySearch),
		)
		if err != nil {
			log.Fatal(err)
		}

		// 输出了新打开标签页的html
		fmt.Println("--->", res)
		fmt.Println("==================================================")
	}
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
