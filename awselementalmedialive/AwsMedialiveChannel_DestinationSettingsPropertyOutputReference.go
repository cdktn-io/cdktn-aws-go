package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMedialiveChannel_DestinationSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AribDestinationSettings() AwsMedialiveChannel_AribDestinationSettingsPropertyOutputReference
	// Experimental.
	AribDestinationSettingsInput() *AwsMedialiveChannel_AribDestinationSettingsProperty
	// Experimental.
	BurnInDestinationSettings() AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference
	// Experimental.
	BurnInDestinationSettingsInput() *AwsMedialiveChannel_BurnInDestinationSettingsProperty
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
	DvbSubDestinationSettings() AwsMedialiveChannel_DvbSubDestinationSettingsPropertyOutputReference
	// Experimental.
	DvbSubDestinationSettingsInput() *AwsMedialiveChannel_DvbSubDestinationSettingsProperty
	// Experimental.
	EbuTtDDestinationSettings() AwsMedialiveChannel_EbuTtDDestinationSettingsPropertyOutputReference
	// Experimental.
	EbuTtDDestinationSettingsInput() *AwsMedialiveChannel_EbuTtDDestinationSettingsProperty
	// Experimental.
	EmbeddedDestinationSettings() AwsMedialiveChannel_EmbeddedDestinationSettingsPropertyOutputReference
	// Experimental.
	EmbeddedDestinationSettingsInput() *AwsMedialiveChannel_EmbeddedDestinationSettingsProperty
	// Experimental.
	EmbeddedPlusScte20DestinationSettings() AwsMedialiveChannel_EmbeddedPlusScte20DestinationSettingsPropertyOutputReference
	// Experimental.
	EmbeddedPlusScte20DestinationSettingsInput() *AwsMedialiveChannel_EmbeddedPlusScte20DestinationSettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsMedialiveChannel_DestinationSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsMedialiveChannel_DestinationSettingsProperty)
	// Experimental.
	RtmpCaptionInfoDestinationSettings() AwsMedialiveChannel_RtmpCaptionInfoDestinationSettingsPropertyOutputReference
	// Experimental.
	RtmpCaptionInfoDestinationSettingsInput() *AwsMedialiveChannel_RtmpCaptionInfoDestinationSettingsProperty
	// Experimental.
	Scte20PlusEmbeddedDestinationSettings() AwsMedialiveChannel_Scte20PlusEmbeddedDestinationSettingsPropertyOutputReference
	// Experimental.
	Scte20PlusEmbeddedDestinationSettingsInput() *AwsMedialiveChannel_Scte20PlusEmbeddedDestinationSettingsProperty
	// Experimental.
	Scte27DestinationSettings() AwsMedialiveChannel_Scte27DestinationSettingsPropertyOutputReference
	// Experimental.
	Scte27DestinationSettingsInput() *AwsMedialiveChannel_Scte27DestinationSettingsProperty
	// Experimental.
	SmpteTtDestinationSettings() AwsMedialiveChannel_SmpteTtDestinationSettingsPropertyOutputReference
	// Experimental.
	SmpteTtDestinationSettingsInput() *AwsMedialiveChannel_SmpteTtDestinationSettingsProperty
	// Experimental.
	TeletextDestinationSettings() AwsMedialiveChannel_TeletextDestinationSettingsPropertyOutputReference
	// Experimental.
	TeletextDestinationSettingsInput() *AwsMedialiveChannel_TeletextDestinationSettingsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TtmlDestinationSettings() AwsMedialiveChannel_TtmlDestinationSettingsPropertyOutputReference
	// Experimental.
	TtmlDestinationSettingsInput() *AwsMedialiveChannel_TtmlDestinationSettingsProperty
	// Experimental.
	WebvttDestinationSettings() AwsMedialiveChannel_WebvttDestinationSettingsPropertyOutputReference
	// Experimental.
	WebvttDestinationSettingsInput() *AwsMedialiveChannel_WebvttDestinationSettingsProperty
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
	PutAribDestinationSettings(value *AwsMedialiveChannel_AribDestinationSettingsProperty)
	// Experimental.
	PutBurnInDestinationSettings(value *AwsMedialiveChannel_BurnInDestinationSettingsProperty)
	// Experimental.
	PutDvbSubDestinationSettings(value *AwsMedialiveChannel_DvbSubDestinationSettingsProperty)
	// Experimental.
	PutEbuTtDDestinationSettings(value *AwsMedialiveChannel_EbuTtDDestinationSettingsProperty)
	// Experimental.
	PutEmbeddedDestinationSettings(value *AwsMedialiveChannel_EmbeddedDestinationSettingsProperty)
	// Experimental.
	PutEmbeddedPlusScte20DestinationSettings(value *AwsMedialiveChannel_EmbeddedPlusScte20DestinationSettingsProperty)
	// Experimental.
	PutRtmpCaptionInfoDestinationSettings(value *AwsMedialiveChannel_RtmpCaptionInfoDestinationSettingsProperty)
	// Experimental.
	PutScte20PlusEmbeddedDestinationSettings(value *AwsMedialiveChannel_Scte20PlusEmbeddedDestinationSettingsProperty)
	// Experimental.
	PutScte27DestinationSettings(value *AwsMedialiveChannel_Scte27DestinationSettingsProperty)
	// Experimental.
	PutSmpteTtDestinationSettings(value *AwsMedialiveChannel_SmpteTtDestinationSettingsProperty)
	// Experimental.
	PutTeletextDestinationSettings(value *AwsMedialiveChannel_TeletextDestinationSettingsProperty)
	// Experimental.
	PutTtmlDestinationSettings(value *AwsMedialiveChannel_TtmlDestinationSettingsProperty)
	// Experimental.
	PutWebvttDestinationSettings(value *AwsMedialiveChannel_WebvttDestinationSettingsProperty)
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

// The jsii proxy struct for AwsMedialiveChannel_DestinationSettingsPropertyOutputReference
type jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) AribDestinationSettings() AwsMedialiveChannel_AribDestinationSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_AribDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"aribDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) AribDestinationSettingsInput() *AwsMedialiveChannel_AribDestinationSettingsProperty {
	var returns *AwsMedialiveChannel_AribDestinationSettingsProperty
	_jsii_.Get(
		j,
		"aribDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) BurnInDestinationSettings() AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_BurnInDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"burnInDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) BurnInDestinationSettingsInput() *AwsMedialiveChannel_BurnInDestinationSettingsProperty {
	var returns *AwsMedialiveChannel_BurnInDestinationSettingsProperty
	_jsii_.Get(
		j,
		"burnInDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) DvbSubDestinationSettings() AwsMedialiveChannel_DvbSubDestinationSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_DvbSubDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"dvbSubDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) DvbSubDestinationSettingsInput() *AwsMedialiveChannel_DvbSubDestinationSettingsProperty {
	var returns *AwsMedialiveChannel_DvbSubDestinationSettingsProperty
	_jsii_.Get(
		j,
		"dvbSubDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) EbuTtDDestinationSettings() AwsMedialiveChannel_EbuTtDDestinationSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_EbuTtDDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"ebuTtDDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) EbuTtDDestinationSettingsInput() *AwsMedialiveChannel_EbuTtDDestinationSettingsProperty {
	var returns *AwsMedialiveChannel_EbuTtDDestinationSettingsProperty
	_jsii_.Get(
		j,
		"ebuTtDDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) EmbeddedDestinationSettings() AwsMedialiveChannel_EmbeddedDestinationSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_EmbeddedDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"embeddedDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) EmbeddedDestinationSettingsInput() *AwsMedialiveChannel_EmbeddedDestinationSettingsProperty {
	var returns *AwsMedialiveChannel_EmbeddedDestinationSettingsProperty
	_jsii_.Get(
		j,
		"embeddedDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) EmbeddedPlusScte20DestinationSettings() AwsMedialiveChannel_EmbeddedPlusScte20DestinationSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_EmbeddedPlusScte20DestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"embeddedPlusScte20DestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) EmbeddedPlusScte20DestinationSettingsInput() *AwsMedialiveChannel_EmbeddedPlusScte20DestinationSettingsProperty {
	var returns *AwsMedialiveChannel_EmbeddedPlusScte20DestinationSettingsProperty
	_jsii_.Get(
		j,
		"embeddedPlusScte20DestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) InternalValue() *AwsMedialiveChannel_DestinationSettingsProperty {
	var returns *AwsMedialiveChannel_DestinationSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) RtmpCaptionInfoDestinationSettings() AwsMedialiveChannel_RtmpCaptionInfoDestinationSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_RtmpCaptionInfoDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rtmpCaptionInfoDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) RtmpCaptionInfoDestinationSettingsInput() *AwsMedialiveChannel_RtmpCaptionInfoDestinationSettingsProperty {
	var returns *AwsMedialiveChannel_RtmpCaptionInfoDestinationSettingsProperty
	_jsii_.Get(
		j,
		"rtmpCaptionInfoDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) Scte20PlusEmbeddedDestinationSettings() AwsMedialiveChannel_Scte20PlusEmbeddedDestinationSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_Scte20PlusEmbeddedDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"scte20PlusEmbeddedDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) Scte20PlusEmbeddedDestinationSettingsInput() *AwsMedialiveChannel_Scte20PlusEmbeddedDestinationSettingsProperty {
	var returns *AwsMedialiveChannel_Scte20PlusEmbeddedDestinationSettingsProperty
	_jsii_.Get(
		j,
		"scte20PlusEmbeddedDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) Scte27DestinationSettings() AwsMedialiveChannel_Scte27DestinationSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_Scte27DestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"scte27DestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) Scte27DestinationSettingsInput() *AwsMedialiveChannel_Scte27DestinationSettingsProperty {
	var returns *AwsMedialiveChannel_Scte27DestinationSettingsProperty
	_jsii_.Get(
		j,
		"scte27DestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) SmpteTtDestinationSettings() AwsMedialiveChannel_SmpteTtDestinationSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_SmpteTtDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"smpteTtDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) SmpteTtDestinationSettingsInput() *AwsMedialiveChannel_SmpteTtDestinationSettingsProperty {
	var returns *AwsMedialiveChannel_SmpteTtDestinationSettingsProperty
	_jsii_.Get(
		j,
		"smpteTtDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) TeletextDestinationSettings() AwsMedialiveChannel_TeletextDestinationSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_TeletextDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"teletextDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) TeletextDestinationSettingsInput() *AwsMedialiveChannel_TeletextDestinationSettingsProperty {
	var returns *AwsMedialiveChannel_TeletextDestinationSettingsProperty
	_jsii_.Get(
		j,
		"teletextDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) TtmlDestinationSettings() AwsMedialiveChannel_TtmlDestinationSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_TtmlDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"ttmlDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) TtmlDestinationSettingsInput() *AwsMedialiveChannel_TtmlDestinationSettingsProperty {
	var returns *AwsMedialiveChannel_TtmlDestinationSettingsProperty
	_jsii_.Get(
		j,
		"ttmlDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) WebvttDestinationSettings() AwsMedialiveChannel_WebvttDestinationSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_WebvttDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"webvttDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) WebvttDestinationSettingsInput() *AwsMedialiveChannel_WebvttDestinationSettingsProperty {
	var returns *AwsMedialiveChannel_WebvttDestinationSettingsProperty
	_jsii_.Get(
		j,
		"webvttDestinationSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMedialiveChannel_DestinationSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMedialiveChannel_DestinationSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMedialiveChannel_DestinationSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.DestinationSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMedialiveChannel_DestinationSettingsPropertyOutputReference_Override(a AwsMedialiveChannel_DestinationSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.DestinationSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference)SetInternalValue(val *AwsMedialiveChannel_DestinationSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) PutAribDestinationSettings(value *AwsMedialiveChannel_AribDestinationSettingsProperty) {
	if err := a.validatePutAribDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAribDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) PutBurnInDestinationSettings(value *AwsMedialiveChannel_BurnInDestinationSettingsProperty) {
	if err := a.validatePutBurnInDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBurnInDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) PutDvbSubDestinationSettings(value *AwsMedialiveChannel_DvbSubDestinationSettingsProperty) {
	if err := a.validatePutDvbSubDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDvbSubDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) PutEbuTtDDestinationSettings(value *AwsMedialiveChannel_EbuTtDDestinationSettingsProperty) {
	if err := a.validatePutEbuTtDDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEbuTtDDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) PutEmbeddedDestinationSettings(value *AwsMedialiveChannel_EmbeddedDestinationSettingsProperty) {
	if err := a.validatePutEmbeddedDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEmbeddedDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) PutEmbeddedPlusScte20DestinationSettings(value *AwsMedialiveChannel_EmbeddedPlusScte20DestinationSettingsProperty) {
	if err := a.validatePutEmbeddedPlusScte20DestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEmbeddedPlusScte20DestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) PutRtmpCaptionInfoDestinationSettings(value *AwsMedialiveChannel_RtmpCaptionInfoDestinationSettingsProperty) {
	if err := a.validatePutRtmpCaptionInfoDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRtmpCaptionInfoDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) PutScte20PlusEmbeddedDestinationSettings(value *AwsMedialiveChannel_Scte20PlusEmbeddedDestinationSettingsProperty) {
	if err := a.validatePutScte20PlusEmbeddedDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScte20PlusEmbeddedDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) PutScte27DestinationSettings(value *AwsMedialiveChannel_Scte27DestinationSettingsProperty) {
	if err := a.validatePutScte27DestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScte27DestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) PutSmpteTtDestinationSettings(value *AwsMedialiveChannel_SmpteTtDestinationSettingsProperty) {
	if err := a.validatePutSmpteTtDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSmpteTtDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) PutTeletextDestinationSettings(value *AwsMedialiveChannel_TeletextDestinationSettingsProperty) {
	if err := a.validatePutTeletextDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTeletextDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) PutTtmlDestinationSettings(value *AwsMedialiveChannel_TtmlDestinationSettingsProperty) {
	if err := a.validatePutTtmlDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTtmlDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) PutWebvttDestinationSettings(value *AwsMedialiveChannel_WebvttDestinationSettingsProperty) {
	if err := a.validatePutWebvttDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWebvttDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) ResetAribDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetAribDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) ResetBurnInDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetBurnInDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) ResetDvbSubDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetDvbSubDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) ResetEbuTtDDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetEbuTtDDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) ResetEmbeddedDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetEmbeddedDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) ResetEmbeddedPlusScte20DestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetEmbeddedPlusScte20DestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) ResetRtmpCaptionInfoDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRtmpCaptionInfoDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) ResetScte20PlusEmbeddedDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetScte20PlusEmbeddedDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) ResetScte27DestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetScte27DestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) ResetSmpteTtDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetSmpteTtDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) ResetTeletextDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetTeletextDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) ResetTtmlDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetTtmlDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) ResetWebvttDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetWebvttDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_DestinationSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

