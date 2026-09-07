package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ec2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/ec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami aws_ami}.
// Experimental.
type AwsAmi interface {
	cdktn.TerraformResource
	// Experimental.
	Architecture() *string
	// Experimental.
	SetArchitecture(val *string)
	// Experimental.
	ArchitectureInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	BootMode() *string
	// Experimental.
	SetBootMode(val *string)
	// Experimental.
	BootModeInput() *string
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
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DeprecationTime() *string
	// Experimental.
	SetDeprecationTime(val *string)
	// Experimental.
	DeprecationTimeInput() *string
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	EbsBlockDevice() AwsAmi_EbsBlockDevicePropertyList
	// Experimental.
	EbsBlockDeviceInput() interface{}
	// Experimental.
	EnaSupport() interface{}
	// Experimental.
	SetEnaSupport(val interface{})
	// Experimental.
	EnaSupportInput() interface{}
	// Experimental.
	EphemeralBlockDevice() AwsAmi_EphemeralBlockDevicePropertyList
	// Experimental.
	EphemeralBlockDeviceInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Hypervisor() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	ImageLocation() *string
	// Experimental.
	SetImageLocation(val *string)
	// Experimental.
	ImageLocationInput() *string
	// Experimental.
	ImageOwnerAlias() *string
	// Experimental.
	ImageType() *string
	// Experimental.
	ImdsSupport() *string
	// Experimental.
	SetImdsSupport(val *string)
	// Experimental.
	ImdsSupportInput() *string
	// Experimental.
	KernelId() *string
	// Experimental.
	SetKernelId(val *string)
	// Experimental.
	KernelIdInput() *string
	// Experimental.
	LastLaunchedTime() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	ManageEbsSnapshots() cdktn.IResolvable
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OwnerId() *string
	// Experimental.
	Platform() *string
	// Experimental.
	PlatformDetails() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	Public() cdktn.IResolvable
	// Experimental.
	RamdiskId() *string
	// Experimental.
	SetRamdiskId(val *string)
	// Experimental.
	RamdiskIdInput() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	RootDeviceName() *string
	// Experimental.
	SetRootDeviceName(val *string)
	// Experimental.
	RootDeviceNameInput() *string
	// Experimental.
	RootSnapshotId() *string
	// Experimental.
	SriovNetSupport() *string
	// Experimental.
	SetSriovNetSupport(val *string)
	// Experimental.
	SriovNetSupportInput() *string
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
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsAmi_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TpmSupport() *string
	// Experimental.
	SetTpmSupport(val *string)
	// Experimental.
	TpmSupportInput() *string
	// Experimental.
	UefiData() *string
	// Experimental.
	SetUefiData(val *string)
	// Experimental.
	UefiDataInput() *string
	// Experimental.
	UsageOperation() *string
	// Experimental.
	VirtualizationType() *string
	// Experimental.
	SetVirtualizationType(val *string)
	// Experimental.
	VirtualizationTypeInput() *string
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
	PutEbsBlockDevice(value interface{})
	// Experimental.
	PutEphemeralBlockDevice(value interface{})
	// Experimental.
	PutTimeouts(value *AwsAmi_TimeoutsProperty)
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
	ResetArchitecture()
	// Experimental.
	ResetBootMode()
	// Experimental.
	ResetDeprecationTime()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetEbsBlockDevice()
	// Experimental.
	ResetEnaSupport()
	// Experimental.
	ResetEphemeralBlockDevice()
	// Experimental.
	ResetId()
	// Experimental.
	ResetImageLocation()
	// Experimental.
	ResetImdsSupport()
	// Experimental.
	ResetKernelId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRamdiskId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRootDeviceName()
	// Experimental.
	ResetSriovNetSupport()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTpmSupport()
	// Experimental.
	ResetUefiData()
	// Experimental.
	ResetVirtualizationType()
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

// The jsii proxy struct for AwsAmi
type jsiiProxy_AwsAmi struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsAmi) Architecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) ArchitectureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) BootMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) BootModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) DeprecationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deprecationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) DeprecationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deprecationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) EbsBlockDevice() AwsAmi_EbsBlockDevicePropertyList {
	var returns AwsAmi_EbsBlockDevicePropertyList
	_jsii_.Get(
		j,
		"ebsBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) EbsBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) EnaSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enaSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) EnaSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enaSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) EphemeralBlockDevice() AwsAmi_EphemeralBlockDevicePropertyList {
	var returns AwsAmi_EphemeralBlockDevicePropertyList
	_jsii_.Get(
		j,
		"ephemeralBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) EphemeralBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) Hypervisor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hypervisor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) ImageLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) ImageLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) ImageOwnerAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageOwnerAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) ImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) ImdsSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imdsSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) ImdsSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imdsSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) KernelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) KernelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) LastLaunchedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastLaunchedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) ManageEbsSnapshots() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"manageEbsSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) PlatformDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) Public() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"public",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) RamdiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramdiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) RamdiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ramdiskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) RootDeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDeviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) RootDeviceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDeviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) RootSnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootSnapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) SriovNetSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sriovNetSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) SriovNetSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sriovNetSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) Timeouts() AwsAmi_TimeoutsPropertyOutputReference {
	var returns AwsAmi_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) TpmSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpmSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) TpmSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpmSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) UefiData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uefiData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) UefiDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uefiDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) UsageOperation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) VirtualizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmi) VirtualizationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualizationTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami aws_ami} Resource.
// Experimental.
func NewAwsAmi(scope constructs.Construct, id *string, config *AwsAmiConfig) AwsAmi {
	_init_.Initialize()

	if err := validateNewAwsAmiParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAmi{}

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsAmi",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ami aws_ami} Resource.
// Experimental.
func NewAwsAmi_Override(a AwsAmi, scope constructs.Construct, id *string, config *AwsAmiConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsAmi",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsAmi)SetArchitecture(val *string) {
	if err := j.validateSetArchitectureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetBootMode(val *string) {
	if err := j.validateSetBootModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootMode",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetDeprecationTime(val *string) {
	if err := j.validateSetDeprecationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deprecationTime",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetEnaSupport(val interface{}) {
	if err := j.validateSetEnaSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enaSupport",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetImageLocation(val *string) {
	if err := j.validateSetImageLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageLocation",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetImdsSupport(val *string) {
	if err := j.validateSetImdsSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imdsSupport",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetKernelId(val *string) {
	if err := j.validateSetKernelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kernelId",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetRamdiskId(val *string) {
	if err := j.validateSetRamdiskIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ramdiskId",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetRootDeviceName(val *string) {
	if err := j.validateSetRootDeviceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootDeviceName",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetSriovNetSupport(val *string) {
	if err := j.validateSetSriovNetSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sriovNetSupport",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetTpmSupport(val *string) {
	if err := j.validateSetTpmSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tpmSupport",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetUefiData(val *string) {
	if err := j.validateSetUefiDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uefiData",
		val,
	)
}

func (j *jsiiProxy_AwsAmi)SetVirtualizationType(val *string) {
	if err := j.validateSetVirtualizationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualizationType",
		val,
	)
}

// Generates CDKTN code for importing a AwsAmi resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsAmi_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsAmi_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsAmi",
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
func AwsAmi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAmi_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsAmi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAmi_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAmi_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsAmi",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAmi_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAmi_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsAmi",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsAmi_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-ec2.AwsAmi",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsAmi) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsAmi) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsAmi) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAmi) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAmi) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAmi) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAmi) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAmi) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAmi) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAmi) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAmi) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAmi) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAmi) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsAmi) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAmi) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsAmi) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAmi) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsAmi) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAmi) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsAmi) PutEbsBlockDevice(value interface{}) {
	if err := a.validatePutEbsBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEbsBlockDevice",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAmi) PutEphemeralBlockDevice(value interface{}) {
	if err := a.validatePutEphemeralBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEphemeralBlockDevice",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAmi) PutTimeouts(value *AwsAmi_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAmi) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsAmi) ResetArchitecture() {
	_jsii_.InvokeVoid(
		a,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetBootMode() {
	_jsii_.InvokeVoid(
		a,
		"resetBootMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetDeprecationTime() {
	_jsii_.InvokeVoid(
		a,
		"resetDeprecationTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetEbsBlockDevice() {
	_jsii_.InvokeVoid(
		a,
		"resetEbsBlockDevice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetEnaSupport() {
	_jsii_.InvokeVoid(
		a,
		"resetEnaSupport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetEphemeralBlockDevice() {
	_jsii_.InvokeVoid(
		a,
		"resetEphemeralBlockDevice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetImageLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetImageLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetImdsSupport() {
	_jsii_.InvokeVoid(
		a,
		"resetImdsSupport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetKernelId() {
	_jsii_.InvokeVoid(
		a,
		"resetKernelId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetRamdiskId() {
	_jsii_.InvokeVoid(
		a,
		"resetRamdiskId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetRootDeviceName() {
	_jsii_.InvokeVoid(
		a,
		"resetRootDeviceName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetSriovNetSupport() {
	_jsii_.InvokeVoid(
		a,
		"resetSriovNetSupport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetTpmSupport() {
	_jsii_.InvokeVoid(
		a,
		"resetTpmSupport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetUefiData() {
	_jsii_.InvokeVoid(
		a,
		"resetUefiData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) ResetVirtualizationType() {
	_jsii_.InvokeVoid(
		a,
		"resetVirtualizationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmi) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAmi) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAmi) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAmi) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAmi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAmi) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAmi) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

