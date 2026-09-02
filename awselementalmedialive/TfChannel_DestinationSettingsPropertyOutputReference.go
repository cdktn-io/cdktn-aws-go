package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_DestinationSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AribDestinationSettings() TfChannel_AribDestinationSettingsPropertyOutputReference
	// Experimental.
	AribDestinationSettingsInput() *TfChannel_AribDestinationSettingsProperty
	// Experimental.
	BurnInDestinationSettings() TfChannel_BurnInDestinationSettingsPropertyOutputReference
	// Experimental.
	BurnInDestinationSettingsInput() *TfChannel_BurnInDestinationSettingsProperty
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
	DvbSubDestinationSettings() TfChannel_DvbSubDestinationSettingsPropertyOutputReference
	// Experimental.
	DvbSubDestinationSettingsInput() *TfChannel_DvbSubDestinationSettingsProperty
	// Experimental.
	EbuTtDDestinationSettings() TfChannel_EbuTtDDestinationSettingsPropertyOutputReference
	// Experimental.
	EbuTtDDestinationSettingsInput() *TfChannel_EbuTtDDestinationSettingsProperty
	// Experimental.
	EmbeddedDestinationSettings() TfChannel_EmbeddedDestinationSettingsPropertyOutputReference
	// Experimental.
	EmbeddedDestinationSettingsInput() *TfChannel_EmbeddedDestinationSettingsProperty
	// Experimental.
	EmbeddedPlusScte20DestinationSettings() TfChannel_EmbeddedPlusScte20DestinationSettingsPropertyOutputReference
	// Experimental.
	EmbeddedPlusScte20DestinationSettingsInput() *TfChannel_EmbeddedPlusScte20DestinationSettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfChannel_DestinationSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_DestinationSettingsProperty)
	// Experimental.
	RtmpCaptionInfoDestinationSettings() TfChannel_RtmpCaptionInfoDestinationSettingsPropertyOutputReference
	// Experimental.
	RtmpCaptionInfoDestinationSettingsInput() *TfChannel_RtmpCaptionInfoDestinationSettingsProperty
	// Experimental.
	Scte20PlusEmbeddedDestinationSettings() TfChannel_Scte20PlusEmbeddedDestinationSettingsPropertyOutputReference
	// Experimental.
	Scte20PlusEmbeddedDestinationSettingsInput() *TfChannel_Scte20PlusEmbeddedDestinationSettingsProperty
	// Experimental.
	Scte27DestinationSettings() TfChannel_Scte27DestinationSettingsPropertyOutputReference
	// Experimental.
	Scte27DestinationSettingsInput() *TfChannel_Scte27DestinationSettingsProperty
	// Experimental.
	SmpteTtDestinationSettings() TfChannel_SmpteTtDestinationSettingsPropertyOutputReference
	// Experimental.
	SmpteTtDestinationSettingsInput() *TfChannel_SmpteTtDestinationSettingsProperty
	// Experimental.
	TeletextDestinationSettings() TfChannel_TeletextDestinationSettingsPropertyOutputReference
	// Experimental.
	TeletextDestinationSettingsInput() *TfChannel_TeletextDestinationSettingsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TtmlDestinationSettings() TfChannel_TtmlDestinationSettingsPropertyOutputReference
	// Experimental.
	TtmlDestinationSettingsInput() *TfChannel_TtmlDestinationSettingsProperty
	// Experimental.
	WebvttDestinationSettings() TfChannel_WebvttDestinationSettingsPropertyOutputReference
	// Experimental.
	WebvttDestinationSettingsInput() *TfChannel_WebvttDestinationSettingsProperty
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
	PutAribDestinationSettings(value *TfChannel_AribDestinationSettingsProperty)
	// Experimental.
	PutBurnInDestinationSettings(value *TfChannel_BurnInDestinationSettingsProperty)
	// Experimental.
	PutDvbSubDestinationSettings(value *TfChannel_DvbSubDestinationSettingsProperty)
	// Experimental.
	PutEbuTtDDestinationSettings(value *TfChannel_EbuTtDDestinationSettingsProperty)
	// Experimental.
	PutEmbeddedDestinationSettings(value *TfChannel_EmbeddedDestinationSettingsProperty)
	// Experimental.
	PutEmbeddedPlusScte20DestinationSettings(value *TfChannel_EmbeddedPlusScte20DestinationSettingsProperty)
	// Experimental.
	PutRtmpCaptionInfoDestinationSettings(value *TfChannel_RtmpCaptionInfoDestinationSettingsProperty)
	// Experimental.
	PutScte20PlusEmbeddedDestinationSettings(value *TfChannel_Scte20PlusEmbeddedDestinationSettingsProperty)
	// Experimental.
	PutScte27DestinationSettings(value *TfChannel_Scte27DestinationSettingsProperty)
	// Experimental.
	PutSmpteTtDestinationSettings(value *TfChannel_SmpteTtDestinationSettingsProperty)
	// Experimental.
	PutTeletextDestinationSettings(value *TfChannel_TeletextDestinationSettingsProperty)
	// Experimental.
	PutTtmlDestinationSettings(value *TfChannel_TtmlDestinationSettingsProperty)
	// Experimental.
	PutWebvttDestinationSettings(value *TfChannel_WebvttDestinationSettingsProperty)
	// Experimental.
	ResetAribDestinationSettings()
	// Experimental.
	ResetBurnInDestinationSettings()
	// Experimental.
	ResetDvbSubDestinationSettings()
	// Experimental.
	ResetEbuTtDDestinationSettings()
	// Experimental.
	ResetEmbeddedDestinationSettings()
	// Experimental.
	ResetEmbeddedPlusScte20DestinationSettings()
	// Experimental.
	ResetRtmpCaptionInfoDestinationSettings()
	// Experimental.
	ResetScte20PlusEmbeddedDestinationSettings()
	// Experimental.
	ResetScte27DestinationSettings()
	// Experimental.
	ResetSmpteTtDestinationSettings()
	// Experimental.
	ResetTeletextDestinationSettings()
	// Experimental.
	ResetTtmlDestinationSettings()
	// Experimental.
	ResetWebvttDestinationSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_DestinationSettingsPropertyOutputReference
type jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) AribDestinationSettings() TfChannel_AribDestinationSettingsPropertyOutputReference {
	var returns TfChannel_AribDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"aribDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) AribDestinationSettingsInput() *TfChannel_AribDestinationSettingsProperty {
	var returns *TfChannel_AribDestinationSettingsProperty
	_jsii_.Get(
		j,
		"aribDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) BurnInDestinationSettings() TfChannel_BurnInDestinationSettingsPropertyOutputReference {
	var returns TfChannel_BurnInDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"burnInDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) BurnInDestinationSettingsInput() *TfChannel_BurnInDestinationSettingsProperty {
	var returns *TfChannel_BurnInDestinationSettingsProperty
	_jsii_.Get(
		j,
		"burnInDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) DvbSubDestinationSettings() TfChannel_DvbSubDestinationSettingsPropertyOutputReference {
	var returns TfChannel_DvbSubDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"dvbSubDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) DvbSubDestinationSettingsInput() *TfChannel_DvbSubDestinationSettingsProperty {
	var returns *TfChannel_DvbSubDestinationSettingsProperty
	_jsii_.Get(
		j,
		"dvbSubDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) EbuTtDDestinationSettings() TfChannel_EbuTtDDestinationSettingsPropertyOutputReference {
	var returns TfChannel_EbuTtDDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"ebuTtDDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) EbuTtDDestinationSettingsInput() *TfChannel_EbuTtDDestinationSettingsProperty {
	var returns *TfChannel_EbuTtDDestinationSettingsProperty
	_jsii_.Get(
		j,
		"ebuTtDDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) EmbeddedDestinationSettings() TfChannel_EmbeddedDestinationSettingsPropertyOutputReference {
	var returns TfChannel_EmbeddedDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"embeddedDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) EmbeddedDestinationSettingsInput() *TfChannel_EmbeddedDestinationSettingsProperty {
	var returns *TfChannel_EmbeddedDestinationSettingsProperty
	_jsii_.Get(
		j,
		"embeddedDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) EmbeddedPlusScte20DestinationSettings() TfChannel_EmbeddedPlusScte20DestinationSettingsPropertyOutputReference {
	var returns TfChannel_EmbeddedPlusScte20DestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"embeddedPlusScte20DestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) EmbeddedPlusScte20DestinationSettingsInput() *TfChannel_EmbeddedPlusScte20DestinationSettingsProperty {
	var returns *TfChannel_EmbeddedPlusScte20DestinationSettingsProperty
	_jsii_.Get(
		j,
		"embeddedPlusScte20DestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) InternalValue() *TfChannel_DestinationSettingsProperty {
	var returns *TfChannel_DestinationSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) RtmpCaptionInfoDestinationSettings() TfChannel_RtmpCaptionInfoDestinationSettingsPropertyOutputReference {
	var returns TfChannel_RtmpCaptionInfoDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rtmpCaptionInfoDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) RtmpCaptionInfoDestinationSettingsInput() *TfChannel_RtmpCaptionInfoDestinationSettingsProperty {
	var returns *TfChannel_RtmpCaptionInfoDestinationSettingsProperty
	_jsii_.Get(
		j,
		"rtmpCaptionInfoDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) Scte20PlusEmbeddedDestinationSettings() TfChannel_Scte20PlusEmbeddedDestinationSettingsPropertyOutputReference {
	var returns TfChannel_Scte20PlusEmbeddedDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"scte20PlusEmbeddedDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) Scte20PlusEmbeddedDestinationSettingsInput() *TfChannel_Scte20PlusEmbeddedDestinationSettingsProperty {
	var returns *TfChannel_Scte20PlusEmbeddedDestinationSettingsProperty
	_jsii_.Get(
		j,
		"scte20PlusEmbeddedDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) Scte27DestinationSettings() TfChannel_Scte27DestinationSettingsPropertyOutputReference {
	var returns TfChannel_Scte27DestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"scte27DestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) Scte27DestinationSettingsInput() *TfChannel_Scte27DestinationSettingsProperty {
	var returns *TfChannel_Scte27DestinationSettingsProperty
	_jsii_.Get(
		j,
		"scte27DestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) SmpteTtDestinationSettings() TfChannel_SmpteTtDestinationSettingsPropertyOutputReference {
	var returns TfChannel_SmpteTtDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"smpteTtDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) SmpteTtDestinationSettingsInput() *TfChannel_SmpteTtDestinationSettingsProperty {
	var returns *TfChannel_SmpteTtDestinationSettingsProperty
	_jsii_.Get(
		j,
		"smpteTtDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) TeletextDestinationSettings() TfChannel_TeletextDestinationSettingsPropertyOutputReference {
	var returns TfChannel_TeletextDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"teletextDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) TeletextDestinationSettingsInput() *TfChannel_TeletextDestinationSettingsProperty {
	var returns *TfChannel_TeletextDestinationSettingsProperty
	_jsii_.Get(
		j,
		"teletextDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) TtmlDestinationSettings() TfChannel_TtmlDestinationSettingsPropertyOutputReference {
	var returns TfChannel_TtmlDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"ttmlDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) TtmlDestinationSettingsInput() *TfChannel_TtmlDestinationSettingsProperty {
	var returns *TfChannel_TtmlDestinationSettingsProperty
	_jsii_.Get(
		j,
		"ttmlDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) WebvttDestinationSettings() TfChannel_WebvttDestinationSettingsPropertyOutputReference {
	var returns TfChannel_WebvttDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"webvttDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) WebvttDestinationSettingsInput() *TfChannel_WebvttDestinationSettingsProperty {
	var returns *TfChannel_WebvttDestinationSettingsProperty
	_jsii_.Get(
		j,
		"webvttDestinationSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_DestinationSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_DestinationSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_DestinationSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.DestinationSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_DestinationSettingsPropertyOutputReference_Override(t TfChannel_DestinationSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.DestinationSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_DestinationSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) PutAribDestinationSettings(value *TfChannel_AribDestinationSettingsProperty) {
	if err := t.validatePutAribDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAribDestinationSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) PutBurnInDestinationSettings(value *TfChannel_BurnInDestinationSettingsProperty) {
	if err := t.validatePutBurnInDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBurnInDestinationSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) PutDvbSubDestinationSettings(value *TfChannel_DvbSubDestinationSettingsProperty) {
	if err := t.validatePutDvbSubDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDvbSubDestinationSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) PutEbuTtDDestinationSettings(value *TfChannel_EbuTtDDestinationSettingsProperty) {
	if err := t.validatePutEbuTtDDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEbuTtDDestinationSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) PutEmbeddedDestinationSettings(value *TfChannel_EmbeddedDestinationSettingsProperty) {
	if err := t.validatePutEmbeddedDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEmbeddedDestinationSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) PutEmbeddedPlusScte20DestinationSettings(value *TfChannel_EmbeddedPlusScte20DestinationSettingsProperty) {
	if err := t.validatePutEmbeddedPlusScte20DestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEmbeddedPlusScte20DestinationSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) PutRtmpCaptionInfoDestinationSettings(value *TfChannel_RtmpCaptionInfoDestinationSettingsProperty) {
	if err := t.validatePutRtmpCaptionInfoDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRtmpCaptionInfoDestinationSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) PutScte20PlusEmbeddedDestinationSettings(value *TfChannel_Scte20PlusEmbeddedDestinationSettingsProperty) {
	if err := t.validatePutScte20PlusEmbeddedDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putScte20PlusEmbeddedDestinationSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) PutScte27DestinationSettings(value *TfChannel_Scte27DestinationSettingsProperty) {
	if err := t.validatePutScte27DestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putScte27DestinationSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) PutSmpteTtDestinationSettings(value *TfChannel_SmpteTtDestinationSettingsProperty) {
	if err := t.validatePutSmpteTtDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSmpteTtDestinationSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) PutTeletextDestinationSettings(value *TfChannel_TeletextDestinationSettingsProperty) {
	if err := t.validatePutTeletextDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTeletextDestinationSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) PutTtmlDestinationSettings(value *TfChannel_TtmlDestinationSettingsProperty) {
	if err := t.validatePutTtmlDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTtmlDestinationSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) PutWebvttDestinationSettings(value *TfChannel_WebvttDestinationSettingsProperty) {
	if err := t.validatePutWebvttDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWebvttDestinationSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) ResetAribDestinationSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetAribDestinationSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) ResetBurnInDestinationSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetBurnInDestinationSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) ResetDvbSubDestinationSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetDvbSubDestinationSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) ResetEbuTtDDestinationSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetEbuTtDDestinationSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) ResetEmbeddedDestinationSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetEmbeddedDestinationSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) ResetEmbeddedPlusScte20DestinationSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetEmbeddedPlusScte20DestinationSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) ResetRtmpCaptionInfoDestinationSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetRtmpCaptionInfoDestinationSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) ResetScte20PlusEmbeddedDestinationSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetScte20PlusEmbeddedDestinationSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) ResetScte27DestinationSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetScte27DestinationSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) ResetSmpteTtDestinationSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetSmpteTtDestinationSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) ResetTeletextDestinationSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetTeletextDestinationSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) ResetTtmlDestinationSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetTtmlDestinationSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) ResetWebvttDestinationSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetWebvttDestinationSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_DestinationSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

