package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AbsentInputAudioBehavior() *string
	// Experimental.
	SetAbsentInputAudioBehavior(val *string)
	// Experimental.
	AbsentInputAudioBehaviorInput() *string
	// Experimental.
	Arib() *string
	// Experimental.
	SetArib(val *string)
	// Experimental.
	AribCaptionsPid() *string
	// Experimental.
	SetAribCaptionsPid(val *string)
	// Experimental.
	AribCaptionsPidControl() *string
	// Experimental.
	SetAribCaptionsPidControl(val *string)
	// Experimental.
	AribCaptionsPidControlInput() *string
	// Experimental.
	AribCaptionsPidInput() *string
	// Experimental.
	AribInput() *string
	// Experimental.
	AudioBufferModel() *string
	// Experimental.
	SetAudioBufferModel(val *string)
	// Experimental.
	AudioBufferModelInput() *string
	// Experimental.
	AudioFramesPerPes() *float64
	// Experimental.
	SetAudioFramesPerPes(val *float64)
	// Experimental.
	AudioFramesPerPesInput() *float64
	// Experimental.
	AudioPids() *string
	// Experimental.
	SetAudioPids(val *string)
	// Experimental.
	AudioPidsInput() *string
	// Experimental.
	AudioStreamType() *string
	// Experimental.
	SetAudioStreamType(val *string)
	// Experimental.
	AudioStreamTypeInput() *string
	// Experimental.
	Bitrate() *float64
	// Experimental.
	SetBitrate(val *float64)
	// Experimental.
	BitrateInput() *float64
	// Experimental.
	BufferModel() *string
	// Experimental.
	SetBufferModel(val *string)
	// Experimental.
	BufferModelInput() *string
	// Experimental.
	CcDescriptor() *string
	// Experimental.
	SetCcDescriptor(val *string)
	// Experimental.
	CcDescriptorInput() *string
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
	DvbNitSettings() TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbNitSettingsPropertyOutputReference
	// Experimental.
	DvbNitSettingsInput() *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbNitSettingsProperty
	// Experimental.
	DvbSdtSettings() TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbSdtSettingsPropertyOutputReference
	// Experimental.
	DvbSdtSettingsInput() *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbSdtSettingsProperty
	// Experimental.
	DvbSubPids() *string
	// Experimental.
	SetDvbSubPids(val *string)
	// Experimental.
	DvbSubPidsInput() *string
	// Experimental.
	DvbTdtSettings() TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbTdtSettingsPropertyOutputReference
	// Experimental.
	DvbTdtSettingsInput() *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbTdtSettingsProperty
	// Experimental.
	DvbTeletextPid() *string
	// Experimental.
	SetDvbTeletextPid(val *string)
	// Experimental.
	DvbTeletextPidInput() *string
	// Experimental.
	Ebif() *string
	// Experimental.
	SetEbif(val *string)
	// Experimental.
	EbifInput() *string
	// Experimental.
	EbpAudioInterval() *string
	// Experimental.
	SetEbpAudioInterval(val *string)
	// Experimental.
	EbpAudioIntervalInput() *string
	// Experimental.
	EbpLookaheadMs() *float64
	// Experimental.
	SetEbpLookaheadMs(val *float64)
	// Experimental.
	EbpLookaheadMsInput() *float64
	// Experimental.
	EbpPlacement() *string
	// Experimental.
	SetEbpPlacement(val *string)
	// Experimental.
	EbpPlacementInput() *string
	// Experimental.
	EcmPid() *string
	// Experimental.
	SetEcmPid(val *string)
	// Experimental.
	EcmPidInput() *string
	// Experimental.
	EsRateInPes() *string
	// Experimental.
	SetEsRateInPes(val *string)
	// Experimental.
	EsRateInPesInput() *string
	// Experimental.
	EtvPlatformPid() *string
	// Experimental.
	SetEtvPlatformPid(val *string)
	// Experimental.
	EtvPlatformPidInput() *string
	// Experimental.
	EtvSignalPid() *string
	// Experimental.
	SetEtvSignalPid(val *string)
	// Experimental.
	EtvSignalPidInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FragmentTime() *float64
	// Experimental.
	SetFragmentTime(val *float64)
	// Experimental.
	FragmentTimeInput() *float64
	// Experimental.
	InternalValue() *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsProperty)
	// Experimental.
	Klv() *string
	// Experimental.
	SetKlv(val *string)
	// Experimental.
	KlvDataPids() *string
	// Experimental.
	SetKlvDataPids(val *string)
	// Experimental.
	KlvDataPidsInput() *string
	// Experimental.
	KlvInput() *string
	// Experimental.
	NielsenId3Behavior() *string
	// Experimental.
	SetNielsenId3Behavior(val *string)
	// Experimental.
	NielsenId3BehaviorInput() *string
	// Experimental.
	NullPacketBitrate() *float64
	// Experimental.
	SetNullPacketBitrate(val *float64)
	// Experimental.
	NullPacketBitrateInput() *float64
	// Experimental.
	PatInterval() *float64
	// Experimental.
	SetPatInterval(val *float64)
	// Experimental.
	PatIntervalInput() *float64
	// Experimental.
	PcrControl() *string
	// Experimental.
	SetPcrControl(val *string)
	// Experimental.
	PcrControlInput() *string
	// Experimental.
	PcrPeriod() *float64
	// Experimental.
	SetPcrPeriod(val *float64)
	// Experimental.
	PcrPeriodInput() *float64
	// Experimental.
	PcrPid() *string
	// Experimental.
	SetPcrPid(val *string)
	// Experimental.
	PcrPidInput() *string
	// Experimental.
	PmtInterval() *float64
	// Experimental.
	SetPmtInterval(val *float64)
	// Experimental.
	PmtIntervalInput() *float64
	// Experimental.
	PmtPid() *string
	// Experimental.
	SetPmtPid(val *string)
	// Experimental.
	PmtPidInput() *string
	// Experimental.
	ProgramNum() *float64
	// Experimental.
	SetProgramNum(val *float64)
	// Experimental.
	ProgramNumInput() *float64
	// Experimental.
	RateMode() *string
	// Experimental.
	SetRateMode(val *string)
	// Experimental.
	RateModeInput() *string
	// Experimental.
	Scte27Pids() *string
	// Experimental.
	SetScte27Pids(val *string)
	// Experimental.
	Scte27PidsInput() *string
	// Experimental.
	Scte35Control() *string
	// Experimental.
	SetScte35Control(val *string)
	// Experimental.
	Scte35ControlInput() *string
	// Experimental.
	Scte35Pid() *string
	// Experimental.
	SetScte35Pid(val *string)
	// Experimental.
	Scte35PidInput() *string
	// Experimental.
	SegmentationMarkers() *string
	// Experimental.
	SetSegmentationMarkers(val *string)
	// Experimental.
	SegmentationMarkersInput() *string
	// Experimental.
	SegmentationStyle() *string
	// Experimental.
	SetSegmentationStyle(val *string)
	// Experimental.
	SegmentationStyleInput() *string
	// Experimental.
	SegmentationTime() *float64
	// Experimental.
	SetSegmentationTime(val *float64)
	// Experimental.
	SegmentationTimeInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimedMetadataBehavior() *string
	// Experimental.
	SetTimedMetadataBehavior(val *string)
	// Experimental.
	TimedMetadataBehaviorInput() *string
	// Experimental.
	TimedMetadataPid() *string
	// Experimental.
	SetTimedMetadataPid(val *string)
	// Experimental.
	TimedMetadataPidInput() *string
	// Experimental.
	TransportStreamId() *float64
	// Experimental.
	SetTransportStreamId(val *float64)
	// Experimental.
	TransportStreamIdInput() *float64
	// Experimental.
	VideoPid() *string
	// Experimental.
	SetVideoPid(val *string)
	// Experimental.
	VideoPidInput() *string
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
	PutDvbNitSettings(value *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbNitSettingsProperty)
	// Experimental.
	PutDvbSdtSettings(value *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbSdtSettingsProperty)
	// Experimental.
	PutDvbTdtSettings(value *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbTdtSettingsProperty)
	// Experimental.
	ResetAbsentInputAudioBehavior()
	// Experimental.
	ResetArib()
	// Experimental.
	ResetAribCaptionsPid()
	// Experimental.
	ResetAribCaptionsPidControl()
	// Experimental.
	ResetAudioBufferModel()
	// Experimental.
	ResetAudioFramesPerPes()
	// Experimental.
	ResetAudioPids()
	// Experimental.
	ResetAudioStreamType()
	// Experimental.
	ResetBitrate()
	// Experimental.
	ResetBufferModel()
	// Experimental.
	ResetCcDescriptor()
	// Experimental.
	ResetDvbNitSettings()
	// Experimental.
	ResetDvbSdtSettings()
	// Experimental.
	ResetDvbSubPids()
	// Experimental.
	ResetDvbTdtSettings()
	// Experimental.
	ResetDvbTeletextPid()
	// Experimental.
	ResetEbif()
	// Experimental.
	ResetEbpAudioInterval()
	// Experimental.
	ResetEbpLookaheadMs()
	// Experimental.
	ResetEbpPlacement()
	// Experimental.
	ResetEcmPid()
	// Experimental.
	ResetEsRateInPes()
	// Experimental.
	ResetEtvPlatformPid()
	// Experimental.
	ResetEtvSignalPid()
	// Experimental.
	ResetFragmentTime()
	// Experimental.
	ResetKlv()
	// Experimental.
	ResetKlvDataPids()
	// Experimental.
	ResetNielsenId3Behavior()
	// Experimental.
	ResetNullPacketBitrate()
	// Experimental.
	ResetPatInterval()
	// Experimental.
	ResetPcrControl()
	// Experimental.
	ResetPcrPeriod()
	// Experimental.
	ResetPcrPid()
	// Experimental.
	ResetPmtInterval()
	// Experimental.
	ResetPmtPid()
	// Experimental.
	ResetProgramNum()
	// Experimental.
	ResetRateMode()
	// Experimental.
	ResetScte27Pids()
	// Experimental.
	ResetScte35Control()
	// Experimental.
	ResetScte35Pid()
	// Experimental.
	ResetSegmentationMarkers()
	// Experimental.
	ResetSegmentationStyle()
	// Experimental.
	ResetSegmentationTime()
	// Experimental.
	ResetTimedMetadataBehavior()
	// Experimental.
	ResetTimedMetadataPid()
	// Experimental.
	ResetTransportStreamId()
	// Experimental.
	ResetVideoPid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference
type jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) AbsentInputAudioBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absentInputAudioBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) AbsentInputAudioBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absentInputAudioBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) Arib() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) AribCaptionsPid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aribCaptionsPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) AribCaptionsPidControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aribCaptionsPidControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) AribCaptionsPidControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aribCaptionsPidControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) AribCaptionsPidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aribCaptionsPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) AribInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aribInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) AudioBufferModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioBufferModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) AudioBufferModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioBufferModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) AudioFramesPerPes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"audioFramesPerPes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) AudioFramesPerPesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"audioFramesPerPesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) AudioPids() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioPids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) AudioPidsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioPidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) AudioStreamType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioStreamType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) AudioStreamTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioStreamTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) Bitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) BitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) BufferModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bufferModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) BufferModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bufferModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) CcDescriptor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ccDescriptor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) CcDescriptorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ccDescriptorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) DvbNitSettings() TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbNitSettingsPropertyOutputReference {
	var returns TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbNitSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"dvbNitSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) DvbNitSettingsInput() *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbNitSettingsProperty {
	var returns *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbNitSettingsProperty
	_jsii_.Get(
		j,
		"dvbNitSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) DvbSdtSettings() TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbSdtSettingsPropertyOutputReference {
	var returns TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbSdtSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"dvbSdtSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) DvbSdtSettingsInput() *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbSdtSettingsProperty {
	var returns *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbSdtSettingsProperty
	_jsii_.Get(
		j,
		"dvbSdtSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) DvbSubPids() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dvbSubPids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) DvbSubPidsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dvbSubPidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) DvbTdtSettings() TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbTdtSettingsPropertyOutputReference {
	var returns TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbTdtSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"dvbTdtSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) DvbTdtSettingsInput() *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbTdtSettingsProperty {
	var returns *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbTdtSettingsProperty
	_jsii_.Get(
		j,
		"dvbTdtSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) DvbTeletextPid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dvbTeletextPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) DvbTeletextPidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dvbTeletextPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) Ebif() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebif",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) EbifInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebifInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) EbpAudioInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebpAudioInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) EbpAudioIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebpAudioIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) EbpLookaheadMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebpLookaheadMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) EbpLookaheadMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebpLookaheadMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) EbpPlacement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebpPlacement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) EbpPlacementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebpPlacementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) EcmPid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecmPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) EcmPidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecmPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) EsRateInPes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"esRateInPes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) EsRateInPesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"esRateInPesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) EtvPlatformPid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etvPlatformPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) EtvPlatformPidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etvPlatformPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) EtvSignalPid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etvSignalPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) EtvSignalPidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etvSignalPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) FragmentTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fragmentTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) FragmentTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fragmentTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) InternalValue() *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsProperty {
	var returns *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) Klv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"klv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) KlvDataPids() *string {
	var returns *string
	_jsii_.Get(
		j,
		"klvDataPids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) KlvDataPidsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"klvDataPidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) KlvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"klvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) NielsenId3Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nielsenId3Behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) NielsenId3BehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nielsenId3BehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) NullPacketBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nullPacketBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) NullPacketBitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nullPacketBitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) PatInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"patInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) PatIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"patIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) PcrControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pcrControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) PcrControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pcrControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) PcrPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pcrPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) PcrPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pcrPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) PcrPid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pcrPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) PcrPidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pcrPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) PmtInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pmtInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) PmtIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pmtIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) PmtPid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pmtPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) PmtPidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pmtPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ProgramNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ProgramNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) RateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) RateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) Scte27Pids() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scte27Pids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) Scte27PidsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scte27PidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) Scte35Control() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scte35Control",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) Scte35ControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scte35ControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) Scte35Pid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scte35Pid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) Scte35PidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scte35PidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) SegmentationMarkers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentationMarkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) SegmentationMarkersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentationMarkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) SegmentationStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentationStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) SegmentationStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentationStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) SegmentationTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) SegmentationTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) TimedMetadataBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timedMetadataBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) TimedMetadataBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timedMetadataBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) TimedMetadataPid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timedMetadataPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) TimedMetadataPidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timedMetadataPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) TransportStreamId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) TransportStreamIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) VideoPid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) VideoPidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoPidInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference_Override(t TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetAbsentInputAudioBehavior(val *string) {
	if err := j.validateSetAbsentInputAudioBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"absentInputAudioBehavior",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetArib(val *string) {
	if err := j.validateSetAribParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arib",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetAribCaptionsPid(val *string) {
	if err := j.validateSetAribCaptionsPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aribCaptionsPid",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetAribCaptionsPidControl(val *string) {
	if err := j.validateSetAribCaptionsPidControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aribCaptionsPidControl",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetAudioBufferModel(val *string) {
	if err := j.validateSetAudioBufferModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioBufferModel",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetAudioFramesPerPes(val *float64) {
	if err := j.validateSetAudioFramesPerPesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioFramesPerPes",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetAudioPids(val *string) {
	if err := j.validateSetAudioPidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioPids",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetAudioStreamType(val *string) {
	if err := j.validateSetAudioStreamTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioStreamType",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetBitrate(val *float64) {
	if err := j.validateSetBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitrate",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetBufferModel(val *string) {
	if err := j.validateSetBufferModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferModel",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetCcDescriptor(val *string) {
	if err := j.validateSetCcDescriptorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ccDescriptor",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetDvbSubPids(val *string) {
	if err := j.validateSetDvbSubPidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dvbSubPids",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetDvbTeletextPid(val *string) {
	if err := j.validateSetDvbTeletextPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dvbTeletextPid",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetEbif(val *string) {
	if err := j.validateSetEbifParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebif",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetEbpAudioInterval(val *string) {
	if err := j.validateSetEbpAudioIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebpAudioInterval",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetEbpLookaheadMs(val *float64) {
	if err := j.validateSetEbpLookaheadMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebpLookaheadMs",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetEbpPlacement(val *string) {
	if err := j.validateSetEbpPlacementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebpPlacement",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetEcmPid(val *string) {
	if err := j.validateSetEcmPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecmPid",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetEsRateInPes(val *string) {
	if err := j.validateSetEsRateInPesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"esRateInPes",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetEtvPlatformPid(val *string) {
	if err := j.validateSetEtvPlatformPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etvPlatformPid",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetEtvSignalPid(val *string) {
	if err := j.validateSetEtvSignalPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etvSignalPid",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetFragmentTime(val *float64) {
	if err := j.validateSetFragmentTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fragmentTime",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetKlv(val *string) {
	if err := j.validateSetKlvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"klv",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetKlvDataPids(val *string) {
	if err := j.validateSetKlvDataPidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"klvDataPids",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetNielsenId3Behavior(val *string) {
	if err := j.validateSetNielsenId3BehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nielsenId3Behavior",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetNullPacketBitrate(val *float64) {
	if err := j.validateSetNullPacketBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullPacketBitrate",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetPatInterval(val *float64) {
	if err := j.validateSetPatIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patInterval",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetPcrControl(val *string) {
	if err := j.validateSetPcrControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pcrControl",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetPcrPeriod(val *float64) {
	if err := j.validateSetPcrPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pcrPeriod",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetPcrPid(val *string) {
	if err := j.validateSetPcrPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pcrPid",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetPmtInterval(val *float64) {
	if err := j.validateSetPmtIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pmtInterval",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetPmtPid(val *string) {
	if err := j.validateSetPmtPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pmtPid",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetProgramNum(val *float64) {
	if err := j.validateSetProgramNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programNum",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetRateMode(val *string) {
	if err := j.validateSetRateModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateMode",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetScte27Pids(val *string) {
	if err := j.validateSetScte27PidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scte27Pids",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetScte35Control(val *string) {
	if err := j.validateSetScte35ControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scte35Control",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetScte35Pid(val *string) {
	if err := j.validateSetScte35PidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scte35Pid",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetSegmentationMarkers(val *string) {
	if err := j.validateSetSegmentationMarkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentationMarkers",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetSegmentationStyle(val *string) {
	if err := j.validateSetSegmentationStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentationStyle",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetSegmentationTime(val *float64) {
	if err := j.validateSetSegmentationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentationTime",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetTimedMetadataBehavior(val *string) {
	if err := j.validateSetTimedMetadataBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timedMetadataBehavior",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetTimedMetadataPid(val *string) {
	if err := j.validateSetTimedMetadataPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timedMetadataPid",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetTransportStreamId(val *float64) {
	if err := j.validateSetTransportStreamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportStreamId",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference)SetVideoPid(val *string) {
	if err := j.validateSetVideoPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"videoPid",
		val,
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) PutDvbNitSettings(value *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbNitSettingsProperty) {
	if err := t.validatePutDvbNitSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDvbNitSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) PutDvbSdtSettings(value *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbSdtSettingsProperty) {
	if err := t.validatePutDvbSdtSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDvbSdtSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) PutDvbTdtSettings(value *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsDvbTdtSettingsProperty) {
	if err := t.validatePutDvbTdtSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDvbTdtSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetAbsentInputAudioBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetAbsentInputAudioBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetArib() {
	_jsii_.InvokeVoid(
		t,
		"resetArib",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetAribCaptionsPid() {
	_jsii_.InvokeVoid(
		t,
		"resetAribCaptionsPid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetAribCaptionsPidControl() {
	_jsii_.InvokeVoid(
		t,
		"resetAribCaptionsPidControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetAudioBufferModel() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioBufferModel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetAudioFramesPerPes() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioFramesPerPes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetAudioPids() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioPids",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetAudioStreamType() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioStreamType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetBitrate() {
	_jsii_.InvokeVoid(
		t,
		"resetBitrate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetBufferModel() {
	_jsii_.InvokeVoid(
		t,
		"resetBufferModel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetCcDescriptor() {
	_jsii_.InvokeVoid(
		t,
		"resetCcDescriptor",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetDvbNitSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetDvbNitSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetDvbSdtSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetDvbSdtSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetDvbSubPids() {
	_jsii_.InvokeVoid(
		t,
		"resetDvbSubPids",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetDvbTdtSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetDvbTdtSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetDvbTeletextPid() {
	_jsii_.InvokeVoid(
		t,
		"resetDvbTeletextPid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetEbif() {
	_jsii_.InvokeVoid(
		t,
		"resetEbif",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetEbpAudioInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetEbpAudioInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetEbpLookaheadMs() {
	_jsii_.InvokeVoid(
		t,
		"resetEbpLookaheadMs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetEbpPlacement() {
	_jsii_.InvokeVoid(
		t,
		"resetEbpPlacement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetEcmPid() {
	_jsii_.InvokeVoid(
		t,
		"resetEcmPid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetEsRateInPes() {
	_jsii_.InvokeVoid(
		t,
		"resetEsRateInPes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetEtvPlatformPid() {
	_jsii_.InvokeVoid(
		t,
		"resetEtvPlatformPid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetEtvSignalPid() {
	_jsii_.InvokeVoid(
		t,
		"resetEtvSignalPid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetFragmentTime() {
	_jsii_.InvokeVoid(
		t,
		"resetFragmentTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetKlv() {
	_jsii_.InvokeVoid(
		t,
		"resetKlv",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetKlvDataPids() {
	_jsii_.InvokeVoid(
		t,
		"resetKlvDataPids",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetNielsenId3Behavior() {
	_jsii_.InvokeVoid(
		t,
		"resetNielsenId3Behavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetNullPacketBitrate() {
	_jsii_.InvokeVoid(
		t,
		"resetNullPacketBitrate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetPatInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetPatInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetPcrControl() {
	_jsii_.InvokeVoid(
		t,
		"resetPcrControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetPcrPeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetPcrPeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetPcrPid() {
	_jsii_.InvokeVoid(
		t,
		"resetPcrPid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetPmtInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetPmtInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetPmtPid() {
	_jsii_.InvokeVoid(
		t,
		"resetPmtPid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetProgramNum() {
	_jsii_.InvokeVoid(
		t,
		"resetProgramNum",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetRateMode() {
	_jsii_.InvokeVoid(
		t,
		"resetRateMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetScte27Pids() {
	_jsii_.InvokeVoid(
		t,
		"resetScte27Pids",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetScte35Control() {
	_jsii_.InvokeVoid(
		t,
		"resetScte35Control",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetScte35Pid() {
	_jsii_.InvokeVoid(
		t,
		"resetScte35Pid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetSegmentationMarkers() {
	_jsii_.InvokeVoid(
		t,
		"resetSegmentationMarkers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetSegmentationStyle() {
	_jsii_.InvokeVoid(
		t,
		"resetSegmentationStyle",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetSegmentationTime() {
	_jsii_.InvokeVoid(
		t,
		"resetSegmentationTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetTimedMetadataBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetTimedMetadataBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetTimedMetadataPid() {
	_jsii_.InvokeVoid(
		t,
		"resetTimedMetadataPid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetTransportStreamId() {
	_jsii_.InvokeVoid(
		t,
		"resetTransportStreamId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ResetVideoPid() {
	_jsii_.InvokeVoid(
		t,
		"resetVideoPid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

