package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbatch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbatch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllocationStrategy() *string
	// Experimental.
	SetAllocationStrategy(val *string)
	// Experimental.
	AllocationStrategyInput() *string
	// Experimental.
	BidPercentage() *float64
	// Experimental.
	SetBidPercentage(val *float64)
	// Experimental.
	BidPercentageInput() *float64
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
	DesiredVcpus() *float64
	// Experimental.
	SetDesiredVcpus(val *float64)
	// Experimental.
	DesiredVcpusInput() *float64
	// Experimental.
	Ec2Configuration() AwsBatchComputeEnvironment_Ec2ConfigurationPropertyList
	// Experimental.
	Ec2ConfigurationInput() interface{}
	// Experimental.
	Ec2KeyPair() *string
	// Experimental.
	SetEc2KeyPair(val *string)
	// Experimental.
	Ec2KeyPairInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	ImageId() *string
	// Experimental.
	SetImageId(val *string)
	// Experimental.
	ImageIdInput() *string
	// Experimental.
	InstanceRole() *string
	// Experimental.
	SetInstanceRole(val *string)
	// Experimental.
	InstanceRoleInput() *string
	// Experimental.
	InstanceType() *[]*string
	// Experimental.
	SetInstanceType(val *[]*string)
	// Experimental.
	InstanceTypeInput() *[]*string
	// Experimental.
	InternalValue() *AwsBatchComputeEnvironment_ComputeResourcesProperty
	// Experimental.
	SetInternalValue(val *AwsBatchComputeEnvironment_ComputeResourcesProperty)
	// Experimental.
	LaunchTemplate() AwsBatchComputeEnvironment_LaunchTemplatePropertyOutputReference
	// Experimental.
	LaunchTemplateInput() *AwsBatchComputeEnvironment_LaunchTemplateProperty
	// Experimental.
	MaxVcpus() *float64
	// Experimental.
	SetMaxVcpus(val *float64)
	// Experimental.
	MaxVcpusInput() *float64
	// Experimental.
	MinVcpus() *float64
	// Experimental.
	SetMinVcpus(val *float64)
	// Experimental.
	MinVcpusInput() *float64
	// Experimental.
	PlacementGroup() *string
	// Experimental.
	SetPlacementGroup(val *string)
	// Experimental.
	PlacementGroupInput() *string
	// Experimental.
	SecurityGroupIds() *[]*string
	// Experimental.
	SetSecurityGroupIds(val *[]*string)
	// Experimental.
	SecurityGroupIdsInput() *[]*string
	// Experimental.
	SpotIamFleetRole() *string
	// Experimental.
	SetSpotIamFleetRole(val *string)
	// Experimental.
	SpotIamFleetRoleInput() *string
	// Experimental.
	Subnets() *[]*string
	// Experimental.
	SetSubnets(val *[]*string)
	// Experimental.
	SubnetsInput() *[]*string
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
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	PutEc2Configuration(value interface{})
	// Experimental.
	PutLaunchTemplate(value *AwsBatchComputeEnvironment_LaunchTemplateProperty)
	// Experimental.
	ResetAllocationStrategy()
	// Experimental.
	ResetBidPercentage()
	// Experimental.
	ResetDesiredVcpus()
	// Experimental.
	ResetEc2Configuration()
	// Experimental.
	ResetEc2KeyPair()
	// Experimental.
	ResetImageId()
	// Experimental.
	ResetInstanceRole()
	// Experimental.
	ResetInstanceType()
	// Experimental.
	ResetLaunchTemplate()
	// Experimental.
	ResetMinVcpus()
	// Experimental.
	ResetPlacementGroup()
	// Experimental.
	ResetSecurityGroupIds()
	// Experimental.
	ResetSpotIamFleetRole()
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

// The jsii proxy struct for AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference
type jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) AllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) AllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) BidPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bidPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) BidPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bidPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) DesiredVcpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredVcpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) DesiredVcpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredVcpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) Ec2Configuration() AwsBatchComputeEnvironment_Ec2ConfigurationPropertyList {
	var returns AwsBatchComputeEnvironment_Ec2ConfigurationPropertyList
	_jsii_.Get(
		j,
		"ec2Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) Ec2ConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) Ec2KeyPair() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2KeyPair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) Ec2KeyPairInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2KeyPairInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) InstanceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) InstanceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) InstanceType() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) InstanceTypeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) InternalValue() *AwsBatchComputeEnvironment_ComputeResourcesProperty {
	var returns *AwsBatchComputeEnvironment_ComputeResourcesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) LaunchTemplate() AwsBatchComputeEnvironment_LaunchTemplatePropertyOutputReference {
	var returns AwsBatchComputeEnvironment_LaunchTemplatePropertyOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) LaunchTemplateInput() *AwsBatchComputeEnvironment_LaunchTemplateProperty {
	var returns *AwsBatchComputeEnvironment_LaunchTemplateProperty
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) MaxVcpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVcpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) MaxVcpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVcpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) MinVcpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minVcpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) MinVcpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minVcpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) PlacementGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) PlacementGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) SpotIamFleetRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotIamFleetRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) SpotIamFleetRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotIamFleetRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) Subnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) SubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-batch.AwsBatchComputeEnvironment.ComputeResourcesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference_Override(a AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-batch.AwsBatchComputeEnvironment.ComputeResourcesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetAllocationStrategy(val *string) {
	if err := j.validateSetAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationStrategy",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetBidPercentage(val *float64) {
	if err := j.validateSetBidPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bidPercentage",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetDesiredVcpus(val *float64) {
	if err := j.validateSetDesiredVcpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredVcpus",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetEc2KeyPair(val *string) {
	if err := j.validateSetEc2KeyPairParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2KeyPair",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetImageId(val *string) {
	if err := j.validateSetImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetInstanceRole(val *string) {
	if err := j.validateSetInstanceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceRole",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetInstanceType(val *[]*string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetInternalValue(val *AwsBatchComputeEnvironment_ComputeResourcesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetMaxVcpus(val *float64) {
	if err := j.validateSetMaxVcpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxVcpus",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetMinVcpus(val *float64) {
	if err := j.validateSetMinVcpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minVcpus",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetPlacementGroup(val *string) {
	if err := j.validateSetPlacementGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementGroup",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetSpotIamFleetRole(val *string) {
	if err := j.validateSetSpotIamFleetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotIamFleetRole",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetSubnets(val *[]*string) {
	if err := j.validateSetSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) PutEc2Configuration(value interface{}) {
	if err := a.validatePutEc2ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEc2Configuration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) PutLaunchTemplate(value *AwsBatchComputeEnvironment_LaunchTemplateProperty) {
	if err := a.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ResetAllocationStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetAllocationStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ResetBidPercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetBidPercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ResetDesiredVcpus() {
	_jsii_.InvokeVoid(
		a,
		"resetDesiredVcpus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ResetEc2Configuration() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2Configuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ResetEc2KeyPair() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2KeyPair",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ResetImageId() {
	_jsii_.InvokeVoid(
		a,
		"resetImageId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ResetInstanceRole() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ResetMinVcpus() {
	_jsii_.InvokeVoid(
		a,
		"resetMinVcpus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ResetPlacementGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetPlacementGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ResetSpotIamFleetRole() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotIamFleetRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBatchComputeEnvironment_ComputeResourcesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

