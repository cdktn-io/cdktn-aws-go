package ecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/ecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Capabilities() AwsDaemonTaskDefinition_CapabilitiesPropertyList
	// Experimental.
	CapabilitiesInput() interface{}
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
	Device() AwsDaemonTaskDefinition_DevicePropertyList
	// Experimental.
	DeviceInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InitProcessEnabled() interface{}
	// Experimental.
	SetInitProcessEnabled(val interface{})
	// Experimental.
	InitProcessEnabledInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Tmpfs() AwsDaemonTaskDefinition_TmpfsPropertyList
	// Experimental.
	TmpfsInput() interface{}
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
	PutCapabilities(value interface{})
	// Experimental.
	PutDevice(value interface{})
	// Experimental.
	PutTmpfs(value interface{})
	// Experimental.
	ResetCapabilities()
	// Experimental.
	ResetDevice()
	// Experimental.
	ResetInitProcessEnabled()
	// Experimental.
	ResetTmpfs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference
type jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) Capabilities() AwsDaemonTaskDefinition_CapabilitiesPropertyList {
	var returns AwsDaemonTaskDefinition_CapabilitiesPropertyList
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) CapabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) Device() AwsDaemonTaskDefinition_DevicePropertyList {
	var returns AwsDaemonTaskDefinition_DevicePropertyList
	_jsii_.Get(
		j,
		"device",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) DeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) InitProcessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initProcessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) InitProcessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initProcessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) Tmpfs() AwsDaemonTaskDefinition_TmpfsPropertyList {
	var returns AwsDaemonTaskDefinition_TmpfsPropertyList
	_jsii_.Get(
		j,
		"tmpfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) TmpfsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tmpfsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDaemonTaskDefinition_LinuxParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsDaemonTaskDefinition.LinuxParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference_Override(a AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsDaemonTaskDefinition.LinuxParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference)SetInitProcessEnabled(val interface{}) {
	if err := j.validateSetInitProcessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initProcessEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) PutCapabilities(value interface{}) {
	if err := a.validatePutCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCapabilities",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) PutDevice(value interface{}) {
	if err := a.validatePutDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDevice",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) PutTmpfs(value interface{}) {
	if err := a.validatePutTmpfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTmpfs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) ResetCapabilities() {
	_jsii_.InvokeVoid(
		a,
		"resetCapabilities",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) ResetDevice() {
	_jsii_.InvokeVoid(
		a,
		"resetDevice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) ResetInitProcessEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetInitProcessEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) ResetTmpfs() {
	_jsii_.InvokeVoid(
		a,
		"resetTmpfs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDaemonTaskDefinition_LinuxParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

