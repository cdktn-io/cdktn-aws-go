package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_H265SettingsPropertyOutputReference interface {
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
	AlternativeTransferFunction() *string
	// Experimental.
	SetAlternativeTransferFunction(val *string)
	// Experimental.
	AlternativeTransferFunctionInput() *string
	// Experimental.
	Bitrate() *float64
	// Experimental.
	SetBitrate(val *float64)
	// Experimental.
	BitrateInput() *float64
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
	// Experimental.
	ColorSpaceSettings() AwsChannel_ColorSpaceSettingsPropertyOutputReference
	// Experimental.
	ColorSpaceSettingsInput() *AwsChannel_ColorSpaceSettingsProperty
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
	FilterSettings() AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsPropertyOutputReference
	// Experimental.
	FilterSettingsInput() *AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsProperty
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
	Fqn() *string
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
	GopClosedCadence() *float64
	// Experimental.
	SetGopClosedCadence(val *float64)
	// Experimental.
	GopClosedCadenceInput() *float64
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
	InternalValue() *AwsChannel_H265SettingsProperty
	// Experimental.
	SetInternalValue(val *AwsChannel_H265SettingsProperty)
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
	MinQp() *float64
	// Experimental.
	SetMinQp(val *float64)
	// Experimental.
	MinQpInput() *float64
	// Experimental.
	MvOverPictureBoundaries() *string
	// Experimental.
	SetMvOverPictureBoundaries(val *string)
	// Experimental.
	MvOverPictureBoundariesInput() *string
	// Experimental.
	MvTemporalPredictor() *string
	// Experimental.
	SetMvTemporalPredictor(val *string)
	// Experimental.
	MvTemporalPredictorInput() *string
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
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Tier() *string
	// Experimental.
	SetTier(val *string)
	// Experimental.
	TierInput() *string
	// Experimental.
	TileHeight() *float64
	// Experimental.
	SetTileHeight(val *float64)
	// Experimental.
	TileHeightInput() *float64
	// Experimental.
	TilePadding() *string
	// Experimental.
	SetTilePadding(val *string)
	// Experimental.
	TilePaddingInput() *string
	// Experimental.
	TileWidth() *float64
	// Experimental.
	SetTileWidth(val *float64)
	// Experimental.
	TileWidthInput() *float64
	// Experimental.
	TimecodeBurninSettings() AwsChannel_TimecodeBurninSettingsPropertyOutputReference
	// Experimental.
	TimecodeBurninSettingsInput() *AwsChannel_TimecodeBurninSettingsProperty
	// Experimental.
	TimecodeInsertion() *string
	// Experimental.
	SetTimecodeInsertion(val *string)
	// Experimental.
	TimecodeInsertionInput() *string
	// Experimental.
	TreeblockSize() *string
	// Experimental.
	SetTreeblockSize(val *string)
	// Experimental.
	TreeblockSizeInput() *string
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
	PutColorSpaceSettings(value *AwsChannel_ColorSpaceSettingsProperty)
	// Experimental.
	PutFilterSettings(value *AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsProperty)
	// Experimental.
	PutTimecodeBurninSettings(value *AwsChannel_TimecodeBurninSettingsProperty)
	// Experimental.
	ResetAdaptiveQuantization()
	// Experimental.
	ResetAfdSignaling()
	// Experimental.
	ResetAlternativeTransferFunction()
	// Experimental.
	ResetBufSize()
	// Experimental.
	ResetColorMetadata()
	// Experimental.
	ResetColorSpaceSettings()
	// Experimental.
	ResetFilterSettings()
	// Experimental.
	ResetFixedAfd()
	// Experimental.
	ResetFlickerAq()
	// Experimental.
	ResetGopClosedCadence()
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
	ResetMinQp()
	// Experimental.
	ResetMvOverPictureBoundaries()
	// Experimental.
	ResetMvTemporalPredictor()
	// Experimental.
	ResetParDenominator()
	// Experimental.
	ResetParNumerator()
	// Experimental.
	ResetProfile()
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
	ResetTier()
	// Experimental.
	ResetTileHeight()
	// Experimental.
	ResetTilePadding()
	// Experimental.
	ResetTileWidth()
	// Experimental.
	ResetTimecodeBurninSettings()
	// Experimental.
	ResetTimecodeInsertion()
	// Experimental.
	ResetTreeblockSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsChannel_H265SettingsPropertyOutputReference
type jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) AdaptiveQuantization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adaptiveQuantization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) AdaptiveQuantizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adaptiveQuantizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) AfdSignaling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afdSignaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) AfdSignalingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afdSignalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) AlternativeTransferFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alternativeTransferFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) AlternativeTransferFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alternativeTransferFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) Bitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) BitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) BufSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) BufSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ColorMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ColorMetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ColorSpaceSettings() AwsChannel_ColorSpaceSettingsPropertyOutputReference {
	var returns AwsChannel_ColorSpaceSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"colorSpaceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ColorSpaceSettingsInput() *AwsChannel_ColorSpaceSettingsProperty {
	var returns *AwsChannel_ColorSpaceSettingsProperty
	_jsii_.Get(
		j,
		"colorSpaceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) FilterSettings() AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsPropertyOutputReference {
	var returns AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"filterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) FilterSettingsInput() *AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsProperty {
	var returns *AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsProperty
	_jsii_.Get(
		j,
		"filterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) FixedAfd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedAfd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) FixedAfdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedAfdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) FlickerAq() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flickerAq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) FlickerAqInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flickerAqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) FramerateDenominator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"framerateDenominator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) FramerateDenominatorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"framerateDenominatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) FramerateNumerator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"framerateNumerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) FramerateNumeratorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"framerateNumeratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) GopClosedCadence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopClosedCadence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) GopClosedCadenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopClosedCadenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) GopSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) GopSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) GopSizeUnits() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gopSizeUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) GopSizeUnitsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gopSizeUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) InternalValue() *AwsChannel_H265SettingsProperty {
	var returns *AwsChannel_H265SettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) LookAheadRateControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookAheadRateControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) LookAheadRateControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookAheadRateControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) MaxBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) MaxBitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) MinIInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) MinIIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) MinQp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minQp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) MinQpInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minQpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) MvOverPictureBoundaries() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mvOverPictureBoundaries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) MvOverPictureBoundariesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mvOverPictureBoundariesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) MvTemporalPredictor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mvTemporalPredictor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) MvTemporalPredictorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mvTemporalPredictorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ParDenominator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parDenominator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ParDenominatorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parDenominatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ParNumerator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parNumerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ParNumeratorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parNumeratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) QvbrQualityLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"qvbrQualityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) QvbrQualityLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"qvbrQualityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) RateControlMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateControlMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) RateControlModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateControlModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ScanType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ScanTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) SceneChangeDetect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sceneChangeDetect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) SceneChangeDetectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sceneChangeDetectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) Slices() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) SlicesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) TileHeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tileHeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) TileHeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tileHeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) TilePadding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tilePadding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) TilePaddingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tilePaddingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) TileWidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tileWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) TileWidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tileWidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) TimecodeBurninSettings() AwsChannel_TimecodeBurninSettingsPropertyOutputReference {
	var returns AwsChannel_TimecodeBurninSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"timecodeBurninSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) TimecodeBurninSettingsInput() *AwsChannel_TimecodeBurninSettingsProperty {
	var returns *AwsChannel_TimecodeBurninSettingsProperty
	_jsii_.Get(
		j,
		"timecodeBurninSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) TimecodeInsertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timecodeInsertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) TimecodeInsertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timecodeInsertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) TreeblockSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treeblockSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) TreeblockSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treeblockSizeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_H265SettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChannel_H265SettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_H265SettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.H265SettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_H265SettingsPropertyOutputReference_Override(a AwsChannel_H265SettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.H265SettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetAdaptiveQuantization(val *string) {
	if err := j.validateSetAdaptiveQuantizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adaptiveQuantization",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetAfdSignaling(val *string) {
	if err := j.validateSetAfdSignalingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afdSignaling",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetAlternativeTransferFunction(val *string) {
	if err := j.validateSetAlternativeTransferFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alternativeTransferFunction",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetBitrate(val *float64) {
	if err := j.validateSetBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitrate",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetBufSize(val *float64) {
	if err := j.validateSetBufSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufSize",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetColorMetadata(val *string) {
	if err := j.validateSetColorMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"colorMetadata",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetFixedAfd(val *string) {
	if err := j.validateSetFixedAfdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fixedAfd",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetFlickerAq(val *string) {
	if err := j.validateSetFlickerAqParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flickerAq",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetFramerateDenominator(val *float64) {
	if err := j.validateSetFramerateDenominatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"framerateDenominator",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetFramerateNumerator(val *float64) {
	if err := j.validateSetFramerateNumeratorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"framerateNumerator",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetGopClosedCadence(val *float64) {
	if err := j.validateSetGopClosedCadenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gopClosedCadence",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetGopSize(val *float64) {
	if err := j.validateSetGopSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gopSize",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetGopSizeUnits(val *string) {
	if err := j.validateSetGopSizeUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gopSizeUnits",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetInternalValue(val *AwsChannel_H265SettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetLookAheadRateControl(val *string) {
	if err := j.validateSetLookAheadRateControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lookAheadRateControl",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetMaxBitrate(val *float64) {
	if err := j.validateSetMaxBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBitrate",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetMinIInterval(val *float64) {
	if err := j.validateSetMinIIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minIInterval",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetMinQp(val *float64) {
	if err := j.validateSetMinQpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minQp",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetMvOverPictureBoundaries(val *string) {
	if err := j.validateSetMvOverPictureBoundariesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mvOverPictureBoundaries",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetMvTemporalPredictor(val *string) {
	if err := j.validateSetMvTemporalPredictorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mvTemporalPredictor",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetParDenominator(val *float64) {
	if err := j.validateSetParDenominatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parDenominator",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetParNumerator(val *float64) {
	if err := j.validateSetParNumeratorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parNumerator",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetProfile(val *string) {
	if err := j.validateSetProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetQvbrQualityLevel(val *float64) {
	if err := j.validateSetQvbrQualityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qvbrQualityLevel",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetRateControlMode(val *string) {
	if err := j.validateSetRateControlModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateControlMode",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetScanType(val *string) {
	if err := j.validateSetScanTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanType",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetSceneChangeDetect(val *string) {
	if err := j.validateSetSceneChangeDetectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sceneChangeDetect",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetSlices(val *float64) {
	if err := j.validateSetSlicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slices",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetTier(val *string) {
	if err := j.validateSetTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetTileHeight(val *float64) {
	if err := j.validateSetTileHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tileHeight",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetTilePadding(val *string) {
	if err := j.validateSetTilePaddingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tilePadding",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetTileWidth(val *float64) {
	if err := j.validateSetTileWidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tileWidth",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetTimecodeInsertion(val *string) {
	if err := j.validateSetTimecodeInsertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timecodeInsertion",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference)SetTreeblockSize(val *string) {
	if err := j.validateSetTreeblockSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"treeblockSize",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) PutColorSpaceSettings(value *AwsChannel_ColorSpaceSettingsProperty) {
	if err := a.validatePutColorSpaceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putColorSpaceSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) PutFilterSettings(value *AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsProperty) {
	if err := a.validatePutFilterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilterSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) PutTimecodeBurninSettings(value *AwsChannel_TimecodeBurninSettingsProperty) {
	if err := a.validatePutTimecodeBurninSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimecodeBurninSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetAdaptiveQuantization() {
	_jsii_.InvokeVoid(
		a,
		"resetAdaptiveQuantization",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetAfdSignaling() {
	_jsii_.InvokeVoid(
		a,
		"resetAfdSignaling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetAlternativeTransferFunction() {
	_jsii_.InvokeVoid(
		a,
		"resetAlternativeTransferFunction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetBufSize() {
	_jsii_.InvokeVoid(
		a,
		"resetBufSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetColorMetadata() {
	_jsii_.InvokeVoid(
		a,
		"resetColorMetadata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetColorSpaceSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetColorSpaceSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetFilterSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetFilterSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetFixedAfd() {
	_jsii_.InvokeVoid(
		a,
		"resetFixedAfd",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetFlickerAq() {
	_jsii_.InvokeVoid(
		a,
		"resetFlickerAq",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetGopClosedCadence() {
	_jsii_.InvokeVoid(
		a,
		"resetGopClosedCadence",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetGopSize() {
	_jsii_.InvokeVoid(
		a,
		"resetGopSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetGopSizeUnits() {
	_jsii_.InvokeVoid(
		a,
		"resetGopSizeUnits",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetLookAheadRateControl() {
	_jsii_.InvokeVoid(
		a,
		"resetLookAheadRateControl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetMaxBitrate() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxBitrate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetMinIInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetMinIInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetMinQp() {
	_jsii_.InvokeVoid(
		a,
		"resetMinQp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetMvOverPictureBoundaries() {
	_jsii_.InvokeVoid(
		a,
		"resetMvOverPictureBoundaries",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetMvTemporalPredictor() {
	_jsii_.InvokeVoid(
		a,
		"resetMvTemporalPredictor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetParDenominator() {
	_jsii_.InvokeVoid(
		a,
		"resetParDenominator",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetParNumerator() {
	_jsii_.InvokeVoid(
		a,
		"resetParNumerator",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetProfile() {
	_jsii_.InvokeVoid(
		a,
		"resetProfile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetQvbrQualityLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetQvbrQualityLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetRateControlMode() {
	_jsii_.InvokeVoid(
		a,
		"resetRateControlMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetScanType() {
	_jsii_.InvokeVoid(
		a,
		"resetScanType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetSceneChangeDetect() {
	_jsii_.InvokeVoid(
		a,
		"resetSceneChangeDetect",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetSlices() {
	_jsii_.InvokeVoid(
		a,
		"resetSlices",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetTier() {
	_jsii_.InvokeVoid(
		a,
		"resetTier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetTileHeight() {
	_jsii_.InvokeVoid(
		a,
		"resetTileHeight",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetTilePadding() {
	_jsii_.InvokeVoid(
		a,
		"resetTilePadding",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetTileWidth() {
	_jsii_.InvokeVoid(
		a,
		"resetTileWidth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetTimecodeBurninSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetTimecodeBurninSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetTimecodeInsertion() {
	_jsii_.InvokeVoid(
		a,
		"resetTimecodeInsertion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ResetTreeblockSize() {
	_jsii_.InvokeVoid(
		a,
		"resetTreeblockSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_H265SettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

