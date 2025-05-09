/**
 * @file xmlvalidate.go
 * @author Krisna Pranav
 * @version 1.0
 * @date 2025-05-09
 *
 * @copyright Copyright (c) 2025 Krisna Pranav
 *
 */

package utils

import (
	"fmt"

	xsdvalidate "github.com/terminalstatic/go-xsd-validate"
)

var xsdhandler *xsdvalidate.XsdHandler

func ValidateXML(inXml []byte) (bool, string) {
	xmlhandler, err := xsdvalidate.NewXmlHandlerMem(inXml, xsdvalidate.ParsErrDefault)
	if err != nil {
		panic(err)
	}

	err = xsdhandler.Validate(xmlhandler, xsdvalidate.ValidErrDefault)

	if err != nil {
		switch err.(type) {
		case xsdvalidate.ValidationError:
			fmt.Println(err)
			fmt.Printf("Error in line: %d\n", err.(xsdvalidate.ValidationError).Errors[0].Line)
			fmt.Println(err.(xsdvalidate.ValidationError).Errors[0].Message)
			return false, err.(xsdvalidate.ValidationError).Errors[0].Message
		default:
			fmt.Println(err)
		}
		return false, err.Error()
	}
	return true, ""
}

func init() {
	xsdvalidate.Init()
	var err error
	
	xsdhandler, err = xsdvalidate.NewXsdHandlerUrl("./TinyMLSchema.xsd", xsdvalidate.ParsErrDefault)

	if err != nil {
		panic(err)
	}
}