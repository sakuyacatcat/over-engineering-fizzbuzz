package provider

// Providerは外部サービスや設定ファイル、DB接続などを提供できるインターフェース。
// 今回はダミー実装。
type Provider interface {
    GetConfig(key string) string
}

type DefaultProvider struct{}

func NewDefaultProvider() *DefaultProvider {
    return &DefaultProvider{}
}

func (d *DefaultProvider) GetConfig(key string) string {
    // 今回は常に空文字、将来的にはファイルや環境変数から取得可能
    return ""
}
