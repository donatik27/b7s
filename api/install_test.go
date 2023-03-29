package api_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"

	"github.com/blocklessnetworking/b7s/api"
	"github.com/blocklessnetworking/b7s/models/api/request"
	"github.com/blocklessnetworking/b7s/testing/mocks"
)

func TestAPI_FunctionInstall(t *testing.T) {
	t.Run("nominal case", func(t *testing.T) {
		t.Parallel()

		api := setupAPI(t)

		req := request.InstallFunction{
			URI: "dummy-function-id",
			CID: "dummy-cid",
		}

		rec, ctx, err := setupRecorder(installEndpoint, req)
		require.NoError(t, err)

		err = api.Install(ctx)
		require.NoError(t, err)

		require.Equal(t, http.StatusOK, rec.Result().StatusCode)
	})
}

func TestAPI_FunctionInstall_HandlesErrors(t *testing.T) {
	t.Run("missing URI and CID", func(t *testing.T) {
		t.Parallel()

		api := setupAPI(t)

		req := request.InstallFunction{
			URI: "",
			CID: "",
		}

		_, ctx, err := setupRecorder(installEndpoint, req)
		require.NoError(t, err)

		err = api.Install(ctx)
		require.Error(t, err)

		echoErr, ok := err.(*echo.HTTPError)
		require.True(t, ok)

		require.Equal(t, http.StatusBadRequest, echoErr.Code)
	})
	t.Run("node install takes too long", func(t *testing.T) {
		t.Parallel()

		const (
			// The API times out after 10 seconds.
			installDuration = 11 * time.Second
		)

		node := mocks.BaselineNode(t)
		node.FunctionInstallFunc = func(context.Context, string, string) error {
			time.Sleep(installDuration)
			return nil
		}

		api := api.New(mocks.NoopLogger, node)

		req := request.InstallFunction{
			URI: "dummy-uri",
			CID: "dummy-cid",
		}

		rec, ctx, err := setupRecorder(installEndpoint, req)
		require.NoError(t, err)

		err = api.Install(ctx)
		require.NoError(t, err)

		require.Equal(t, http.StatusRequestTimeout, rec.Result().StatusCode)
	})
	t.Run("node fails to install function", func(t *testing.T) {
		t.Parallel()

		node := mocks.BaselineNode(t)
		node.FunctionInstallFunc = func(context.Context, string, string) error {
			return mocks.GenericError
		}

		api := api.New(mocks.NoopLogger, node)

		req := request.InstallFunction{
			URI: "dummy-uri",
			CID: "dummy-cid",
		}

		_, ctx, err := setupRecorder(installEndpoint, req)
		require.NoError(t, err)

		err = api.Install(ctx)
		require.Error(t, err)

		echoErr, ok := err.(*echo.HTTPError)
		require.True(t, ok)

		require.Equal(t, http.StatusInternalServerError, echoErr.Code)
	})
}

func TestAPI_InstallFunction_HandlesMalformedRequests(t *testing.T) {

	api := setupAPI(t)

	const (
		wrongFieldType = `
		{
			"uri": "dummy-uri",
			"cid": 14
		}`

		unclosedBracket = `
		{
			"uri": "dummy-uri",
			"cid": "dummy-cid"
		`

		validJSON = `
		{
			"uri": "dummy-uri",
			"cid": "dummy-cid"
		}`
	)

	tests := []struct {
		name        string
		payload     []byte
		contentType string
	}{
		{
			name:        "wrong field type",
			payload:     []byte(wrongFieldType),
			contentType: echo.MIMEApplicationJSON,
		},
		{
			name:        "malformed JSON",
			payload:     []byte(unclosedBracket),
			contentType: echo.MIMEApplicationJSON,
		},
		{
			name:    "valid JSON with no MIME type",
			payload: []byte(validJSON),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			prepare := func(req *http.Request) {
				req.Header.Set(echo.HeaderContentType, test.contentType)
			}

			_, ctx, err := setupRecorder(installEndpoint, test.payload, prepare)
			require.NoError(t, err)

			err = api.Install(ctx)
			require.Error(t, err)

			echoErr, ok := err.(*echo.HTTPError)
			require.True(t, ok)

			require.Equal(t, http.StatusBadRequest, echoErr.Code)
		})
	}
}
