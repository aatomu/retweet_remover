# tweet remover
Golangで自分のために作成  
開発環境:rpi4(8gb) golang(go1.17.8 linux/arm64)  

## -設定-  
同ディレクトリに以下のことを準備/設定する  
twitterAPIKeys:  
```
{
  "accessToken": "<Your TwitterAPI key>",
  "accessSecret": "<Your TwitterAPI key>",
  "APIKey": "<Your TwitterAPI key>",
  "APISecret": "<Your TwitterAPI key>",
  "Token":"<Your TwitterAPI key>"
}
```
tweet.js:  
twitterにて "設定とプライバシー" => "データのアーカイブをダウンロード" => "データのアーカイブをリクエスト"  
してzipをDL後 tweet.jsの最初の`window.YTD.tweet.part0 = `を削除して保存  

## -起動-  
```go run main.go```
## コード元:  
API Code   : https://github.com/ChimeraCoder/anaconda  
Language   : https://golang.org/  