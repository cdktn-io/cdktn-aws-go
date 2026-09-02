package awseventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPipe_OverridesPropertyOutputReference interface {
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
	// Experimental.
	ContainerOverride() TfPipe_ContainerOverridePropertyList
	// Experimental.
	ContainerOverrideInput() interface{}
	// Experimental.
	Cpu() *string
	// Experimental.
	SetCpu(val *string)
	// Experimental.
	CpuInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EphemeralStorage() TfPipe_EphemeralStoragePropertyOutputReference
	// Experimental.
	EphemeralStorageInput() *TfPipe_EphemeralStorageProperty
	// Experimental.
	ExecutionRoleArn() *string
	// Experimental.
	SetExecutionRoleArn(val *string)
	// Experimental.
	ExecutionRoleArnInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InferenceAcceleratorOverride() TfPipe_InferenceAcceleratorOverridePropertyList
	// Experimental.
	InferenceAcceleratorOverrideInput() interface{}
	// Experimental.
	InternalValue() *TfPipe_OverridesProperty
	// Experimental.
	SetInternalValue(val *TfPipe_OverridesProperty)
	// Experimental.
	Memory() *string
	// Experimental.
	SetMemory(val *string)
	// Experimental.
	MemoryInput() *string
	// Experimental.
	TaskRoleArn() *string
	// Experimental.
	SetTaskRoleArn(val *string)
	// Experimental.
	TaskRoleArnInput() *string
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
	PutContainerOverride(value interface{})
	// Experimental.
	PutEphemeralStorage(value *TfPipe_EphemeralStorageProperty)
	// Experimental.
	PutInferenceAcceleratorOverride(value interface{})
	// Experimental.
	ResetContainerOverride()
	// Experimental.
	ResetCpu()
	// Experimental.
	ResetEphemeralStorage()
	// Experimental.
	ResetExecutionRoleArn()
	// Experimental.
	ResetInferenceAcceleratorOverride()
	// Experimental.
	ResetMemory()
	// Experimental.
	ResetTaskRoleArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPipe_OverridesPropertyOutputReference
type jsiiProxy_TfPipe_OverridesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) ContainerOverride() TfPipe_ContainerOverridePropertyList {
	var returns TfPipe_ContainerOverridePropertyList
	_jsii_.Get(
		j,
		"containerOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) ContainerOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) Cpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) CpuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) EphemeralStorage() TfPipe_EphemeralStoragePropertyOutputReference {
	var returns TfPipe_EphemeralStoragePropertyOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) EphemeralStorageInput() *TfPipe_EphemeralStorageProperty {
	var returns *TfPipe_EphemeralStorageProperty
	_jsii_.Get(
		j,
		"ephemeralStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) InferenceAcceleratorOverride() TfPipe_InferenceAcceleratorOverridePropertyList {
	var returns TfPipe_InferenceAcceleratorOverridePropertyList
	_jsii_.Get(
		j,
		"inferenceAcceleratorOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) InferenceAcceleratorOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceAcceleratorOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) InternalValue() *TfPipe_OverridesProperty {
	var returns *TfPipe_OverridesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) Memory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) MemoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) TaskRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) TaskRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPipe_OverridesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPipe_OverridesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPipe_OverridesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPipe_OverridesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.OverridesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPipe_OverridesPropertyOutputReference_Override(t TfPipe_OverridesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.OverridesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference)SetCpu(val *string) {
	if err := j.validateSetCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference)SetInternalValue(val *TfPipe_OverridesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference)SetMemory(val *string) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference)SetTaskRoleArn(val *string) {
	if err := j.validateSetTaskRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPipe_OverridesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) PutContainerOverride(value interface{}) {
	if err := t.validatePutContainerOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putContainerOverride",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) PutEphemeralStorage(value *TfPipe_EphemeralStorageProperty) {
	if err := t.validatePutEphemeralStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEphemeralStorage",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) PutInferenceAcceleratorOverride(value interface{}) {
	if err := t.validatePutInferenceAcceleratorOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInferenceAcceleratorOverride",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) ResetContainerOverride() {
	_jsii_.InvokeVoid(
		t,
		"resetContainerOverride",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) ResetCpu() {
	_jsii_.InvokeVoid(
		t,
		"resetCpu",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) ResetEphemeralStorage() {
	_jsii_.InvokeVoid(
		t,
		"resetEphemeralStorage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) ResetExecutionRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetExecutionRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) ResetInferenceAcceleratorOverride() {
	_jsii_.InvokeVoid(
		t,
		"resetInferenceAcceleratorOverride",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) ResetMemory() {
	_jsii_.InvokeVoid(
		t,
		"resetMemory",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) ResetTaskRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetTaskRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPipe_OverridesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

