package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfLaunchTemplate_NetworkInterfacesPropertyOutputReference interface {
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
	ConnectionTrackingSpecification() TfLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference
	// Experimental.
	ConnectionTrackingSpecificationInput() *TfLaunchTemplate_ConnectionTrackingSpecificationProperty
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
	EnaSrdSpecification() TfLaunchTemplate_EnaSrdSpecificationPropertyOutputReference
	// Experimental.
	EnaSrdSpecificationInput() *TfLaunchTemplate_EnaSrdSpecificationProperty
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
	PutConnectionTrackingSpecification(value *TfLaunchTemplate_ConnectionTrackingSpecificationProperty)
	// Experimental.
	PutEnaSrdSpecification(value *TfLaunchTemplate_EnaSrdSpecificationProperty)
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

// The jsii proxy struct for TfLaunchTemplate_NetworkInterfacesPropertyOutputReference
type jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) AssociateCarrierIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associateCarrierIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) AssociateCarrierIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associateCarrierIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) AssociatePublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) AssociatePublicIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatePublicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ConnectionTrackingSpecification() TfLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference {
	var returns TfLaunchTemplate_ConnectionTrackingSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"connectionTrackingSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ConnectionTrackingSpecificationInput() *TfLaunchTemplate_ConnectionTrackingSpecificationProperty {
	var returns *TfLaunchTemplate_ConnectionTrackingSpecificationProperty
	_jsii_.Get(
		j,
		"connectionTrackingSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) DeleteOnTermination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteOnTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) DeleteOnTerminationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteOnTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) DeviceIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) DeviceIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deviceIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) EnaQueueCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enaQueueCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) EnaQueueCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enaQueueCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) EnaSrdSpecification() TfLaunchTemplate_EnaSrdSpecificationPropertyOutputReference {
	var returns TfLaunchTemplate_EnaSrdSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"enaSrdSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) EnaSrdSpecificationInput() *TfLaunchTemplate_EnaSrdSpecificationProperty {
	var returns *TfLaunchTemplate_EnaSrdSpecificationProperty
	_jsii_.Get(
		j,
		"enaSrdSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) InterfaceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) InterfaceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv4AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv4AddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4AddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv4Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv4AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv4PrefixCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4PrefixCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv4PrefixCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4PrefixCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv4Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv4PrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4PrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv6AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv6AddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv6Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv6AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv6PrefixCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6PrefixCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv6PrefixCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6PrefixCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv6Prefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Prefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Ipv6PrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6PrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) NetworkCardIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkCardIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) NetworkCardIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkCardIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) NetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) NetworkInterfaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) PrimaryIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) PrimaryIpv6Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) PrivateIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfLaunchTemplate_NetworkInterfacesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfLaunchTemplate_NetworkInterfacesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfLaunchTemplate_NetworkInterfacesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2.TfLaunchTemplate.NetworkInterfacesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfLaunchTemplate_NetworkInterfacesPropertyOutputReference_Override(t TfLaunchTemplate_NetworkInterfacesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.TfLaunchTemplate.NetworkInterfacesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetAssociateCarrierIpAddress(val *string) {
	if err := j.validateSetAssociateCarrierIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associateCarrierIpAddress",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetAssociatePublicIpAddress(val *string) {
	if err := j.validateSetAssociatePublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetDeleteOnTermination(val *string) {
	if err := j.validateSetDeleteOnTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteOnTermination",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetDeviceIndex(val *float64) {
	if err := j.validateSetDeviceIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceIndex",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetEnaQueueCount(val *float64) {
	if err := j.validateSetEnaQueueCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enaQueueCount",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetInterfaceType(val *string) {
	if err := j.validateSetInterfaceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interfaceType",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetIpv4AddressCount(val *float64) {
	if err := j.validateSetIpv4AddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4AddressCount",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetIpv4Addresses(val *[]*string) {
	if err := j.validateSetIpv4AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Addresses",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetIpv4PrefixCount(val *float64) {
	if err := j.validateSetIpv4PrefixCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4PrefixCount",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetIpv4Prefixes(val *[]*string) {
	if err := j.validateSetIpv4PrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4Prefixes",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetIpv6AddressCount(val *float64) {
	if err := j.validateSetIpv6AddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6AddressCount",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetIpv6Addresses(val *[]*string) {
	if err := j.validateSetIpv6AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Addresses",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetIpv6PrefixCount(val *float64) {
	if err := j.validateSetIpv6PrefixCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6PrefixCount",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetIpv6Prefixes(val *[]*string) {
	if err := j.validateSetIpv6PrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Prefixes",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetNetworkCardIndex(val *float64) {
	if err := j.validateSetNetworkCardIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkCardIndex",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetNetworkInterfaceId(val *string) {
	if err := j.validateSetNetworkInterfaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaceId",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetPrimaryIpv6(val *string) {
	if err := j.validateSetPrimaryIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryIpv6",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetPrivateIpAddress(val *string) {
	if err := j.validateSetPrivateIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddress",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) PutConnectionTrackingSpecification(value *TfLaunchTemplate_ConnectionTrackingSpecificationProperty) {
	if err := t.validatePutConnectionTrackingSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConnectionTrackingSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) PutEnaSrdSpecification(value *TfLaunchTemplate_EnaSrdSpecificationProperty) {
	if err := t.validatePutEnaSrdSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEnaSrdSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetAssociateCarrierIpAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetAssociateCarrierIpAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetAssociatePublicIpAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetAssociatePublicIpAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetConnectionTrackingSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectionTrackingSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetDeleteOnTermination() {
	_jsii_.InvokeVoid(
		t,
		"resetDeleteOnTermination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetDeviceIndex() {
	_jsii_.InvokeVoid(
		t,
		"resetDeviceIndex",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetEnaQueueCount() {
	_jsii_.InvokeVoid(
		t,
		"resetEnaQueueCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetEnaSrdSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetEnaSrdSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetInterfaceType() {
	_jsii_.InvokeVoid(
		t,
		"resetInterfaceType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetIpv4AddressCount() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv4AddressCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetIpv4Addresses() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv4Addresses",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetIpv4PrefixCount() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv4PrefixCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetIpv4Prefixes() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv4Prefixes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetIpv6AddressCount() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv6AddressCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetIpv6Addresses() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv6Addresses",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetIpv6PrefixCount() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv6PrefixCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetIpv6Prefixes() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv6Prefixes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetNetworkCardIndex() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkCardIndex",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetNetworkInterfaceId() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkInterfaceId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetPrimaryIpv6() {
	_jsii_.InvokeVoid(
		t,
		"resetPrimaryIpv6",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetPrivateIpAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetPrivateIpAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		t,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfLaunchTemplate_NetworkInterfacesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

