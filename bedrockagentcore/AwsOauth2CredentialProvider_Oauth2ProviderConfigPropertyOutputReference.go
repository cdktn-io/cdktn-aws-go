package bedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference interface {
	cdktn.ComplexObject
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
	CustomOauth2ProviderConfig() AwsOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyList
	// Experimental.
	CustomOauth2ProviderConfigInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	GithubOauth2ProviderConfig() AwsOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyList
	// Experimental.
	GithubOauth2ProviderConfigInput() interface{}
	// Experimental.
	GoogleOauth2ProviderConfig() AwsOauth2CredentialProvider_GoogleOauth2ProviderConfigPropertyList
	// Experimental.
	GoogleOauth2ProviderConfigInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MicrosoftOauth2ProviderConfig() AwsOauth2CredentialProvider_MicrosoftOauth2ProviderConfigPropertyList
	// Experimental.
	MicrosoftOauth2ProviderConfigInput() interface{}
	// Experimental.
	SalesforceOauth2ProviderConfig() AwsOauth2CredentialProvider_SalesforceOauth2ProviderConfigPropertyList
	// Experimental.
	SalesforceOauth2ProviderConfigInput() interface{}
	// Experimental.
	SlackOauth2ProviderConfig() AwsOauth2CredentialProvider_SlackOauth2ProviderConfigPropertyList
	// Experimental.
	SlackOauth2ProviderConfigInput() interface{}
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
	PutCustomOauth2ProviderConfig(value interface{})
	// Experimental.
	PutGithubOauth2ProviderConfig(value interface{})
	// Experimental.
	PutGoogleOauth2ProviderConfig(value interface{})
	// Experimental.
	PutMicrosoftOauth2ProviderConfig(value interface{})
	// Experimental.
	PutSalesforceOauth2ProviderConfig(value interface{})
	// Experimental.
	PutSlackOauth2ProviderConfig(value interface{})
	// Experimental.
	ResetCustomOauth2ProviderConfig()
	// Experimental.
	ResetGithubOauth2ProviderConfig()
	// Experimental.
	ResetGoogleOauth2ProviderConfig()
	// Experimental.
	ResetMicrosoftOauth2ProviderConfig()
	// Experimental.
	ResetSalesforceOauth2ProviderConfig()
	// Experimental.
	ResetSlackOauth2ProviderConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference
type jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) CustomOauth2ProviderConfig() AwsOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyList {
	var returns AwsOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyList
	_jsii_.Get(
		j,
		"customOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) CustomOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) GithubOauth2ProviderConfig() AwsOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyList {
	var returns AwsOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyList
	_jsii_.Get(
		j,
		"githubOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) GithubOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) GoogleOauth2ProviderConfig() AwsOauth2CredentialProvider_GoogleOauth2ProviderConfigPropertyList {
	var returns AwsOauth2CredentialProvider_GoogleOauth2ProviderConfigPropertyList
	_jsii_.Get(
		j,
		"googleOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) GoogleOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"googleOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) MicrosoftOauth2ProviderConfig() AwsOauth2CredentialProvider_MicrosoftOauth2ProviderConfigPropertyList {
	var returns AwsOauth2CredentialProvider_MicrosoftOauth2ProviderConfigPropertyList
	_jsii_.Get(
		j,
		"microsoftOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) MicrosoftOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"microsoftOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) SalesforceOauth2ProviderConfig() AwsOauth2CredentialProvider_SalesforceOauth2ProviderConfigPropertyList {
	var returns AwsOauth2CredentialProvider_SalesforceOauth2ProviderConfigPropertyList
	_jsii_.Get(
		j,
		"salesforceOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) SalesforceOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"salesforceOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) SlackOauth2ProviderConfig() AwsOauth2CredentialProvider_SlackOauth2ProviderConfigPropertyList {
	var returns AwsOauth2CredentialProvider_SlackOauth2ProviderConfigPropertyList
	_jsii_.Get(
		j,
		"slackOauth2ProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) SlackOauth2ProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slackOauth2ProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsOauth2CredentialProvider.Oauth2ProviderConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference_Override(a AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsOauth2CredentialProvider.Oauth2ProviderConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) PutCustomOauth2ProviderConfig(value interface{}) {
	if err := a.validatePutCustomOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) PutGithubOauth2ProviderConfig(value interface{}) {
	if err := a.validatePutGithubOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGithubOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) PutGoogleOauth2ProviderConfig(value interface{}) {
	if err := a.validatePutGoogleOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) PutMicrosoftOauth2ProviderConfig(value interface{}) {
	if err := a.validatePutMicrosoftOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMicrosoftOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) PutSalesforceOauth2ProviderConfig(value interface{}) {
	if err := a.validatePutSalesforceOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSalesforceOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) PutSlackOauth2ProviderConfig(value interface{}) {
	if err := a.validatePutSlackOauth2ProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSlackOauth2ProviderConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) ResetCustomOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) ResetGithubOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetGithubOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) ResetGoogleOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) ResetMicrosoftOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetMicrosoftOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) ResetSalesforceOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSalesforceOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) ResetSlackOauth2ProviderConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSlackOauth2ProviderConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

