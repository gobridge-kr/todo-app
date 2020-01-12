package middleware_test

//go:generate mockgen -destination=../../test/mocks/mock_response_writer.go -package=mocks net/http ResponseWriter
import (
	"testing"

	"github.com/gobridge-kr/todo-app/server/middleware"
	"github.com/gobridge-kr/todo-app/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestCors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	w := mocks.NewMockResponseWriter(ctrl)
	h := map[string][]string{}
	w.EXPECT().Header().Return(h).AnyTimes()

	middleware.Cors(w)

	if m := map[string][]string{
		"Access-Control-Allow-Origin":  []string{"*"},
		"Access-Control-Allow-Methods": []string{"GET, POST, PATCH, DELETE"},
		"Access-Control-Allow-Headers": []string{"accept, content-type"},
		"Content-Type":                 []string{"application/json; charset=UTF-8"},
	}; !cmp.Equal(h, m) {
		t.Errorf("expected: %v\nactual: %v\ndiff: %v", m, h, cmp.Diff(h, m))
	}
}
