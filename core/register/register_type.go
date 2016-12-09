package register

import (
	"log"

	"github.com/fudali113/doob/utils/reflect"
)

/**
 * 函数类别
 */
const (
	PARAM_NONE = iota
	CTX
	ORIGIN
	CI_PATHVARIABLE
	CI_PATHVARIABLE_CTX
)

const (
	RETURN_NONE = iota
	FILE
	JSON
	RETURN_TYPE
)

type RegisterType struct {
	Handler    interface{}
	ParamType  *ParamType
	ReturnType *ReturnType
}

type ParamType struct {
	Type  int
	CiLen int
}

type ReturnType struct {
	Type int
}

func GetFuncRegisterType(function interface{}) *RegisterType {
	paramType, returnType := GetFuncParam3ReturnType(function)
	if paramType.Type == ORIGIN && returnType.Type == RETURN_NONE {
		log.Panic("支持原始http函数只为兼容，不允许设置返回值")
	}
	return &RegisterType{
		Handler:    function,
		ParamType:  paramType,
		ReturnType: returnType,
	}
}

func GetFuncParam3ReturnType(function interface{}) (*ParamType, *ReturnType) {
	params, returns := reflect.GetFuncParams(function)
	return getParamType(params), getReturnType(returns)
}

/**
 * 获取参数类型
 */
func getParamType(params []string) *ParamType {
	stringTypeLen := 0
	hasCTX := 0
	hasOringin := 0
	for _, param := range params {
		log.Print(param)
		switch param {
		case "string":
			stringTypeLen++
			if hasOringin > 0 || hasCTX > 0 {
				log.Panic("自动注入url参数必须放在参数最前面")
			}
		case "*core.Context":
			hasCTX++
		case "*http.Request":
			hasOringin++
		case "http.ResponseWriter":
			hasOringin++
		default:
		}
	}
	if stringTypeLen > 0 {
		if hasCTX > 0 {
			return &ParamType{
				Type:  CI_PATHVARIABLE_CTX,
				CiLen: stringTypeLen,
			}
		}
		if hasOringin > 0 {
			log.Panic("您的函数doob不支持")
		}
		return &ParamType{
			Type:  CI_PATHVARIABLE,
			CiLen: stringTypeLen,
		}
	}
	if hasCTX > 0 {
		return &ParamType{
			Type: CTX,
		}
	}
	if hasOringin > 0 {
		return &ParamType{
			Type: ORIGIN,
		}
	}
	return &ParamType{
		Type: PARAM_NONE,
	}
}

/**
 * 获取返回值类型
 */
func getReturnType(returns []string) *ReturnType {
	Type := func(returns []string) int {
		switch len(returns) {
		case 1:
			switch returns[0] {
			case "string":
				return FILE
			default:
				return JSON

			}
		case 2:
			switch returns[0] {
			case "string":
				return RETURN_TYPE
			default:
				log.Panic("您的函数doob不支持")
			}
		default:
			return RETURN_NONE
		}
		log.Panic("您的函数doob不支持")
		return RETURN_NONE
	}(returns)
	return &ReturnType{Type: Type}
}
