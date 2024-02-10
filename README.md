# Historical Finanacial Data API

ABOUT:
This program serves to allow other interenal programs a way to request the historical data on financial securites. This works by spinning up the gin engine and making requests to the gin local server. This script takes requests from the gin server and sends them to the Yahoo Finance API (implemented under the library "finance-go"). The Yahoo API will return the closing data for the past year in JSON format. 

USE:
to request equity data:
endpoint: /equity 
   
Example request =  0.0.0.0:8080/equity?tickers=AAPL,GOOGL,D
    
    notes: this parses query and will return data for every ticker in json format

to request a graph of an equity:
endpoint: /graph
    
Example graph request: 0.0.0.0:8080/graph?ticker=AAPL ... will chart AAPLs close for past year
    
    notes: the chart has a height of 60 for legibility.  Timeframe = Length: 1 year from today, Interval: 1 day.
    
