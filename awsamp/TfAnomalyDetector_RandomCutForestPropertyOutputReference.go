package awsamp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsamp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsamp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAnomalyDetector_RandomCutForestPropertyOutputReference interface {
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
	IgnoreNearExpectedFromAbove() TfAnomalyDetector_IgnoreNearExpectedFromAbovePropertyList
	// Experimental.
	IgnoreNearExpectedFromAboveInput() interface{}
	// Experimental.
	IgnoreNearExpectedFromBelow() TfAnomalyDetector_IgnoreNearExpectedFromBelowPropertyList
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

// The jsii proxy struct for TfAnomalyDetector_RandomCutForestPropertyOutputReference
type jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) IgnoreNearExpectedFromAbove() TfAnomalyDetector_IgnoreNearExpectedFromAbovePropertyList {
	var returns TfAnomalyDetector_IgnoreNearExpectedFromAbovePropertyList
	_jsii_.Get(
		j,
		"ignoreNearExpectedFromAbove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) IgnoreNearExpectedFromAboveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNearExpectedFromAboveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) IgnoreNearExpectedFromBelow() TfAnomalyDetector_IgnoreNearExpectedFromBelowPropertyList {
	var returns TfAnomalyDetector_IgnoreNearExpectedFromBelowPropertyList
	_jsii_.Get(
		j,
		"ignoreNearExpectedFromBelow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) IgnoreNearExpectedFromBelowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNearExpectedFromBelowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) SampleSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) SampleSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) ShingleSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shingleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) ShingleSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shingleSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAnomalyDetector_RandomCutForestPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfAnomalyDetector_RandomCutForestPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAnomalyDetector_RandomCutForestPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-amp.TfAnomalyDetector.RandomCutForestPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAnomalyDetector_RandomCutForestPropertyOutputReference_Override(t TfAnomalyDetector_RandomCutForestPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-amp.TfAnomalyDetector.RandomCutForestPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference)SetSampleSize(val *float64) {
	if err := j.validateSetSampleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleSize",
		val,
	)
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference)SetShingleSize(val *float64) {
	if err := j.validateSetShingleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shingleSize",
		val,
	)
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) PutIgnoreNearExpectedFromAbove(value interface{}) {
	if err := t.validatePutIgnoreNearExpectedFromAboveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIgnoreNearExpectedFromAbove",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) PutIgnoreNearExpectedFromBelow(value interface{}) {
	if err := t.validatePutIgnoreNearExpectedFromBelowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIgnoreNearExpectedFromBelow",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) ResetIgnoreNearExpectedFromAbove() {
	_jsii_.InvokeVoid(
		t,
		"resetIgnoreNearExpectedFromAbove",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) ResetIgnoreNearExpectedFromBelow() {
	_jsii_.InvokeVoid(
		t,
		"resetIgnoreNearExpectedFromBelow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) ResetSampleSize() {
	_jsii_.InvokeVoid(
		t,
		"resetSampleSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) ResetShingleSize() {
	_jsii_.InvokeVoid(
		t,
		"resetShingleSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfAnomalyDetector_RandomCutForestPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

