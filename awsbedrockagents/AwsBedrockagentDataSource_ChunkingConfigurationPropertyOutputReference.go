package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ChunkingStrategy() *string
	// Experimental.
	SetChunkingStrategy(val *string)
	// Experimental.
	ChunkingStrategyInput() *string
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
	FixedSizeChunkingConfiguration() AwsBedrockagentDataSource_FixedSizeChunkingConfigurationPropertyList
	// Experimental.
	FixedSizeChunkingConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	HierarchicalChunkingConfiguration() AwsBedrockagentDataSource_HierarchicalChunkingConfigurationPropertyList
	// Experimental.
	HierarchicalChunkingConfigurationInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	SemanticChunkingConfiguration() AwsBedrockagentDataSource_SemanticChunkingConfigurationPropertyList
	// Experimental.
	SemanticChunkingConfigurationInput() interface{}
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
	PutFixedSizeChunkingConfiguration(value interface{})
	// Experimental.
	PutHierarchicalChunkingConfiguration(value interface{})
	// Experimental.
	PutSemanticChunkingConfiguration(value interface{})
	// Experimental.
	ResetFixedSizeChunkingConfiguration()
	// Experimental.
	ResetHierarchicalChunkingConfiguration()
	// Experimental.
	ResetSemanticChunkingConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference
type jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) ChunkingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chunkingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) ChunkingStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chunkingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) FixedSizeChunkingConfiguration() AwsBedrockagentDataSource_FixedSizeChunkingConfigurationPropertyList {
	var returns AwsBedrockagentDataSource_FixedSizeChunkingConfigurationPropertyList
	_jsii_.Get(
		j,
		"fixedSizeChunkingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) FixedSizeChunkingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fixedSizeChunkingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) HierarchicalChunkingConfiguration() AwsBedrockagentDataSource_HierarchicalChunkingConfigurationPropertyList {
	var returns AwsBedrockagentDataSource_HierarchicalChunkingConfigurationPropertyList
	_jsii_.Get(
		j,
		"hierarchicalChunkingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) HierarchicalChunkingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hierarchicalChunkingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) SemanticChunkingConfiguration() AwsBedrockagentDataSource_SemanticChunkingConfigurationPropertyList {
	var returns AwsBedrockagentDataSource_SemanticChunkingConfigurationPropertyList
	_jsii_.Get(
		j,
		"semanticChunkingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) SemanticChunkingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"semanticChunkingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentDataSource.ChunkingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference_Override(a AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentDataSource.ChunkingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference)SetChunkingStrategy(val *string) {
	if err := j.validateSetChunkingStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chunkingStrategy",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) PutFixedSizeChunkingConfiguration(value interface{}) {
	if err := a.validatePutFixedSizeChunkingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFixedSizeChunkingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) PutHierarchicalChunkingConfiguration(value interface{}) {
	if err := a.validatePutHierarchicalChunkingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHierarchicalChunkingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) PutSemanticChunkingConfiguration(value interface{}) {
	if err := a.validatePutSemanticChunkingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSemanticChunkingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) ResetFixedSizeChunkingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetFixedSizeChunkingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) ResetHierarchicalChunkingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetHierarchicalChunkingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) ResetSemanticChunkingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSemanticChunkingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBedrockagentDataSource_ChunkingConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

