package eks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/eks/jsii"

	"github.com/cdktn-io/cdktn-aws-go/eks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCluster_VpcConfigPropertyOutputReference interface {
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
	InternalValue() *AwsCluster_VpcConfigProperty
	// Experimental.
	SetInternalValue(val *AwsCluster_VpcConfigProperty)
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

// The jsii proxy struct for AwsCluster_VpcConfigPropertyOutputReference
type jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) ClusterSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) ControlPlaneEgressMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneEgressMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) ControlPlaneEgressModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneEgressModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) EndpointPrivateAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointPrivateAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) EndpointPrivateAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointPrivateAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) EndpointPublicAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointPublicAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) EndpointPublicAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointPublicAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) InternalValue() *AwsCluster_VpcConfigProperty {
	var returns *AwsCluster_VpcConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) PublicAccessCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicAccessCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) PublicAccessCidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicAccessCidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCluster_VpcConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCluster_VpcConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCluster_VpcConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eks.AwsCluster.VpcConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCluster_VpcConfigPropertyOutputReference_Override(a AwsCluster_VpcConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.AwsCluster.VpcConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference)SetControlPlaneEgressMode(val *string) {
	if err := j.validateSetControlPlaneEgressModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controlPlaneEgressMode",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference)SetEndpointPrivateAccess(val interface{}) {
	if err := j.validateSetEndpointPrivateAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointPrivateAccess",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference)SetEndpointPublicAccess(val interface{}) {
	if err := j.validateSetEndpointPublicAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointPublicAccess",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference)SetInternalValue(val *AwsCluster_VpcConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference)SetPublicAccessCidrs(val *[]*string) {
	if err := j.validateSetPublicAccessCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicAccessCidrs",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) ResetControlPlaneEgressMode() {
	_jsii_.InvokeVoid(
		a,
		"resetControlPlaneEgressMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) ResetEndpointPrivateAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointPrivateAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) ResetEndpointPublicAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointPublicAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) ResetPublicAccessCidrs() {
	_jsii_.InvokeVoid(
		a,
		"resetPublicAccessCidrs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCluster_VpcConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

