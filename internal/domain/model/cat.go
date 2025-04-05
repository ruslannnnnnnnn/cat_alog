package model

import (
	"errors"
	"regexp"
	"time"
)

const urlRegex = "^(?:(https?|ftp):\\/\\/)?(?:www\\.)?(?:[a-zA-Z0-9-]+\\.)+[a-zA-Z]{2,}(?::\\d+)?(?:\\/[^\\s?#]*)?(?:\\?[^#\\s]*)?(?:#[^\\s]*)?$"

type Cat struct {
	Id          string
	Name        string
	DateOfBirth time.Time
	ImageUrl    string
}

func (c *Cat) IsValid() error {
	if c.ImageUrl == "" {
		return errors.New("cat has no image url")
	}

	regex := regexp.MustCompile(urlRegex)
	if !regex.MatchString(c.ImageUrl) {
		return errors.New("cat has invalid image url")
	}
	return nil
}
