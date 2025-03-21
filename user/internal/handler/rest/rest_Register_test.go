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
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_RestAPI_Register(t *testing.T) {
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
		reqUrl    = "/api/v1/user/register"

		name     = "eDOT"
		phone    = "081234567890"
		email    = faker.Email()
		password = "P@5sw0rD"
		id       = uuid.New()
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
			name: "should return response bad request when name is invalid",
			reqBody: model.RestAPIRegisterRequest{
				Name:         "UT",
				PhoneOrEmail: email,
				Password:     password,
			},
			expectedHttpCode: http.StatusBadRequest,
			expectedResponseBody: model.RestAPIErrorResponse{
				Error: "The name must be at least 3 characters long. Please enter a valid name.",
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
			reqBody: model.RestAPIRegisterRequest{
				Name:         name,
				PhoneOrEmail: "",
				Password:     password,
			},
			expectedHttpCode: http.StatusBadRequest,
			expectedResponseBody: model.RestAPIErrorResponse{
				Error: "Please provide either a phone number or an email address to register.",
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
			reqBody: model.RestAPIRegisterRequest{
				Name:         name,
				PhoneOrEmail: name,
				Password:     password,
			},
			expectedHttpCode: http.StatusBadRequest,
			expectedResponseBody: model.RestAPIErrorResponse{
				Error: "Please provide either a phone number or an email address to register.",
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
			reqBody: model.RestAPIRegisterRequest{
				Name:         name,
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
			name: "should return response internal server error if unmapped error occured in AuthUsecase.Register",
			reqBody: model.RestAPIRegisterRequest{
				Name:         name,
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

				authUsecaseMock.EXPECT().Register(
					test_util.ContextTypeMatcher,
					model.RegisterInput{
						Name:     name,
						Email:    email,
						Password: password,
					},
				).Return(
					model.RegisterOutput{},
					errors.Join(in_err.ErrInsertUser, errors.New("expected InsertUser error")),
				).Times(1)
			},
		},
		{
			name: "should return response id if register with email was successful",
			reqBody: model.RestAPIRegisterRequest{
				Name:         name,
				PhoneOrEmail: email,
				Password:     password,
			},
			expectedHttpCode: http.StatusOK,
			expectedResponseBody: model.RestAPIRegisterResponse{
				Id: id,
			},
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				authUsecaseMock := auth.NewMockAuthUsecaseInterface(mockCtrl)

				tt.fields.Usecase = &usecase.Usecase{
					AuthUsecase: authUsecaseMock,
				}

				authUsecaseMock.EXPECT().Register(
					test_util.ContextTypeMatcher,
					model.RegisterInput{
						Name:     name,
						Email:    email,
						Password: password,
					},
				).Return(
					model.RegisterOutput{
						Id: id,
					},
					nil,
				).Times(1)
			},
		},
		{
			name: "should return response id if register with phone was successful",
			reqBody: model.RestAPIRegisterRequest{
				Name:         name,
				PhoneOrEmail: phone,
				Password:     password,
			},
			expectedHttpCode: http.StatusOK,
			expectedResponseBody: model.RestAPIRegisterResponse{
				Id: id,
			},
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				authUsecaseMock := auth.NewMockAuthUsecaseInterface(mockCtrl)

				tt.fields.Usecase = &usecase.Usecase{
					AuthUsecase: authUsecaseMock,
				}

				authUsecaseMock.EXPECT().Register(
					test_util.ContextTypeMatcher,
					model.RegisterInput{
						Name:     name,
						Phone:    phone,
						Password: password,
					},
				).Return(
					model.RegisterOutput{
						Id: id,
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
			h.Register(tt.args.c)
			assert.Equal(t, tt.expectedHttpCode, tt.fields.Recorder.Code)
			assert.JSONEq(t, string_util.ParseObjectToJsonString(tt.expectedResponseBody), tt.fields.Recorder.Body.String())
		})
	}
}
