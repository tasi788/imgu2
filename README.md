# imgu2

[简体中文](https://github.com/sduoduo233/imgu2/blob/master/README_zh_cn.md)

An image sharing platform powered by Golang

# Screenshots

<details>
  <summary>Show screenshots</summary>

![image preview page](https://github.com/sduoduo233/imgu2/blob/master/screenshots/1.png?raw=true)

![my uploads page](https://github.com/sduoduo233/imgu2/blob/master/screenshots/2.png?raw=true)

![user list](https://github.com/sduoduo233/imgu2/blob/master/screenshots/3.png?raw=true)

![image list](https://github.com/sduoduo233/imgu2/blob/master/screenshots/4.png?raw=true)

![login page](https://github.com/sduoduo233/imgu2/blob/master/screenshots/5.png?raw=true)

</details>

# Features

- Lightweight design
- OAuth login with Google & GitHub
- Image re-encoding
- SQLite database integration
- Multiple storage options supported, including S3-compatible, FTP, and local file systems

# How to install

1. Get the latest binary from the GitHub releases.

2. Transfer the downloaded file to your Linux web server.

3. In the same directory as the executable file, create a `.env` file with the following content:

```
IMGU2_SMTP_USERNAME=mailer@example.com
IMGU2_SMTP_PASSWORD=example_password
IMGU2_SMTP_HOST=example.com
IMGU2_SMTP_PORT=25
IMGU2_SMTP_SENDER=mailer@example.com
IMGU2_JWT_SECRET=example_secret_string
```

`IMGU2_JWT_SECRET` should be a hard-to-guess string, which can be generated using `openssl rand -hex 8` on Linux.

`IMGU2_SMTP_*` are the SMTP configurations required for sending verification emails.

4. Run the executable using `./imgu2-linux-amd64`.

5. Configure NGINX for Reverse Proxying and SSL. Add the following configuration to NGINX:

```nginx
server {
  listen 443 ssl;
  ssl_certificate /path/to/ssl/certificate;
  ssl_certificate_key /path/to/ssl/certificate/key;

  server_name example.com;
  root /var/www/html;
  index index.php index.html;

  client_max_body_size 32M;

  location / {
    proxy_pass http://127.0.0.1:3000;
  }
}
```

6. The default admin email is `admin@example.com`, and the password is `admin`. It is crucial to change these credentials as soon as you deploy the platform for security reasons.