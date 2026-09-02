package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_H264SettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdaptiveQuantization() *string
	// Experimental.
	SetAdaptiveQuantization(val *string)
	// Experimental.
	AdaptiveQuantizationInput() *string
	// Experimental.
	AfdSignaling() *string
	// Experimental.
	SetAfdSignaling(val *string)
	// Experimental.
	AfdSignalingInput() *string
	// Experimental.
	Bitrate() *float64
	// Experimental.
	SetBitrate(val *float64)
	// Experimental.
	BitrateInput() *float64
	// Experimental.
	BufFillPct() *float64
	// Experimental.
	SetBufFillPct(val *float64)
	// Experimental.
	BufFillPctInput() *float64
	// Experimental.
	BufSize() *float64
	// Experimental.
	SetBufSize(val *float64)
	// Experimental.
	BufSizeInput() *float64
	// Experimental.
	ColorMetadata() *string
	// Experimental.
	SetColorMetadata(val *string)
	// Experimental.
	ColorMetadataInput() *string
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
	EntropyEncoding() *string
	// Experimental.
	SetEntropyEncoding(val *string)
	// Experimental.
	EntropyEncodingInput() *string
	// Experimental.
	FilterSettings() TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsPropertyOutputReference
	// Experimental.
	FilterSettingsInput() *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsProperty
	// Experimental.
	FixedAfd() *string
	// Experimental.
	SetFixedAfd(val *string)
	// Experimental.
	FixedAfdInput() *string
	// Experimental.
	FlickerAq() *string
	// Experimental.
	SetFlickerAq(val *string)
	// Experimental.
	FlickerAqInput() *string
	// Experimental.
	ForceFieldPictures() *string
	// Experimental.
	SetForceFieldPictures(val *string)
	// Experimental.
	ForceFieldPicturesInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FramerateControl() *string
	// Experimental.
	SetFramerateControl(val *string)
	// Experimental.
	FramerateControlInput() *string
	// Experimental.
	FramerateDenominator() *float64
	// Experimental.
	SetFramerateDenominator(val *float64)
	// Experimental.
	FramerateDenominatorInput() *float64
	// Experimental.
	FramerateNumerator() *float64
	// Experimental.
	SetFramerateNumerator(val *float64)
	// Experimental.
	FramerateNumeratorInput() *float64
	// Experimental.
	GopBReference() *string
	// Experimental.
	SetGopBReference(val *string)
	// Experimental.
	GopBReferenceInput() *string
	// Experimental.
	GopClosedCadence() *float64
	// Experimental.
	SetGopClosedCadence(val *float64)
	// Experimental.
	GopClosedCadenceInput() *float64
	// Experimental.
	GopNumBFrames() *float64
	// Experimental.
	SetGopNumBFrames(val *float64)
	// Experimental.
	GopNumBFramesInput() *float64
	// Experimental.
	GopSize() *float64
	// Experimental.
	SetGopSize(val *float64)
	// Experimental.
	GopSizeInput() *float64
	// Experimental.
	GopSizeUnits() *string
	// Experimental.
	SetGopSizeUnits(val *string)
	// Experimental.
	GopSizeUnitsInput() *string
	// Experimental.
	InternalValue() *TfChannel_H264SettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_H264SettingsProperty)
	// Experimental.
	Level() *string
	// Experimental.
	SetLevel(val *string)
	// Experimental.
	LevelInput() *string
	// Experimental.
	LookAheadRateControl() *string
	// Experimental.
	SetLookAheadRateControl(val *string)
	// Experimental.
	LookAheadRateControlInput() *string
	// Experimental.
	MaxBitrate() *float64
	// Experimental.
	SetMaxBitrate(val *float64)
	// Experimental.
	MaxBitrateInput() *float64
	// Experimental.
	MinIInterval() *float64
	// Experimental.
	SetMinIInterval(val *float64)
	// Experimental.
	MinIIntervalInput() *float64
	// Experimental.
	NumRefFrames() *float64
	// Experimental.
	SetNumRefFrames(val *float64)
	// Experimental.
	NumRefFramesInput() *float64
	// Experimental.
	ParControl() *string
	// Experimental.
	SetParControl(val *string)
	// Experimental.
	ParControlInput() *string
	// Experimental.
	ParDenominator() *float64
	// Experimental.
	SetParDenominator(val *float64)
	// Experimental.
	ParDenominatorInput() *float64
	// Experimental.
	ParNumerator() *float64
	// Experimental.
	SetParNumerator(val *float64)
	// Experimental.
	ParNumeratorInput() *float64
	// Experimental.
	Profile() *string
	// Experimental.
	SetProfile(val *string)
	// Experimental.
	ProfileInput() *string
	// Experimental.
	QualityLevel() *string
	// Experimental.
	SetQualityLevel(val *string)
	// Experimental.
	QualityLevelInput() *string
	// Experimental.
	QvbrQualityLevel() *float64
	// Experimental.
	SetQvbrQualityLevel(val *float64)
	// Experimental.
	QvbrQualityLevelInput() *float64
	// Experimental.
	RateControlMode() *string
	// Experimental.
	SetRateControlMode(val *string)
	// Experimental.
	RateControlModeInput() *string
	// Experimental.
	ScanType() *string
	// Experimental.
	SetScanType(val *string)
	// Experimental.
	ScanTypeInput() *string
	// Experimental.
	SceneChangeDetect() *string
	// Experimental.
	SetSceneChangeDetect(val *string)
	// Experimental.
	SceneChangeDetectInput() *string
	// Experimental.
	Slices() *float64
	// Experimental.
	SetSlices(val *float64)
	// Experimental.
	SlicesInput() *float64
	// Experimental.
	Softness() *float64
	// Experimental.
	SetSoftness(val *float64)
	// Experimental.
	SoftnessInput() *float64
	// Experimental.
	SpatialAq() *string
	// Experimental.
	SetSpatialAq(val *string)
	// Experimental.
	SpatialAqInput() *string
	// Experimental.
	SubgopLength() *string
	// Experimental.
	SetSubgopLength(val *string)
	// Experimental.
	SubgopLengthInput() *string
	// Experimental.
	Syntax() *string
	// Experimental.
	SetSyntax(val *string)
	// Experimental.
	SyntaxInput() *string
	// Experimental.
	TemporalAq() *string
	// Experimental.
	SetTemporalAq(val *string)
	// Experimental.
	TemporalAqInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimecodeInsertion() *string
	// Experimental.
	SetTimecodeInsertion(val *string)
	// Experimental.
	TimecodeInsertionInput() *string
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
	PutFilterSettings(value *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsProperty)
	// Experimental.
	ResetAdaptiveQuantization()
	// Experimental.
	ResetAfdSignaling()
	// Experimental.
	ResetBitrate()
	// Experimental.
	ResetBufFillPct()
	// Experimental.
	ResetBufSize()
	// Experimental.
	ResetColorMetadata()
	// Experimental.
	ResetEntropyEncoding()
	// Experimental.
	ResetFilterSettings()
	// Experimental.
	ResetFixedAfd()
	// Experimental.
	ResetFlickerAq()
	// Experimental.
	ResetForceFieldPictures()
	// Experimental.
	ResetFramerateControl()
	// Experimental.
	ResetFramerateDenominator()
	// Experimental.
	ResetFramerateNumerator()
	// Experimental.
	ResetGopBReference()
	// Experimental.
	ResetGopClosedCadence()
	// Experimental.
	ResetGopNumBFrames()
	// Experimental.
	ResetGopSize()
	// Experimental.
	ResetGopSizeUnits()
	// Experimental.
	ResetLevel()
	// Experimental.
	ResetLookAheadRateControl()
	// Experimental.
	ResetMaxBitrate()
	// Experimental.
	ResetMinIInterval()
	// Experimental.
	ResetNumRefFrames()
	// Experimental.
	ResetParControl()
	// Experimental.
	ResetParDenominator()
	// Experimental.
	ResetParNumerator()
	// Experimental.
	ResetProfile()
	// Experimental.
	ResetQualityLevel()
	// Experimental.
	ResetQvbrQualityLevel()
	// Experimental.
	ResetRateControlMode()
	// Experimental.
	ResetScanType()
	// Experimental.
	ResetSceneChangeDetect()
	// Experimental.
	ResetSlices()
	// Experimental.
	ResetSoftness()
	// Experimental.
	ResetSpatialAq()
	// Experimental.
	ResetSubgopLength()
	// Experimental.
	ResetSyntax()
	// Experimental.
	ResetTemporalAq()
	// Experimental.
	ResetTimecodeInsertion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_H264SettingsPropertyOutputReference
type jsiiProxy_TfChannel_H264SettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) AdaptiveQuantization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adaptiveQuantization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) AdaptiveQuantizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adaptiveQuantizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) AfdSignaling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afdSignaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) AfdSignalingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afdSignalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) Bitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) BitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) BufFillPct() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufFillPct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) BufFillPctInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufFillPctInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) BufSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) BufSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ColorMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ColorMetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) EntropyEncoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entropyEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) EntropyEncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entropyEncodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) FilterSettings() TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsPropertyOutputReference {
	var returns TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"filterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) FilterSettingsInput() *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsProperty {
	var returns *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsProperty
	_jsii_.Get(
		j,
		"filterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) FixedAfd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedAfd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) FixedAfdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedAfdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) FlickerAq() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flickerAq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) FlickerAqInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flickerAqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ForceFieldPictures() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceFieldPictures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ForceFieldPicturesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceFieldPicturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) FramerateControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"framerateControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) FramerateControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"framerateControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) FramerateDenominator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"framerateDenominator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) FramerateDenominatorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"framerateDenominatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) FramerateNumerator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"framerateNumerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) FramerateNumeratorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"framerateNumeratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GopBReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gopBReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GopBReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gopBReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GopClosedCadence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopClosedCadence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GopClosedCadenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopClosedCadenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GopNumBFrames() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopNumBFrames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GopNumBFramesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopNumBFramesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GopSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GopSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GopSizeUnits() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gopSizeUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GopSizeUnitsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gopSizeUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) InternalValue() *TfChannel_H264SettingsProperty {
	var returns *TfChannel_H264SettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) LookAheadRateControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookAheadRateControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) LookAheadRateControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookAheadRateControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) MaxBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) MaxBitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) MinIInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) MinIIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) NumRefFrames() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRefFrames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) NumRefFramesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRefFramesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ParControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ParControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ParDenominator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parDenominator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ParDenominatorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parDenominatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ParNumerator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parNumerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ParNumeratorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parNumeratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) QualityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) QualityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) QvbrQualityLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"qvbrQualityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) QvbrQualityLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"qvbrQualityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) RateControlMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateControlMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) RateControlModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateControlModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ScanType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ScanTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) SceneChangeDetect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sceneChangeDetect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) SceneChangeDetectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sceneChangeDetectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) Slices() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) SlicesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) Softness() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"softness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) SoftnessInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"softnessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) SpatialAq() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spatialAq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) SpatialAqInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spatialAqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) SubgopLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subgopLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) SubgopLengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subgopLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) Syntax() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syntax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) SyntaxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syntaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) TemporalAq() *string {
	var returns *string
	_jsii_.Get(
		j,
		"temporalAq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) TemporalAqInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"temporalAqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) TimecodeInsertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timecodeInsertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) TimecodeInsertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timecodeInsertionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_H264SettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_H264SettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_H264SettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_H264SettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.H264SettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_H264SettingsPropertyOutputReference_Override(t TfChannel_H264SettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.H264SettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetAdaptiveQuantization(val *string) {
	if err := j.validateSetAdaptiveQuantizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adaptiveQuantization",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetAfdSignaling(val *string) {
	if err := j.validateSetAfdSignalingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afdSignaling",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetBitrate(val *float64) {
	if err := j.validateSetBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitrate",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetBufFillPct(val *float64) {
	if err := j.validateSetBufFillPctParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufFillPct",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetBufSize(val *float64) {
	if err := j.validateSetBufSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufSize",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetColorMetadata(val *string) {
	if err := j.validateSetColorMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"colorMetadata",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetEntropyEncoding(val *string) {
	if err := j.validateSetEntropyEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entropyEncoding",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetFixedAfd(val *string) {
	if err := j.validateSetFixedAfdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fixedAfd",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetFlickerAq(val *string) {
	if err := j.validateSetFlickerAqParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flickerAq",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetForceFieldPictures(val *string) {
	if err := j.validateSetForceFieldPicturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceFieldPictures",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetFramerateControl(val *string) {
	if err := j.validateSetFramerateControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"framerateControl",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetFramerateDenominator(val *float64) {
	if err := j.validateSetFramerateDenominatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"framerateDenominator",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetFramerateNumerator(val *float64) {
	if err := j.validateSetFramerateNumeratorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"framerateNumerator",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetGopBReference(val *string) {
	if err := j.validateSetGopBReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gopBReference",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetGopClosedCadence(val *float64) {
	if err := j.validateSetGopClosedCadenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gopClosedCadence",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetGopNumBFrames(val *float64) {
	if err := j.validateSetGopNumBFramesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gopNumBFrames",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetGopSize(val *float64) {
	if err := j.validateSetGopSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gopSize",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetGopSizeUnits(val *string) {
	if err := j.validateSetGopSizeUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gopSizeUnits",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetInternalValue(val *TfChannel_H264SettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetLookAheadRateControl(val *string) {
	if err := j.validateSetLookAheadRateControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lookAheadRateControl",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetMaxBitrate(val *float64) {
	if err := j.validateSetMaxBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBitrate",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetMinIInterval(val *float64) {
	if err := j.validateSetMinIIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minIInterval",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetNumRefFrames(val *float64) {
	if err := j.validateSetNumRefFramesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numRefFrames",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetParControl(val *string) {
	if err := j.validateSetParControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parControl",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetParDenominator(val *float64) {
	if err := j.validateSetParDenominatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parDenominator",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetParNumerator(val *float64) {
	if err := j.validateSetParNumeratorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parNumerator",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetProfile(val *string) {
	if err := j.validateSetProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetQualityLevel(val *string) {
	if err := j.validateSetQualityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qualityLevel",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetQvbrQualityLevel(val *float64) {
	if err := j.validateSetQvbrQualityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qvbrQualityLevel",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetRateControlMode(val *string) {
	if err := j.validateSetRateControlModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateControlMode",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetScanType(val *string) {
	if err := j.validateSetScanTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanType",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetSceneChangeDetect(val *string) {
	if err := j.validateSetSceneChangeDetectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sceneChangeDetect",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetSlices(val *float64) {
	if err := j.validateSetSlicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slices",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetSoftness(val *float64) {
	if err := j.validateSetSoftnessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"softness",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetSpatialAq(val *string) {
	if err := j.validateSetSpatialAqParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spatialAq",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetSubgopLength(val *string) {
	if err := j.validateSetSubgopLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subgopLength",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetSyntax(val *string) {
	if err := j.validateSetSyntaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syntax",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetTemporalAq(val *string) {
	if err := j.validateSetTemporalAqParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"temporalAq",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference)SetTimecodeInsertion(val *string) {
	if err := j.validateSetTimecodeInsertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timecodeInsertion",
		val,
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) PutFilterSettings(value *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsProperty) {
	if err := t.validatePutFilterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFilterSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetAdaptiveQuantization() {
	_jsii_.InvokeVoid(
		t,
		"resetAdaptiveQuantization",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetAfdSignaling() {
	_jsii_.InvokeVoid(
		t,
		"resetAfdSignaling",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetBitrate() {
	_jsii_.InvokeVoid(
		t,
		"resetBitrate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetBufFillPct() {
	_jsii_.InvokeVoid(
		t,
		"resetBufFillPct",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetBufSize() {
	_jsii_.InvokeVoid(
		t,
		"resetBufSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetColorMetadata() {
	_jsii_.InvokeVoid(
		t,
		"resetColorMetadata",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetEntropyEncoding() {
	_jsii_.InvokeVoid(
		t,
		"resetEntropyEncoding",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetFilterSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetFilterSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetFixedAfd() {
	_jsii_.InvokeVoid(
		t,
		"resetFixedAfd",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetFlickerAq() {
	_jsii_.InvokeVoid(
		t,
		"resetFlickerAq",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetForceFieldPictures() {
	_jsii_.InvokeVoid(
		t,
		"resetForceFieldPictures",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetFramerateControl() {
	_jsii_.InvokeVoid(
		t,
		"resetFramerateControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetFramerateDenominator() {
	_jsii_.InvokeVoid(
		t,
		"resetFramerateDenominator",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetFramerateNumerator() {
	_jsii_.InvokeVoid(
		t,
		"resetFramerateNumerator",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetGopBReference() {
	_jsii_.InvokeVoid(
		t,
		"resetGopBReference",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetGopClosedCadence() {
	_jsii_.InvokeVoid(
		t,
		"resetGopClosedCadence",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetGopNumBFrames() {
	_jsii_.InvokeVoid(
		t,
		"resetGopNumBFrames",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetGopSize() {
	_jsii_.InvokeVoid(
		t,
		"resetGopSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetGopSizeUnits() {
	_jsii_.InvokeVoid(
		t,
		"resetGopSizeUnits",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetLevel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetLookAheadRateControl() {
	_jsii_.InvokeVoid(
		t,
		"resetLookAheadRateControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetMaxBitrate() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxBitrate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetMinIInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetMinIInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetNumRefFrames() {
	_jsii_.InvokeVoid(
		t,
		"resetNumRefFrames",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetParControl() {
	_jsii_.InvokeVoid(
		t,
		"resetParControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetParDenominator() {
	_jsii_.InvokeVoid(
		t,
		"resetParDenominator",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetParNumerator() {
	_jsii_.InvokeVoid(
		t,
		"resetParNumerator",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetProfile() {
	_jsii_.InvokeVoid(
		t,
		"resetProfile",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetQualityLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetQualityLevel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetQvbrQualityLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetQvbrQualityLevel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetRateControlMode() {
	_jsii_.InvokeVoid(
		t,
		"resetRateControlMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetScanType() {
	_jsii_.InvokeVoid(
		t,
		"resetScanType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetSceneChangeDetect() {
	_jsii_.InvokeVoid(
		t,
		"resetSceneChangeDetect",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetSlices() {
	_jsii_.InvokeVoid(
		t,
		"resetSlices",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetSoftness() {
	_jsii_.InvokeVoid(
		t,
		"resetSoftness",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetSpatialAq() {
	_jsii_.InvokeVoid(
		t,
		"resetSpatialAq",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetSubgopLength() {
	_jsii_.InvokeVoid(
		t,
		"resetSubgopLength",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetSyntax() {
	_jsii_.InvokeVoid(
		t,
		"resetSyntax",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetTemporalAq() {
	_jsii_.InvokeVoid(
		t,
		"resetTemporalAq",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ResetTimecodeInsertion() {
	_jsii_.InvokeVoid(
		t,
		"resetTimecodeInsertion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_H264SettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

