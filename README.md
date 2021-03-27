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

EPGStationの`config/config.yml` にコマンドをセットする．

```yaml
# 録画予約の新規追加時に実行されるコマンド
reserveNewAddtionCommand: "/path/to/epgstation-slack-notification reserve-new-addition"
# 録画情報の更新時に実行されるコマンド
reserveUpdateCommand: "/path/to/epgstation-slack-notification reserve-update"
# 録画予約の削除時に実行されるコマンド
reservedeletedCommand: "/path/to/epgstation-slack-notification reserve-deleted"
# 録画準備の開始時に実行されるコマンド
recordingPreStartCommand: "/path/to/epgstation-slack-notification recording-pre-start"
# 録画準備の失敗時に実行されるコマンド
recordingPrepRecFailedCommand: "/path/to/epgstation-slack-notification recording-prep-rec-failed"
# 録画開始時に実行するコマンド
recordingStartCommand: "/path/to/epgstation-slack-notification recording-start"
# 録画終了時に実行するコマンド
recordingFinishCommand: "/path/to/epgstation-slack-notification recording-finish"
# 録画中のエラー発生時に実行するコマンド
recordingFailedCommand: "/path/to/epgstation-slack-notification recording-failed"
```

EPGStation v1を使っている場合は`config/config.json`にコマンドをセットする．

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
