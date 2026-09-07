package appsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appsync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appsync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthorizerResultTtlInSeconds() *float64
	// Experimental.
	SetAuthorizerResultTtlInSeconds(val *float64)
	// Experimental.
	AuthorizerResultTtlInSecondsInput() *float64
	// Experimental.
	AuthorizerUri() *string
	// Experimental.
	SetAuthorizerUri(val *string)
	// Experimental.
	AuthorizerUriInput() *string
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
	IdentityValidationExpression() *string
	// Experimental.
	SetIdentityValidationExpression(val *string)
	// Experimental.
	IdentityValidationExpressionInput() *string
	// Experimental.
	InternalValue() *AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigProperty
	// Experimental.
	SetInternalValue(val *AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigProperty)
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
	ResetAuthorizerResultTtlInSeconds()
	// Experimental.
	ResetIdentityValidationExpression()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference
type jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) AuthorizerResultTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorizerResultTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) AuthorizerResultTtlInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorizerResultTtlInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) AuthorizerUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) AuthorizerUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) IdentityValidationExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityValidationExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) IdentityValidationExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityValidationExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) InternalValue() *AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigProperty {
	var returns *AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsGraphqlApi.AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference_Override(a AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsGraphqlApi.AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference)SetAuthorizerResultTtlInSeconds(val *float64) {
	if err := j.validateSetAuthorizerResultTtlInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizerResultTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference)SetAuthorizerUri(val *string) {
	if err := j.validateSetAuthorizerUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizerUri",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference)SetIdentityValidationExpression(val *string) {
	if err := j.validateSetIdentityValidationExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityValidationExpression",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference)SetInternalValue(val *AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) ResetAuthorizerResultTtlInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerResultTtlInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) ResetIdentityValidationExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityValidationExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

