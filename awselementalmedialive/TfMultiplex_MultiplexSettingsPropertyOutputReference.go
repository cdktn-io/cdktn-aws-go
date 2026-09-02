package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMultiplex_MultiplexSettingsPropertyOutputReference interface {
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
	InternalValue() *TfMultiplex_MultiplexSettingsProperty
	// Experimental.
	SetInternalValue(val *TfMultiplex_MultiplexSettingsProperty)
	// Experimental.
	MaximumVideoBufferDelayMilliseconds() *float64
	// Experimental.
	SetMaximumVideoBufferDelayMilliseconds(val *float64)
	// Experimental.
	MaximumVideoBufferDelayMillisecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TransportStreamBitrate() *float64
	// Experimental.
	SetTransportStreamBitrate(val *float64)
	// Experimental.
	TransportStreamBitrateInput() *float64
	// Experimental.
	TransportStreamId() *float64
	// Experimental.
	SetTransportStreamId(val *float64)
	// Experimental.
	TransportStreamIdInput() *float64
	// Experimental.
	TransportStreamReservedBitrate() *float64
	// Experimental.
	SetTransportStreamReservedBitrate(val *float64)
	// Experimental.
	TransportStreamReservedBitrateInput() *float64
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
	ResetMaximumVideoBufferDelayMilliseconds()
	// Experimental.
	ResetTransportStreamReservedBitrate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfMultiplex_MultiplexSettingsPropertyOutputReference
type jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) InternalValue() *TfMultiplex_MultiplexSettingsProperty {
	var returns *TfMultiplex_MultiplexSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) MaximumVideoBufferDelayMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumVideoBufferDelayMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) MaximumVideoBufferDelayMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumVideoBufferDelayMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) TransportStreamBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) TransportStreamBitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamBitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) TransportStreamId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) TransportStreamIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) TransportStreamReservedBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamReservedBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) TransportStreamReservedBitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamReservedBitrateInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfMultiplex_MultiplexSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfMultiplex_MultiplexSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfMultiplex_MultiplexSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfMultiplex.MultiplexSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfMultiplex_MultiplexSettingsPropertyOutputReference_Override(t TfMultiplex_MultiplexSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfMultiplex.MultiplexSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference)SetInternalValue(val *TfMultiplex_MultiplexSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference)SetMaximumVideoBufferDelayMilliseconds(val *float64) {
	if err := j.validateSetMaximumVideoBufferDelayMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumVideoBufferDelayMilliseconds",
		val,
	)
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference)SetTransportStreamBitrate(val *float64) {
	if err := j.validateSetTransportStreamBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportStreamBitrate",
		val,
	)
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference)SetTransportStreamId(val *float64) {
	if err := j.validateSetTransportStreamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportStreamId",
		val,
	)
}

func (j *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference)SetTransportStreamReservedBitrate(val *float64) {
	if err := j.validateSetTransportStreamReservedBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportStreamReservedBitrate",
		val,
	)
}

func (t *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) ResetMaximumVideoBufferDelayMilliseconds() {
	_jsii_.InvokeVoid(
		t,
		"resetMaximumVideoBufferDelayMilliseconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) ResetTransportStreamReservedBitrate() {
	_jsii_.InvokeVoid(
		t,
		"resetTransportStreamReservedBitrate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfMultiplex_MultiplexSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

