package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccessToken() *string
	// Experimental.
	SetAccessToken(val *string)
	// Experimental.
	AccessTokenInput() *string
	// Experimental.
	ClientCredentialsArn() *string
	// Experimental.
	SetClientCredentialsArn(val *string)
	// Experimental.
	ClientCredentialsArnInput() *string
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
	InternalValue() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty
	// Experimental.
	SetInternalValue(val *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty)
	// Experimental.
	JwtToken() *string
	// Experimental.
	SetJwtToken(val *string)
	// Experimental.
	JwtTokenInput() *string
	// Experimental.
	Oauth2GrantType() *string
	// Experimental.
	SetOauth2GrantType(val *string)
	// Experimental.
	Oauth2GrantTypeInput() *string
	// Experimental.
	OauthRequest() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceOauthRequestPropertyOutputReference
	// Experimental.
	OauthRequestInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceOauthRequestProperty
	// Experimental.
	RefreshToken() *string
	// Experimental.
	SetRefreshToken(val *string)
	// Experimental.
	RefreshTokenInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutOauthRequest(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceOauthRequestProperty)
	// Experimental.
	ResetAccessToken()
	// Experimental.
	ResetClientCredentialsArn()
	// Experimental.
	ResetJwtToken()
	// Experimental.
	ResetOauth2GrantType()
	// Experimental.
	ResetOauthRequest()
	// Experimental.
	ResetRefreshToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference
type jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) ClientCredentialsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCredentialsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) ClientCredentialsArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCredentialsArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) InternalValue() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) JwtToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) JwtTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) Oauth2GrantType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2GrantType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) Oauth2GrantTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2GrantTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) OauthRequest() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceOauthRequestPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceOauthRequestPropertyOutputReference
	_jsii_.Get(
		j,
		"oauthRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) OauthRequestInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceOauthRequestProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceOauthRequestProperty
	_jsii_.Get(
		j,
		"oauthRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) RefreshToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) RefreshTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowConnectorProfile.ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference_Override(a AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowConnectorProfile.ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference)SetAccessToken(val *string) {
	if err := j.validateSetAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference)SetClientCredentialsArn(val *string) {
	if err := j.validateSetClientCredentialsArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCredentialsArn",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference)SetInternalValue(val *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference)SetJwtToken(val *string) {
	if err := j.validateSetJwtTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtToken",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference)SetOauth2GrantType(val *string) {
	if err := j.validateSetOauth2GrantTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauth2GrantType",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference)SetRefreshToken(val *string) {
	if err := j.validateSetRefreshTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshToken",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) PutOauthRequest(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceOauthRequestProperty) {
	if err := a.validatePutOauthRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOauthRequest",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) ResetAccessToken() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) ResetClientCredentialsArn() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCredentialsArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) ResetJwtToken() {
	_jsii_.InvokeVoid(
		a,
		"resetJwtToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) ResetOauth2GrantType() {
	_jsii_.InvokeVoid(
		a,
		"resetOauth2GrantType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) ResetOauthRequest() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthRequest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) ResetRefreshToken() {
	_jsii_.InvokeVoid(
		a,
		"resetRefreshToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

