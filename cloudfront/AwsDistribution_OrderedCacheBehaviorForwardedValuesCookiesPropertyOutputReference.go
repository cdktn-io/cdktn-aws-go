package cloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference interface {
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
	Forward() *string
	// Experimental.
	SetForward(val *string)
	// Experimental.
	ForwardInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty
	// Experimental.
	SetInternalValue(val *AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WhitelistedNames() *[]*string
	// Experimental.
	SetWhitelistedNames(val *[]*string)
	// Experimental.
	WhitelistedNamesInput() *[]*string
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
	ResetWhitelistedNames()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference
type jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) Forward() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) ForwardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) InternalValue() *AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty {
	var returns *AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) WhitelistedNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelistedNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) WhitelistedNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelistedNamesInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsDistribution.OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference_Override(a AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsDistribution.OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference)SetForward(val *string) {
	if err := j.validateSetForwardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forward",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference)SetInternalValue(val *AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference)SetWhitelistedNames(val *[]*string) {
	if err := j.validateSetWhitelistedNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whitelistedNames",
		val,
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) ResetWhitelistedNames() {
	_jsii_.InvokeVoid(
		a,
		"resetWhitelistedNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

