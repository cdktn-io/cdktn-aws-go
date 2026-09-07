package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_OutputSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArchiveOutputSettings() AwsChannel_ArchiveOutputSettingsPropertyOutputReference
	// Experimental.
	ArchiveOutputSettingsInput() *AwsChannel_ArchiveOutputSettingsProperty
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
	FrameCaptureOutputSettings() AwsChannel_FrameCaptureOutputSettingsPropertyOutputReference
	// Experimental.
	FrameCaptureOutputSettingsInput() *AwsChannel_FrameCaptureOutputSettingsProperty
	// Experimental.
	HlsOutputSettings() AwsChannel_HlsOutputSettingsPropertyOutputReference
	// Experimental.
	HlsOutputSettingsInput() *AwsChannel_HlsOutputSettingsProperty
	// Experimental.
	InternalValue() *AwsChannel_OutputSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsChannel_OutputSettingsProperty)
	// Experimental.
	MediaPackageOutputSettings() AwsChannel_MediaPackageOutputSettingsPropertyOutputReference
	// Experimental.
	MediaPackageOutputSettingsInput() *AwsChannel_MediaPackageOutputSettingsProperty
	// Experimental.
	MsSmoothOutputSettings() AwsChannel_MsSmoothOutputSettingsPropertyOutputReference
	// Experimental.
	MsSmoothOutputSettingsInput() *AwsChannel_MsSmoothOutputSettingsProperty
	// Experimental.
	MultiplexOutputSettings() AwsChannel_MultiplexOutputSettingsPropertyOutputReference
	// Experimental.
	MultiplexOutputSettingsInput() *AwsChannel_MultiplexOutputSettingsProperty
	// Experimental.
	RtmpOutputSettings() AwsChannel_RtmpOutputSettingsPropertyOutputReference
	// Experimental.
	RtmpOutputSettingsInput() *AwsChannel_RtmpOutputSettingsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UdpOutputSettings() AwsChannel_UdpOutputSettingsPropertyOutputReference
	// Experimental.
	UdpOutputSettingsInput() *AwsChannel_UdpOutputSettingsProperty
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
	PutArchiveOutputSettings(value *AwsChannel_ArchiveOutputSettingsProperty)
	// Experimental.
	PutFrameCaptureOutputSettings(value *AwsChannel_FrameCaptureOutputSettingsProperty)
	// Experimental.
	PutHlsOutputSettings(value *AwsChannel_HlsOutputSettingsProperty)
	// Experimental.
	PutMediaPackageOutputSettings(value *AwsChannel_MediaPackageOutputSettingsProperty)
	// Experimental.
	PutMsSmoothOutputSettings(value *AwsChannel_MsSmoothOutputSettingsProperty)
	// Experimental.
	PutMultiplexOutputSettings(value *AwsChannel_MultiplexOutputSettingsProperty)
	// Experimental.
	PutRtmpOutputSettings(value *AwsChannel_RtmpOutputSettingsProperty)
	// Experimental.
	PutUdpOutputSettings(value *AwsChannel_UdpOutputSettingsProperty)
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

// The jsii proxy struct for AwsChannel_OutputSettingsPropertyOutputReference
type jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) ArchiveOutputSettings() AwsChannel_ArchiveOutputSettingsPropertyOutputReference {
	var returns AwsChannel_ArchiveOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"archiveOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) ArchiveOutputSettingsInput() *AwsChannel_ArchiveOutputSettingsProperty {
	var returns *AwsChannel_ArchiveOutputSettingsProperty
	_jsii_.Get(
		j,
		"archiveOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) FrameCaptureOutputSettings() AwsChannel_FrameCaptureOutputSettingsPropertyOutputReference {
	var returns AwsChannel_FrameCaptureOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"frameCaptureOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) FrameCaptureOutputSettingsInput() *AwsChannel_FrameCaptureOutputSettingsProperty {
	var returns *AwsChannel_FrameCaptureOutputSettingsProperty
	_jsii_.Get(
		j,
		"frameCaptureOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) HlsOutputSettings() AwsChannel_HlsOutputSettingsPropertyOutputReference {
	var returns AwsChannel_HlsOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hlsOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) HlsOutputSettingsInput() *AwsChannel_HlsOutputSettingsProperty {
	var returns *AwsChannel_HlsOutputSettingsProperty
	_jsii_.Get(
		j,
		"hlsOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) InternalValue() *AwsChannel_OutputSettingsProperty {
	var returns *AwsChannel_OutputSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) MediaPackageOutputSettings() AwsChannel_MediaPackageOutputSettingsPropertyOutputReference {
	var returns AwsChannel_MediaPackageOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"mediaPackageOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) MediaPackageOutputSettingsInput() *AwsChannel_MediaPackageOutputSettingsProperty {
	var returns *AwsChannel_MediaPackageOutputSettingsProperty
	_jsii_.Get(
		j,
		"mediaPackageOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) MsSmoothOutputSettings() AwsChannel_MsSmoothOutputSettingsPropertyOutputReference {
	var returns AwsChannel_MsSmoothOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"msSmoothOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) MsSmoothOutputSettingsInput() *AwsChannel_MsSmoothOutputSettingsProperty {
	var returns *AwsChannel_MsSmoothOutputSettingsProperty
	_jsii_.Get(
		j,
		"msSmoothOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) MultiplexOutputSettings() AwsChannel_MultiplexOutputSettingsPropertyOutputReference {
	var returns AwsChannel_MultiplexOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"multiplexOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) MultiplexOutputSettingsInput() *AwsChannel_MultiplexOutputSettingsProperty {
	var returns *AwsChannel_MultiplexOutputSettingsProperty
	_jsii_.Get(
		j,
		"multiplexOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) RtmpOutputSettings() AwsChannel_RtmpOutputSettingsPropertyOutputReference {
	var returns AwsChannel_RtmpOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rtmpOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) RtmpOutputSettingsInput() *AwsChannel_RtmpOutputSettingsProperty {
	var returns *AwsChannel_RtmpOutputSettingsProperty
	_jsii_.Get(
		j,
		"rtmpOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) UdpOutputSettings() AwsChannel_UdpOutputSettingsPropertyOutputReference {
	var returns AwsChannel_UdpOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"udpOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) UdpOutputSettingsInput() *AwsChannel_UdpOutputSettingsProperty {
	var returns *AwsChannel_UdpOutputSettingsProperty
	_jsii_.Get(
		j,
		"udpOutputSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_OutputSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChannel_OutputSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_OutputSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.OutputSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_OutputSettingsPropertyOutputReference_Override(a AwsChannel_OutputSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.OutputSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference)SetInternalValue(val *AwsChannel_OutputSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) PutArchiveOutputSettings(value *AwsChannel_ArchiveOutputSettingsProperty) {
	if err := a.validatePutArchiveOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putArchiveOutputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) PutFrameCaptureOutputSettings(value *AwsChannel_FrameCaptureOutputSettingsProperty) {
	if err := a.validatePutFrameCaptureOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFrameCaptureOutputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) PutHlsOutputSettings(value *AwsChannel_HlsOutputSettingsProperty) {
	if err := a.validatePutHlsOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHlsOutputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) PutMediaPackageOutputSettings(value *AwsChannel_MediaPackageOutputSettingsProperty) {
	if err := a.validatePutMediaPackageOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMediaPackageOutputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) PutMsSmoothOutputSettings(value *AwsChannel_MsSmoothOutputSettingsProperty) {
	if err := a.validatePutMsSmoothOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMsSmoothOutputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) PutMultiplexOutputSettings(value *AwsChannel_MultiplexOutputSettingsProperty) {
	if err := a.validatePutMultiplexOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMultiplexOutputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) PutRtmpOutputSettings(value *AwsChannel_RtmpOutputSettingsProperty) {
	if err := a.validatePutRtmpOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRtmpOutputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) PutUdpOutputSettings(value *AwsChannel_UdpOutputSettingsProperty) {
	if err := a.validatePutUdpOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUdpOutputSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) ResetArchiveOutputSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetArchiveOutputSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) ResetFrameCaptureOutputSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetFrameCaptureOutputSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) ResetHlsOutputSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetHlsOutputSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) ResetMediaPackageOutputSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetMediaPackageOutputSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) ResetMsSmoothOutputSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetMsSmoothOutputSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) ResetMultiplexOutputSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetMultiplexOutputSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) ResetRtmpOutputSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRtmpOutputSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) ResetUdpOutputSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetUdpOutputSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_OutputSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

