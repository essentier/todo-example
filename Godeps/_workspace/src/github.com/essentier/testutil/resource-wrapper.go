package testutil

import (
	"io"
	"net/http"

	"github.com/essentier/gopencils"
)

type ResWrapper struct {
	Resource   *gopencils.Resource
	ErrHandler RestErrorHandler
}

// Creates a new Resource.
func (rw *ResWrapper) NewChildResource(resourceName string) *ResWrapper {
	newRes := rw.Resource.NewChildResource(resourceName, nil)
	newRW := &ResWrapper{ErrHandler: rw.ErrHandler, Resource: newRes}
	return newRW
}

// Same as Res() Method, but returns a Resource with url resource/:id
func (rw *ResWrapper) NewChildIdResource(id string) *ResWrapper {
	newRes := rw.Resource.NewChildIdResource(id)
	newRW := &ResWrapper{ErrHandler: rw.ErrHandler, Resource: newRes}
	return newRW
}

// Sets QueryValues for current Resource
func (rw *ResWrapper) SetQuery(querystring map[string]string) *ResWrapper {
	rw.Resource.SetQuery(querystring)
	return rw
}

// Performs a GET request on given Resource
// Call SetQuery beforehand if you want to set the query string of the GET request.
func (rw *ResWrapper) Get(responseBody interface{}) *ResWrapper {
	rw.Resource.Response = responseBody
	_, err := rw.Resource.Get()
	if rw.ErrHandler != nil {
		rw.ErrHandler.HandleError(err, "REST GET failed.")
	}
	return rw
}

// Performs a HEAD request on given Resource
// Call SetQuery beforehand if you want to set the query string of the HEAD request.
func (rw *ResWrapper) Head() *ResWrapper {
	_, err := rw.Resource.Head()
	if rw.ErrHandler != nil {
		rw.ErrHandler.HandleError(err, "REST HEAD failed.")
	}
	return rw
}

// Performs a PUT request on given Resource.
// Accepts interface{} as parameter, will be used as payload.
func (rw *ResWrapper) Put(payload interface{}, responseBody interface{}) *ResWrapper {
	rw.Resource.Response = responseBody
	_, err := rw.Resource.Put(payload)
	if rw.ErrHandler != nil {
		rw.ErrHandler.HandleError(err, "REST PUT failed.")
	}
	return rw
}

// Performs a POST request on given Resource.
// Accepts interface{} as parameter, will be used as payload.
func (rw *ResWrapper) Post(payload interface{}, responseBody interface{}) *ResWrapper {
	rw.Resource.Response = responseBody
	_, err := rw.Resource.Post(payload)
	if rw.ErrHandler != nil {
		rw.ErrHandler.HandleError(err, "REST POST failed.")
	}
	return rw
}

// Performs a Delete request on given Resource.
// Call SetQuery beforehand if you want to set the query string of the DELETE request.
func (rw *ResWrapper) Delete(responseBody interface{}) *ResWrapper {
	rw.Resource.Response = responseBody
	_, err := rw.Resource.Delete()
	if rw.ErrHandler != nil {
		rw.ErrHandler.HandleError(err, "REST DELETE failed.")
	}
	return rw
}

// Performs a OPTIONS request on given Resource.
// Call SetQuery beforehand if you want to set the query string of the OPTIONS request.
func (rw *ResWrapper) Options(responseBody interface{}) *ResWrapper {
	rw.Resource.Response = responseBody
	_, err := rw.Resource.Options()
	if rw.ErrHandler != nil {
		rw.ErrHandler.HandleError(err, "REST OPTIONS failed.")
	}
	return rw
}

// Performs a PATCH request on given Resource.
// Accepts interface{} as parameter, will be used as payload.
func (rw *ResWrapper) Patch(payload interface{}, responseBody interface{}) *ResWrapper {
	rw.Resource.Response = responseBody
	_, err := rw.Resource.Patch(payload)
	if rw.ErrHandler != nil {
		rw.ErrHandler.HandleError(err, "REST PATCH failed.")
	}
	return rw
}

// Sets Payload for current Resource
func (rw *ResWrapper) SetPayload(args interface{}) io.Reader {
	return rw.Resource.SetPayload(args)
}

// Sets Headers
func (rw *ResWrapper) SetHeader(key string, value string) {
	rw.Resource.SetHeader(key, value)
}

// Overwrites the client that will be used for requests.
// For example if you want to use your own client with OAuth2
func (rw *ResWrapper) SetClient(c *http.Client) {
	rw.Resource.SetClient(c)
}
