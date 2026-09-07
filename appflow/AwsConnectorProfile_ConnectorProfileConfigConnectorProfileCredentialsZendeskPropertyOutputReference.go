package appflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference interface {
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
	InternalValue() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty
	// Experimental.
	SetInternalValue(val *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty)
	// Experimental.
	OauthRequest() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskOauthRequestPropertyOutputReference
	// Experimental.
	OauthRequestInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskOauthRequestProperty
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
	PutOauthRequest(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskOauthRequestProperty)
	// Experimental.
	ResetAccessToken()
	// Experimental.
	ResetOauthRequest()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference
type jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) InternalValue() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) OauthRequest() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskOauthRequestPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskOauthRequestPropertyOutputReference
	_jsii_.Get(
		j,
		"oauthRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) OauthRequestInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskOauthRequestProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskOauthRequestProperty
	_jsii_.Get(
		j,
		"oauthRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsConnectorProfile.ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference_Override(a AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsConnectorProfile.ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference)SetAccessToken(val *string) {
	if err := j.validateSetAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference)SetInternalValue(val *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) PutOauthRequest(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskOauthRequestProperty) {
	if err := a.validatePutOauthRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOauthRequest",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) ResetAccessToken() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) ResetOauthRequest() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthRequest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

