package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Capabilities() TfDaemonTaskDefinition_CapabilitiesPropertyList
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
	Device() TfDaemonTaskDefinition_DevicePropertyList
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
	Tmpfs() TfDaemonTaskDefinition_TmpfsPropertyList
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

// The jsii proxy struct for TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference
type jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) Capabilities() TfDaemonTaskDefinition_CapabilitiesPropertyList {
	var returns TfDaemonTaskDefinition_CapabilitiesPropertyList
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) CapabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) Device() TfDaemonTaskDefinition_DevicePropertyList {
	var returns TfDaemonTaskDefinition_DevicePropertyList
	_jsii_.Get(
		j,
		"device",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) DeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) InitProcessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initProcessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) InitProcessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initProcessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) Tmpfs() TfDaemonTaskDefinition_TmpfsPropertyList {
	var returns TfDaemonTaskDefinition_TmpfsPropertyList
	_jsii_.Get(
		j,
		"tmpfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) TmpfsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tmpfsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDaemonTaskDefinition_LinuxParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDaemonTaskDefinition_LinuxParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.TfDaemonTaskDefinition.LinuxParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDaemonTaskDefinition_LinuxParametersPropertyOutputReference_Override(t TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.TfDaemonTaskDefinition.LinuxParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference)SetInitProcessEnabled(val interface{}) {
	if err := j.validateSetInitProcessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initProcessEnabled",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) PutCapabilities(value interface{}) {
	if err := t.validatePutCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCapabilities",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) PutDevice(value interface{}) {
	if err := t.validatePutDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDevice",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) PutTmpfs(value interface{}) {
	if err := t.validatePutTmpfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTmpfs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) ResetCapabilities() {
	_jsii_.InvokeVoid(
		t,
		"resetCapabilities",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) ResetDevice() {
	_jsii_.InvokeVoid(
		t,
		"resetDevice",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) ResetInitProcessEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetInitProcessEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) ResetTmpfs() {
	_jsii_.InvokeVoid(
		t,
		"resetTmpfs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDaemonTaskDefinition_LinuxParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

