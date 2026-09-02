package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_Eac3SettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AttenuationControl() *string
	// Experimental.
	SetAttenuationControl(val *string)
	// Experimental.
	AttenuationControlInput() *string
	// Experimental.
	Bitrate() *float64
	// Experimental.
	SetBitrate(val *float64)
	// Experimental.
	BitrateInput() *float64
	// Experimental.
	BitstreamMode() *string
	// Experimental.
	SetBitstreamMode(val *string)
	// Experimental.
	BitstreamModeInput() *string
	// Experimental.
	CodingMode() *string
	// Experimental.
	SetCodingMode(val *string)
	// Experimental.
	CodingModeInput() *string
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
	DcFilter() *string
	// Experimental.
	SetDcFilter(val *string)
	// Experimental.
	DcFilterInput() *string
	// Experimental.
	Dialnorm() *float64
	// Experimental.
	SetDialnorm(val *float64)
	// Experimental.
	DialnormInput() *float64
	// Experimental.
	DrcLine() *string
	// Experimental.
	SetDrcLine(val *string)
	// Experimental.
	DrcLineInput() *string
	// Experimental.
	DrcRf() *string
	// Experimental.
	SetDrcRf(val *string)
	// Experimental.
	DrcRfInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfChannel_Eac3SettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_Eac3SettingsProperty)
	// Experimental.
	LfeControl() *string
	// Experimental.
	SetLfeControl(val *string)
	// Experimental.
	LfeControlInput() *string
	// Experimental.
	LfeFilter() *string
	// Experimental.
	SetLfeFilter(val *string)
	// Experimental.
	LfeFilterInput() *string
	// Experimental.
	LoRoCenterMixLevel() *float64
	// Experimental.
	SetLoRoCenterMixLevel(val *float64)
	// Experimental.
	LoRoCenterMixLevelInput() *float64
	// Experimental.
	LoRoSurroundMixLevel() *float64
	// Experimental.
	SetLoRoSurroundMixLevel(val *float64)
	// Experimental.
	LoRoSurroundMixLevelInput() *float64
	// Experimental.
	LtRtCenterMixLevel() *float64
	// Experimental.
	SetLtRtCenterMixLevel(val *float64)
	// Experimental.
	LtRtCenterMixLevelInput() *float64
	// Experimental.
	LtRtSurroundMixLevel() *float64
	// Experimental.
	SetLtRtSurroundMixLevel(val *float64)
	// Experimental.
	LtRtSurroundMixLevelInput() *float64
	// Experimental.
	MetadataControl() *string
	// Experimental.
	SetMetadataControl(val *string)
	// Experimental.
	MetadataControlInput() *string
	// Experimental.
	PassthroughControl() *string
	// Experimental.
	SetPassthroughControl(val *string)
	// Experimental.
	PassthroughControlInput() *string
	// Experimental.
	PhaseControl() *string
	// Experimental.
	SetPhaseControl(val *string)
	// Experimental.
	PhaseControlInput() *string
	// Experimental.
	StereoDownmix() *string
	// Experimental.
	SetStereoDownmix(val *string)
	// Experimental.
	StereoDownmixInput() *string
	// Experimental.
	SurroundExMode() *string
	// Experimental.
	SetSurroundExMode(val *string)
	// Experimental.
	SurroundExModeInput() *string
	// Experimental.
	SurroundMode() *string
	// Experimental.
	SetSurroundMode(val *string)
	// Experimental.
	SurroundModeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetAttenuationControl()
	// Experimental.
	ResetBitrate()
	// Experimental.
	ResetBitstreamMode()
	// Experimental.
	ResetCodingMode()
	// Experimental.
	ResetDcFilter()
	// Experimental.
	ResetDialnorm()
	// Experimental.
	ResetDrcLine()
	// Experimental.
	ResetDrcRf()
	// Experimental.
	ResetLfeControl()
	// Experimental.
	ResetLfeFilter()
	// Experimental.
	ResetLoRoCenterMixLevel()
	// Experimental.
	ResetLoRoSurroundMixLevel()
	// Experimental.
	ResetLtRtCenterMixLevel()
	// Experimental.
	ResetLtRtSurroundMixLevel()
	// Experimental.
	ResetMetadataControl()
	// Experimental.
	ResetPassthroughControl()
	// Experimental.
	ResetPhaseControl()
	// Experimental.
	ResetStereoDownmix()
	// Experimental.
	ResetSurroundExMode()
	// Experimental.
	ResetSurroundMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_Eac3SettingsPropertyOutputReference
type jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) AttenuationControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attenuationControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) AttenuationControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attenuationControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) Bitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) BitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) BitstreamMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitstreamMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) BitstreamModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitstreamModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) CodingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) CodingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) DcFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dcFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) DcFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dcFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) Dialnorm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dialnorm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) DialnormInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dialnormInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) DrcLine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drcLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) DrcLineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drcLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) DrcRf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drcRf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) DrcRfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drcRfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) InternalValue() *TfChannel_Eac3SettingsProperty {
	var returns *TfChannel_Eac3SettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) LfeControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lfeControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) LfeControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lfeControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) LfeFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lfeFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) LfeFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lfeFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) LoRoCenterMixLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loRoCenterMixLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) LoRoCenterMixLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loRoCenterMixLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) LoRoSurroundMixLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loRoSurroundMixLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) LoRoSurroundMixLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loRoSurroundMixLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) LtRtCenterMixLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ltRtCenterMixLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) LtRtCenterMixLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ltRtCenterMixLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) LtRtSurroundMixLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ltRtSurroundMixLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) LtRtSurroundMixLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ltRtSurroundMixLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) MetadataControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) MetadataControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) PassthroughControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passthroughControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) PassthroughControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passthroughControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) PhaseControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phaseControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) PhaseControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phaseControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) StereoDownmix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stereoDownmix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) StereoDownmixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stereoDownmixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) SurroundExMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surroundExMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) SurroundExModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surroundExModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) SurroundMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surroundMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) SurroundModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surroundModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_Eac3SettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_Eac3SettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_Eac3SettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.Eac3SettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_Eac3SettingsPropertyOutputReference_Override(t TfChannel_Eac3SettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.Eac3SettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetAttenuationControl(val *string) {
	if err := j.validateSetAttenuationControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attenuationControl",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetBitrate(val *float64) {
	if err := j.validateSetBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitrate",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetBitstreamMode(val *string) {
	if err := j.validateSetBitstreamModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitstreamMode",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetCodingMode(val *string) {
	if err := j.validateSetCodingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codingMode",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetDcFilter(val *string) {
	if err := j.validateSetDcFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dcFilter",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetDialnorm(val *float64) {
	if err := j.validateSetDialnormParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dialnorm",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetDrcLine(val *string) {
	if err := j.validateSetDrcLineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drcLine",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetDrcRf(val *string) {
	if err := j.validateSetDrcRfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drcRf",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetInternalValue(val *TfChannel_Eac3SettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetLfeControl(val *string) {
	if err := j.validateSetLfeControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lfeControl",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetLfeFilter(val *string) {
	if err := j.validateSetLfeFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lfeFilter",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetLoRoCenterMixLevel(val *float64) {
	if err := j.validateSetLoRoCenterMixLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loRoCenterMixLevel",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetLoRoSurroundMixLevel(val *float64) {
	if err := j.validateSetLoRoSurroundMixLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loRoSurroundMixLevel",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetLtRtCenterMixLevel(val *float64) {
	if err := j.validateSetLtRtCenterMixLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ltRtCenterMixLevel",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetLtRtSurroundMixLevel(val *float64) {
	if err := j.validateSetLtRtSurroundMixLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ltRtSurroundMixLevel",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetMetadataControl(val *string) {
	if err := j.validateSetMetadataControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataControl",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetPassthroughControl(val *string) {
	if err := j.validateSetPassthroughControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passthroughControl",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetPhaseControl(val *string) {
	if err := j.validateSetPhaseControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phaseControl",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetStereoDownmix(val *string) {
	if err := j.validateSetStereoDownmixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stereoDownmix",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetSurroundExMode(val *string) {
	if err := j.validateSetSurroundExModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"surroundExMode",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetSurroundMode(val *string) {
	if err := j.validateSetSurroundModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"surroundMode",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetAttenuationControl() {
	_jsii_.InvokeVoid(
		t,
		"resetAttenuationControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetBitrate() {
	_jsii_.InvokeVoid(
		t,
		"resetBitrate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetBitstreamMode() {
	_jsii_.InvokeVoid(
		t,
		"resetBitstreamMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetCodingMode() {
	_jsii_.InvokeVoid(
		t,
		"resetCodingMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetDcFilter() {
	_jsii_.InvokeVoid(
		t,
		"resetDcFilter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetDialnorm() {
	_jsii_.InvokeVoid(
		t,
		"resetDialnorm",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetDrcLine() {
	_jsii_.InvokeVoid(
		t,
		"resetDrcLine",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetDrcRf() {
	_jsii_.InvokeVoid(
		t,
		"resetDrcRf",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetLfeControl() {
	_jsii_.InvokeVoid(
		t,
		"resetLfeControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetLfeFilter() {
	_jsii_.InvokeVoid(
		t,
		"resetLfeFilter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetLoRoCenterMixLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetLoRoCenterMixLevel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetLoRoSurroundMixLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetLoRoSurroundMixLevel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetLtRtCenterMixLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetLtRtCenterMixLevel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetLtRtSurroundMixLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetLtRtSurroundMixLevel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetMetadataControl() {
	_jsii_.InvokeVoid(
		t,
		"resetMetadataControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetPassthroughControl() {
	_jsii_.InvokeVoid(
		t,
		"resetPassthroughControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetPhaseControl() {
	_jsii_.InvokeVoid(
		t,
		"resetPhaseControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetStereoDownmix() {
	_jsii_.InvokeVoid(
		t,
		"resetStereoDownmix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetSurroundExMode() {
	_jsii_.InvokeVoid(
		t,
		"resetSurroundExMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ResetSurroundMode() {
	_jsii_.InvokeVoid(
		t,
		"resetSurroundMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_Eac3SettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

