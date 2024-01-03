package selenium

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

var _ = Describe("MyProject", func() {
	var (
		webDriver selenium.WebDriver
		service   *selenium.Service
		err       error
	)

	BeforeEach(func() {
		// Start the Selenium service using the local Chromedriver binary
		service, err = selenium.NewChromeDriverService("/tmp/chromedriver", 4444)
		Expect(err).NotTo(HaveOccurred())

		// Set up Chrome WebDriver
		caps := selenium.Capabilities{
			"browserName": "chrome",
		}
		chromeCaps := chrome.Capabilities{
			Args: []string{
				// Add any Chrome command-line flags you need here
			},
		}
		caps.AddChrome(chromeCaps)

		webDriver, err = selenium.NewRemote(caps, "")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		// Stop the Selenium service
		if service != nil {
			service.Stop()
		}

		// Close the WebDriver
		if webDriver != nil {
			webDriver.Quit()
		}
	})

	Context("when running tests", func() {
		It("should interact with Selenium", func() {
			// Your Selenium test code here
			err := webDriver.Get("https://www.example.com")
			Expect(err).NotTo(HaveOccurred())

			title, err := webDriver.Title()
			Expect(err).NotTo(HaveOccurred())
			Expect(title).To(Equal("Example Domain"))
		})
	})
})
