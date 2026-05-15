import { test, expect, Page } from '@playwright/test';

const uniqueEmail = () => `test-${Date.now()}-${Math.random().toString(36).slice(2, 8)}@example.com`;
const TEST_PASSWORD = 'testpass123';

async function loginAsNewUser(page: Page): Promise<string> {
  const email = uniqueEmail();
  await page.goto('/register');
  await page.fill('#email', email);
  await page.fill('#password', TEST_PASSWORD);
  await page.fill('#passwordConfirm', TEST_PASSWORD);
  await page.click('button[type="submit"]');
  await page.waitForURL('/timer');
  return email;
}

test.describe('Timer start/stop', () => {
  test.beforeEach(async ({ page }) => {
    await loginAsNewUser(page);
  });

  test('kan starte en timer', async ({ page }) => {
    await expect(page.locator('#timer-section')).toContainText('Ingen aktiv timer');

    await page.click('button:has-text("Start")');
    await expect(page.locator('#timer-section')).toContainText('Timer aktiv');
    await expect(page.locator('#timer-section')).toContainText('Stopp');
  });

  test('kan stoppe en timer', async ({ page }) => {
    await page.click('button:has-text("Start")');
    await expect(page.locator('#timer-section')).toContainText('Timer aktiv');

    await page.click('button:has-text("Stopp")');
    await expect(page.locator('#timer-section')).toContainText('Ingen aktiv timer');
  });

  test('stoppad timer vises i listen', async ({ page }) => {
    await page.click('button:has-text("Start")');
    await expect(page.locator('#timer-section')).toContainText('Timer aktiv');

    // Wait a moment for the timer to have some duration
    await page.waitForTimeout(1000);

    await page.click('button:has-text("Stopp")');
    await expect(page.locator('#timer-section')).toContainText('Ingen aktiv timer');

    // Should appear in list
    await expect(page.locator('#timings-list .bg-white')).toHaveCount(1);
  });
});

test.describe('Beskrivelse', () => {
  test('kan legge til beskrivelse på aktiv timing', async ({ page }) => {
    await loginAsNewUser(page);

    await page.click('button:has-text("Start")');
    await expect(page.locator('#timer-section')).toContainText('Timer aktiv');

    // Type description on active timing using pressSequentially to trigger input events
    const descInput = page.locator('#timer-section input[name="description"]');
    await descInput.pressSequentially('Jobbe med prosjekt', { delay: 20 });

    // Wait for HTMX to save (debounced at 500ms + roundtrip)
    await page.waitForResponse(resp => resp.url().includes('/edit-description'), { timeout: 5000 });

    // Stop timer
    await page.click('button:has-text("Stopp")');

    // Check that description is in the list
    await expect(page.locator('#timings-list')).toContainText('Jobbe med prosjekt');
  });
});

test.describe('Tagger', () => {
  test('kan legge til og fjerne tagger', async ({ page }) => {
    await loginAsNewUser(page);

    // Start and stop a timer
    await page.click('button:has-text("Start")');
    await expect(page.locator('#timer-section')).toContainText('Timer aktiv');

    // Add a tag via the tag input in timer section
    const tagInput = page.locator('#timer-section input[name="tag"]');
    await tagInput.fill('arbeid');
    await page.locator('#timer-section button:has-text("+")').click();

    // Tag should appear
    await expect(page.locator('#timer-section')).toContainText('arbeid');

    // Stop timer
    await page.click('button:has-text("Stopp")');

    // Tag should be in the list item
    await expect(page.locator('#timings-list')).toContainText('arbeid');

    // Remove the tag
    const removeBtn = page.locator('#timings-list .bg-blue-50 button').first();
    await removeBtn.click();

    // Tag should be removed
    await expect(page.locator('#timings-list .bg-blue-50')).toHaveCount(0);
  });
});

test.describe('Redigering', () => {
  test('kan redigere beskrivelse på eksisterende timing', async ({ page }) => {
    await loginAsNewUser(page);

    // Create a timing
    await page.click('button:has-text("Start")');
    await page.waitForTimeout(500);
    await page.click('button:has-text("Stopp")');

    // Click edit button on the timing item
    await page.locator('#timings-list button:has-text("✏️")').first().click();

    // Fill in description using pressSequentially to trigger real input events for HTMX
    const descInput = page.locator('#timings-list .edit-form input[name="description"]').first();
    await descInput.pressSequentially('Oppdatert beskrivelse', { delay: 20 });

    // Wait for HTMX to send the debounced request and swap the DOM
    await page.waitForResponse(resp => resp.url().includes('/edit-description'), { timeout: 5000 });

    // Verify description is updated
    await expect(page.locator('#timings-list').first()).toContainText('Oppdatert beskrivelse');
  });
});

test.describe('Slett timing', () => {
  test('kan slette en timing', async ({ page }) => {
    await loginAsNewUser(page);

    // Create a timing
    await page.click('button:has-text("Start")');
    await page.waitForTimeout(500);
    await page.click('button:has-text("Stopp")');

    await expect(page.locator('#timings-list .bg-white')).toHaveCount(1);

    // Set up dialog handler before clicking delete
    page.once('dialog', dialog => dialog.accept());

    // Remove hx-confirm to avoid dialog timing issues and click delete
    const deleteBtn = page.locator('#timings-list button[title="Slett"]').first();
    await deleteBtn.evaluate(el => el.removeAttribute('hx-confirm'));

    // Click delete button and wait for response
    const [deleteResp] = await Promise.all([
      page.waitForResponse(resp => resp.url().includes('/api/timings/'), { timeout: 5000 }),
      deleteBtn.click(),
    ]);

    // Timing should be removed
    await expect(page.locator('#timings-list .bg-white')).toHaveCount(0);
  });
});
