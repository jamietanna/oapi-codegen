// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package models

// Error defines model for Error.
type Error struct {
	// Code Error code
	Code int32 `json:"code"`

	// Message Error message
	Message string `json:"message"`
}

// NewPet defines model for NewPet.
type NewPet struct {
	// Name Name of the pet
	Name string `json:"name"`

	// Tag Type of the pet
	Tag *string `json:"tag,omitempty"`
}

// Pet defines model for Pet.
type Pet struct {
	// Id Unique id of the pet
	Id int64 `json:"id"`

	// Name Name of the pet
	Name string `json:"name"`

	// Tag Type of the pet
	Tag *string `json:"tag,omitempty"`
}

// FindPetsResponseBody200 defines model for the response body for GET /pets (200, application/json).
type FindPetsResponseBody200 = []Pet

// FindPetsResponseBodyDefault defines model for the response body for GET /pets (default, application/json).
type FindPetsResponseBodyDefault = Error

// AddPetRequestBody defines model for the request body for POST /pets (application/json).
type AddPetRequestBody = NewPet

// AddPetResponseBody200 defines model for the response body for POST /pets (200, application/json).
type AddPetResponseBody200 = Pet

// AddPetResponseBodyDefault defines model for the response body for POST /pets (default, application/json).
type AddPetResponseBodyDefault = Error

// DeletePetResponseBodyDefault defines model for the response body for DELETE /pets/{id} (default, application/json).
type DeletePetResponseBodyDefault = Error

// FindPetByIDResponseBody200 defines model for the response body for GET /pets/{id} (200, application/json).
type FindPetByIDResponseBody200 = Pet

// FindPetByIDResponseBodyDefault defines model for the response body for GET /pets/{id} (default, application/json).
type FindPetByIDResponseBodyDefault = Error

// FindPetsParams defines parameters for FindPets.
type FindPetsParams struct {
	// Tags tags to filter by
	Tags *[]string `form:"tags,omitempty" json:"tags,omitempty"`

	// Limit maximum number of results to return
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

// AddPetJSONRequestBody defines body for AddPet for application/json ContentType.
type AddPetJSONRequestBody = NewPet
