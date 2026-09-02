package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAlgorithm_TrainingSpecificationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdditionalS3DataSource() TfAlgorithm_TrainingSpecificationAdditionalS3DataSourcePropertyList
	// Experimental.
	AdditionalS3DataSourceInput() interface{}
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MetricDefinitions() TfAlgorithm_MetricDefinitionsPropertyList
	// Experimental.
	MetricDefinitionsInput() interface{}
	// Experimental.
	SupportedHyperParameters() TfAlgorithm_SupportedHyperParametersPropertyList
	// Experimental.
	SupportedHyperParametersInput() interface{}
	// Experimental.
	SupportedTrainingInstanceTypes() *[]*string
	// Experimental.
	SetSupportedTrainingInstanceTypes(val *[]*string)
	// Experimental.
	SupportedTrainingInstanceTypesInput() *[]*string
	// Experimental.
	SupportedTuningJobObjectiveMetrics() TfAlgorithm_SupportedTuningJobObjectiveMetricsPropertyList
	// Experimental.
	SupportedTuningJobObjectiveMetricsInput() interface{}
	// Experimental.
	SupportsDistributedTraining() interface{}
	// Experimental.
	SetSupportsDistributedTraining(val interface{})
	// Experimental.
	SupportsDistributedTrainingInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TrainingChannels() TfAlgorithm_TrainingChannelsPropertyList
	// Experimental.
	TrainingChannelsInput() interface{}
	// Experimental.
	TrainingImage() *string
	// Experimental.
	SetTrainingImage(val *string)
	// Experimental.
	TrainingImageDigest() *string
	// Experimental.
	SetTrainingImageDigest(val *string)
	// Experimental.
	TrainingImageDigestInput() *string
	// Experimental.
	TrainingImageInput() *string
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
	PutAdditionalS3DataSource(value interface{})
	// Experimental.
	PutMetricDefinitions(value interface{})
	// Experimental.
	PutSupportedHyperParameters(value interface{})
	// Experimental.
	PutSupportedTuningJobObjectiveMetrics(value interface{})
	// Experimental.
	PutTrainingChannels(value interface{})
	// Experimental.
	ResetAdditionalS3DataSource()
	// Experimental.
	ResetMetricDefinitions()
	// Experimental.
	ResetSupportedHyperParameters()
	// Experimental.
	ResetSupportedTuningJobObjectiveMetrics()
	// Experimental.
	ResetSupportsDistributedTraining()
	// Experimental.
	ResetTrainingChannels()
	// Experimental.
	ResetTrainingImageDigest()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfAlgorithm_TrainingSpecificationPropertyOutputReference
type jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) AdditionalS3DataSource() TfAlgorithm_TrainingSpecificationAdditionalS3DataSourcePropertyList {
	var returns TfAlgorithm_TrainingSpecificationAdditionalS3DataSourcePropertyList
	_jsii_.Get(
		j,
		"additionalS3DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) AdditionalS3DataSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalS3DataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) MetricDefinitions() TfAlgorithm_MetricDefinitionsPropertyList {
	var returns TfAlgorithm_MetricDefinitionsPropertyList
	_jsii_.Get(
		j,
		"metricDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) MetricDefinitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricDefinitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) SupportedHyperParameters() TfAlgorithm_SupportedHyperParametersPropertyList {
	var returns TfAlgorithm_SupportedHyperParametersPropertyList
	_jsii_.Get(
		j,
		"supportedHyperParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) SupportedHyperParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportedHyperParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) SupportedTrainingInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedTrainingInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) SupportedTrainingInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedTrainingInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) SupportedTuningJobObjectiveMetrics() TfAlgorithm_SupportedTuningJobObjectiveMetricsPropertyList {
	var returns TfAlgorithm_SupportedTuningJobObjectiveMetricsPropertyList
	_jsii_.Get(
		j,
		"supportedTuningJobObjectiveMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) SupportedTuningJobObjectiveMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportedTuningJobObjectiveMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) SupportsDistributedTraining() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsDistributedTraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) SupportsDistributedTrainingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsDistributedTrainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) TrainingChannels() TfAlgorithm_TrainingChannelsPropertyList {
	var returns TfAlgorithm_TrainingChannelsPropertyList
	_jsii_.Get(
		j,
		"trainingChannels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) TrainingChannelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trainingChannelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) TrainingImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) TrainingImageDigest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingImageDigest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) TrainingImageDigestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingImageDigestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) TrainingImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingImageInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAlgorithm_TrainingSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfAlgorithm_TrainingSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAlgorithm_TrainingSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAlgorithm.TrainingSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAlgorithm_TrainingSpecificationPropertyOutputReference_Override(t TfAlgorithm_TrainingSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAlgorithm.TrainingSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference)SetSupportedTrainingInstanceTypes(val *[]*string) {
	if err := j.validateSetSupportedTrainingInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedTrainingInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference)SetSupportsDistributedTraining(val interface{}) {
	if err := j.validateSetSupportsDistributedTrainingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsDistributedTraining",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference)SetTrainingImage(val *string) {
	if err := j.validateSetTrainingImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingImage",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference)SetTrainingImageDigest(val *string) {
	if err := j.validateSetTrainingImageDigestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingImageDigest",
		val,
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) PutAdditionalS3DataSource(value interface{}) {
	if err := t.validatePutAdditionalS3DataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAdditionalS3DataSource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) PutMetricDefinitions(value interface{}) {
	if err := t.validatePutMetricDefinitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMetricDefinitions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) PutSupportedHyperParameters(value interface{}) {
	if err := t.validatePutSupportedHyperParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSupportedHyperParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) PutSupportedTuningJobObjectiveMetrics(value interface{}) {
	if err := t.validatePutSupportedTuningJobObjectiveMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSupportedTuningJobObjectiveMetrics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) PutTrainingChannels(value interface{}) {
	if err := t.validatePutTrainingChannelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTrainingChannels",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) ResetAdditionalS3DataSource() {
	_jsii_.InvokeVoid(
		t,
		"resetAdditionalS3DataSource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) ResetMetricDefinitions() {
	_jsii_.InvokeVoid(
		t,
		"resetMetricDefinitions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) ResetSupportedHyperParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetSupportedHyperParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) ResetSupportedTuningJobObjectiveMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetSupportedTuningJobObjectiveMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) ResetSupportsDistributedTraining() {
	_jsii_.InvokeVoid(
		t,
		"resetSupportsDistributedTraining",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) ResetTrainingChannels() {
	_jsii_.InvokeVoid(
		t,
		"resetTrainingChannels",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) ResetTrainingImageDigest() {
	_jsii_.InvokeVoid(
		t,
		"resetTrainingImageDigest",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfAlgorithm_TrainingSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

