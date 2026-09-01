package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Agent() AwsBedrockagentFlow_AgentPropertyList
	// Experimental.
	AgentInput() interface{}
	// Experimental.
	Collector() AwsBedrockagentFlow_CollectorPropertyList
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
	Condition() AwsBedrockagentFlow_DefinitionNodeConfigurationConditionPropertyList
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
	InlineCode() AwsBedrockagentFlow_InlineCodePropertyList
	// Experimental.
	InlineCodeInput() interface{}
	// Experimental.
	Input() AwsBedrockagentFlow_DefinitionNodeConfigurationInputPropertyList
	// Experimental.
	InputInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Iterator() AwsBedrockagentFlow_IteratorPropertyList
	// Experimental.
	IteratorInput() interface{}
	// Experimental.
	KnowledgeBase() AwsBedrockagentFlow_KnowledgeBasePropertyList
	// Experimental.
	KnowledgeBaseInput() interface{}
	// Experimental.
	LambdaFunction() AwsBedrockagentFlow_LambdaFunctionPropertyList
	// Experimental.
	LambdaFunctionInput() interface{}
	// Experimental.
	Lex() AwsBedrockagentFlow_LexPropertyList
	// Experimental.
	LexInput() interface{}
	// Experimental.
	Output() AwsBedrockagentFlow_DefinitionNodeConfigurationOutputPropertyList
	// Experimental.
	OutputInput() interface{}
	// Experimental.
	Prompt() AwsBedrockagentFlow_PromptPropertyList
	// Experimental.
	PromptInput() interface{}
	// Experimental.
	Retrieval() AwsBedrockagentFlow_RetrievalPropertyList
	// Experimental.
	RetrievalInput() interface{}
	// Experimental.
	Storage() AwsBedrockagentFlow_StoragePropertyList
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

// The jsii proxy struct for AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference
type jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) Agent() AwsBedrockagentFlow_AgentPropertyList {
	var returns AwsBedrockagentFlow_AgentPropertyList
	_jsii_.Get(
		j,
		"agent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) AgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) Collector() AwsBedrockagentFlow_CollectorPropertyList {
	var returns AwsBedrockagentFlow_CollectorPropertyList
	_jsii_.Get(
		j,
		"collector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) CollectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) Condition() AwsBedrockagentFlow_DefinitionNodeConfigurationConditionPropertyList {
	var returns AwsBedrockagentFlow_DefinitionNodeConfigurationConditionPropertyList
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) InlineCode() AwsBedrockagentFlow_InlineCodePropertyList {
	var returns AwsBedrockagentFlow_InlineCodePropertyList
	_jsii_.Get(
		j,
		"inlineCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) InlineCodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inlineCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) Input() AwsBedrockagentFlow_DefinitionNodeConfigurationInputPropertyList {
	var returns AwsBedrockagentFlow_DefinitionNodeConfigurationInputPropertyList
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) InputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) Iterator() AwsBedrockagentFlow_IteratorPropertyList {
	var returns AwsBedrockagentFlow_IteratorPropertyList
	_jsii_.Get(
		j,
		"iterator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) IteratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iteratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) KnowledgeBase() AwsBedrockagentFlow_KnowledgeBasePropertyList {
	var returns AwsBedrockagentFlow_KnowledgeBasePropertyList
	_jsii_.Get(
		j,
		"knowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) KnowledgeBaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"knowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) LambdaFunction() AwsBedrockagentFlow_LambdaFunctionPropertyList {
	var returns AwsBedrockagentFlow_LambdaFunctionPropertyList
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) LambdaFunctionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) Lex() AwsBedrockagentFlow_LexPropertyList {
	var returns AwsBedrockagentFlow_LexPropertyList
	_jsii_.Get(
		j,
		"lex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) LexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) Output() AwsBedrockagentFlow_DefinitionNodeConfigurationOutputPropertyList {
	var returns AwsBedrockagentFlow_DefinitionNodeConfigurationOutputPropertyList
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) OutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) Prompt() AwsBedrockagentFlow_PromptPropertyList {
	var returns AwsBedrockagentFlow_PromptPropertyList
	_jsii_.Get(
		j,
		"prompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) PromptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) Retrieval() AwsBedrockagentFlow_RetrievalPropertyList {
	var returns AwsBedrockagentFlow_RetrievalPropertyList
	_jsii_.Get(
		j,
		"retrieval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) RetrievalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retrievalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) Storage() AwsBedrockagentFlow_StoragePropertyList {
	var returns AwsBedrockagentFlow_StoragePropertyList
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) StorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentFlow.DefinitionNodeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference_Override(a AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentFlow.DefinitionNodeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) PutAgent(value interface{}) {
	if err := a.validatePutAgentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAgent",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) PutCollector(value interface{}) {
	if err := a.validatePutCollectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCollector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) PutCondition(value interface{}) {
	if err := a.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCondition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) PutInlineCode(value interface{}) {
	if err := a.validatePutInlineCodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInlineCode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) PutInput(value interface{}) {
	if err := a.validatePutInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) PutIterator(value interface{}) {
	if err := a.validatePutIteratorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIterator",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) PutKnowledgeBase(value interface{}) {
	if err := a.validatePutKnowledgeBaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKnowledgeBase",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) PutLambdaFunction(value interface{}) {
	if err := a.validatePutLambdaFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaFunction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) PutLex(value interface{}) {
	if err := a.validatePutLexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLex",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) PutOutput(value interface{}) {
	if err := a.validatePutOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) PutPrompt(value interface{}) {
	if err := a.validatePutPromptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPrompt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) PutRetrieval(value interface{}) {
	if err := a.validatePutRetrievalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRetrieval",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) PutStorage(value interface{}) {
	if err := a.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStorage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetAgent() {
	_jsii_.InvokeVoid(
		a,
		"resetAgent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetCollector() {
	_jsii_.InvokeVoid(
		a,
		"resetCollector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		a,
		"resetCondition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetInlineCode() {
	_jsii_.InvokeVoid(
		a,
		"resetInlineCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetInput() {
	_jsii_.InvokeVoid(
		a,
		"resetInput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetIterator() {
	_jsii_.InvokeVoid(
		a,
		"resetIterator",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetKnowledgeBase() {
	_jsii_.InvokeVoid(
		a,
		"resetKnowledgeBase",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetLambdaFunction() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFunction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetLex() {
	_jsii_.InvokeVoid(
		a,
		"resetLex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetOutput() {
	_jsii_.InvokeVoid(
		a,
		"resetOutput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetPrompt() {
	_jsii_.InvokeVoid(
		a,
		"resetPrompt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetRetrieval() {
	_jsii_.InvokeVoid(
		a,
		"resetRetrieval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ResetStorage() {
	_jsii_.InvokeVoid(
		a,
		"resetStorage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBedrockagentFlow_DefinitionNodeConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

