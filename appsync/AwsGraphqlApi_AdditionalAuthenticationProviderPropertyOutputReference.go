package appsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appsync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appsync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthenticationType() *string
	// Experimental.
	SetAuthenticationType(val *string)
	// Experimental.
	AuthenticationTypeInput() *string
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
	LambdaAuthorizerConfig() AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference
	// Experimental.
	LambdaAuthorizerConfigInput() *AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigProperty
	// Experimental.
	OpenidConnectConfig() AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference
	// Experimental.
	OpenidConnectConfigInput() *AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UserPoolConfig() AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference
	// Experimental.
	UserPoolConfigInput() *AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty
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
	PutLambdaAuthorizerConfig(value *AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigProperty)
	// Experimental.
	PutOpenidConnectConfig(value *AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty)
	// Experimental.
	PutUserPoolConfig(value *AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty)
	// Experimental.
	ResetLambdaAuthorizerConfig()
	// Experimental.
	ResetOpenidConnectConfig()
	// Experimental.
	ResetUserPoolConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference
type jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) LambdaAuthorizerConfig() AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference {
	var returns AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"lambdaAuthorizerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) LambdaAuthorizerConfigInput() *AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigProperty {
	var returns *AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigProperty
	_jsii_.Get(
		j,
		"lambdaAuthorizerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) OpenidConnectConfig() AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference {
	var returns AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"openidConnectConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) OpenidConnectConfigInput() *AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty {
	var returns *AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty
	_jsii_.Get(
		j,
		"openidConnectConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) UserPoolConfig() AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference {
	var returns AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"userPoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) UserPoolConfigInput() *AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty {
	var returns *AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty
	_jsii_.Get(
		j,
		"userPoolConfigInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsGraphqlApi.AdditionalAuthenticationProviderPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference_Override(a AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsGraphqlApi.AdditionalAuthenticationProviderPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) PutLambdaAuthorizerConfig(value *AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigProperty) {
	if err := a.validatePutLambdaAuthorizerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaAuthorizerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) PutOpenidConnectConfig(value *AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty) {
	if err := a.validatePutOpenidConnectConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpenidConnectConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) PutUserPoolConfig(value *AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty) {
	if err := a.validatePutUserPoolConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUserPoolConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) ResetLambdaAuthorizerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaAuthorizerConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) ResetOpenidConnectConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetOpenidConnectConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) ResetUserPoolConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetUserPoolConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

