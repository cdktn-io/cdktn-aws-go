package bedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/bedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/bedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFlow_DefinitionNodeConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Agent() AwsFlow_AgentPropertyList
	// Experimental.
	AgentInput() interface{}
	// Experimental.
	Collector() AwsFlow_CollectorPropertyList
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
	Condition() AwsFlow_DefinitionNodeConfigurationConditionPropertyList
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
	InlineCode() AwsFlow_InlineCodePropertyList
	// Experimental.
	InlineCodeInput() interface{}
	// Experimental.
	Input() AwsFlow_DefinitionNodeConfigurationInputPropertyList
	// Experimental.
	InputInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Iterator() AwsFlow_IteratorPropertyList
	// Experimental.
	IteratorInput() interface{}
	// Experimental.
	KnowledgeBase() AwsFlow_KnowledgeBasePropertyList
	// Experimental.
	KnowledgeBaseInput() interface{}
	// Experimental.
	LambdaFunction() AwsFlow_LambdaFunctionPropertyList
	// Experimental.
	LambdaFunctionInput() interface{}
	// Experimental.
	Lex() AwsFlow_LexPropertyList
	// Experimental.
	LexInput() interface{}
	// Experimental.
	Output() AwsFlow_DefinitionNodeConfigurationOutputPropertyList
	// Experimental.
	OutputInput() interface{}
	// Experimental.
	Prompt() AwsFlow_PromptPropertyList
	// Experimental.
	PromptInput() interface{}
	// Experimental.
	Retrieval() AwsFlow_RetrievalPropertyList
	// Experimental.
	RetrievalInput() interface{}
	// Experimental.
	Storage() AwsFlow_StoragePropertyList
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

// The jsii proxy struct for AwsFlow_DefinitionNodeConfigurationPropertyOutputReference
type jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) Agent() AwsFlow_AgentPropertyList {
	var returns AwsFlow_AgentPropertyList
	_jsii_.Get(
		j,
		"agent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) AgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) Collector() AwsFlow_CollectorPropertyList {
	var returns AwsFlow_CollectorPropertyList
	_jsii_.Get(
		j,
		"collector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) CollectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) Condition() AwsFlow_DefinitionNodeConfigurationConditionPropertyList {
	var returns AwsFlow_DefinitionNodeConfigurationConditionPropertyList
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) InlineCode() AwsFlow_InlineCodePropertyList {
	var returns AwsFlow_InlineCodePropertyList
	_jsii_.Get(
		j,
		"inlineCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) InlineCodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inlineCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) Input() AwsFlow_DefinitionNodeConfigurationInputPropertyList {
	var returns AwsFlow_DefinitionNodeConfigurationInputPropertyList
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) InputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) Iterator() AwsFlow_IteratorPropertyList {
	var returns AwsFlow_IteratorPropertyList
	_jsii_.Get(
		j,
		"iterator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) IteratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iteratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) KnowledgeBase() AwsFlow_KnowledgeBasePropertyList {
	var returns AwsFlow_KnowledgeBasePropertyList
	_jsii_.Get(
		j,
		"knowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) KnowledgeBaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"knowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) LambdaFunction() AwsFlow_LambdaFunctionPropertyList {
	var returns AwsFlow_LambdaFunctionPropertyList
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) LambdaFunctionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) Lex() AwsFlow_LexPropertyList {
	var returns AwsFlow_LexPropertyList
	_jsii_.Get(
		j,
		"lex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) LexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) Output() AwsFlow_DefinitionNodeConfigurationOutputPropertyList {
	var returns AwsFlow_DefinitionNodeConfigurationOutputPropertyList
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) OutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) Prompt() AwsFlow_PromptPropertyList {
	var returns AwsFlow_PromptPropertyList
	_jsii_.Get(
		j,
		"prompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) PromptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) Retrieval() AwsFlow_RetrievalPropertyList {
	var returns AwsFlow_RetrievalPropertyList
	_jsii_.Get(
		j,
		"retrieval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) RetrievalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retrievalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) Storage() AwsFlow_StoragePropertyList {
	var returns AwsFlow_StoragePropertyList
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) StorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFlow_DefinitionNodeConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsFlow_DefinitionNodeConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFlow_DefinitionNodeConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsFlow.DefinitionNodeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFlow_DefinitionNodeConfigurationPropertyOutputReference_Override(a AwsFlow_DefinitionNodeConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsFlow.DefinitionNodeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) PutAgent(value interface{}) {
	if err := a.validatePutAgentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAgent",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) PutCollector(value interface{}) {
	if err := a.validatePutCollectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCollector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) PutCondition(value interface{}) {
	if err := a.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCondition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) PutInlineCode(value interface{}) {
	if err := a.validatePutInlineCodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInlineCode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) PutInput(value interface{}) {
	if err := a.validatePutInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) PutIterator(value interface{}) {
	if err := a.validatePutIteratorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIterator",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) PutKnowledgeBase(value interface{}) {
	if err := a.validatePutKnowledgeBaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKnowledgeBase",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) PutLambdaFunction(value interface{}) {
	if err := a.validatePutLambdaFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaFunction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) PutLex(value interface{}) {
	if err := a.validatePutLexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLex",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) PutOutput(value interface{}) {
	if err := a.validatePutOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) PutPrompt(value interface{}) {
	if err := a.validatePutPromptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPrompt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) PutRetrieval(value interface{}) {
	if err := a.validatePutRetrievalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRetrieval",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) PutStorage(value interface{}) {
	if err := a.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStorage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetAgent() {
	_jsii_.InvokeVoid(
		a,
		"resetAgent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetCollector() {
	_jsii_.InvokeVoid(
		a,
		"resetCollector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		a,
		"resetCondition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetInlineCode() {
	_jsii_.InvokeVoid(
		a,
		"resetInlineCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetInput() {
	_jsii_.InvokeVoid(
		a,
		"resetInput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetIterator() {
	_jsii_.InvokeVoid(
		a,
		"resetIterator",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetKnowledgeBase() {
	_jsii_.InvokeVoid(
		a,
		"resetKnowledgeBase",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetLambdaFunction() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFunction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetLex() {
	_jsii_.InvokeVoid(
		a,
		"resetLex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetOutput() {
	_jsii_.InvokeVoid(
		a,
		"resetOutput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetPrompt() {
	_jsii_.InvokeVoid(
		a,
		"resetPrompt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetRetrieval() {
	_jsii_.InvokeVoid(
		a,
		"resetRetrieval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetStorage() {
	_jsii_.InvokeVoid(
		a,
		"resetStorage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFlow_DefinitionNodeConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

