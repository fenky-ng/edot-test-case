package rest_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fenky-ng/edot-test-case/user/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/user/internal/error"
	"github.com/fenky-ng/edot-test-case/user/internal/handler/rest"
	"github.com/fenky-ng/edot-test-case/user/internal/model"
	"github.com/fenky-ng/edot-test-case/user/internal/usecase"
	"github.com/fenky-ng/edot-test-case/user/internal/usecase/profile"
	jwt_util "github.com/fenky-ng/edot-test-case/user/internal/utility/jwt"
	string_util "github.com/fenky-ng/edot-test-case/user/internal/utility/string"
	test_util "github.com/fenky-ng/edot-test-case/user/internal/utility/test"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_RestAPI_GetProfile(t *testing.T) {
	type fields struct {
		Usecase  *usecase.Usecase
		Recorder *httptest.ResponseRecorder
	}
	type args struct {
		c *gin.Context
	}
	type test struct {
		name                 string
		fields               fields
		args                 args
		headerAuth           string
		expectedHttpCode     int
		expectedResponseBody any
		mock                 func(tt *test)
	}

	var (
		reqMethod = http.MethodGet
		reqUrl    = "/api/v1/user/profile"

		id    = uuid.New()
		name  = faker.Name()
		phone = faker.Phonenumber()
		email = faker.Email()

		jwt, _ = jwt_util.GenerateJWT(id)
	)

	prepareReq := func(tt *test) {
		gin.SetMode(gin.TestMode)

		rec := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(rec)

		reqBodyBuf := bytes.NewBuffer([]byte{})

		c.Request = httptest.NewRequest(reqMethod, reqUrl, reqBodyBuf)
		c.Request.Header.Set(constant.HeaderAuth, tt.headerAuth)

		tt.args.c = c
		tt.fields.Recorder = rec
	}

	tests := []test{
		{
			name:             "should return error if header auth was empty",
			headerAuth:       "",
			expectedHttpCode: http.StatusUnauthorized,
			expectedResponseBody: model.RestAPIErrorResponse{
				Error: "Authentication token is required.",
			},
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				profileUsecaseMock := profile.NewMockProfileUsecaseInterface(mockCtrl)

				tt.fields.Usecase = &usecase.Usecase{
					ProfileUsecase: profileUsecaseMock,
				}
			},
		},
		{
			name:             "should return error if no token provided",
			headerAuth:       constant.AuthBearer + " ",
			expectedHttpCode: http.StatusUnauthorized,
			expectedResponseBody: model.RestAPIErrorResponse{
				Error: "Authentication token is required.",
			},
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				profileUsecaseMock := profile.NewMockProfileUsecaseInterface(mockCtrl)

				tt.fields.Usecase = &usecase.Usecase{
					ProfileUsecase: profileUsecaseMock,
				}
			},
		},
		{
			name:             "should return response internal server error if unmapped error occured in ProfileUsecase.GetProfile",
			headerAuth:       constant.AuthBearer + " " + jwt,
			expectedHttpCode: http.StatusInternalServerError,
			expectedResponseBody: model.RestAPIErrorResponse{
				Error: "Something went wrong. Please try again later.",
			},
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				profileUsecaseMock := profile.NewMockProfileUsecaseInterface(mockCtrl)

				tt.fields.Usecase = &usecase.Usecase{
					ProfileUsecase: profileUsecaseMock,
				}

				profileUsecaseMock.EXPECT().GetProfile(
					test_util.ContextTypeMatcher,
					model.GetProfileInput{
						Id: id,
					},
				).Return(
					model.GetProfileOutput{},
					errors.Join(in_err.ErrGetUser, errors.New("expected GetUser error")),
				).Times(1)
			},
		},
		{
			name:             "should return data if get profile was successful",
			headerAuth:       constant.AuthBearer + " " + jwt,
			expectedHttpCode: http.StatusOK,
			expectedResponseBody: model.RestAPIGetProfileResponse{
				Id:    id,
				Name:  name,
				Phone: phone,
				Email: email,
			},
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				profileUsecaseMock := profile.NewMockProfileUsecaseInterface(mockCtrl)

				tt.fields.Usecase = &usecase.Usecase{
					ProfileUsecase: profileUsecaseMock,
				}

				profileUsecaseMock.EXPECT().GetProfile(
					test_util.ContextTypeMatcher,
					model.GetProfileInput{
						Id: id,
					},
				).Return(
					model.GetProfileOutput{
						Id:    id,
						Name:  name,
						Phone: phone,
						Email: email,
					},
					nil,
				).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prepareReq(&tt)

			if tt.mock != nil {
				tt.mock(&tt)
			}

			h := rest.InitRestAPI(rest.InitRestAPIOptions{
				Usecase: tt.fields.Usecase,
			})
			h.GetProfile(tt.args.c)
			assert.Equal(t, tt.expectedHttpCode, tt.fields.Recorder.Code)
			assert.JSONEq(t, string_util.ParseObjectToJsonString(tt.expectedResponseBody), tt.fields.Recorder.Body.String())
		})
	}
}
