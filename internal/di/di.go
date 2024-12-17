package di

import (
	"github.com/sakuyacatcat/over-engineering-fizzbuzz/internal/app"
	"github.com/sakuyacatcat/over-engineering-fizzbuzz/internal/formatter"
	"github.com/sakuyacatcat/over-engineering-fizzbuzz/internal/provider"
	"github.com/sakuyacatcat/over-engineering-fizzbuzz/pkg/domain"
	"github.com/sakuyacatcat/over-engineering-fizzbuzz/pkg/usecase"
)

func InitializeApp() (*app.App, error) {
    // Provider層で設定可能なインフラを提供できる想定(今回はダミー)
    p := provider.NewDefaultProvider()

    // StrategyパターンでFormatterを差し替え可能に
    // ここで複雑にインターフェースを差し込めるが、今回は1種類だけ
    f := formatter.NewCompositeFormatter(
        formatter.NewChain(
            formatter.NewDefaultStrategy(),
        ),
    )

    // UseCaseにドメインロジックとフォーマッタを注入
    uc := usecase.NewFizzBuzzUseCase(
        domain.NewFizzBuzzDomainService(),
        f,
    )

    // AppへUseCaseを注入
    return app.NewApp(uc, p), nil
}
