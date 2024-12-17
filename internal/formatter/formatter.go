package formatter

type Formatter interface {
    Format(number int) string
}

type CompositeFormatter struct {
    chain Chain
}

func NewCompositeFormatter(chain Chain) *CompositeFormatter {
    return &CompositeFormatter{chain: chain}
}

func (c *CompositeFormatter) Format(number int) string {
    return c.chain.Execute(number)
}
