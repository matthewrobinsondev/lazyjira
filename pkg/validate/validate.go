package validate

import (
	"fmt"
	"net/url"
)

func IsValidUrl(v string) error {
	u, err := url.Parse(v)
	if err != nil {
		return err
	}

	if u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("URL %s is not valid", v)
	}

	return nil
}
