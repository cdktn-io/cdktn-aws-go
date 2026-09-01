package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference interface {
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
	MapBlockKey() *string
	// Experimental.
	SetMapBlockKey(val *string)
	// Experimental.
	MapBlockKeyInput() *string
	// Experimental.
	RelevanceScore() *float64
	// Experimental.
	SetRelevanceScore(val *float64)
	// Experimental.
	RelevanceScoreInput() *float64
	// Experimental.
	StrategyId() *string
	// Experimental.
	SetStrategyId(val *string)
	// Experimental.
	StrategyIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TopK() *float64
	// Experimental.
	SetTopK(val *float64)
	// Experimental.
	TopKInput() *float64
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
	ResetRelevanceScore()
	// Experimental.
	ResetStrategyId()
	// Experimental.
	ResetTopK()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference
type jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) MapBlockKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapBlockKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) MapBlockKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapBlockKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) RelevanceScore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"relevanceScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) RelevanceScoreInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"relevanceScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) StrategyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) StrategyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) TopK() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topK",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) TopKInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"topKInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreHarness.MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference_Override(a AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreHarness.MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference)SetMapBlockKey(val *string) {
	if err := j.validateSetMapBlockKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapBlockKey",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference)SetRelevanceScore(val *float64) {
	if err := j.validateSetRelevanceScoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"relevanceScore",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference)SetStrategyId(val *string) {
	if err := j.validateSetStrategyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strategyId",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference)SetTopK(val *float64) {
	if err := j.validateSetTopKParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topK",
		val,
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) ResetRelevanceScore() {
	_jsii_.InvokeVoid(
		a,
		"resetRelevanceScore",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) ResetStrategyId() {
	_jsii_.InvokeVoid(
		a,
		"resetStrategyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) ResetTopK() {
	_jsii_.InvokeVoid(
		a,
		"resetTopK",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_MemoryAgentcoreMemoryConfigurationRetrievalConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

