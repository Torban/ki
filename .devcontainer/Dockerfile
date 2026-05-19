FROM mcr.microsoft.com/devcontainers/go:1

# installer nodejs, npm, fish shell, rg, fd-find, opencode og playwright mcp
RUN apt update \
  && apt install -y nodejs npm fish ripgrep fd-find \
  && ln -s $(which fdfind) /usr/local/bin/fd \
  && npm install --global opencode-ai @playwright/mcp
