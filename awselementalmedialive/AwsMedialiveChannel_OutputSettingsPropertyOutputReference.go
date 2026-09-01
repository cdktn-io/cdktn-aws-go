package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMedialiveChannel_OutputSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArchiveOutputSettings() AwsMedialiveChannel_ArchiveOutputSettingsPropertyOutputReference
	// Experimental.
	ArchiveOutputSettingsInput() *AwsMedialiveChannel_ArchiveOutputSettingsProperty
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
	FrameCaptureOutputSettings() AwsMedialiveChannel_FrameCaptureOutputSettingsPropertyOutputReference
	// Experimental.
	FrameCaptureOutputSettingsInput() *AwsMedialiveChannel_FrameCaptureOutputSettingsProperty
	// Experimental.
	HlsOutputSettings() AwsMedialiveChannel_HlsOutputSettingsPropertyOutputReference
	// Experimental.
	HlsOutputSettingsInput() *AwsMedialiveChannel_HlsOutputSettingsProperty
	// Experimental.
	InternalValue() *AwsMedialiveChannel_OutputSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsMedialiveChannel_OutputSettingsProperty)
	// Experimental.
	MediaPackageOutputSettings() AwsMedialiveChannel_MediaPackageOutputSettingsPropertyOutputReference
	// Experimental.
	MediaPackageOutputSettingsInput() *AwsMedialiveChannel_MediaPackageOutputSettingsProperty
	// Experimental.
	MsSmoothOutputSettings() AwsMedialiveChannel_MsSmoothOutputSettingsPropertyOutputReference
	// Experimental.
	MsSmoothOutputSettingsInput() *AwsMedialiveChannel_MsSmoothOutputSettingsProperty
	// Experimental.
	MultiplexOutputSettings() AwsMedialiveChannel_MultiplexOutputSettingsPropertyOutputReference
	// Experimental.
	MultiplexOutputSettingsInput() *AwsMedialiveChannel_MultiplexOutputSettingsProperty
	// Experimental.
	RtmpOutputSettings() AwsMedialiveChannel_RtmpOutputSettingsPropertyOutputReference
	// Experimental.
	RtmpOutputSettingsInput() *AwsMedialiveChannel_RtmpOutputSettingsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UdpOutputSettings() AwsMedialiveChannel_UdpOutputSettingsPropertyOutputReference
	// Experimental.
	UdpOutputSettingsInput() *AwsMedialiveChannel_UdpOutputSettingsProperty
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
	PutArchiveOutputSettings(value *AwsMedialiveChannel_ArchiveOutputSettingsProperty)
	// Experimental.
	PutFrameCaptureOutputSettings(value *AwsMedialiveChannel_FrameCaptureOutputSettingsProperty)
	// Experimental.
	PutHlsOutputSettings(value *AwsMedialiveChannel_HlsOutputSettingsProperty)
	// Experimental.
	PutMediaPackageOutputSettings(value *AwsMedialiveChannel_MediaPackageOutputSettingsProperty)
	// Experimental.
	PutMsSmoothOutputSettings(value *AwsMedialiveChannel_MsSmoothOutputSettingsProperty)
	// Experimental.
	PutMultiplexOutputSettings(value *AwsMedialiveChannel_MultiplexOutputSettingsProperty)
	// Experimental.
	PutRtmpOutputSettings(value *AwsMedialiveChannel_RtmpOutputSettingsProperty)
	// Experimental.
	PutUdpOutputSettings(value *AwsMedialiveChannel_UdpOutputSettingsProperty)
	// Experimental.
	ResetArchiveOutputSettings()
	// Experimental.
	ResetFrameCaptureOutputSettings()
	// Experimental.
	ResetHlsOutputSettings()
	// Experimental.
	ResetMediaPackageOutputSettings()
	// Experimental.
	ResetMsSmoothOutputSettings()
	// Experimental.
	ResetMultiplexOutputSettings()
	// Experimental.
	ResetRtmpOutputSettings()
	// Experimental.
	ResetUdpOutputSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMedialiveChannel_OutputSettingsPropertyOutputReference
type jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) ArchiveOutputSettings() AwsMedialiveChannel_ArchiveOutputSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_ArchiveOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"archiveOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) ArchiveOutputSettingsInput() *AwsMedialiveChannel_ArchiveOutputSettingsProperty {
	var returns *AwsMedialiveChannel_ArchiveOutputSettingsProperty
	_jsii_.Get(
		j,
		"archiveOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) FrameCaptureOutputSettings() AwsMedialiveChannel_FrameCaptureOutputSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_FrameCaptureOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"frameCaptureOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) FrameCaptureOutputSettingsInput() *AwsMedialiveChannel_FrameCaptureOutputSettingsProperty {
	var returns *AwsMedialiveChannel_FrameCaptureOutputSettingsProperty
	_jsii_.Get(
		j,
		"frameCaptureOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) HlsOutputSettings() AwsMedialiveChannel_HlsOutputSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_HlsOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hlsOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) HlsOutputSettingsInput() *AwsMedialiveChannel_HlsOutputSettingsProperty {
	var returns *AwsMedialiveChannel_HlsOutputSettingsProperty
	_jsii_.Get(
		j,
		"hlsOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) InternalValue() *AwsMedialiveChannel_OutputSettingsProperty {
	var returns *AwsMedialiveChannel_OutputSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) MediaPackageOutputSettings() AwsMedialiveChannel_MediaPackageOutputSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_MediaPackageOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"mediaPackageOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) MediaPackageOutputSettingsInput() *AwsMedialiveChannel_MediaPackageOutputSettingsProperty {
	var returns *AwsMedialiveChannel_MediaPackageOutputSettingsProperty
	_jsii_.Get(
		j,
		"mediaPackageOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) MsSmoothOutputSettings() AwsMedialiveChannel_MsSmoothOutputSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_MsSmoothOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"msSmoothOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) MsSmoothOutputSettingsInput() *AwsMedialiveChannel_MsSmoothOutputSettingsProperty {
	var returns *AwsMedialiveChannel_MsSmoothOutputSettingsProperty
	_jsii_.Get(
		j,
		"msSmoothOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) MultiplexOutputSettings() AwsMedialiveChannel_MultiplexOutputSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_MultiplexOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"multiplexOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) MultiplexOutputSettingsInput() *AwsMedialiveChannel_MultiplexOutputSettingsProperty {
	var returns *AwsMedialiveChannel_MultiplexOutputSettingsProperty
	_jsii_.Get(
		j,
		"multiplexOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) RtmpOutputSettings() AwsMedialiveChannel_RtmpOutputSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_RtmpOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rtmpOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) RtmpOutputSettingsInput() *AwsMedialiveChannel_RtmpOutputSettingsProperty {
	var returns *AwsMedialiveChannel_RtmpOutputSettingsProperty
	_jsii_.Get(
		j,
		"rtmpOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) UdpOutputSettings() AwsMedialiveChannel_UdpOutputSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_UdpOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"udpOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) UdpOutputSettingsInput() *AwsMedialiveChannel_UdpOutputSettingsProperty {
	var returns *AwsMedialiveChannel_UdpOutputSettingsProperty
	_jsii_.Get(
		j,
		"udpOutputSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMedialiveChannel_OutputSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMedialiveChannel_OutputSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMedialiveChannel_OutputSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.OutputSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMedialiveChannel_OutputSettingsPropertyOutputReference_Override(a AwsMedialiveChannel_OutputSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.OutputSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference)SetInternalValue(val *AwsMedialiveChannel_OutputSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) PutArchiveOutputSettings(value *AwsMedialiveChannel_ArchiveOutputSettingsProperty) {
	if err := a.validatePutArchiveOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putArchiveOutputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) PutFrameCaptureOutputSettings(value *AwsMedialiveChannel_FrameCaptureOutputSettingsProperty) {
	if err := a.validatePutFrameCaptureOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFrameCaptureOutputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) PutHlsOutputSettings(value *AwsMedialiveChannel_HlsOutputSettingsProperty) {
	if err := a.validatePutHlsOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHlsOutputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) PutMediaPackageOutputSettings(value *AwsMedialiveChannel_MediaPackageOutputSettingsProperty) {
	if err := a.validatePutMediaPackageOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMediaPackageOutputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) PutMsSmoothOutputSettings(value *AwsMedialiveChannel_MsSmoothOutputSettingsProperty) {
	if err := a.validatePutMsSmoothOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMsSmoothOutputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) PutMultiplexOutputSettings(value *AwsMedialiveChannel_MultiplexOutputSettingsProperty) {
	if err := a.validatePutMultiplexOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMultiplexOutputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) PutRtmpOutputSettings(value *AwsMedialiveChannel_RtmpOutputSettingsProperty) {
	if err := a.validatePutRtmpOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRtmpOutputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) PutUdpOutputSettings(value *AwsMedialiveChannel_UdpOutputSettingsProperty) {
	if err := a.validatePutUdpOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUdpOutputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) ResetArchiveOutputSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetArchiveOutputSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) ResetFrameCaptureOutputSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetFrameCaptureOutputSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) ResetHlsOutputSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetHlsOutputSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) ResetMediaPackageOutputSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetMediaPackageOutputSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) ResetMsSmoothOutputSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetMsSmoothOutputSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) ResetMultiplexOutputSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetMultiplexOutputSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) ResetRtmpOutputSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRtmpOutputSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) ResetUdpOutputSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetUdpOutputSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_OutputSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

