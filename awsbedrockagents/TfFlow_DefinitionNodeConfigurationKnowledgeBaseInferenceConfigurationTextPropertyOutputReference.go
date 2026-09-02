package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MaxTokens() *float64
	// Experimental.
	SetMaxTokens(val *float64)
	// Experimental.
	MaxTokensInput() *float64
	// Experimental.
	StopSequences() *[]*string
	// Experimental.
	SetStopSequences(val *[]*string)
	// Experimental.
	StopSequencesInput() *[]*string
	// Experimental.
	Temperature() *float64
	// Experimental.
	SetTemperature(val *float64)
	// Experimental.
	TemperatureInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TopP() *float64
	// Experimental.
	SetTopP(val *float64)
	// Experimental.
	TopPInput() *float64
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
	ResetMaxTokens()
	// Experimental.
	ResetStopSequences()
	// Experimental.
	ResetTemperature()
	// Experimental.
	ResetTopP()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference
type jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) MaxTokens() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) MaxTokensInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) StopSequences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stopSequences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) StopSequencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stopSequencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) Temperature() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temperature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) TemperatureInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"temperatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) TopP() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topP",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) TopPInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topPInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfFlow.DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference_Override(t TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfFlow.DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference)SetMaxTokens(val *float64) {
	if err := j.validateSetMaxTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTokens",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference)SetStopSequences(val *[]*string) {
	if err := j.validateSetStopSequencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stopSequences",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference)SetTemperature(val *float64) {
	if err := j.validateSetTemperatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"temperature",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference)SetTopP(val *float64) {
	if err := j.validateSetTopPParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topP",
		val,
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) ResetMaxTokens() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxTokens",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) ResetStopSequences() {
	_jsii_.InvokeVoid(
		t,
		"resetStopSequences",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) ResetTemperature() {
	_jsii_.InvokeVoid(
		t,
		"resetTemperature",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) ResetTopP() {
	_jsii_.InvokeVoid(
		t,
		"resetTopP",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationKnowledgeBaseInferenceConfigurationTextPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

