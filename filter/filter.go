package filter

import (
	"fmt"
	"github.com/tanakrid/accounting/transaction"
)

func ByDate(start string, end string, records map[string]transaction.Record) map[string]transaction.Record {
	fmt.Println("filter transaction's data by date")
	result := map[string]transaction.Record{}
	for k, val := range records {
		if val.Date <= end && val.Date >= start {
			result[k] = val
		}
	}
	return result
}

func ByType(typeName string, records map[string]transaction.Record) map[string]transaction.Record{
	fmt.Println("filter transaction's data by type")
	result := map[string]transaction.Record{}
	for k, val := range records {
		if val.TypeName == typeName {
			result[k] = val
		}
	}
	return result
}