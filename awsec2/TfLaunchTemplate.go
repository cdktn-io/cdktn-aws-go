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
type TfLaunchTemplate interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	BlockDeviceMappings() TfLaunchTemplate_BlockDeviceMappingsPropertyList
	// Experimental.
	BlockDeviceMappingsInput() interface{}
	// Experimental.
	CapacityReservationSpecification() TfLaunchTemplate_CapacityReservationSpecificationPropertyOutputReference
	// Experimental.
	CapacityReservationSpecificationInput() *TfLaunchTemplate_CapacityReservationSpecificationProperty
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
	CpuOptions() TfLaunchTemplate_CpuOptionsPropertyOutputReference
	// Experimental.
	CpuOptionsInput() *TfLaunchTemplate_CpuOptionsProperty
	// Experimental.
	CreditSpecification() TfLaunchTemplate_CreditSpecificationPropertyOutputReference
	// Experimental.
	CreditSpecificationInput() *TfLaunchTemplate_CreditSpecificationProperty
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
	EnclaveOptions() TfLaunchTemplate_EnclaveOptionsPropertyOutputReference
	// Experimental.
	EnclaveOptionsInput() *TfLaunchTemplate_EnclaveOptionsProperty
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HibernationOptions() TfLaunchTemplate_HibernationOptionsPropertyOutputReference
	// Experimental.
	HibernationOptionsInput() *TfLaunchTemplate_HibernationOptionsProperty
	// Experimental.
	IamInstanceProfile() TfLaunchTemplate_IamInstanceProfilePropertyOutputReference
	// Experimental.
	IamInstanceProfileInput() *TfLaunchTemplate_IamInstanceProfileProperty
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
	InstanceMarketOptions() TfLaunchTemplate_InstanceMarketOptionsPropertyOutputReference
	// Experimental.
	InstanceMarketOptionsInput() *TfLaunchTemplate_InstanceMarketOptionsProperty
	// Experimental.
	InstanceRequirements() TfLaunchTemplate_InstanceRequirementsPropertyOutputReference
	// Experimental.
	InstanceRequirementsInput() *TfLaunchTemplate_InstanceRequirementsProperty
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
	LicenseSpecification() TfLaunchTemplate_LicenseSpecificationPropertyList
	// Experimental.
	LicenseSpecificationInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaintenanceOptions() TfLaunchTemplate_MaintenanceOptionsPropertyOutputReference
	// Experimental.
	MaintenanceOptionsInput() *TfLaunchTemplate_MaintenanceOptionsProperty
	// Experimental.
	MetadataOptions() TfLaunchTemplate_MetadataOptionsPropertyOutputReference
	// Experimental.
	MetadataOptionsInput() *TfLaunchTemplate_MetadataOptionsProperty
	// Experimental.
	Monitoring() TfLaunchTemplate_MonitoringPropertyOutputReference
	// Experimental.
	MonitoringInput() *TfLaunchTemplate_MonitoringProperty
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
	NetworkInterfaces() TfLaunchTemplate_NetworkInterfacesPropertyList
	// Experimental.
	NetworkInterfacesInput() interface{}
	// Experimental.
	NetworkPerformanceOptions() TfLaunchTemplate_NetworkPerformanceOptionsPropertyOutputReference
	// Experimental.
	NetworkPerformanceOptionsInput() *TfLaunchTemplate_NetworkPerformanceOptionsProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Placement() TfLaunchTemplate_PlacementPropertyOutputReference
	// Experimental.
	PlacementInput() *TfLaunchTemplate_PlacementProperty
	// Experimental.
	PrivateDnsNameOptions() TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference
	// Experimental.
	PrivateDnsNameOptionsInput() *TfLaunchTemplate_PrivateDnsNameOptionsProperty
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
	SecondaryInterfaces() TfLaunchTemplate_SecondaryInterfacesPropertyList
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
	TagSpecifications() TfLaunchTemplate_TagSpecificationsPropertyList
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
	PutCapacityReservationSpecification(value *TfLaunchTemplate_CapacityReservationSpecificationProperty)
	// Experimental.
	PutCpuOptions(value *TfLaunchTemplate_CpuOptionsProperty)
	// Experimental.
	PutCreditSpecification(value *TfLaunchTemplate_CreditSpecificationProperty)
	// Experimental.
	PutEnclaveOptions(value *TfLaunchTemplate_EnclaveOptionsProperty)
	// Experimental.
	PutHibernationOptions(value *TfLaunchTemplate_HibernationOptionsProperty)
	// Experimental.
	PutIamInstanceProfile(value *TfLaunchTemplate_IamInstanceProfileProperty)
	// Experimental.
	PutInstanceMarketOptions(value *TfLaunchTemplate_InstanceMarketOptionsProperty)
	// Experimental.
	PutInstanceRequirements(value *TfLaunchTemplate_InstanceRequirementsProperty)
	// Experimental.
	PutLicenseSpecification(value interface{})
	// Experimental.
	PutMaintenanceOptions(value *TfLaunchTemplate_MaintenanceOptionsProperty)
	// Experimental.
	PutMetadataOptions(value *TfLaunchTemplate_MetadataOptionsProperty)
	// Experimental.
	PutMonitoring(value *TfLaunchTemplate_MonitoringProperty)
	// Experimental.
	PutNetworkInterfaces(value interface{})
	// Experimental.
	PutNetworkPerformanceOptions(value *TfLaunchTemplate_NetworkPerformanceOptionsProperty)
	// Experimental.
	PutPlacement(value *TfLaunchTemplate_PlacementProperty)
	// Experimental.
	PutPrivateDnsNameOptions(value *TfLaunchTemplate_PrivateDnsNameOptionsProperty)
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

// The jsii proxy struct for TfLaunchTemplate
type jsiiProxy_TfLaunchTemplate struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfLaunchTemplate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) BlockDeviceMappings() TfLaunchTemplate_BlockDeviceMappingsPropertyList {
	var returns TfLaunchTemplate_BlockDeviceMappingsPropertyList
	_jsii_.Get(
		j,
		"blockDeviceMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) BlockDeviceMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockDeviceMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) CapacityReservationSpecification() TfLaunchTemplate_CapacityReservationSpecificationPropertyOutputReference {
	var returns TfLaunchTemplate_CapacityReservationSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"capacityReservationSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) CapacityReservationSpecificationInput() *TfLaunchTemplate_CapacityReservationSpecificationProperty {
	var returns *TfLaunchTemplate_CapacityReservationSpecificationProperty
	_jsii_.Get(
		j,
		"capacityReservationSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) CpuOptions() TfLaunchTemplate_CpuOptionsPropertyOutputReference {
	var returns TfLaunchTemplate_CpuOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"cpuOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) CpuOptionsInput() *TfLaunchTemplate_CpuOptionsProperty {
	var returns *TfLaunchTemplate_CpuOptionsProperty
	_jsii_.Get(
		j,
		"cpuOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) CreditSpecification() TfLaunchTemplate_CreditSpecificationPropertyOutputReference {
	var returns TfLaunchTemplate_CreditSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"creditSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) CreditSpecificationInput() *TfLaunchTemplate_CreditSpecificationProperty {
	var returns *TfLaunchTemplate_CreditSpecificationProperty
	_jsii_.Get(
		j,
		"creditSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) DefaultVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) DefaultVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) DisableApiStop() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) DisableApiStopInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiStopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) DisableApiTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) DisableApiTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableApiTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) EbsOptimized() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) EbsOptimizedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) EnclaveOptions() TfLaunchTemplate_EnclaveOptionsPropertyOutputReference {
	var returns TfLaunchTemplate_EnclaveOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"enclaveOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) EnclaveOptionsInput() *TfLaunchTemplate_EnclaveOptionsProperty {
	var returns *TfLaunchTemplate_EnclaveOptionsProperty
	_jsii_.Get(
		j,
		"enclaveOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) HibernationOptions() TfLaunchTemplate_HibernationOptionsPropertyOutputReference {
	var returns TfLaunchTemplate_HibernationOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"hibernationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) HibernationOptionsInput() *TfLaunchTemplate_HibernationOptionsProperty {
	var returns *TfLaunchTemplate_HibernationOptionsProperty
	_jsii_.Get(
		j,
		"hibernationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) IamInstanceProfile() TfLaunchTemplate_IamInstanceProfilePropertyOutputReference {
	var returns TfLaunchTemplate_IamInstanceProfilePropertyOutputReference
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) IamInstanceProfileInput() *TfLaunchTemplate_IamInstanceProfileProperty {
	var returns *TfLaunchTemplate_IamInstanceProfileProperty
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) InstanceInitiatedShutdownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) InstanceInitiatedShutdownBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInitiatedShutdownBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) InstanceMarketOptions() TfLaunchTemplate_InstanceMarketOptionsPropertyOutputReference {
	var returns TfLaunchTemplate_InstanceMarketOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"instanceMarketOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) InstanceMarketOptionsInput() *TfLaunchTemplate_InstanceMarketOptionsProperty {
	var returns *TfLaunchTemplate_InstanceMarketOptionsProperty
	_jsii_.Get(
		j,
		"instanceMarketOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) InstanceRequirements() TfLaunchTemplate_InstanceRequirementsPropertyOutputReference {
	var returns TfLaunchTemplate_InstanceRequirementsPropertyOutputReference
	_jsii_.Get(
		j,
		"instanceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) InstanceRequirementsInput() *TfLaunchTemplate_InstanceRequirementsProperty {
	var returns *TfLaunchTemplate_InstanceRequirementsProperty
	_jsii_.Get(
		j,
		"instanceRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) KernelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) KernelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) LatestVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) LicenseSpecification() TfLaunchTemplate_LicenseSpecificationPropertyList {
	var returns TfLaunchTemplate_LicenseSpecificationPropertyList
	_jsii_.Get(
		j,
		"licenseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) LicenseSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"licenseSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) MaintenanceOptions() TfLaunchTemplate_MaintenanceOptionsPropertyOutputReference {
	var returns TfLaunchTemplate_MaintenanceOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"maintenanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) MaintenanceOptionsInput() *TfLaunchTemplate_MaintenanceOptionsProperty {
	var returns *TfLaunchTemplate_MaintenanceOptionsProperty
	_jsii_.Get(
		j,
		"maintenanceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) MetadataOptions() TfLaunchTemplate_MetadataOptionsPropertyOutputReference {
	var returns TfLaunchTemplate_MetadataOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) MetadataOptionsInput() *TfLaunchTemplate_MetadataOptionsProperty {
	var returns *TfLaunchTemplate_MetadataOptionsProperty
	_jsii_.Get(
		j,
		"metadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) Monitoring() TfLaunchTemplate_MonitoringPropertyOutputReference {
	var returns TfLaunchTemplate_MonitoringPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) MonitoringInput() *TfLaunchTemplate_MonitoringProperty {
	var returns *TfLaunchTemplate_MonitoringProperty
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) NetworkInterfaces() TfLaunchTemplate_NetworkInterfacesPropertyList {
	var returns TfLaunchTemplate_NetworkInterfacesPropertyList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) NetworkInterfacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) NetworkPerformanceOptions() TfLaunchTemplate_NetworkPerformanceOptionsPropertyOutputReference {
	var returns TfLaunchTemplate_NetworkPerformanceOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"networkPerformanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) NetworkPerformanceOptionsInput() *TfLaunchTemplate_NetworkPerformanceOptionsProperty {
	var returns *TfLaunchTemplate_NetworkPerformanceOptionsProperty
	_jsii_.Get(
		j,
		"networkPerformanceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) Placement() TfLaunchTemplate_PlacementPropertyOutputReference {
	var returns TfLaunchTemplate_PlacementPropertyOutputReference
	_jsii_.Get(
		j,
		"placement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) PlacementInput() *TfLaunchTemplate_PlacementProperty {
	var returns *TfLaunchTemplate_PlacementProperty
	_jsii_.Get(
		j,
		"placementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) PrivateDnsNameOptions() TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference {
	var returns TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"privateDnsNameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) PrivateDnsNameOptionsInput() *TfLaunchTemplate_PrivateDnsNameOptionsProperty {
	var returns *TfLaunchTemplate_PrivateDnsNameOptionsProperty
	_jsii_.Get(
		j,
		"privateDnsNameOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) RamDiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramDiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) RamDiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramDiskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) SecondaryInterfaces() TfLaunchTemplate_SecondaryInterfacesPropertyList {
	var returns TfLaunchTemplate_SecondaryInterfacesPropertyList
	_jsii_.Get(
		j,
		"secondaryInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) SecondaryInterfacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) SecurityGroupNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) SecurityGroupNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) TagSpecifications() TfLaunchTemplate_TagSpecificationsPropertyList {
	var returns TfLaunchTemplate_TagSpecificationsPropertyList
	_jsii_.Get(
		j,
		"tagSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) TagSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) UpdateDefaultVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateDefaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) UpdateDefaultVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateDefaultVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate) VpcSecurityGroupIdsInput() *[]*string {
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
func NewTfLaunchTemplate(scope constructs.Construct, id *string, config *TfLaunchTemplateConfig) TfLaunchTemplate {
	_init_.Initialize()

	if err := validateNewTfLaunchTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfLaunchTemplate{}

	_jsii_.Create(
		"@cdktn/aws-ec2.TfLaunchTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template aws_launch_template} Resource.
// Experimental.
func NewTfLaunchTemplate_Override(t TfLaunchTemplate, scope constructs.Construct, id *string, config *TfLaunchTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.TfLaunchTemplate",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetDefaultVersion(val *float64) {
	if err := j.validateSetDefaultVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultVersion",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetDisableApiStop(val interface{}) {
	if err := j.validateSetDisableApiStopParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiStop",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetDisableApiTermination(val interface{}) {
	if err := j.validateSetDisableApiTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableApiTermination",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetEbsOptimized(val *string) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetImageId(val *string) {
	if err := j.validateSetImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetInstanceInitiatedShutdownBehavior(val *string) {
	if err := j.validateSetInstanceInitiatedShutdownBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceInitiatedShutdownBehavior",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetKernelId(val *string) {
	if err := j.validateSetKernelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kernelId",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetRamDiskId(val *string) {
	if err := j.validateSetRamDiskIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ramDiskId",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetSecurityGroupNames(val *[]*string) {
	if err := j.validateSetSecurityGroupNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupNames",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetUpdateDefaultVersion(val interface{}) {
	if err := j.validateSetUpdateDefaultVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateDefaultVersion",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Generates CDKTN code for importing a TfLaunchTemplate resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfLaunchTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfLaunchTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.TfLaunchTemplate",
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
func TfLaunchTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfLaunchTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.TfLaunchTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfLaunchTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfLaunchTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.TfLaunchTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfLaunchTemplate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfLaunchTemplate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.TfLaunchTemplate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfLaunchTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-ec2.TfLaunchTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfLaunchTemplate) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfLaunchTemplate) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLaunchTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfLaunchTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfLaunchTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfLaunchTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfLaunchTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfLaunchTemplate) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfLaunchTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfLaunchTemplate) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLaunchTemplate) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLaunchTemplate) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfLaunchTemplate) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutBlockDeviceMappings(value interface{}) {
	if err := t.validatePutBlockDeviceMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBlockDeviceMappings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutCapacityReservationSpecification(value *TfLaunchTemplate_CapacityReservationSpecificationProperty) {
	if err := t.validatePutCapacityReservationSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCapacityReservationSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutCpuOptions(value *TfLaunchTemplate_CpuOptionsProperty) {
	if err := t.validatePutCpuOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCpuOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutCreditSpecification(value *TfLaunchTemplate_CreditSpecificationProperty) {
	if err := t.validatePutCreditSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCreditSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutEnclaveOptions(value *TfLaunchTemplate_EnclaveOptionsProperty) {
	if err := t.validatePutEnclaveOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEnclaveOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutHibernationOptions(value *TfLaunchTemplate_HibernationOptionsProperty) {
	if err := t.validatePutHibernationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHibernationOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutIamInstanceProfile(value *TfLaunchTemplate_IamInstanceProfileProperty) {
	if err := t.validatePutIamInstanceProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIamInstanceProfile",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutInstanceMarketOptions(value *TfLaunchTemplate_InstanceMarketOptionsProperty) {
	if err := t.validatePutInstanceMarketOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInstanceMarketOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutInstanceRequirements(value *TfLaunchTemplate_InstanceRequirementsProperty) {
	if err := t.validatePutInstanceRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInstanceRequirements",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutLicenseSpecification(value interface{}) {
	if err := t.validatePutLicenseSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLicenseSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutMaintenanceOptions(value *TfLaunchTemplate_MaintenanceOptionsProperty) {
	if err := t.validatePutMaintenanceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMaintenanceOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutMetadataOptions(value *TfLaunchTemplate_MetadataOptionsProperty) {
	if err := t.validatePutMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMetadataOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutMonitoring(value *TfLaunchTemplate_MonitoringProperty) {
	if err := t.validatePutMonitoringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMonitoring",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutNetworkInterfaces(value interface{}) {
	if err := t.validatePutNetworkInterfacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkInterfaces",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutNetworkPerformanceOptions(value *TfLaunchTemplate_NetworkPerformanceOptionsProperty) {
	if err := t.validatePutNetworkPerformanceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkPerformanceOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutPlacement(value *TfLaunchTemplate_PlacementProperty) {
	if err := t.validatePutPlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPlacement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutPrivateDnsNameOptions(value *TfLaunchTemplate_PrivateDnsNameOptionsProperty) {
	if err := t.validatePutPrivateDnsNameOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPrivateDnsNameOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutSecondaryInterfaces(value interface{}) {
	if err := t.validatePutSecondaryInterfacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSecondaryInterfaces",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) PutTagSpecifications(value interface{}) {
	if err := t.validatePutTagSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTagSpecifications",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetBlockDeviceMappings() {
	_jsii_.InvokeVoid(
		t,
		"resetBlockDeviceMappings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetCapacityReservationSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetCapacityReservationSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetCpuOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetCpuOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetCreditSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetCreditSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetDefaultVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetDisableApiStop() {
	_jsii_.InvokeVoid(
		t,
		"resetDisableApiStop",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetDisableApiTermination() {
	_jsii_.InvokeVoid(
		t,
		"resetDisableApiTermination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		t,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetEnclaveOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetEnclaveOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetHibernationOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetHibernationOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetIamInstanceProfile() {
	_jsii_.InvokeVoid(
		t,
		"resetIamInstanceProfile",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetImageId() {
	_jsii_.InvokeVoid(
		t,
		"resetImageId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetInstanceInitiatedShutdownBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceInitiatedShutdownBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetInstanceMarketOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceMarketOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetInstanceRequirements() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceRequirements",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetInstanceType() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetKernelId() {
	_jsii_.InvokeVoid(
		t,
		"resetKernelId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetKeyName() {
	_jsii_.InvokeVoid(
		t,
		"resetKeyName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetLicenseSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetLicenseSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetMaintenanceOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetMaintenanceOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetMetadataOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetMetadataOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetMonitoring() {
	_jsii_.InvokeVoid(
		t,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetNetworkInterfaces() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkInterfaces",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetNetworkPerformanceOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkPerformanceOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetPlacement() {
	_jsii_.InvokeVoid(
		t,
		"resetPlacement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetPrivateDnsNameOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetPrivateDnsNameOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetRamDiskId() {
	_jsii_.InvokeVoid(
		t,
		"resetRamDiskId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetSecondaryInterfaces() {
	_jsii_.InvokeVoid(
		t,
		"resetSecondaryInterfaces",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetSecurityGroupNames() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroupNames",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetTagSpecifications() {
	_jsii_.InvokeVoid(
		t,
		"resetTagSpecifications",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetUpdateDefaultVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetUpdateDefaultVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetUserData() {
	_jsii_.InvokeVoid(
		t,
		"resetUserData",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLaunchTemplate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLaunchTemplate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLaunchTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLaunchTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLaunchTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLaunchTemplate) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

