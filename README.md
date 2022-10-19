# api-temp-go-gin

API template by golang and gin

## Use

step1. setting app config

```shell
cp config/app.conf.example.yaml app.conf.yaml
```

step2. download depend library

```shell
go mod download
```

step3. run app

```shell
go run cmd/app/main.go -c [path to app.conf.yaml]
```

## config & environment

|參數|說明|
|----|---|
| PROJECT | 服務名稱，預設`ambipom` |
| APP_ADDR | 服務運行port，預設`8999` |
| APP_MODE | 服務運行模式(debug, release, test)，預設`debug` |
| APP_READ_TIMEOUT | http 服務讀取時timeout設定，預設`10s` |
| APP_WRITE_TIMEOUT | http 服務寫入時timeout設定，預設`10s` |
| DB_DRIVER | 資料庫類型，預設`mysql` |
| DB_DATABASE | 資料庫名稱， 預設`platform` |
| DB_MASTER_USERNAME | 寫入用DB主機位址 |
| DB_MASTER_PASSWORD | 寫入用DB帳號 |
| DB_MASTER_ADDRESS | 寫入用DB主機位址 |
| DB_SLAVE_USERNAME | 讀取用DB帳號 |
| DB_SLAVE_PASSWORD | 讀取用DB密碼 |
| DB_SLAVE_ADDRESS | 讀取用DB主機位址 |
| DB_DIAL_TIMEOUT | DB連線超時時間設定，預設`10s` |
| DB_READ_TIMEOUT | DB讀取時超時時間設定，預設`30s` |
| DB_WRITE_TIMEOUT | DB寫入時超時時間設定，預設`60s` |
| DB_DB_TIMEZONE | DB運行時時區，預設`UTC` |
| DB_APP_TIMEZONE | DB連線時時區，預設`UTC` |
| DB_CONN_MAX_LIFE_TIME | DB連線最大時間，預設`0s` |
| DB_MAX_IDLE_CONNS | 最大空閑連接數，預設`2` |
| DB_MAX_OPEN_CONNS | 最大連接數，預設`0`|
| REDIS_ADDRESS | redis 位址 |
| REDIS_PASSWORD | redis 密碼 |
| REDIS_DB | redis 開始索引位址，預設`0` |

> 可以透過app.cong.yaml進行設定，也可以透過系統環境變數進行設定

