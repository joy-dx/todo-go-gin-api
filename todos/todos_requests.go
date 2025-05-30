/*
 * Symbolic Boilerplate
 *
 * An API spec for managing todo lists
 *
 * API version: 1.0.0
 * Contact: contact@simonball.me
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package todos

type TodosIdURI struct {
	TodoId int64 `json:"id" uri:"TodoId" binding:"required"`
}

type TodosPostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type TodosPutByIdRequest struct {
	Title    string `json:"title,omitempty"`
	Content  string `json:"content"`
	Complete bool   `json:"complete"`
}
