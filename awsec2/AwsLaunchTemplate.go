package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template aws_launch_template}.
// Experimental.
type AwsLaunchTemplate interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	BlockDeviceMappings() AwsLaunchTemplate_BlockDeviceMappingsPropertyList
	// Experimental.
	BlockDeviceMappingsInput() interface{}
	// Experimental.
	CapacityReservationSpecification() AwsLaunchTemplate_CapacityReservationSpecificationPropertyOutputReference
	// Experimental.
	CapacityReservationSpecificationInput() *AwsLaunchTemplate_CapacityReservationSpecificationProperty
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
	CpuOptions() AwsLaunchTemplate_CpuOptionsPropertyOutputReference
	// Experimental.
	CpuOptionsInput() *AwsLaunchTemplate_CpuOptionsProperty
	// Experimental.
	CreditSpecification() AwsLaunchTemplate_CreditSpecificationPropertyOutputReference
	// Experimental.
	CreditSpecificationInput() *AwsLaunchTemplate_CreditSpecificationProperty
	// Experimental.
	DefaultVersion() *float64
	// Experimental.
	SetDefaultVersion(val *float64)
	// Experimental.
	DefaultVersionInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
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
	EbsOptimized() *string
	// Experimental.
	SetEbsOptimized(val *string)
	// Experimental.
	EbsOptimizedInput() *string
	// Experimental.
	EnclaveOptions() AwsLaunchTemplate_EnclaveOptionsPropertyOutputReference
	// Experimental.
	EnclaveOptionsInput() *AwsLaunchTemplate_EnclaveOptionsProperty
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HibernationOptions() AwsLaunchTemplate_HibernationOptionsPropertyOutputReference
	// Experimental.
	HibernationOptionsInput() *AwsLaunchTemplate_HibernationOptionsProperty
	// Experimental.
	IamInstanceProfile() AwsLaunchTemplate_IamInstanceProfilePropertyOutputReference
	// Experimental.
	IamInstanceProfileInput() *AwsLaunchTemplate_IamInstanceProfileProperty
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	ImageId() *string
	// Experimental.
	SetImageId(val *string)
	// Experimental.
	ImageIdInput() *string
	// Experimental.
	InstanceInitiatedShutdownBehavior() *string
	// Experimental.
	SetInstanceInitiatedShutdownBehavior(val *string)
	// Experimental.
	InstanceInitiatedShutdownBehaviorInput() *string
	// Experimental.
	InstanceMarketOptions() AwsLaunchTemplate_InstanceMarketOptionsPropertyOutputReference
	// Experimental.
	InstanceMarketOptionsInput() *AwsLaunchTemplate_InstanceMarketOptionsProperty
	// Experimental.
	InstanceRequirements() AwsLaunchTemplate_InstanceRequirementsPropertyOutputReference
	// Experimental.
	InstanceRequirementsInput() *AwsLaunchTemplate_InstanceRequirementsProperty
	// Experimental.
	InstanceType() *string
	// Experimental.
	SetInstanceType(val *string)
	// Experimental.
	InstanceTypeInput() *string
	// Experimental.
	KernelId() *string
	// Experimental.
	SetKernelId(val *string)
	// Experimental.
	KernelIdInput() *string
	// Experimental.
	KeyName() *string
	// Experimental.
	SetKeyName(val *string)
	// Experimental.
	KeyNameInput() *string
	// Experimental.
	LatestVersion() *float64
	// Experimental.
	LicenseSpecification() AwsLaunchTemplate_LicenseSpecificationPropertyList
	// Experimental.
	LicenseSpecificationInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaintenanceOptions() AwsLaunchTemplate_MaintenanceOptionsPropertyOutputReference
	// Experimental.
	MaintenanceOptionsInput() *AwsLaunchTemplate_MaintenanceOptionsProperty
	// Experimental.
	MetadataOptions() AwsLaunchTemplate_MetadataOptionsPropertyOutputReference
	// Experimental.
	MetadataOptionsInput() *AwsLaunchTemplate_MetadataOptionsProperty
	// Experimental.
	Monitoring() AwsLaunchTemplate_MonitoringPropertyOutputReference
	// Experimental.
	MonitoringInput() *AwsLaunchTemplate_MonitoringProperty
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NamePrefix() *string
	// Experimental.
	SetNamePrefix(val *string)
	// Experimental.
	NamePrefixInput() *string
	// Experimental.
	NetworkInterfaces() AwsLaunchTemplate_NetworkInterfacesPropertyList
	// Experimental.
	NetworkInterfacesInput() interface{}
	// Experimental.
	NetworkPerformanceOptions() AwsLaunchTemplate_NetworkPerformanceOptionsPropertyOutputReference
	// Experimental.
	NetworkPerformanceOptionsInput() *AwsLaunchTemplate_NetworkPerformanceOptionsProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Placement() AwsLaunchTemplate_PlacementPropertyOutputReference
	// Experimental.
	PlacementInput() *AwsLaunchTemplate_PlacementProperty
	// Experimental.
	PrivateDnsNameOptions() AwsLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference
	// Experimental.
	PrivateDnsNameOptionsInput() *AwsLaunchTemplate_PrivateDnsNameOptionsProperty
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RamDiskId() *string
	// Experimental.
	SetRamDiskId(val *string)
	// Experimental.
	RamDiskIdInput() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	SecondaryInterfaces() AwsLaunchTemplate_SecondaryInterfacesPropertyList
	// Experimental.
	SecondaryInterfacesInput() interface{}
	// Experimental.
	SecurityGroupNames() *[]*string
	// Experimental.
	SetSecurityGroupNames(val *[]*string)
	// Experimental.
	SecurityGroupNamesInput() *[]*string
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
	TagSpecifications() AwsLaunchTemplate_TagSpecificationsPropertyList
	// Experimental.
	TagSpecificationsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	UpdateDefaultVersion() interface{}
	// Experimental.
	SetUpdateDefaultVersion(val interface{})
	// Experimental.
	UpdateDefaultVersionInput() interface{}
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
	PutBlockDeviceMappings(value interface{})
	// Experimental.
	PutCapacityReservationSpecification(value *AwsLaunchTemplate_CapacityReservationSpecificationProperty)
	// Experimental.
	PutCpuOptions(value *AwsLaunchTemplate_CpuOptionsProperty)
	// Experimental.
	PutCreditSpecification(value *AwsLaunchTemplate_CreditSpecificationProperty)
	// Experimental.
	PutEnclaveOptions(value *AwsLaunchTemplate_EnclaveOptionsProperty)
	// Experimental.
	PutHibernationOptions(value *AwsLaunchTemplate_HibernationOptionsProperty)
	// Experimental.
	PutIamInstanceProfile(value *AwsLaunchTemplate_IamInstanceProfileProperty)
	// Experimental.
	PutInstanceMarketOptions(value *AwsLaunchTemplate_InstanceMarketOptionsProperty)
	// Experimental.
	PutInstanceRequirements(value *AwsLaunchTemplate_InstanceRequirementsProperty)
	// Experimental.
	PutLicenseSpecification(value interface{})
	// Experimental.
	PutMaintenanceOptions(value *AwsLaunchTemplate_MaintenanceOptionsProperty)
	// Experimental.
	PutMetadataOptions(value *AwsLaunchTemplate_MetadataOptionsProperty)
	// Experimental.
	PutMonitoring(value *AwsLaunchTemplate_MonitoringProperty)
	// Experimental.
	PutNetworkInterfaces(value interface{})
	// Experimental.
	PutNetworkPerformanceOptions(value *AwsLaunchTemplate_NetworkPerformanceOptionsProperty)
	// Experimental.
	PutPlacement(value *AwsLaunchTemplate_PlacementProperty)
	// Experimental.
	PutPrivateDnsNameOptions(value *AwsLaunchTemplate_PrivateDnsNameOptionsProperty)
	// Experimental.
	PutSecondaryInterfaces(value interface{})
	// Experimental.
	PutTagSpecifications(value interface{})
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
	ResetBlockDeviceMappings()
	// Experimental.
	ResetCapacityReservationSpecification()
	// Experimental.
	ResetCpuOptions()
	// Experimental.
	ResetCreditSpecification()
	// Experimental.
	ResetDefaultVersion()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetDisableApiStop()
	// Experimental.
	ResetDisableApiTermination()
	// Experimental.
	ResetEbsOptimized()
	// Experimental.
	ResetEnclaveOptions()
	// Experimental.
	ResetHibernationOptions()
	// Experimental.
	ResetIamInstanceProfile()
	// Experimental.
	ResetId()
	// Experimental.
	ResetImageId()
	// Experimental.
	ResetInstanceInitiatedShutdownBehavior()
	// Experimental.
	ResetInstanceMarketOptions()
	// Experimental.
	ResetInstanceRequirements()
	// Experimental.
	ResetInstanceType()
	// Experimental.
	ResetKernelId()
	// Experimental.
	ResetKeyName()
	// Experimental.
	ResetLicenseSpecification()
	// Experimental.
	ResetMaintenanceOptions()
	// Experimental.
	ResetMetadataOptions()
	// Experimental.
	ResetMonitoring()
	// Experimental.
	ResetName()
	// Experimental.
	ResetNamePrefix()
	// Experimental.
	ResetNetworkInterfaces()
	// Experimental.
	ResetNetworkPerformanceOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPlacement()
	// Experimental.
	ResetPrivateDnsNameOptions()
	// Experimental.
	ResetRamDiskId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSecondaryInterfaces()
	// Experimental.
	ResetSecurityGroupNames()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTagSpecifications()
	// Experimental.
	ResetUpdateDefaultVersion()
	// Experimental.
	ResetUserData()
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

// The jsii proxy struct for AwsLaunchTemplate
type jsiiProxy_AwsLaunchTemplate struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsLaunchTemplate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) BlockDeviceMappings() AwsLaunchTemplate_BlockDeviceMappingsPropertyList {
	var returns AwsLaunchTemplate_BlockDeviceMappingsPropertyList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) BlockDeviceMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) CapacityReservationSpecification() AwsLaunchTemplate_CapacityReservationSpecificationPropertyOutputReference {
	var returns AwsLaunchTemplate_CapacityReservationSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"capacityReservationSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) CapacityReservationSpecificationInput() *AwsLaunchTemplate_CapacityReservationSpecificationProperty {
	var returns *AwsLaunchTemplate_CapacityReservationSpecificationProperty
	_jsii_.Get(
		j,
		"capacityReservationSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) CpuOptions() AwsLaunchTemplate_CpuOptionsPropertyOutputReference {
	var returns AwsLaunchTemplate_CpuOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"cpuOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) CpuOptionsInput() *AwsLaunchTemplate_CpuOptionsProperty {
	var returns *AwsLaunchTemplate_CpuOptionsProperty
	_jsii_.Get(
		j,
		"cpuOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) CreditSpecification() AwsLaunchTemplate_CreditSpecificationPropertyOutputReference {
	var returns AwsLaunchTemplate_CreditSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"creditSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) CreditSpecificationInput() *AwsLaunchTemplate_CreditSpecificationProperty {
	var returns *AwsLaunchTemplate_CreditSpecificationProperty
	_jsii_.Get(
		j,
		"creditSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) DefaultVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) DefaultVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) DisableApiStop() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) DisableApiStopInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) DisableApiTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) DisableApiTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) EbsOptimized() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) EbsOptimizedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) EnclaveOptions() AwsLaunchTemplate_EnclaveOptionsPropertyOutputReference {
	var returns AwsLaunchTemplate_EnclaveOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"enclaveOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) EnclaveOptionsInput() *AwsLaunchTemplate_EnclaveOptionsProperty {
	var returns *AwsLaunchTemplate_EnclaveOptionsProperty
	_jsii_.Get(
		j,
		"enclaveOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) HibernationOptions() AwsLaunchTemplate_HibernationOptionsPropertyOutputReference {
	var returns AwsLaunchTemplate_HibernationOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"hibernationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) HibernationOptionsInput() *AwsLaunchTemplate_HibernationOptionsProperty {
	var returns *AwsLaunchTemplate_HibernationOptionsProperty
	_jsii_.Get(
		j,
		"hibernationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) IamInstanceProfile() AwsLaunchTemplate_IamInstanceProfilePropertyOutputReference {
	var returns AwsLaunchTemplate_IamInstanceProfilePropertyOutputReference
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) IamInstanceProfileInput() *AwsLaunchTemplate_IamInstanceProfileProperty {
	var returns *AwsLaunchTemplate_IamInstanceProfileProperty
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) InstanceInitiatedShutdownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) InstanceInitiatedShutdownBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) InstanceMarketOptions() AwsLaunchTemplate_InstanceMarketOptionsPropertyOutputReference {
	var returns AwsLaunchTemplate_InstanceMarketOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"instanceMarketOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) InstanceMarketOptionsInput() *AwsLaunchTemplate_InstanceMarketOptionsProperty {
	var returns *AwsLaunchTemplate_InstanceMarketOptionsProperty
	_jsii_.Get(
		j,
		"instanceMarketOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) InstanceRequirements() AwsLaunchTemplate_InstanceRequirementsPropertyOutputReference {
	var returns AwsLaunchTemplate_InstanceRequirementsPropertyOutputReference
	_jsii_.Get(
		j,
		"instanceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) InstanceRequirementsInput() *AwsLaunchTemplate_InstanceRequirementsProperty {
	var returns *AwsLaunchTemplate_InstanceRequirementsProperty
	_jsii_.Get(
		j,
		"instanceRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) KernelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) KernelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) LatestVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) LicenseSpecification() AwsLaunchTemplate_LicenseSpecificationPropertyList {
	var returns AwsLaunchTemplate_LicenseSpecificationPropertyList
	_jsii_.Get(
		j,
		"licenseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) LicenseSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"licenseSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) MaintenanceOptions() AwsLaunchTemplate_MaintenanceOptionsPropertyOutputReference {
	var returns AwsLaunchTemplate_MaintenanceOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"maintenanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) MaintenanceOptionsInput() *AwsLaunchTemplate_MaintenanceOptionsProperty {
	var returns *AwsLaunchTemplate_MaintenanceOptionsProperty
	_jsii_.Get(
		j,
		"maintenanceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) MetadataOptions() AwsLaunchTemplate_MetadataOptionsPropertyOutputReference {
	var returns AwsLaunchTemplate_MetadataOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) MetadataOptionsInput() *AwsLaunchTemplate_MetadataOptionsProperty {
	var returns *AwsLaunchTemplate_MetadataOptionsProperty
	_jsii_.Get(
		j,
		"metadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) Monitoring() AwsLaunchTemplate_MonitoringPropertyOutputReference {
	var returns AwsLaunchTemplate_MonitoringPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) MonitoringInput() *AwsLaunchTemplate_MonitoringProperty {
	var returns *AwsLaunchTemplate_MonitoringProperty
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) NetworkInterfaces() AwsLaunchTemplate_NetworkInterfacesPropertyList {
	var returns AwsLaunchTemplate_NetworkInterfacesPropertyList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) NetworkInterfacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) NetworkPerformanceOptions() AwsLaunchTemplate_NetworkPerformanceOptionsPropertyOutputReference {
	var returns AwsLaunchTemplate_NetworkPerformanceOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"networkPerformanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) NetworkPerformanceOptionsInput() *AwsLaunchTemplate_NetworkPerformanceOptionsProperty {
	var returns *AwsLaunchTemplate_NetworkPerformanceOptionsProperty
	_jsii_.Get(
		j,
		"networkPerformanceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) Placement() AwsLaunchTemplate_PlacementPropertyOutputReference {
	var returns AwsLaunchTemplate_PlacementPropertyOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) PlacementInput() *AwsLaunchTemplate_PlacementProperty {
	var returns *AwsLaunchTemplate_PlacementProperty
	_jsii_.Get(
		j,
		"placementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) PrivateDnsNameOptions() AwsLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference {
	var returns AwsLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"privateDnsNameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) PrivateDnsNameOptionsInput() *AwsLaunchTemplate_PrivateDnsNameOptionsProperty {
	var returns *AwsLaunchTemplate_PrivateDnsNameOptionsProperty
	_jsii_.Get(
		j,
		"privateDnsNameOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) RamDiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramDiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) RamDiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramDiskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) SecondaryInterfaces() AwsLaunchTemplate_SecondaryInterfacesPropertyList {
	var returns AwsLaunchTemplate_SecondaryInterfacesPropertyList
	_jsii_.Get(
		j,
		"secondaryInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) SecondaryInterfacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) SecurityGroupNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) SecurityGroupNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) TagSpecifications() AwsLaunchTemplate_TagSpecificationsPropertyList {
	var returns AwsLaunchTemplate_TagSpecificationsPropertyList
	_jsii_.Get(
		j,
		"tagSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) TagSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) UpdateDefaultVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateDefaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) UpdateDefaultVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateDefaultVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLaunchTemplate) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template aws_launch_template} Resource.
// Experimental.
func NewAwsLaunchTemplate(scope constructs.Construct, id *string, config *AwsLaunchTemplateConfig) AwsLaunchTemplate {
	_init_.Initialize()

	if err := validateNewAwsLaunchTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLaunchTemplate{}

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsLaunchTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template aws_launch_template} Resource.
// Experimental.
func NewAwsLaunchTemplate_Override(a AwsLaunchTemplate, scope constructs.Construct, id *string, config *AwsLaunchTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsLaunchTemplate",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetDefaultVersion(val *float64) {
	if err := j.validateSetDefaultVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultVersion",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetDisableApiStop(val interface{}) {
	if err := j.validateSetDisableApiStopParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiStop",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetDisableApiTermination(val interface{}) {
	if err := j.validateSetDisableApiTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiTermination",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetEbsOptimized(val *string) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetImageId(val *string) {
	if err := j.validateSetImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetInstanceInitiatedShutdownBehavior(val *string) {
	if err := j.validateSetInstanceInitiatedShutdownBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceInitiatedShutdownBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetKernelId(val *string) {
	if err := j.validateSetKernelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kernelId",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetRamDiskId(val *string) {
	if err := j.validateSetRamDiskIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ramDiskId",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetSecurityGroupNames(val *[]*string) {
	if err := j.validateSetSecurityGroupNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupNames",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetUpdateDefaultVersion(val interface{}) {
	if err := j.validateSetUpdateDefaultVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateDefaultVersion",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_AwsLaunchTemplate)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Generates CDKTN code for importing a AwsLaunchTemplate resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsLaunchTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsLaunchTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsLaunchTemplate",
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
func AwsLaunchTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsLaunchTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsLaunchTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsLaunchTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsLaunchTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsLaunchTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsLaunchTemplate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsLaunchTemplate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsLaunchTemplate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsLaunchTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-ec2.AwsLaunchTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLaunchTemplate) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLaunchTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLaunchTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLaunchTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLaunchTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLaunchTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLaunchTemplate) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLaunchTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLaunchTemplate) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLaunchTemplate) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsLaunchTemplate) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutBlockDeviceMappings(value interface{}) {
	if err := a.validatePutBlockDeviceMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBlockDeviceMappings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutCapacityReservationSpecification(value *AwsLaunchTemplate_CapacityReservationSpecificationProperty) {
	if err := a.validatePutCapacityReservationSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCapacityReservationSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutCpuOptions(value *AwsLaunchTemplate_CpuOptionsProperty) {
	if err := a.validatePutCpuOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCpuOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutCreditSpecification(value *AwsLaunchTemplate_CreditSpecificationProperty) {
	if err := a.validatePutCreditSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCreditSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutEnclaveOptions(value *AwsLaunchTemplate_EnclaveOptionsProperty) {
	if err := a.validatePutEnclaveOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnclaveOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutHibernationOptions(value *AwsLaunchTemplate_HibernationOptionsProperty) {
	if err := a.validatePutHibernationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHibernationOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutIamInstanceProfile(value *AwsLaunchTemplate_IamInstanceProfileProperty) {
	if err := a.validatePutIamInstanceProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIamInstanceProfile",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutInstanceMarketOptions(value *AwsLaunchTemplate_InstanceMarketOptionsProperty) {
	if err := a.validatePutInstanceMarketOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInstanceMarketOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutInstanceRequirements(value *AwsLaunchTemplate_InstanceRequirementsProperty) {
	if err := a.validatePutInstanceRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInstanceRequirements",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutLicenseSpecification(value interface{}) {
	if err := a.validatePutLicenseSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLicenseSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutMaintenanceOptions(value *AwsLaunchTemplate_MaintenanceOptionsProperty) {
	if err := a.validatePutMaintenanceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMaintenanceOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutMetadataOptions(value *AwsLaunchTemplate_MetadataOptionsProperty) {
	if err := a.validatePutMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetadataOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutMonitoring(value *AwsLaunchTemplate_MonitoringProperty) {
	if err := a.validatePutMonitoringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonitoring",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutNetworkInterfaces(value interface{}) {
	if err := a.validatePutNetworkInterfacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkInterfaces",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutNetworkPerformanceOptions(value *AwsLaunchTemplate_NetworkPerformanceOptionsProperty) {
	if err := a.validatePutNetworkPerformanceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkPerformanceOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutPlacement(value *AwsLaunchTemplate_PlacementProperty) {
	if err := a.validatePutPlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPlacement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutPrivateDnsNameOptions(value *AwsLaunchTemplate_PrivateDnsNameOptionsProperty) {
	if err := a.validatePutPrivateDnsNameOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPrivateDnsNameOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutSecondaryInterfaces(value interface{}) {
	if err := a.validatePutSecondaryInterfacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecondaryInterfaces",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) PutTagSpecifications(value interface{}) {
	if err := a.validatePutTagSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTagSpecifications",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetBlockDeviceMappings() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockDeviceMappings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetCapacityReservationSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCapacityReservationSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetCpuOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetCpuOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetCreditSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCreditSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetDefaultVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetDisableApiStop() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableApiStop",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetDisableApiTermination() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableApiTermination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		a,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetEnclaveOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetEnclaveOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetHibernationOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetHibernationOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		a,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetImageId() {
	_jsii_.InvokeVoid(
		a,
		"resetImageId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetInstanceInitiatedShutdownBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceInitiatedShutdownBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetInstanceMarketOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceMarketOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetInstanceRequirements() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceRequirements",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetInstanceType() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetKernelId() {
	_jsii_.InvokeVoid(
		a,
		"resetKernelId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetKeyName() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetLicenseSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetLicenseSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetMaintenanceOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetMaintenanceOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetMetadataOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadataOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetMonitoring() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetNetworkInterfaces() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkInterfaces",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetNetworkPerformanceOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkPerformanceOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetPlacement() {
	_jsii_.InvokeVoid(
		a,
		"resetPlacement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetPrivateDnsNameOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateDnsNameOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetRamDiskId() {
	_jsii_.InvokeVoid(
		a,
		"resetRamDiskId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetSecondaryInterfaces() {
	_jsii_.InvokeVoid(
		a,
		"resetSecondaryInterfaces",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetSecurityGroupNames() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroupNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetTagSpecifications() {
	_jsii_.InvokeVoid(
		a,
		"resetTagSpecifications",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetUpdateDefaultVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetUpdateDefaultVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetUserData() {
	_jsii_.InvokeVoid(
		a,
		"resetUserData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLaunchTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLaunchTemplate) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

