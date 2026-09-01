package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AssociateCarrierIpAddress() *string
	// Experimental.
	SetAssociateCarrierIpAddress(val *string)
	// Experimental.
	AssociateCarrierIpAddressInput() *string
	// Experimental.
	AssociatePublicIpAddress() *string
	// Experimental.
	SetAssociatePublicIpAddress(val *string)
	// Experimental.
	AssociatePublicIpAddressInput() *string
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
	ConnectionTrackingSpecification() AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference
	// Experimental.
	ConnectionTrackingSpecificationInput() *AwsLaunchTemplate_ConnectionTrackingSpecificationProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DeleteOnTermination() *string
	// Experimental.
	SetDeleteOnTermination(val *string)
	// Experimental.
	DeleteOnTerminationInput() *string
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	DeviceIndex() *float64
	// Experimental.
	SetDeviceIndex(val *float64)
	// Experimental.
	DeviceIndexInput() *float64
	// Experimental.
	EnaQueueCount() *float64
	// Experimental.
	SetEnaQueueCount(val *float64)
	// Experimental.
	EnaQueueCountInput() *float64
	// Experimental.
	EnaSrdSpecification() AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference
	// Experimental.
	EnaSrdSpecificationInput() *AwsLaunchTemplate_EnaSrdSpecificationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InterfaceType() *string
	// Experimental.
	SetInterfaceType(val *string)
	// Experimental.
	InterfaceTypeInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Ipv4AddressCount() *float64
	// Experimental.
	SetIpv4AddressCount(val *float64)
	// Experimental.
	Ipv4AddressCountInput() *float64
	// Experimental.
	Ipv4Addresses() *[]*string
	// Experimental.
	SetIpv4Addresses(val *[]*string)
	// Experimental.
	Ipv4AddressesInput() *[]*string
	// Experimental.
	Ipv4PrefixCount() *float64
	// Experimental.
	SetIpv4PrefixCount(val *float64)
	// Experimental.
	Ipv4PrefixCountInput() *float64
	// Experimental.
	Ipv4Prefixes() *[]*string
	// Experimental.
	SetIpv4Prefixes(val *[]*string)
	// Experimental.
	Ipv4PrefixesInput() *[]*string
	// Experimental.
	Ipv6AddressCount() *float64
	// Experimental.
	SetIpv6AddressCount(val *float64)
	// Experimental.
	Ipv6AddressCountInput() *float64
	// Experimental.
	Ipv6Addresses() *[]*string
	// Experimental.
	SetIpv6Addresses(val *[]*string)
	// Experimental.
	Ipv6AddressesInput() *[]*string
	// Experimental.
	Ipv6PrefixCount() *float64
	// Experimental.
	SetIpv6PrefixCount(val *float64)
	// Experimental.
	Ipv6PrefixCountInput() *float64
	// Experimental.
	Ipv6Prefixes() *[]*string
	// Experimental.
	SetIpv6Prefixes(val *[]*string)
	// Experimental.
	Ipv6PrefixesInput() *[]*string
	// Experimental.
	NetworkCardIndex() *float64
	// Experimental.
	SetNetworkCardIndex(val *float64)
	// Experimental.
	NetworkCardIndexInput() *float64
	// Experimental.
	NetworkInterfaceId() *string
	// Experimental.
	SetNetworkInterfaceId(val *string)
	// Experimental.
	NetworkInterfaceIdInput() *string
	// Experimental.
	PrimaryIpv6() *string
	// Experimental.
	SetPrimaryIpv6(val *string)
	// Experimental.
	PrimaryIpv6Input() *string
	// Experimental.
	PrivateIpAddress() *string
	// Experimental.
	SetPrivateIpAddress(val *string)
	// Experimental.
	PrivateIpAddressInput() *string
	// Experimental.
	SecurityGroups() *[]*string
	// Experimental.
	SetSecurityGroups(val *[]*string)
	// Experimental.
	SecurityGroupsInput() *[]*string
	// Experimental.
	SubnetId() *string
	// Experimental.
	SetSubnetId(val *string)
	// Experimental.
	SubnetIdInput() *string
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
	PutConnectionTrackingSpecification(value *AwsLaunchTemplate_ConnectionTrackingSpecificationProperty)
	// Experimental.
	PutEnaSrdSpecification(value *AwsLaunchTemplate_EnaSrdSpecificationProperty)
	// Experimental.
	ResetAssociateCarrierIpAddress()
	// Experimental.
	ResetAssociatePublicIpAddress()
	// Experimental.
	ResetConnectionTrackingSpecification()
	// Experimental.
	ResetDeleteOnTermination()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetDeviceIndex()
	// Experimental.
	ResetEnaQueueCount()
	// Experimental.
	ResetEnaSrdSpecification()
	// Experimental.
	ResetInterfaceType()
	// Experimental.
	ResetIpv4AddressCount()
	// Experimental.
	ResetIpv4Addresses()
	// Experimental.
	ResetIpv4PrefixCount()
	// Experimental.
	ResetIpv4Prefixes()
	// Experimental.
	ResetIpv6AddressCount()
	// Experimental.
	ResetIpv6Addresses()
	// Experimental.
	ResetIpv6PrefixCount()
	// Experimental.
	ResetIpv6Prefixes()
	// Experimental.
	ResetNetworkCardIndex()
	// Experimental.
	ResetNetworkInterfaceId()
	// Experimental.
	ResetPrimaryIpv6()
	// Experimental.
	ResetPrivateIpAddress()
	// Experimental.
	ResetSecurityGroups()
	// Experimental.
	ResetSubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference
type jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) AssociateCarrierIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associateCarrierIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) AssociateCarrierIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associateCarrierIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) AssociatePublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) AssociatePublicIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatePublicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ConnectionTrackingSpecification() AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference {
	var returns AwsLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"connectionTrackingSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ConnectionTrackingSpecificationInput() *AwsLaunchTemplate_ConnectionTrackingSpecificationProperty {
	var returns *AwsLaunchTemplate_ConnectionTrackingSpecificationProperty
	_jsii_.Get(
		j,
		"connectionTrackingSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) DeleteOnTermination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteOnTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) DeleteOnTerminationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteOnTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) DeviceIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) DeviceIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) EnaQueueCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enaQueueCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) EnaQueueCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enaQueueCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) EnaSrdSpecification() AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference {
	var returns AwsLaunchTemplate_EnaSrdSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"enaSrdSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) EnaSrdSpecificationInput() *AwsLaunchTemplate_EnaSrdSpecificationProperty {
	var returns *AwsLaunchTemplate_EnaSrdSpecificationProperty
	_jsii_.Get(
		j,
		"enaSrdSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) InterfaceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) InterfaceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv4AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv4AddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4AddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv4Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv4AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv4PrefixCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4PrefixCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv4PrefixCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4PrefixCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv4Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv4PrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4PrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv6AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv6AddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv6Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv6AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv6PrefixCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6PrefixCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv6PrefixCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6PrefixCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv6Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv6PrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6PrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) NetworkCardIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkCardIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) NetworkCardIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkCardIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) NetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) NetworkInterfaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) PrimaryIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) PrimaryIpv6Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) PrivateIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLaunchTemplate_NetworkInterfacesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLaunchTemplate_NetworkInterfacesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsLaunchTemplate.NetworkInterfacesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLaunchTemplate_NetworkInterfacesPropertyOutputReference_Override(a AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsLaunchTemplate.NetworkInterfacesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetAssociateCarrierIpAddress(val *string) {
	if err := j.validateSetAssociateCarrierIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associateCarrierIpAddress",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetAssociatePublicIpAddress(val *string) {
	if err := j.validateSetAssociatePublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetDeleteOnTermination(val *string) {
	if err := j.validateSetDeleteOnTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOnTermination",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetDeviceIndex(val *float64) {
	if err := j.validateSetDeviceIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetEnaQueueCount(val *float64) {
	if err := j.validateSetEnaQueueCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enaQueueCount",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetInterfaceType(val *string) {
	if err := j.validateSetInterfaceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interfaceType",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetIpv4AddressCount(val *float64) {
	if err := j.validateSetIpv4AddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4AddressCount",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetIpv4Addresses(val *[]*string) {
	if err := j.validateSetIpv4AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Addresses",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetIpv4PrefixCount(val *float64) {
	if err := j.validateSetIpv4PrefixCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4PrefixCount",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetIpv4Prefixes(val *[]*string) {
	if err := j.validateSetIpv4PrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Prefixes",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetIpv6AddressCount(val *float64) {
	if err := j.validateSetIpv6AddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6AddressCount",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetIpv6Addresses(val *[]*string) {
	if err := j.validateSetIpv6AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Addresses",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetIpv6PrefixCount(val *float64) {
	if err := j.validateSetIpv6PrefixCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6PrefixCount",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetIpv6Prefixes(val *[]*string) {
	if err := j.validateSetIpv6PrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Prefixes",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetNetworkCardIndex(val *float64) {
	if err := j.validateSetNetworkCardIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkCardIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetNetworkInterfaceId(val *string) {
	if err := j.validateSetNetworkInterfaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaceId",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetPrimaryIpv6(val *string) {
	if err := j.validateSetPrimaryIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryIpv6",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetPrivateIpAddress(val *string) {
	if err := j.validateSetPrivateIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddress",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) PutConnectionTrackingSpecification(value *AwsLaunchTemplate_ConnectionTrackingSpecificationProperty) {
	if err := a.validatePutConnectionTrackingSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConnectionTrackingSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) PutEnaSrdSpecification(value *AwsLaunchTemplate_EnaSrdSpecificationProperty) {
	if err := a.validatePutEnaSrdSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnaSrdSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetAssociateCarrierIpAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetAssociateCarrierIpAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetAssociatePublicIpAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetAssociatePublicIpAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetConnectionTrackingSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionTrackingSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetDeleteOnTermination() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteOnTermination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetDeviceIndex() {
	_jsii_.InvokeVoid(
		a,
		"resetDeviceIndex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetEnaQueueCount() {
	_jsii_.InvokeVoid(
		a,
		"resetEnaQueueCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetEnaSrdSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetEnaSrdSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetInterfaceType() {
	_jsii_.InvokeVoid(
		a,
		"resetInterfaceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetIpv4AddressCount() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv4AddressCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetIpv4Addresses() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv4Addresses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetIpv4PrefixCount() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv4PrefixCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetIpv4Prefixes() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv4Prefixes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetIpv6AddressCount() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv6AddressCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetIpv6Addresses() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv6Addresses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetIpv6PrefixCount() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv6PrefixCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetIpv6Prefixes() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv6Prefixes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetNetworkCardIndex() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkCardIndex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetNetworkInterfaceId() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkInterfaceId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetPrimaryIpv6() {
	_jsii_.InvokeVoid(
		a,
		"resetPrimaryIpv6",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetPrivateIpAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateIpAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLaunchTemplate_NetworkInterfacesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

