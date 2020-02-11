package ping

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"ginexample/common"
	"ginexample/utils/log"
)

func TestMain(m *testing.M) {
	err := common.LoadTestEnv()
	if err != nil {
		log.Fatal(err)
	}
	common.Setup()
	if err != nil {
		log.Fatal(err)
	}
	exitCode := m.Run()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(exitCode)
}
func TestPing(t *testing.T) {

	router := pingRouter()
	type reqParam map[string]string
	tests := []struct {
		name      string
		reqParams reqParam
		respCode  int
		respBody  []byte
	}{
		{name: "case 1", respCode: http.StatusOK, respBody: []byte(`{"message":"5.7.28"}`)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", PingUri, nil)
			q := req.URL.Query()
			for key, value := range tt.reqParams {
				q.Add(key, value)
			}
			req.URL.RawQuery = q.Encode()
			router.ServeHTTP(w, req)
			if w.Code != tt.respCode {
				t.Errorf("%s got %v, want %v", PingUri, w.Code, tt.respCode)
			} else {
				body := w.Body.Bytes()
				body = bytes.Trim(body, "\n")
				if !bytes.Equal(body, tt.respBody) {
					t.Errorf("%s got %s, want %s", PingUri, body, tt.respBody)
				}
			}
		})
	}
}
