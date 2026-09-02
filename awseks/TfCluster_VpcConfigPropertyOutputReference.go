package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseks/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCluster_VpcConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ClusterSecurityGroupId() *string
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
	// Experimental.
	ControlPlaneEgressMode() *string
	// Experimental.
	SetControlPlaneEgressMode(val *string)
	// Experimental.
	ControlPlaneEgressModeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EndpointPrivateAccess() interface{}
	// Experimental.
	SetEndpointPrivateAccess(val interface{})
	// Experimental.
	EndpointPrivateAccessInput() interface{}
	// Experimental.
	EndpointPublicAccess() interface{}
	// Experimental.
	SetEndpointPublicAccess(val interface{})
	// Experimental.
	EndpointPublicAccessInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfCluster_VpcConfigProperty
	// Experimental.
	SetInternalValue(val *TfCluster_VpcConfigProperty)
	// Experimental.
	PublicAccessCidrs() *[]*string
	// Experimental.
	SetPublicAccessCidrs(val *[]*string)
	// Experimental.
	PublicAccessCidrsInput() *[]*string
	// Experimental.
	SecurityGroupIds() *[]*string
	// Experimental.
	SetSecurityGroupIds(val *[]*string)
	// Experimental.
	SecurityGroupIdsInput() *[]*string
	// Experimental.
	SubnetIds() *[]*string
	// Experimental.
	SetSubnetIds(val *[]*string)
	// Experimental.
	SubnetIdsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VpcId() *string
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
	ResetControlPlaneEgressMode()
	// Experimental.
	ResetEndpointPrivateAccess()
	// Experimental.
	ResetEndpointPublicAccess()
	// Experimental.
	ResetPublicAccessCidrs()
	// Experimental.
	ResetSecurityGroupIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCluster_VpcConfigPropertyOutputReference
type jsiiProxy_TfCluster_VpcConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) ClusterSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) ControlPlaneEgressMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneEgressMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) ControlPlaneEgressModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneEgressModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) EndpointPrivateAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointPrivateAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) EndpointPrivateAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointPrivateAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) EndpointPublicAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointPublicAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) EndpointPublicAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointPublicAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) InternalValue() *TfCluster_VpcConfigProperty {
	var returns *TfCluster_VpcConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) PublicAccessCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicAccessCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) PublicAccessCidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicAccessCidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCluster_VpcConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCluster_VpcConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCluster_VpcConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCluster_VpcConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eks.TfCluster.VpcConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCluster_VpcConfigPropertyOutputReference_Override(t TfCluster_VpcConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.TfCluster.VpcConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference)SetControlPlaneEgressMode(val *string) {
	if err := j.validateSetControlPlaneEgressModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controlPlaneEgressMode",
		val,
	)
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference)SetEndpointPrivateAccess(val interface{}) {
	if err := j.validateSetEndpointPrivateAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointPrivateAccess",
		val,
	)
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference)SetEndpointPublicAccess(val interface{}) {
	if err := j.validateSetEndpointPublicAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointPublicAccess",
		val,
	)
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference)SetInternalValue(val *TfCluster_VpcConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference)SetPublicAccessCidrs(val *[]*string) {
	if err := j.validateSetPublicAccessCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicAccessCidrs",
		val,
	)
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) ResetControlPlaneEgressMode() {
	_jsii_.InvokeVoid(
		t,
		"resetControlPlaneEgressMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) ResetEndpointPrivateAccess() {
	_jsii_.InvokeVoid(
		t,
		"resetEndpointPrivateAccess",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) ResetEndpointPublicAccess() {
	_jsii_.InvokeVoid(
		t,
		"resetEndpointPublicAccess",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) ResetPublicAccessCidrs() {
	_jsii_.InvokeVoid(
		t,
		"resetPublicAccessCidrs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCluster_VpcConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

