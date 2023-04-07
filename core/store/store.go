package store

type Store interface {
	// Set sets the digits for the captcha id.
	Set(id string, value string) error
	// Get returns stored digits for the captcha id. Clear indicates
	// whether the captcha must be deleted from the store.
	Get(id string, clear bool) string
	//Verify captcha's answer directly
	Verify(id, answer string, clear bool) bool
}
