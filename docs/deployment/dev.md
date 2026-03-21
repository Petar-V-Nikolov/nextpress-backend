# Dev

Deploy nextpress-backend for development or integration testing on a server. Fixed folder, `dev` branch, systemd, Nginx, TLS. This environment can be on its own server or share a server with others.

| Item   | Value                                   |
|--------|-----------------------------------------|
| Branch | `dev`                                   |
| Folder | `/var/www/nextpress-backend-dev`        |
| Port   | Set in `.env` on this server (e.g. 9092)|
| Domain | e.g. `cms-dev.yourdomain.com`           |

---

## 1. Environment file (.env)

The app reads configuration from `.env` in the project folder. Create it from the example:

```bash
cd /var/www/nextpress-backend-dev
cp .env.example .env
```

Edit `.env`. Required:

| Variable                    | Purpose |
|----------------------------|---------|
| `APP_PORT`                 | TCP port the API listens on (this server). Nginx will proxy to this port. |
| `DB_*`                     | PostgreSQL connection (required for full API). |
| `JWT_SECRET`               | Signing key for access/refresh tokens. |
| `MEDIA_*` / `RATE_LIMIT_*` | Optional; see `.env.example`. |

Set `APP_ENV=dev` in `.env`.

---

## 2. Deploy

From the dev folder, run the deploy script with the `dev` argument. It expects `.env` to exist.

```bash
./scripts/deploy dev
```

The script checks out latest `dev`, builds `bin/server`, and restarts the systemd service `nextpress-backend@dev` after you complete the Systemd step below.

---

## 3. Systemd

Run the API as a systemd service so it starts on boot and restarts on failure. The service name is **nextpress-backend@dev**. Create the unit file `/etc/systemd/system/nextpress-backend@.service` (the `%i` in the file is replaced by systemd with the instance name `dev`):

```ini
[Unit]
Description=nextpress-backend (%i)
After=network.target

[Service]
Type=simple
WorkingDirectory=/var/www/nextpress-backend-%i
Environment=APP_ENV=%i
EnvironmentFile=/var/www/nextpress-backend-%i/.env
ExecStart=/var/www/nextpress-backend-%i/bin/server
Restart=always
RestartSec=5
User=www-data
Group=www-data

[Install]
WantedBy=multi-user.target
```

Copy from the repo:

```bash
sudo cp /var/www/nextpress-backend-dev/deploy/systemd/nextpress-backend@.service /etc/systemd/system/
sudo systemctl daemon-reload
```

Enable and start the service:

```bash
sudo systemctl enable nextpress-backend@dev
sudo systemctl start nextpress-backend@dev
```

Check status: `sudo systemctl status nextpress-backend@dev`. The deploy script will restart this service on future runs.

---

## 4. Nginx

Nginx receives HTTP/HTTPS and forwards to nextpress-backend. Copy the config, enable it, test, and reload:

```bash
sudo cp deploy/nginx/dev.conf /etc/nginx/sites-available/nextpress-backend-dev.conf
sudo ln -sf /etc/nginx/sites-available/nextpress-backend-dev.conf /etc/nginx/sites-enabled/
sudo nginx -t && sudo systemctl reload nginx
```

Edit the copied file: set `server_name` to your dev domain and `proxy_pass` to the app port (must match `APP_PORT` in `.env`). See `deploy/nginx/README.md` for details.

---

## 5. TLS

After DNS for your dev domain points to this server and Nginx serves HTTP, get a certificate with Let's Encrypt:

```bash
sudo apt install -y certbot python3-certbot-nginx
sudo certbot --nginx -d <your-domain>
```

Replace `<your-domain>` with your dev domain (e.g. `cms-dev.yourdomain.com`). certbot will update the Nginx config and reload it.

---

[← Menu](../DEPLOYMENT.md)

