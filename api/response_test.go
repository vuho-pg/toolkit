package api

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/vuho-pg/toolkit/pagination"
	"github.com/vuho-pg/toolkit/test"
	"net/http"
	"testing"
)

func Test_transformMessage(t *testing.T) {
	cases := []test.Case{
		{"with message", transformMessage(http.StatusOK, "message"), "message"},
		{"200 no message", transformMessage(http.StatusOK, ""), "success"},
		{"400 no message", transformMessage(http.StatusBadRequest, ""), "bad request"},
		{"500 no message", transformMessage(http.StatusInternalServerError, ""), "error"},
		{"other", transformMessage(http.StatusBadGateway, ""), "status code 502"},
	}
	t.Parallel()
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Input, tc.Expect)
		})
	}
}

func TestBadRequest(t *testing.T) {
	req := BadRequest("bad request")
	reqJson, err := json.Marshal(req)
	assert.Nil(t, err)
	expect := Response{
		Meta: Metadata{
			Code:    http.StatusBadRequest,
			Message: "bad request",
		},
	}

	assert.Equal(t, req, expect)
	assert.True(t, test.CompareToJSON(expect, string(reqJson)))
}

func TestSuccessPagination(t *testing.T) {
	req := SuccessPagination("example data", pagination.Pagination{
		Page:        1,
		TotalPage:   10,
		TotalRecord: 100,
		Limit:       10,
	},
		"hello world")
	reqJson, err := json.Marshal(req)
	assert.Nil(t, err)
	expect := Response{
		Meta: Metadata{
			Code:    http.StatusOK,
			Message: "hello world",
			Pagination: &pagination.Pagination{
				Page:        1,
				TotalPage:   10,
				TotalRecord: 100,
				Limit:       10,
			},
		},
		Data: "example data",
	}

	assert.Equal(t, req, expect)
	assert.True(t, test.CompareToJSON(expect, string(reqJson)))
}

func TestSuccess(t *testing.T) {
	req := Success("example data", "hello world")
	reqJson, err := json.Marshal(req)
	assert.Nil(t, err)
	expect := Response{
		Meta: Metadata{
			Code:    http.StatusOK,
			Message: "hello world",
		},
		Data: "example data",
	}

	assert.Equal(t, req, expect)
	assert.True(t, test.CompareToJSON(expect, string(reqJson)))
}
