# epgstation-slack-notification

![GitHub Actions Go CI Status](https://github.com/hiroxto/epgstation-slack-notification/workflows/Go%20CI/badge.svg)

EPGStation の通知を Slack に送るコマンドラインツール

## 設定

### 設定ファイルを作成

`epgstation-slack-config.example.yml` ファイルの中身を `epgstation-slack-config.yml` コピーして Slack の API キーとチャンネル名を書き込む.

`epgstation-slack-config.yml` ファイルは `epgstation-slack-notification` のバイナリと同じ場所に配置する.

```bash
$ wget -O epgstation-slack-config.yml https://raw.githubusercontent.com/hiroxto/epgstation-slack-notification/master/epgstation-slack-config.example.yml
```

### コマンドをセットする

EPGStation の `config.json` にコマンドをセットする.

```json
{
  "reservationAddedCommand": "/path/to/epgstation-slack-notification reservation-added",
  "recordedPreStartCommand": "/path/to/epgstation-slack-notification recorded-pre-start",
  "recordedPrepRecFailedCommand": "/path/to/epgstation-slack-notification recorded-prep-rec-failed",

  "recordedStartCommand": "/path/to/epgstation-slack-notification recorded-start",
  "recordedEndCommand": "/path/to/epgstation-slack-notification recorded-end",
  "recordedFailedCommand": "/path/to/epgstation-slack-notification recorded-failed"
}
```

### EPGStation を再起動

設定をした後, EPGStation を再起動する

```bash
$ pm2 restart epgstation
```

## Licence

[MIT Licence](https://raw.githubusercontent.com/hiroxto/epgstation-slack-notification/master/LICENSE)
