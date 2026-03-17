# Staging

Deploy nextpress-backend for staging (pre-production) testing on a server. Fixed folder, `staging` branch, systemd, Nginx, TLS. This environment can be on its own server or share a server with others.

| Item   | Value                                      |
|--------|--------------------------------------------|
| Branch | `staging`                                  |
| Folder | `/var/www/nextpress-backend-staging`       |
| Port   | Set in `.env` on this server (e.g. 9091)   |
| Domain | e.g. `cms-staging.yourdomain.com`          |

---

## 1. Environment file (.env)

The app reads configuration from `.env` in the project folder. Create it from the example:

```bash
cd /var/www/nextpress-backend-staging
cp .env.example .env
```

Edit `.env`. Required:

| Variable                    | Purpose |
|----------------------------|---------|
| `APP_PORT`                 | TCP port the API listens on (this server). Nginx will proxy to this port. |
| `DB_DRIVER`, `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_USER`, `DB_PASSWORD`, `DB_SSLMODE` | PostgreSQL connection (future phases). |

Set `APP_ENV=staging` in `.env`.

---

## 2. Deploy

From the staging folder, run the deploy script with the `staging` argument. It expects `.env` to exist.

```bash
./scripts/deploy staging
```

The script checks out latest `staging`, builds `bin/server`, and restarts the systemd service `nextpress-backend@staging` after you complete the Systemd step below.

---

## 3. Systemd

Run the API as a systemd service so it starts on boot and restarts on failure. The service name is **nextpress-backend@staging**. Create the unit file `/etc/systemd/system/nextpress-backend@.service` (the `%i` in the file is replaced by systemd with the instance name `staging`):

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
sudo cp /var/www/nextpress-backend-staging/deploy/systemd/nextpress-backend@.service /etc/systemd/system/
sudo systemctl daemon-reload
```

Enable and start the service:

```bash
sudo systemctl enable nextpress-backend@staging
sudo systemctl start nextpress-backend@staging
```

Check status: `sudo systemctl status nextpress-backend@staging`. The deploy script will restart this service on future runs.

---

## 4. Nginx

Nginx receives HTTP/HTTPS and forwards to nextpress-backend. Copy the config, enable it, test, and reload:

```bash
sudo cp deploy/nginx/staging.conf /etc/nginx/sites-available/nextpress-backend-staging.conf
sudo ln -sf /etc/nginx/sites-available/nextpress-backend-staging.conf /etc/nginx/sites-enabled/
sudo nginx -t && sudo systemctl reload nginx
```

Edit the copied file: set `server_name` to your staging domain and `proxy_pass` to the app port (must match `APP_PORT` in `.env`). See `deploy/nginx/README.md` for details.

---

## 5. TLS

After DNS for your staging domain points to this server and Nginx serves HTTP, get a certificate with Let's Encrypt:

```bash
sudo apt install -y certbot python3-certbot-nginx
sudo certbot --nginx -d <your-domain>
```

Replace `<your-domain>` with your staging domain (e.g. `cms-staging.yourdomain.com`). certbot will update the Nginx config and reload it.

---

[← Menu](../DEPLOYMENT.md)

