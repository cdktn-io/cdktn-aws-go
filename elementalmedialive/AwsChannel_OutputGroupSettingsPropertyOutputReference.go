package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_OutputGroupSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArchiveGroupSettings() AwsChannel_ArchiveGroupSettingsPropertyList
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
	FrameCaptureGroupSettings() AwsChannel_FrameCaptureGroupSettingsPropertyOutputReference
	// Experimental.
	FrameCaptureGroupSettingsInput() *AwsChannel_FrameCaptureGroupSettingsProperty
	// Experimental.
	HlsGroupSettings() AwsChannel_HlsGroupSettingsPropertyOutputReference
	// Experimental.
	HlsGroupSettingsInput() *AwsChannel_HlsGroupSettingsProperty
	// Experimental.
	InternalValue() *AwsChannel_OutputGroupSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsChannel_OutputGroupSettingsProperty)
	// Experimental.
	MediaPackageGroupSettings() AwsChannel_MediaPackageGroupSettingsPropertyOutputReference
	// Experimental.
	MediaPackageGroupSettingsInput() *AwsChannel_MediaPackageGroupSettingsProperty
	// Experimental.
	MsSmoothGroupSettings() AwsChannel_MsSmoothGroupSettingsPropertyOutputReference
	// Experimental.
	MsSmoothGroupSettingsInput() *AwsChannel_MsSmoothGroupSettingsProperty
	// Experimental.
	MultiplexGroupSettings() AwsChannel_MultiplexGroupSettingsPropertyOutputReference
	// Experimental.
	MultiplexGroupSettingsInput() *AwsChannel_MultiplexGroupSettingsProperty
	// Experimental.
	RtmpGroupSettings() AwsChannel_RtmpGroupSettingsPropertyOutputReference
	// Experimental.
	RtmpGroupSettingsInput() *AwsChannel_RtmpGroupSettingsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UdpGroupSettings() AwsChannel_UdpGroupSettingsPropertyOutputReference
	// Experimental.
	UdpGroupSettingsInput() *AwsChannel_UdpGroupSettingsProperty
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
	PutFrameCaptureGroupSettings(value *AwsChannel_FrameCaptureGroupSettingsProperty)
	// Experimental.
	PutHlsGroupSettings(value *AwsChannel_HlsGroupSettingsProperty)
	// Experimental.
	PutMediaPackageGroupSettings(value *AwsChannel_MediaPackageGroupSettingsProperty)
	// Experimental.
	PutMsSmoothGroupSettings(value *AwsChannel_MsSmoothGroupSettingsProperty)
	// Experimental.
	PutMultiplexGroupSettings(value *AwsChannel_MultiplexGroupSettingsProperty)
	// Experimental.
	PutRtmpGroupSettings(value *AwsChannel_RtmpGroupSettingsProperty)
	// Experimental.
	PutUdpGroupSettings(value *AwsChannel_UdpGroupSettingsProperty)
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

// The jsii proxy struct for AwsChannel_OutputGroupSettingsPropertyOutputReference
type jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) ArchiveGroupSettings() AwsChannel_ArchiveGroupSettingsPropertyList {
	var returns AwsChannel_ArchiveGroupSettingsPropertyList
	_jsii_.Get(
		j,
		"archiveGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) ArchiveGroupSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archiveGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) FrameCaptureGroupSettings() AwsChannel_FrameCaptureGroupSettingsPropertyOutputReference {
	var returns AwsChannel_FrameCaptureGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"frameCaptureGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) FrameCaptureGroupSettingsInput() *AwsChannel_FrameCaptureGroupSettingsProperty {
	var returns *AwsChannel_FrameCaptureGroupSettingsProperty
	_jsii_.Get(
		j,
		"frameCaptureGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) HlsGroupSettings() AwsChannel_HlsGroupSettingsPropertyOutputReference {
	var returns AwsChannel_HlsGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hlsGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) HlsGroupSettingsInput() *AwsChannel_HlsGroupSettingsProperty {
	var returns *AwsChannel_HlsGroupSettingsProperty
	_jsii_.Get(
		j,
		"hlsGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) InternalValue() *AwsChannel_OutputGroupSettingsProperty {
	var returns *AwsChannel_OutputGroupSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) MediaPackageGroupSettings() AwsChannel_MediaPackageGroupSettingsPropertyOutputReference {
	var returns AwsChannel_MediaPackageGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"mediaPackageGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) MediaPackageGroupSettingsInput() *AwsChannel_MediaPackageGroupSettingsProperty {
	var returns *AwsChannel_MediaPackageGroupSettingsProperty
	_jsii_.Get(
		j,
		"mediaPackageGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) MsSmoothGroupSettings() AwsChannel_MsSmoothGroupSettingsPropertyOutputReference {
	var returns AwsChannel_MsSmoothGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"msSmoothGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) MsSmoothGroupSettingsInput() *AwsChannel_MsSmoothGroupSettingsProperty {
	var returns *AwsChannel_MsSmoothGroupSettingsProperty
	_jsii_.Get(
		j,
		"msSmoothGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) MultiplexGroupSettings() AwsChannel_MultiplexGroupSettingsPropertyOutputReference {
	var returns AwsChannel_MultiplexGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"multiplexGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) MultiplexGroupSettingsInput() *AwsChannel_MultiplexGroupSettingsProperty {
	var returns *AwsChannel_MultiplexGroupSettingsProperty
	_jsii_.Get(
		j,
		"multiplexGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) RtmpGroupSettings() AwsChannel_RtmpGroupSettingsPropertyOutputReference {
	var returns AwsChannel_RtmpGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rtmpGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) RtmpGroupSettingsInput() *AwsChannel_RtmpGroupSettingsProperty {
	var returns *AwsChannel_RtmpGroupSettingsProperty
	_jsii_.Get(
		j,
		"rtmpGroupSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) UdpGroupSettings() AwsChannel_UdpGroupSettingsPropertyOutputReference {
	var returns AwsChannel_UdpGroupSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"udpGroupSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) UdpGroupSettingsInput() *AwsChannel_UdpGroupSettingsProperty {
	var returns *AwsChannel_UdpGroupSettingsProperty
	_jsii_.Get(
		j,
		"udpGroupSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_OutputGroupSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChannel_OutputGroupSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_OutputGroupSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.OutputGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_OutputGroupSettingsPropertyOutputReference_Override(a AwsChannel_OutputGroupSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.OutputGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference)SetInternalValue(val *AwsChannel_OutputGroupSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) PutArchiveGroupSettings(value interface{}) {
	if err := a.validatePutArchiveGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putArchiveGroupSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) PutFrameCaptureGroupSettings(value *AwsChannel_FrameCaptureGroupSettingsProperty) {
	if err := a.validatePutFrameCaptureGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFrameCaptureGroupSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) PutHlsGroupSettings(value *AwsChannel_HlsGroupSettingsProperty) {
	if err := a.validatePutHlsGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHlsGroupSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) PutMediaPackageGroupSettings(value *AwsChannel_MediaPackageGroupSettingsProperty) {
	if err := a.validatePutMediaPackageGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMediaPackageGroupSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) PutMsSmoothGroupSettings(value *AwsChannel_MsSmoothGroupSettingsProperty) {
	if err := a.validatePutMsSmoothGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMsSmoothGroupSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) PutMultiplexGroupSettings(value *AwsChannel_MultiplexGroupSettingsProperty) {
	if err := a.validatePutMultiplexGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMultiplexGroupSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) PutRtmpGroupSettings(value *AwsChannel_RtmpGroupSettingsProperty) {
	if err := a.validatePutRtmpGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRtmpGroupSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) PutUdpGroupSettings(value *AwsChannel_UdpGroupSettingsProperty) {
	if err := a.validatePutUdpGroupSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUdpGroupSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) ResetArchiveGroupSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetArchiveGroupSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) ResetFrameCaptureGroupSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetFrameCaptureGroupSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) ResetHlsGroupSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetHlsGroupSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) ResetMediaPackageGroupSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetMediaPackageGroupSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) ResetMsSmoothGroupSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetMsSmoothGroupSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) ResetMultiplexGroupSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetMultiplexGroupSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) ResetRtmpGroupSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRtmpGroupSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) ResetUdpGroupSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetUdpGroupSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_OutputGroupSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

