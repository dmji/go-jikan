package retry

import (
	"errors"
	"net/http"
	"strconv"
	"time"
)

// BackoffStrategy defines the parameters for exponential backoff. This can be
// used to drive a retry loop for example.
type BackoffStrategy struct {
	InitialInterval int
	MaxInterval     int
	Exponent        float64
	MaxElapsedTime  int
}

// Config configures a retry policy.
type Config struct {
	// Strategy sets the algorithm to use for a retry loop. It can be one of:
	//   - "backoff": retry with exponential backoff and random jitter.
	//   - "none" or "": disables retries.
	Strategy              string
	Backoff               *BackoffStrategy
	RetryConnectionErrors bool
}

// PermanentError is an error that signals that some operation has terminally
// failed and should not be retried.
type PermanentError struct {
	cause error
}

// Permanent creates a PermanentError that signals to a retry loop that it
// should stop retrying an operation and return the underlying error.
func Permanent(cause error) error {
	if IsPermanentError(cause) {
		return cause
	}

	return &PermanentError{
		cause: cause,
	}
}

func (e *PermanentError) Error() string {
	return e.cause.Error()
}

func (e *PermanentError) Unwrap() error {
	return e.cause
}

// TemporaryError represents a retryable error and signals to a retry loop that
// an operation may be retried with an optional wait interval.
type TemporaryError struct {
	wait    time.Duration
	message string
}

// Temporary creates a TemporaryError that signals to a retry loop that an
// operation can be retried. The error may also carry details about how long to
// wait before retrying. This wait interval may be used to override the retry
// policy in use.
func Temporary(message string) error {
	return &TemporaryError{
		message: message,
	}
}

// TemporaryFromResponse creates a TemporaryError similar to Temporary but
// additionally parses the Retry-After header from a response to determine the
// wait interval before the next retry attempt.
func TemporaryFromResponse(message string, res *http.Response) error {
	return &TemporaryError{
		wait:    retryIntervalFromResponse(res),
		message: message,
	}
}

func (e *TemporaryError) Error() string {
	return e.message
}

// RetryAfter returns the time to wait before retrying the request. The zero
// value should be interpreted by retry loops to mean they should fallback on
// their default policy whether expenonential, constant backoff or something
// else. It does not mean that an operation should be retried immediately.
func (e *TemporaryError) RetryAfter() time.Duration {
	return e.wait
}

func retryIntervalFromResponse(res *http.Response) time.Duration {
	if res == nil {
		return 0
	}

	retryVal := res.Header.Get("retry-after")
	if retryVal == "" {
		return 0
	}

	parsedNumber, err := strconv.ParseInt(retryVal, 10, 64)
	if err == nil {
		if parsedNumber < 0 {
			return 0
		} else {
			return time.Duration(parsedNumber) * time.Second
		}
	}

	parsedDate, err := time.Parse(time.RFC1123, retryVal)
	if err == nil {
		delta := parsedDate.Sub(time.Now())
		if delta < 0 {
			return 0
		} else {
			return delta
		}
	}

	return 0
}

// IsPermanentError returns true if an error value is or contains a
// PermanentError in its chain of errors.
func IsPermanentError(err error) bool {
	if err == nil {
		return false
	}

	var pe *PermanentError
	return errors.As(err, &pe)
}

// IsTemporaryError returns true if an error value is or contains a
// TemporaryError in its chain of errors.
func IsTemporaryError(err error) bool {
	if err == nil {
		return false
	}

	var pe *TemporaryError
	return errors.As(err, &pe)
}
