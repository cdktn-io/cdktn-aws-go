package eventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/eventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/eventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPipe_OverridesPropertyOutputReference interface {
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
	ContainerOverride() AwsPipe_ContainerOverridePropertyList
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
	EphemeralStorage() AwsPipe_EphemeralStoragePropertyOutputReference
	// Experimental.
	EphemeralStorageInput() *AwsPipe_EphemeralStorageProperty
	// Experimental.
	ExecutionRoleArn() *string
	// Experimental.
	SetExecutionRoleArn(val *string)
	// Experimental.
	ExecutionRoleArnInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InferenceAcceleratorOverride() AwsPipe_InferenceAcceleratorOverridePropertyList
	// Experimental.
	InferenceAcceleratorOverrideInput() interface{}
	// Experimental.
	InternalValue() *AwsPipe_OverridesProperty
	// Experimental.
	SetInternalValue(val *AwsPipe_OverridesProperty)
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
	PutEphemeralStorage(value *AwsPipe_EphemeralStorageProperty)
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

// The jsii proxy struct for AwsPipe_OverridesPropertyOutputReference
type jsiiProxy_AwsPipe_OverridesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) ContainerOverride() AwsPipe_ContainerOverridePropertyList {
	var returns AwsPipe_ContainerOverridePropertyList
	_jsii_.Get(
		j,
		"containerOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) ContainerOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) Cpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) CpuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) EphemeralStorage() AwsPipe_EphemeralStoragePropertyOutputReference {
	var returns AwsPipe_EphemeralStoragePropertyOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) EphemeralStorageInput() *AwsPipe_EphemeralStorageProperty {
	var returns *AwsPipe_EphemeralStorageProperty
	_jsii_.Get(
		j,
		"ephemeralStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) InferenceAcceleratorOverride() AwsPipe_InferenceAcceleratorOverridePropertyList {
	var returns AwsPipe_InferenceAcceleratorOverridePropertyList
	_jsii_.Get(
		j,
		"inferenceAcceleratorOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) InferenceAcceleratorOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceAcceleratorOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) InternalValue() *AwsPipe_OverridesProperty {
	var returns *AwsPipe_OverridesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) Memory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) MemoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) TaskRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) TaskRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPipe_OverridesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsPipe_OverridesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPipe_OverridesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPipe_OverridesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipe.OverridesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPipe_OverridesPropertyOutputReference_Override(a AwsPipe_OverridesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipe.OverridesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference)SetCpu(val *string) {
	if err := j.validateSetCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference)SetInternalValue(val *AwsPipe_OverridesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference)SetMemory(val *string) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference)SetTaskRoleArn(val *string) {
	if err := j.validateSetTaskRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_OverridesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) PutContainerOverride(value interface{}) {
	if err := a.validatePutContainerOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putContainerOverride",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) PutEphemeralStorage(value *AwsPipe_EphemeralStorageProperty) {
	if err := a.validatePutEphemeralStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEphemeralStorage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) PutInferenceAcceleratorOverride(value interface{}) {
	if err := a.validatePutInferenceAcceleratorOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInferenceAcceleratorOverride",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) ResetContainerOverride() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerOverride",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) ResetCpu() {
	_jsii_.InvokeVoid(
		a,
		"resetCpu",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) ResetEphemeralStorage() {
	_jsii_.InvokeVoid(
		a,
		"resetEphemeralStorage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) ResetExecutionRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) ResetInferenceAcceleratorOverride() {
	_jsii_.InvokeVoid(
		a,
		"resetInferenceAcceleratorOverride",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) ResetMemory() {
	_jsii_.InvokeVoid(
		a,
		"resetMemory",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) ResetTaskRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetTaskRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPipe_OverridesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

