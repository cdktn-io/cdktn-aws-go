package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappsync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappsync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference interface {
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
	LambdaAuthorizerConfig() TfGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference
	// Experimental.
	LambdaAuthorizerConfigInput() *TfGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigProperty
	// Experimental.
	OpenidConnectConfig() TfGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference
	// Experimental.
	OpenidConnectConfigInput() *TfGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UserPoolConfig() TfGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference
	// Experimental.
	UserPoolConfigInput() *TfGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty
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
	PutLambdaAuthorizerConfig(value *TfGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigProperty)
	// Experimental.
	PutOpenidConnectConfig(value *TfGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty)
	// Experimental.
	PutUserPoolConfig(value *TfGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty)
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

// The jsii proxy struct for TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference
type jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) LambdaAuthorizerConfig() TfGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference {
	var returns TfGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"lambdaAuthorizerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) LambdaAuthorizerConfigInput() *TfGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigProperty {
	var returns *TfGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigProperty
	_jsii_.Get(
		j,
		"lambdaAuthorizerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) OpenidConnectConfig() TfGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference {
	var returns TfGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"openidConnectConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) OpenidConnectConfigInput() *TfGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty {
	var returns *TfGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty
	_jsii_.Get(
		j,
		"openidConnectConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) UserPoolConfig() TfGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference {
	var returns TfGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"userPoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) UserPoolConfigInput() *TfGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty {
	var returns *TfGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty
	_jsii_.Get(
		j,
		"userPoolConfigInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appsync.TfGraphqlApi.AdditionalAuthenticationProviderPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference_Override(t TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appsync.TfGraphqlApi.AdditionalAuthenticationProviderPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) PutLambdaAuthorizerConfig(value *TfGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigProperty) {
	if err := t.validatePutLambdaAuthorizerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambdaAuthorizerConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) PutOpenidConnectConfig(value *TfGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty) {
	if err := t.validatePutOpenidConnectConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOpenidConnectConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) PutUserPoolConfig(value *TfGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty) {
	if err := t.validatePutUserPoolConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUserPoolConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) ResetLambdaAuthorizerConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaAuthorizerConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) ResetOpenidConnectConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetOpenidConnectConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) ResetUserPoolConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetUserPoolConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGraphqlApi_AdditionalAuthenticationProviderPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

