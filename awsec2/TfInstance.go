package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance aws_instance}.
// Experimental.
type TfInstance interface {
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
	CapacityReservationSpecification() TfInstance_CapacityReservationSpecificationPropertyOutputReference
	// Experimental.
	CapacityReservationSpecificationInput() *TfInstance_CapacityReservationSpecificationProperty
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
	CpuOptions() TfInstance_CpuOptionsPropertyOutputReference
	// Experimental.
	CpuOptionsInput() *TfInstance_CpuOptionsProperty
	// Experimental.
	CreditSpecification() TfInstance_CreditSpecificationPropertyOutputReference
	// Experimental.
	CreditSpecificationInput() *TfInstance_CreditSpecificationProperty
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
	EbsBlockDevice() TfInstance_EbsBlockDevicePropertyList
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
	EnclaveOptions() TfInstance_EnclaveOptionsPropertyOutputReference
	// Experimental.
	EnclaveOptionsInput() *TfInstance_EnclaveOptionsProperty
	// Experimental.
	EphemeralBlockDevice() TfInstance_EphemeralBlockDevicePropertyList
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
	InstanceLifecycle() *string
	// Experimental.
	InstanceMarketOptions() TfInstance_InstanceMarketOptionsPropertyOutputReference
	// Experimental.
	InstanceMarketOptionsInput() *TfInstance_InstanceMarketOptionsProperty
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
	LaunchTemplate() TfInstance_LaunchTemplatePropertyOutputReference
	// Experimental.
	LaunchTemplateInput() *TfInstance_LaunchTemplateProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaintenanceOptions() TfInstance_MaintenanceOptionsPropertyOutputReference
	// Experimental.
	MaintenanceOptionsInput() *TfInstance_MaintenanceOptionsProperty
	// Experimental.
	MetadataOptions() TfInstance_MetadataOptionsPropertyOutputReference
	// Experimental.
	MetadataOptionsInput() *TfInstance_MetadataOptionsProperty
	// Experimental.
	Monitoring() interface{}
	// Experimental.
	SetMonitoring(val interface{})
	// Experimental.
	MonitoringInput() interface{}
	// Experimental.
	NetworkInterface() TfInstance_NetworkInterfacePropertyList
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
	PrimaryNetworkInterface() TfInstance_PrimaryNetworkInterfacePropertyOutputReference
	// Experimental.
	PrimaryNetworkInterfaceId() *string
	// Experimental.
	PrimaryNetworkInterfaceInput() *TfInstance_PrimaryNetworkInterfaceProperty
	// Experimental.
	PrivateDns() *string
	// Experimental.
	PrivateDnsNameOptions() TfInstance_PrivateDnsNameOptionsPropertyOutputReference
	// Experimental.
	PrivateDnsNameOptionsInput() *TfInstance_PrivateDnsNameOptionsProperty
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
	RootBlockDevice() TfInstance_RootBlockDevicePropertyOutputReference
	// Experimental.
	RootBlockDeviceInput() *TfInstance_RootBlockDeviceProperty
	// Experimental.
	SecondaryNetworkInterface() TfInstance_SecondaryNetworkInterfacePropertyList
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
	SpotInstanceRequestId() *string
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
	Timeouts() TfInstance_TimeoutsPropertyOutputReference
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
	PutCapacityReservationSpecification(value *TfInstance_CapacityReservationSpecificationProperty)
	// Experimental.
	PutCpuOptions(value *TfInstance_CpuOptionsProperty)
	// Experimental.
	PutCreditSpecification(value *TfInstance_CreditSpecificationProperty)
	// Experimental.
	PutEbsBlockDevice(value interface{})
	// Experimental.
	PutEnclaveOptions(value *TfInstance_EnclaveOptionsProperty)
	// Experimental.
	PutEphemeralBlockDevice(value interface{})
	// Experimental.
	PutInstanceMarketOptions(value *TfInstance_InstanceMarketOptionsProperty)
	// Experimental.
	PutLaunchTemplate(value *TfInstance_LaunchTemplateProperty)
	// Experimental.
	PutMaintenanceOptions(value *TfInstance_MaintenanceOptionsProperty)
	// Experimental.
	PutMetadataOptions(value *TfInstance_MetadataOptionsProperty)
	// Experimental.
	PutNetworkInterface(value interface{})
	// Experimental.
	PutPrimaryNetworkInterface(value *TfInstance_PrimaryNetworkInterfaceProperty)
	// Experimental.
	PutPrivateDnsNameOptions(value *TfInstance_PrivateDnsNameOptionsProperty)
	// Experimental.
	PutRootBlockDevice(value *TfInstance_RootBlockDeviceProperty)
	// Experimental.
	PutSecondaryNetworkInterface(value interface{})
	// Experimental.
	PutTimeouts(value *TfInstance_TimeoutsProperty)
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
	ResetInstanceMarketOptions()
	// Experimental.
	ResetInstanceType()
	// Experimental.
	ResetIpv6AddressCount()
	// Experimental.
	ResetIpv6Addresses()
	// Experimental.
	ResetKeyName()
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
	ResetPrimaryNetworkInterface()
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
	ResetVolumeTags()
	// Experimental.
	ResetVpcSecurityGroupIds()
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

// The jsii proxy struct for TfInstance
type jsiiProxy_TfInstance struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfInstance) Ami() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ami",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) AmiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) AssociatePublicIpAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) AssociatePublicIpAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatePublicIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) CapacityReservationSpecification() TfInstance_CapacityReservationSpecificationPropertyOutputReference {
	var returns TfInstance_CapacityReservationSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"capacityReservationSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) CapacityReservationSpecificationInput() *TfInstance_CapacityReservationSpecificationProperty {
	var returns *TfInstance_CapacityReservationSpecificationProperty
	_jsii_.Get(
		j,
		"capacityReservationSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) CpuOptions() TfInstance_CpuOptionsPropertyOutputReference {
	var returns TfInstance_CpuOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"cpuOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) CpuOptionsInput() *TfInstance_CpuOptionsProperty {
	var returns *TfInstance_CpuOptionsProperty
	_jsii_.Get(
		j,
		"cpuOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) CreditSpecification() TfInstance_CreditSpecificationPropertyOutputReference {
	var returns TfInstance_CreditSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"creditSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) CreditSpecificationInput() *TfInstance_CreditSpecificationProperty {
	var returns *TfInstance_CreditSpecificationProperty
	_jsii_.Get(
		j,
		"creditSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) DisableApiStop() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) DisableApiStopInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) DisableApiTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) DisableApiTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) EbsBlockDevice() TfInstance_EbsBlockDevicePropertyList {
	var returns TfInstance_EbsBlockDevicePropertyList
	_jsii_.Get(
		j,
		"ebsBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) EbsBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) EnablePrimaryIpv6() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrimaryIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) EnablePrimaryIpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrimaryIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) EnclaveOptions() TfInstance_EnclaveOptionsPropertyOutputReference {
	var returns TfInstance_EnclaveOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"enclaveOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) EnclaveOptionsInput() *TfInstance_EnclaveOptionsProperty {
	var returns *TfInstance_EnclaveOptionsProperty
	_jsii_.Get(
		j,
		"enclaveOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) EphemeralBlockDevice() TfInstance_EphemeralBlockDevicePropertyList {
	var returns TfInstance_EphemeralBlockDevicePropertyList
	_jsii_.Get(
		j,
		"ephemeralBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) EphemeralBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) FetchPasswordData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchPasswordData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) FetchPasswordDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchPasswordDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Hibernation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hibernation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) HibernationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hibernationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) HostId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) HostIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) HostResourceGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostResourceGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) HostResourceGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostResourceGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) InstanceInitiatedShutdownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) InstanceInitiatedShutdownBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) InstanceLifecycle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceLifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) InstanceMarketOptions() TfInstance_InstanceMarketOptionsPropertyOutputReference {
	var returns TfInstance_InstanceMarketOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"instanceMarketOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) InstanceMarketOptionsInput() *TfInstance_InstanceMarketOptionsProperty {
	var returns *TfInstance_InstanceMarketOptionsProperty
	_jsii_.Get(
		j,
		"instanceMarketOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) InstanceState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Ipv6AddressCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Ipv6AddressCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6AddressCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Ipv6Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Ipv6AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) LaunchTemplate() TfInstance_LaunchTemplatePropertyOutputReference {
	var returns TfInstance_LaunchTemplatePropertyOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) LaunchTemplateInput() *TfInstance_LaunchTemplateProperty {
	var returns *TfInstance_LaunchTemplateProperty
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) MaintenanceOptions() TfInstance_MaintenanceOptionsPropertyOutputReference {
	var returns TfInstance_MaintenanceOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"maintenanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) MaintenanceOptionsInput() *TfInstance_MaintenanceOptionsProperty {
	var returns *TfInstance_MaintenanceOptionsProperty
	_jsii_.Get(
		j,
		"maintenanceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) MetadataOptions() TfInstance_MetadataOptionsPropertyOutputReference {
	var returns TfInstance_MetadataOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) MetadataOptionsInput() *TfInstance_MetadataOptionsProperty {
	var returns *TfInstance_MetadataOptionsProperty
	_jsii_.Get(
		j,
		"metadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Monitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) MonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) NetworkInterface() TfInstance_NetworkInterfacePropertyList {
	var returns TfInstance_NetworkInterfacePropertyList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) OutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) PasswordData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) PlacementGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) PlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) PlacementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) PlacementGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) PlacementPartitionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"placementPartitionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) PlacementPartitionNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"placementPartitionNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) PrimaryNetworkInterface() TfInstance_PrimaryNetworkInterfacePropertyOutputReference {
	var returns TfInstance_PrimaryNetworkInterfacePropertyOutputReference
	_jsii_.Get(
		j,
		"primaryNetworkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) PrimaryNetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryNetworkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) PrimaryNetworkInterfaceInput() *TfInstance_PrimaryNetworkInterfaceProperty {
	var returns *TfInstance_PrimaryNetworkInterfaceProperty
	_jsii_.Get(
		j,
		"primaryNetworkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) PrivateDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) PrivateDnsNameOptions() TfInstance_PrivateDnsNameOptionsPropertyOutputReference {
	var returns TfInstance_PrivateDnsNameOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"privateDnsNameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) PrivateDnsNameOptionsInput() *TfInstance_PrivateDnsNameOptionsProperty {
	var returns *TfInstance_PrivateDnsNameOptionsProperty
	_jsii_.Get(
		j,
		"privateDnsNameOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) PrivateIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) PublicDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) PublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) RootBlockDevice() TfInstance_RootBlockDevicePropertyOutputReference {
	var returns TfInstance_RootBlockDevicePropertyOutputReference
	_jsii_.Get(
		j,
		"rootBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) RootBlockDeviceInput() *TfInstance_RootBlockDeviceProperty {
	var returns *TfInstance_RootBlockDeviceProperty
	_jsii_.Get(
		j,
		"rootBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) SecondaryNetworkInterface() TfInstance_SecondaryNetworkInterfacePropertyList {
	var returns TfInstance_SecondaryNetworkInterfacePropertyList
	_jsii_.Get(
		j,
		"secondaryNetworkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) SecondaryNetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryNetworkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) SecondaryPrivateIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secondaryPrivateIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) SecondaryPrivateIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"secondaryPrivateIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) SourceDestCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceDestCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) SourceDestCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceDestCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) SpotInstanceRequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotInstanceRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Tenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) TenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Timeouts() TfInstance_TimeoutsPropertyOutputReference {
	var returns TfInstance_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) UserDataBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) UserDataBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) UserDataReplaceOnChange() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDataReplaceOnChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) UserDataReplaceOnChangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDataReplaceOnChangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) VolumeTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) VolumeTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"volumeTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance aws_instance} Resource.
// Experimental.
func NewTfInstance(scope constructs.Construct, id *string, config *TfInstanceConfig) TfInstance {
	_init_.Initialize()

	if err := validateNewTfInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfInstance{}

	_jsii_.Create(
		"@cdktn/aws-ec2.TfInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance aws_instance} Resource.
// Experimental.
func NewTfInstance_Override(t TfInstance, scope constructs.Construct, id *string, config *TfInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.TfInstance",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfInstance)SetAmi(val *string) {
	if err := j.validateSetAmiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ami",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetAssociatePublicIpAddress(val interface{}) {
	if err := j.validateSetAssociatePublicIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatePublicIpAddress",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetDisableApiStop(val interface{}) {
	if err := j.validateSetDisableApiStopParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiStop",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetDisableApiTermination(val interface{}) {
	if err := j.validateSetDisableApiTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiTermination",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetEnablePrimaryIpv6(val interface{}) {
	if err := j.validateSetEnablePrimaryIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrimaryIpv6",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetFetchPasswordData(val interface{}) {
	if err := j.validateSetFetchPasswordDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fetchPasswordData",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetHibernation(val interface{}) {
	if err := j.validateSetHibernationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hibernation",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetHostId(val *string) {
	if err := j.validateSetHostIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostId",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetHostResourceGroupArn(val *string) {
	if err := j.validateSetHostResourceGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostResourceGroupArn",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetInstanceInitiatedShutdownBehavior(val *string) {
	if err := j.validateSetInstanceInitiatedShutdownBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceInitiatedShutdownBehavior",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetIpv6AddressCount(val *float64) {
	if err := j.validateSetIpv6AddressCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6AddressCount",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetIpv6Addresses(val *[]*string) {
	if err := j.validateSetIpv6AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Addresses",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetMonitoring(val interface{}) {
	if err := j.validateSetMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoring",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetPlacementGroup(val *string) {
	if err := j.validateSetPlacementGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementGroup",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetPlacementGroupId(val *string) {
	if err := j.validateSetPlacementGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementGroupId",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetPlacementPartitionNumber(val *float64) {
	if err := j.validateSetPlacementPartitionNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementPartitionNumber",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetPrivateIp(val *string) {
	if err := j.validateSetPrivateIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIp",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetSecondaryPrivateIps(val *[]*string) {
	if err := j.validateSetSecondaryPrivateIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryPrivateIps",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetSourceDestCheck(val interface{}) {
	if err := j.validateSetSourceDestCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDestCheck",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetTenancy(val *string) {
	if err := j.validateSetTenancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenancy",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetUserDataBase64(val *string) {
	if err := j.validateSetUserDataBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userDataBase64",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetUserDataReplaceOnChange(val interface{}) {
	if err := j.validateSetUserDataReplaceOnChangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userDataReplaceOnChange",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetVolumeTags(val *map[string]*string) {
	if err := j.validateSetVolumeTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeTags",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Generates CDKTN code for importing a TfInstance resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.TfInstance",
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
func TfInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.TfInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.TfInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.TfInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-ec2.TfInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfInstance) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfInstance) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfInstance) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstance) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfInstance) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfInstance) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := t.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstance) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfInstance) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfInstance) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfInstance) PutCapacityReservationSpecification(value *TfInstance_CapacityReservationSpecificationProperty) {
	if err := t.validatePutCapacityReservationSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCapacityReservationSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstance) PutCpuOptions(value *TfInstance_CpuOptionsProperty) {
	if err := t.validatePutCpuOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCpuOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstance) PutCreditSpecification(value *TfInstance_CreditSpecificationProperty) {
	if err := t.validatePutCreditSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCreditSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstance) PutEbsBlockDevice(value interface{}) {
	if err := t.validatePutEbsBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEbsBlockDevice",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstance) PutEnclaveOptions(value *TfInstance_EnclaveOptionsProperty) {
	if err := t.validatePutEnclaveOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEnclaveOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstance) PutEphemeralBlockDevice(value interface{}) {
	if err := t.validatePutEphemeralBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEphemeralBlockDevice",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstance) PutInstanceMarketOptions(value *TfInstance_InstanceMarketOptionsProperty) {
	if err := t.validatePutInstanceMarketOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInstanceMarketOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstance) PutLaunchTemplate(value *TfInstance_LaunchTemplateProperty) {
	if err := t.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstance) PutMaintenanceOptions(value *TfInstance_MaintenanceOptionsProperty) {
	if err := t.validatePutMaintenanceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMaintenanceOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstance) PutMetadataOptions(value *TfInstance_MetadataOptionsProperty) {
	if err := t.validatePutMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMetadataOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstance) PutNetworkInterface(value interface{}) {
	if err := t.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstance) PutPrimaryNetworkInterface(value *TfInstance_PrimaryNetworkInterfaceProperty) {
	if err := t.validatePutPrimaryNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPrimaryNetworkInterface",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstance) PutPrivateDnsNameOptions(value *TfInstance_PrivateDnsNameOptionsProperty) {
	if err := t.validatePutPrivateDnsNameOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPrivateDnsNameOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstance) PutRootBlockDevice(value *TfInstance_RootBlockDeviceProperty) {
	if err := t.validatePutRootBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRootBlockDevice",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstance) PutSecondaryNetworkInterface(value interface{}) {
	if err := t.validatePutSecondaryNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSecondaryNetworkInterface",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstance) PutTimeouts(value *TfInstance_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstance) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfInstance) ResetAmi() {
	_jsii_.InvokeVoid(
		t,
		"resetAmi",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetAssociatePublicIpAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetAssociatePublicIpAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		t,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetCapacityReservationSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetCapacityReservationSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetCpuOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetCpuOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetCreditSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetCreditSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetDisableApiStop() {
	_jsii_.InvokeVoid(
		t,
		"resetDisableApiStop",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetDisableApiTermination() {
	_jsii_.InvokeVoid(
		t,
		"resetDisableApiTermination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetEbsBlockDevice() {
	_jsii_.InvokeVoid(
		t,
		"resetEbsBlockDevice",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		t,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetEnablePrimaryIpv6() {
	_jsii_.InvokeVoid(
		t,
		"resetEnablePrimaryIpv6",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetEnclaveOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetEnclaveOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetEphemeralBlockDevice() {
	_jsii_.InvokeVoid(
		t,
		"resetEphemeralBlockDevice",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetFetchPasswordData() {
	_jsii_.InvokeVoid(
		t,
		"resetFetchPasswordData",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		t,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetHibernation() {
	_jsii_.InvokeVoid(
		t,
		"resetHibernation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetHostId() {
	_jsii_.InvokeVoid(
		t,
		"resetHostId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetHostResourceGroupArn() {
	_jsii_.InvokeVoid(
		t,
		"resetHostResourceGroupArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		t,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetInstanceInitiatedShutdownBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceInitiatedShutdownBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetInstanceMarketOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceMarketOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetInstanceType() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetIpv6AddressCount() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv6AddressCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetIpv6Addresses() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv6Addresses",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetKeyName() {
	_jsii_.InvokeVoid(
		t,
		"resetKeyName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		t,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetMaintenanceOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetMaintenanceOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetMetadataOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetMetadataOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetMonitoring() {
	_jsii_.InvokeVoid(
		t,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetPlacementGroup() {
	_jsii_.InvokeVoid(
		t,
		"resetPlacementGroup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetPlacementGroupId() {
	_jsii_.InvokeVoid(
		t,
		"resetPlacementGroupId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetPlacementPartitionNumber() {
	_jsii_.InvokeVoid(
		t,
		"resetPlacementPartitionNumber",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetPrimaryNetworkInterface() {
	_jsii_.InvokeVoid(
		t,
		"resetPrimaryNetworkInterface",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetPrivateDnsNameOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetPrivateDnsNameOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetPrivateIp() {
	_jsii_.InvokeVoid(
		t,
		"resetPrivateIp",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetRootBlockDevice() {
	_jsii_.InvokeVoid(
		t,
		"resetRootBlockDevice",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetSecondaryNetworkInterface() {
	_jsii_.InvokeVoid(
		t,
		"resetSecondaryNetworkInterface",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetSecondaryPrivateIps() {
	_jsii_.InvokeVoid(
		t,
		"resetSecondaryPrivateIps",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetSourceDestCheck() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceDestCheck",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetSubnetId() {
	_jsii_.InvokeVoid(
		t,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetTenancy() {
	_jsii_.InvokeVoid(
		t,
		"resetTenancy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetUserData() {
	_jsii_.InvokeVoid(
		t,
		"resetUserData",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetUserDataBase64() {
	_jsii_.InvokeVoid(
		t,
		"resetUserDataBase64",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetUserDataReplaceOnChange() {
	_jsii_.InvokeVoid(
		t,
		"resetUserDataReplaceOnChange",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetVolumeTags() {
	_jsii_.InvokeVoid(
		t,
		"resetVolumeTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstance) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		t,
		"with",
		args,
		&returns,
	)

	return returns
}

