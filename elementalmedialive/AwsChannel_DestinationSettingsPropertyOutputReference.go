package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_DestinationSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AribDestinationSettings() AwsChannel_AribDestinationSettingsPropertyOutputReference
	// Experimental.
	AribDestinationSettingsInput() *AwsChannel_AribDestinationSettingsProperty
	// Experimental.
	BurnInDestinationSettings() AwsChannel_BurnInDestinationSettingsPropertyOutputReference
	// Experimental.
	BurnInDestinationSettingsInput() *AwsChannel_BurnInDestinationSettingsProperty
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
	DvbSubDestinationSettings() AwsChannel_DvbSubDestinationSettingsPropertyOutputReference
	// Experimental.
	DvbSubDestinationSettingsInput() *AwsChannel_DvbSubDestinationSettingsProperty
	// Experimental.
	EbuTtDDestinationSettings() AwsChannel_EbuTtDDestinationSettingsPropertyOutputReference
	// Experimental.
	EbuTtDDestinationSettingsInput() *AwsChannel_EbuTtDDestinationSettingsProperty
	// Experimental.
	EmbeddedDestinationSettings() AwsChannel_EmbeddedDestinationSettingsPropertyOutputReference
	// Experimental.
	EmbeddedDestinationSettingsInput() *AwsChannel_EmbeddedDestinationSettingsProperty
	// Experimental.
	EmbeddedPlusScte20DestinationSettings() AwsChannel_EmbeddedPlusScte20DestinationSettingsPropertyOutputReference
	// Experimental.
	EmbeddedPlusScte20DestinationSettingsInput() *AwsChannel_EmbeddedPlusScte20DestinationSettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsChannel_DestinationSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsChannel_DestinationSettingsProperty)
	// Experimental.
	RtmpCaptionInfoDestinationSettings() AwsChannel_RtmpCaptionInfoDestinationSettingsPropertyOutputReference
	// Experimental.
	RtmpCaptionInfoDestinationSettingsInput() *AwsChannel_RtmpCaptionInfoDestinationSettingsProperty
	// Experimental.
	Scte20PlusEmbeddedDestinationSettings() AwsChannel_Scte20PlusEmbeddedDestinationSettingsPropertyOutputReference
	// Experimental.
	Scte20PlusEmbeddedDestinationSettingsInput() *AwsChannel_Scte20PlusEmbeddedDestinationSettingsProperty
	// Experimental.
	Scte27DestinationSettings() AwsChannel_Scte27DestinationSettingsPropertyOutputReference
	// Experimental.
	Scte27DestinationSettingsInput() *AwsChannel_Scte27DestinationSettingsProperty
	// Experimental.
	SmpteTtDestinationSettings() AwsChannel_SmpteTtDestinationSettingsPropertyOutputReference
	// Experimental.
	SmpteTtDestinationSettingsInput() *AwsChannel_SmpteTtDestinationSettingsProperty
	// Experimental.
	TeletextDestinationSettings() AwsChannel_TeletextDestinationSettingsPropertyOutputReference
	// Experimental.
	TeletextDestinationSettingsInput() *AwsChannel_TeletextDestinationSettingsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TtmlDestinationSettings() AwsChannel_TtmlDestinationSettingsPropertyOutputReference
	// Experimental.
	TtmlDestinationSettingsInput() *AwsChannel_TtmlDestinationSettingsProperty
	// Experimental.
	WebvttDestinationSettings() AwsChannel_WebvttDestinationSettingsPropertyOutputReference
	// Experimental.
	WebvttDestinationSettingsInput() *AwsChannel_WebvttDestinationSettingsProperty
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
	PutAribDestinationSettings(value *AwsChannel_AribDestinationSettingsProperty)
	// Experimental.
	PutBurnInDestinationSettings(value *AwsChannel_BurnInDestinationSettingsProperty)
	// Experimental.
	PutDvbSubDestinationSettings(value *AwsChannel_DvbSubDestinationSettingsProperty)
	// Experimental.
	PutEbuTtDDestinationSettings(value *AwsChannel_EbuTtDDestinationSettingsProperty)
	// Experimental.
	PutEmbeddedDestinationSettings(value *AwsChannel_EmbeddedDestinationSettingsProperty)
	// Experimental.
	PutEmbeddedPlusScte20DestinationSettings(value *AwsChannel_EmbeddedPlusScte20DestinationSettingsProperty)
	// Experimental.
	PutRtmpCaptionInfoDestinationSettings(value *AwsChannel_RtmpCaptionInfoDestinationSettingsProperty)
	// Experimental.
	PutScte20PlusEmbeddedDestinationSettings(value *AwsChannel_Scte20PlusEmbeddedDestinationSettingsProperty)
	// Experimental.
	PutScte27DestinationSettings(value *AwsChannel_Scte27DestinationSettingsProperty)
	// Experimental.
	PutSmpteTtDestinationSettings(value *AwsChannel_SmpteTtDestinationSettingsProperty)
	// Experimental.
	PutTeletextDestinationSettings(value *AwsChannel_TeletextDestinationSettingsProperty)
	// Experimental.
	PutTtmlDestinationSettings(value *AwsChannel_TtmlDestinationSettingsProperty)
	// Experimental.
	PutWebvttDestinationSettings(value *AwsChannel_WebvttDestinationSettingsProperty)
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

// The jsii proxy struct for AwsChannel_DestinationSettingsPropertyOutputReference
type jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) AribDestinationSettings() AwsChannel_AribDestinationSettingsPropertyOutputReference {
	var returns AwsChannel_AribDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"aribDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) AribDestinationSettingsInput() *AwsChannel_AribDestinationSettingsProperty {
	var returns *AwsChannel_AribDestinationSettingsProperty
	_jsii_.Get(
		j,
		"aribDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) BurnInDestinationSettings() AwsChannel_BurnInDestinationSettingsPropertyOutputReference {
	var returns AwsChannel_BurnInDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"burnInDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) BurnInDestinationSettingsInput() *AwsChannel_BurnInDestinationSettingsProperty {
	var returns *AwsChannel_BurnInDestinationSettingsProperty
	_jsii_.Get(
		j,
		"burnInDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) DvbSubDestinationSettings() AwsChannel_DvbSubDestinationSettingsPropertyOutputReference {
	var returns AwsChannel_DvbSubDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"dvbSubDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) DvbSubDestinationSettingsInput() *AwsChannel_DvbSubDestinationSettingsProperty {
	var returns *AwsChannel_DvbSubDestinationSettingsProperty
	_jsii_.Get(
		j,
		"dvbSubDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) EbuTtDDestinationSettings() AwsChannel_EbuTtDDestinationSettingsPropertyOutputReference {
	var returns AwsChannel_EbuTtDDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"ebuTtDDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) EbuTtDDestinationSettingsInput() *AwsChannel_EbuTtDDestinationSettingsProperty {
	var returns *AwsChannel_EbuTtDDestinationSettingsProperty
	_jsii_.Get(
		j,
		"ebuTtDDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) EmbeddedDestinationSettings() AwsChannel_EmbeddedDestinationSettingsPropertyOutputReference {
	var returns AwsChannel_EmbeddedDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"embeddedDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) EmbeddedDestinationSettingsInput() *AwsChannel_EmbeddedDestinationSettingsProperty {
	var returns *AwsChannel_EmbeddedDestinationSettingsProperty
	_jsii_.Get(
		j,
		"embeddedDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) EmbeddedPlusScte20DestinationSettings() AwsChannel_EmbeddedPlusScte20DestinationSettingsPropertyOutputReference {
	var returns AwsChannel_EmbeddedPlusScte20DestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"embeddedPlusScte20DestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) EmbeddedPlusScte20DestinationSettingsInput() *AwsChannel_EmbeddedPlusScte20DestinationSettingsProperty {
	var returns *AwsChannel_EmbeddedPlusScte20DestinationSettingsProperty
	_jsii_.Get(
		j,
		"embeddedPlusScte20DestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) InternalValue() *AwsChannel_DestinationSettingsProperty {
	var returns *AwsChannel_DestinationSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) RtmpCaptionInfoDestinationSettings() AwsChannel_RtmpCaptionInfoDestinationSettingsPropertyOutputReference {
	var returns AwsChannel_RtmpCaptionInfoDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rtmpCaptionInfoDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) RtmpCaptionInfoDestinationSettingsInput() *AwsChannel_RtmpCaptionInfoDestinationSettingsProperty {
	var returns *AwsChannel_RtmpCaptionInfoDestinationSettingsProperty
	_jsii_.Get(
		j,
		"rtmpCaptionInfoDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) Scte20PlusEmbeddedDestinationSettings() AwsChannel_Scte20PlusEmbeddedDestinationSettingsPropertyOutputReference {
	var returns AwsChannel_Scte20PlusEmbeddedDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"scte20PlusEmbeddedDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) Scte20PlusEmbeddedDestinationSettingsInput() *AwsChannel_Scte20PlusEmbeddedDestinationSettingsProperty {
	var returns *AwsChannel_Scte20PlusEmbeddedDestinationSettingsProperty
	_jsii_.Get(
		j,
		"scte20PlusEmbeddedDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) Scte27DestinationSettings() AwsChannel_Scte27DestinationSettingsPropertyOutputReference {
	var returns AwsChannel_Scte27DestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"scte27DestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) Scte27DestinationSettingsInput() *AwsChannel_Scte27DestinationSettingsProperty {
	var returns *AwsChannel_Scte27DestinationSettingsProperty
	_jsii_.Get(
		j,
		"scte27DestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) SmpteTtDestinationSettings() AwsChannel_SmpteTtDestinationSettingsPropertyOutputReference {
	var returns AwsChannel_SmpteTtDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"smpteTtDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) SmpteTtDestinationSettingsInput() *AwsChannel_SmpteTtDestinationSettingsProperty {
	var returns *AwsChannel_SmpteTtDestinationSettingsProperty
	_jsii_.Get(
		j,
		"smpteTtDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) TeletextDestinationSettings() AwsChannel_TeletextDestinationSettingsPropertyOutputReference {
	var returns AwsChannel_TeletextDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"teletextDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) TeletextDestinationSettingsInput() *AwsChannel_TeletextDestinationSettingsProperty {
	var returns *AwsChannel_TeletextDestinationSettingsProperty
	_jsii_.Get(
		j,
		"teletextDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) TtmlDestinationSettings() AwsChannel_TtmlDestinationSettingsPropertyOutputReference {
	var returns AwsChannel_TtmlDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"ttmlDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) TtmlDestinationSettingsInput() *AwsChannel_TtmlDestinationSettingsProperty {
	var returns *AwsChannel_TtmlDestinationSettingsProperty
	_jsii_.Get(
		j,
		"ttmlDestinationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) WebvttDestinationSettings() AwsChannel_WebvttDestinationSettingsPropertyOutputReference {
	var returns AwsChannel_WebvttDestinationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"webvttDestinationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) WebvttDestinationSettingsInput() *AwsChannel_WebvttDestinationSettingsProperty {
	var returns *AwsChannel_WebvttDestinationSettingsProperty
	_jsii_.Get(
		j,
		"webvttDestinationSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_DestinationSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChannel_DestinationSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_DestinationSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.DestinationSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_DestinationSettingsPropertyOutputReference_Override(a AwsChannel_DestinationSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.DestinationSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference)SetInternalValue(val *AwsChannel_DestinationSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) PutAribDestinationSettings(value *AwsChannel_AribDestinationSettingsProperty) {
	if err := a.validatePutAribDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAribDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) PutBurnInDestinationSettings(value *AwsChannel_BurnInDestinationSettingsProperty) {
	if err := a.validatePutBurnInDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBurnInDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) PutDvbSubDestinationSettings(value *AwsChannel_DvbSubDestinationSettingsProperty) {
	if err := a.validatePutDvbSubDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDvbSubDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) PutEbuTtDDestinationSettings(value *AwsChannel_EbuTtDDestinationSettingsProperty) {
	if err := a.validatePutEbuTtDDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEbuTtDDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) PutEmbeddedDestinationSettings(value *AwsChannel_EmbeddedDestinationSettingsProperty) {
	if err := a.validatePutEmbeddedDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEmbeddedDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) PutEmbeddedPlusScte20DestinationSettings(value *AwsChannel_EmbeddedPlusScte20DestinationSettingsProperty) {
	if err := a.validatePutEmbeddedPlusScte20DestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEmbeddedPlusScte20DestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) PutRtmpCaptionInfoDestinationSettings(value *AwsChannel_RtmpCaptionInfoDestinationSettingsProperty) {
	if err := a.validatePutRtmpCaptionInfoDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRtmpCaptionInfoDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) PutScte20PlusEmbeddedDestinationSettings(value *AwsChannel_Scte20PlusEmbeddedDestinationSettingsProperty) {
	if err := a.validatePutScte20PlusEmbeddedDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScte20PlusEmbeddedDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) PutScte27DestinationSettings(value *AwsChannel_Scte27DestinationSettingsProperty) {
	if err := a.validatePutScte27DestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScte27DestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) PutSmpteTtDestinationSettings(value *AwsChannel_SmpteTtDestinationSettingsProperty) {
	if err := a.validatePutSmpteTtDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSmpteTtDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) PutTeletextDestinationSettings(value *AwsChannel_TeletextDestinationSettingsProperty) {
	if err := a.validatePutTeletextDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTeletextDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) PutTtmlDestinationSettings(value *AwsChannel_TtmlDestinationSettingsProperty) {
	if err := a.validatePutTtmlDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTtmlDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) PutWebvttDestinationSettings(value *AwsChannel_WebvttDestinationSettingsProperty) {
	if err := a.validatePutWebvttDestinationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWebvttDestinationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) ResetAribDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetAribDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) ResetBurnInDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetBurnInDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) ResetDvbSubDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetDvbSubDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) ResetEbuTtDDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetEbuTtDDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) ResetEmbeddedDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetEmbeddedDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) ResetEmbeddedPlusScte20DestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetEmbeddedPlusScte20DestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) ResetRtmpCaptionInfoDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRtmpCaptionInfoDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) ResetScte20PlusEmbeddedDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetScte20PlusEmbeddedDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) ResetScte27DestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetScte27DestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) ResetSmpteTtDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetSmpteTtDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) ResetTeletextDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetTeletextDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) ResetTtmlDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetTtmlDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) ResetWebvttDestinationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetWebvttDestinationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_DestinationSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

