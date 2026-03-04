FROM nginx

COPY ./filebrowser /usr/share/nginx/html/

# COPY ./nginx/server.pem /usr/share/nginx/ssl/server.pem
# COPY ./nginx/server.key /usr/share/nginx/ssl/server.key
# COPY ./nginx/web.conf /etc/nginx/conf.d/web.conf

WORKDIR /usr/share/nginx/html/

EXPOSE 8080 8443

CMD ["nginx", "-g", "daemon off;"]
