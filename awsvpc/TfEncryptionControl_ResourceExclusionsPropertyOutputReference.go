package awsvpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsvpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEncryptionControl_ResourceExclusionsPropertyOutputReference interface {
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
	EgressOnlyInternetGateway() TfEncryptionControl_EgressOnlyInternetGatewayPropertyOutputReference
	// Experimental.
	ElasticFileSystem() TfEncryptionControl_ElasticFileSystemPropertyOutputReference
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfEncryptionControl_ResourceExclusionsProperty
	// Experimental.
	SetInternalValue(val *TfEncryptionControl_ResourceExclusionsProperty)
	// Experimental.
	InternetGateway() TfEncryptionControl_InternetGatewayPropertyOutputReference
	// Experimental.
	Lambda() TfEncryptionControl_LambdaPropertyOutputReference
	// Experimental.
	NatGateway() TfEncryptionControl_NatGatewayPropertyOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VirtualPrivateGateway() TfEncryptionControl_VirtualPrivateGatewayPropertyOutputReference
	// Experimental.
	VpcLattice() TfEncryptionControl_VpcLatticePropertyOutputReference
	// Experimental.
	VpcPeering() TfEncryptionControl_VpcPeeringPropertyOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEncryptionControl_ResourceExclusionsPropertyOutputReference
type jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) EgressOnlyInternetGateway() TfEncryptionControl_EgressOnlyInternetGatewayPropertyOutputReference {
	var returns TfEncryptionControl_EgressOnlyInternetGatewayPropertyOutputReference
	_jsii_.Get(
		j,
		"egressOnlyInternetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) ElasticFileSystem() TfEncryptionControl_ElasticFileSystemPropertyOutputReference {
	var returns TfEncryptionControl_ElasticFileSystemPropertyOutputReference
	_jsii_.Get(
		j,
		"elasticFileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) InternalValue() *TfEncryptionControl_ResourceExclusionsProperty {
	var returns *TfEncryptionControl_ResourceExclusionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) InternetGateway() TfEncryptionControl_InternetGatewayPropertyOutputReference {
	var returns TfEncryptionControl_InternetGatewayPropertyOutputReference
	_jsii_.Get(
		j,
		"internetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) Lambda() TfEncryptionControl_LambdaPropertyOutputReference {
	var returns TfEncryptionControl_LambdaPropertyOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) NatGateway() TfEncryptionControl_NatGatewayPropertyOutputReference {
	var returns TfEncryptionControl_NatGatewayPropertyOutputReference
	_jsii_.Get(
		j,
		"natGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) VirtualPrivateGateway() TfEncryptionControl_VirtualPrivateGatewayPropertyOutputReference {
	var returns TfEncryptionControl_VirtualPrivateGatewayPropertyOutputReference
	_jsii_.Get(
		j,
		"virtualPrivateGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) VpcLattice() TfEncryptionControl_VpcLatticePropertyOutputReference {
	var returns TfEncryptionControl_VpcLatticePropertyOutputReference
	_jsii_.Get(
		j,
		"vpcLattice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) VpcPeering() TfEncryptionControl_VpcPeeringPropertyOutputReference {
	var returns TfEncryptionControl_VpcPeeringPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcPeering",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEncryptionControl_ResourceExclusionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEncryptionControl_ResourceExclusionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEncryptionControl_ResourceExclusionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc.TfEncryptionControl.ResourceExclusionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEncryptionControl_ResourceExclusionsPropertyOutputReference_Override(t TfEncryptionControl_ResourceExclusionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.TfEncryptionControl.ResourceExclusionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference)SetInternalValue(val *TfEncryptionControl_ResourceExclusionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEncryptionControl_ResourceExclusionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

