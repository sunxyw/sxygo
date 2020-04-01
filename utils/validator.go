/*
 * Package utils
 * File: validator.go
 * Project: SXYGo
 * File Created: 2020-04-01 12:9:32
 * Author: sunxyw <xy2496419818@gmail.com>
 * -----
 * Last Modified: 2020-04-01 12:13:03
 * Modified By: sunxyw <xy2496419818@gmail.com>
 */

package utils

import (
	gvalidator "github.com/go-playground/validator/v10"
)

var validator *gvalidator.Validate

// Validator 返回验证器对象
func Validator() *gvalidator.Validate {
	if validator == nil {
		validator = gvalidator.New()
	}

	return validator
}
