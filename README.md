# FinAPI

/equity endpoint for requesting data
    usage: 0.0.0.0:8080/equity?tickers=AAPL,GOOGL,D
    
    notes: this parses query and will return data for every ticker in json format
    
/graph end point for graphing a ticker:

    Timeframe = Length: 1 year from today, Interval: 1 day.
    
    usage: 0.0.0.0:8080/graph?ticker=AAPL ... will chart AAPLs close for past year
    
    notes: the chart has a height of 60 for legibility.
    
