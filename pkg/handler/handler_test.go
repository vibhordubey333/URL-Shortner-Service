package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"os"
	"testing"
)

var mockContext *gin.Context

func init() {
	os.Setenv("GIN_MODE", "release")
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

/*
func TestHandleShortUrlRedirect(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandleShortUrlRedirect(tt.args.c)
		})
	}
}*/

// ########
type MockUsers struct {
	ctrl     *gomock.Controller
	recorder *MockUsersMockRecorders
}

type MockUsersMockRecorders struct {
	mock *MockUsers
}

func NewMockUsers(ctrl *gomock.Controller) *MockUsers {
	mock := &MockUsers{ctrl: ctrl}
	mock.recorder = &MockUsersMockRecorders{mock: mock}
	return mock
}

func TestURLShortnerProcessor_CreateShortUrl(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//mockFetcher := mock_handler.NewMockURLShortnerHandlerInterface(mockCtrl)
	mockContext, err := gin.CreateTestContext(httptest.NewRecorder())
	if err != nil {
		//log.Fatal("Error while initializing mockContext: ", err)
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "CreateShortURL_Success",
			args: args{mockContext},
		},
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := new(URLShortnerProcessor)
			o.CreateShortUrl(tt.args.c)
		})

	}
}
