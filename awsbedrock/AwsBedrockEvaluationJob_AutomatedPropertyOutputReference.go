package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrock/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrock/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBedrockEvaluationJob_AutomatedPropertyOutputReference interface {
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
	CustomMetricConfig() AwsBedrockEvaluationJob_CustomMetricConfigPropertyList
	// Experimental.
	CustomMetricConfigInput() interface{}
	// Experimental.
	DatasetMetricConfig() AwsBedrockEvaluationJob_EvaluationConfigAutomatedDatasetMetricConfigPropertyList
	// Experimental.
	DatasetMetricConfigInput() interface{}
	// Experimental.
	EvaluatorModelConfig() AwsBedrockEvaluationJob_EvaluationConfigAutomatedEvaluatorModelConfigPropertyList
	// Experimental.
	EvaluatorModelConfigInput() interface{}
	// Experimental.
	Fqn() *string
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
	PutCustomMetricConfig(value interface{})
	// Experimental.
	PutDatasetMetricConfig(value interface{})
	// Experimental.
	PutEvaluatorModelConfig(value interface{})
	// Experimental.
	ResetCustomMetricConfig()
	// Experimental.
	ResetDatasetMetricConfig()
	// Experimental.
	ResetEvaluatorModelConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBedrockEvaluationJob_AutomatedPropertyOutputReference
type jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) CustomMetricConfig() AwsBedrockEvaluationJob_CustomMetricConfigPropertyList {
	var returns AwsBedrockEvaluationJob_CustomMetricConfigPropertyList
	_jsii_.Get(
		j,
		"customMetricConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) CustomMetricConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customMetricConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) DatasetMetricConfig() AwsBedrockEvaluationJob_EvaluationConfigAutomatedDatasetMetricConfigPropertyList {
	var returns AwsBedrockEvaluationJob_EvaluationConfigAutomatedDatasetMetricConfigPropertyList
	_jsii_.Get(
		j,
		"datasetMetricConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) DatasetMetricConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datasetMetricConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) EvaluatorModelConfig() AwsBedrockEvaluationJob_EvaluationConfigAutomatedEvaluatorModelConfigPropertyList {
	var returns AwsBedrockEvaluationJob_EvaluationConfigAutomatedEvaluatorModelConfigPropertyList
	_jsii_.Get(
		j,
		"evaluatorModelConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) EvaluatorModelConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluatorModelConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBedrockEvaluationJob_AutomatedPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBedrockEvaluationJob_AutomatedPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBedrockEvaluationJob_AutomatedPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock.AwsBedrockEvaluationJob.AutomatedPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBedrockEvaluationJob_AutomatedPropertyOutputReference_Override(a AwsBedrockEvaluationJob_AutomatedPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock.AwsBedrockEvaluationJob.AutomatedPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) PutCustomMetricConfig(value interface{}) {
	if err := a.validatePutCustomMetricConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomMetricConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) PutDatasetMetricConfig(value interface{}) {
	if err := a.validatePutDatasetMetricConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatasetMetricConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) PutEvaluatorModelConfig(value interface{}) {
	if err := a.validatePutEvaluatorModelConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEvaluatorModelConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) ResetCustomMetricConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomMetricConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) ResetDatasetMetricConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDatasetMetricConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) ResetEvaluatorModelConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEvaluatorModelConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBedrockEvaluationJob_AutomatedPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

