package handlers

import (
  "net/http"
  "net/http/httptest"
  "testing"
	"github.com/stretchr/testify/assert"
)

func TestRouterHomeWithSuccess(t *testing.T) {
  r := Router()
  w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/home", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome home", w.Body.String())

}

func TestFunc(t *testing.T) {
  assert.True(t, true, "ข้อความแสดงเวลา error")
  // assert.True(t, false, "ข้อความแสดงเวลา error")
}
