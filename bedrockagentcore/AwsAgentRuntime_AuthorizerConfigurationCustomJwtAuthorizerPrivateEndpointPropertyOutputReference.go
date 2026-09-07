package bedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference interface {
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
	ManagedVpcResource() AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyList
	// Experimental.
	ManagedVpcResourceInput() interface{}
	// Experimental.
	SelfManagedLatticeResource() AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyList
	// Experimental.
	SelfManagedLatticeResourceInput() interface{}
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
	PutManagedVpcResource(value interface{})
	// Experimental.
	PutSelfManagedLatticeResource(value interface{})
	// Experimental.
	ResetManagedVpcResource()
	// Experimental.
	ResetSelfManagedLatticeResource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference
type jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) ManagedVpcResource() AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyList {
	var returns AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyList
	_jsii_.Get(
		j,
		"managedVpcResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) ManagedVpcResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedVpcResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) SelfManagedLatticeResource() AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyList {
	var returns AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyList
	_jsii_.Get(
		j,
		"selfManagedLatticeResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) SelfManagedLatticeResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedLatticeResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsAgentRuntime.AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference_Override(a AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsAgentRuntime.AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) PutManagedVpcResource(value interface{}) {
	if err := a.validatePutManagedVpcResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedVpcResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) PutSelfManagedLatticeResource(value interface{}) {
	if err := a.validatePutSelfManagedLatticeResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSelfManagedLatticeResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) ResetManagedVpcResource() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedVpcResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) ResetSelfManagedLatticeResource() {
	_jsii_.InvokeVoid(
		a,
		"resetSelfManagedLatticeResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

