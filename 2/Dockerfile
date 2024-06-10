FROM nginx:alpine

CMD sh -c 'echo "<html><body><h1>${MY_ENV}</h1></body></html>" > /usr/share/nginx/html/index.html && nginx -g "daemon off;"'

