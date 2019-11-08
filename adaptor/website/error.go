// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package website

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	Message    string    `json:"message"`
	Extensions Extension `json:"extensions"`
}

type Extension struct {
	Code uint `json:"code"`
}
