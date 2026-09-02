package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMonitoringSchedule_EndpointInputPropertyOutputReference interface {
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
	EndpointName() *string
	// Experimental.
	SetEndpointName(val *string)
	// Experimental.
	EndpointNameInput() *string
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
	InternalValue() *TfMonitoringSchedule_EndpointInputProperty
	// Experimental.
	SetInternalValue(val *TfMonitoringSchedule_EndpointInputProperty)
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

// The jsii proxy struct for TfMonitoringSchedule_EndpointInputPropertyOutputReference
type jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) EndpointNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) EndTimeOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) EndTimeOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ExcludeFeaturesAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeFeaturesAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ExcludeFeaturesAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeFeaturesAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) FeaturesAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featuresAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) FeaturesAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featuresAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) InferenceAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) InferenceAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) InternalValue() *TfMonitoringSchedule_EndpointInputProperty {
	var returns *TfMonitoringSchedule_EndpointInputProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) LocalPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) LocalPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ProbabilityAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probabilityAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ProbabilityAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probabilityAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ProbabilityThresholdAttribute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"probabilityThresholdAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ProbabilityThresholdAttributeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"probabilityThresholdAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) S3DataDistributionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataDistributionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) S3DataDistributionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataDistributionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) S3InputMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3InputMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) S3InputModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3InputModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) StartTimeOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) StartTimeOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfMonitoringSchedule_EndpointInputPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfMonitoringSchedule_EndpointInputPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfMonitoringSchedule_EndpointInputPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfMonitoringSchedule.EndpointInputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfMonitoringSchedule_EndpointInputPropertyOutputReference_Override(t TfMonitoringSchedule_EndpointInputPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfMonitoringSchedule.EndpointInputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference)SetEndpointName(val *string) {
	if err := j.validateSetEndpointNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointName",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference)SetEndTimeOffset(val *string) {
	if err := j.validateSetEndTimeOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTimeOffset",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference)SetExcludeFeaturesAttribute(val *string) {
	if err := j.validateSetExcludeFeaturesAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeFeaturesAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference)SetFeaturesAttribute(val *string) {
	if err := j.validateSetFeaturesAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"featuresAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference)SetInferenceAttribute(val *string) {
	if err := j.validateSetInferenceAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference)SetInternalValue(val *TfMonitoringSchedule_EndpointInputProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference)SetLocalPath(val *string) {
	if err := j.validateSetLocalPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localPath",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference)SetProbabilityAttribute(val *string) {
	if err := j.validateSetProbabilityAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"probabilityAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference)SetProbabilityThresholdAttribute(val *float64) {
	if err := j.validateSetProbabilityThresholdAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"probabilityThresholdAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference)SetS3DataDistributionType(val *string) {
	if err := j.validateSetS3DataDistributionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3DataDistributionType",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference)SetS3InputMode(val *string) {
	if err := j.validateSetS3InputModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3InputMode",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference)SetStartTimeOffset(val *string) {
	if err := j.validateSetStartTimeOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTimeOffset",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ResetEndTimeOffset() {
	_jsii_.InvokeVoid(
		t,
		"resetEndTimeOffset",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ResetExcludeFeaturesAttribute() {
	_jsii_.InvokeVoid(
		t,
		"resetExcludeFeaturesAttribute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ResetFeaturesAttribute() {
	_jsii_.InvokeVoid(
		t,
		"resetFeaturesAttribute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ResetInferenceAttribute() {
	_jsii_.InvokeVoid(
		t,
		"resetInferenceAttribute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ResetProbabilityAttribute() {
	_jsii_.InvokeVoid(
		t,
		"resetProbabilityAttribute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ResetProbabilityThresholdAttribute() {
	_jsii_.InvokeVoid(
		t,
		"resetProbabilityThresholdAttribute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ResetS3DataDistributionType() {
	_jsii_.InvokeVoid(
		t,
		"resetS3DataDistributionType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ResetS3InputMode() {
	_jsii_.InvokeVoid(
		t,
		"resetS3InputMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ResetStartTimeOffset() {
	_jsii_.InvokeVoid(
		t,
		"resetStartTimeOffset",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfMonitoringSchedule_EndpointInputPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

