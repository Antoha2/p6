 docker build -t nginx-local .
 docker run -d -p 8080:80 -v html:/usr/share/nginx/html nginx-local  