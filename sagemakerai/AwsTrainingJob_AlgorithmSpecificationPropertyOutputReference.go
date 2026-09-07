package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AlgorithmName() *string
	// Experimental.
	SetAlgorithmName(val *string)
	// Experimental.
	AlgorithmNameInput() *string
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
	ContainerArguments() *[]*string
	// Experimental.
	SetContainerArguments(val *[]*string)
	// Experimental.
	ContainerArgumentsInput() *[]*string
	// Experimental.
	ContainerEntrypoint() *[]*string
	// Experimental.
	SetContainerEntrypoint(val *[]*string)
	// Experimental.
	ContainerEntrypointInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EnableSagemakerMetricsTimeSeries() interface{}
	// Experimental.
	SetEnableSagemakerMetricsTimeSeries(val interface{})
	// Experimental.
	EnableSagemakerMetricsTimeSeriesInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MetricDefinitions() AwsTrainingJob_MetricDefinitionsPropertyList
	// Experimental.
	MetricDefinitionsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TrainingImage() *string
	// Experimental.
	SetTrainingImage(val *string)
	// Experimental.
	TrainingImageConfig() AwsTrainingJob_TrainingImageConfigPropertyList
	// Experimental.
	TrainingImageConfigInput() interface{}
	// Experimental.
	TrainingImageInput() *string
	// Experimental.
	TrainingInputMode() *string
	// Experimental.
	SetTrainingInputMode(val *string)
	// Experimental.
	TrainingInputModeInput() *string
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
	PutMetricDefinitions(value interface{})
	// Experimental.
	PutTrainingImageConfig(value interface{})
	// Experimental.
	ResetAlgorithmName()
	// Experimental.
	ResetContainerArguments()
	// Experimental.
	ResetContainerEntrypoint()
	// Experimental.
	ResetEnableSagemakerMetricsTimeSeries()
	// Experimental.
	ResetMetricDefinitions()
	// Experimental.
	ResetTrainingImage()
	// Experimental.
	ResetTrainingImageConfig()
	// Experimental.
	ResetTrainingInputMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference
type jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) AlgorithmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) AlgorithmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) ContainerArguments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) ContainerArgumentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) ContainerEntrypoint() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerEntrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) ContainerEntrypointInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerEntrypointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) EnableSagemakerMetricsTimeSeries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSagemakerMetricsTimeSeries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) EnableSagemakerMetricsTimeSeriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSagemakerMetricsTimeSeriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) MetricDefinitions() AwsTrainingJob_MetricDefinitionsPropertyList {
	var returns AwsTrainingJob_MetricDefinitionsPropertyList
	_jsii_.Get(
		j,
		"metricDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) MetricDefinitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricDefinitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) TrainingImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) TrainingImageConfig() AwsTrainingJob_TrainingImageConfigPropertyList {
	var returns AwsTrainingJob_TrainingImageConfigPropertyList
	_jsii_.Get(
		j,
		"trainingImageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) TrainingImageConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trainingImageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) TrainingImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) TrainingInputMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingInputMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) TrainingInputModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingInputModeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTrainingJob_AlgorithmSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTrainingJob_AlgorithmSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsTrainingJob.AlgorithmSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTrainingJob_AlgorithmSpecificationPropertyOutputReference_Override(a AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsTrainingJob.AlgorithmSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference)SetAlgorithmName(val *string) {
	if err := j.validateSetAlgorithmNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithmName",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference)SetContainerArguments(val *[]*string) {
	if err := j.validateSetContainerArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerArguments",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference)SetContainerEntrypoint(val *[]*string) {
	if err := j.validateSetContainerEntrypointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerEntrypoint",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference)SetEnableSagemakerMetricsTimeSeries(val interface{}) {
	if err := j.validateSetEnableSagemakerMetricsTimeSeriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSagemakerMetricsTimeSeries",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference)SetTrainingImage(val *string) {
	if err := j.validateSetTrainingImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingImage",
		val,
	)
}

func (j *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference)SetTrainingInputMode(val *string) {
	if err := j.validateSetTrainingInputModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingInputMode",
		val,
	)
}

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) PutMetricDefinitions(value interface{}) {
	if err := a.validatePutMetricDefinitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetricDefinitions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) PutTrainingImageConfig(value interface{}) {
	if err := a.validatePutTrainingImageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrainingImageConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) ResetAlgorithmName() {
	_jsii_.InvokeVoid(
		a,
		"resetAlgorithmName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) ResetContainerArguments() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerArguments",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) ResetContainerEntrypoint() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerEntrypoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) ResetEnableSagemakerMetricsTimeSeries() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableSagemakerMetricsTimeSeries",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) ResetMetricDefinitions() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricDefinitions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) ResetTrainingImage() {
	_jsii_.InvokeVoid(
		a,
		"resetTrainingImage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) ResetTrainingImageConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetTrainingImageConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) ResetTrainingInputMode() {
	_jsii_.InvokeVoid(
		a,
		"resetTrainingInputMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTrainingJob_AlgorithmSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

