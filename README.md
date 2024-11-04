# GoShortURL

### Requirements:

1. Git
2. Docker & Docker-Compose
3. Web-server (Nginx or Apache or smth else)

### How to run this project on server:

1. Clone repo:

```shell
git clone https://github.com/ej-you/GoShortURL.git
```

2. Add `.env` file (into root of project) and fill it like this:

```dotenv
# host and port for local start (in docker container)
HOST_IP=0.0.0.0
HOST_PORT=8000

# external host for internet (it must be set up on web-server)
HOST_FOR_SHORT_LINK=https://your-domain.com
```

3. Run app in background:

```shell
docker-compose up --build -d
```

4. Set up Nginx (this is just example, you can use another config or web-server):

```nginx configuration
## URL Shorter on Golang + html/css/js


server {
    # external port
    listen 8090 ssl;
    #http2 on;
    # server name
    server_name your-domain.com;

    ssl_certificate     /full/path/to/your-domain.com/fullchain.pem;
    ssl_certificate_key /full/path/to/your-domain.com/privkey.pem;

    # proxy to golang app
    location / {
    	# ip and listening port of docker container
        proxy_pass http://172.20.1.5:8085;
    }
}
```

5. Add allow rule to firewall (if you need):

```shell
ufw allow 8090/tcp
ufw reload
```
