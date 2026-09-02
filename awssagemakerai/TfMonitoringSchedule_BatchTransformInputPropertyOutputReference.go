package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMonitoringSchedule_BatchTransformInputPropertyOutputReference interface {
	cdktn.ComplexObject
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
	DataCapturedDestinationS3Uri() *string
	// Experimental.
	SetDataCapturedDestinationS3Uri(val *string)
	// Experimental.
	DataCapturedDestinationS3UriInput() *string
	// Experimental.
	DatasetFormat() TfMonitoringSchedule_DatasetFormatPropertyOutputReference
	// Experimental.
	DatasetFormatInput() *TfMonitoringSchedule_DatasetFormatProperty
	// Experimental.
	EndTimeOffset() *string
	// Experimental.
	SetEndTimeOffset(val *string)
	// Experimental.
	EndTimeOffsetInput() *string
	// Experimental.
	ExcludeFeaturesAttribute() *string
	// Experimental.
	SetExcludeFeaturesAttribute(val *string)
	// Experimental.
	ExcludeFeaturesAttributeInput() *string
	// Experimental.
	FeaturesAttribute() *string
	// Experimental.
	SetFeaturesAttribute(val *string)
	// Experimental.
	FeaturesAttributeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InferenceAttribute() *string
	// Experimental.
	SetInferenceAttribute(val *string)
	// Experimental.
	InferenceAttributeInput() *string
	// Experimental.
	InternalValue() *TfMonitoringSchedule_BatchTransformInputProperty
	// Experimental.
	SetInternalValue(val *TfMonitoringSchedule_BatchTransformInputProperty)
	// Experimental.
	LocalPath() *string
	// Experimental.
	SetLocalPath(val *string)
	// Experimental.
	LocalPathInput() *string
	// Experimental.
	ProbabilityAttribute() *string
	// Experimental.
	SetProbabilityAttribute(val *string)
	// Experimental.
	ProbabilityAttributeInput() *string
	// Experimental.
	ProbabilityThresholdAttribute() *float64
	// Experimental.
	SetProbabilityThresholdAttribute(val *float64)
	// Experimental.
	ProbabilityThresholdAttributeInput() *float64
	// Experimental.
	S3DataDistributionType() *string
	// Experimental.
	SetS3DataDistributionType(val *string)
	// Experimental.
	S3DataDistributionTypeInput() *string
	// Experimental.
	S3InputMode() *string
	// Experimental.
	SetS3InputMode(val *string)
	// Experimental.
	S3InputModeInput() *string
	// Experimental.
	StartTimeOffset() *string
	// Experimental.
	SetStartTimeOffset(val *string)
	// Experimental.
	StartTimeOffsetInput() *string
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
	PutDatasetFormat(value *TfMonitoringSchedule_DatasetFormatProperty)
	// Experimental.
	ResetEndTimeOffset()
	// Experimental.
	ResetExcludeFeaturesAttribute()
	// Experimental.
	ResetFeaturesAttribute()
	// Experimental.
	ResetInferenceAttribute()
	// Experimental.
	ResetProbabilityAttribute()
	// Experimental.
	ResetProbabilityThresholdAttribute()
	// Experimental.
	ResetS3DataDistributionType()
	// Experimental.
	ResetS3InputMode()
	// Experimental.
	ResetStartTimeOffset()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfMonitoringSchedule_BatchTransformInputPropertyOutputReference
type jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) DataCapturedDestinationS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCapturedDestinationS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) DataCapturedDestinationS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCapturedDestinationS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) DatasetFormat() TfMonitoringSchedule_DatasetFormatPropertyOutputReference {
	var returns TfMonitoringSchedule_DatasetFormatPropertyOutputReference
	_jsii_.Get(
		j,
		"datasetFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) DatasetFormatInput() *TfMonitoringSchedule_DatasetFormatProperty {
	var returns *TfMonitoringSchedule_DatasetFormatProperty
	_jsii_.Get(
		j,
		"datasetFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) EndTimeOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) EndTimeOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ExcludeFeaturesAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeFeaturesAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ExcludeFeaturesAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeFeaturesAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) FeaturesAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featuresAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) FeaturesAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featuresAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) InferenceAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) InferenceAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) InternalValue() *TfMonitoringSchedule_BatchTransformInputProperty {
	var returns *TfMonitoringSchedule_BatchTransformInputProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) LocalPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) LocalPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ProbabilityAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probabilityAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ProbabilityAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probabilityAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ProbabilityThresholdAttribute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"probabilityThresholdAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ProbabilityThresholdAttributeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"probabilityThresholdAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) S3DataDistributionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataDistributionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) S3DataDistributionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataDistributionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) S3InputMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3InputMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) S3InputModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3InputModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) StartTimeOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) StartTimeOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfMonitoringSchedule_BatchTransformInputPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfMonitoringSchedule_BatchTransformInputPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfMonitoringSchedule_BatchTransformInputPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfMonitoringSchedule.BatchTransformInputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfMonitoringSchedule_BatchTransformInputPropertyOutputReference_Override(t TfMonitoringSchedule_BatchTransformInputPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfMonitoringSchedule.BatchTransformInputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference)SetDataCapturedDestinationS3Uri(val *string) {
	if err := j.validateSetDataCapturedDestinationS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCapturedDestinationS3Uri",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference)SetEndTimeOffset(val *string) {
	if err := j.validateSetEndTimeOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTimeOffset",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference)SetExcludeFeaturesAttribute(val *string) {
	if err := j.validateSetExcludeFeaturesAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeFeaturesAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference)SetFeaturesAttribute(val *string) {
	if err := j.validateSetFeaturesAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"featuresAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference)SetInferenceAttribute(val *string) {
	if err := j.validateSetInferenceAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference)SetInternalValue(val *TfMonitoringSchedule_BatchTransformInputProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference)SetLocalPath(val *string) {
	if err := j.validateSetLocalPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localPath",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference)SetProbabilityAttribute(val *string) {
	if err := j.validateSetProbabilityAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"probabilityAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference)SetProbabilityThresholdAttribute(val *float64) {
	if err := j.validateSetProbabilityThresholdAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"probabilityThresholdAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference)SetS3DataDistributionType(val *string) {
	if err := j.validateSetS3DataDistributionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3DataDistributionType",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference)SetS3InputMode(val *string) {
	if err := j.validateSetS3InputModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3InputMode",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference)SetStartTimeOffset(val *string) {
	if err := j.validateSetStartTimeOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTimeOffset",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) PutDatasetFormat(value *TfMonitoringSchedule_DatasetFormatProperty) {
	if err := t.validatePutDatasetFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDatasetFormat",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ResetEndTimeOffset() {
	_jsii_.InvokeVoid(
		t,
		"resetEndTimeOffset",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ResetExcludeFeaturesAttribute() {
	_jsii_.InvokeVoid(
		t,
		"resetExcludeFeaturesAttribute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ResetFeaturesAttribute() {
	_jsii_.InvokeVoid(
		t,
		"resetFeaturesAttribute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ResetInferenceAttribute() {
	_jsii_.InvokeVoid(
		t,
		"resetInferenceAttribute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ResetProbabilityAttribute() {
	_jsii_.InvokeVoid(
		t,
		"resetProbabilityAttribute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ResetProbabilityThresholdAttribute() {
	_jsii_.InvokeVoid(
		t,
		"resetProbabilityThresholdAttribute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ResetS3DataDistributionType() {
	_jsii_.InvokeVoid(
		t,
		"resetS3DataDistributionType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ResetS3InputMode() {
	_jsii_.InvokeVoid(
		t,
		"resetS3InputMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ResetStartTimeOffset() {
	_jsii_.InvokeVoid(
		t,
		"resetStartTimeOffset",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfMonitoringSchedule_BatchTransformInputPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

