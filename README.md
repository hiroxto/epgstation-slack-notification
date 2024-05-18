# epgstation-slack-notification

![GitHub Actions Test Status](https://github.com/hiroxto/epgstation-slack-notification/workflows/Test/badge.svg)

EPGStation の通知を Slack に送るコマンドラインツール

## ダウンロード

[Releases](https://github.com/hiroxto/epgstation-slack-notification/releases) から OS とアーキテクチャに合ったバイナリをダウンロード。

Linux / macOS の amd64 / arm64 / armv7 のみビルドしているのでそれ以外の環境では自分でビルドする。

## 設定

### Slack API Token の準備

メッセージの投稿に API Token を使うので準備する。
詳しくは [Slack のドキュメント](https://api.slack.com/authentication/token-types)を参照。

作成した API Token は設定ファイルに書き込むので控えておく。

### 投稿するチャンネル ID の確認

投稿するチャンネルをチャンネル ID で指定するので確認して控えておく。

### 設定ファイルを作成

`epgstation-slack-config.example.yml` ファイルを `epgstation-slack-config.yml` へコピーして Slack の API キーとチャンネル名を書き込む。

設定ファイルはデフォルトではバイナリと同じディレクトリに配置した `epgstation-slack-config.yml` ファイルを利用し，オプションで指定された場合はオプションの値が優先される。

```bash
$ wget -O epgstation-slack-config.yml https://raw.githubusercontent.com/hiroxto/epgstation-slack-notification/master/epgstation-slack-config.example.yml
```

デフォルト以外の場所に配置した設定ファイルを利用する場合は `--config`, `-c` オプションを利用する。

```bash
$ ./epgstation-slack-notification --config /path/to/config.yml reserve-new-addition
```

### コマンドをセットする

EPGStation の `config/config.yml` にコマンドをセットする。

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
# エンコード終了時に実行するコマンド
encodingFinishCommand: "/path/to/epgstation-slack-notification encoding-finish"
```

EPGStation v1 を使っている場合は`config/config.json`にコマンドをセットする。

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

設定をした後, EPGStation を再起動する。

```bash
$ pm2 restart epgstation
```

## 通知内容のカスタマイズ

通知設定の `message` と `template` は Go の text/template を使って任意の内容を出力できる。

テンプレートに渡されるデータの構造と中身は `dump:detail` コマンドを使って確認可能。

渡されるデータはコマンド毎に分かれていて以下のように使い分けされる。
- `ReserveDetail` を使うコマンド
  - `reserve-new-addition`
  - `reserve-update`
  - `reserve-deleted`
  - `recording-pre-start`
  - `recording-prep-rec-failed`
- `RecordingDetail` を使うコマンド
  - `recording-start`
  - `recording-finish`
  - `recording-failed`
- `EncodingDetail` を使うコマンド
  - `encoding-finish`

データの構造は基本的に [EPGStation のドキュメント通り](https://github.com/l3tnun/EPGStation/blob/master/doc/conf-manual.md#%E5%A4%96%E9%83%A8%E3%82%B3%E3%83%9E%E3%83%B3%E3%83%89%E5%AE%9F%E8%A1%8C)の値となり，一部独自の値を持っている。
- `StartAtTime`
  - 環境変数 `STARTAT` を `time.Time` 型に変換した値
- `EndAtTime`
  - 環境変数 `ENDAT` を `time.Time` 型に変換した値
- `DurationMin`
  - 環境変数 `DURATION` を分単位の `int64` 型に変換した値

## Licence

[MIT Licence](https://raw.githubusercontent.com/hiroxto/epgstation-slack-notification/master/LICENSE)
