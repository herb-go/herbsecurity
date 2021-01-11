package roleparser

import (
	"net/url"
)

//Escape escape Reserved Characters wirh urlencode
// https://tools.ietf.org/html/rfc3986#section-2.2
func Escape(unescaped string) string {
	return url.QueryEscape(unescaped)
}

//Unescape  unescape Reserved Characters wirh urlencode
//https://tools.ietf.org/html/rfc3986#section-2.2
func Unescape(escaped string) (string, error) {
	return url.QueryUnescape(escaped)
}

//TokenRoleSep token separate roles
const TokenRoleSep = ";"

//TokenAttributesStart token marks attributes start
const TokenAttributesStart = ":"

//TokenAttributeSep token separate attribites
const TokenAttributeSep = ","

//TokenAttributeVauleStart token marks attribute vaues start
const TokenAttributeVauleStart = "="
