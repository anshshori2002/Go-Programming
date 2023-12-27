Build a webserver in Go to receive data external and route through messaging systems
● The webserver has two endpoints : one to get data, and another to push data
● A POST request to the webserver has one endpoint that receives a request and routes it
to kafka topic via goroutine-1. Another goroutine-2 will receive this data from kafka topic
and routes it to REDIS database
● A GET request will retrieve the data from REDIS to Kafka via goroutine-3, and delivers
the response
Design a Request / Response to address this requirement.
