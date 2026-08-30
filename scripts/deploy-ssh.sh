#!/usr/bin/env bash
set -euo pipefail

# Manual deploy of the SSH portfolio server to an EC2 host.
#
# This mirrors what .github/workflows/deploy.yml does, but runs from your local
# machine — useful for a one-off or when GitHub Actions isn't set up yet.
#
# Requirements:
#   - SSH access to the EC2 host (ssh <user>@<host> must work / bashrc)
#   - Go toolchain
#   - sudo on the remote host for systemd
#
# Usage:
#   scripts/deploy-ssh.sh <user>@<host> [ssh-port]
#
# Example:
#   scripts/deploy-ssh.sh ubuntu@habibiahmada.dev 22

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

if [ "$#" -lt 1 ]; then
  echo "Usage: $0 <user>@<host> [ssh-port]" >&2
  exit 1
fi

HOST="$1"
SSH_PORT="${2:-22}"
REMOTE_DIR="/opt/habibiahmada"

echo "==> Building SSH server binary"
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o habibiahmada-ssh ./cmd/ssh

echo "==> Copying binary + systemd unit to $HOST"
ssh -p "$SSH_PORT" "$HOST" "sudo mkdir -p '$REMOTE_DIR' && sudo chown \$USER '$REMOTE_DIR'"
scp -P "$SSH_PORT" \
  habibiahmada-ssh \
  deploy/portfolio-ssh.service \
  "$HOST:$REMOTE_DIR/"

echo "==> Installing & starting service on $HOST"
ssh -p "$SSH_PORT" "$HOST" "sudo bash -s" <<'REMOTE'
set -euo pipefail
cd /opt/habibiahmada
mkdir -p .ssh
chmod 755 habibiahmada-ssh

# Provision the Wish SSH host key if it does not exist yet.
if [ ! -f .ssh/term_info_ed25519 ]; then
  echo "  Generating Wish SSH host key (.ssh/term_info_ed25519)"
  ssh-keygen -t ed25519 -f .ssh/term_info_ed25519 -N "" -C "wish-portfolio@habibiahmada.dev"
  chmod 600 .ssh/term_info_ed25519*
fi

sudo cp portfolio-ssh.service /etc/systemd/system/portfolio-ssh.service
sudo systemctl daemon-reload
sudo systemctl enable portfolio-ssh.service || true
sudo systemctl restart portfolio-ssh.service
sleep 2
sudo systemctl --no-pager --lines=20 status portfolio-ssh.service || true

if ss -ltn | grep -q ':2222'; then
  echo "OK: SSH portfolio listening on :2222"
else
  echo "WARN: port 2222 not detected — check logs" >&2
  exit 1
fi
REMOTE

echo ""
echo "Deploy complete. Test with: ssh -p 2222 <some-user>@$HOST"
