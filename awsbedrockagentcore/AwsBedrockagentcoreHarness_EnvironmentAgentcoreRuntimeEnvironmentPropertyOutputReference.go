package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AgentRuntimeArn() *string
	// Experimental.
	AgentRuntimeId() *string
	// Experimental.
	AgentRuntimeName() *string
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
	FilesystemConfiguration() AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyList
	// Experimental.
	FilesystemConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LifecycleConfiguration() AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentLifecycleConfigurationPropertyList
	// Experimental.
	LifecycleConfigurationInput() interface{}
	// Experimental.
	NetworkConfiguration() AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentNetworkConfigurationPropertyList
	// Experimental.
	NetworkConfigurationInput() interface{}
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
	PutFilesystemConfiguration(value interface{})
	// Experimental.
	PutLifecycleConfiguration(value interface{})
	// Experimental.
	PutNetworkConfiguration(value interface{})
	// Experimental.
	ResetFilesystemConfiguration()
	// Experimental.
	ResetLifecycleConfiguration()
	// Experimental.
	ResetNetworkConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference
type jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) AgentRuntimeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) AgentRuntimeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) AgentRuntimeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) FilesystemConfiguration() AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyList {
	var returns AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyList
	_jsii_.Get(
		j,
		"filesystemConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) FilesystemConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filesystemConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) LifecycleConfiguration() AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentLifecycleConfigurationPropertyList {
	var returns AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentLifecycleConfigurationPropertyList
	_jsii_.Get(
		j,
		"lifecycleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) LifecycleConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecycleConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) NetworkConfiguration() AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentNetworkConfigurationPropertyList {
	var returns AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentNetworkConfigurationPropertyList
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) NetworkConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreHarness.EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference_Override(a AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreHarness.EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) PutFilesystemConfiguration(value interface{}) {
	if err := a.validatePutFilesystemConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilesystemConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) PutLifecycleConfiguration(value interface{}) {
	if err := a.validatePutLifecycleConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLifecycleConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) PutNetworkConfiguration(value interface{}) {
	if err := a.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) ResetFilesystemConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetFilesystemConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) ResetLifecycleConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLifecycleConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness_EnvironmentAgentcoreRuntimeEnvironmentPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

