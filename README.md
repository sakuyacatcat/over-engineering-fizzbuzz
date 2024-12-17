# FizzBuzz Overengineered

本プロジェクトは、単純な FizzBuzz 問題を Go 言語でオーバーエンジニアリング気味に実装したサンプルコードです。
アプリケーションを階層化し、依存性注入、DI コンテナ、ストラテジーパターン、ユースケース層とドメイン層の分離などを行っています。
本来シンプルに書ける FizzBuzz を、拡張性や変更容易性、テスト性を過剰に考慮したアーキテクチャを示すための実装例です。

## 構成

```
fizzbuzz/
├─ cmd/
│  └─ fizzbuzz/
│     ├─ main.go                // エントリポイント
│     └─ integration_test.go    // Integrationテスト例
├─ internal/
│  ├─ app/
│  │  └─ app.go                 // アプリケーション層: UseCaseを呼び出す
│  ├─ fizzbuzz/
│  │  ├─ fizzbuzz.go            // 将来的な拡張用。現在はラッパ的存在
│  ├─ formatter/
│  │  ├─ formatter.go           // Formatterインターフェース実装とCompositeクラス
│  │  ├─ formatter_test.go
│  │  └─ strategy.go            // StrategyパターンによるFizzBuzz文字列化ロジック
│  ├─ provider/
│  │  └─ provider.go            // 設定や外部サービス接続を抽象化
│  └─ di/
│     └─ di.go                  // 依存性注入(コンストラクションルート)
├─ pkg/
│  ├─ domain/
│  │  └─ domain.go              // ドメインサービス層
│  └─ usecase/
│     └─ usecase.go             // UseCase層(FizzBuzz出力)
│
└─ Makefile
```

### 各レイヤの役割

- **cmd 層**:
  アプリケーションのエントリポイント。`main.go`から DI コンテナを通してアプリケーションを実行します。
  また、統合テストを行う`integration_test.go`もここに配置。

- **internal/app 層**:
  ユースケースを用いてアプリケーションの制御を行う層。
  `app.App`が`usecase`を呼び出して、結果を返します。

- **internal/fizzbuzz 層**:
  今回は大きなロジックは無いが、将来的な拡張（別の FizzBuzz バリエーション等）を想定して用意した層。

- **internal/formatter 層**:
  数値を"Fizz", "Buzz", "FizzBuzz", あるいはそのまま文字列化する戦略を実装しています。
  `Strategy`パターンでロジックを分離し、`CompositeFormatter`や`Chain`で柔軟な拡張を可能にしています。

- **internal/provider 層**:
  設定や外部サービス接続などを抽象化するための層。
  現状はダミー実装ですが、将来的に設定ファイル読み込みや環境変数管理を行うことを想定しています。

- **internal/di 層**:
  依存性注入コンテナ的な実装が含まれています。`InitializeApp()`で全ての依存を組み立て、`app.App`オブジェクトを返します。

- **pkg/domain 層**:
  ビジネスロジック(ドメインロジック)を扱う層。
  今回は FizzBuzz の単純なロジックしかないため、ほぼダミーですが、将来の複雑なビジネスルール拡張を想定しています。

- **pkg/usecase 層**:
  ドメインロジック、フォーマッタ、その他依存を利用し、ビジネス的な操作(今回であれば指定範囲の FizzBuzz 生成)を行う層。
  `Execute()`メソッドで FizzBuzz 結果を生成します。

## 開発・ビルド手順

1. Go 1.20 以上がインストールされていることを確認してください。
2. プロジェクトルートで以下を実行します。

```sh
make build
```

`build/fizzbuzz` という実行ファイルが生成されます。

## 実行

```sh
make run
```

`1`から`100`までの FizzBuzz 結果が標準出力に表示されます。

## テスト

```sh
make test
```

`go test ./...` により、全てのテストが実行されます。

## オーバーエンジニアリングポイント

- レイヤードアーキテクチャ: `cmd`, `internal`, `pkg`と役割ごとに階層を分離
- 依存性注入(DI): `internal/di`で必要なオブジェクトを集約
- Strategy パターンを用いたフォーマットロジック分割
- provider 層による外部依存抽象化
- 拡張性・テスト容易性の確保(Integration Test, Unit Test)

実用プロジェクトでは過剰な分割ですが、アーキテクチャのサンプルとして参考にしてください。
