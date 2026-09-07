package appsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appsync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appsync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AppIdClientRegex() *string
	// Experimental.
	SetAppIdClientRegex(val *string)
	// Experimental.
	AppIdClientRegexInput() *string
	// Experimental.
	AwsRegion() *string
	// Experimental.
	SetAwsRegion(val *string)
	// Experimental.
	AwsRegionInput() *string
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
	InternalValue() *AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty
	// Experimental.
	SetInternalValue(val *AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UserPoolId() *string
	// Experimental.
	SetUserPoolId(val *string)
	// Experimental.
	UserPoolIdInput() *string
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
	ResetAppIdClientRegex()
	// Experimental.
	ResetAwsRegion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference
type jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) AppIdClientRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdClientRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) AppIdClientRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdClientRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) AwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) AwsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) InternalValue() *AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty {
	var returns *AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) UserPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolIdInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsGraphqlApi.AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference_Override(a AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsGraphqlApi.AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference)SetAppIdClientRegex(val *string) {
	if err := j.validateSetAppIdClientRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appIdClientRegex",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference)SetAwsRegion(val *string) {
	if err := j.validateSetAwsRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsRegion",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference)SetInternalValue(val *AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference)SetUserPoolId(val *string) {
	if err := j.validateSetUserPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) ResetAppIdClientRegex() {
	_jsii_.InvokeVoid(
		a,
		"resetAppIdClientRegex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) ResetAwsRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

