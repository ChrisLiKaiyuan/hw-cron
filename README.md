# hw-cron ![Go](https://github.com/ChrisLiKaiyuan/hw-cron/workflows/Go/badge.svg?branch=master) ![cron](https://github.com/ChrisLiKaiyuan/hw-cron/workflows/cron/badge.svg?branch=master)
定时推送提醒我及时完成泛雅上的作业。

运行所需环境变量：
```
export HDU_NO=<REDACTED>
export HDU_PASSWORD=<REDACTED>
export ALERT_URL=<REDACTED>
```

## 编译运行
```bash
go mod tidy
go build . && ./hw-cron
```