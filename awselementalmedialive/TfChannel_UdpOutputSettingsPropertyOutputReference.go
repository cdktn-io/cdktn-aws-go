package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_UdpOutputSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BufferMsec() *float64
	// Experimental.
	SetBufferMsec(val *float64)
	// Experimental.
	BufferMsecInput() *float64
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
	ContainerSettings() TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference
	// Experimental.
	ContainerSettingsInput() *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Destination() TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsDestinationPropertyOutputReference
	// Experimental.
	DestinationInput() *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsDestinationProperty
	// Experimental.
	FecOutputSettings() TfChannel_FecOutputSettingsPropertyOutputReference
	// Experimental.
	FecOutputSettingsInput() *TfChannel_FecOutputSettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfChannel_UdpOutputSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_UdpOutputSettingsProperty)
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
	PutContainerSettings(value *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsProperty)
	// Experimental.
	PutDestination(value *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsDestinationProperty)
	// Experimental.
	PutFecOutputSettings(value *TfChannel_FecOutputSettingsProperty)
	// Experimental.
	ResetBufferMsec()
	// Experimental.
	ResetFecOutputSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_UdpOutputSettingsPropertyOutputReference
type jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) BufferMsec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferMsec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) BufferMsecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferMsecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) ContainerSettings() TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference {
	var returns TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"containerSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) ContainerSettingsInput() *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsProperty {
	var returns *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsProperty
	_jsii_.Get(
		j,
		"containerSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) Destination() TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsDestinationPropertyOutputReference {
	var returns TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) DestinationInput() *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsDestinationProperty {
	var returns *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsDestinationProperty
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) FecOutputSettings() TfChannel_FecOutputSettingsPropertyOutputReference {
	var returns TfChannel_FecOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"fecOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) FecOutputSettingsInput() *TfChannel_FecOutputSettingsProperty {
	var returns *TfChannel_FecOutputSettingsProperty
	_jsii_.Get(
		j,
		"fecOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) InternalValue() *TfChannel_UdpOutputSettingsProperty {
	var returns *TfChannel_UdpOutputSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_UdpOutputSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_UdpOutputSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_UdpOutputSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.UdpOutputSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_UdpOutputSettingsPropertyOutputReference_Override(t TfChannel_UdpOutputSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.UdpOutputSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference)SetBufferMsec(val *float64) {
	if err := j.validateSetBufferMsecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferMsec",
		val,
	)
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_UdpOutputSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) PutContainerSettings(value *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsProperty) {
	if err := t.validatePutContainerSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putContainerSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) PutDestination(value *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsDestinationProperty) {
	if err := t.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDestination",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) PutFecOutputSettings(value *TfChannel_FecOutputSettingsProperty) {
	if err := t.validatePutFecOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFecOutputSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) ResetBufferMsec() {
	_jsii_.InvokeVoid(
		t,
		"resetBufferMsec",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) ResetFecOutputSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetFecOutputSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_UdpOutputSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

