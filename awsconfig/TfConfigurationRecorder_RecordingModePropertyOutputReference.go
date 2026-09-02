package awsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconfig/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConfigurationRecorder_RecordingModePropertyOutputReference interface {
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
	InternalValue() *TfConfigurationRecorder_RecordingModeProperty
	// Experimental.
	SetInternalValue(val *TfConfigurationRecorder_RecordingModeProperty)
	// Experimental.
	RecordingFrequency() *string
	// Experimental.
	SetRecordingFrequency(val *string)
	// Experimental.
	RecordingFrequencyInput() *string
	// Experimental.
	RecordingModeOverride() TfConfigurationRecorder_RecordingModeOverridePropertyOutputReference
	// Experimental.
	RecordingModeOverrideInput() *TfConfigurationRecorder_RecordingModeOverrideProperty
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
	PutRecordingModeOverride(value *TfConfigurationRecorder_RecordingModeOverrideProperty)
	// Experimental.
	ResetRecordingFrequency()
	// Experimental.
	ResetRecordingModeOverride()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfConfigurationRecorder_RecordingModePropertyOutputReference
type jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) InternalValue() *TfConfigurationRecorder_RecordingModeProperty {
	var returns *TfConfigurationRecorder_RecordingModeProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) RecordingFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordingFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) RecordingFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordingFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) RecordingModeOverride() TfConfigurationRecorder_RecordingModeOverridePropertyOutputReference {
	var returns TfConfigurationRecorder_RecordingModeOverridePropertyOutputReference
	_jsii_.Get(
		j,
		"recordingModeOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) RecordingModeOverrideInput() *TfConfigurationRecorder_RecordingModeOverrideProperty {
	var returns *TfConfigurationRecorder_RecordingModeOverrideProperty
	_jsii_.Get(
		j,
		"recordingModeOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConfigurationRecorder_RecordingModePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConfigurationRecorder_RecordingModePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConfigurationRecorder_RecordingModePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-config.TfConfigurationRecorder.RecordingModePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConfigurationRecorder_RecordingModePropertyOutputReference_Override(t TfConfigurationRecorder_RecordingModePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-config.TfConfigurationRecorder.RecordingModePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference)SetInternalValue(val *TfConfigurationRecorder_RecordingModeProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference)SetRecordingFrequency(val *string) {
	if err := j.validateSetRecordingFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordingFrequency",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) PutRecordingModeOverride(value *TfConfigurationRecorder_RecordingModeOverrideProperty) {
	if err := t.validatePutRecordingModeOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRecordingModeOverride",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) ResetRecordingFrequency() {
	_jsii_.InvokeVoid(
		t,
		"resetRecordingFrequency",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) ResetRecordingModeOverride() {
	_jsii_.InvokeVoid(
		t,
		"resetRecordingModeOverride",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingModePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

