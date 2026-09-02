package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference interface {
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
	InternalValue() *TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesProperty
	// Experimental.
	SetInternalValue(val *TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesProperty)
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

// The jsii proxy struct for TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference
type jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) Forward() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) ForwardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) InternalValue() *TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesProperty {
	var returns *TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) WhitelistedNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelistedNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) WhitelistedNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelistedNamesInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfDistribution.DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference_Override(t TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfDistribution.DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference)SetForward(val *string) {
	if err := j.validateSetForwardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forward",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference)SetInternalValue(val *TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference)SetWhitelistedNames(val *[]*string) {
	if err := j.validateSetWhitelistedNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whitelistedNames",
		val,
	)
}

func (t *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) ResetWhitelistedNames() {
	_jsii_.InvokeVoid(
		t,
		"resetWhitelistedNames",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

