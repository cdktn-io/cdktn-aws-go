package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_HlsGroupSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdMarkers() *[]*string
	// Experimental.
	SetAdMarkers(val *[]*string)
	// Experimental.
	AdMarkersInput() *[]*string
	// Experimental.
	BaseUrlContent() *string
	// Experimental.
	SetBaseUrlContent(val *string)
	// Experimental.
	BaseUrlContent1() *string
	// Experimental.
	SetBaseUrlContent1(val *string)
	// Experimental.
	BaseUrlContent1Input() *string
	// Experimental.
	BaseUrlContentInput() *string
	// Experimental.
	BaseUrlManifest() *string
	// Experimental.
	SetBaseUrlManifest(val *string)
	// Experimental.
	BaseUrlManifest1() *string
	// Experimental.
	SetBaseUrlManifest1(val *string)
	// Experimental.
	BaseUrlManifest1Input() *string
	// Experimental.
	BaseUrlManifestInput() *string
	// Experimental.
	CaptionLanguageMappings() TfChannel_CaptionLanguageMappingsPropertyList
	// Experimental.
	CaptionLanguageMappingsInput() interface{}
	// Experimental.
	CaptionLanguageSetting() *string
	// Experimental.
	SetCaptionLanguageSetting(val *string)
	// Experimental.
	CaptionLanguageSettingInput() *string
	// Experimental.
	ClientCache() *string
	// Experimental.
	SetClientCache(val *string)
	// Experimental.
	ClientCacheInput() *string
	// Experimental.
	CodecSpecification() *string
	// Experimental.
	SetCodecSpecification(val *string)
	// Experimental.
	CodecSpecificationInput() *string
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
	// Experimental.
	ConstantIv() *string
	// Experimental.
	SetConstantIv(val *string)
	// Experimental.
	ConstantIvInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Destination() TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationPropertyOutputReference
	// Experimental.
	DestinationInput() *TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationProperty
	// Experimental.
	DirectoryStructure() *string
	// Experimental.
	SetDirectoryStructure(val *string)
	// Experimental.
	DirectoryStructureInput() *string
	// Experimental.
	DiscontinuityTags() *string
	// Experimental.
	SetDiscontinuityTags(val *string)
	// Experimental.
	DiscontinuityTagsInput() *string
	// Experimental.
	EncryptionType() *string
	// Experimental.
	SetEncryptionType(val *string)
	// Experimental.
	EncryptionTypeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	HlsCdnSettings() TfChannel_HlsCdnSettingsPropertyList
	// Experimental.
	HlsCdnSettingsInput() interface{}
	// Experimental.
	HlsId3SegmentTagging() *string
	// Experimental.
	SetHlsId3SegmentTagging(val *string)
	// Experimental.
	HlsId3SegmentTaggingInput() *string
	// Experimental.
	IframeOnlyPlaylists() *string
	// Experimental.
	SetIframeOnlyPlaylists(val *string)
	// Experimental.
	IframeOnlyPlaylistsInput() *string
	// Experimental.
	IncompleteSegmentBehavior() *string
	// Experimental.
	SetIncompleteSegmentBehavior(val *string)
	// Experimental.
	IncompleteSegmentBehaviorInput() *string
	// Experimental.
	IndexNSegments() *float64
	// Experimental.
	SetIndexNSegments(val *float64)
	// Experimental.
	IndexNSegmentsInput() *float64
	// Experimental.
	InputLossAction() *string
	// Experimental.
	SetInputLossAction(val *string)
	// Experimental.
	InputLossActionInput() *string
	// Experimental.
	InternalValue() *TfChannel_HlsGroupSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_HlsGroupSettingsProperty)
	// Experimental.
	IvInManifest() *string
	// Experimental.
	SetIvInManifest(val *string)
	// Experimental.
	IvInManifestInput() *string
	// Experimental.
	IvSource() *string
	// Experimental.
	SetIvSource(val *string)
	// Experimental.
	IvSourceInput() *string
	// Experimental.
	KeepSegments() *float64
	// Experimental.
	SetKeepSegments(val *float64)
	// Experimental.
	KeepSegmentsInput() *float64
	// Experimental.
	KeyFormat() *string
	// Experimental.
	SetKeyFormat(val *string)
	// Experimental.
	KeyFormatInput() *string
	// Experimental.
	KeyFormatVersions() *string
	// Experimental.
	SetKeyFormatVersions(val *string)
	// Experimental.
	KeyFormatVersionsInput() *string
	// Experimental.
	KeyProviderSettings() TfChannel_KeyProviderSettingsPropertyOutputReference
	// Experimental.
	KeyProviderSettingsInput() *TfChannel_KeyProviderSettingsProperty
	// Experimental.
	ManifestCompression() *string
	// Experimental.
	SetManifestCompression(val *string)
	// Experimental.
	ManifestCompressionInput() *string
	// Experimental.
	ManifestDurationFormat() *string
	// Experimental.
	SetManifestDurationFormat(val *string)
	// Experimental.
	ManifestDurationFormatInput() *string
	// Experimental.
	MinSegmentLength() *float64
	// Experimental.
	SetMinSegmentLength(val *float64)
	// Experimental.
	MinSegmentLengthInput() *float64
	// Experimental.
	Mode() *string
	// Experimental.
	SetMode(val *string)
	// Experimental.
	ModeInput() *string
	// Experimental.
	OutputSelection() *string
	// Experimental.
	SetOutputSelection(val *string)
	// Experimental.
	OutputSelectionInput() *string
	// Experimental.
	ProgramDateTime() *string
	// Experimental.
	SetProgramDateTime(val *string)
	// Experimental.
	ProgramDateTimeClock() *string
	// Experimental.
	SetProgramDateTimeClock(val *string)
	// Experimental.
	ProgramDateTimeClockInput() *string
	// Experimental.
	ProgramDateTimeInput() *string
	// Experimental.
	ProgramDateTimePeriod() *float64
	// Experimental.
	SetProgramDateTimePeriod(val *float64)
	// Experimental.
	ProgramDateTimePeriodInput() *float64
	// Experimental.
	RedundantManifest() *string
	// Experimental.
	SetRedundantManifest(val *string)
	// Experimental.
	RedundantManifestInput() *string
	// Experimental.
	SegmentLength() *float64
	// Experimental.
	SetSegmentLength(val *float64)
	// Experimental.
	SegmentLengthInput() *float64
	// Experimental.
	SegmentsPerSubdirectory() *float64
	// Experimental.
	SetSegmentsPerSubdirectory(val *float64)
	// Experimental.
	SegmentsPerSubdirectoryInput() *float64
	// Experimental.
	StreamInfResolution() *string
	// Experimental.
	SetStreamInfResolution(val *string)
	// Experimental.
	StreamInfResolutionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimedMetadataId3Frame() *string
	// Experimental.
	SetTimedMetadataId3Frame(val *string)
	// Experimental.
	TimedMetadataId3FrameInput() *string
	// Experimental.
	TimedMetadataId3Period() *float64
	// Experimental.
	SetTimedMetadataId3Period(val *float64)
	// Experimental.
	TimedMetadataId3PeriodInput() *float64
	// Experimental.
	TimestampDeltaMilliseconds() *float64
	// Experimental.
	SetTimestampDeltaMilliseconds(val *float64)
	// Experimental.
	TimestampDeltaMillisecondsInput() *float64
	// Experimental.
	TsFileMode() *string
	// Experimental.
	SetTsFileMode(val *string)
	// Experimental.
	TsFileModeInput() *string
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
	PutCaptionLanguageMappings(value interface{})
	// Experimental.
	PutDestination(value *TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationProperty)
	// Experimental.
	PutHlsCdnSettings(value interface{})
	// Experimental.
	PutKeyProviderSettings(value *TfChannel_KeyProviderSettingsProperty)
	// Experimental.
	ResetAdMarkers()
	// Experimental.
	ResetBaseUrlContent()
	// Experimental.
	ResetBaseUrlContent1()
	// Experimental.
	ResetBaseUrlManifest()
	// Experimental.
	ResetBaseUrlManifest1()
	// Experimental.
	ResetCaptionLanguageMappings()
	// Experimental.
	ResetCaptionLanguageSetting()
	// Experimental.
	ResetClientCache()
	// Experimental.
	ResetCodecSpecification()
	// Experimental.
	ResetConstantIv()
	// Experimental.
	ResetDirectoryStructure()
	// Experimental.
	ResetDiscontinuityTags()
	// Experimental.
	ResetEncryptionType()
	// Experimental.
	ResetHlsCdnSettings()
	// Experimental.
	ResetHlsId3SegmentTagging()
	// Experimental.
	ResetIframeOnlyPlaylists()
	// Experimental.
	ResetIncompleteSegmentBehavior()
	// Experimental.
	ResetIndexNSegments()
	// Experimental.
	ResetInputLossAction()
	// Experimental.
	ResetIvInManifest()
	// Experimental.
	ResetIvSource()
	// Experimental.
	ResetKeepSegments()
	// Experimental.
	ResetKeyFormat()
	// Experimental.
	ResetKeyFormatVersions()
	// Experimental.
	ResetKeyProviderSettings()
	// Experimental.
	ResetManifestCompression()
	// Experimental.
	ResetManifestDurationFormat()
	// Experimental.
	ResetMinSegmentLength()
	// Experimental.
	ResetMode()
	// Experimental.
	ResetOutputSelection()
	// Experimental.
	ResetProgramDateTime()
	// Experimental.
	ResetProgramDateTimeClock()
	// Experimental.
	ResetProgramDateTimePeriod()
	// Experimental.
	ResetRedundantManifest()
	// Experimental.
	ResetSegmentLength()
	// Experimental.
	ResetSegmentsPerSubdirectory()
	// Experimental.
	ResetStreamInfResolution()
	// Experimental.
	ResetTimedMetadataId3Frame()
	// Experimental.
	ResetTimedMetadataId3Period()
	// Experimental.
	ResetTimestampDeltaMilliseconds()
	// Experimental.
	ResetTsFileMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_HlsGroupSettingsPropertyOutputReference
type jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) AdMarkers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adMarkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) AdMarkersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adMarkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) BaseUrlContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) BaseUrlContent1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlContent1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) BaseUrlContent1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlContent1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) BaseUrlContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) BaseUrlManifest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlManifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) BaseUrlManifest1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlManifest1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) BaseUrlManifest1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlManifest1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) BaseUrlManifestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlManifestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) CaptionLanguageMappings() TfChannel_CaptionLanguageMappingsPropertyList {
	var returns TfChannel_CaptionLanguageMappingsPropertyList
	_jsii_.Get(
		j,
		"captionLanguageMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) CaptionLanguageMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captionLanguageMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) CaptionLanguageSetting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"captionLanguageSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) CaptionLanguageSettingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"captionLanguageSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ClientCache() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ClientCacheInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) CodecSpecification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codecSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) CodecSpecificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codecSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ConstantIv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constantIv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ConstantIvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constantIvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) Destination() TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationPropertyOutputReference {
	var returns TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) DestinationInput() *TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationProperty {
	var returns *TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationProperty
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) DirectoryStructure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryStructure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) DirectoryStructureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryStructureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) DiscontinuityTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discontinuityTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) DiscontinuityTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discontinuityTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) EncryptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) EncryptionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) HlsCdnSettings() TfChannel_HlsCdnSettingsPropertyList {
	var returns TfChannel_HlsCdnSettingsPropertyList
	_jsii_.Get(
		j,
		"hlsCdnSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) HlsCdnSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsCdnSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) HlsId3SegmentTagging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hlsId3SegmentTagging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) HlsId3SegmentTaggingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hlsId3SegmentTaggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) IframeOnlyPlaylists() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iframeOnlyPlaylists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) IframeOnlyPlaylistsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iframeOnlyPlaylistsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) IncompleteSegmentBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"incompleteSegmentBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) IncompleteSegmentBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"incompleteSegmentBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) IndexNSegments() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indexNSegments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) IndexNSegmentsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indexNSegmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) InputLossAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) InputLossActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) InternalValue() *TfChannel_HlsGroupSettingsProperty {
	var returns *TfChannel_HlsGroupSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) IvInManifest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ivInManifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) IvInManifestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ivInManifestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) IvSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ivSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) IvSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ivSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) KeepSegments() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepSegments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) KeepSegmentsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepSegmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) KeyFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) KeyFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) KeyFormatVersions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFormatVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) KeyFormatVersionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFormatVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) KeyProviderSettings() TfChannel_KeyProviderSettingsPropertyOutputReference {
	var returns TfChannel_KeyProviderSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"keyProviderSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) KeyProviderSettingsInput() *TfChannel_KeyProviderSettingsProperty {
	var returns *TfChannel_KeyProviderSettingsProperty
	_jsii_.Get(
		j,
		"keyProviderSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ManifestCompression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestCompression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ManifestCompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestCompressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ManifestDurationFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestDurationFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ManifestDurationFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestDurationFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) MinSegmentLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSegmentLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) MinSegmentLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSegmentLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) OutputSelection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) OutputSelectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ProgramDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ProgramDateTimeClock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programDateTimeClock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ProgramDateTimeClockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programDateTimeClockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ProgramDateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programDateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ProgramDateTimePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programDateTimePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ProgramDateTimePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programDateTimePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) RedundantManifest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redundantManifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) RedundantManifestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redundantManifestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) SegmentLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) SegmentLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) SegmentsPerSubdirectory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentsPerSubdirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) SegmentsPerSubdirectoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentsPerSubdirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) StreamInfResolution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamInfResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) StreamInfResolutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamInfResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) TimedMetadataId3Frame() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timedMetadataId3Frame",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) TimedMetadataId3FrameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timedMetadataId3FrameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) TimedMetadataId3Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timedMetadataId3Period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) TimedMetadataId3PeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timedMetadataId3PeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) TimestampDeltaMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timestampDeltaMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) TimestampDeltaMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timestampDeltaMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) TsFileMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tsFileMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) TsFileModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tsFileModeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_HlsGroupSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_HlsGroupSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_HlsGroupSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.HlsGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_HlsGroupSettingsPropertyOutputReference_Override(t TfChannel_HlsGroupSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.HlsGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetAdMarkers(val *[]*string) {
	if err := j.validateSetAdMarkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adMarkers",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetBaseUrlContent(val *string) {
	if err := j.validateSetBaseUrlContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseUrlContent",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetBaseUrlContent1(val *string) {
	if err := j.validateSetBaseUrlContent1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseUrlContent1",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetBaseUrlManifest(val *string) {
	if err := j.validateSetBaseUrlManifestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseUrlManifest",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetBaseUrlManifest1(val *string) {
	if err := j.validateSetBaseUrlManifest1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseUrlManifest1",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetCaptionLanguageSetting(val *string) {
	if err := j.validateSetCaptionLanguageSettingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"captionLanguageSetting",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetClientCache(val *string) {
	if err := j.validateSetClientCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCache",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetCodecSpecification(val *string) {
	if err := j.validateSetCodecSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codecSpecification",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetConstantIv(val *string) {
	if err := j.validateSetConstantIvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"constantIv",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetDirectoryStructure(val *string) {
	if err := j.validateSetDirectoryStructureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryStructure",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetDiscontinuityTags(val *string) {
	if err := j.validateSetDiscontinuityTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discontinuityTags",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetEncryptionType(val *string) {
	if err := j.validateSetEncryptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionType",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetHlsId3SegmentTagging(val *string) {
	if err := j.validateSetHlsId3SegmentTaggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hlsId3SegmentTagging",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetIframeOnlyPlaylists(val *string) {
	if err := j.validateSetIframeOnlyPlaylistsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iframeOnlyPlaylists",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetIncompleteSegmentBehavior(val *string) {
	if err := j.validateSetIncompleteSegmentBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"incompleteSegmentBehavior",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetIndexNSegments(val *float64) {
	if err := j.validateSetIndexNSegmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexNSegments",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetInputLossAction(val *string) {
	if err := j.validateSetInputLossActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputLossAction",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_HlsGroupSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetIvInManifest(val *string) {
	if err := j.validateSetIvInManifestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ivInManifest",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetIvSource(val *string) {
	if err := j.validateSetIvSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ivSource",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetKeepSegments(val *float64) {
	if err := j.validateSetKeepSegmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepSegments",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetKeyFormat(val *string) {
	if err := j.validateSetKeyFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyFormat",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetKeyFormatVersions(val *string) {
	if err := j.validateSetKeyFormatVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyFormatVersions",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetManifestCompression(val *string) {
	if err := j.validateSetManifestCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestCompression",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetManifestDurationFormat(val *string) {
	if err := j.validateSetManifestDurationFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestDurationFormat",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetMinSegmentLength(val *float64) {
	if err := j.validateSetMinSegmentLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSegmentLength",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetOutputSelection(val *string) {
	if err := j.validateSetOutputSelectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputSelection",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetProgramDateTime(val *string) {
	if err := j.validateSetProgramDateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programDateTime",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetProgramDateTimeClock(val *string) {
	if err := j.validateSetProgramDateTimeClockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programDateTimeClock",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetProgramDateTimePeriod(val *float64) {
	if err := j.validateSetProgramDateTimePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programDateTimePeriod",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetRedundantManifest(val *string) {
	if err := j.validateSetRedundantManifestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redundantManifest",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetSegmentLength(val *float64) {
	if err := j.validateSetSegmentLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentLength",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetSegmentsPerSubdirectory(val *float64) {
	if err := j.validateSetSegmentsPerSubdirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentsPerSubdirectory",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetStreamInfResolution(val *string) {
	if err := j.validateSetStreamInfResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamInfResolution",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetTimedMetadataId3Frame(val *string) {
	if err := j.validateSetTimedMetadataId3FrameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timedMetadataId3Frame",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetTimedMetadataId3Period(val *float64) {
	if err := j.validateSetTimedMetadataId3PeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timedMetadataId3Period",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetTimestampDeltaMilliseconds(val *float64) {
	if err := j.validateSetTimestampDeltaMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampDeltaMilliseconds",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference)SetTsFileMode(val *string) {
	if err := j.validateSetTsFileModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tsFileMode",
		val,
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) PutCaptionLanguageMappings(value interface{}) {
	if err := t.validatePutCaptionLanguageMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCaptionLanguageMappings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) PutDestination(value *TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationProperty) {
	if err := t.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDestination",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) PutHlsCdnSettings(value interface{}) {
	if err := t.validatePutHlsCdnSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHlsCdnSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) PutKeyProviderSettings(value *TfChannel_KeyProviderSettingsProperty) {
	if err := t.validatePutKeyProviderSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKeyProviderSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetAdMarkers() {
	_jsii_.InvokeVoid(
		t,
		"resetAdMarkers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetBaseUrlContent() {
	_jsii_.InvokeVoid(
		t,
		"resetBaseUrlContent",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetBaseUrlContent1() {
	_jsii_.InvokeVoid(
		t,
		"resetBaseUrlContent1",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetBaseUrlManifest() {
	_jsii_.InvokeVoid(
		t,
		"resetBaseUrlManifest",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetBaseUrlManifest1() {
	_jsii_.InvokeVoid(
		t,
		"resetBaseUrlManifest1",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetCaptionLanguageMappings() {
	_jsii_.InvokeVoid(
		t,
		"resetCaptionLanguageMappings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetCaptionLanguageSetting() {
	_jsii_.InvokeVoid(
		t,
		"resetCaptionLanguageSetting",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetClientCache() {
	_jsii_.InvokeVoid(
		t,
		"resetClientCache",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetCodecSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetCodecSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetConstantIv() {
	_jsii_.InvokeVoid(
		t,
		"resetConstantIv",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetDirectoryStructure() {
	_jsii_.InvokeVoid(
		t,
		"resetDirectoryStructure",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetDiscontinuityTags() {
	_jsii_.InvokeVoid(
		t,
		"resetDiscontinuityTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetEncryptionType() {
	_jsii_.InvokeVoid(
		t,
		"resetEncryptionType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetHlsCdnSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetHlsCdnSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetHlsId3SegmentTagging() {
	_jsii_.InvokeVoid(
		t,
		"resetHlsId3SegmentTagging",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetIframeOnlyPlaylists() {
	_jsii_.InvokeVoid(
		t,
		"resetIframeOnlyPlaylists",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetIncompleteSegmentBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetIncompleteSegmentBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetIndexNSegments() {
	_jsii_.InvokeVoid(
		t,
		"resetIndexNSegments",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetInputLossAction() {
	_jsii_.InvokeVoid(
		t,
		"resetInputLossAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetIvInManifest() {
	_jsii_.InvokeVoid(
		t,
		"resetIvInManifest",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetIvSource() {
	_jsii_.InvokeVoid(
		t,
		"resetIvSource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetKeepSegments() {
	_jsii_.InvokeVoid(
		t,
		"resetKeepSegments",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetKeyFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetKeyFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetKeyFormatVersions() {
	_jsii_.InvokeVoid(
		t,
		"resetKeyFormatVersions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetKeyProviderSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetKeyProviderSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetManifestCompression() {
	_jsii_.InvokeVoid(
		t,
		"resetManifestCompression",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetManifestDurationFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetManifestDurationFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetMinSegmentLength() {
	_jsii_.InvokeVoid(
		t,
		"resetMinSegmentLength",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		t,
		"resetMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetOutputSelection() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputSelection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetProgramDateTime() {
	_jsii_.InvokeVoid(
		t,
		"resetProgramDateTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetProgramDateTimeClock() {
	_jsii_.InvokeVoid(
		t,
		"resetProgramDateTimeClock",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetProgramDateTimePeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetProgramDateTimePeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetRedundantManifest() {
	_jsii_.InvokeVoid(
		t,
		"resetRedundantManifest",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetSegmentLength() {
	_jsii_.InvokeVoid(
		t,
		"resetSegmentLength",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetSegmentsPerSubdirectory() {
	_jsii_.InvokeVoid(
		t,
		"resetSegmentsPerSubdirectory",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetStreamInfResolution() {
	_jsii_.InvokeVoid(
		t,
		"resetStreamInfResolution",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetTimedMetadataId3Frame() {
	_jsii_.InvokeVoid(
		t,
		"resetTimedMetadataId3Frame",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetTimedMetadataId3Period() {
	_jsii_.InvokeVoid(
		t,
		"resetTimedMetadataId3Period",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetTimestampDeltaMilliseconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTimestampDeltaMilliseconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ResetTsFileMode() {
	_jsii_.InvokeVoid(
		t,
		"resetTsFileMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_HlsGroupSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

