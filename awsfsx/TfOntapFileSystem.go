package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_file_system aws_fsx_ontap_file_system}.
// Experimental.
type TfOntapFileSystem interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	AutomaticBackupRetentionDays() *float64
	// Experimental.
	SetAutomaticBackupRetentionDays(val *float64)
	// Experimental.
	AutomaticBackupRetentionDaysInput() *float64
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
	DailyAutomaticBackupStartTime() *string
	// Experimental.
	SetDailyAutomaticBackupStartTime(val *string)
	// Experimental.
	DailyAutomaticBackupStartTimeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DeploymentType() *string
	// Experimental.
	SetDeploymentType(val *string)
	// Experimental.
	DeploymentTypeInput() *string
	// Experimental.
	DiskIopsConfiguration() TfOntapFileSystem_DiskIopsConfigurationPropertyOutputReference
	// Experimental.
	DiskIopsConfigurationInput() *TfOntapFileSystem_DiskIopsConfigurationProperty
	// Experimental.
	DnsName() *string
	// Experimental.
	EndpointIpAddressRange() *string
	// Experimental.
	SetEndpointIpAddressRange(val *string)
	// Experimental.
	EndpointIpAddressRangeInput() *string
	// Experimental.
	Endpoints() TfOntapFileSystem_EndpointsPropertyList
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	FsxAdminPassword() *string
	// Experimental.
	SetFsxAdminPassword(val *string)
	// Experimental.
	FsxAdminPasswordInput() *string
	// Experimental.
	HaPairs() *float64
	// Experimental.
	SetHaPairs(val *float64)
	// Experimental.
	HaPairsInput() *float64
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	KmsKeyId() *string
	// Experimental.
	SetKmsKeyId(val *string)
	// Experimental.
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	NetworkInterfaceIds() *[]*string
	// Experimental.
	NetworkType() *string
	// Experimental.
	SetNetworkType(val *string)
	// Experimental.
	NetworkTypeInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OwnerId() *string
	// Experimental.
	PreferredSubnetId() *string
	// Experimental.
	SetPreferredSubnetId(val *string)
	// Experimental.
	PreferredSubnetIdInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	RouteTableIds() *[]*string
	// Experimental.
	SetRouteTableIds(val *[]*string)
	// Experimental.
	RouteTableIdsInput() *[]*string
	// Experimental.
	SecurityGroupIds() *[]*string
	// Experimental.
	SetSecurityGroupIds(val *[]*string)
	// Experimental.
	SecurityGroupIdsInput() *[]*string
	// Experimental.
	StorageCapacity() *float64
	// Experimental.
	SetStorageCapacity(val *float64)
	// Experimental.
	StorageCapacityInput() *float64
	// Experimental.
	StorageType() *string
	// Experimental.
	SetStorageType(val *string)
	// Experimental.
	StorageTypeInput() *string
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
	ThroughputCapacity() *float64
	// Experimental.
	SetThroughputCapacity(val *float64)
	// Experimental.
	ThroughputCapacityInput() *float64
	// Experimental.
	ThroughputCapacityPerHaPair() *float64
	// Experimental.
	SetThroughputCapacityPerHaPair(val *float64)
	// Experimental.
	ThroughputCapacityPerHaPairInput() *float64
	// Experimental.
	Timeouts() TfOntapFileSystem_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	VpcId() *string
	// Experimental.
	WeeklyMaintenanceStartTime() *string
	// Experimental.
	SetWeeklyMaintenanceStartTime(val *string)
	// Experimental.
	WeeklyMaintenanceStartTimeInput() *string
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
	PutDiskIopsConfiguration(value *TfOntapFileSystem_DiskIopsConfigurationProperty)
	// Experimental.
	PutTimeouts(value *TfOntapFileSystem_TimeoutsProperty)
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
	ResetAutomaticBackupRetentionDays()
	// Experimental.
	ResetDailyAutomaticBackupStartTime()
	// Experimental.
	ResetDiskIopsConfiguration()
	// Experimental.
	ResetEndpointIpAddressRange()
	// Experimental.
	ResetFsxAdminPassword()
	// Experimental.
	ResetHaPairs()
	// Experimental.
	ResetId()
	// Experimental.
	ResetKmsKeyId()
	// Experimental.
	ResetNetworkType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRouteTableIds()
	// Experimental.
	ResetSecurityGroupIds()
	// Experimental.
	ResetStorageType()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetThroughputCapacity()
	// Experimental.
	ResetThroughputCapacityPerHaPair()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetWeeklyMaintenanceStartTime()
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

// The jsii proxy struct for TfOntapFileSystem
type jsiiProxy_TfOntapFileSystem struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfOntapFileSystem) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) AutomaticBackupRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) AutomaticBackupRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) DailyAutomaticBackupStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) DailyAutomaticBackupStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) DeploymentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) DeploymentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) DiskIopsConfiguration() TfOntapFileSystem_DiskIopsConfigurationPropertyOutputReference {
	var returns TfOntapFileSystem_DiskIopsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"diskIopsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) DiskIopsConfigurationInput() *TfOntapFileSystem_DiskIopsConfigurationProperty {
	var returns *TfOntapFileSystem_DiskIopsConfigurationProperty
	_jsii_.Get(
		j,
		"diskIopsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) EndpointIpAddressRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIpAddressRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) EndpointIpAddressRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIpAddressRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) Endpoints() TfOntapFileSystem_EndpointsPropertyList {
	var returns TfOntapFileSystem_EndpointsPropertyList
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) FsxAdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsxAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) FsxAdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsxAdminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) HaPairs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haPairs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) HaPairsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"haPairsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) NetworkInterfaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) PreferredSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) PreferredSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) RouteTableIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routeTableIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) RouteTableIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routeTableIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) StorageCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) StorageCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) ThroughputCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) ThroughputCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) ThroughputCapacityPerHaPair() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacityPerHaPair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) ThroughputCapacityPerHaPairInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacityPerHaPairInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) Timeouts() TfOntapFileSystem_TimeoutsPropertyOutputReference {
	var returns TfOntapFileSystem_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) WeeklyMaintenanceStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapFileSystem) WeeklyMaintenanceStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTimeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_file_system aws_fsx_ontap_file_system} Resource.
// Experimental.
func NewTfOntapFileSystem(scope constructs.Construct, id *string, config *TfOntapFileSystemConfig) TfOntapFileSystem {
	_init_.Initialize()

	if err := validateNewTfOntapFileSystemParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfOntapFileSystem{}

	_jsii_.Create(
		"@cdktn/aws-fsx.TfOntapFileSystem",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_file_system aws_fsx_ontap_file_system} Resource.
// Experimental.
func NewTfOntapFileSystem_Override(t TfOntapFileSystem, scope constructs.Construct, id *string, config *TfOntapFileSystemConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.TfOntapFileSystem",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetAutomaticBackupRetentionDays(val *float64) {
	if err := j.validateSetAutomaticBackupRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticBackupRetentionDays",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetDailyAutomaticBackupStartTime(val *string) {
	if err := j.validateSetDailyAutomaticBackupStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailyAutomaticBackupStartTime",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetDeploymentType(val *string) {
	if err := j.validateSetDeploymentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentType",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetEndpointIpAddressRange(val *string) {
	if err := j.validateSetEndpointIpAddressRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointIpAddressRange",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetFsxAdminPassword(val *string) {
	if err := j.validateSetFsxAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fsxAdminPassword",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetHaPairs(val *float64) {
	if err := j.validateSetHaPairsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"haPairs",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetPreferredSubnetId(val *string) {
	if err := j.validateSetPreferredSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredSubnetId",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetRouteTableIds(val *[]*string) {
	if err := j.validateSetRouteTableIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeTableIds",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetStorageCapacity(val *float64) {
	if err := j.validateSetStorageCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCapacity",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetThroughputCapacity(val *float64) {
	if err := j.validateSetThroughputCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughputCapacity",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetThroughputCapacityPerHaPair(val *float64) {
	if err := j.validateSetThroughputCapacityPerHaPairParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughputCapacityPerHaPair",
		val,
	)
}

func (j *jsiiProxy_TfOntapFileSystem)SetWeeklyMaintenanceStartTime(val *string) {
	if err := j.validateSetWeeklyMaintenanceStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeklyMaintenanceStartTime",
		val,
	)
}

// Generates CDKTN code for importing a TfOntapFileSystem resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfOntapFileSystem_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfOntapFileSystem_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.TfOntapFileSystem",
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
func TfOntapFileSystem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfOntapFileSystem_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.TfOntapFileSystem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfOntapFileSystem_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfOntapFileSystem_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.TfOntapFileSystem",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfOntapFileSystem_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfOntapFileSystem_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.TfOntapFileSystem",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfOntapFileSystem_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-fsx.TfOntapFileSystem",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfOntapFileSystem) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfOntapFileSystem) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfOntapFileSystem) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfOntapFileSystem) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOntapFileSystem) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfOntapFileSystem) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfOntapFileSystem) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfOntapFileSystem) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfOntapFileSystem) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfOntapFileSystem) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfOntapFileSystem) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfOntapFileSystem) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapFileSystem) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfOntapFileSystem) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOntapFileSystem) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfOntapFileSystem) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfOntapFileSystem) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfOntapFileSystem) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfOntapFileSystem) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfOntapFileSystem) PutDiskIopsConfiguration(value *TfOntapFileSystem_DiskIopsConfigurationProperty) {
	if err := t.validatePutDiskIopsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDiskIopsConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOntapFileSystem) PutTimeouts(value *TfOntapFileSystem_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOntapFileSystem) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetAutomaticBackupRetentionDays() {
	_jsii_.InvokeVoid(
		t,
		"resetAutomaticBackupRetentionDays",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetDailyAutomaticBackupStartTime() {
	_jsii_.InvokeVoid(
		t,
		"resetDailyAutomaticBackupStartTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetDiskIopsConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetDiskIopsConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetEndpointIpAddressRange() {
	_jsii_.InvokeVoid(
		t,
		"resetEndpointIpAddressRange",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetFsxAdminPassword() {
	_jsii_.InvokeVoid(
		t,
		"resetFsxAdminPassword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetHaPairs() {
	_jsii_.InvokeVoid(
		t,
		"resetHaPairs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetNetworkType() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetRouteTableIds() {
	_jsii_.InvokeVoid(
		t,
		"resetRouteTableIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetStorageType() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetThroughputCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetThroughputCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetThroughputCapacityPerHaPair() {
	_jsii_.InvokeVoid(
		t,
		"resetThroughputCapacityPerHaPair",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) ResetWeeklyMaintenanceStartTime() {
	_jsii_.InvokeVoid(
		t,
		"resetWeeklyMaintenanceStartTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapFileSystem) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapFileSystem) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapFileSystem) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapFileSystem) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapFileSystem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapFileSystem) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapFileSystem) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

