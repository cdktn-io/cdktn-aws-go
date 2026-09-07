package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_HlsGroupSettingsPropertyOutputReference interface {
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
	CaptionLanguageMappings() AwsChannel_CaptionLanguageMappingsPropertyList
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
	Destination() AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationPropertyOutputReference
	// Experimental.
	DestinationInput() *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationProperty
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
	HlsCdnSettings() AwsChannel_HlsCdnSettingsPropertyList
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
	InternalValue() *AwsChannel_HlsGroupSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsChannel_HlsGroupSettingsProperty)
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
	KeyProviderSettings() AwsChannel_KeyProviderSettingsPropertyOutputReference
	// Experimental.
	KeyProviderSettingsInput() *AwsChannel_KeyProviderSettingsProperty
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
	PutDestination(value *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationProperty)
	// Experimental.
	PutHlsCdnSettings(value interface{})
	// Experimental.
	PutKeyProviderSettings(value *AwsChannel_KeyProviderSettingsProperty)
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

// The jsii proxy struct for AwsChannel_HlsGroupSettingsPropertyOutputReference
type jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) AdMarkers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adMarkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) AdMarkersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adMarkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) BaseUrlContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) BaseUrlContent1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlContent1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) BaseUrlContent1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlContent1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) BaseUrlContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) BaseUrlManifest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlManifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) BaseUrlManifest1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlManifest1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) BaseUrlManifest1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlManifest1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) BaseUrlManifestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlManifestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) CaptionLanguageMappings() AwsChannel_CaptionLanguageMappingsPropertyList {
	var returns AwsChannel_CaptionLanguageMappingsPropertyList
	_jsii_.Get(
		j,
		"captionLanguageMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) CaptionLanguageMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captionLanguageMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) CaptionLanguageSetting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"captionLanguageSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) CaptionLanguageSettingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"captionLanguageSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ClientCache() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ClientCacheInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) CodecSpecification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codecSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) CodecSpecificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codecSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ConstantIv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constantIv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ConstantIvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constantIvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) Destination() AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationPropertyOutputReference {
	var returns AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) DestinationInput() *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationProperty {
	var returns *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationProperty
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) DirectoryStructure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryStructure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) DirectoryStructureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryStructureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) DiscontinuityTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discontinuityTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) DiscontinuityTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discontinuityTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) EncryptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) EncryptionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) HlsCdnSettings() AwsChannel_HlsCdnSettingsPropertyList {
	var returns AwsChannel_HlsCdnSettingsPropertyList
	_jsii_.Get(
		j,
		"hlsCdnSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) HlsCdnSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsCdnSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) HlsId3SegmentTagging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hlsId3SegmentTagging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) HlsId3SegmentTaggingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hlsId3SegmentTaggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) IframeOnlyPlaylists() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iframeOnlyPlaylists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) IframeOnlyPlaylistsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iframeOnlyPlaylistsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) IncompleteSegmentBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"incompleteSegmentBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) IncompleteSegmentBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"incompleteSegmentBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) IndexNSegments() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indexNSegments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) IndexNSegmentsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indexNSegmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) InputLossAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) InputLossActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) InternalValue() *AwsChannel_HlsGroupSettingsProperty {
	var returns *AwsChannel_HlsGroupSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) IvInManifest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ivInManifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) IvInManifestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ivInManifestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) IvSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ivSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) IvSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ivSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) KeepSegments() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepSegments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) KeepSegmentsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepSegmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) KeyFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) KeyFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) KeyFormatVersions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFormatVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) KeyFormatVersionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFormatVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) KeyProviderSettings() AwsChannel_KeyProviderSettingsPropertyOutputReference {
	var returns AwsChannel_KeyProviderSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"keyProviderSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) KeyProviderSettingsInput() *AwsChannel_KeyProviderSettingsProperty {
	var returns *AwsChannel_KeyProviderSettingsProperty
	_jsii_.Get(
		j,
		"keyProviderSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ManifestCompression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestCompression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ManifestCompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestCompressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ManifestDurationFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestDurationFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ManifestDurationFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestDurationFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) MinSegmentLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSegmentLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) MinSegmentLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSegmentLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) OutputSelection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) OutputSelectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ProgramDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ProgramDateTimeClock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programDateTimeClock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ProgramDateTimeClockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programDateTimeClockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ProgramDateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programDateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ProgramDateTimePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programDateTimePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ProgramDateTimePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programDateTimePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) RedundantManifest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redundantManifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) RedundantManifestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redundantManifestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) SegmentLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) SegmentLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) SegmentsPerSubdirectory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentsPerSubdirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) SegmentsPerSubdirectoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentsPerSubdirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) StreamInfResolution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamInfResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) StreamInfResolutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamInfResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) TimedMetadataId3Frame() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timedMetadataId3Frame",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) TimedMetadataId3FrameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timedMetadataId3FrameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) TimedMetadataId3Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timedMetadataId3Period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) TimedMetadataId3PeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timedMetadataId3PeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) TimestampDeltaMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timestampDeltaMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) TimestampDeltaMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timestampDeltaMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) TsFileMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tsFileMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) TsFileModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tsFileModeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_HlsGroupSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChannel_HlsGroupSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_HlsGroupSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.HlsGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_HlsGroupSettingsPropertyOutputReference_Override(a AwsChannel_HlsGroupSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.HlsGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetAdMarkers(val *[]*string) {
	if err := j.validateSetAdMarkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adMarkers",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetBaseUrlContent(val *string) {
	if err := j.validateSetBaseUrlContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseUrlContent",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetBaseUrlContent1(val *string) {
	if err := j.validateSetBaseUrlContent1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseUrlContent1",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetBaseUrlManifest(val *string) {
	if err := j.validateSetBaseUrlManifestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseUrlManifest",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetBaseUrlManifest1(val *string) {
	if err := j.validateSetBaseUrlManifest1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseUrlManifest1",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetCaptionLanguageSetting(val *string) {
	if err := j.validateSetCaptionLanguageSettingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"captionLanguageSetting",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetClientCache(val *string) {
	if err := j.validateSetClientCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCache",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetCodecSpecification(val *string) {
	if err := j.validateSetCodecSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codecSpecification",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetConstantIv(val *string) {
	if err := j.validateSetConstantIvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"constantIv",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetDirectoryStructure(val *string) {
	if err := j.validateSetDirectoryStructureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryStructure",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetDiscontinuityTags(val *string) {
	if err := j.validateSetDiscontinuityTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discontinuityTags",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetEncryptionType(val *string) {
	if err := j.validateSetEncryptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionType",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetHlsId3SegmentTagging(val *string) {
	if err := j.validateSetHlsId3SegmentTaggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hlsId3SegmentTagging",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetIframeOnlyPlaylists(val *string) {
	if err := j.validateSetIframeOnlyPlaylistsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iframeOnlyPlaylists",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetIncompleteSegmentBehavior(val *string) {
	if err := j.validateSetIncompleteSegmentBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"incompleteSegmentBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetIndexNSegments(val *float64) {
	if err := j.validateSetIndexNSegmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexNSegments",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetInputLossAction(val *string) {
	if err := j.validateSetInputLossActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputLossAction",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetInternalValue(val *AwsChannel_HlsGroupSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetIvInManifest(val *string) {
	if err := j.validateSetIvInManifestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ivInManifest",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetIvSource(val *string) {
	if err := j.validateSetIvSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ivSource",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetKeepSegments(val *float64) {
	if err := j.validateSetKeepSegmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepSegments",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetKeyFormat(val *string) {
	if err := j.validateSetKeyFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyFormat",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetKeyFormatVersions(val *string) {
	if err := j.validateSetKeyFormatVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyFormatVersions",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetManifestCompression(val *string) {
	if err := j.validateSetManifestCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestCompression",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetManifestDurationFormat(val *string) {
	if err := j.validateSetManifestDurationFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestDurationFormat",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetMinSegmentLength(val *float64) {
	if err := j.validateSetMinSegmentLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSegmentLength",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetOutputSelection(val *string) {
	if err := j.validateSetOutputSelectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputSelection",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetProgramDateTime(val *string) {
	if err := j.validateSetProgramDateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programDateTime",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetProgramDateTimeClock(val *string) {
	if err := j.validateSetProgramDateTimeClockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programDateTimeClock",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetProgramDateTimePeriod(val *float64) {
	if err := j.validateSetProgramDateTimePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programDateTimePeriod",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetRedundantManifest(val *string) {
	if err := j.validateSetRedundantManifestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redundantManifest",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetSegmentLength(val *float64) {
	if err := j.validateSetSegmentLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentLength",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetSegmentsPerSubdirectory(val *float64) {
	if err := j.validateSetSegmentsPerSubdirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentsPerSubdirectory",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetStreamInfResolution(val *string) {
	if err := j.validateSetStreamInfResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamInfResolution",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetTimedMetadataId3Frame(val *string) {
	if err := j.validateSetTimedMetadataId3FrameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timedMetadataId3Frame",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetTimedMetadataId3Period(val *float64) {
	if err := j.validateSetTimedMetadataId3PeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timedMetadataId3Period",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetTimestampDeltaMilliseconds(val *float64) {
	if err := j.validateSetTimestampDeltaMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampDeltaMilliseconds",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference)SetTsFileMode(val *string) {
	if err := j.validateSetTsFileModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tsFileMode",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) PutCaptionLanguageMappings(value interface{}) {
	if err := a.validatePutCaptionLanguageMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCaptionLanguageMappings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) PutDestination(value *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationProperty) {
	if err := a.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) PutHlsCdnSettings(value interface{}) {
	if err := a.validatePutHlsCdnSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHlsCdnSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) PutKeyProviderSettings(value *AwsChannel_KeyProviderSettingsProperty) {
	if err := a.validatePutKeyProviderSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKeyProviderSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetAdMarkers() {
	_jsii_.InvokeVoid(
		a,
		"resetAdMarkers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetBaseUrlContent() {
	_jsii_.InvokeVoid(
		a,
		"resetBaseUrlContent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetBaseUrlContent1() {
	_jsii_.InvokeVoid(
		a,
		"resetBaseUrlContent1",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetBaseUrlManifest() {
	_jsii_.InvokeVoid(
		a,
		"resetBaseUrlManifest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetBaseUrlManifest1() {
	_jsii_.InvokeVoid(
		a,
		"resetBaseUrlManifest1",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetCaptionLanguageMappings() {
	_jsii_.InvokeVoid(
		a,
		"resetCaptionLanguageMappings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetCaptionLanguageSetting() {
	_jsii_.InvokeVoid(
		a,
		"resetCaptionLanguageSetting",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetClientCache() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCache",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetCodecSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCodecSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetConstantIv() {
	_jsii_.InvokeVoid(
		a,
		"resetConstantIv",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetDirectoryStructure() {
	_jsii_.InvokeVoid(
		a,
		"resetDirectoryStructure",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetDiscontinuityTags() {
	_jsii_.InvokeVoid(
		a,
		"resetDiscontinuityTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetEncryptionType() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetHlsCdnSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetHlsCdnSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetHlsId3SegmentTagging() {
	_jsii_.InvokeVoid(
		a,
		"resetHlsId3SegmentTagging",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetIframeOnlyPlaylists() {
	_jsii_.InvokeVoid(
		a,
		"resetIframeOnlyPlaylists",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetIncompleteSegmentBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetIncompleteSegmentBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetIndexNSegments() {
	_jsii_.InvokeVoid(
		a,
		"resetIndexNSegments",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetInputLossAction() {
	_jsii_.InvokeVoid(
		a,
		"resetInputLossAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetIvInManifest() {
	_jsii_.InvokeVoid(
		a,
		"resetIvInManifest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetIvSource() {
	_jsii_.InvokeVoid(
		a,
		"resetIvSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetKeepSegments() {
	_jsii_.InvokeVoid(
		a,
		"resetKeepSegments",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetKeyFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetKeyFormatVersions() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyFormatVersions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetKeyProviderSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyProviderSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetManifestCompression() {
	_jsii_.InvokeVoid(
		a,
		"resetManifestCompression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetManifestDurationFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetManifestDurationFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetMinSegmentLength() {
	_jsii_.InvokeVoid(
		a,
		"resetMinSegmentLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		a,
		"resetMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetOutputSelection() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputSelection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetProgramDateTime() {
	_jsii_.InvokeVoid(
		a,
		"resetProgramDateTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetProgramDateTimeClock() {
	_jsii_.InvokeVoid(
		a,
		"resetProgramDateTimeClock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetProgramDateTimePeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetProgramDateTimePeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetRedundantManifest() {
	_jsii_.InvokeVoid(
		a,
		"resetRedundantManifest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetSegmentLength() {
	_jsii_.InvokeVoid(
		a,
		"resetSegmentLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetSegmentsPerSubdirectory() {
	_jsii_.InvokeVoid(
		a,
		"resetSegmentsPerSubdirectory",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetStreamInfResolution() {
	_jsii_.InvokeVoid(
		a,
		"resetStreamInfResolution",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetTimedMetadataId3Frame() {
	_jsii_.InvokeVoid(
		a,
		"resetTimedMetadataId3Frame",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetTimedMetadataId3Period() {
	_jsii_.InvokeVoid(
		a,
		"resetTimedMetadataId3Period",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetTimestampDeltaMilliseconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTimestampDeltaMilliseconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ResetTsFileMode() {
	_jsii_.InvokeVoid(
		a,
		"resetTsFileMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_HlsGroupSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

