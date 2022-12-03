package dto

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type ContentHeaders struct {
	data map[string][]string
}

type EmailInfo struct {
	Name  string
	Email string
}

func (headers *ContentHeaders) Has(key string) bool {
	_, ok := headers.data[key]

	return ok
}

func (headers *ContentHeaders) Get(key string) ([]string, error) {
	if val, ok := headers.data[key]; ok {
		return val, nil
	}

	return []string{}, errors.New(fmt.Sprintf("Header '%s' not found", key))
}

func (headers *ContentHeaders) GetOne(key string) (string, error) {
	values, err := headers.Get(key)
	if err == nil && len(values) > 0 {
		return values[0], nil
	}

	return "", errors.New(fmt.Sprintf("Header '%s' not found", key))
}

func (headers *ContentHeaders) Set(key string, val []string) *ContentHeaders {
	headers.data[key] = val

	return headers
}

func (headers *ContentHeaders) All() map[string][]string {
	return headers.data
}

func getEmailInfoFromString(item string) *EmailInfo {
	item = strings.TrimSpace(item)
	match := regexp.MustCompile(`^(?U)(.*)\s*<(.*@.*)>$`).FindStringSubmatch(item)
	if len(match) == 3 {
		return &EmailInfo{
			Name:  match[1],
			Email: match[2],
		}
	}
	match = regexp.MustCompile(`^[^\s]*@[^\s]*$`).FindStringSubmatch(item)
	if len(match) == 1 {
		return &EmailInfo{
			Name:  "",
			Email: match[0],
		}
	}

	return nil
}

func (headers *ContentHeaders) From() (EmailInfo, error) {
	from, err := headers.GetOne("From")
	if err == nil {
		return *getEmailInfoFromString(from), nil
	}

	return EmailInfo{}, errors.New("form not found")
}

func (headers *ContentHeaders) To() []EmailInfo {
	var result []EmailInfo
	toList, err := headers.Get("To")
	if err == nil {
		for _, toRow := range toList {
			toItems := strings.Split(toRow, ",")
			for _, toItem := range toItems {
				item := getEmailInfoFromString(toItem)
				if item != nil {
					result = append(result, *item)
				}
			}
		}
	}

	return result
}

func (headers *ContentHeaders) DateUTC() (time.Time, error) {
	date, err := headers.GetOne("DateUTC")
	if err != nil {
		return time.Time{}, errors.New("date not found")
	}

	t, err := time.Parse("Mon, _2 Jan 2006 15:04:05 -0700", date)

	if err != nil {
		return time.Time{}, errors.New("date has not valid format")
	}

	return t.UTC(), nil

}
