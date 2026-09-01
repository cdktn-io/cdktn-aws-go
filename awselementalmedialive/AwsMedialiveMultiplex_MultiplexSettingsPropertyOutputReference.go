package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference interface {
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
	InternalValue() *AwsMedialiveMultiplex_MultiplexSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsMedialiveMultiplex_MultiplexSettingsProperty)
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

// The jsii proxy struct for AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference
type jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) InternalValue() *AwsMedialiveMultiplex_MultiplexSettingsProperty {
	var returns *AwsMedialiveMultiplex_MultiplexSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) MaximumVideoBufferDelayMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumVideoBufferDelayMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) MaximumVideoBufferDelayMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumVideoBufferDelayMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) TransportStreamBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) TransportStreamBitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamBitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) TransportStreamId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) TransportStreamIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) TransportStreamReservedBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamReservedBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) TransportStreamReservedBitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamReservedBitrateInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveMultiplex.MultiplexSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference_Override(a AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveMultiplex.MultiplexSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference)SetInternalValue(val *AwsMedialiveMultiplex_MultiplexSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference)SetMaximumVideoBufferDelayMilliseconds(val *float64) {
	if err := j.validateSetMaximumVideoBufferDelayMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumVideoBufferDelayMilliseconds",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference)SetTransportStreamBitrate(val *float64) {
	if err := j.validateSetTransportStreamBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportStreamBitrate",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference)SetTransportStreamId(val *float64) {
	if err := j.validateSetTransportStreamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportStreamId",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference)SetTransportStreamReservedBitrate(val *float64) {
	if err := j.validateSetTransportStreamReservedBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportStreamReservedBitrate",
		val,
	)
}

func (a *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) ResetMaximumVideoBufferDelayMilliseconds() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumVideoBufferDelayMilliseconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) ResetTransportStreamReservedBitrate() {
	_jsii_.InvokeVoid(
		a,
		"resetTransportStreamReservedBitrate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMedialiveMultiplex_MultiplexSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

