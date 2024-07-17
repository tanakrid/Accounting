package report

import (
	"fmt"
	"github.com/tanakrid/accounting/transaction"
)

func Show()  {
	fmt.Println("show report")
}

func Sum(records map[string]transaction.Record) float64 {
	var sum float64
	for _, val := range records {
		sum += val.Amount
	}
	return sum
}

func ShowPieChart(records map[string]transaction.Record) map[string]float64 {
	pieChart := map[string]float64{}
	amountSum := Sum(records)
	for _, val := range records {
		pieChart[val.TypeName] = pieChart[val.TypeName] + val.Amount
	}

	for k, val := range pieChart {
		pieChart[k] = val / amountSum * 100
	}
	return pieChart
}