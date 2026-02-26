package validation

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

func InitTrans() {
	// 注册中文翻译器
	zhT := zh.New()
	uni := ut.New(zhT, zhT)
	Trans, _ = uni.GetTranslator("zh")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// --- 新增核心代码：注册自定义标签获取函数 ---
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			// 获取 json 标签的值
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "" {
				name = fld.Tag.Get("form")
			}
			// 如果 json 标签是 "-" (忽略字段)，则返回原字段名
			if name == "-" {
				return ""
			}
			return name
		})
		_ = zhTranslations.RegisterDefaultTranslations(v, Trans)
	}
}

// Translate 提取并翻译错误信息
func Translate(err error) string {
	var errs validator.ValidationErrors
	// 判断错误是否为验证错误（类型断言）
	if errors.As(err, &errs) {
		// 翻译错误，并只提取出第一个字段的错误信息返回，这样前端更好展示
		for _, e := range errs {
			return e.Translate(Trans)
		}
	}
	// 如果不是验证错误（例如 JSON 格式传错了），原样返回
	return err.Error()
}

func BindJSON[T any](c *gin.Context) (*T, error) {
	var req T
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func BindQuery[T any](c *gin.Context) (*T, error) {
	var req T
	if err := c.ShouldBindQuery(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
