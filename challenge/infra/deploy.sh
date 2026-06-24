#!/usr/bin/env bash
set -euo pipefail

HOST="10.21.4.31"
PORT="2204"
SSH_USER="release"
SSH_PASSWORD="Qm8p!2zT5r"
TARGET="/srv/mesh-flow"

sshpass -p "$SSH_PASSWORD" ssh -p "$PORT" -o StrictHostKeyChecking=no "$SSH_USER@$HOST" "cd $TARGET && ./rollout"
