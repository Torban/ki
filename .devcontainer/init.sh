#!/usr/bin/env bash
set -euo pipefail

root=$(git rev-parse --show-toplevel)
failure=false

if [[ ! -f "$root/.devcontainer/copilot-ok" ]]; then
  echo '🔍 Sjekker GitHub Copilot-tilgang...'
  if copilot --prompt 'hva er 1+1?' --model claude-sonnet-4.6 --disable-builtin-mcps --no-custom-instructions &> "$root/.devcontainer/copilot-logg" \
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

opencode_ok=false
playwright_mcp_ok=false
command -v opencode &> /dev/null && opencode_ok=true
npm list -g --depth=0 @playwright/mcp &> /dev/null && playwright_mcp_ok=true

if [[ "$opencode_ok" == true && "$playwright_mcp_ok" == true ]]; then
  echo '✅ opencode og Playwright MCP er allerede installert.'
else
  echo '📦 Installerer opencode og playwright mcp...'
  if npm install --global opencode-ai @playwright/mcp &> "$root/.devcontainer/npm-install-logg"; then
    echo '✅ opencode og Playwright MCP er installert.'
  else
    echo '❌  Feil ved installasjon av opencode eller Playwright MCP. Logg i .devcontainer/npm-install-logg'
    failure=true
  fi
fi

echo "⏳ Installerer rg og fd-find..."
if sudo apt-get update && sudo apt-get install -y ripgrep fd-find &> "$root/.devcontainer/apt-install-logg"; then
  echo '✅ rg og fd-find er installert.'
else
  echo '❌  Feil ved installasjon av rg eller fd-find. Logg i .devcontainer/apt-install-logg'
  failure=true
fi

if [[ "$failure" == true ]]; then
  echo '⚠️  init.sh feilet.'
  exit 1
else
  echo '✅ init.sh fullført uten feil.'
fi
