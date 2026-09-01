package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Ami() *string
	// Experimental.
	SetAmi(val *string)
	// Experimental.
	AmiInput() *string
	// Experimental.
	AssociatePublicIpAddress() interface{}
	// Experimental.
	SetAssociatePublicIpAddress(val interface{})
	// Experimental.
	AssociatePublicIpAddressInput() interface{}
	// Experimental.
	AvailabilityZone() *string
	// Experimental.
	SetAvailabilityZone(val *string)
	// Experimental.
	AvailabilityZoneInput() *string
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
	EbsBlockDevice() AwsSpotFleetRequest_EbsBlockDevicePropertyList
	// Experimental.
	EbsBlockDeviceInput() interface{}
	// Experimental.
	EbsOptimized() interface{}
	// Experimental.
	SetEbsOptimized(val interface{})
	// Experimental.
	EbsOptimizedInput() interface{}
	// Experimental.
	EphemeralBlockDevice() AwsSpotFleetRequest_EphemeralBlockDevicePropertyList
	// Experimental.
	EphemeralBlockDeviceInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	IamInstanceProfile() *string
	// Experimental.
	SetIamInstanceProfile(val *string)
	// Experimental.
	IamInstanceProfileArn() *string
	// Experimental.
	SetIamInstanceProfileArn(val *string)
	// Experimental.
	IamInstanceProfileArnInput() *string
	// Experimental.
	IamInstanceProfileInput() *string
	// Experimental.
	InstanceType() *string
	// Experimental.
	SetInstanceType(val *string)
	// Experimental.
	InstanceTypeInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	KeyName() *string
	// Experimental.
	SetKeyName(val *string)
	// Experimental.
	KeyNameInput() *string
	// Experimental.
	Monitoring() interface{}
	// Experimental.
	SetMonitoring(val interface{})
	// Experimental.
	MonitoringInput() interface{}
	// Experimental.
	PlacementGroup() *string
	// Experimental.
	SetPlacementGroup(val *string)
	// Experimental.
	PlacementGroupInput() *string
	// Experimental.
	PlacementTenancy() *string
	// Experimental.
	SetPlacementTenancy(val *string)
	// Experimental.
	PlacementTenancyInput() *string
	// Experimental.
	RootBlockDevice() AwsSpotFleetRequest_RootBlockDevicePropertyList
	// Experimental.
	RootBlockDeviceInput() interface{}
	// Experimental.
	SpotPrice() *string
	// Experimental.
	SetSpotPrice(val *string)
	// Experimental.
	SpotPriceInput() *string
	// Experimental.
	SubnetId() *string
	// Experimental.
	SetSubnetId(val *string)
	// Experimental.
	SubnetIdInput() *string
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
	UserData() *string
	// Experimental.
	SetUserData(val *string)
	// Experimental.
	UserDataInput() *string
	// Experimental.
	VpcSecurityGroupIds() *[]*string
	// Experimental.
	SetVpcSecurityGroupIds(val *[]*string)
	// Experimental.
	VpcSecurityGroupIdsInput() *[]*string
	// Experimental.
	WeightedCapacity() *string
	// Experimental.
	SetWeightedCapacity(val *string)
	// Experimental.
	WeightedCapacityInput() *string
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
	PutEbsBlockDevice(value interface{})
	// Experimental.
	PutEphemeralBlockDevice(value interface{})
	// Experimental.
	PutRootBlockDevice(value interface{})
	// Experimental.
	ResetAssociatePublicIpAddress()
	// Experimental.
	ResetAvailabilityZone()
	// Experimental.
	ResetEbsBlockDevice()
	// Experimental.
	ResetEbsOptimized()
	// Experimental.
	ResetEphemeralBlockDevice()
	// Experimental.
	ResetIamInstanceProfile()
	// Experimental.
	ResetIamInstanceProfileArn()
	// Experimental.
	ResetKeyName()
	// Experimental.
	ResetMonitoring()
	// Experimental.
	ResetPlacementGroup()
	// Experimental.
	ResetPlacementTenancy()
	// Experimental.
	ResetRootBlockDevice()
	// Experimental.
	ResetSpotPrice()
	// Experimental.
	ResetSubnetId()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetUserData()
	// Experimental.
	ResetVpcSecurityGroupIds()
	// Experimental.
	ResetWeightedCapacity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference
type jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) Ami() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ami",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) AmiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) AssociatePublicIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) AssociatePublicIpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) EbsBlockDevice() AwsSpotFleetRequest_EbsBlockDevicePropertyList {
	var returns AwsSpotFleetRequest_EbsBlockDevicePropertyList
	_jsii_.Get(
		j,
		"ebsBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) EbsBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) EphemeralBlockDevice() AwsSpotFleetRequest_EphemeralBlockDevicePropertyList {
	var returns AwsSpotFleetRequest_EphemeralBlockDevicePropertyList
	_jsii_.Get(
		j,
		"ephemeralBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) EphemeralBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) IamInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) IamInstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) Monitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) MonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) PlacementGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) PlacementGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) PlacementTenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementTenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) PlacementTenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementTenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) RootBlockDevice() AwsSpotFleetRequest_RootBlockDevicePropertyList {
	var returns AwsSpotFleetRequest_RootBlockDevicePropertyList
	_jsii_.Get(
		j,
		"rootBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) RootBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) SpotPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) SpotPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) WeightedCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weightedCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) WeightedCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weightedCapacityInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSpotFleetRequest_LaunchSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsSpotFleetRequest.LaunchSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference_Override(a AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsSpotFleetRequest.LaunchSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetAmi(val *string) {
	if err := j.validateSetAmiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ami",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetAssociatePublicIpAddress(val interface{}) {
	if err := j.validateSetAssociatePublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetIamInstanceProfileArn(val *string) {
	if err := j.validateSetIamInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetMonitoring(val interface{}) {
	if err := j.validateSetMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoring",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetPlacementGroup(val *string) {
	if err := j.validateSetPlacementGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementGroup",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetPlacementTenancy(val *string) {
	if err := j.validateSetPlacementTenancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementTenancy",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetSpotPrice(val *string) {
	if err := j.validateSetSpotPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotPrice",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference)SetWeightedCapacity(val *string) {
	if err := j.validateSetWeightedCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weightedCapacity",
		val,
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) PutEbsBlockDevice(value interface{}) {
	if err := a.validatePutEbsBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEbsBlockDevice",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) PutEphemeralBlockDevice(value interface{}) {
	if err := a.validatePutEphemeralBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEphemeralBlockDevice",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) PutRootBlockDevice(value interface{}) {
	if err := a.validatePutRootBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRootBlockDevice",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetAssociatePublicIpAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetAssociatePublicIpAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetEbsBlockDevice() {
	_jsii_.InvokeVoid(
		a,
		"resetEbsBlockDevice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		a,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetEphemeralBlockDevice() {
	_jsii_.InvokeVoid(
		a,
		"resetEphemeralBlockDevice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		a,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetIamInstanceProfileArn() {
	_jsii_.InvokeVoid(
		a,
		"resetIamInstanceProfileArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetKeyName() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetMonitoring() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetPlacementGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetPlacementGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetPlacementTenancy() {
	_jsii_.InvokeVoid(
		a,
		"resetPlacementTenancy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetRootBlockDevice() {
	_jsii_.InvokeVoid(
		a,
		"resetRootBlockDevice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetSpotPrice() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotPrice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetUserData() {
	_jsii_.InvokeVoid(
		a,
		"resetUserData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ResetWeightedCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetWeightedCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSpotFleetRequest_LaunchSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

