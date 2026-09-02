package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference interface {
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
	OauthDiscovery() TfOauth2CredentialProvider_Oauth2ProviderConfigCustomOauth2ProviderConfigOauthDiscoveryPropertyList
	// Experimental.
	OauthDiscoveryInput() interface{}
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
	PutOauthDiscovery(value interface{})
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
	// Experimental.
	ResetOauthDiscovery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference
type jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ClientCredentialsWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientCredentialsWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ClientCredentialsWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientCredentialsWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ClientIdWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ClientIdWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ClientSecretWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ClientSecretWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) OauthDiscovery() TfOauth2CredentialProvider_Oauth2ProviderConfigCustomOauth2ProviderConfigOauthDiscoveryPropertyList {
	var returns TfOauth2CredentialProvider_Oauth2ProviderConfigCustomOauth2ProviderConfigOauthDiscoveryPropertyList
	_jsii_.Get(
		j,
		"oauthDiscovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) OauthDiscoveryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oauthDiscoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfOauth2CredentialProvider.CustomOauth2ProviderConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference_Override(t TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfOauth2CredentialProvider.CustomOauth2ProviderConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference)SetClientCredentialsWoVersion(val *float64) {
	if err := j.validateSetClientCredentialsWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCredentialsWoVersion",
		val,
	)
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference)SetClientIdWo(val *string) {
	if err := j.validateSetClientIdWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientIdWo",
		val,
	)
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference)SetClientSecretWo(val *string) {
	if err := j.validateSetClientSecretWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecretWo",
		val,
	)
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) PutOauthDiscovery(value interface{}) {
	if err := t.validatePutOauthDiscoveryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOauthDiscovery",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ResetClientCredentialsWoVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetClientCredentialsWoVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		t,
		"resetClientId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ResetClientIdWo() {
	_jsii_.InvokeVoid(
		t,
		"resetClientIdWo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		t,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ResetClientSecretWo() {
	_jsii_.InvokeVoid(
		t,
		"resetClientSecretWo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ResetOauthDiscovery() {
	_jsii_.InvokeVoid(
		t,
		"resetOauthDiscovery",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfOauth2CredentialProvider_CustomOauth2ProviderConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

