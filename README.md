# stock price checker

## Go Binary
#### Building
```
$ go get github.com/gorilla/mux
$ go get github.com/savaki/jq
$ go build
```

#### Running
```
$ source .env
$ ./stock_price_checker
```

## Dockerfile

Note: currently produces an image thats 784MB... hah.

#### Building
`docker build -t stock_price_checker:0.1 . -f Dockerfile`

#### Running
`docker run -it -p 8000:8000 stock_price_checker:0.1`

# Usage
```
curl 0.0.0.0:8000
```

## task

```
xxxxxxxxx DevOps Cloud Challenge

This is a simple exercise that should take you no more than 3-4 hours to complete. This is intended to be fun. If you enjoy this exercise, you will like the Cloud Engineering role at xxxxxxxxx. 


Write a web service that looks up the closing price of a stock. You can use any language you like (slight preference for Go, python, Java, etc. - but if you are more comfortable in another language that is fine). 

In response to a GET request, the service should return the last NDAYS days of data along with the average closing price over those  days. For example, for MSFT and NDAYS=3, the GET response might look something like this:

MSFT  data=[110.56, 111.25, 115.78], average=112.50 

The stock SYMBOL (the symbol to look up) and NDAYS (the number of days) are environment variables provided to your program. 

Use this free quote service:

Sample query:

https://www.alphavantage.co/query?apikey=demo&function=TIME_SERIES_DAILY_ADJUSTED&symbol=MSFT 

Note Use the apikey="C227WD9W3LUVKVV9" instead of demo, or create your own apikey.

No error handling is required. You can assume the stock exists, the number of days is valid, etc..  In other words, this can be a quick and dirty program - it does not need to be polished.

Create a docker image that runs your web service. Publish your docker image, your code,  and provide instructions on how to build the image and run it.


Part 2: Optional, time permitting


Create a Kubernetes manifest that deploys your web service, creates a service for it, and exposes it as an ingress.

Use a configmap to pass in all environment variables. For the exercise use SYMBOL=MSFT and NDAYS=7

Use a secret to pass in the api key APIKEY=C227WD9W3LUVKVV9

Publish your manifests, etc. to git (github, bitbucket, gitlab, etc) , and send us the link along with instructions on how to deploy it. The sample provided should run on a vanilla Kubernetes environment (minikube, for example).

If you have any questions feel free to reach out to us for assistance. 
```
