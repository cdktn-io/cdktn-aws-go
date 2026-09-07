package amp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/amp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/amp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAnomalyDetector_RandomCutForestPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	IgnoreNearExpectedFromAbove() AwsAnomalyDetector_IgnoreNearExpectedFromAbovePropertyList
	// Experimental.
	IgnoreNearExpectedFromAboveInput() interface{}
	// Experimental.
	IgnoreNearExpectedFromBelow() AwsAnomalyDetector_IgnoreNearExpectedFromBelowPropertyList
	// Experimental.
	IgnoreNearExpectedFromBelowInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Query() *string
	// Experimental.
	SetQuery(val *string)
	// Experimental.
	QueryInput() *string
	// Experimental.
	SampleSize() *float64
	// Experimental.
	SetSampleSize(val *float64)
	// Experimental.
	SampleSizeInput() *float64
	// Experimental.
	ShingleSize() *float64
	// Experimental.
	SetShingleSize(val *float64)
	// Experimental.
	ShingleSizeInput() *float64
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
	PutIgnoreNearExpectedFromAbove(value interface{})
	// Experimental.
	PutIgnoreNearExpectedFromBelow(value interface{})
	// Experimental.
	ResetIgnoreNearExpectedFromAbove()
	// Experimental.
	ResetIgnoreNearExpectedFromBelow()
	// Experimental.
	ResetSampleSize()
	// Experimental.
	ResetShingleSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsAnomalyDetector_RandomCutForestPropertyOutputReference
type jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) IgnoreNearExpectedFromAbove() AwsAnomalyDetector_IgnoreNearExpectedFromAbovePropertyList {
	var returns AwsAnomalyDetector_IgnoreNearExpectedFromAbovePropertyList
	_jsii_.Get(
		j,
		"ignoreNearExpectedFromAbove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) IgnoreNearExpectedFromAboveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNearExpectedFromAboveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) IgnoreNearExpectedFromBelow() AwsAnomalyDetector_IgnoreNearExpectedFromBelowPropertyList {
	var returns AwsAnomalyDetector_IgnoreNearExpectedFromBelowPropertyList
	_jsii_.Get(
		j,
		"ignoreNearExpectedFromBelow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) IgnoreNearExpectedFromBelowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNearExpectedFromBelowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) SampleSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) SampleSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) ShingleSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shingleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) ShingleSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shingleSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAnomalyDetector_RandomCutForestPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsAnomalyDetector_RandomCutForestPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAnomalyDetector_RandomCutForestPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-amp.AwsAnomalyDetector.RandomCutForestPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAnomalyDetector_RandomCutForestPropertyOutputReference_Override(a AwsAnomalyDetector_RandomCutForestPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-amp.AwsAnomalyDetector.RandomCutForestPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference)SetSampleSize(val *float64) {
	if err := j.validateSetSampleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleSize",
		val,
	)
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference)SetShingleSize(val *float64) {
	if err := j.validateSetShingleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shingleSize",
		val,
	)
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) PutIgnoreNearExpectedFromAbove(value interface{}) {
	if err := a.validatePutIgnoreNearExpectedFromAboveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIgnoreNearExpectedFromAbove",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) PutIgnoreNearExpectedFromBelow(value interface{}) {
	if err := a.validatePutIgnoreNearExpectedFromBelowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIgnoreNearExpectedFromBelow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) ResetIgnoreNearExpectedFromAbove() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoreNearExpectedFromAbove",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) ResetIgnoreNearExpectedFromBelow() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoreNearExpectedFromBelow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) ResetSampleSize() {
	_jsii_.InvokeVoid(
		a,
		"resetSampleSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) ResetShingleSize() {
	_jsii_.InvokeVoid(
		a,
		"resetShingleSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAnomalyDetector_RandomCutForestPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

