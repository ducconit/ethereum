package util

import (
	"github.com/shopspring/decimal"
)

func EtherToWei(amount interface{}) decimal.Decimal {
	return ToWei(amount, DecimalEtherWei)
}

func WeiToEther(amount interface{}) decimal.Decimal {
	return ToEther(amount, DecimalEtherWei)
}

func ParseToDecimal(iamount interface{}) decimal.Decimal {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromInt(v)
	case int:
		amount = decimal.NewFromInt(int64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}
	return amount
}

func ToWei(iamount interface{}, decimals int64) decimal.Decimal {
	mul := decimal.NewFromInt(10).Pow(decimal.NewFromInt(decimals))
	return ParseToDecimal(iamount).Mul(mul)
}

func ToEther(iamount interface{}, decimals int64) decimal.Decimal {
	mul := decimal.NewFromInt(10).Pow(decimal.NewFromInt(decimals))
	return ParseToDecimal(iamount).Div(mul)
}
