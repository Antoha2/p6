docker build -t greeting .
docker run -e MY_ENV="hello, world" -p 80:80 greeting