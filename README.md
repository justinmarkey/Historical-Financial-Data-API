# Historical Financial Data API

## Overview
This project provides a local API for retrieving historical equity price data. It spins up a Gin web server that accepts HTTP requests and proxies them to the Yahoo Finance API via the finance-go library.
The API currently returns daily closing prices for the past year, formatted as JSON, and can also generate simple price charts for individual equities.
This service is designed to be consumed by internal programs or tools that need quick access to historical market data without directly calling external APIs. 

## Features
- Retrieve 1 year of historical closing price data
- Support for multiple equity tickers per request
- JSON-formatted responses
- Simple equity price chart generation
- Lightweight local server using Gin


### Example of an equity request
http://0.0.0.0:8080/equity?tickers=AAPL,GOOGL,D

### Example of a graph request
http://0.0.0.0:8080/graph?ticker=AAPL

## Test in terminal

'''bash
curl "http://0.0.0.0:8080/graph?ticker=AAPL"
'''
