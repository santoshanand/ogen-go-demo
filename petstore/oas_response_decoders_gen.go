// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

func decodeAddPetResponse(resp *http.Response) (res Pet, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Pet
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "decode \"application/json\"")
			}
			if err := d.Skip(); err != io.EOF {
				return res, errors.New("unexpected trailing data")
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeDeletePetResponse(resp *http.Response) (res DeletePetOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return DeletePetOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeGetPetByIdResponse(resp *http.Response) (res GetPetByIdRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response Pet
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "decode \"application/json\"")
			}
			if err := d.Skip(); err != io.EOF {
				return res, errors.New("unexpected trailing data")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 404:
		// Code 404.
		return &GetPetByIdNotFound{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeUpdatePetResponse(resp *http.Response) (res UpdatePetOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return UpdatePetOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}
