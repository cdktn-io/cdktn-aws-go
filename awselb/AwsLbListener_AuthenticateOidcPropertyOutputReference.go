package awselb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLbListener_AuthenticateOidcPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthenticationRequestExtraParams() *map[string]*string
	// Experimental.
	SetAuthenticationRequestExtraParams(val *map[string]*string)
	// Experimental.
	AuthenticationRequestExtraParamsInput() *map[string]*string
	// Experimental.
	AuthorizationEndpoint() *string
	// Experimental.
	SetAuthorizationEndpoint(val *string)
	// Experimental.
	AuthorizationEndpointInput() *string
	// Experimental.
	ClientId() *string
	// Experimental.
	SetClientId(val *string)
	// Experimental.
	ClientIdInput() *string
	// Experimental.
	ClientSecret() *string
	// Experimental.
	SetClientSecret(val *string)
	// Experimental.
	ClientSecretInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsLbListener_AuthenticateOidcProperty
	// Experimental.
	SetInternalValue(val *AwsLbListener_AuthenticateOidcProperty)
	// Experimental.
	Issuer() *string
	// Experimental.
	SetIssuer(val *string)
	// Experimental.
	IssuerInput() *string
	// Experimental.
	OnUnauthenticatedRequest() *string
	// Experimental.
	SetOnUnauthenticatedRequest(val *string)
	// Experimental.
	OnUnauthenticatedRequestInput() *string
	// Experimental.
	Scope() *string
	// Experimental.
	SetScope(val *string)
	// Experimental.
	ScopeInput() *string
	// Experimental.
	SessionCookieName() *string
	// Experimental.
	SetSessionCookieName(val *string)
	// Experimental.
	SessionCookieNameInput() *string
	// Experimental.
	SessionTimeout() *float64
	// Experimental.
	SetSessionTimeout(val *float64)
	// Experimental.
	SessionTimeoutInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TokenEndpoint() *string
	// Experimental.
	SetTokenEndpoint(val *string)
	// Experimental.
	TokenEndpointInput() *string
	// Experimental.
	UserInfoEndpoint() *string
	// Experimental.
	SetUserInfoEndpoint(val *string)
	// Experimental.
	UserInfoEndpointInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	ResetAuthenticationRequestExtraParams()
	// Experimental.
	ResetOnUnauthenticatedRequest()
	// Experimental.
	ResetScope()
	// Experimental.
	ResetSessionCookieName()
	// Experimental.
	ResetSessionTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLbListener_AuthenticateOidcPropertyOutputReference
type jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) AuthenticationRequestExtraParams() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"authenticationRequestExtraParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) AuthenticationRequestExtraParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"authenticationRequestExtraParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) AuthorizationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) AuthorizationEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) InternalValue() *AwsLbListener_AuthenticateOidcProperty {
	var returns *AwsLbListener_AuthenticateOidcProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) OnUnauthenticatedRequest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onUnauthenticatedRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) OnUnauthenticatedRequestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onUnauthenticatedRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) SessionCookieName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionCookieName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) SessionCookieNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionCookieNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) SessionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) SessionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) TokenEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) TokenEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) UserInfoEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInfoEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) UserInfoEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInfoEndpointInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLbListener_AuthenticateOidcPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsLbListener_AuthenticateOidcPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLbListener_AuthenticateOidcPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elb.AwsLbListener.AuthenticateOidcPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLbListener_AuthenticateOidcPropertyOutputReference_Override(a AwsLbListener_AuthenticateOidcPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.AwsLbListener.AuthenticateOidcPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference)SetAuthenticationRequestExtraParams(val *map[string]*string) {
	if err := j.validateSetAuthenticationRequestExtraParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationRequestExtraParams",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference)SetAuthorizationEndpoint(val *string) {
	if err := j.validateSetAuthorizationEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference)SetInternalValue(val *AwsLbListener_AuthenticateOidcProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference)SetOnUnauthenticatedRequest(val *string) {
	if err := j.validateSetOnUnauthenticatedRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onUnauthenticatedRequest",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference)SetSessionCookieName(val *string) {
	if err := j.validateSetSessionCookieNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionCookieName",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference)SetSessionTimeout(val *float64) {
	if err := j.validateSetSessionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference)SetTokenEndpoint(val *string) {
	if err := j.validateSetTokenEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference)SetUserInfoEndpoint(val *string) {
	if err := j.validateSetUserInfoEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userInfoEndpoint",
		val,
	)
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) ResetAuthenticationRequestExtraParams() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticationRequestExtraParams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) ResetOnUnauthenticatedRequest() {
	_jsii_.InvokeVoid(
		a,
		"resetOnUnauthenticatedRequest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		a,
		"resetScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) ResetSessionCookieName() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionCookieName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) ResetSessionTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLbListener_AuthenticateOidcPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

