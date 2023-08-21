# panforyou-test-1

## アプリの概要

contentful API で 3 つのパンの情報を取得し、Cloud Firebase に保存する CLI です。

## 環境構築

### モジュールのインストール

`go mod download`

### .env ファイルの作成

以下のコマンドを実行して.env ファイルを作成し、各環境変数を設定してください。  
`cp .env.dist .env`

### Firebase でサービスアカウントの秘密鍵の作成

以下の流れで秘密鍵(JSON ファイル)を作成し、ルートディレクトリに `path/to/serviceAccount.json` として保存してください。

↓Firebase コンソール  
↓ プロジェクトの設定  
↓ サービスアカウント  
↓ 新しい秘密鍵を生成ボタン押下  
→JSON ファイルのダウンロード

## アプリの使い方

### CLI コマンドの実行

以下のコマンドに引数となる 1 つの entry_id を指定し、実行します。  
`go run main.go <entry_id>`
