package eks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/eks/jsii"

	"github.com/cdktn-io/cdktn-aws-go/eks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCluster_KubernetesNetworkConfigPropertyOutputReference interface {
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
	ElasticLoadBalancing() AwsCluster_ElasticLoadBalancingPropertyOutputReference
	// Experimental.
	ElasticLoadBalancingInput() *AwsCluster_ElasticLoadBalancingProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsCluster_KubernetesNetworkConfigProperty
	// Experimental.
	SetInternalValue(val *AwsCluster_KubernetesNetworkConfigProperty)
	// Experimental.
	IpFamily() *string
	// Experimental.
	SetIpFamily(val *string)
	// Experimental.
	IpFamilyInput() *string
	// Experimental.
	ServiceIpv4Cidr() *string
	// Experimental.
	SetServiceIpv4Cidr(val *string)
	// Experimental.
	ServiceIpv4CidrInput() *string
	// Experimental.
	ServiceIpv6Cidr() *string
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
	PutElasticLoadBalancing(value *AwsCluster_ElasticLoadBalancingProperty)
	// Experimental.
	ResetElasticLoadBalancing()
	// Experimental.
	ResetIpFamily()
	// Experimental.
	ResetServiceIpv4Cidr()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCluster_KubernetesNetworkConfigPropertyOutputReference
type jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) ElasticLoadBalancing() AwsCluster_ElasticLoadBalancingPropertyOutputReference {
	var returns AwsCluster_ElasticLoadBalancingPropertyOutputReference
	_jsii_.Get(
		j,
		"elasticLoadBalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) ElasticLoadBalancingInput() *AwsCluster_ElasticLoadBalancingProperty {
	var returns *AwsCluster_ElasticLoadBalancingProperty
	_jsii_.Get(
		j,
		"elasticLoadBalancingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) InternalValue() *AwsCluster_KubernetesNetworkConfigProperty {
	var returns *AwsCluster_KubernetesNetworkConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) IpFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) IpFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) ServiceIpv4Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceIpv4Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) ServiceIpv4CidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceIpv4CidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) ServiceIpv6Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceIpv6Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCluster_KubernetesNetworkConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCluster_KubernetesNetworkConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCluster_KubernetesNetworkConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eks.AwsCluster.KubernetesNetworkConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCluster_KubernetesNetworkConfigPropertyOutputReference_Override(a AwsCluster_KubernetesNetworkConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.AwsCluster.KubernetesNetworkConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference)SetInternalValue(val *AwsCluster_KubernetesNetworkConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference)SetIpFamily(val *string) {
	if err := j.validateSetIpFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipFamily",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference)SetServiceIpv4Cidr(val *string) {
	if err := j.validateSetServiceIpv4CidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceIpv4Cidr",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) PutElasticLoadBalancing(value *AwsCluster_ElasticLoadBalancingProperty) {
	if err := a.validatePutElasticLoadBalancingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putElasticLoadBalancing",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) ResetElasticLoadBalancing() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticLoadBalancing",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) ResetIpFamily() {
	_jsii_.InvokeVoid(
		a,
		"resetIpFamily",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) ResetServiceIpv4Cidr() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceIpv4Cidr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCluster_KubernetesNetworkConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

