package app

import (
	"github.com/sakuyacatcat/over-engineering-fizzbuzz/internal/provider"
	"github.com/sakuyacatcat/over-engineering-fizzbuzz/pkg/usecase"
)

type App struct {
    useCase  usecase.FizzBuzzUseCase
    provider provider.Provider
}

func NewApp(uc usecase.FizzBuzzUseCase, p provider.Provider) *App {
    return &App{
        useCase:  uc,
        provider: p,
    }
}

func (a *App) Run(start, end int) ([]string, error) {
    return a.useCase.Execute(start, end)
}
