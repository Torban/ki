import { defineConfig } from '@playwright/test';

export default defineConfig({
  testDir: './tests/e2e',
  timeout: 30000,
  retries: 0,
  use: {
    baseURL: process.env.BASE_URL || 'http://127.0.0.1:8091',
    headless: true,
    screenshot: 'only-on-failure',
    trace: 'on-first-retry',
  },
  projects: [
    {
      name: 'chromium',
      use: { browserName: 'chromium' },
    },
  ],
  webServer: {
    command: 'scripts/dev-test.sh',
    url: 'http://127.0.0.1:8091/api/health',
    reuseExistingServer: !!process.env.BASE_URL,
    timeout: 30000,
  },
});
