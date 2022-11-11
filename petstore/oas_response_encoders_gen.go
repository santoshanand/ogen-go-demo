// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeAddPetResponse(response Pet, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	e := jx.GetEncoder()

	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil

}

func encodeDeletePetResponse(response DeletePetOK, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	return nil

}

func encodeGetPetByIdResponse(response GetPetByIdRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Pet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *GetPetByIdNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeUpdatePetResponse(response UpdatePetOK, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))
	return nil

}
