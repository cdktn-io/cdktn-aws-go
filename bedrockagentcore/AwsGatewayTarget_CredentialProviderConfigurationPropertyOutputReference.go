package bedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ApiKey() AwsGatewayTarget_ApiKeyPropertyList
	// Experimental.
	ApiKeyInput() interface{}
	// Experimental.
	CallerIamCredentials() AwsGatewayTarget_CallerIamCredentialsPropertyList
	// Experimental.
	CallerIamCredentialsInput() interface{}
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
	GatewayIamRole() AwsGatewayTarget_GatewayIamRolePropertyList
	// Experimental.
	GatewayIamRoleInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	JwtPassthrough() AwsGatewayTarget_JwtPassthroughPropertyList
	// Experimental.
	JwtPassthroughInput() interface{}
	// Experimental.
	Oauth() AwsGatewayTarget_OauthPropertyList
	// Experimental.
	OauthInput() interface{}
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
	PutApiKey(value interface{})
	// Experimental.
	PutCallerIamCredentials(value interface{})
	// Experimental.
	PutGatewayIamRole(value interface{})
	// Experimental.
	PutJwtPassthrough(value interface{})
	// Experimental.
	PutOauth(value interface{})
	// Experimental.
	ResetApiKey()
	// Experimental.
	ResetCallerIamCredentials()
	// Experimental.
	ResetGatewayIamRole()
	// Experimental.
	ResetJwtPassthrough()
	// Experimental.
	ResetOauth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference
type jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ApiKey() AwsGatewayTarget_ApiKeyPropertyList {
	var returns AwsGatewayTarget_ApiKeyPropertyList
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ApiKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) CallerIamCredentials() AwsGatewayTarget_CallerIamCredentialsPropertyList {
	var returns AwsGatewayTarget_CallerIamCredentialsPropertyList
	_jsii_.Get(
		j,
		"callerIamCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) CallerIamCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"callerIamCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GatewayIamRole() AwsGatewayTarget_GatewayIamRolePropertyList {
	var returns AwsGatewayTarget_GatewayIamRolePropertyList
	_jsii_.Get(
		j,
		"gatewayIamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GatewayIamRoleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatewayIamRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) JwtPassthrough() AwsGatewayTarget_JwtPassthroughPropertyList {
	var returns AwsGatewayTarget_JwtPassthroughPropertyList
	_jsii_.Get(
		j,
		"jwtPassthrough",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) JwtPassthroughInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jwtPassthroughInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) Oauth() AwsGatewayTarget_OauthPropertyList {
	var returns AwsGatewayTarget_OauthPropertyList
	_jsii_.Get(
		j,
		"oauth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) OauthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oauthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsGatewayTarget.CredentialProviderConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference_Override(a AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsGatewayTarget.CredentialProviderConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) PutApiKey(value interface{}) {
	if err := a.validatePutApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApiKey",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) PutCallerIamCredentials(value interface{}) {
	if err := a.validatePutCallerIamCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCallerIamCredentials",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) PutGatewayIamRole(value interface{}) {
	if err := a.validatePutGatewayIamRoleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGatewayIamRole",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) PutJwtPassthrough(value interface{}) {
	if err := a.validatePutJwtPassthroughParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJwtPassthrough",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) PutOauth(value interface{}) {
	if err := a.validatePutOauthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOauth",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ResetApiKey() {
	_jsii_.InvokeVoid(
		a,
		"resetApiKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ResetCallerIamCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetCallerIamCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ResetGatewayIamRole() {
	_jsii_.InvokeVoid(
		a,
		"resetGatewayIamRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ResetJwtPassthrough() {
	_jsii_.InvokeVoid(
		a,
		"resetJwtPassthrough",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ResetOauth() {
	_jsii_.InvokeVoid(
		a,
		"resetOauth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

