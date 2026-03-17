# Nginx

One config per environment. Edit `server_name` and `proxy_pass` to match your domain and app port.

| File              | Domain (example)            | Port (example) |
|-------------------|----------------------------|----------------|
| `production.conf` | cms.yourdomain.com         | 9090           |
| `staging.conf`    | cms-staging.yourdomain.com | 9091           |
| `dev.conf`        | cms-dev.yourdomain.com     | 9092           |

## Enable

```bash
sudo cp deploy/nginx/production.conf /etc/nginx/sites-available/nextpress-backend-production.conf
sudo ln -sf /etc/nginx/sites-available/nextpress-backend-production.conf /etc/nginx/sites-enabled/
sudo nginx -t && sudo systemctl reload nginx
```

Repeat for staging/dev if needed, adjusting filenames.

## TLS (Let's Encrypt)

```bash
sudo apt install -y certbot python3-certbot-nginx
sudo certbot --nginx -d <your-domain>
```

Replace `<your-domain>` with the domain in your config (e.g. `cms.yourdomain.com`). certbot will update the Nginx config and reload it.

