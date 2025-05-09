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

import xsdvalidate "github.com/terminalstatic/go-xsd-validate"

var xsdhandler *xsdvalidate.XsdHandler

func ValidateXML(inXml []byte) (bool, string) {
	xmlhander, err := xsdvalidate.NewXmlHandlerMem(inXml, xsdvalidate.ParsErrDefault)
	
	if err != nil {
		panic(err)
	}

	err = xsdhandler.Validate(xmlhander, xsdvalidate.ValidErrDefault)

	if err != nil {

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