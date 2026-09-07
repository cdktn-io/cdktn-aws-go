package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/bedrock/jsii"

	"github.com/cdktn-io/cdktn-aws-go/bedrock/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEvaluationJob_HumanPropertyOutputReference interface {
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
	CustomMetric() AwsEvaluationJob_EvaluationConfigHumanCustomMetricPropertyList
	// Experimental.
	CustomMetricInput() interface{}
	// Experimental.
	DatasetMetricConfig() AwsEvaluationJob_EvaluationConfigHumanDatasetMetricConfigPropertyList
	// Experimental.
	DatasetMetricConfigInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	HumanWorkflowConfig() AwsEvaluationJob_HumanWorkflowConfigPropertyList
	// Experimental.
	HumanWorkflowConfigInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
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
	PutCustomMetric(value interface{})
	// Experimental.
	PutDatasetMetricConfig(value interface{})
	// Experimental.
	PutHumanWorkflowConfig(value interface{})
	// Experimental.
	ResetCustomMetric()
	// Experimental.
	ResetDatasetMetricConfig()
	// Experimental.
	ResetHumanWorkflowConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEvaluationJob_HumanPropertyOutputReference
type jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) CustomMetric() AwsEvaluationJob_EvaluationConfigHumanCustomMetricPropertyList {
	var returns AwsEvaluationJob_EvaluationConfigHumanCustomMetricPropertyList
	_jsii_.Get(
		j,
		"customMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) CustomMetricInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) DatasetMetricConfig() AwsEvaluationJob_EvaluationConfigHumanDatasetMetricConfigPropertyList {
	var returns AwsEvaluationJob_EvaluationConfigHumanDatasetMetricConfigPropertyList
	_jsii_.Get(
		j,
		"datasetMetricConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) DatasetMetricConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datasetMetricConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) HumanWorkflowConfig() AwsEvaluationJob_HumanWorkflowConfigPropertyList {
	var returns AwsEvaluationJob_HumanWorkflowConfigPropertyList
	_jsii_.Get(
		j,
		"humanWorkflowConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) HumanWorkflowConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"humanWorkflowConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEvaluationJob_HumanPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsEvaluationJob_HumanPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEvaluationJob_HumanPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock.AwsEvaluationJob.HumanPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEvaluationJob_HumanPropertyOutputReference_Override(a AwsEvaluationJob_HumanPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock.AwsEvaluationJob.HumanPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) PutCustomMetric(value interface{}) {
	if err := a.validatePutCustomMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomMetric",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) PutDatasetMetricConfig(value interface{}) {
	if err := a.validatePutDatasetMetricConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatasetMetricConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) PutHumanWorkflowConfig(value interface{}) {
	if err := a.validatePutHumanWorkflowConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHumanWorkflowConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) ResetCustomMetric() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomMetric",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) ResetDatasetMetricConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDatasetMetricConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) ResetHumanWorkflowConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetHumanWorkflowConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEvaluationJob_HumanPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

