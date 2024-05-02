package server_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/playwright-community/playwright-go"
	testify "github.com/stretchr/testify/assert"
)

// var testURL = "http://localhost:8080"

var testURL = os.Getenv("TEST_URL")

var pw *playwright.Playwright
var browser playwright.Browser
var browserCtx playwright.BrowserContext
var page playwright.Page

var headless = true
var exact = true
var err error

var assert *testify.Assertions

func InitializeAssert(t *testing.T) {
	assert = testify.New(t)
}

func InitializePlaywright() {
	pw, err = playwright.Run()

	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	browser, err = pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: &headless})

	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	browserCtx, err = browser.NewContext(playwright.BrowserNewContextOptions{BaseURL: &testURL})

	if err != nil {
		log.Fatalf("could not create context: %v", err)
	}

	page, err = browserCtx.NewPage()

	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		browserCtx.ClearCookies()

		return ctx, nil
	})

	sc.Given(`^I am on (.+) page$`, func(route string) {
		_, err = page.Goto(route)

		assert.NoError(err)
	})

	sc.Then(`^the page title should be (.+)$`, func(expected string) {
		actual, _ := page.Title()

		assert.Equal(expected, actual)
	})

	sc.Then(`^I fill the (.+) with (.+)$`, func(placeholder, value string) {
		element := page.GetByPlaceholder(placeholder, playwright.PageGetByPlaceholderOptions{Exact: &exact})

		err = element.Fill(value)

		assert.NoError(err)
	})

	sc.Then(`^I select the (.+) in (.+)$`, func(dropdown, value string) {
		element := page.Locator("#" + dropdown).GetByRole("textbox")

		err = element.Click()

		assert.NoError(err)

		element = page.GetByText(value)

		err = element.Click()

		assert.NoError(err)
	})

	sc.Then(`^I click the (.+) button$`, func(name string) {
		element := page.GetByRole("button", playwright.PageGetByRoleOptions{Name: name})

		err = element.Click()

		assert.NoError(err)
	})

	sc.Then(`^I see (.+) notification$`, func(expected string) {
		element := page.GetByText(expected)

		actual, _ := element.InnerText()

		assert.Equal(expected, actual)
	})
}

func TestAcceptanceServerFeatures(t *testing.T) {
	InitializeAssert(t)

	InitializePlaywright()

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
