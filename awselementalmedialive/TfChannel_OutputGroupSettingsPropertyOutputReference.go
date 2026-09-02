package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_OutputGroupSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArchiveGroupSettings() TfChannel_ArchiveGroupSettingsPropertyList
	// Experimental.
	ArchiveGroupSettingsInput() interface{}
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
	FrameCaptureGroupSettings() TfChannel_FrameCaptureGroupSettingsPropertyOutputReference
	// Experimental.
	FrameCaptureGroupSettingsInput() *TfChannel_FrameCaptureGroupSettingsProperty
	// Experimental.
	HlsGroupSettings() TfChannel_HlsGroupSettingsPropertyOutputReference
	// Experimental.
	HlsGroupSettingsInput() *TfChannel_HlsGroupSettingsProperty
	// Experimental.
	InternalValue() *TfChannel_OutputGroupSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_OutputGroupSettingsProperty)
	// Experimental.
	MediaPackageGroupSettings() TfChannel_MediaPackageGroupSettingsPropertyOutputReference
	// Experimental.
	MediaPackageGroupSettingsInput() *TfChannel_MediaPackageGroupSettingsProperty
	// Experimental.
	MsSmoothGroupSettings() TfChannel_MsSmoothGroupSettingsPropertyOutputReference
	// Experimental.
	MsSmoothGroupSettingsInput() *TfChannel_MsSmoothGroupSettingsProperty
	// Experimental.
	MultiplexGroupSettings() TfChannel_MultiplexGroupSettingsPropertyOutputReference
	// Experimental.
	MultiplexGroupSettingsInput() *TfChannel_MultiplexGroupSettingsProperty
	// Experimental.
	RtmpGroupSettings() TfChannel_RtmpGroupSettingsPropertyOutputReference
	// Experimental.
	RtmpGroupSettingsInput() *TfChannel_RtmpGroupSettingsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UdpGroupSettings() TfChannel_UdpGroupSettingsPropertyOutputReference
	// Experimental.
	UdpGroupSettingsInput() *TfChannel_UdpGroupSettingsProperty
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
	PutArchiveGroupSettings(value interface{})
	// Experimental.
	PutFrameCaptureGroupSettings(value *TfChannel_FrameCaptureGroupSettingsProperty)
	// Experimental.
	PutHlsGroupSettings(value *TfChannel_HlsGroupSettingsProperty)
	// Experimental.
	PutMediaPackageGroupSettings(value *TfChannel_MediaPackageGroupSettingsProperty)
	// Experimental.
	PutMsSmoothGroupSettings(value *TfChannel_MsSmoothGroupSettingsProperty)
	// Experimental.
	PutMultiplexGroupSettings(value *TfChannel_MultiplexGroupSettingsProperty)
	// Experimental.
	PutRtmpGroupSettings(value *TfChannel_RtmpGroupSettingsProperty)
	// Experimental.
	PutUdpGroupSettings(value *TfChannel_UdpGroupSettingsProperty)
	// Experimental.
	ResetArchiveGroupSettings()
	// Experimental.
	ResetFrameCaptureGroupSettings()
	// Experimental.
	ResetHlsGroupSettings()
	// Experimental.
	ResetMediaPackageGroupSettings()
	// Experimental.
	ResetMsSmoothGroupSettings()
	// Experimental.
	ResetMultiplexGroupSettings()
	// Experimental.
	ResetRtmpGroupSettings()
	// Experimental.
	ResetUdpGroupSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_OutputGroupSettingsPropertyOutputReference
type jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) ArchiveGroupSettings() TfChannel_ArchiveGroupSettingsPropertyList {
	var returns TfChannel_ArchiveGroupSettingsPropertyList
	_jsii_.Get(
		j,
		"archiveGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) ArchiveGroupSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archiveGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) FrameCaptureGroupSettings() TfChannel_FrameCaptureGroupSettingsPropertyOutputReference {
	var returns TfChannel_FrameCaptureGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"frameCaptureGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) FrameCaptureGroupSettingsInput() *TfChannel_FrameCaptureGroupSettingsProperty {
	var returns *TfChannel_FrameCaptureGroupSettingsProperty
	_jsii_.Get(
		j,
		"frameCaptureGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) HlsGroupSettings() TfChannel_HlsGroupSettingsPropertyOutputReference {
	var returns TfChannel_HlsGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hlsGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) HlsGroupSettingsInput() *TfChannel_HlsGroupSettingsProperty {
	var returns *TfChannel_HlsGroupSettingsProperty
	_jsii_.Get(
		j,
		"hlsGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) InternalValue() *TfChannel_OutputGroupSettingsProperty {
	var returns *TfChannel_OutputGroupSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) MediaPackageGroupSettings() TfChannel_MediaPackageGroupSettingsPropertyOutputReference {
	var returns TfChannel_MediaPackageGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"mediaPackageGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) MediaPackageGroupSettingsInput() *TfChannel_MediaPackageGroupSettingsProperty {
	var returns *TfChannel_MediaPackageGroupSettingsProperty
	_jsii_.Get(
		j,
		"mediaPackageGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) MsSmoothGroupSettings() TfChannel_MsSmoothGroupSettingsPropertyOutputReference {
	var returns TfChannel_MsSmoothGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"msSmoothGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) MsSmoothGroupSettingsInput() *TfChannel_MsSmoothGroupSettingsProperty {
	var returns *TfChannel_MsSmoothGroupSettingsProperty
	_jsii_.Get(
		j,
		"msSmoothGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) MultiplexGroupSettings() TfChannel_MultiplexGroupSettingsPropertyOutputReference {
	var returns TfChannel_MultiplexGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"multiplexGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) MultiplexGroupSettingsInput() *TfChannel_MultiplexGroupSettingsProperty {
	var returns *TfChannel_MultiplexGroupSettingsProperty
	_jsii_.Get(
		j,
		"multiplexGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) RtmpGroupSettings() TfChannel_RtmpGroupSettingsPropertyOutputReference {
	var returns TfChannel_RtmpGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rtmpGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) RtmpGroupSettingsInput() *TfChannel_RtmpGroupSettingsProperty {
	var returns *TfChannel_RtmpGroupSettingsProperty
	_jsii_.Get(
		j,
		"rtmpGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) UdpGroupSettings() TfChannel_UdpGroupSettingsPropertyOutputReference {
	var returns TfChannel_UdpGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"udpGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) UdpGroupSettingsInput() *TfChannel_UdpGroupSettingsProperty {
	var returns *TfChannel_UdpGroupSettingsProperty
	_jsii_.Get(
		j,
		"udpGroupSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_OutputGroupSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_OutputGroupSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_OutputGroupSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.OutputGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_OutputGroupSettingsPropertyOutputReference_Override(t TfChannel_OutputGroupSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.OutputGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_OutputGroupSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) PutArchiveGroupSettings(value interface{}) {
	if err := t.validatePutArchiveGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putArchiveGroupSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) PutFrameCaptureGroupSettings(value *TfChannel_FrameCaptureGroupSettingsProperty) {
	if err := t.validatePutFrameCaptureGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFrameCaptureGroupSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) PutHlsGroupSettings(value *TfChannel_HlsGroupSettingsProperty) {
	if err := t.validatePutHlsGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHlsGroupSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) PutMediaPackageGroupSettings(value *TfChannel_MediaPackageGroupSettingsProperty) {
	if err := t.validatePutMediaPackageGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMediaPackageGroupSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) PutMsSmoothGroupSettings(value *TfChannel_MsSmoothGroupSettingsProperty) {
	if err := t.validatePutMsSmoothGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMsSmoothGroupSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) PutMultiplexGroupSettings(value *TfChannel_MultiplexGroupSettingsProperty) {
	if err := t.validatePutMultiplexGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMultiplexGroupSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) PutRtmpGroupSettings(value *TfChannel_RtmpGroupSettingsProperty) {
	if err := t.validatePutRtmpGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRtmpGroupSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) PutUdpGroupSettings(value *TfChannel_UdpGroupSettingsProperty) {
	if err := t.validatePutUdpGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUdpGroupSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) ResetArchiveGroupSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetArchiveGroupSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) ResetFrameCaptureGroupSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetFrameCaptureGroupSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) ResetHlsGroupSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetHlsGroupSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) ResetMediaPackageGroupSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetMediaPackageGroupSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) ResetMsSmoothGroupSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetMsSmoothGroupSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) ResetMultiplexGroupSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetMultiplexGroupSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) ResetRtmpGroupSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetRtmpGroupSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) ResetUdpGroupSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetUdpGroupSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_OutputGroupSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

