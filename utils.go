package sqlmgo

import "strings"

const(
	FuncName_SUM = "sum"
)

func ConvertOperator(strOperator string) string {
	switch strOperator {
	case "=":
		return "$eq"
	case ">":
		return "$gt"
	case "<":
		return "$lt"
	case ">=":
		return "$gte"
	case "<=":
		return "$lte"
	case "!=":
		return "$ne"
	}
	return strOperator
}

func ConvertFunc(strFuncName string) string {
	if strings.EqualFold(strFuncName, FuncName_SUM) {
		return "$sum"
	}
	return ""
}