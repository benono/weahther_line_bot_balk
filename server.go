/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"work/weather"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/microsoft/vscode-remote-try-go/hello"
)

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, hello.Hello())
}

func main() {
	// LINE Botクライアント生成する
	// BOT にはチャネルシークレットとチャネルトークンを環境変数から読み込み引数に渡す
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	// エラーに値があればログに出力し終了する
	if err != nil {
		log.Fatal(err)
	}
	// weatherパッケージパッケージから天気情報の文字列をを取得する
	result, err := weather.GetWeather()
	// エラーに値があればログに出力し終了する
	if err != nil {
		log.Fatal(err)
	}
	// テキストメッセージを生成する
	message := linebot.NewTextMessage(result)
	// テキストメッセージを友達登録しているユーザー全員に配信する
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
