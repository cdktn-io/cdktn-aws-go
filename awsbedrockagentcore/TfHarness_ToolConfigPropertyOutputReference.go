package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfHarness_ToolConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AgentcoreBrowser() TfHarness_AgentcoreBrowserPropertyList
	// Experimental.
	AgentcoreBrowserInput() interface{}
	// Experimental.
	AgentcoreCodeInterpreter() TfHarness_AgentcoreCodeInterpreterPropertyList
	// Experimental.
	AgentcoreCodeInterpreterInput() interface{}
	// Experimental.
	AgentcoreGateway() TfHarness_AgentcoreGatewayPropertyList
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
	InlineFunction() TfHarness_InlineFunctionPropertyList
	// Experimental.
	InlineFunctionInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	RemoteMcp() TfHarness_RemoteMcpPropertyList
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

// The jsii proxy struct for TfHarness_ToolConfigPropertyOutputReference
type jsiiProxy_TfHarness_ToolConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) AgentcoreBrowser() TfHarness_AgentcoreBrowserPropertyList {
	var returns TfHarness_AgentcoreBrowserPropertyList
	_jsii_.Get(
		j,
		"agentcoreBrowser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) AgentcoreBrowserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentcoreBrowserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) AgentcoreCodeInterpreter() TfHarness_AgentcoreCodeInterpreterPropertyList {
	var returns TfHarness_AgentcoreCodeInterpreterPropertyList
	_jsii_.Get(
		j,
		"agentcoreCodeInterpreter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) AgentcoreCodeInterpreterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentcoreCodeInterpreterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) AgentcoreGateway() TfHarness_AgentcoreGatewayPropertyList {
	var returns TfHarness_AgentcoreGatewayPropertyList
	_jsii_.Get(
		j,
		"agentcoreGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) AgentcoreGatewayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentcoreGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) InlineFunction() TfHarness_InlineFunctionPropertyList {
	var returns TfHarness_InlineFunctionPropertyList
	_jsii_.Get(
		j,
		"inlineFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) InlineFunctionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inlineFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) RemoteMcp() TfHarness_RemoteMcpPropertyList {
	var returns TfHarness_RemoteMcpPropertyList
	_jsii_.Get(
		j,
		"remoteMcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) RemoteMcpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteMcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfHarness_ToolConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfHarness_ToolConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfHarness_ToolConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfHarness_ToolConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfHarness.ToolConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfHarness_ToolConfigPropertyOutputReference_Override(t TfHarness_ToolConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfHarness.ToolConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) PutAgentcoreBrowser(value interface{}) {
	if err := t.validatePutAgentcoreBrowserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAgentcoreBrowser",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) PutAgentcoreCodeInterpreter(value interface{}) {
	if err := t.validatePutAgentcoreCodeInterpreterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAgentcoreCodeInterpreter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) PutAgentcoreGateway(value interface{}) {
	if err := t.validatePutAgentcoreGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAgentcoreGateway",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) PutInlineFunction(value interface{}) {
	if err := t.validatePutInlineFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInlineFunction",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) PutRemoteMcp(value interface{}) {
	if err := t.validatePutRemoteMcpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRemoteMcp",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) ResetAgentcoreBrowser() {
	_jsii_.InvokeVoid(
		t,
		"resetAgentcoreBrowser",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) ResetAgentcoreCodeInterpreter() {
	_jsii_.InvokeVoid(
		t,
		"resetAgentcoreCodeInterpreter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) ResetAgentcoreGateway() {
	_jsii_.InvokeVoid(
		t,
		"resetAgentcoreGateway",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) ResetInlineFunction() {
	_jsii_.InvokeVoid(
		t,
		"resetInlineFunction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) ResetRemoteMcp() {
	_jsii_.InvokeVoid(
		t,
		"resetRemoteMcp",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfHarness_ToolConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

