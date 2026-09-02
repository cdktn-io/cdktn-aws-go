package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrock/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrock/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEvaluationJob_AutomatedPropertyOutputReference interface {
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
	CustomMetricConfig() TfEvaluationJob_CustomMetricConfigPropertyList
	// Experimental.
	CustomMetricConfigInput() interface{}
	// Experimental.
	DatasetMetricConfig() TfEvaluationJob_EvaluationConfigAutomatedDatasetMetricConfigPropertyList
	// Experimental.
	DatasetMetricConfigInput() interface{}
	// Experimental.
	EvaluatorModelConfig() TfEvaluationJob_EvaluationConfigAutomatedEvaluatorModelConfigPropertyList
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

// The jsii proxy struct for TfEvaluationJob_AutomatedPropertyOutputReference
type jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) CustomMetricConfig() TfEvaluationJob_CustomMetricConfigPropertyList {
	var returns TfEvaluationJob_CustomMetricConfigPropertyList
	_jsii_.Get(
		j,
		"customMetricConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) CustomMetricConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customMetricConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) DatasetMetricConfig() TfEvaluationJob_EvaluationConfigAutomatedDatasetMetricConfigPropertyList {
	var returns TfEvaluationJob_EvaluationConfigAutomatedDatasetMetricConfigPropertyList
	_jsii_.Get(
		j,
		"datasetMetricConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) DatasetMetricConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datasetMetricConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) EvaluatorModelConfig() TfEvaluationJob_EvaluationConfigAutomatedEvaluatorModelConfigPropertyList {
	var returns TfEvaluationJob_EvaluationConfigAutomatedEvaluatorModelConfigPropertyList
	_jsii_.Get(
		j,
		"evaluatorModelConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) EvaluatorModelConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluatorModelConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEvaluationJob_AutomatedPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfEvaluationJob_AutomatedPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEvaluationJob_AutomatedPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock.TfEvaluationJob.AutomatedPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEvaluationJob_AutomatedPropertyOutputReference_Override(t TfEvaluationJob_AutomatedPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock.TfEvaluationJob.AutomatedPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) PutCustomMetricConfig(value interface{}) {
	if err := t.validatePutCustomMetricConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomMetricConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) PutDatasetMetricConfig(value interface{}) {
	if err := t.validatePutDatasetMetricConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDatasetMetricConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) PutEvaluatorModelConfig(value interface{}) {
	if err := t.validatePutEvaluatorModelConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEvaluatorModelConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) ResetCustomMetricConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomMetricConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) ResetDatasetMetricConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetDatasetMetricConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) ResetEvaluatorModelConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetEvaluatorModelConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEvaluationJob_AutomatedPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

