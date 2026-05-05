# NextPressKit dev CLI for Windows (PowerShell). From repo root:
#   .\scripts\nextpresskit.ps1 setup
#   .\scripts\nextpresskit.ps1 run
# Requires: Go on PATH, PostgreSQL for migrate/seed.

$ErrorActionPreference = "Stop"
$RootDir = Split-Path -Parent $PSScriptRoot
Set-Location $RootDir
$env:CGO_ENABLED = if ($env:CGO_ENABLED) { $env:CGO_ENABLED } else { "0" }

function Get-AppPort {
    $envFile = Join-Path $RootDir ".env"
    if (-not (Test-Path $envFile)) { return "9090" }
    foreach ($line in Get-Content $envFile) {
        if ($line -match '^\s*APP_PORT\s*=\s*(.+)$') {
            return $matches[1].Trim()
        }
    }
    "9090"
}

function Get-DevRuntimeBasename {
    $envFile = Join-Path $RootDir ".env"
    if (-not (Test-Path $envFile)) { return "nextpresskit" }
    foreach ($line in Get-Content $envFile) {
        if ($line -match '^\s*APP_DEV_RUNTIME_BASENAME\s*=\s*(.+)$') {
            $b = $matches[1].Trim()
            if (-not [string]::IsNullOrWhiteSpace($b)) { return $b }
        }
    }
    "nextpresskit"
}

function Test-PortListen([int]$Port) {
    $c = Get-NetTCPConnection -LocalPort $Port -State Listen -ErrorAction SilentlyContinue
    return $null -ne $c
}

function Assert-PortFree([int]$Port) {
    if (Test-PortListen $Port) {
        Write-Host "Port $Port is already in use. Stop the process or change APP_PORT in .env." -ForegroundColor Yellow
        Get-NetTCPConnection -LocalPort $Port -State Listen -ErrorAction SilentlyContinue |
            ForEach-Object { Get-Process -Id $_.OwningProcess -ErrorAction SilentlyContinue }
        exit 1
    }
}

function Get-Version {
    Push-Location $RootDir
    try {
        $out = & git describe --tags --always --dirty 2>$null
        if ($LASTEXITCODE -eq 0 -and $out) { return $out }
    } finally {
        Pop-Location
    }
    return "dev"
}

function Show-Help {
    @"
NextPressKit dev CLI (Windows PowerShell). On Linux/macOS use: ./scripts/nextpresskit

Commands:
  help install deps tidy build build-all setup (text menu on console; NP_SETUP_NONINTERACTIVE=1 = linear only)
  migrate-up migrate-down migrate-version migrate-steps <N> seed
  run start stop deploy checks
  test test-coverage test-integration security-check clean postman-sync

Examples:
  .\scripts\nextpresskit.ps1 setup
  .\scripts\nextpresskit.ps1 run
"@
}

function Ensure-Go {
    if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
        Write-Error "Go is not installed or not on PATH."
    }
}

function Test-NpSetupLinearOnly {
    if ($env:NP_SETUP_NONINTERACTIVE -eq "1") { return $true }
    try { return [Console]::IsInputRedirected } catch { return $true }
}

function Get-SetupMenuStepOrder {
    return @(
        "install", "deps", "tidy", "clean", "build", "build-all",
        "migrate-drop", "db-fresh", "migrate-up", "seed",
        "test", "test-coverage", "test-integration", "security-check", "checks",
        "local-https", "deploy-nginx", "deploy", "start", "stop"
    )
}

function Get-SetupMenuHint([string]$id) {
    switch ($id) {
        "install" { return "go mod download; .env from .env.example if missing" }
        "deps" { return "go mod download only" }
        "tidy" { return "go mod tidy" }
        "clean" { return "rm bin/; go clean" }
        "build" { return "go build → bin/server" }
        "build-all" { return "bin/server, bin/migrate, bin/seed" }
        "migrate-up" { return "cmd/migrate up (MODULES)" }
        "seed" { return "cmd/seed" }
        "migrate-drop" { return "drop public tables (confirm)" }
        "db-fresh" { return "migrate-drop + migrate-up" }
        "test" { return "go test -v ./..." }
        "test-coverage" { return "go test -cover ./..." }
        "test-integration" { return "go test -tags=integration (needs DB_*)" }
        "security-check" { return "govulncheck ./..." }
        "checks" { return "test, vet, integration, openapi validate, govulncheck" }
        "local-https" { return "scripts/setup-local-https.sh" }
        "deploy-nginx" { return "scripts/deploy apply-nginx" }
        "deploy" { return "scripts/deploy" }
        "start" { return "background API (Unix)" }
        "stop" { return "stop background API" }
        default { return $id }
    }
}

function Normalize-SetupMenuSteps([string[]]$raw) {
    $dbFresh = $raw -contains "db-fresh"
    $out = [System.Collections.Generic.List[string]]::new()
    foreach ($id in (Get-SetupMenuStepOrder)) {
        if ($id -notin $raw) { continue }
        if ($dbFresh -and ($id -eq "migrate-drop" -or $id -eq "migrate-up")) { continue }
        $out.Add($id)
    }
    return $out.ToArray()
}

function Invoke-SetupMenuRunStep([string]$id) {
    switch ($id) {
        "local-https" {
            $sh = Join-Path $RootDir "scripts\setup-local-https.sh"
            $bash = Get-Command bash -ErrorAction SilentlyContinue
            if ($bash) { & $bash.Path $sh } else { Write-Host "Need Git Bash/WSL for local-https script." -ForegroundColor Yellow }
            return
        }
        "deploy-nginx" {
            $bash = Get-Command bash -ErrorAction SilentlyContinue
            if ($bash) { & $bash.Path (Join-Path $RootDir "scripts\deploy") "apply-nginx" }
            else { Write-Warning "bash not found for deploy-apply-nginx" }
            return
        }
        "deploy" {
            & (Join-Path $RootDir "scripts\deploy.ps1")
            return
        }
        default { & $PSCommandPath $id }
    }
}

function Get-KitModuleOrder {
    return @("user", "rbac", "auth", "taxonomy", "media", "posts", "pages")
}

function Get-EnvModulesValue {
    $p = Join-Path $RootDir ".env"
    if (-not (Test-Path $p)) { return "" }
    foreach ($line in Get-Content $p) {
        if ($line -match '^\s*MODULES\s*=\s*(.*)$') {
            return $matches[1].Trim()
        }
    }
    return ""
}

function Set-EnvModulesValue([string]$newVal) {
    $p = Join-Path $RootDir ".env"
    if (-not (Test-Path $p)) {
        Write-Host "No .env — option 1 (full setup) or: .\scripts\nextpresskit.ps1 install" -ForegroundColor Yellow
        return $false
    }
    $lines = @(Get-Content $p -ErrorAction Stop)
    $out = [System.Collections.Generic.List[string]]::new()
    $rep = $false
    foreach ($line in $lines) {
        if ($line -match '^\s*MODULES\s*=') {
            [void]$out.Add("MODULES=$newVal")
            $rep = $true
        } else {
            [void]$out.Add($line)
        }
    }
    if (-not $rep) { [void]$out.Add("MODULES=$newVal") }
    Set-Content -Path $p -Value ($out -join "`n") -NoNewline:$false -Encoding utf8
    return $true
}

function Test-TokenInModuleList([string]$needle, [string[]]$hay) {
    $n = $needle.ToLowerInvariant().Trim()
    foreach ($h in $hay) {
        if ($h.ToLowerInvariant().Trim() -eq $n) { return $true }
    }
    return $false
}

function Invoke-ModulesAddMenu {
    $raw = (Get-EnvModulesValue).Trim()
    if ([string]::IsNullOrWhiteSpace($raw)) {
        Write-Host "MODULES empty = all modules. Use option 7 first for an explicit list, then 6 to add."
        return
    }
    $have = @($raw.Split(",") | ForEach-Object { $_.Trim().ToLowerInvariant() } | Where-Object { $_ })
    $order = Get-KitModuleOrder
    $avail = [System.Collections.Generic.List[string]]::new()
    foreach ($m in $order) {
        if (-not (Test-TokenInModuleList $m $have)) { $avail.Add($m) }
    }
    if ($avail.Count -eq 0) {
        Write-Host "All ids already in MODULES."
        return
    }
    Write-Host "`nAdd MODULES id:"
    for ($i = 0; $i -lt $avail.Count; $i++) {
        Write-Host ("{0,3}) {1}" -f ($i + 1), $avail[$i])
    }
    $pick = Read-Host "Number [0=cancel]"
    if ($pick -eq "0" -or [string]::IsNullOrWhiteSpace($pick)) { return }
    $ix = 0
    if (-not [int]::TryParse($pick, [ref]$ix)) { Write-Host "Invalid." -ForegroundColor Red; return }
    $ix--
    if ($ix -lt 0 -or $ix -ge $avail.Count) { Write-Host "Invalid." -ForegroundColor Red; return }
    $add = $avail[$ix]
    $newcsv = "$raw,$add"
    if (Set-EnvModulesValue $newcsv) {
        Write-Host "MODULES=$newcsv"
        Write-Host "Then: migrate-up, seed. docs/MODULES.md"
    }
}

function Invoke-ModulesRemoveMenu {
    $raw = (Get-EnvModulesValue).Trim()
    $order = Get-KitModuleOrder
    if ([string]::IsNullOrWhiteSpace($raw)) {
        Write-Host "MODULES empty = full kit. Remove one (rest → .env):"
        for ($i = 0; $i -lt $order.Count; $i++) {
            Write-Host ("{0,3}) {1}" -f ($i + 1), $order[$i])
        }
        $pick = Read-Host "Number [0=cancel]"
        if ($pick -eq "0" -or [string]::IsNullOrWhiteSpace($pick)) { return }
        $ix = 0
        if (-not [int]::TryParse($pick, [ref]$ix)) { Write-Host "Invalid." -ForegroundColor Red; return }
        $ix--
        if ($ix -lt 0 -or $ix -ge $order.Count) { Write-Host "Invalid." -ForegroundColor Red; return }
        $rem = $order[$ix]
        $parts = [System.Collections.Generic.List[string]]::new()
        foreach ($m in $order) {
            if ($m -ne $rem) { $parts.Add($m) }
        }
        $newcsv = ($parts -join ",")
        if (Set-EnvModulesValue $newcsv) {
            Write-Host "MODULES=$newcsv"
        }
    } else {
        $toks = @($raw.Split(",") | ForEach-Object { $_.Trim().ToLowerInvariant() } | Where-Object { $_ })
        if ($toks.Count -eq 0) { Write-Host "Could not parse MODULES." -ForegroundColor Red; return }
        Write-Host "`nMODULES tokens — remove one:"
        for ($i = 0; $i -lt $toks.Count; $i++) {
            Write-Host ("{0,3}) {1}" -f ($i + 1), $toks[$i])
        }
        $pick = Read-Host "Number [0=cancel]"
        if ($pick -eq "0" -or [string]::IsNullOrWhiteSpace($pick)) { return }
        $ix = 0
        if (-not [int]::TryParse($pick, [ref]$ix)) { Write-Host "Invalid." -ForegroundColor Red; return }
        $ix--
        if ($ix -lt 0 -or $ix -ge $toks.Count) { Write-Host "Invalid." -ForegroundColor Red; return }
        $out = [System.Collections.Generic.List[string]]::new()
        for ($j = 0; $j -lt $toks.Count; $j++) {
            if ($j -ne $ix) { $out.Add($toks[$j]) }
        }
        $newcsv = ($out -join ",")
        if (Set-EnvModulesValue $newcsv) {
            if ([string]::IsNullOrWhiteSpace($newcsv)) {
                Write-Host "MODULES= (empty → full default)"
            } else {
                Write-Host "MODULES=$newcsv"
            }
        }
    }
    Write-Host "migrate-up if needed. docs/MODULES.md"
}

function Invoke-LinearSetupPS {
    & $PSCommandPath install
    & $PSCommandPath build-all
    & $PSCommandPath migrate-up
    & $PSCommandPath seed
    Write-Host "Local HTTPS: use Git Bash ./scripts/nextpresskit setup or mkcert manually on Windows."
    Write-Host "Setup complete. Run: .\scripts\nextpresskit.ps1 run"
}

function Invoke-SetupTextMenu {
    $ordAll = Get-SetupMenuStepOrder
    while ($true) {
        Write-Host ""
        Write-Host "NextPressKit setup"
        Write-Host "------------------"
        Write-Host " 1) Full setup — install, build-all, migrate-up, seed; mkcert/HTTPS if interactive (skip when piped/CI)"
        Write-Host " 2) Refresh database — migrate-up + seed (after pulling code or changing MODULES)"
        Write-Host " 3) Compile only — install + build-all (binaries; no DB)"
        Write-Host " 4) Quality gate — tests, vet, integration, OpenAPI check, govulncheck"
        Write-Host " 5) Custom — run chosen nextpresskit steps (numbered list)"
        Write-Host " 6) Enable a kit module — add id to MODULES in .env (then run migrate-up / seed)"
        Write-Host " 7) Disable a kit module — remove id from MODULES in .env (then migrate-up if needed)"
        Write-Host " 8) Exit"
        Write-Host ""
        $c = Read-Host "[1-8]"
        $raw = @()
        switch ($c) {
            "1" { $raw = @("install", "build-all", "migrate-up", "seed", "local-https") }
            "2" { $raw = @("migrate-up", "seed") }
            "3" { $raw = @("install", "build-all") }
            "4" { $raw = @("checks") }
            "5" {
                Write-Host ""
                Write-Host "Step numbers (space/comma). 0 or empty = back."
                $i = 1
                foreach ($id in $ordAll) {
                    Write-Host ("{0,3}) {1,-14} {2}" -f $i, $id, (Get-SetupMenuHint $id))
                    $i++
                }
                Write-Host ""
                $line = Read-Host ">"
                if ([string]::IsNullOrWhiteSpace($line)) { continue }
                if ($line.Trim() -eq "0") { continue }
                $pick = [System.Collections.Generic.List[string]]::new()
                foreach ($tok in $line.Replace(",", " ").Split(" ", [StringSplitOptions]::RemoveEmptyEntries)) {
                    $n = 0
                    if (-not [int]::TryParse($tok.Trim(), [ref]$n)) { continue }
                    $ix = $n - 1
                    if ($ix -ge 0 -and $ix -lt $ordAll.Count) { $pick.Add($ordAll[$ix]) }
                }
                $raw = Normalize-SetupMenuSteps ($pick.ToArray())
                if ($raw.Count -eq 0) {
                    Write-Host "No valid step numbers (use 1-$($ordAll.Count))." -ForegroundColor Yellow
                    continue
                }
                foreach ($id in $raw) {
                    if ($id -eq "db-fresh" -or $id -eq "migrate-drop") {
                        $h = Get-SetupMenuHint $id
                        $ok = Read-Host "Run '$id' — $h ? [y/N]"
                        if ($ok -ne "y") { continue }
                    }
                    Write-Host "==> $id"
                    Invoke-SetupMenuRunStep $id
                }
                Write-Host "`nDone. .\scripts\nextpresskit.ps1 run"
                return
            }
            "6" {
                Invoke-ModulesAddMenu
                return
            }
            "7" {
                Invoke-ModulesRemoveMenu
                return
            }
            "8" { return }
            "" { return }
            default {
                Write-Host "Invalid choice (1-8)." -ForegroundColor Red
                continue
            }
        }
        $ordered = Normalize-SetupMenuSteps $raw
        foreach ($id in $ordered) {
            if ($id -eq "db-fresh" -or $id -eq "migrate-drop") {
                $h = Get-SetupMenuHint $id
                $ok = Read-Host "Run '$id' — $h ? [y/N]"
                if ($ok -ne "y") { continue }
            }
            Write-Host "==> $id"
            Invoke-SetupMenuRunStep $id
        }
        Write-Host "`nDone. .\scripts\nextpresskit.ps1 run"
        return
    }
}

$cmd = if ($args.Count -ge 1) { $args[0] } else { "help" }
if ($cmd -eq "-h" -or $cmd -eq "--help") { $cmd = "help" }

switch ($cmd) {
    "help" { Show-Help }
    "install" {
        Ensure-Go
        go mod download
        $envPath = Join-Path $RootDir ".env"
        if (-not (Test-Path $envPath)) {
            Copy-Item (Join-Path $RootDir ".env.example") $envPath
            Write-Host "Created .env from .env.example — edit DB_* and JWT_SECRET."
        } else {
            Write-Host ".env already exists."
        }
    }
    "deps" { Ensure-Go; go mod download }
    "tidy" { Ensure-Go; go mod tidy }
    "build" {
        Ensure-Go
        New-Item -ItemType Directory -Force -Path (Join-Path $RootDir "bin") | Out-Null
        $v = Get-Version
        go build -ldflags "-X main.version=$v" -o (Join-Path $RootDir "bin\server.exe") ./cmd/api
        Write-Host "Built bin\server.exe"
    }
    "build-all" {
        Ensure-Go
        New-Item -ItemType Directory -Force -Path (Join-Path $RootDir "bin") | Out-Null
        $v = Get-Version
        go build -ldflags "-X main.version=$v" -o (Join-Path $RootDir "bin\server.exe") ./cmd/api
        go build -o (Join-Path $RootDir "bin\migrate.exe") ./cmd/migrate
        go build -o (Join-Path $RootDir "bin\seed.exe") ./cmd/seed
        Write-Host "Built bin\server.exe, bin\migrate.exe, bin\seed.exe"
    }
    "setup" {
        if (Test-NpSetupLinearOnly) {
            Invoke-LinearSetupPS
        } else {
            Invoke-SetupTextMenu
        }
    }
    "migrate-up" { Ensure-Go; go run ./cmd/migrate -command=up }
    "migrate-down" {
        Write-Error "migrate-down is removed (GORM AutoMigrate). For dev reset use: db-fresh"
    }
    "migrate-version" {
        Write-Host "No migration version row: schema is internal/platform/dbmigrate + cmd/migrate up."
    }
    "migrate-steps" {
        Write-Error "migrate-steps is removed (no numbered SQL migrations)."
    }
    "migrate-drop" {
        Ensure-Go
        Write-Warning "This drops all tables."
        $c = Read-Host "Type y to continue"
        if ($c -ne "y") { exit 1 }
        $env:ALLOW_SCHEMA_DROP = "1"
        go run ./cmd/migrate -command=drop
    }
    "db-fresh" {
        & $PSCommandPath migrate-drop
        & $PSCommandPath migrate-up
    }
    "seed" { Ensure-Go; go run ./cmd/seed }
    "run" {
        Ensure-Go
        if (-not (Test-Path (Join-Path $RootDir ".env"))) {
            Write-Error ".env missing. Run: .\scripts\nextpresskit.ps1 install"
        }
        $port = [int](Get-AppPort)
        Assert-PortFree $port
        go run ./cmd/api
    }
    "start" {
        Ensure-Go
        $rt = Get-DevRuntimeBasename
        $pidFile = Join-Path $RootDir ".tmp\$rt-api.pid"
        $logFile = Join-Path $RootDir ".tmp\$rt-api.log"
        New-Item -ItemType Directory -Force -Path (Split-Path $pidFile) | Out-Null
        if (Test-Path $pidFile) {
            $oldId = Get-Content $pidFile -ErrorAction SilentlyContinue
            if ($oldId -and (Get-Process -Id $oldId -ErrorAction SilentlyContinue)) {
                Write-Host "API already running (pid=$oldId)."
                exit 0
            }
            Remove-Item $pidFile -Force -ErrorAction SilentlyContinue
        }
        $port = [int](Get-AppPort)
        Assert-PortFree $port
        $goExe = (Get-Command go -ErrorAction Stop).Source
        $proc = Start-Process -FilePath $goExe -ArgumentList "run","./cmd/api" -WorkingDirectory $RootDir `
            -WindowStyle Hidden -RedirectStandardOutput $logFile -RedirectStandardError $logFile -PassThru
        Set-Content -Path $pidFile -Value $proc.Id
        Start-Sleep -Seconds 1
        if (Get-Process -Id $proc.Id -ErrorAction SilentlyContinue) {
            Write-Host "API started in background (pid=$($proc.Id))."
            Write-Host "Logs: $logFile"
        } else {
            Write-Host "API failed to start. See $logFile" -ForegroundColor Red
            Remove-Item $pidFile -Force -ErrorAction SilentlyContinue
            exit 1
        }
    }
    "stop" {
        $rt = Get-DevRuntimeBasename
        $pidFile = Join-Path $RootDir ".tmp\$rt-api.pid"
        if (Test-Path $pidFile) {
            $apiPid = Get-Content $pidFile
            $running = Get-Process -Id $apiPid -ErrorAction SilentlyContinue
            if ($running) {
                Stop-Process -Id $apiPid -Force -ErrorAction SilentlyContinue
                Write-Host "API stopped."
            } else {
                Write-Host "Stale pid file removed."
            }
            Remove-Item $pidFile -Force -ErrorAction SilentlyContinue
        } else {
            Write-Host "No .tmp\$rt-api.pid (nothing to stop from start)."
        }
    }
    "deploy" {
        & (Join-Path $RootDir "scripts\deploy.ps1")
    }
    "checks" {
        Ensure-Go
        go test ./...
        go vet ./...
        go test -tags=integration -v ./internal/platform/database
        go run github.com/getkin/kin-openapi/cmd/validate@latest (Join-Path $RootDir "docs\openapi.yaml")
        go run golang.org/x/vuln/cmd/govulncheck@latest ./...
    }
    "test" { Ensure-Go; go test -v ./... }
    "test-coverage" { Ensure-Go; go test -cover ./... }
    "test-integration" { Ensure-Go; go test -tags=integration -v ./internal/platform/database }
    "security-check" { Ensure-Go; go run golang.org/x/vuln/cmd/govulncheck@latest ./... }
    "postman-sync" {
        $py = Get-Command python3 -ErrorAction SilentlyContinue
        if (-not $py) { $py = Get-Command python -ErrorAction SilentlyContinue }
        if (-not $py) {
            Write-Error "Python 3 is required for postman-sync (install Python or use Git Bash: ./scripts/nextpresskit postman-sync)."
        }
        $extra = @()
        if ($args.Count -gt 1) { $extra = $args[1..($args.Count - 1)] }
        & $py.Source (Join-Path $RootDir "scripts\sync-postman.py") @extra
    }
    "clean" {
        Remove-Item -Recurse -Force (Join-Path $RootDir "bin") -ErrorAction SilentlyContinue
        go clean
    }
    default {
        Write-Host "Unknown command: $cmd" -ForegroundColor Red
        Show-Help
        exit 1
    }
}
