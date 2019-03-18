# stock price checker

Dear reviewer
- I have not implemented NDAYS. The programme will always return the last 100 days of results.
- The dockerfile/image works, but is sickeningly large at 784MB. There are some attempts at a much smaller image on branch `feature/smaller_dockerfile` but results in `standard_init_linux.go:178: exec user process caused "no such file or directory"`. The binary seems to exist in the expected place within the image but I haven't investigated further. The working image can be pulled from andycsoka/stock_price_checker:0.4
- Deploying to MiniKube results in `404 - default backend` from the ingress when hitting the hostname:80

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

Note:
- currently produces an image thats 784MB... hah.
- tested on Ubuntu 18:04 & Docker version 18.09.2
- published to `andycsoka/stock_price_checker:0.4`

#### Building
`docker build -t stock_price_checker:0.4 . -f Dockerfile`

#### Running
`docker run -it -p 8000:8000 --env-file docker.env stock_price_checker:0.4`

## Kubernetes

##### problem: curl hostname:80 gives default backend 404

Note: for minikube you'll need
- `minikube addons enable ingress`
- `minikube ip` and update your hosts file to map the ingress' host

```
kubectl create -f k8s_manifests/configmap.yaml \
  -f k8s_manifests/secrets.yaml
  -f k8s_manifests/service.yaml
  -f k8s_manifests/ingress.yaml
  -f k8s_manifests/deployment.yaml
```

# Usage
### go binary or docker:
```
curl 0.0.0.0:8000
```

### MiniKube:
Assuming you've got minikube up and running with an ingress controller enabled and have deployed the manifests
```
minikube ip
sudo echo "<minikube ip output> stockprice" >> /etc/hosts
curl stockprice:80
and witness it not working :(
```


# Pushing to docker hub (personal notes)
```
docker login --username=andycsoka --email=andrewcsoka@gmail.com
docker tag 5250d83bc977 andycsoka/stock_price_checker:0.4
docker push andycsoka/stock_price_checker
```

# task

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
