package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ClientCredentialsWoVersion() *float64
	// Experimental.
	SetClientCredentialsWoVersion(val *float64)
	// Experimental.
	ClientCredentialsWoVersionInput() *float64
	// Experimental.
	ClientId() *string
	// Experimental.
	SetClientId(val *string)
	// Experimental.
	ClientIdInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	ClientIdWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetClientIdWo(val *string)
	// Experimental.
	ClientIdWoInput() *string
	// Experimental.
	ClientSecret() *string
	// Experimental.
	SetClientSecret(val *string)
	// Experimental.
	ClientSecretInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	ClientSecretWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetClientSecretWo(val *string)
	// Experimental.
	ClientSecretWoInput() *string
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	OauthDiscovery() AwsBedrockagentcoreOauth2CredentialProvider_Oauth2ProviderConfigGithubOauth2ProviderConfigOauthDiscoveryPropertyList
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
	ResetClientCredentialsWoVersion()
	// Experimental.
	ResetClientId()
	// Experimental.
	ResetClientIdWo()
	// Experimental.
	ResetClientSecret()
	// Experimental.
	ResetClientSecretWo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference
type jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ClientCredentialsWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientCredentialsWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ClientCredentialsWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientCredentialsWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ClientIdWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ClientIdWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ClientSecretWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ClientSecretWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) OauthDiscovery() AwsBedrockagentcoreOauth2CredentialProvider_Oauth2ProviderConfigGithubOauth2ProviderConfigOauthDiscoveryPropertyList {
	var returns AwsBedrockagentcoreOauth2CredentialProvider_Oauth2ProviderConfigGithubOauth2ProviderConfigOauthDiscoveryPropertyList
	_jsii_.Get(
		j,
		"oauthDiscovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreOauth2CredentialProvider.GithubOauth2ProviderConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference_Override(a AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreOauth2CredentialProvider.GithubOauth2ProviderConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference)SetClientCredentialsWoVersion(val *float64) {
	if err := j.validateSetClientCredentialsWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCredentialsWoVersion",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference)SetClientIdWo(val *string) {
	if err := j.validateSetClientIdWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientIdWo",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference)SetClientSecretWo(val *string) {
	if err := j.validateSetClientSecretWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecretWo",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ResetClientCredentialsWoVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCredentialsWoVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ResetClientIdWo() {
	_jsii_.InvokeVoid(
		a,
		"resetClientIdWo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ResetClientSecretWo() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecretWo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBedrockagentcoreOauth2CredentialProvider_GithubOauth2ProviderConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

