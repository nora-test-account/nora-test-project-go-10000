// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package noratestproject10000

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/nora-test-project-10000-go/internal/apijson"
	"github.com/stainless-sdks/nora-test-project-10000-go/internal/apiquery"
	"github.com/stainless-sdks/nora-test-project-10000-go/internal/requestconfig"
	"github.com/stainless-sdks/nora-test-project-10000-go/option"
	"github.com/stainless-sdks/nora-test-project-10000-go/packages/param"
	"github.com/stainless-sdks/nora-test-project-10000-go/packages/respjson"
)

// UserService contains methods and other services that help with interacting with
// the nora-test-project-10000 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	return
}

// This can only be done by the logged in user.
func (r *UserService) New(ctx context.Context, body UserNewParams, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get user by user name
func (r *UserService) Get(ctx context.Context, username string, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	if username == "" {
		err = errors.New("missing required username parameter")
		return
	}
	path := fmt.Sprintf("user/%s", username)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This can only be done by the logged in user.
func (r *UserService) Update(ctx context.Context, existingUsername string, body UserUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if existingUsername == "" {
		err = errors.New("missing required existingUsername parameter")
		return
	}
	path := fmt.Sprintf("user/%s", existingUsername)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// This can only be done by the logged in user.
func (r *UserService) Delete(ctx context.Context, username string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if username == "" {
		err = errors.New("missing required username parameter")
		return
	}
	path := fmt.Sprintf("user/%s", username)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Creates list of users with given input array
func (r *UserService) NewWithList(ctx context.Context, body UserNewWithListParams, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/createWithList"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Logs user into the system
func (r *UserService) Login(ctx context.Context, query UserLoginParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Logs out current logged in user session
func (r *UserService) Logout(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "user/logout"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Username  string `json:"username"`
	// User Status
	UserStatus int64 `json:"userStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		Password    respjson.Field
		Phone       respjson.Field
		Username    respjson.Field
		UserStatus  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this User to a UserParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UserParam.Overrides()
func (r User) ToParam() UserParam {
	return param.Override[UserParam](r.RawJSON())
}

type UserParam struct {
	ID        param.Opt[int64]  `json:"id,omitzero"`
	Email     param.Opt[string] `json:"email,omitzero"`
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	LastName  param.Opt[string] `json:"lastName,omitzero"`
	Password  param.Opt[string] `json:"password,omitzero"`
	Phone     param.Opt[string] `json:"phone,omitzero"`
	Username  param.Opt[string] `json:"username,omitzero"`
	// User Status
	UserStatus param.Opt[int64] `json:"userStatus,omitzero"`
	paramObj
}

func (r UserParam) MarshalJSON() (data []byte, err error) {
	type shadow UserParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewParams struct {
	User UserParam
	paramObj
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.User)
}
func (r *UserNewParams) UnmarshalJSON(data []byte) error {
	return r.User.UnmarshalJSON(data)
}

type UserUpdateParams struct {
	User UserParam
	paramObj
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.User)
}
func (r *UserUpdateParams) UnmarshalJSON(data []byte) error {
	return r.User.UnmarshalJSON(data)
}

type UserNewWithListParams struct {
	Items []UserParam
	paramObj
}

func (r UserNewWithListParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Items)
}
func (r *UserNewWithListParams) UnmarshalJSON(data []byte) error {
	return r.Items.UnmarshalJSON(data)
}

type UserLoginParams struct {
	// The password for login in clear text
	Password param.Opt[string] `query:"password,omitzero" json:"-"`
	// The user name for login
	Username param.Opt[string] `query:"username,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserLoginParams]'s query parameters as `url.Values`.
func (r UserLoginParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
