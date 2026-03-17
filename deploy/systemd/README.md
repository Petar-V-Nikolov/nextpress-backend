# Systemd

One template unit for all environments. Instance name (`%i`) = production, staging, or dev.

- **Template:** `nextpress-backend@.service` → install to `/etc/systemd/system/`
- **Instances:** `nextpress-backend@production`, `nextpress-backend@staging`, `nextpress-backend@dev`
- **Folders:** `/var/www/nextpress-backend-%i` (e.g. `/var/www/nextpress-backend-production`)
- **APP_ENV:** The unit sets `Environment=APP_ENV=%i` so the process gets `APP_ENV=production`, `APP_ENV=staging`, or `APP_ENV=dev` from the instance name. `.env` can override if needed.

Install once:

```bash
sudo cp deploy/systemd/nextpress-backend@.service /etc/systemd/system/
sudo systemctl daemon-reload
```

Then per environment:

```bash
sudo systemctl enable nextpress-backend@<env>
sudo systemctl start nextpress-backend@<env>
```

See deployment guides under `docs/deployment/` for full steps.

