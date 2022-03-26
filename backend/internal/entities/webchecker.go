package entities

import "exam.com/webchecker/pkg/null"

type Websites []string

type WebsiteStatus struct {
	Name   string          `json:"name"`
	OK     bool            `json:"ok"`
	ErrMsg null.NullString `json:"err_msg"`
}
