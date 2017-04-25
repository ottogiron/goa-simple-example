// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "comics": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/ottogiron/goa-heroes/design
// --out=$(GOPATH)/src/github.com/ottogiron/goa-heroes
// --version=v1.2.0

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A comic hero (default view)
//
// Identifier: application/vnd.goa.example.heroe+json; view=default
type GoaExampleHeroe struct {
	// Unique heroe ID
	ID int `form:"id" json:"id" xml:"id"`
	// Name of hero
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GoaExampleHeroe media type instance.
func (mt *GoaExampleHeroe) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// DecodeGoaExampleHeroe decodes the GoaExampleHeroe instance encoded in resp body.
func (c *Client) DecodeGoaExampleHeroe(resp *http.Response) (*GoaExampleHeroe, error) {
	var decoded GoaExampleHeroe
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// GoaExampleHeroeCollection is the media type for an array of GoaExampleHeroe (default view)
//
// Identifier: application/vnd.goa.example.heroe+json; type=collection; view=default
type GoaExampleHeroeCollection []*GoaExampleHeroe

// Validate validates the GoaExampleHeroeCollection media type instance.
func (mt GoaExampleHeroeCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeGoaExampleHeroeCollection decodes the GoaExampleHeroeCollection instance encoded in resp body.
func (c *Client) DecodeGoaExampleHeroeCollection(resp *http.Response) (GoaExampleHeroeCollection, error) {
	var decoded GoaExampleHeroeCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}