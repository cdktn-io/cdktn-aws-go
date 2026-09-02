package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference interface {
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
	ManagedVpcResource() TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointManagedVpcResourcePropertyList
	// Experimental.
	ManagedVpcResourceInput() interface{}
	// Experimental.
	SelfManagedLatticeResource() TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointSelfManagedLatticeResourcePropertyList
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

// The jsii proxy struct for TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference
type jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) ManagedVpcResource() TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointManagedVpcResourcePropertyList {
	var returns TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointManagedVpcResourcePropertyList
	_jsii_.Get(
		j,
		"managedVpcResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) ManagedVpcResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedVpcResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) SelfManagedLatticeResource() TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointSelfManagedLatticeResourcePropertyList {
	var returns TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointSelfManagedLatticeResourcePropertyList
	_jsii_.Get(
		j,
		"selfManagedLatticeResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) SelfManagedLatticeResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedLatticeResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfGateway.AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference_Override(t TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfGateway.AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) PutManagedVpcResource(value interface{}) {
	if err := t.validatePutManagedVpcResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putManagedVpcResource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) PutSelfManagedLatticeResource(value interface{}) {
	if err := t.validatePutSelfManagedLatticeResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSelfManagedLatticeResource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) ResetManagedVpcResource() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedVpcResource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) ResetSelfManagedLatticeResource() {
	_jsii_.InvokeVoid(
		t,
		"resetSelfManagedLatticeResource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

