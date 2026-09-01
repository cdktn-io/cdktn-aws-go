package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AgentcoreBrowser() AwsBedrockagentcoreHarness_AgentcoreBrowserPropertyList
	// Experimental.
	AgentcoreBrowserInput() interface{}
	// Experimental.
	AgentcoreCodeInterpreter() AwsBedrockagentcoreHarness_AgentcoreCodeInterpreterPropertyList
	// Experimental.
	AgentcoreCodeInterpreterInput() interface{}
	// Experimental.
	AgentcoreGateway() AwsBedrockagentcoreHarness_AgentcoreGatewayPropertyList
	// Experimental.
	AgentcoreGatewayInput() interface{}
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
	InlineFunction() AwsBedrockagentcoreHarness_InlineFunctionPropertyList
	// Experimental.
	InlineFunctionInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	RemoteMcp() AwsBedrockagentcoreHarness_RemoteMcpPropertyList
	// Experimental.
	RemoteMcpInput() interface{}
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
	PutAgentcoreBrowser(value interface{})
	// Experimental.
	PutAgentcoreCodeInterpreter(value interface{})
	// Experimental.
	PutAgentcoreGateway(value interface{})
	// Experimental.
	PutInlineFunction(value interface{})
	// Experimental.
	PutRemoteMcp(value interface{})
	// Experimental.
	ResetAgentcoreBrowser()
	// Experimental.
	ResetAgentcoreCodeInterpreter()
	// Experimental.
	ResetAgentcoreGateway()
	// Experimental.
	ResetInlineFunction()
	// Experimental.
	ResetRemoteMcp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference
type jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) AgentcoreBrowser() AwsBedrockagentcoreHarness_AgentcoreBrowserPropertyList {
	var returns AwsBedrockagentcoreHarness_AgentcoreBrowserPropertyList
	_jsii_.Get(
		j,
		"agentcoreBrowser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) AgentcoreBrowserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentcoreBrowserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) AgentcoreCodeInterpreter() AwsBedrockagentcoreHarness_AgentcoreCodeInterpreterPropertyList {
	var returns AwsBedrockagentcoreHarness_AgentcoreCodeInterpreterPropertyList
	_jsii_.Get(
		j,
		"agentcoreCodeInterpreter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) AgentcoreCodeInterpreterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentcoreCodeInterpreterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) AgentcoreGateway() AwsBedrockagentcoreHarness_AgentcoreGatewayPropertyList {
	var returns AwsBedrockagentcoreHarness_AgentcoreGatewayPropertyList
	_jsii_.Get(
		j,
		"agentcoreGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) AgentcoreGatewayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentcoreGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) InlineFunction() AwsBedrockagentcoreHarness_InlineFunctionPropertyList {
	var returns AwsBedrockagentcoreHarness_InlineFunctionPropertyList
	_jsii_.Get(
		j,
		"inlineFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) InlineFunctionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inlineFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) RemoteMcp() AwsBedrockagentcoreHarness_RemoteMcpPropertyList {
	var returns AwsBedrockagentcoreHarness_RemoteMcpPropertyList
	_jsii_.Get(
		j,
		"remoteMcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) RemoteMcpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteMcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBedrockagentcoreHarness_ToolConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreHarness.ToolConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference_Override(a AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreHarness.ToolConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) PutAgentcoreBrowser(value interface{}) {
	if err := a.validatePutAgentcoreBrowserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAgentcoreBrowser",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) PutAgentcoreCodeInterpreter(value interface{}) {
	if err := a.validatePutAgentcoreCodeInterpreterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAgentcoreCodeInterpreter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) PutAgentcoreGateway(value interface{}) {
	if err := a.validatePutAgentcoreGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAgentcoreGateway",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) PutInlineFunction(value interface{}) {
	if err := a.validatePutInlineFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInlineFunction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) PutRemoteMcp(value interface{}) {
	if err := a.validatePutRemoteMcpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRemoteMcp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) ResetAgentcoreBrowser() {
	_jsii_.InvokeVoid(
		a,
		"resetAgentcoreBrowser",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) ResetAgentcoreCodeInterpreter() {
	_jsii_.InvokeVoid(
		a,
		"resetAgentcoreCodeInterpreter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) ResetAgentcoreGateway() {
	_jsii_.InvokeVoid(
		a,
		"resetAgentcoreGateway",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) ResetInlineFunction() {
	_jsii_.InvokeVoid(
		a,
		"resetInlineFunction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) ResetRemoteMcp() {
	_jsii_.InvokeVoid(
		a,
		"resetRemoteMcp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_ToolConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

