package formatter

import "fmt"

// Strategyインターフェース(拡張性確保)
type Strategy interface {
    Apply(number int) (string, bool)
}

// デフォルトのFizzBuzz戦略
type DefaultStrategy struct{}

func NewDefaultStrategy() *DefaultStrategy {
    return &DefaultStrategy{}
}

func (d *DefaultStrategy) Apply(number int) (string, bool) {
    if number%15 == 0 {
        return "FizzBuzz", true
    } else if number%3 == 0 {
        return "Fizz", true
    } else if number%5 == 0 {
        return "Buzz", true
    }
    return "", false
}

// ChainはStrategyを複数繋げられるようにするラッパ
type Chain struct {
    strategies []Strategy
}

func NewChain(strategies ...Strategy) Chain {
    return Chain{strategies: strategies}
}

func (c Chain) Execute(number int) string {
    for _, s := range c.strategies {
        if res, ok := s.Apply(number); ok {
            return res
        }
    }
    // FizzBuzzに該当しなければ数字を文字列化
    return defaultNumberToString(number)
}

func defaultNumberToString(n int) string {
    return fmtInt(n)
}

// 数値から文字列への変換関数を抽象化(将来的なロケール対応を想定)
func fmtInt(n int) string {
    return fmt.Sprintf("%d", n)
}
