package bedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthorizationServerMetadata() AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryAuthorizationServerMetadataPropertyList
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
	DiscoveryUrl() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryProperty
	// Experimental.
	SetInternalValue(val *AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryProperty)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference
type jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) AuthorizationServerMetadata() AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryAuthorizationServerMetadataPropertyList {
	var returns AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryAuthorizationServerMetadataPropertyList
	_jsii_.Get(
		j,
		"authorizationServerMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) DiscoveryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) InternalValue() *AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryProperty {
	var returns *AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsOauth2CredentialProvider.Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference_Override(a AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsOauth2CredentialProvider.Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference)SetInternalValue(val *AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsOauth2CredentialProvider_Oauth2ProviderConfigMicrosoftOauth2ProviderConfigOauthDiscoveryPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

