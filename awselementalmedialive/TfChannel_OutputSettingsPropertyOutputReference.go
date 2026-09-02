package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_OutputSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArchiveOutputSettings() TfChannel_ArchiveOutputSettingsPropertyOutputReference
	// Experimental.
	ArchiveOutputSettingsInput() *TfChannel_ArchiveOutputSettingsProperty
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
	FrameCaptureOutputSettings() TfChannel_FrameCaptureOutputSettingsPropertyOutputReference
	// Experimental.
	FrameCaptureOutputSettingsInput() *TfChannel_FrameCaptureOutputSettingsProperty
	// Experimental.
	HlsOutputSettings() TfChannel_HlsOutputSettingsPropertyOutputReference
	// Experimental.
	HlsOutputSettingsInput() *TfChannel_HlsOutputSettingsProperty
	// Experimental.
	InternalValue() *TfChannel_OutputSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_OutputSettingsProperty)
	// Experimental.
	MediaPackageOutputSettings() TfChannel_MediaPackageOutputSettingsPropertyOutputReference
	// Experimental.
	MediaPackageOutputSettingsInput() *TfChannel_MediaPackageOutputSettingsProperty
	// Experimental.
	MsSmoothOutputSettings() TfChannel_MsSmoothOutputSettingsPropertyOutputReference
	// Experimental.
	MsSmoothOutputSettingsInput() *TfChannel_MsSmoothOutputSettingsProperty
	// Experimental.
	MultiplexOutputSettings() TfChannel_MultiplexOutputSettingsPropertyOutputReference
	// Experimental.
	MultiplexOutputSettingsInput() *TfChannel_MultiplexOutputSettingsProperty
	// Experimental.
	RtmpOutputSettings() TfChannel_RtmpOutputSettingsPropertyOutputReference
	// Experimental.
	RtmpOutputSettingsInput() *TfChannel_RtmpOutputSettingsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UdpOutputSettings() TfChannel_UdpOutputSettingsPropertyOutputReference
	// Experimental.
	UdpOutputSettingsInput() *TfChannel_UdpOutputSettingsProperty
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
	PutArchiveOutputSettings(value *TfChannel_ArchiveOutputSettingsProperty)
	// Experimental.
	PutFrameCaptureOutputSettings(value *TfChannel_FrameCaptureOutputSettingsProperty)
	// Experimental.
	PutHlsOutputSettings(value *TfChannel_HlsOutputSettingsProperty)
	// Experimental.
	PutMediaPackageOutputSettings(value *TfChannel_MediaPackageOutputSettingsProperty)
	// Experimental.
	PutMsSmoothOutputSettings(value *TfChannel_MsSmoothOutputSettingsProperty)
	// Experimental.
	PutMultiplexOutputSettings(value *TfChannel_MultiplexOutputSettingsProperty)
	// Experimental.
	PutRtmpOutputSettings(value *TfChannel_RtmpOutputSettingsProperty)
	// Experimental.
	PutUdpOutputSettings(value *TfChannel_UdpOutputSettingsProperty)
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

// The jsii proxy struct for TfChannel_OutputSettingsPropertyOutputReference
type jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) ArchiveOutputSettings() TfChannel_ArchiveOutputSettingsPropertyOutputReference {
	var returns TfChannel_ArchiveOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"archiveOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) ArchiveOutputSettingsInput() *TfChannel_ArchiveOutputSettingsProperty {
	var returns *TfChannel_ArchiveOutputSettingsProperty
	_jsii_.Get(
		j,
		"archiveOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) FrameCaptureOutputSettings() TfChannel_FrameCaptureOutputSettingsPropertyOutputReference {
	var returns TfChannel_FrameCaptureOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"frameCaptureOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) FrameCaptureOutputSettingsInput() *TfChannel_FrameCaptureOutputSettingsProperty {
	var returns *TfChannel_FrameCaptureOutputSettingsProperty
	_jsii_.Get(
		j,
		"frameCaptureOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) HlsOutputSettings() TfChannel_HlsOutputSettingsPropertyOutputReference {
	var returns TfChannel_HlsOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hlsOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) HlsOutputSettingsInput() *TfChannel_HlsOutputSettingsProperty {
	var returns *TfChannel_HlsOutputSettingsProperty
	_jsii_.Get(
		j,
		"hlsOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) InternalValue() *TfChannel_OutputSettingsProperty {
	var returns *TfChannel_OutputSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) MediaPackageOutputSettings() TfChannel_MediaPackageOutputSettingsPropertyOutputReference {
	var returns TfChannel_MediaPackageOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"mediaPackageOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) MediaPackageOutputSettingsInput() *TfChannel_MediaPackageOutputSettingsProperty {
	var returns *TfChannel_MediaPackageOutputSettingsProperty
	_jsii_.Get(
		j,
		"mediaPackageOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) MsSmoothOutputSettings() TfChannel_MsSmoothOutputSettingsPropertyOutputReference {
	var returns TfChannel_MsSmoothOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"msSmoothOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) MsSmoothOutputSettingsInput() *TfChannel_MsSmoothOutputSettingsProperty {
	var returns *TfChannel_MsSmoothOutputSettingsProperty
	_jsii_.Get(
		j,
		"msSmoothOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) MultiplexOutputSettings() TfChannel_MultiplexOutputSettingsPropertyOutputReference {
	var returns TfChannel_MultiplexOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"multiplexOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) MultiplexOutputSettingsInput() *TfChannel_MultiplexOutputSettingsProperty {
	var returns *TfChannel_MultiplexOutputSettingsProperty
	_jsii_.Get(
		j,
		"multiplexOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) RtmpOutputSettings() TfChannel_RtmpOutputSettingsPropertyOutputReference {
	var returns TfChannel_RtmpOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rtmpOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) RtmpOutputSettingsInput() *TfChannel_RtmpOutputSettingsProperty {
	var returns *TfChannel_RtmpOutputSettingsProperty
	_jsii_.Get(
		j,
		"rtmpOutputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) UdpOutputSettings() TfChannel_UdpOutputSettingsPropertyOutputReference {
	var returns TfChannel_UdpOutputSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"udpOutputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) UdpOutputSettingsInput() *TfChannel_UdpOutputSettingsProperty {
	var returns *TfChannel_UdpOutputSettingsProperty
	_jsii_.Get(
		j,
		"udpOutputSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_OutputSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_OutputSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_OutputSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.OutputSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_OutputSettingsPropertyOutputReference_Override(t TfChannel_OutputSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.OutputSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_OutputSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) PutArchiveOutputSettings(value *TfChannel_ArchiveOutputSettingsProperty) {
	if err := t.validatePutArchiveOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putArchiveOutputSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) PutFrameCaptureOutputSettings(value *TfChannel_FrameCaptureOutputSettingsProperty) {
	if err := t.validatePutFrameCaptureOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFrameCaptureOutputSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) PutHlsOutputSettings(value *TfChannel_HlsOutputSettingsProperty) {
	if err := t.validatePutHlsOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHlsOutputSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) PutMediaPackageOutputSettings(value *TfChannel_MediaPackageOutputSettingsProperty) {
	if err := t.validatePutMediaPackageOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMediaPackageOutputSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) PutMsSmoothOutputSettings(value *TfChannel_MsSmoothOutputSettingsProperty) {
	if err := t.validatePutMsSmoothOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMsSmoothOutputSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) PutMultiplexOutputSettings(value *TfChannel_MultiplexOutputSettingsProperty) {
	if err := t.validatePutMultiplexOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMultiplexOutputSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) PutRtmpOutputSettings(value *TfChannel_RtmpOutputSettingsProperty) {
	if err := t.validatePutRtmpOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRtmpOutputSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) PutUdpOutputSettings(value *TfChannel_UdpOutputSettingsProperty) {
	if err := t.validatePutUdpOutputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUdpOutputSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) ResetArchiveOutputSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetArchiveOutputSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) ResetFrameCaptureOutputSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetFrameCaptureOutputSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) ResetHlsOutputSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetHlsOutputSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) ResetMediaPackageOutputSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetMediaPackageOutputSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) ResetMsSmoothOutputSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetMsSmoothOutputSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) ResetMultiplexOutputSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetMultiplexOutputSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) ResetRtmpOutputSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetRtmpOutputSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) ResetUdpOutputSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetUdpOutputSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_OutputSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

