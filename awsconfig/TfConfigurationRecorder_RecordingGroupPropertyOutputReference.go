package awsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconfig/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConfigurationRecorder_RecordingGroupPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllSupported() interface{}
	// Experimental.
	SetAllSupported(val interface{})
	// Experimental.
	AllSupportedInput() interface{}
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
	ExclusionByResourceTypes() TfConfigurationRecorder_ExclusionByResourceTypesPropertyList
	// Experimental.
	ExclusionByResourceTypesInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	IncludeGlobalResourceTypes() interface{}
	// Experimental.
	SetIncludeGlobalResourceTypes(val interface{})
	// Experimental.
	IncludeGlobalResourceTypesInput() interface{}
	// Experimental.
	InternalValue() *TfConfigurationRecorder_RecordingGroupProperty
	// Experimental.
	SetInternalValue(val *TfConfigurationRecorder_RecordingGroupProperty)
	// Experimental.
	RecordingStrategy() TfConfigurationRecorder_RecordingStrategyPropertyList
	// Experimental.
	RecordingStrategyInput() interface{}
	// Experimental.
	ResourceTypes() *[]*string
	// Experimental.
	SetResourceTypes(val *[]*string)
	// Experimental.
	ResourceTypesInput() *[]*string
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
	PutExclusionByResourceTypes(value interface{})
	// Experimental.
	PutRecordingStrategy(value interface{})
	// Experimental.
	ResetAllSupported()
	// Experimental.
	ResetExclusionByResourceTypes()
	// Experimental.
	ResetIncludeGlobalResourceTypes()
	// Experimental.
	ResetRecordingStrategy()
	// Experimental.
	ResetResourceTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfConfigurationRecorder_RecordingGroupPropertyOutputReference
type jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) AllSupported() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) AllSupportedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allSupportedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) ExclusionByResourceTypes() TfConfigurationRecorder_ExclusionByResourceTypesPropertyList {
	var returns TfConfigurationRecorder_ExclusionByResourceTypesPropertyList
	_jsii_.Get(
		j,
		"exclusionByResourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) ExclusionByResourceTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exclusionByResourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) IncludeGlobalResourceTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalResourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) IncludeGlobalResourceTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalResourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) InternalValue() *TfConfigurationRecorder_RecordingGroupProperty {
	var returns *TfConfigurationRecorder_RecordingGroupProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) RecordingStrategy() TfConfigurationRecorder_RecordingStrategyPropertyList {
	var returns TfConfigurationRecorder_RecordingStrategyPropertyList
	_jsii_.Get(
		j,
		"recordingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) RecordingStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) ResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) ResourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConfigurationRecorder_RecordingGroupPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConfigurationRecorder_RecordingGroupPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConfigurationRecorder_RecordingGroupPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-config.TfConfigurationRecorder.RecordingGroupPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConfigurationRecorder_RecordingGroupPropertyOutputReference_Override(t TfConfigurationRecorder_RecordingGroupPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-config.TfConfigurationRecorder.RecordingGroupPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference)SetAllSupported(val interface{}) {
	if err := j.validateSetAllSupportedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allSupported",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference)SetIncludeGlobalResourceTypes(val interface{}) {
	if err := j.validateSetIncludeGlobalResourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeGlobalResourceTypes",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference)SetInternalValue(val *TfConfigurationRecorder_RecordingGroupProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference)SetResourceTypes(val *[]*string) {
	if err := j.validateSetResourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypes",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) PutExclusionByResourceTypes(value interface{}) {
	if err := t.validatePutExclusionByResourceTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExclusionByResourceTypes",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) PutRecordingStrategy(value interface{}) {
	if err := t.validatePutRecordingStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRecordingStrategy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) ResetAllSupported() {
	_jsii_.InvokeVoid(
		t,
		"resetAllSupported",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) ResetExclusionByResourceTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetExclusionByResourceTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) ResetIncludeGlobalResourceTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeGlobalResourceTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) ResetRecordingStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetRecordingStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) ResetResourceTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfConfigurationRecorder_RecordingGroupPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

