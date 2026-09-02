package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InstanceCount() *float64
	// Experimental.
	SetInstanceCount(val *float64)
	// Experimental.
	InstanceCountInput() *float64
	// Experimental.
	InstanceGroups() TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigInstanceGroupsPropertyList
	// Experimental.
	InstanceGroupsInput() interface{}
	// Experimental.
	InstancePlacementConfig() TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigInstancePlacementConfigPropertyList
	// Experimental.
	InstancePlacementConfigInput() interface{}
	// Experimental.
	InstanceType() *string
	// Experimental.
	SetInstanceType(val *string)
	// Experimental.
	InstanceTypeInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	KeepAlivePeriodInSeconds() *float64
	// Experimental.
	SetKeepAlivePeriodInSeconds(val *float64)
	// Experimental.
	KeepAlivePeriodInSecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TrainingPlanArn() *string
	// Experimental.
	SetTrainingPlanArn(val *string)
	// Experimental.
	TrainingPlanArnInput() *string
	// Experimental.
	VolumeKmsKeyId() *string
	// Experimental.
	SetVolumeKmsKeyId(val *string)
	// Experimental.
	VolumeKmsKeyIdInput() *string
	// Experimental.
	VolumeSizeInGb() *float64
	// Experimental.
	SetVolumeSizeInGb(val *float64)
	// Experimental.
	VolumeSizeInGbInput() *float64
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
	PutInstanceGroups(value interface{})
	// Experimental.
	PutInstancePlacementConfig(value interface{})
	// Experimental.
	ResetInstanceCount()
	// Experimental.
	ResetInstanceGroups()
	// Experimental.
	ResetInstancePlacementConfig()
	// Experimental.
	ResetInstanceType()
	// Experimental.
	ResetKeepAlivePeriodInSeconds()
	// Experimental.
	ResetTrainingPlanArn()
	// Experimental.
	ResetVolumeKmsKeyId()
	// Experimental.
	ResetVolumeSizeInGb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference
type jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) InstanceGroups() TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigInstanceGroupsPropertyList {
	var returns TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigInstanceGroupsPropertyList
	_jsii_.Get(
		j,
		"instanceGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) InstanceGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) InstancePlacementConfig() TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigInstancePlacementConfigPropertyList {
	var returns TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigInstancePlacementConfigPropertyList
	_jsii_.Get(
		j,
		"instancePlacementConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) InstancePlacementConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instancePlacementConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) KeepAlivePeriodInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAlivePeriodInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) KeepAlivePeriodInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAlivePeriodInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) TrainingPlanArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingPlanArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) TrainingPlanArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingPlanArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) VolumeKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) VolumeKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) VolumeSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) VolumeSizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInGbInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfHyperParameterTuningJob.TrainingJobDefinitionsResourceConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference_Override(t TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfHyperParameterTuningJob.TrainingJobDefinitionsResourceConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference)SetInstanceCount(val *float64) {
	if err := j.validateSetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference)SetKeepAlivePeriodInSeconds(val *float64) {
	if err := j.validateSetKeepAlivePeriodInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepAlivePeriodInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference)SetTrainingPlanArn(val *string) {
	if err := j.validateSetTrainingPlanArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingPlanArn",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference)SetVolumeKmsKeyId(val *string) {
	if err := j.validateSetVolumeKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference)SetVolumeSizeInGb(val *float64) {
	if err := j.validateSetVolumeSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSizeInGb",
		val,
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) PutInstanceGroups(value interface{}) {
	if err := t.validatePutInstanceGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInstanceGroups",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) PutInstancePlacementConfig(value interface{}) {
	if err := t.validatePutInstancePlacementConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInstancePlacementConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) ResetInstanceCount() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) ResetInstanceGroups() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceGroups",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) ResetInstancePlacementConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetInstancePlacementConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) ResetKeepAlivePeriodInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetKeepAlivePeriodInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) ResetTrainingPlanArn() {
	_jsii_.InvokeVoid(
		t,
		"resetTrainingPlanArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) ResetVolumeKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetVolumeKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) ResetVolumeSizeInGb() {
	_jsii_.InvokeVoid(
		t,
		"resetVolumeSizeInGb",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

