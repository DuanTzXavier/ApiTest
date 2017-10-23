package parseAPI

import (
	"net/http"
	"reflect"
	"fmt"
	header2 "ApiTest/ReadingIN/base/communication/apiStructure/header"
)

func ParseHeader(header http.Header) {
	var requestHeader header2.RequestHeaderV1

	object := reflect.ValueOf(&requestHeader)
	fields := object.Elem()
	fieldsOfType := fields.Type()
	for i:=0; i<fields.NumField(); i++{
		fields.FieldByName(fieldsOfType.Field(i).Name).SetString(header.Get(fieldsOfType.Field(i).Name))
		fmt.Print(fieldsOfType.Field(i).Name + " : " + header.Get(fieldsOfType.Field(i).Name) + " \n")

	}
	fmt.Print(header)
	fmt.Print(requestHeader.ClientID)
	fmt.Println("\n")
}
