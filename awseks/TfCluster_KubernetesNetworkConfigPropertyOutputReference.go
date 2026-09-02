package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseks/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCluster_KubernetesNetworkConfigPropertyOutputReference interface {
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
	ElasticLoadBalancing() TfCluster_ElasticLoadBalancingPropertyOutputReference
	// Experimental.
	ElasticLoadBalancingInput() *TfCluster_ElasticLoadBalancingProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfCluster_KubernetesNetworkConfigProperty
	// Experimental.
	SetInternalValue(val *TfCluster_KubernetesNetworkConfigProperty)
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
	PutElasticLoadBalancing(value *TfCluster_ElasticLoadBalancingProperty)
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

// The jsii proxy struct for TfCluster_KubernetesNetworkConfigPropertyOutputReference
type jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) ElasticLoadBalancing() TfCluster_ElasticLoadBalancingPropertyOutputReference {
	var returns TfCluster_ElasticLoadBalancingPropertyOutputReference
	_jsii_.Get(
		j,
		"elasticLoadBalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) ElasticLoadBalancingInput() *TfCluster_ElasticLoadBalancingProperty {
	var returns *TfCluster_ElasticLoadBalancingProperty
	_jsii_.Get(
		j,
		"elasticLoadBalancingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) InternalValue() *TfCluster_KubernetesNetworkConfigProperty {
	var returns *TfCluster_KubernetesNetworkConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) IpFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) IpFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) ServiceIpv4Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceIpv4Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) ServiceIpv4CidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceIpv4CidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) ServiceIpv6Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceIpv6Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCluster_KubernetesNetworkConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCluster_KubernetesNetworkConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCluster_KubernetesNetworkConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eks.TfCluster.KubernetesNetworkConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCluster_KubernetesNetworkConfigPropertyOutputReference_Override(t TfCluster_KubernetesNetworkConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.TfCluster.KubernetesNetworkConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference)SetInternalValue(val *TfCluster_KubernetesNetworkConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference)SetIpFamily(val *string) {
	if err := j.validateSetIpFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipFamily",
		val,
	)
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference)SetServiceIpv4Cidr(val *string) {
	if err := j.validateSetServiceIpv4CidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceIpv4Cidr",
		val,
	)
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) PutElasticLoadBalancing(value *TfCluster_ElasticLoadBalancingProperty) {
	if err := t.validatePutElasticLoadBalancingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putElasticLoadBalancing",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) ResetElasticLoadBalancing() {
	_jsii_.InvokeVoid(
		t,
		"resetElasticLoadBalancing",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) ResetIpFamily() {
	_jsii_.InvokeVoid(
		t,
		"resetIpFamily",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) ResetServiceIpv4Cidr() {
	_jsii_.InvokeVoid(
		t,
		"resetServiceIpv4Cidr",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCluster_KubernetesNetworkConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

