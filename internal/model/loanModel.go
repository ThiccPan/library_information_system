package model

import "time"

type CommandLoanRequest struct {
	Id            uint ``
	User_id       uint ``
	Book_id       uint ``
	LoanStatus_id uint `json:"status"`
	Deadline      time.Time
}

type QueryLoanRequest struct {
	Id          uint
	Search      string
	QueryParams map[string]any
}
