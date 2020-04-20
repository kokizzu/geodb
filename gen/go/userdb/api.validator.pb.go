// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Account) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	if this.Payment != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Payment); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Payment", err)
		}
	}
	return nil
}
func (this *Payment) Validate() error {
	for _, item := range this.Subscriptions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Subscriptions", err)
			}
		}
	}
	return nil
}
func (this *Subscription) Validate() error {
	return nil
}

var _regex_UserDetail_Email = regexp.MustCompile(`^.{1,225}$`)
var _regex_UserDetail_Name = regexp.MustCompile(`^.{1,225}$`)

func (this *UserDetail) Validate() error {
	if !_regex_UserDetail_Email.MatchString(this.Email) {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Email))
	}
	if !_regex_UserDetail_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	// Validation of proto3 map<> fields is unsupported.
	if this.Account != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Account); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Account", err)
		}
	}
	return nil
}

var _regex_LoginRequest_Code = regexp.MustCompile(`^.{1,225}$`)

func (this *LoginRequest) Validate() error {
	if !_regex_LoginRequest_Code.MatchString(this.Code) {
		return github_com_mwitkow_go_proto_validators.FieldError("Code", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Code))
	}
	return nil
}
func (this *LoginResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}

var _regex_LoginJWTRequest_Jwt = regexp.MustCompile(`^.{1,225}$`)

func (this *LoginJWTRequest) Validate() error {
	if !_regex_LoginJWTRequest_Jwt.MatchString(this.Jwt) {
		return github_com_mwitkow_go_proto_validators.FieldError("Jwt", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Jwt))
	}
	return nil
}
func (this *LoginJWTResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *SetRequest) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *SetResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *GetEmailsRequest) Validate() error {
	return nil
}
func (this *GetEmailsResponse) Validate() error {
	return nil
}

var _regex_GetRegexEmailsRequest_Regex = regexp.MustCompile(`^.{1,225}$`)

func (this *GetRegexEmailsRequest) Validate() error {
	if !_regex_GetRegexEmailsRequest_Regex.MatchString(this.Regex) {
		return github_com_mwitkow_go_proto_validators.FieldError("Regex", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Regex))
	}
	return nil
}
func (this *GetRegexEmailsResponse) Validate() error {
	return nil
}
func (this *GetRequest) Validate() error {
	return nil
}
func (this *GetResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}

var _regex_GetRegexRequest_Regex = regexp.MustCompile(`^.{1,225}$`)

func (this *GetRegexRequest) Validate() error {
	if !_regex_GetRegexRequest_Regex.MatchString(this.Regex) {
		return github_com_mwitkow_go_proto_validators.FieldError("Regex", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Regex))
	}
	return nil
}
func (this *GetRegexResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *DeleteRequest) Validate() error {
	return nil
}
func (this *DeleteResponse) Validate() error {
	return nil
}
func (this *PingRequest) Validate() error {
	return nil
}
func (this *PingResponse) Validate() error {
	return nil
}

var _regex_NewAccountRequest_Name = regexp.MustCompile(`^.{1,225}$`)
var _regex_NewAccountRequest_AdminEmail = regexp.MustCompile(`^.{1,225}$`)

func (this *NewAccountRequest) Validate() error {
	if !_regex_NewAccountRequest_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	if !_regex_NewAccountRequest_AdminEmail.MatchString(this.AdminEmail) {
		return github_com_mwitkow_go_proto_validators.FieldError("AdminEmail", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.AdminEmail))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *NewAccountResponse) Validate() error {
	if this.Account != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Account); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Account", err)
		}
	}
	return nil
}
func (this *GetAccountRequest) Validate() error {
	return nil
}
func (this *GetAccountResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}

var _regex_GetAccountRegexRequest_Regex = regexp.MustCompile(`^.{1,225}$`)

func (this *GetAccountRegexRequest) Validate() error {
	if !_regex_GetAccountRegexRequest_Regex.MatchString(this.Regex) {
		return github_com_mwitkow_go_proto_validators.FieldError("Regex", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Regex))
	}
	return nil
}
func (this *GetAccountRegexResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}

var _regex_SetAccountPlanRequest_AccountName = regexp.MustCompile(`^.{1,225}$`)
var _regex_SetAccountPlanRequest_Plan = regexp.MustCompile(`^.{1,225}$`)

func (this *SetAccountPlanRequest) Validate() error {
	if !_regex_SetAccountPlanRequest_AccountName.MatchString(this.AccountName) {
		return github_com_mwitkow_go_proto_validators.FieldError("AccountName", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.AccountName))
	}
	if !_regex_SetAccountPlanRequest_Plan.MatchString(this.Plan) {
		return github_com_mwitkow_go_proto_validators.FieldError("Plan", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Plan))
	}
	return nil
}
func (this *SetAccountPlanResponse) Validate() error {
	if this.Account != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Account); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Account", err)
		}
	}
	return nil
}

var _regex_SetAccountSourceRequest_AccountName = regexp.MustCompile(`^.{1,225}$`)
var _regex_SetAccountSourceRequest_Source = regexp.MustCompile(`^.{1,225}$`)

func (this *SetAccountSourceRequest) Validate() error {
	if !_regex_SetAccountSourceRequest_AccountName.MatchString(this.AccountName) {
		return github_com_mwitkow_go_proto_validators.FieldError("AccountName", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.AccountName))
	}
	if !_regex_SetAccountSourceRequest_Source.MatchString(this.Source) {
		return github_com_mwitkow_go_proto_validators.FieldError("Source", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Source))
	}
	return nil
}
func (this *SetAccountSourceResponse) Validate() error {
	if this.Account != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Account); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Account", err)
		}
	}
	return nil
}
func (this *DeleteAccountRequest) Validate() error {
	return nil
}
func (this *DeleteAccountResponse) Validate() error {
	return nil
}
func (this *GetAccountNamesRequest) Validate() error {
	return nil
}
func (this *GetAccountNamesResponse) Validate() error {
	return nil
}

var _regex_GetAccountNamesRegexRequest_Regex = regexp.MustCompile(`^.{1,225}$`)

func (this *GetAccountNamesRegexRequest) Validate() error {
	if !_regex_GetAccountNamesRegexRequest_Regex.MatchString(this.Regex) {
		return github_com_mwitkow_go_proto_validators.FieldError("Regex", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Regex))
	}
	return nil
}
func (this *GetAccountNamesRegexResponse) Validate() error {
	return nil
}

var _regex_SetUserAccountRequest_UserEmail = regexp.MustCompile(`^.{1,225}$`)
var _regex_SetUserAccountRequest_AccountName = regexp.MustCompile(`^.{1,225}$`)

func (this *SetUserAccountRequest) Validate() error {
	if !_regex_SetUserAccountRequest_UserEmail.MatchString(this.UserEmail) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserEmail", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.UserEmail))
	}
	if !_regex_SetUserAccountRequest_AccountName.MatchString(this.AccountName) {
		return github_com_mwitkow_go_proto_validators.FieldError("AccountName", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.AccountName))
	}
	return nil
}
func (this *SetUserAccountResponse) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}

var _regex_IncAccountPlanUsageRequest_AccountName = regexp.MustCompile(`^.{1,225}$`)
var _regex_IncAccountPlanUsageRequest_Plan = regexp.MustCompile(`^.{1,225}$`)

func (this *IncAccountPlanUsageRequest) Validate() error {
	if !(this.Increment > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Increment", fmt.Errorf(`value '%v' must be greater than '0'`, this.Increment))
	}
	if !_regex_IncAccountPlanUsageRequest_AccountName.MatchString(this.AccountName) {
		return github_com_mwitkow_go_proto_validators.FieldError("AccountName", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.AccountName))
	}
	if !_regex_IncAccountPlanUsageRequest_Plan.MatchString(this.Plan) {
		return github_com_mwitkow_go_proto_validators.FieldError("Plan", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Plan))
	}
	return nil
}
func (this *IncAccountPlanUsageResponse) Validate() error {
	return nil
}
