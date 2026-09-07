package bedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference interface {
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
	EndpointIpAddressType() *string
	// Experimental.
	SetEndpointIpAddressType(val *string)
	// Experimental.
	EndpointIpAddressTypeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	RoutingDomain() *string
	// Experimental.
	SetRoutingDomain(val *string)
	// Experimental.
	RoutingDomainInput() *string
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
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VpcIdentifier() *string
	// Experimental.
	SetVpcIdentifier(val *string)
	// Experimental.
	VpcIdentifierInput() *string
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
	ResetRoutingDomain()
	// Experimental.
	ResetSecurityGroupIds()
	// Experimental.
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference
type jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) EndpointIpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIpAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) EndpointIpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIpAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) RoutingDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) RoutingDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) VpcIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) VpcIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdentifierInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsHarness.AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference_Override(a AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsHarness.AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference)SetEndpointIpAddressType(val *string) {
	if err := j.validateSetEndpointIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointIpAddressType",
		val,
	)
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference)SetRoutingDomain(val *string) {
	if err := j.validateSetRoutingDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingDomain",
		val,
	)
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference)SetVpcIdentifier(val *string) {
	if err := j.validateSetVpcIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcIdentifier",
		val,
	)
}

func (a *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) ResetRoutingDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

