package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfVirtualNode_OutlierDetectionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BaseEjectionDuration() TfVirtualNode_BaseEjectionDurationPropertyOutputReference
	// Experimental.
	BaseEjectionDurationInput() *TfVirtualNode_BaseEjectionDurationProperty
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
	InternalValue() *TfVirtualNode_OutlierDetectionProperty
	// Experimental.
	SetInternalValue(val *TfVirtualNode_OutlierDetectionProperty)
	// Experimental.
	Interval() TfVirtualNode_IntervalPropertyOutputReference
	// Experimental.
	IntervalInput() *TfVirtualNode_IntervalProperty
	// Experimental.
	MaxEjectionPercent() *float64
	// Experimental.
	SetMaxEjectionPercent(val *float64)
	// Experimental.
	MaxEjectionPercentInput() *float64
	// Experimental.
	MaxServerErrors() *float64
	// Experimental.
	SetMaxServerErrors(val *float64)
	// Experimental.
	MaxServerErrorsInput() *float64
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
	PutBaseEjectionDuration(value *TfVirtualNode_BaseEjectionDurationProperty)
	// Experimental.
	PutInterval(value *TfVirtualNode_IntervalProperty)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfVirtualNode_OutlierDetectionPropertyOutputReference
type jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) BaseEjectionDuration() TfVirtualNode_BaseEjectionDurationPropertyOutputReference {
	var returns TfVirtualNode_BaseEjectionDurationPropertyOutputReference
	_jsii_.Get(
		j,
		"baseEjectionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) BaseEjectionDurationInput() *TfVirtualNode_BaseEjectionDurationProperty {
	var returns *TfVirtualNode_BaseEjectionDurationProperty
	_jsii_.Get(
		j,
		"baseEjectionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) InternalValue() *TfVirtualNode_OutlierDetectionProperty {
	var returns *TfVirtualNode_OutlierDetectionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) Interval() TfVirtualNode_IntervalPropertyOutputReference {
	var returns TfVirtualNode_IntervalPropertyOutputReference
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) IntervalInput() *TfVirtualNode_IntervalProperty {
	var returns *TfVirtualNode_IntervalProperty
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) MaxEjectionPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEjectionPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) MaxEjectionPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEjectionPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) MaxServerErrors() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxServerErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) MaxServerErrorsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxServerErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfVirtualNode_OutlierDetectionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfVirtualNode_OutlierDetectionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfVirtualNode_OutlierDetectionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualNode.OutlierDetectionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfVirtualNode_OutlierDetectionPropertyOutputReference_Override(t TfVirtualNode_OutlierDetectionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualNode.OutlierDetectionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference)SetInternalValue(val *TfVirtualNode_OutlierDetectionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference)SetMaxEjectionPercent(val *float64) {
	if err := j.validateSetMaxEjectionPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxEjectionPercent",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference)SetMaxServerErrors(val *float64) {
	if err := j.validateSetMaxServerErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxServerErrors",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) PutBaseEjectionDuration(value *TfVirtualNode_BaseEjectionDurationProperty) {
	if err := t.validatePutBaseEjectionDurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBaseEjectionDuration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) PutInterval(value *TfVirtualNode_IntervalProperty) {
	if err := t.validatePutIntervalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInterval",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfVirtualNode_OutlierDetectionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

