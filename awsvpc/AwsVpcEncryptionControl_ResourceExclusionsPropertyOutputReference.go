package awsvpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsvpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference interface {
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
	EgressOnlyInternetGateway() AwsVpcEncryptionControl_EgressOnlyInternetGatewayPropertyOutputReference
	// Experimental.
	ElasticFileSystem() AwsVpcEncryptionControl_ElasticFileSystemPropertyOutputReference
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsVpcEncryptionControl_ResourceExclusionsProperty
	// Experimental.
	SetInternalValue(val *AwsVpcEncryptionControl_ResourceExclusionsProperty)
	// Experimental.
	InternetGateway() AwsVpcEncryptionControl_InternetGatewayPropertyOutputReference
	// Experimental.
	Lambda() AwsVpcEncryptionControl_LambdaPropertyOutputReference
	// Experimental.
	NatGateway() AwsVpcEncryptionControl_NatGatewayPropertyOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VirtualPrivateGateway() AwsVpcEncryptionControl_VirtualPrivateGatewayPropertyOutputReference
	// Experimental.
	VpcLattice() AwsVpcEncryptionControl_VpcLatticePropertyOutputReference
	// Experimental.
	VpcPeering() AwsVpcEncryptionControl_VpcPeeringPropertyOutputReference
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

// The jsii proxy struct for AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference
type jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) EgressOnlyInternetGateway() AwsVpcEncryptionControl_EgressOnlyInternetGatewayPropertyOutputReference {
	var returns AwsVpcEncryptionControl_EgressOnlyInternetGatewayPropertyOutputReference
	_jsii_.Get(
		j,
		"egressOnlyInternetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) ElasticFileSystem() AwsVpcEncryptionControl_ElasticFileSystemPropertyOutputReference {
	var returns AwsVpcEncryptionControl_ElasticFileSystemPropertyOutputReference
	_jsii_.Get(
		j,
		"elasticFileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) InternalValue() *AwsVpcEncryptionControl_ResourceExclusionsProperty {
	var returns *AwsVpcEncryptionControl_ResourceExclusionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) InternetGateway() AwsVpcEncryptionControl_InternetGatewayPropertyOutputReference {
	var returns AwsVpcEncryptionControl_InternetGatewayPropertyOutputReference
	_jsii_.Get(
		j,
		"internetGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) Lambda() AwsVpcEncryptionControl_LambdaPropertyOutputReference {
	var returns AwsVpcEncryptionControl_LambdaPropertyOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) NatGateway() AwsVpcEncryptionControl_NatGatewayPropertyOutputReference {
	var returns AwsVpcEncryptionControl_NatGatewayPropertyOutputReference
	_jsii_.Get(
		j,
		"natGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) VirtualPrivateGateway() AwsVpcEncryptionControl_VirtualPrivateGatewayPropertyOutputReference {
	var returns AwsVpcEncryptionControl_VirtualPrivateGatewayPropertyOutputReference
	_jsii_.Get(
		j,
		"virtualPrivateGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) VpcLattice() AwsVpcEncryptionControl_VpcLatticePropertyOutputReference {
	var returns AwsVpcEncryptionControl_VpcLatticePropertyOutputReference
	_jsii_.Get(
		j,
		"vpcLattice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) VpcPeering() AwsVpcEncryptionControl_VpcPeeringPropertyOutputReference {
	var returns AwsVpcEncryptionControl_VpcPeeringPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcPeering",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsVpcEncryptionControl.ResourceExclusionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference_Override(a AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsVpcEncryptionControl.ResourceExclusionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference)SetInternalValue(val *AwsVpcEncryptionControl_ResourceExclusionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsVpcEncryptionControl_ResourceExclusionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

