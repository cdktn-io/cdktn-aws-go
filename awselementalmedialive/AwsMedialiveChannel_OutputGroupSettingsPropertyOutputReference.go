package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArchiveGroupSettings() AwsMedialiveChannel_ArchiveGroupSettingsPropertyList
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
	FrameCaptureGroupSettings() AwsMedialiveChannel_FrameCaptureGroupSettingsPropertyOutputReference
	// Experimental.
	FrameCaptureGroupSettingsInput() *AwsMedialiveChannel_FrameCaptureGroupSettingsProperty
	// Experimental.
	HlsGroupSettings() AwsMedialiveChannel_HlsGroupSettingsPropertyOutputReference
	// Experimental.
	HlsGroupSettingsInput() *AwsMedialiveChannel_HlsGroupSettingsProperty
	// Experimental.
	InternalValue() *AwsMedialiveChannel_OutputGroupSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsMedialiveChannel_OutputGroupSettingsProperty)
	// Experimental.
	MediaPackageGroupSettings() AwsMedialiveChannel_MediaPackageGroupSettingsPropertyOutputReference
	// Experimental.
	MediaPackageGroupSettingsInput() *AwsMedialiveChannel_MediaPackageGroupSettingsProperty
	// Experimental.
	MsSmoothGroupSettings() AwsMedialiveChannel_MsSmoothGroupSettingsPropertyOutputReference
	// Experimental.
	MsSmoothGroupSettingsInput() *AwsMedialiveChannel_MsSmoothGroupSettingsProperty
	// Experimental.
	MultiplexGroupSettings() AwsMedialiveChannel_MultiplexGroupSettingsPropertyOutputReference
	// Experimental.
	MultiplexGroupSettingsInput() *AwsMedialiveChannel_MultiplexGroupSettingsProperty
	// Experimental.
	RtmpGroupSettings() AwsMedialiveChannel_RtmpGroupSettingsPropertyOutputReference
	// Experimental.
	RtmpGroupSettingsInput() *AwsMedialiveChannel_RtmpGroupSettingsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UdpGroupSettings() AwsMedialiveChannel_UdpGroupSettingsPropertyOutputReference
	// Experimental.
	UdpGroupSettingsInput() *AwsMedialiveChannel_UdpGroupSettingsProperty
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
	PutFrameCaptureGroupSettings(value *AwsMedialiveChannel_FrameCaptureGroupSettingsProperty)
	// Experimental.
	PutHlsGroupSettings(value *AwsMedialiveChannel_HlsGroupSettingsProperty)
	// Experimental.
	PutMediaPackageGroupSettings(value *AwsMedialiveChannel_MediaPackageGroupSettingsProperty)
	// Experimental.
	PutMsSmoothGroupSettings(value *AwsMedialiveChannel_MsSmoothGroupSettingsProperty)
	// Experimental.
	PutMultiplexGroupSettings(value *AwsMedialiveChannel_MultiplexGroupSettingsProperty)
	// Experimental.
	PutRtmpGroupSettings(value *AwsMedialiveChannel_RtmpGroupSettingsProperty)
	// Experimental.
	PutUdpGroupSettings(value *AwsMedialiveChannel_UdpGroupSettingsProperty)
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

// The jsii proxy struct for AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference
type jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) ArchiveGroupSettings() AwsMedialiveChannel_ArchiveGroupSettingsPropertyList {
	var returns AwsMedialiveChannel_ArchiveGroupSettingsPropertyList
	_jsii_.Get(
		j,
		"archiveGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) ArchiveGroupSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archiveGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) FrameCaptureGroupSettings() AwsMedialiveChannel_FrameCaptureGroupSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_FrameCaptureGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"frameCaptureGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) FrameCaptureGroupSettingsInput() *AwsMedialiveChannel_FrameCaptureGroupSettingsProperty {
	var returns *AwsMedialiveChannel_FrameCaptureGroupSettingsProperty
	_jsii_.Get(
		j,
		"frameCaptureGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) HlsGroupSettings() AwsMedialiveChannel_HlsGroupSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_HlsGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hlsGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) HlsGroupSettingsInput() *AwsMedialiveChannel_HlsGroupSettingsProperty {
	var returns *AwsMedialiveChannel_HlsGroupSettingsProperty
	_jsii_.Get(
		j,
		"hlsGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) InternalValue() *AwsMedialiveChannel_OutputGroupSettingsProperty {
	var returns *AwsMedialiveChannel_OutputGroupSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) MediaPackageGroupSettings() AwsMedialiveChannel_MediaPackageGroupSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_MediaPackageGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"mediaPackageGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) MediaPackageGroupSettingsInput() *AwsMedialiveChannel_MediaPackageGroupSettingsProperty {
	var returns *AwsMedialiveChannel_MediaPackageGroupSettingsProperty
	_jsii_.Get(
		j,
		"mediaPackageGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) MsSmoothGroupSettings() AwsMedialiveChannel_MsSmoothGroupSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_MsSmoothGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"msSmoothGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) MsSmoothGroupSettingsInput() *AwsMedialiveChannel_MsSmoothGroupSettingsProperty {
	var returns *AwsMedialiveChannel_MsSmoothGroupSettingsProperty
	_jsii_.Get(
		j,
		"msSmoothGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) MultiplexGroupSettings() AwsMedialiveChannel_MultiplexGroupSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_MultiplexGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"multiplexGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) MultiplexGroupSettingsInput() *AwsMedialiveChannel_MultiplexGroupSettingsProperty {
	var returns *AwsMedialiveChannel_MultiplexGroupSettingsProperty
	_jsii_.Get(
		j,
		"multiplexGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) RtmpGroupSettings() AwsMedialiveChannel_RtmpGroupSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_RtmpGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rtmpGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) RtmpGroupSettingsInput() *AwsMedialiveChannel_RtmpGroupSettingsProperty {
	var returns *AwsMedialiveChannel_RtmpGroupSettingsProperty
	_jsii_.Get(
		j,
		"rtmpGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) UdpGroupSettings() AwsMedialiveChannel_UdpGroupSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_UdpGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"udpGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) UdpGroupSettingsInput() *AwsMedialiveChannel_UdpGroupSettingsProperty {
	var returns *AwsMedialiveChannel_UdpGroupSettingsProperty
	_jsii_.Get(
		j,
		"udpGroupSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMedialiveChannel_OutputGroupSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.OutputGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference_Override(a AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.OutputGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference)SetInternalValue(val *AwsMedialiveChannel_OutputGroupSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) PutArchiveGroupSettings(value interface{}) {
	if err := a.validatePutArchiveGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putArchiveGroupSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) PutFrameCaptureGroupSettings(value *AwsMedialiveChannel_FrameCaptureGroupSettingsProperty) {
	if err := a.validatePutFrameCaptureGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFrameCaptureGroupSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) PutHlsGroupSettings(value *AwsMedialiveChannel_HlsGroupSettingsProperty) {
	if err := a.validatePutHlsGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHlsGroupSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) PutMediaPackageGroupSettings(value *AwsMedialiveChannel_MediaPackageGroupSettingsProperty) {
	if err := a.validatePutMediaPackageGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMediaPackageGroupSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) PutMsSmoothGroupSettings(value *AwsMedialiveChannel_MsSmoothGroupSettingsProperty) {
	if err := a.validatePutMsSmoothGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMsSmoothGroupSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) PutMultiplexGroupSettings(value *AwsMedialiveChannel_MultiplexGroupSettingsProperty) {
	if err := a.validatePutMultiplexGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMultiplexGroupSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) PutRtmpGroupSettings(value *AwsMedialiveChannel_RtmpGroupSettingsProperty) {
	if err := a.validatePutRtmpGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRtmpGroupSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) PutUdpGroupSettings(value *AwsMedialiveChannel_UdpGroupSettingsProperty) {
	if err := a.validatePutUdpGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUdpGroupSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) ResetArchiveGroupSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetArchiveGroupSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) ResetFrameCaptureGroupSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetFrameCaptureGroupSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) ResetHlsGroupSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetHlsGroupSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) ResetMediaPackageGroupSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetMediaPackageGroupSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) ResetMsSmoothGroupSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetMsSmoothGroupSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) ResetMultiplexGroupSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetMultiplexGroupSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) ResetRtmpGroupSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRtmpGroupSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) ResetUdpGroupSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetUdpGroupSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputGroupSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

