#!/usr/bin/env bash
set -euo pipefail

root=$(git rev-parse --show-toplevel)

if [[ ! -f "$root/.devcontainer/copilot-ok" ]]; then
  echo '🔍 Sjekker GitHub Copilot-tilgang...'
  if copilot -p 'hva er 1+1?' --model gpt-4o-mini > /dev/null 2>&1; then
    echo '✅ GitHub Copilot-tilgang er OK.'
    touch "$root/.devcontainer/copilot-ok"
  else
    echo '❌  Ingen GitHub Copilot-tilgang funnet. Se kurs/01-start.md for instruksjoner.'
  fi
else
  echo '✅ GitHub Copilot-tilgang er OK (cache).'
fi

echo '📦 Installerer opencode og playwright mcp...'
npm install opencode-ai @playwright/mcp
