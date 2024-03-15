package validation

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

var (
	ErrOriginPath     = fmt.Errorf("path not allowed")
	ErrOriginQuery    = fmt.Errorf("query not allowed")
	ErrOriginFragment = fmt.Errorf("fragment not allowed")
)

// NewOrigin creates a new Origin from a string. It normalizes the origin and returns an error if the origin is invalid.
func NewOrigin(s string) (Origin, error) {
	url, err := url.Parse(s)
	if err != nil {
		return "", fmt.Errorf("invalid origin: %w", err)
	}
	var errs []error
	if url.Path != "" && url.Path != "/" {
		errs = append(errs, ErrOriginPath)
	}
	if url.RawQuery != "" {
		errs = append(errs, ErrOriginQuery)
	}
	if url.Fragment != "" {
		errs = append(errs, ErrOriginFragment)
	}
	if len(errs) > 0 {
		return "", fmt.Errorf("invalid origin: %w", errors.Join(errs...))
	}
	return Origin(url.Scheme + "://" + url.Host), nil
}

// Origin represents a normalized origin.
type Origin string

// String returns the string representation of the origin.
func (o Origin) String() string {
	return string(o)
}

// Normalize removes trailing slashes and converts the origin to lowercase.
func (o Origin) Normalize() Origin {
	s := o.String()
	s = strings.ToLower(s)
	s = strings.TrimRight(s, "/")
	return Origin(s)
}

// Matches uses the origin as a pattern to match against the given origin.
func (o Origin) Matches(origin string) bool {
	pattern := o.String()
	if pattern == "*" {
		return true
	}

	prefix, suffix, hasWildcard := strings.Cut(pattern, "*")
	if !hasWildcard {
		return origin == pattern
	}

	if len(origin) <= len(prefix+suffix) {
		return false
	}
	if !strings.HasPrefix(origin, prefix) {
		return false
	}
	if !strings.HasSuffix(origin, suffix) {
		return false
	}
	return true
}
