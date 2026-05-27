package helpers

func FindIdx(matcher string) func(s string) bool {
	return func(s string) bool {
		return s == matcher
	}
}
