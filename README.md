# HTTP-Buddy
HTTP-Buddy allows you to easily test your HTTP clients. Simply make a request
to one of the below endpoints:
* /myip
* /useragent
* /images/jpg
* /images/png
* /methods/post
* /methods/get
* /methods/delete
* /methods/put
* /methods/patch

# Building
You can either build HTTP-Buddy from source, or deploy the docker container.

## Building From Source
```
/code/http-buddy $ go build -o http-buddy 
/code/http-buddy $ ./http-buddy
```

## Deploy with Docker
```
/code/http-buddy $ docker build -t http-buddy .
/code/http-buddy $ docker run -d -p 8080:8080 http-buddy
```

# Usage
```
/code/http-buddy $ curl http://localhost:8080/myip
/code/http-buddy $ {"IP":"172.17.0.1","Port":"55120"}

/code/http-buddy $ curl http://localhost:8080/headers
/code/http-buddy $ {"Headers":[{"User-Agent":"curl/7.47.0"},{"Accept":"*/*"}],"count":2}
```

# Demo Server
If you would like to try HTTP-Buddy out without running it yourself, there is an instance running at:
```
http://159.65.177.76:8888/
```
