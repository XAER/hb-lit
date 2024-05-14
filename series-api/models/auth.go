package models

type LoginBodyStruct struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
