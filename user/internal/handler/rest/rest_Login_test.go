package rest_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fenky-ng/edot-test-case/user/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/user/internal/error"
	"github.com/fenky-ng/edot-test-case/user/internal/handler/rest"
	"github.com/fenky-ng/edot-test-case/user/internal/model"
	"github.com/fenky-ng/edot-test-case/user/internal/usecase"
	"github.com/fenky-ng/edot-test-case/user/internal/usecase/auth"
	string_util "github.com/fenky-ng/edot-test-case/user/internal/utility/string"
	test_util "github.com/fenky-ng/edot-test-case/user/internal/utility/test"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_RestAPI_Login(t *testing.T) {
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
		reqBody              any
		expectedHttpCode     int
		expectedResponseBody any
		mock                 func(tt *test)
	}

	var (
		reqMethod = http.MethodPost
		reqUrl    = "/api/v1/user/login"

		name     = "eDOT"
		phone    = "081234567890"
		email    = faker.Email()
		password = "P@5sw0rD"
		jwt      = faker.Jwt()
	)

	prepareReq := func(tt *test) {
		gin.SetMode(gin.TestMode)

		rec := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(rec)

		reqBodyBuf := bytes.NewBuffer([]byte{})
		if tt.reqBody != nil {
			reqBodyBytes, err := json.Marshal(tt.reqBody)
			require.NoError(t, err)
			reqBodyBuf = bytes.NewBuffer(reqBodyBytes)
		}

		c.Request = httptest.NewRequest(reqMethod, reqUrl, reqBodyBuf)
		if tt.reqBody != nil {
			c.Request.Header.Set(constant.ContentType, constant.ApplicationJSON)
		}

		tt.args.c = c
		tt.fields.Recorder = rec
	}

	tests := []test{
		{
			name: "should return response bad request when failed to bind request body json",
			reqBody: struct {
				Password int `json:"password"`
			}{
				Password: 123456,
			},
			expectedHttpCode: http.StatusBadRequest,
			expectedResponseBody: model.RestAPIErrorResponse{
				Error: "Something went wrong. Please try again later.",
			},
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				authUsecaseMock := auth.NewMockAuthUsecaseInterface(mockCtrl)

				tt.fields.Usecase = &usecase.Usecase{
					AuthUsecase: authUsecaseMock,
				}
			},
		},
		{
			name: "should return response bad request when phone or email is empty",
			reqBody: model.RestAPILoginRequest{
				PhoneOrEmail: "",
				Password:     password,
			},
			expectedHttpCode: http.StatusBadRequest,
			expectedResponseBody: model.RestAPIErrorResponse{
				Error: "Please provide either a phone number or an email address to login.",
			},
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				authUsecaseMock := auth.NewMockAuthUsecaseInterface(mockCtrl)

				tt.fields.Usecase = &usecase.Usecase{
					AuthUsecase: authUsecaseMock,
				}
			},
		},
		{
			name: "should return response bad request when no phone or email provided",
			reqBody: model.RestAPILoginRequest{
				PhoneOrEmail: name,
				Password:     password,
			},
			expectedHttpCode: http.StatusBadRequest,
			expectedResponseBody: model.RestAPIErrorResponse{
				Error: "Please provide either a phone number or an email address to login.",
			},
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				authUsecaseMock := auth.NewMockAuthUsecaseInterface(mockCtrl)

				tt.fields.Usecase = &usecase.Usecase{
					AuthUsecase: authUsecaseMock,
				}
			},
		},
		{
			name: "should return response bad request when password is invalid",
			reqBody: model.RestAPILoginRequest{
				PhoneOrEmail: email,
				Password:     "P@5sw",
			},
			expectedHttpCode: http.StatusBadRequest,
			expectedResponseBody: model.RestAPIErrorResponse{
				Error: "The password must be at least 6 characters long. Please enter a valid password.",
			},
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				authUsecaseMock := auth.NewMockAuthUsecaseInterface(mockCtrl)

				tt.fields.Usecase = &usecase.Usecase{
					AuthUsecase: authUsecaseMock,
				}
			},
		},
		{
			name: "should return response internal server error if unmapped error occurred in AuthUsecase.Login",
			reqBody: model.RestAPILoginRequest{
				PhoneOrEmail: email,
				Password:     password,
			},
			expectedHttpCode: http.StatusInternalServerError,
			expectedResponseBody: model.RestAPIErrorResponse{
				Error: "Something went wrong. Please try again later.",
			},
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				authUsecaseMock := auth.NewMockAuthUsecaseInterface(mockCtrl)

				tt.fields.Usecase = &usecase.Usecase{
					AuthUsecase: authUsecaseMock,
				}

				authUsecaseMock.EXPECT().Login(
					test_util.ContextTypeMatcher,
					model.LoginInput{
						Email:    email,
						Password: password,
					},
				).Return(
					model.LoginOutput{},
					errors.Join(in_err.ErrGetUser, errors.New("expected GetUser error")),
				).Times(1)
			},
		},
		{
			name: "should return response id if login with email was successful",
			reqBody: model.RestAPILoginRequest{
				PhoneOrEmail: email,
				Password:     password,
			},
			expectedHttpCode: http.StatusOK,
			expectedResponseBody: model.RestAPILoginResponse{
				JWT: jwt,
			},
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				authUsecaseMock := auth.NewMockAuthUsecaseInterface(mockCtrl)

				tt.fields.Usecase = &usecase.Usecase{
					AuthUsecase: authUsecaseMock,
				}

				authUsecaseMock.EXPECT().Login(
					test_util.ContextTypeMatcher,
					model.LoginInput{
						Email:    email,
						Password: password,
					},
				).Return(
					model.LoginOutput{
						JWT: jwt,
					},
					nil,
				).Times(1)
			},
		},
		{
			name: "should return response id if login with phone was successful",
			reqBody: model.RestAPILoginRequest{
				PhoneOrEmail: phone,
				Password:     password,
			},
			expectedHttpCode: http.StatusOK,
			expectedResponseBody: model.RestAPILoginResponse{
				JWT: jwt,
			},
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				authUsecaseMock := auth.NewMockAuthUsecaseInterface(mockCtrl)

				tt.fields.Usecase = &usecase.Usecase{
					AuthUsecase: authUsecaseMock,
				}

				authUsecaseMock.EXPECT().Login(
					test_util.ContextTypeMatcher,
					model.LoginInput{
						Phone:    phone,
						Password: password,
					},
				).Return(
					model.LoginOutput{
						JWT: jwt,
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
			h.Login(tt.args.c)
			assert.Equal(t, tt.expectedHttpCode, tt.fields.Recorder.Code)
			assert.JSONEq(t, string_util.ParseObjectToJsonString(tt.expectedResponseBody), tt.fields.Recorder.Body.String())
		})
	}
}
