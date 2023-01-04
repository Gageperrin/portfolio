package techpalace
import "strings"


// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) (msg string) {
	msg = "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	return
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) (msg string) {
	stars := strings.Repeat("*", numStarsPerLine)
	msg = stars + "\n" + welcomeMsg + "\n" + stars
	return
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) (msg string) {
	msg = strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
	return 
}
