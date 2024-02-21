package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var buf []byte

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://map.baidu.com/search/%E6%B2%B3%E4%B8%9C%E5%8C%BA%E5%85%AD%E7%BA%AC%E8%B7%AF%E8%BD%BB%E7%BA%BA%E5%9F%8E/@12978592,4861519,13z?querytype=s&wd=%E6%B2%B3%E4%B8%9C%E5%8C%BA%E5%85%AD%E7%BA%AC%E8%B7%AF%E8%BD%BB%E7%BA%BA%E5%9F%8E&c=131&pn=0&device_ratio=2&da_src=shareurl"), // 替换为你的目标URL
		chromedp.WaitVisible(`body`, chromedp.BySearch), // 等待body元素可见
		chromedp.CaptureScreenshot(&buf),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("screenshot.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
}
