# GO pricing microservice

This is a simple tutorial that demonstrates how to build
a microservice in GoLang. It includes routes, controllers,
services and tests.

This tutorial makes use of github.com/gin-gonic/gin for routing.

The code accepts a start and end date and returns a price back
by applying an hourly / daily rate to the given interval.

### Usage

To start the microservice:

1. Clone the repository into your GOPATH
2. Install `realize` (optional)
```
go get github.com/oxequa/realize
```
3. Start the project
```
realize start
```
4. Head over to http://127.0.0.1:3000/status and ensure the page loads.
5. Send a post request
```
curl -X POST http://127.0.0.1:3000/pricing/calculate \
     -d 'start_date=2018-11-01T10:00:00%2B00:00&end_date=2018-11-01T15:00:00%2B00:00'
```
6. You should get a price back.
