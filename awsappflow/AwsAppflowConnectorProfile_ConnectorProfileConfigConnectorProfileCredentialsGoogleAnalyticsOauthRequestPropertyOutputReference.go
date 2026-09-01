package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthCode() *string
	// Experimental.
	SetAuthCode(val *string)
	// Experimental.
	AuthCodeInput() *string
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
	InternalValue() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestProperty
	// Experimental.
	SetInternalValue(val *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestProperty)
	// Experimental.
	RedirectUri() *string
	// Experimental.
	SetRedirectUri(val *string)
	// Experimental.
	RedirectUriInput() *string
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
	ResetAuthCode()
	// Experimental.
	ResetRedirectUri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference
type jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) AuthCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) AuthCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) InternalValue() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) RedirectUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) RedirectUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowConnectorProfile.ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference_Override(a AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowConnectorProfile.ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference)SetAuthCode(val *string) {
	if err := j.validateSetAuthCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authCode",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference)SetInternalValue(val *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference)SetRedirectUri(val *string) {
	if err := j.validateSetRedirectUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectUri",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) ResetAuthCode() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) ResetRedirectUri() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirectUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsOauthRequestPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

