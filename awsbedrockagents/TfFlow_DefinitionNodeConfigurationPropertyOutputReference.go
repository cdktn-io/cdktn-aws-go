package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFlow_DefinitionNodeConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Agent() TfFlow_AgentPropertyList
	// Experimental.
	AgentInput() interface{}
	// Experimental.
	Collector() TfFlow_CollectorPropertyList
	// Experimental.
	CollectorInput() interface{}
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
	Condition() TfFlow_DefinitionNodeConfigurationConditionPropertyList
	// Experimental.
	ConditionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InlineCode() TfFlow_InlineCodePropertyList
	// Experimental.
	InlineCodeInput() interface{}
	// Experimental.
	Input() TfFlow_DefinitionNodeConfigurationInputPropertyList
	// Experimental.
	InputInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Iterator() TfFlow_IteratorPropertyList
	// Experimental.
	IteratorInput() interface{}
	// Experimental.
	KnowledgeBase() TfFlow_KnowledgeBasePropertyList
	// Experimental.
	KnowledgeBaseInput() interface{}
	// Experimental.
	LambdaFunction() TfFlow_LambdaFunctionPropertyList
	// Experimental.
	LambdaFunctionInput() interface{}
	// Experimental.
	Lex() TfFlow_LexPropertyList
	// Experimental.
	LexInput() interface{}
	// Experimental.
	Output() TfFlow_DefinitionNodeConfigurationOutputPropertyList
	// Experimental.
	OutputInput() interface{}
	// Experimental.
	Prompt() TfFlow_PromptPropertyList
	// Experimental.
	PromptInput() interface{}
	// Experimental.
	Retrieval() TfFlow_RetrievalPropertyList
	// Experimental.
	RetrievalInput() interface{}
	// Experimental.
	Storage() TfFlow_StoragePropertyList
	// Experimental.
	StorageInput() interface{}
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
	PutAgent(value interface{})
	// Experimental.
	PutCollector(value interface{})
	// Experimental.
	PutCondition(value interface{})
	// Experimental.
	PutInlineCode(value interface{})
	// Experimental.
	PutInput(value interface{})
	// Experimental.
	PutIterator(value interface{})
	// Experimental.
	PutKnowledgeBase(value interface{})
	// Experimental.
	PutLambdaFunction(value interface{})
	// Experimental.
	PutLex(value interface{})
	// Experimental.
	PutOutput(value interface{})
	// Experimental.
	PutPrompt(value interface{})
	// Experimental.
	PutRetrieval(value interface{})
	// Experimental.
	PutStorage(value interface{})
	// Experimental.
	ResetAgent()
	// Experimental.
	ResetCollector()
	// Experimental.
	ResetCondition()
	// Experimental.
	ResetInlineCode()
	// Experimental.
	ResetInput()
	// Experimental.
	ResetIterator()
	// Experimental.
	ResetKnowledgeBase()
	// Experimental.
	ResetLambdaFunction()
	// Experimental.
	ResetLex()
	// Experimental.
	ResetOutput()
	// Experimental.
	ResetPrompt()
	// Experimental.
	ResetRetrieval()
	// Experimental.
	ResetStorage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFlow_DefinitionNodeConfigurationPropertyOutputReference
type jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) Agent() TfFlow_AgentPropertyList {
	var returns TfFlow_AgentPropertyList
	_jsii_.Get(
		j,
		"agent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) AgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) Collector() TfFlow_CollectorPropertyList {
	var returns TfFlow_CollectorPropertyList
	_jsii_.Get(
		j,
		"collector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) CollectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) Condition() TfFlow_DefinitionNodeConfigurationConditionPropertyList {
	var returns TfFlow_DefinitionNodeConfigurationConditionPropertyList
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) InlineCode() TfFlow_InlineCodePropertyList {
	var returns TfFlow_InlineCodePropertyList
	_jsii_.Get(
		j,
		"inlineCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) InlineCodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inlineCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) Input() TfFlow_DefinitionNodeConfigurationInputPropertyList {
	var returns TfFlow_DefinitionNodeConfigurationInputPropertyList
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) InputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) Iterator() TfFlow_IteratorPropertyList {
	var returns TfFlow_IteratorPropertyList
	_jsii_.Get(
		j,
		"iterator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) IteratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iteratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) KnowledgeBase() TfFlow_KnowledgeBasePropertyList {
	var returns TfFlow_KnowledgeBasePropertyList
	_jsii_.Get(
		j,
		"knowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) KnowledgeBaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"knowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) LambdaFunction() TfFlow_LambdaFunctionPropertyList {
	var returns TfFlow_LambdaFunctionPropertyList
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) LambdaFunctionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) Lex() TfFlow_LexPropertyList {
	var returns TfFlow_LexPropertyList
	_jsii_.Get(
		j,
		"lex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) LexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) Output() TfFlow_DefinitionNodeConfigurationOutputPropertyList {
	var returns TfFlow_DefinitionNodeConfigurationOutputPropertyList
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) OutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) Prompt() TfFlow_PromptPropertyList {
	var returns TfFlow_PromptPropertyList
	_jsii_.Get(
		j,
		"prompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) PromptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) Retrieval() TfFlow_RetrievalPropertyList {
	var returns TfFlow_RetrievalPropertyList
	_jsii_.Get(
		j,
		"retrieval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) RetrievalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retrievalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) Storage() TfFlow_StoragePropertyList {
	var returns TfFlow_StoragePropertyList
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) StorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFlow_DefinitionNodeConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfFlow_DefinitionNodeConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFlow_DefinitionNodeConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfFlow.DefinitionNodeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFlow_DefinitionNodeConfigurationPropertyOutputReference_Override(t TfFlow_DefinitionNodeConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfFlow.DefinitionNodeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) PutAgent(value interface{}) {
	if err := t.validatePutAgentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAgent",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) PutCollector(value interface{}) {
	if err := t.validatePutCollectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCollector",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) PutCondition(value interface{}) {
	if err := t.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCondition",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) PutInlineCode(value interface{}) {
	if err := t.validatePutInlineCodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInlineCode",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) PutInput(value interface{}) {
	if err := t.validatePutInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInput",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) PutIterator(value interface{}) {
	if err := t.validatePutIteratorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIterator",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) PutKnowledgeBase(value interface{}) {
	if err := t.validatePutKnowledgeBaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKnowledgeBase",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) PutLambdaFunction(value interface{}) {
	if err := t.validatePutLambdaFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambdaFunction",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) PutLex(value interface{}) {
	if err := t.validatePutLexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLex",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) PutOutput(value interface{}) {
	if err := t.validatePutOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOutput",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) PutPrompt(value interface{}) {
	if err := t.validatePutPromptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPrompt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) PutRetrieval(value interface{}) {
	if err := t.validatePutRetrievalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRetrieval",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) PutStorage(value interface{}) {
	if err := t.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStorage",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetAgent() {
	_jsii_.InvokeVoid(
		t,
		"resetAgent",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetCollector() {
	_jsii_.InvokeVoid(
		t,
		"resetCollector",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		t,
		"resetCondition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetInlineCode() {
	_jsii_.InvokeVoid(
		t,
		"resetInlineCode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetInput() {
	_jsii_.InvokeVoid(
		t,
		"resetInput",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetIterator() {
	_jsii_.InvokeVoid(
		t,
		"resetIterator",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetKnowledgeBase() {
	_jsii_.InvokeVoid(
		t,
		"resetKnowledgeBase",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetLambdaFunction() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaFunction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetLex() {
	_jsii_.InvokeVoid(
		t,
		"resetLex",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetOutput() {
	_jsii_.InvokeVoid(
		t,
		"resetOutput",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetPrompt() {
	_jsii_.InvokeVoid(
		t,
		"resetPrompt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetRetrieval() {
	_jsii_.InvokeVoid(
		t,
		"resetRetrieval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetStorage() {
	_jsii_.InvokeVoid(
		t,
		"resetStorage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFlow_DefinitionNodeConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

