// Package of handler aggregates abstration of
// function from model to controller/presentation
package handler

type Handler struct {
	WaitForKeyPress func() (int, error)
}

// Create a new Handler given a waitForKeyPress function.
func New(waitForKeyPress func() (int, error)) *Handler {
	return &Handler{
		WaitForKeyPress: waitForKeyPress,
	}
}

// Deep Copy a Handler
func Copy(handler *Handler) *Handler {
	return &Handler{
		WaitForKeyPress: handler.WaitForKeyPress,
	}
}

// Destroy a Handler.
func Destroy(handler *Handler) {
	handler.WaitForKeyPress = nil
	handler = nil
}
