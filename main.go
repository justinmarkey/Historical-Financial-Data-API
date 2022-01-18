package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guptarohit/asciigraph"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"github.com/piquette/finance-go/equity"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}
}
func graphrequestHandler (c *gin.Context) {
	ticker := c.Query("ticker")
	todaymonth := time.Now().Month()
	todayday := time.Now().Day()
	todayyear := time.Now().Year()

	chartparams := &chart.Params {
		Symbol: ticker,
  		Start: &datetime.Datetime{Month: int(todaymonth), Day: todayday, Year: todayyear-1},
 		End: &datetime.Datetime{Month: int(todaymonth), Day: todayday, Year: todayyear},
  		Interval: datetime.OneDay,
	}
	
	genchart := chart.Get(chartparams)

	data := []float64{}
	for genchart.Next() {
		bar := genchart.Bar()
		float, _ := bar.Close.Round(2).Float64()
		data = append(data, float)
	}
	checkErr(genchart.Err())
	
	graph := asciigraph.Plot(data, asciigraph.Height(20))

	
	c.String(200,"%v",graph)
}

func equityrequestHandler (c *gin.Context) {
	tickers := c.QueryArray("tickers")
	response:= equity.List(tickers)
	
	for response.Next() {
		data := response.Equity()
		
		c.IndentedJSON(200, data)
	}
	checkErr(response.Err())
}

func main() {

	router := gin.Default()
	router.GET("/equity", equityrequestHandler)
	router.GET("/graph", graphrequestHandler)
	router.Run()
}