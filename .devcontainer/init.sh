#!/usr/bin/env bash
set -euo pipefail

root=$(git rev-parse --show-toplevel)
failure=false

if [[ ! -f "$root/.devcontainer/copilot-ok" ]]; then
  echo '🔍 Sjekker GitHub Copilot-tilgang...'
  if copilot -p 'hva er 1+1?' --model claude-sonnet-4.6 &> "$root/.devcontainer/copilot-logg" \
    && ! grep -i -E '(error|fail)' "$root/.devcontainer/copilot-logg" &> /dev/null \
    && grep 2 "$root/.devcontainer/copilot-logg" &> /dev/null
  then
    echo '✅ GitHub Copilot-tilgang er OK.'
    touch "$root/.devcontainer/copilot-ok"
  else
    grep -i -E '(error|fail)' "$root/.devcontainer/copilot-logg" || true
    echo '❌  Ingen GitHub Copilot-tilgang funnet. Se kurs/01-start.md for instruksjoner. Logg i .devcontainer/copilot-logg'
    failure=true
  fi
else
  echo '✅ GitHub Copilot-tilgang er OK (cache).'
fi

echo '📦 Installerer opencode og playwright mcp...'
if npm install --global opencode-ai @playwright/mcp &> "$root/.devcontainer/npm-install-logg"; then
  echo '✅ opencode og Playwright MCP er installert.'
else
  echo '❌  Feil ved installasjon av opencode eller Playwright MCP. Logg i .devcontainer/npm-install-logg'
  failure=true
fi

if [[ "$failure" == true ]]; then
  echo '⚠️  init.sh feilet.'
  exit 1
else
  echo '✅ init.sh fullført uten feil.'
fi
