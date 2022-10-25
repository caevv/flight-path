package server

import (
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/steinfletcher/apitest"
)

func TestAPICall(t *testing.T) {
	apitest.New().
		HandlerFunc(FiberToHandlerFunc(New(fiber.New()).fiber)).
		Post("/calculate").
		Body(`[["SFO", "EWR"]]`).
		Expect(t).
		Body(`["SFO", "EWR"]`).
		Status(http.StatusOK).
		End()
}

func TestAPICallMultiplePaths(t *testing.T) {
	apitest.New().
		HandlerFunc(FiberToHandlerFunc(New(fiber.New()).fiber)).
		Post("/calculate").
		Body(`[["ATL", "EWR"], ["SFO", "ATL"]]`).
		Expect(t).
		Body(`["SFO", "EWR"]`).
		Status(http.StatusOK).
		End()
}

func TestAPICallInvalid(t *testing.T) {
	apitest.New().
		HandlerFunc(FiberToHandlerFunc(New(fiber.New()).fiber)).
		Post("/calculate").
		Body(`[["IND", "EWR"], ["EWR", "IND"]]`).
		Expect(t).
		Body(`failed to calculate: invalid path`).
		Status(http.StatusInternalServerError).
		End()
}

func FiberToHandlerFunc(app *fiber.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := app.Test(r)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// copy headers
		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)

		// copy body
		if _, err := io.Copy(w, resp.Body); err != nil {
			panic(err)
		}
	}
}
