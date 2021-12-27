<!--
Title: Dell Tech Direct Client API Go
Author: vnikolin
-->

# DellTDAPI

This Library supports fetching Dell Tech Direct API Warranty Information

## Usage

```go
//main.go
package main

import (
	"fmt"
	"github.com/vnikolin/delltdapi"
)

func main() {
	
	client, err := delltdapi.NewDellTDClient("dell-url-fqdn", "client-id", "client-secret", "client-token")
	//Example below (client-token expires every 3600 and is an optional parameter):
	//client, err := delltdapi.NewDellTDClient("apigtwb2c.us.dell.com", "123456", "654321", "")
	
	dellReturn, err := client.FetchWarrantyInfo("hostname", "asset-tag")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dellReturn)
	}
}
```
