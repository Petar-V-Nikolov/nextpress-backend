# Production

Deploy nextpress-backend for production traffic on a server. Fixed folder, `main` branch, systemd, Nginx, TLS. This environment can be on its own server or share a server with others.

| Item   | Value                                       |
|--------|---------------------------------------------|
| Branch | `main`                                      |
| Folder | `/var/www/nextpress-backend-production`     |
| Port   | Set in `.env` on this server (e.g. 9090)    |
| Domain | e.g. `cms.yourdomain.com`                   |

---

## 1. Environment file (.env)

The app reads configuration from `.env` in the project folder. Create it from the example:

```bash
cd /var/www/nextpress-backend-production
cp .env.example .env
```

Edit `.env`. Required:

| Variable                    | Purpose |
|----------------------------|---------|
| `APP_PORT`                 | TCP port the API listens on (this server). Nginx will proxy to this port. |
| `DB_*`                     | PostgreSQL connection. |
| `JWT_SECRET`               | Strong secret for JWT signing (never commit real values). |

Set `APP_ENV=production` in `.env`. See `.env.example` for media, rate limits, and optional `RUN_SEED_ON_DEPLOY`.

---

## 2. Deploy

From the production folder, run the deploy script. It expects `.env` to exist.

```bash
chmod +x scripts/deploy
./scripts/deploy
```

The script checks out latest `main`, builds `bin/server`, and restarts the systemd service `nextpress-backend@production` after you complete the Systemd step below.

---

## 3. Systemd

Run the API as a systemd service so it starts on boot and restarts on failure. The service name is **nextpress-backend@production**. Create the unit file `/etc/systemd/system/nextpress-backend@.service` (the `%i` in the file is replaced by systemd with the instance name `production`):

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
sudo cp /var/www/nextpress-backend-production/deploy/systemd/nextpress-backend@.service /etc/systemd/system/
sudo systemctl daemon-reload
```

Enable and start the service:

```bash
sudo systemctl enable nextpress-backend@production
sudo systemctl start nextpress-backend@production
```

Check status: `sudo systemctl status nextpress-backend@production`. The deploy script will restart this service on future runs.

---

## 4. Nginx

Nginx receives HTTP/HTTPS and forwards to nextpress-backend. Copy the config, enable it, test, and reload:

```bash
sudo cp deploy/nginx/production.conf /etc/nginx/sites-available/nextpress-backend-production.conf
sudo ln -sf /etc/nginx/sites-available/nextpress-backend-production.conf /etc/nginx/sites-enabled/
sudo nginx -t && sudo systemctl reload nginx
```

Edit the copied file: set `server_name` to your domain and `proxy_pass` to the app port (must match `APP_PORT` in `.env`). See `deploy/nginx/README.md` for details.

---

## 5. TLS

After DNS points to this server and Nginx serves HTTP, get a certificate with Let’s Encrypt:

```bash
sudo apt install -y certbot python3-certbot-nginx
sudo certbot --nginx -d <your-domain>
```

Replace `<your-domain>` with your domain (e.g. `cms.yourdomain.com`). certbot will update the Nginx config and reload it.

---

[← Menu](../DEPLOYMENT.md)

