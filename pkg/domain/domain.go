package domain

// domain層：FizzBuzzに関連するビジネスロジックが増えた場合、この層で表現

type FizzBuzzDomainService struct{}

func NewFizzBuzzDomainService() *FizzBuzzDomainService {
    return &FizzBuzzDomainService{}
}

// 今回は特にドメインロジックを外部化せず、そのまま利用可能。
// 将来的に条件変更や別のパターン対応などをここで対処可能。
