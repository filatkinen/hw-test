package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:"`
	}

	Token struct {
		Header    []byte `validate:"min:18|max:50"`
		Payload   []byte `validate:"in:23,45,78"`
		Signature []byte `validate:"minus:10"`
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     "272",
				Name:   "User1",
				Age:    75,
				Email:  "fen@ya.ru",
				Role:   "admin",
				Phones: []string{"12345678901", "12345678901"},
				meta:   nil,
			},
			expectedErr: ValidationErrors{
				{
					Field: "ID",
					Err:   ErrorValidation,
				},
				{
					Field: "Age",
					Err:   ErrorValidation,
				},
			},
		},
		{
			in: App{Version: "one"},
			expectedErr: ValidationErrors{{
				Field: "Version",
				Err:   ErrorParsValidateTag,
			}},
		},
		{
			in: Token{
				Header:    []byte{7, 19, 100},
				Payload:   []byte{23, 89},
				Signature: []byte{64},
			},
			expectedErr: ValidationErrors{
				{
					Field: "Header",
					Err:   ErrorValidation,
				},
				{
					Field: "Payload",
					Err:   ErrorValidation,
				},
				{
					Field: "Signature",
					Err:   ErrorNotImplemented,
				},
			},
		},
		{
			in: Response{
				Code: 403,
			},
			expectedErr: ValidationErrors{
				{
					Field: "Code",
					Err:   ErrorValidation,
				},
			},
		},
		{
			in:          []byte{},
			expectedErr: ErrorNotStructure,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			err := Validate(tt.in)
			var e ValidationErrors
			if errors.As(err, &e) {
				fmt.Println(err)
				var v ValidationErrors
				if errors.As(tt.expectedErr, &v) {
					require.Equal(t, len(e), len(v))
					for k := 0; k < len(e); k++ {
						require.ErrorIs(t, e[k].Err, v[k].Err)
					}
				}
			} else {
				fmt.Println(err)
				require.ErrorIs(t, err, tt.expectedErr)
			}
		})
	}
}
