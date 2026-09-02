package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConnectorProfile_Oauth2PropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccessToken() *string
	// Experimental.
	SetAccessToken(val *string)
	// Experimental.
	AccessTokenInput() *string
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
	InternalValue() *TfConnectorProfile_Oauth2Property
	// Experimental.
	SetInternalValue(val *TfConnectorProfile_Oauth2Property)
	// Experimental.
	OauthRequest() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OauthRequestPropertyOutputReference
	// Experimental.
	OauthRequestInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OauthRequestProperty
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
	PutOauthRequest(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OauthRequestProperty)
	// Experimental.
	ResetAccessToken()
	// Experimental.
	ResetClientId()
	// Experimental.
	ResetClientSecret()
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

// The jsii proxy struct for TfConnectorProfile_Oauth2PropertyOutputReference
type jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) InternalValue() *TfConnectorProfile_Oauth2Property {
	var returns *TfConnectorProfile_Oauth2Property
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) OauthRequest() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OauthRequestPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OauthRequestPropertyOutputReference
	_jsii_.Get(
		j,
		"oauthRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) OauthRequestInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OauthRequestProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OauthRequestProperty
	_jsii_.Get(
		j,
		"oauthRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) RefreshToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) RefreshTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConnectorProfile_Oauth2PropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConnectorProfile_Oauth2PropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConnectorProfile_Oauth2PropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.TfConnectorProfile.Oauth2PropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConnectorProfile_Oauth2PropertyOutputReference_Override(t TfConnectorProfile_Oauth2PropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.TfConnectorProfile.Oauth2PropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference)SetAccessToken(val *string) {
	if err := j.validateSetAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference)SetInternalValue(val *TfConnectorProfile_Oauth2Property) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference)SetRefreshToken(val *string) {
	if err := j.validateSetRefreshTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshToken",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) PutOauthRequest(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OauthRequestProperty) {
	if err := t.validatePutOauthRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOauthRequest",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) ResetAccessToken() {
	_jsii_.InvokeVoid(
		t,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		t,
		"resetClientId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		t,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) ResetOauthRequest() {
	_jsii_.InvokeVoid(
		t,
		"resetOauthRequest",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) ResetRefreshToken() {
	_jsii_.InvokeVoid(
		t,
		"resetRefreshToken",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_Oauth2PropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

