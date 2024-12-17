package usecase

import "github.com/sakuyacatcat/over-engineering-fizzbuzz/internal/formatter"

// FizzBuzzUseCaseはFizzBuzzのユースケースを表すインターフェース
type FizzBuzzUseCase interface {
    Execute(start, end int) ([]string, error)
}

type fizzBuzzUseCase struct {
    formatter formatter.Formatter
}

func NewFizzBuzzUseCase(domainService interface{}, f formatter.Formatter) FizzBuzzUseCase {
    return &fizzBuzzUseCase{
        formatter: f,
    }
}

func (uc *fizzBuzzUseCase) Execute(start, end int) ([]string, error) {
    var results []string
    for i := start; i <= end; i++ {
        results = append(results, uc.formatter.Format(i))
    }
    return results, nil
}
