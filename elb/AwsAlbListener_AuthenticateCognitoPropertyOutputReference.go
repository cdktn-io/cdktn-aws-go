package elb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAlbListener_AuthenticateCognitoPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthenticationRequestExtraParams() *map[string]*string
	// Experimental.
	SetAuthenticationRequestExtraParams(val *map[string]*string)
	// Experimental.
	AuthenticationRequestExtraParamsInput() *map[string]*string
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
	InternalValue() *AwsAlbListener_AuthenticateCognitoProperty
	// Experimental.
	SetInternalValue(val *AwsAlbListener_AuthenticateCognitoProperty)
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
	UserPoolArn() *string
	// Experimental.
	SetUserPoolArn(val *string)
	// Experimental.
	UserPoolArnInput() *string
	// Experimental.
	UserPoolClientId() *string
	// Experimental.
	SetUserPoolClientId(val *string)
	// Experimental.
	UserPoolClientIdInput() *string
	// Experimental.
	UserPoolDomain() *string
	// Experimental.
	SetUserPoolDomain(val *string)
	// Experimental.
	UserPoolDomainInput() *string
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

// The jsii proxy struct for AwsAlbListener_AuthenticateCognitoPropertyOutputReference
type jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) AuthenticationRequestExtraParams() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"authenticationRequestExtraParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) AuthenticationRequestExtraParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"authenticationRequestExtraParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) InternalValue() *AwsAlbListener_AuthenticateCognitoProperty {
	var returns *AwsAlbListener_AuthenticateCognitoProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) OnUnauthenticatedRequest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onUnauthenticatedRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) OnUnauthenticatedRequestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onUnauthenticatedRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) SessionCookieName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionCookieName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) SessionCookieNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionCookieNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) SessionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) SessionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) UserPoolArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) UserPoolArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) UserPoolClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) UserPoolClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) UserPoolDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) UserPoolDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolDomainInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAlbListener_AuthenticateCognitoPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAlbListener_AuthenticateCognitoPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAlbListener_AuthenticateCognitoPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elb.AwsAlbListener.AuthenticateCognitoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAlbListener_AuthenticateCognitoPropertyOutputReference_Override(a AwsAlbListener_AuthenticateCognitoPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.AwsAlbListener.AuthenticateCognitoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference)SetAuthenticationRequestExtraParams(val *map[string]*string) {
	if err := j.validateSetAuthenticationRequestExtraParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationRequestExtraParams",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference)SetInternalValue(val *AwsAlbListener_AuthenticateCognitoProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference)SetOnUnauthenticatedRequest(val *string) {
	if err := j.validateSetOnUnauthenticatedRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onUnauthenticatedRequest",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference)SetSessionCookieName(val *string) {
	if err := j.validateSetSessionCookieNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionCookieName",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference)SetSessionTimeout(val *float64) {
	if err := j.validateSetSessionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference)SetUserPoolArn(val *string) {
	if err := j.validateSetUserPoolArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPoolArn",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference)SetUserPoolClientId(val *string) {
	if err := j.validateSetUserPoolClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPoolClientId",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference)SetUserPoolDomain(val *string) {
	if err := j.validateSetUserPoolDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPoolDomain",
		val,
	)
}

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) ResetAuthenticationRequestExtraParams() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticationRequestExtraParams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) ResetOnUnauthenticatedRequest() {
	_jsii_.InvokeVoid(
		a,
		"resetOnUnauthenticatedRequest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		a,
		"resetScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) ResetSessionCookieName() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionCookieName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) ResetSessionTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAlbListener_AuthenticateCognitoPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

