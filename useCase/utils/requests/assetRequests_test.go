package requests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeRequestAsset(t *testing.T) {
	tests := []struct {
		name       string
		assetType  string
		assetCode  string
		mockStatus int
		wantErr    error
	}{
		{
			name:       "valid COIN request",
			assetType:  "COIN",
			assetCode:  "USD-BRL",
			mockStatus: http.StatusOK,
			wantErr:    nil,
		},
		{
			name:       "valid CRYPTO request",
			assetType:  "CRYPTO",
			assetCode:  "BTC",
			mockStatus: http.StatusInternalServerError,
			wantErr:    errors.New("status code diferente de 200: 400"),
		},
		{
			name:       "invalid asset type",
			assetType:  "STOCK",
			assetCode:  "AAPL",
			mockStatus: http.StatusOK,
			wantErr:    errors.New("erro ao construir a URL"),
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockStatus)
			}))
			defer server.Close()

			_, err := MakeRequestAsset(tt.assetType, tt.assetCode)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}

		})
	}
}

func TestGetValueFromCoin(t *testing.T) {
	tests := []struct {
		name     string
		response string
		code     string
		want     float64
		wantErr  error
	}{
		{
			name:     "valid coin response",
			response: `{"USDBRL":{"bid":"5.25"}}`,
			code:     "USD-BRL",
			want:     5.25,
			wantErr:  nil,
		},
		{
			name:     "invalid coin response",
			response: "",
			code:     "USD-BRL",
			want:     0,
			wantErr:  errors.New("moeda inválida"),
		},
		{
			name:     "zero value coin",
			response: `{"USDBRL":{"bid":"0"}}`,
			code:     "USD-BRL",
			want:     0,
			wantErr:  errors.New("valor do ativo é zero"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetValueFromCoin(tt.response, tt.code)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetValueFromCrypto(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     float64
		wantErr  error
	}{
		{
			name:     "valid crypto response",
			response: `[{"last":"45000.00"}]`,
			want:     45000.00,
			wantErr:  nil,
		},
		{
			name:     "invalid crypto response",
			response: "",
			want:     0,
			wantErr:  errors.New("crypto inválida"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetValueFromCrypto(tt.response)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
