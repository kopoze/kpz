# Kopoze

DevOps toolikts made with go.

**Note:** The project is still in active development so there will be a lot of bug. And all step here was only tested on Ubuntu 20.04.

## Requirements

- PostresSQL
- Nginx
- Mkcert

## Install

### For local use

For local use, this cli is to simplify subdomain-port mapping with https of your local project.

#### Get the project

```sh
wget -O - https://github.com/hantsaniala/kopoze/raw/main/install.sh | sudo bash
```

The following step will be added to `install.sh` later.

#### Configure your hosts

Add this line to your `/etc/hosts` file.

```text
127.0.0.1  project.mg  app1.project.mg
```

**Note:** For now, all subdomain must sepecified manually. And musti be under the domain `project.mg` for now until next update.

#### Generate certificate for local https

Generate certificate for a wildcard subdomain.

```sh
mkcert *.project.mg
```

To install and configure Mkcert, you can read more [here](https://www.howtoforge.com/how-to-create-locally-trusted-ssl-certificates-with-mkcert-on-ubuntu/).

#### Configure nginx

Now create subdomain configuration under `/etc/nginx/sites-availble/sub.project.mg` with the following value.

```conf
server {
    server_name *.project.mg;

    location / {
    proxy_pass http://localhost:8080;
    proxy_pass_header Set-Cookie;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_set_header X-Forwarded-Host $host;
    proxy_set_header Cookie $http_cookie;

    proxy_cookie_path / /;
    client_max_body_size 5m;
    }


    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/_wildcard.project.mg/_wildcard.project.mg.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/_wildcard.project.mg/_wildcard.project.mg-key.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

}
server {
    if ($host = *.project.mg) {
        return 301 https://$host$request_uri;
    } # managed by Certbot

    server_name *.project.mg;
    listen 80;
    return 404; # managed by Certbot
}
```

Don't forget to change the value of `ssl_certificate` and `ssl_certificate_key` with the path of the generated certificate with `mkcert` earlier.

Note that the port used inside your config is the port that will be used when starting our script with `kpz serve`.

At the end, validate your configuration with `sudo nginx -t` and if everything is ok, restart your nginx server with `sudo systemctl restart nginx`.

#### PostgreSQL configuration

Create a `.env` file with the following keys:

```ini
DB_HOST=localhost
DB_PORT=5432
DB_USER=i-am-groot
DB_NAME=kopoze
DB_PASSWORD=MyStr0ngP@ssw0rD
```

Of course feel free to change it according to your settings.

#### Run the program

Now you can run the cli with

```sh
kpz serve
```

#### Add your app

You can use the existing API under `http://locahost:8080/cli/apps/` or add directly your entry inside PostresSQL.

Here is the JSON format to create app:

```http
POST http://localhost:8080/cli/apps/
Content-Type: application/json

{
    "name": "App1",
    "subdomain": "app1",
    "port": 9000
}
```

After adding your app, don't forget to add this subdomain inside your `/etc/hosts`.

## Limitation

For know, accessing Django admin with this tool need specific configuration on Django side because `sessionId` cookies is not set. Still can't figure out the way out.

## Author

&copy; [Kopoze](https://t.me/hantsaniala3) 2023
