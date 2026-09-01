package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BasePromptTemplate() *string
	// Experimental.
	SetBasePromptTemplate(val *string)
	// Experimental.
	BasePromptTemplateInput() *string
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
	InferenceConfiguration() AwsBedrockagentAgent_InferenceConfigurationPropertyList
	// Experimental.
	InferenceConfigurationInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	ParserMode() *string
	// Experimental.
	SetParserMode(val *string)
	// Experimental.
	ParserModeInput() *string
	// Experimental.
	PromptCreationMode() *string
	// Experimental.
	SetPromptCreationMode(val *string)
	// Experimental.
	PromptCreationModeInput() *string
	// Experimental.
	PromptState() *string
	// Experimental.
	SetPromptState(val *string)
	// Experimental.
	PromptStateInput() *string
	// Experimental.
	PromptType() *string
	// Experimental.
	SetPromptType(val *string)
	// Experimental.
	PromptTypeInput() *string
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
	PutInferenceConfiguration(value interface{})
	// Experimental.
	ResetBasePromptTemplate()
	// Experimental.
	ResetInferenceConfiguration()
	// Experimental.
	ResetParserMode()
	// Experimental.
	ResetPromptCreationMode()
	// Experimental.
	ResetPromptState()
	// Experimental.
	ResetPromptType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference
type jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) BasePromptTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basePromptTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) BasePromptTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basePromptTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) InferenceConfiguration() AwsBedrockagentAgent_InferenceConfigurationPropertyList {
	var returns AwsBedrockagentAgent_InferenceConfigurationPropertyList
	_jsii_.Get(
		j,
		"inferenceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) InferenceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) ParserMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parserMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) ParserModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parserModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) PromptCreationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptCreationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) PromptCreationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptCreationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) PromptState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) PromptStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) PromptType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) PromptTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBedrockagentAgent_PromptConfigurationsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentAgent.PromptConfigurationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference_Override(a AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentAgent.PromptConfigurationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference)SetBasePromptTemplate(val *string) {
	if err := j.validateSetBasePromptTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"basePromptTemplate",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference)SetParserMode(val *string) {
	if err := j.validateSetParserModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parserMode",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference)SetPromptCreationMode(val *string) {
	if err := j.validateSetPromptCreationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"promptCreationMode",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference)SetPromptState(val *string) {
	if err := j.validateSetPromptStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"promptState",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference)SetPromptType(val *string) {
	if err := j.validateSetPromptTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"promptType",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) PutInferenceConfiguration(value interface{}) {
	if err := a.validatePutInferenceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInferenceConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) ResetBasePromptTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetBasePromptTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) ResetInferenceConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetInferenceConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) ResetParserMode() {
	_jsii_.InvokeVoid(
		a,
		"resetParserMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) ResetPromptCreationMode() {
	_jsii_.InvokeVoid(
		a,
		"resetPromptCreationMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) ResetPromptState() {
	_jsii_.InvokeVoid(
		a,
		"resetPromptState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) ResetPromptType() {
	_jsii_.InvokeVoid(
		a,
		"resetPromptType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBedrockagentAgent_PromptConfigurationsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

