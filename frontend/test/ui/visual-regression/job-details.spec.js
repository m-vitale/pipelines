const assert = require('assert');

function assertDiffs(results) {
  results.forEach(r => assert.ok(r.isExactSameImage));
}

describe('view job details', () => {

  beforeAll(() => {
    browser.url('/');
  });

  it('opens job details on double click', () => {
    // Find a job with runs that can also be cloned. The 3rd job is one.
    // TODO: Explore making this more reliable
    const selector = 'app-shell job-list item-list #listContainer paper-item:nth-of-type(3)';

    browser.waitForVisible(selector);
    browser.doubleClick(selector);
    assertDiffs(browser.checkDocument());
  });

  it('can switch to run list tab', () => {
    const selector = 'app-shell job-details paper-tab:last-child';

    browser.click(selector);
    assertDiffs(browser.checkDocument());
  });

  it('loads next page on "next" button press', () => {
    const nextButtonselector =
        'app-shell job-details run-list item-list #nextPage';

    browser.click(nextButtonselector);
    assertDiffs(browser.checkDocument());
  });

  it('loads previous page on "previous" button press after pressing "next"', () => {
    const previousButtonselector =
        'app-shell job-details run-list item-list #previousPage';
    browser.click(previousButtonselector);

    assertDiffs(browser.checkDocument());
  });

  it('loads additional runs after changing page size', () => {
    // Default is 20, but we'll change it to 50.
    const pageSizeDropdownSelector =
        'app-shell job-details run-list item-list paper-dropdown-menu';
    browser.click(pageSizeDropdownSelector);

    const pageSizeSelector =
        'app-shell job-details run-list item-list ' +
        'paper-dropdown-menu::paper-item:nth-of-type(2)';
    browser.click(pageSizeSelector);

    assertDiffs(browser.checkDocument());
  });

  it('allows the list to be sorted. Defaults to ascending order', () => {
    // Sort by Run Name column (ascending)
    const runNameColumnButtonSelector =
        'app-shell job-details run-list item-list #header::div:nth-of-type(2)::paper-button';
    browser.click(runNameColumnButtonSelector);

    assertDiffs(browser.checkDocument());
  });

  it('sorts in descending order on second time a column is clicked', () => {
    // Sort by Run Name column (descending)
    // Sort will be descending now since it has already been clicked once in the previous test.
    const runNameColumnButtonSelector =
        'app-shell job-details run-list item-list #header::div:nth-of-type(2)::paper-button';

    browser.click(runNameColumnButtonSelector);

    // List should be reset to first page of results
    assertDiffs(browser.checkDocument());
  });

  it('allows the list to be filtered by Run name', () => {
    // Open up the filter box
    const filterButtonSelector = 'app-shell job-details run-list item-list paper-icon-button';
    browser.click(filterButtonSelector);

    const filterBoxSelector =
        'app-shell job-details run-list item-list #headerContainer::div:nth-of-type(2)::input';
    browser.setValue(filterBoxSelector, 'hello');

    assertDiffs(browser.checkDocument());
  });

  it('allows the list to be filtered and sorted', () => {
    // List is already filtered from previous test
    // Sort by run creation time column.
    const createdAtColumnButtonSelector =
        'app-shell job-details run-list item-list #header::div:nth-of-type(3)::paper-button';
    browser.click(createdAtColumnButtonSelector);

    assertDiffs(browser.checkDocument());

    // Clear filter by clicking the button again.
    const filterButtonSelector = 'app-shell job-details run-list item-list paper-icon-button';
    browser.click(filterButtonSelector);
  });

  it('populates new job on clone', () => {
    const cloneBtnSelector = 'app-shell job-details paper-button#cloneBtn';

    browser.click(cloneBtnSelector);
    assertDiffs(browser.checkDocument());
  });
});
