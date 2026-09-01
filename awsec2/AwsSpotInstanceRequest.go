package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request aws_spot_instance_request}.
// Experimental.
type AwsSpotInstanceRequest interface {
	cdktn.TerraformResource
	// Experimental.
	Ami() *string
	// Experimental.
	SetAmi(val *string)
	// Experimental.
	AmiInput() *string
	// Experimental.
	Arn() *string
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
	// Experimental.
	CapacityReservationSpecification() AwsSpotInstanceRequest_CapacityReservationSpecificationPropertyOutputReference
	// Experimental.
	CapacityReservationSpecificationInput() *AwsSpotInstanceRequest_CapacityReservationSpecificationProperty
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CpuOptions() AwsSpotInstanceRequest_CpuOptionsPropertyOutputReference
	// Experimental.
	CpuOptionsInput() *AwsSpotInstanceRequest_CpuOptionsProperty
	// Experimental.
	CreditSpecification() AwsSpotInstanceRequest_CreditSpecificationPropertyOutputReference
	// Experimental.
	CreditSpecificationInput() *AwsSpotInstanceRequest_CreditSpecificationProperty
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DisableApiStop() interface{}
	// Experimental.
	SetDisableApiStop(val interface{})
	// Experimental.
	DisableApiStopInput() interface{}
	// Experimental.
	DisableApiTermination() interface{}
	// Experimental.
	SetDisableApiTermination(val interface{})
	// Experimental.
	DisableApiTerminationInput() interface{}
	// Experimental.
	EbsBlockDevice() AwsSpotInstanceRequest_EbsBlockDevicePropertyList
	// Experimental.
	EbsBlockDeviceInput() interface{}
	// Experimental.
	EbsOptimized() interface{}
	// Experimental.
	SetEbsOptimized(val interface{})
	// Experimental.
	EbsOptimizedInput() interface{}
	// Experimental.
	EnablePrimaryIpv6() interface{}
	// Experimental.
	SetEnablePrimaryIpv6(val interface{})
	// Experimental.
	EnablePrimaryIpv6Input() interface{}
	// Experimental.
	EnclaveOptions() AwsSpotInstanceRequest_EnclaveOptionsPropertyOutputReference
	// Experimental.
	EnclaveOptionsInput() *AwsSpotInstanceRequest_EnclaveOptionsProperty
	// Experimental.
	EphemeralBlockDevice() AwsSpotInstanceRequest_EphemeralBlockDevicePropertyList
	// Experimental.
	EphemeralBlockDeviceInput() interface{}
	// Experimental.
	FetchPasswordData() interface{}
	// Experimental.
	SetFetchPasswordData(val interface{})
	// Experimental.
	FetchPasswordDataInput() interface{}
	// Experimental.
	ForceDestroy() interface{}
	// Experimental.
	SetForceDestroy(val interface{})
	// Experimental.
	ForceDestroyInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Hibernation() interface{}
	// Experimental.
	SetHibernation(val interface{})
	// Experimental.
	HibernationInput() interface{}
	// Experimental.
	HostId() *string
	// Experimental.
	SetHostId(val *string)
	// Experimental.
	HostIdInput() *string
	// Experimental.
	HostResourceGroupArn() *string
	// Experimental.
	SetHostResourceGroupArn(val *string)
	// Experimental.
	HostResourceGroupArnInput() *string
	// Experimental.
	IamInstanceProfile() *string
	// Experimental.
	SetIamInstanceProfile(val *string)
	// Experimental.
	IamInstanceProfileInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	InstanceInitiatedShutdownBehavior() *string
	// Experimental.
	SetInstanceInitiatedShutdownBehavior(val *string)
	// Experimental.
	InstanceInitiatedShutdownBehaviorInput() *string
	// Experimental.
	InstanceInterruptionBehavior() *string
	// Experimental.
	SetInstanceInterruptionBehavior(val *string)
	// Experimental.
	InstanceInterruptionBehaviorInput() *string
	// Experimental.
	InstanceState() *string
	// Experimental.
	InstanceType() *string
	// Experimental.
	SetInstanceType(val *string)
	// Experimental.
	InstanceTypeInput() *string
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
	KeyName() *string
	// Experimental.
	SetKeyName(val *string)
	// Experimental.
	KeyNameInput() *string
	// Experimental.
	LaunchGroup() *string
	// Experimental.
	SetLaunchGroup(val *string)
	// Experimental.
	LaunchGroupInput() *string
	// Experimental.
	LaunchTemplate() AwsSpotInstanceRequest_LaunchTemplatePropertyOutputReference
	// Experimental.
	LaunchTemplateInput() *AwsSpotInstanceRequest_LaunchTemplateProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaintenanceOptions() AwsSpotInstanceRequest_MaintenanceOptionsPropertyOutputReference
	// Experimental.
	MaintenanceOptionsInput() *AwsSpotInstanceRequest_MaintenanceOptionsProperty
	// Experimental.
	MetadataOptions() AwsSpotInstanceRequest_MetadataOptionsPropertyOutputReference
	// Experimental.
	MetadataOptionsInput() *AwsSpotInstanceRequest_MetadataOptionsProperty
	// Experimental.
	Monitoring() interface{}
	// Experimental.
	SetMonitoring(val interface{})
	// Experimental.
	MonitoringInput() interface{}
	// Experimental.
	NetworkInterface() AwsSpotInstanceRequest_NetworkInterfacePropertyList
	// Experimental.
	NetworkInterfaceInput() interface{}
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OutpostArn() *string
	// Experimental.
	PasswordData() *string
	// Experimental.
	PlacementGroup() *string
	// Experimental.
	SetPlacementGroup(val *string)
	// Experimental.
	PlacementGroupId() *string
	// Experimental.
	SetPlacementGroupId(val *string)
	// Experimental.
	PlacementGroupIdInput() *string
	// Experimental.
	PlacementGroupInput() *string
	// Experimental.
	PlacementPartitionNumber() *float64
	// Experimental.
	SetPlacementPartitionNumber(val *float64)
	// Experimental.
	PlacementPartitionNumberInput() *float64
	// Experimental.
	PrimaryNetworkInterface() AwsSpotInstanceRequest_PrimaryNetworkInterfacePropertyList
	// Experimental.
	PrimaryNetworkInterfaceId() *string
	// Experimental.
	PrivateDns() *string
	// Experimental.
	PrivateDnsNameOptions() AwsSpotInstanceRequest_PrivateDnsNameOptionsPropertyOutputReference
	// Experimental.
	PrivateDnsNameOptionsInput() *AwsSpotInstanceRequest_PrivateDnsNameOptionsProperty
	// Experimental.
	PrivateIp() *string
	// Experimental.
	SetPrivateIp(val *string)
	// Experimental.
	PrivateIpInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	PublicDns() *string
	// Experimental.
	PublicIp() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	RootBlockDevice() AwsSpotInstanceRequest_RootBlockDevicePropertyOutputReference
	// Experimental.
	RootBlockDeviceInput() *AwsSpotInstanceRequest_RootBlockDeviceProperty
	// Experimental.
	SecondaryNetworkInterface() AwsSpotInstanceRequest_SecondaryNetworkInterfacePropertyList
	// Experimental.
	SecondaryNetworkInterfaceInput() interface{}
	// Experimental.
	SecondaryPrivateIps() *[]*string
	// Experimental.
	SetSecondaryPrivateIps(val *[]*string)
	// Experimental.
	SecondaryPrivateIpsInput() *[]*string
	// Experimental.
	SecurityGroups() *[]*string
	// Experimental.
	SetSecurityGroups(val *[]*string)
	// Experimental.
	SecurityGroupsInput() *[]*string
	// Experimental.
	SourceDestCheck() interface{}
	// Experimental.
	SetSourceDestCheck(val interface{})
	// Experimental.
	SourceDestCheckInput() interface{}
	// Experimental.
	SpotBidStatus() *string
	// Experimental.
	SpotInstanceId() *string
	// Experimental.
	SpotPrice() *string
	// Experimental.
	SetSpotPrice(val *string)
	// Experimental.
	SpotPriceInput() *string
	// Experimental.
	SpotRequestState() *string
	// Experimental.
	SpotType() *string
	// Experimental.
	SetSpotType(val *string)
	// Experimental.
	SpotTypeInput() *string
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
	TagsAll() *map[string]*string
	// Experimental.
	SetTagsAll(val *map[string]*string)
	// Experimental.
	TagsAllInput() *map[string]*string
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	Tenancy() *string
	// Experimental.
	SetTenancy(val *string)
	// Experimental.
	TenancyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsSpotInstanceRequest_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	UserData() *string
	// Experimental.
	SetUserData(val *string)
	// Experimental.
	UserDataBase64() *string
	// Experimental.
	SetUserDataBase64(val *string)
	// Experimental.
	UserDataBase64Input() *string
	// Experimental.
	UserDataInput() *string
	// Experimental.
	UserDataReplaceOnChange() interface{}
	// Experimental.
	SetUserDataReplaceOnChange(val interface{})
	// Experimental.
	UserDataReplaceOnChangeInput() interface{}
	// Experimental.
	ValidFrom() *string
	// Experimental.
	SetValidFrom(val *string)
	// Experimental.
	ValidFromInput() *string
	// Experimental.
	ValidUntil() *string
	// Experimental.
	SetValidUntil(val *string)
	// Experimental.
	ValidUntilInput() *string
	// Experimental.
	VolumeTags() *map[string]*string
	// Experimental.
	SetVolumeTags(val *map[string]*string)
	// Experimental.
	VolumeTagsInput() *map[string]*string
	// Experimental.
	VpcSecurityGroupIds() *[]*string
	// Experimental.
	SetVpcSecurityGroupIds(val *[]*string)
	// Experimental.
	VpcSecurityGroupIdsInput() *[]*string
	// Experimental.
	WaitForFulfillment() interface{}
	// Experimental.
	SetWaitForFulfillment(val interface{})
	// Experimental.
	WaitForFulfillmentInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Experimental.
	PutCapacityReservationSpecification(value *AwsSpotInstanceRequest_CapacityReservationSpecificationProperty)
	// Experimental.
	PutCpuOptions(value *AwsSpotInstanceRequest_CpuOptionsProperty)
	// Experimental.
	PutCreditSpecification(value *AwsSpotInstanceRequest_CreditSpecificationProperty)
	// Experimental.
	PutEbsBlockDevice(value interface{})
	// Experimental.
	PutEnclaveOptions(value *AwsSpotInstanceRequest_EnclaveOptionsProperty)
	// Experimental.
	PutEphemeralBlockDevice(value interface{})
	// Experimental.
	PutLaunchTemplate(value *AwsSpotInstanceRequest_LaunchTemplateProperty)
	// Experimental.
	PutMaintenanceOptions(value *AwsSpotInstanceRequest_MaintenanceOptionsProperty)
	// Experimental.
	PutMetadataOptions(value *AwsSpotInstanceRequest_MetadataOptionsProperty)
	// Experimental.
	PutNetworkInterface(value interface{})
	// Experimental.
	PutPrivateDnsNameOptions(value *AwsSpotInstanceRequest_PrivateDnsNameOptionsProperty)
	// Experimental.
	PutRootBlockDevice(value *AwsSpotInstanceRequest_RootBlockDeviceProperty)
	// Experimental.
	PutSecondaryNetworkInterface(value interface{})
	// Experimental.
	PutTimeouts(value *AwsSpotInstanceRequest_TimeoutsProperty)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	// Experimental.
	ResetAmi()
	// Experimental.
	ResetAssociatePublicIpAddress()
	// Experimental.
	ResetAvailabilityZone()
	// Experimental.
	ResetCapacityReservationSpecification()
	// Experimental.
	ResetCpuOptions()
	// Experimental.
	ResetCreditSpecification()
	// Experimental.
	ResetDisableApiStop()
	// Experimental.
	ResetDisableApiTermination()
	// Experimental.
	ResetEbsBlockDevice()
	// Experimental.
	ResetEbsOptimized()
	// Experimental.
	ResetEnablePrimaryIpv6()
	// Experimental.
	ResetEnclaveOptions()
	// Experimental.
	ResetEphemeralBlockDevice()
	// Experimental.
	ResetFetchPasswordData()
	// Experimental.
	ResetForceDestroy()
	// Experimental.
	ResetHibernation()
	// Experimental.
	ResetHostId()
	// Experimental.
	ResetHostResourceGroupArn()
	// Experimental.
	ResetIamInstanceProfile()
	// Experimental.
	ResetId()
	// Experimental.
	ResetInstanceInitiatedShutdownBehavior()
	// Experimental.
	ResetInstanceInterruptionBehavior()
	// Experimental.
	ResetInstanceType()
	// Experimental.
	ResetIpv6AddressCount()
	// Experimental.
	ResetIpv6Addresses()
	// Experimental.
	ResetKeyName()
	// Experimental.
	ResetLaunchGroup()
	// Experimental.
	ResetLaunchTemplate()
	// Experimental.
	ResetMaintenanceOptions()
	// Experimental.
	ResetMetadataOptions()
	// Experimental.
	ResetMonitoring()
	// Experimental.
	ResetNetworkInterface()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPlacementGroup()
	// Experimental.
	ResetPlacementGroupId()
	// Experimental.
	ResetPlacementPartitionNumber()
	// Experimental.
	ResetPrivateDnsNameOptions()
	// Experimental.
	ResetPrivateIp()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRootBlockDevice()
	// Experimental.
	ResetSecondaryNetworkInterface()
	// Experimental.
	ResetSecondaryPrivateIps()
	// Experimental.
	ResetSecurityGroups()
	// Experimental.
	ResetSourceDestCheck()
	// Experimental.
	ResetSpotPrice()
	// Experimental.
	ResetSpotType()
	// Experimental.
	ResetSubnetId()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTenancy()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetUserData()
	// Experimental.
	ResetUserDataBase64()
	// Experimental.
	ResetUserDataReplaceOnChange()
	// Experimental.
	ResetValidFrom()
	// Experimental.
	ResetValidUntil()
	// Experimental.
	ResetVolumeTags()
	// Experimental.
	ResetVpcSecurityGroupIds()
	// Experimental.
	ResetWaitForFulfillment()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for AwsSpotInstanceRequest
type jsiiProxy_AwsSpotInstanceRequest struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Ami() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ami",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) AmiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) AssociatePublicIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) AssociatePublicIpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) CapacityReservationSpecification() AwsSpotInstanceRequest_CapacityReservationSpecificationPropertyOutputReference {
	var returns AwsSpotInstanceRequest_CapacityReservationSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"capacityReservationSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) CapacityReservationSpecificationInput() *AwsSpotInstanceRequest_CapacityReservationSpecificationProperty {
	var returns *AwsSpotInstanceRequest_CapacityReservationSpecificationProperty
	_jsii_.Get(
		j,
		"capacityReservationSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) CpuOptions() AwsSpotInstanceRequest_CpuOptionsPropertyOutputReference {
	var returns AwsSpotInstanceRequest_CpuOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"cpuOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) CpuOptionsInput() *AwsSpotInstanceRequest_CpuOptionsProperty {
	var returns *AwsSpotInstanceRequest_CpuOptionsProperty
	_jsii_.Get(
		j,
		"cpuOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) CreditSpecification() AwsSpotInstanceRequest_CreditSpecificationPropertyOutputReference {
	var returns AwsSpotInstanceRequest_CreditSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"creditSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) CreditSpecificationInput() *AwsSpotInstanceRequest_CreditSpecificationProperty {
	var returns *AwsSpotInstanceRequest_CreditSpecificationProperty
	_jsii_.Get(
		j,
		"creditSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) DisableApiStop() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) DisableApiStopInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) DisableApiTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) DisableApiTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) EbsBlockDevice() AwsSpotInstanceRequest_EbsBlockDevicePropertyList {
	var returns AwsSpotInstanceRequest_EbsBlockDevicePropertyList
	_jsii_.Get(
		j,
		"ebsBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) EbsBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) EnablePrimaryIpv6() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrimaryIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) EnablePrimaryIpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrimaryIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) EnclaveOptions() AwsSpotInstanceRequest_EnclaveOptionsPropertyOutputReference {
	var returns AwsSpotInstanceRequest_EnclaveOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"enclaveOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) EnclaveOptionsInput() *AwsSpotInstanceRequest_EnclaveOptionsProperty {
	var returns *AwsSpotInstanceRequest_EnclaveOptionsProperty
	_jsii_.Get(
		j,
		"enclaveOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) EphemeralBlockDevice() AwsSpotInstanceRequest_EphemeralBlockDevicePropertyList {
	var returns AwsSpotInstanceRequest_EphemeralBlockDevicePropertyList
	_jsii_.Get(
		j,
		"ephemeralBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) EphemeralBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) FetchPasswordData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchPasswordData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) FetchPasswordDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchPasswordDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Hibernation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hibernation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) HibernationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hibernationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) HostId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) HostIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) HostResourceGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostResourceGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) HostResourceGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostResourceGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) InstanceInitiatedShutdownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) InstanceInitiatedShutdownBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) InstanceInterruptionBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInterruptionBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) InstanceInterruptionBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInterruptionBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) InstanceState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Ipv6AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Ipv6AddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Ipv6Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Ipv6AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) LaunchGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) LaunchGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) LaunchTemplate() AwsSpotInstanceRequest_LaunchTemplatePropertyOutputReference {
	var returns AwsSpotInstanceRequest_LaunchTemplatePropertyOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) LaunchTemplateInput() *AwsSpotInstanceRequest_LaunchTemplateProperty {
	var returns *AwsSpotInstanceRequest_LaunchTemplateProperty
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) MaintenanceOptions() AwsSpotInstanceRequest_MaintenanceOptionsPropertyOutputReference {
	var returns AwsSpotInstanceRequest_MaintenanceOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"maintenanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) MaintenanceOptionsInput() *AwsSpotInstanceRequest_MaintenanceOptionsProperty {
	var returns *AwsSpotInstanceRequest_MaintenanceOptionsProperty
	_jsii_.Get(
		j,
		"maintenanceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) MetadataOptions() AwsSpotInstanceRequest_MetadataOptionsPropertyOutputReference {
	var returns AwsSpotInstanceRequest_MetadataOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) MetadataOptionsInput() *AwsSpotInstanceRequest_MetadataOptionsProperty {
	var returns *AwsSpotInstanceRequest_MetadataOptionsProperty
	_jsii_.Get(
		j,
		"metadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Monitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) MonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) NetworkInterface() AwsSpotInstanceRequest_NetworkInterfacePropertyList {
	var returns AwsSpotInstanceRequest_NetworkInterfacePropertyList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) OutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) PasswordData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) PlacementGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) PlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) PlacementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) PlacementGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) PlacementPartitionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"placementPartitionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) PlacementPartitionNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"placementPartitionNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) PrimaryNetworkInterface() AwsSpotInstanceRequest_PrimaryNetworkInterfacePropertyList {
	var returns AwsSpotInstanceRequest_PrimaryNetworkInterfacePropertyList
	_jsii_.Get(
		j,
		"primaryNetworkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) PrimaryNetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryNetworkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) PrivateDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) PrivateDnsNameOptions() AwsSpotInstanceRequest_PrivateDnsNameOptionsPropertyOutputReference {
	var returns AwsSpotInstanceRequest_PrivateDnsNameOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"privateDnsNameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) PrivateDnsNameOptionsInput() *AwsSpotInstanceRequest_PrivateDnsNameOptionsProperty {
	var returns *AwsSpotInstanceRequest_PrivateDnsNameOptionsProperty
	_jsii_.Get(
		j,
		"privateDnsNameOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) PrivateIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) PublicDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) PublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) RootBlockDevice() AwsSpotInstanceRequest_RootBlockDevicePropertyOutputReference {
	var returns AwsSpotInstanceRequest_RootBlockDevicePropertyOutputReference
	_jsii_.Get(
		j,
		"rootBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) RootBlockDeviceInput() *AwsSpotInstanceRequest_RootBlockDeviceProperty {
	var returns *AwsSpotInstanceRequest_RootBlockDeviceProperty
	_jsii_.Get(
		j,
		"rootBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) SecondaryNetworkInterface() AwsSpotInstanceRequest_SecondaryNetworkInterfacePropertyList {
	var returns AwsSpotInstanceRequest_SecondaryNetworkInterfacePropertyList
	_jsii_.Get(
		j,
		"secondaryNetworkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) SecondaryNetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryNetworkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) SecondaryPrivateIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secondaryPrivateIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) SecondaryPrivateIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secondaryPrivateIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) SourceDestCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceDestCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) SourceDestCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceDestCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) SpotBidStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotBidStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) SpotInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) SpotPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) SpotPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) SpotRequestState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotRequestState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) SpotType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) SpotTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Tenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) TenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) Timeouts() AwsSpotInstanceRequest_TimeoutsPropertyOutputReference {
	var returns AwsSpotInstanceRequest_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) UserDataBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) UserDataBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) UserDataReplaceOnChange() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDataReplaceOnChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) UserDataReplaceOnChangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDataReplaceOnChangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) ValidFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) ValidFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) ValidUntil() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntil",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) ValidUntilInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntilInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) VolumeTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) VolumeTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) WaitForFulfillment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForFulfillment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotInstanceRequest) WaitForFulfillmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForFulfillmentInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request aws_spot_instance_request} Resource.
// Experimental.
func NewAwsSpotInstanceRequest(scope constructs.Construct, id *string, config *AwsSpotInstanceRequestConfig) AwsSpotInstanceRequest {
	_init_.Initialize()

	if err := validateNewAwsSpotInstanceRequestParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSpotInstanceRequest{}

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsSpotInstanceRequest",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request aws_spot_instance_request} Resource.
// Experimental.
func NewAwsSpotInstanceRequest_Override(a AwsSpotInstanceRequest, scope constructs.Construct, id *string, config *AwsSpotInstanceRequestConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsSpotInstanceRequest",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetAmi(val *string) {
	if err := j.validateSetAmiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ami",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetAssociatePublicIpAddress(val interface{}) {
	if err := j.validateSetAssociatePublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetDisableApiStop(val interface{}) {
	if err := j.validateSetDisableApiStopParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiStop",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetDisableApiTermination(val interface{}) {
	if err := j.validateSetDisableApiTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiTermination",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetEnablePrimaryIpv6(val interface{}) {
	if err := j.validateSetEnablePrimaryIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrimaryIpv6",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetFetchPasswordData(val interface{}) {
	if err := j.validateSetFetchPasswordDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fetchPasswordData",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetHibernation(val interface{}) {
	if err := j.validateSetHibernationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hibernation",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetHostId(val *string) {
	if err := j.validateSetHostIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostId",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetHostResourceGroupArn(val *string) {
	if err := j.validateSetHostResourceGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostResourceGroupArn",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetInstanceInitiatedShutdownBehavior(val *string) {
	if err := j.validateSetInstanceInitiatedShutdownBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceInitiatedShutdownBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetInstanceInterruptionBehavior(val *string) {
	if err := j.validateSetInstanceInterruptionBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceInterruptionBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetIpv6AddressCount(val *float64) {
	if err := j.validateSetIpv6AddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6AddressCount",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetIpv6Addresses(val *[]*string) {
	if err := j.validateSetIpv6AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Addresses",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetLaunchGroup(val *string) {
	if err := j.validateSetLaunchGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchGroup",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetMonitoring(val interface{}) {
	if err := j.validateSetMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoring",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetPlacementGroup(val *string) {
	if err := j.validateSetPlacementGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementGroup",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetPlacementGroupId(val *string) {
	if err := j.validateSetPlacementGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementGroupId",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetPlacementPartitionNumber(val *float64) {
	if err := j.validateSetPlacementPartitionNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementPartitionNumber",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetPrivateIp(val *string) {
	if err := j.validateSetPrivateIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIp",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetSecondaryPrivateIps(val *[]*string) {
	if err := j.validateSetSecondaryPrivateIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryPrivateIps",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetSourceDestCheck(val interface{}) {
	if err := j.validateSetSourceDestCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDestCheck",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetSpotPrice(val *string) {
	if err := j.validateSetSpotPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotPrice",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetSpotType(val *string) {
	if err := j.validateSetSpotTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotType",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetTenancy(val *string) {
	if err := j.validateSetTenancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenancy",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetUserDataBase64(val *string) {
	if err := j.validateSetUserDataBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userDataBase64",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetUserDataReplaceOnChange(val interface{}) {
	if err := j.validateSetUserDataReplaceOnChangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userDataReplaceOnChange",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetValidFrom(val *string) {
	if err := j.validateSetValidFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validFrom",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetValidUntil(val *string) {
	if err := j.validateSetValidUntilParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validUntil",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetVolumeTags(val *map[string]*string) {
	if err := j.validateSetVolumeTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeTags",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AwsSpotInstanceRequest)SetWaitForFulfillment(val interface{}) {
	if err := j.validateSetWaitForFulfillmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForFulfillment",
		val,
	)
}

// Generates CDKTN code for importing a AwsSpotInstanceRequest resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsSpotInstanceRequest_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsSpotInstanceRequest_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsSpotInstanceRequest",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func AwsSpotInstanceRequest_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSpotInstanceRequest_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsSpotInstanceRequest",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsSpotInstanceRequest_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSpotInstanceRequest_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsSpotInstanceRequest",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsSpotInstanceRequest_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSpotInstanceRequest_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsSpotInstanceRequest",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsSpotInstanceRequest_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-ec2.AwsSpotInstanceRequest",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsSpotInstanceRequest) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSpotInstanceRequest) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSpotInstanceRequest) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSpotInstanceRequest) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSpotInstanceRequest) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSpotInstanceRequest) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSpotInstanceRequest) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSpotInstanceRequest) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSpotInstanceRequest) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSpotInstanceRequest) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSpotInstanceRequest) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := a.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpotInstanceRequest) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) PutCapacityReservationSpecification(value *AwsSpotInstanceRequest_CapacityReservationSpecificationProperty) {
	if err := a.validatePutCapacityReservationSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCapacityReservationSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) PutCpuOptions(value *AwsSpotInstanceRequest_CpuOptionsProperty) {
	if err := a.validatePutCpuOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCpuOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) PutCreditSpecification(value *AwsSpotInstanceRequest_CreditSpecificationProperty) {
	if err := a.validatePutCreditSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCreditSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) PutEbsBlockDevice(value interface{}) {
	if err := a.validatePutEbsBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEbsBlockDevice",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) PutEnclaveOptions(value *AwsSpotInstanceRequest_EnclaveOptionsProperty) {
	if err := a.validatePutEnclaveOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnclaveOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) PutEphemeralBlockDevice(value interface{}) {
	if err := a.validatePutEphemeralBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEphemeralBlockDevice",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) PutLaunchTemplate(value *AwsSpotInstanceRequest_LaunchTemplateProperty) {
	if err := a.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) PutMaintenanceOptions(value *AwsSpotInstanceRequest_MaintenanceOptionsProperty) {
	if err := a.validatePutMaintenanceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMaintenanceOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) PutMetadataOptions(value *AwsSpotInstanceRequest_MetadataOptionsProperty) {
	if err := a.validatePutMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetadataOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) PutNetworkInterface(value interface{}) {
	if err := a.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) PutPrivateDnsNameOptions(value *AwsSpotInstanceRequest_PrivateDnsNameOptionsProperty) {
	if err := a.validatePutPrivateDnsNameOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPrivateDnsNameOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) PutRootBlockDevice(value *AwsSpotInstanceRequest_RootBlockDeviceProperty) {
	if err := a.validatePutRootBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRootBlockDevice",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) PutSecondaryNetworkInterface(value interface{}) {
	if err := a.validatePutSecondaryNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecondaryNetworkInterface",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) PutTimeouts(value *AwsSpotInstanceRequest_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetAmi() {
	_jsii_.InvokeVoid(
		a,
		"resetAmi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetAssociatePublicIpAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetAssociatePublicIpAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetCapacityReservationSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCapacityReservationSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetCpuOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetCpuOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetCreditSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCreditSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetDisableApiStop() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableApiStop",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetDisableApiTermination() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableApiTermination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetEbsBlockDevice() {
	_jsii_.InvokeVoid(
		a,
		"resetEbsBlockDevice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		a,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetEnablePrimaryIpv6() {
	_jsii_.InvokeVoid(
		a,
		"resetEnablePrimaryIpv6",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetEnclaveOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetEnclaveOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetEphemeralBlockDevice() {
	_jsii_.InvokeVoid(
		a,
		"resetEphemeralBlockDevice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetFetchPasswordData() {
	_jsii_.InvokeVoid(
		a,
		"resetFetchPasswordData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		a,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetHibernation() {
	_jsii_.InvokeVoid(
		a,
		"resetHibernation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetHostId() {
	_jsii_.InvokeVoid(
		a,
		"resetHostId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetHostResourceGroupArn() {
	_jsii_.InvokeVoid(
		a,
		"resetHostResourceGroupArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		a,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetInstanceInitiatedShutdownBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceInitiatedShutdownBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetInstanceInterruptionBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceInterruptionBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetInstanceType() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetIpv6AddressCount() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv6AddressCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetIpv6Addresses() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv6Addresses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetKeyName() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetLaunchGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetMaintenanceOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetMaintenanceOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetMetadataOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadataOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetMonitoring() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetPlacementGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetPlacementGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetPlacementGroupId() {
	_jsii_.InvokeVoid(
		a,
		"resetPlacementGroupId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetPlacementPartitionNumber() {
	_jsii_.InvokeVoid(
		a,
		"resetPlacementPartitionNumber",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetPrivateDnsNameOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateDnsNameOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetPrivateIp() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateIp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetRootBlockDevice() {
	_jsii_.InvokeVoid(
		a,
		"resetRootBlockDevice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetSecondaryNetworkInterface() {
	_jsii_.InvokeVoid(
		a,
		"resetSecondaryNetworkInterface",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetSecondaryPrivateIps() {
	_jsii_.InvokeVoid(
		a,
		"resetSecondaryPrivateIps",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetSourceDestCheck() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceDestCheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetSpotPrice() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotPrice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetSpotType() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetSubnetId() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetTenancy() {
	_jsii_.InvokeVoid(
		a,
		"resetTenancy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetUserData() {
	_jsii_.InvokeVoid(
		a,
		"resetUserData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetUserDataBase64() {
	_jsii_.InvokeVoid(
		a,
		"resetUserDataBase64",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetUserDataReplaceOnChange() {
	_jsii_.InvokeVoid(
		a,
		"resetUserDataReplaceOnChange",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetValidFrom() {
	_jsii_.InvokeVoid(
		a,
		"resetValidFrom",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetValidUntil() {
	_jsii_.InvokeVoid(
		a,
		"resetValidUntil",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetVolumeTags() {
	_jsii_.InvokeVoid(
		a,
		"resetVolumeTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ResetWaitForFulfillment() {
	_jsii_.InvokeVoid(
		a,
		"resetWaitForFulfillment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotInstanceRequest) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpotInstanceRequest) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpotInstanceRequest) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpotInstanceRequest) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

