package model

import (
	"time"
)

type Currency struct {
	Id    int
	Title string
	Code  string
	Value float64
	Date  time.Time
}

type CurrencyXml struct {
	Date       string `xml:"date"`
	Currencies []struct {
		Code  string  `xml:"title"`
		Title string  `xml:"fullname"`
		Value float64 `xml:"description"`
	} `xml:"item"`
}

type SaveCurrency struct {
	Date  time.Time
	Code  string
	Title string
	Value float64
}

type SaveCurrencyResponse struct {
	Success bool `json:"success"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type GetCurrencyResponse struct {
	Title string  `json:"title"`
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}
