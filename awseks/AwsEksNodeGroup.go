package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseks/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awseks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group aws_eks_node_group}.
// Experimental.
type AwsEksNodeGroup interface {
	cdktn.TerraformResource
	// Experimental.
	AmiType() *string
	// Experimental.
	SetAmiType(val *string)
	// Experimental.
	AmiTypeInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	CapacityType() *string
	// Experimental.
	SetCapacityType(val *string)
	// Experimental.
	CapacityTypeInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClusterName() *string
	// Experimental.
	SetClusterName(val *string)
	// Experimental.
	ClusterNameInput() *string
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
	DiskSize() *float64
	// Experimental.
	SetDiskSize(val *float64)
	// Experimental.
	DiskSizeInput() *float64
	// Experimental.
	ForceUpdateVersion() interface{}
	// Experimental.
	SetForceUpdateVersion(val interface{})
	// Experimental.
	ForceUpdateVersionInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	InstanceTypes() *[]*string
	// Experimental.
	SetInstanceTypes(val *[]*string)
	// Experimental.
	InstanceTypesInput() *[]*string
	// Experimental.
	Labels() *map[string]*string
	// Experimental.
	SetLabels(val *map[string]*string)
	// Experimental.
	LabelsInput() *map[string]*string
	// Experimental.
	LaunchTemplate() AwsEksNodeGroup_LaunchTemplatePropertyOutputReference
	// Experimental.
	LaunchTemplateInput() *AwsEksNodeGroup_LaunchTemplateProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	NodeGroupName() *string
	// Experimental.
	SetNodeGroupName(val *string)
	// Experimental.
	NodeGroupNameInput() *string
	// Experimental.
	NodeGroupNamePrefix() *string
	// Experimental.
	SetNodeGroupNamePrefix(val *string)
	// Experimental.
	NodeGroupNamePrefixInput() *string
	// Experimental.
	NodeRepairConfig() AwsEksNodeGroup_NodeRepairConfigPropertyOutputReference
	// Experimental.
	NodeRepairConfigInput() *AwsEksNodeGroup_NodeRepairConfigProperty
	// Experimental.
	NodeRoleArn() *string
	// Experimental.
	SetNodeRoleArn(val *string)
	// Experimental.
	NodeRoleArnInput() *string
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
	ReleaseVersion() *string
	// Experimental.
	SetReleaseVersion(val *string)
	// Experimental.
	ReleaseVersionInput() *string
	// Experimental.
	RemoteAccess() AwsEksNodeGroup_RemoteAccessPropertyOutputReference
	// Experimental.
	RemoteAccessInput() *AwsEksNodeGroup_RemoteAccessProperty
	// Experimental.
	Resources() AwsEksNodeGroup_ResourcesPropertyList
	// Experimental.
	ScalingConfig() AwsEksNodeGroup_ScalingConfigPropertyOutputReference
	// Experimental.
	ScalingConfigInput() *AwsEksNodeGroup_ScalingConfigProperty
	// Experimental.
	Status() *string
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
	Taint() AwsEksNodeGroup_TaintPropertyList
	// Experimental.
	TaintInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsEksNodeGroup_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	UpdateConfig() AwsEksNodeGroup_UpdateConfigPropertyOutputReference
	// Experimental.
	UpdateConfigInput() *AwsEksNodeGroup_UpdateConfigProperty
	// Experimental.
	Version() *string
	// Experimental.
	SetVersion(val *string)
	// Experimental.
	VersionInput() *string
	// Experimental.
	WarmPoolConfig() AwsEksNodeGroup_WarmPoolConfigPropertyOutputReference
	// Experimental.
	WarmPoolConfigInput() *AwsEksNodeGroup_WarmPoolConfigProperty
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
	PutLaunchTemplate(value *AwsEksNodeGroup_LaunchTemplateProperty)
	// Experimental.
	PutNodeRepairConfig(value *AwsEksNodeGroup_NodeRepairConfigProperty)
	// Experimental.
	PutRemoteAccess(value *AwsEksNodeGroup_RemoteAccessProperty)
	// Experimental.
	PutScalingConfig(value *AwsEksNodeGroup_ScalingConfigProperty)
	// Experimental.
	PutTaint(value interface{})
	// Experimental.
	PutTimeouts(value *AwsEksNodeGroup_TimeoutsProperty)
	// Experimental.
	PutUpdateConfig(value *AwsEksNodeGroup_UpdateConfigProperty)
	// Experimental.
	PutWarmPoolConfig(value *AwsEksNodeGroup_WarmPoolConfigProperty)
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
	ResetAmiType()
	// Experimental.
	ResetCapacityType()
	// Experimental.
	ResetDiskSize()
	// Experimental.
	ResetForceUpdateVersion()
	// Experimental.
	ResetId()
	// Experimental.
	ResetInstanceTypes()
	// Experimental.
	ResetLabels()
	// Experimental.
	ResetLaunchTemplate()
	// Experimental.
	ResetNodeGroupName()
	// Experimental.
	ResetNodeGroupNamePrefix()
	// Experimental.
	ResetNodeRepairConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetReleaseVersion()
	// Experimental.
	ResetRemoteAccess()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTaint()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetUpdateConfig()
	// Experimental.
	ResetVersion()
	// Experimental.
	ResetWarmPoolConfig()
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

// The jsii proxy struct for AwsEksNodeGroup
type jsiiProxy_AwsEksNodeGroup struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsEksNodeGroup) AmiType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) AmiTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) CapacityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) CapacityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) DiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) DiskSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) ForceUpdateVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) ForceUpdateVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) InstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) InstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) LaunchTemplate() AwsEksNodeGroup_LaunchTemplatePropertyOutputReference {
	var returns AwsEksNodeGroup_LaunchTemplatePropertyOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) LaunchTemplateInput() *AwsEksNodeGroup_LaunchTemplateProperty {
	var returns *AwsEksNodeGroup_LaunchTemplateProperty
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) NodeGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) NodeGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) NodeGroupNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) NodeGroupNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) NodeRepairConfig() AwsEksNodeGroup_NodeRepairConfigPropertyOutputReference {
	var returns AwsEksNodeGroup_NodeRepairConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"nodeRepairConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) NodeRepairConfigInput() *AwsEksNodeGroup_NodeRepairConfigProperty {
	var returns *AwsEksNodeGroup_NodeRepairConfigProperty
	_jsii_.Get(
		j,
		"nodeRepairConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) NodeRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) NodeRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) ReleaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) ReleaseVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) RemoteAccess() AwsEksNodeGroup_RemoteAccessPropertyOutputReference {
	var returns AwsEksNodeGroup_RemoteAccessPropertyOutputReference
	_jsii_.Get(
		j,
		"remoteAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) RemoteAccessInput() *AwsEksNodeGroup_RemoteAccessProperty {
	var returns *AwsEksNodeGroup_RemoteAccessProperty
	_jsii_.Get(
		j,
		"remoteAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) Resources() AwsEksNodeGroup_ResourcesPropertyList {
	var returns AwsEksNodeGroup_ResourcesPropertyList
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) ScalingConfig() AwsEksNodeGroup_ScalingConfigPropertyOutputReference {
	var returns AwsEksNodeGroup_ScalingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"scalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) ScalingConfigInput() *AwsEksNodeGroup_ScalingConfigProperty {
	var returns *AwsEksNodeGroup_ScalingConfigProperty
	_jsii_.Get(
		j,
		"scalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) Taint() AwsEksNodeGroup_TaintPropertyList {
	var returns AwsEksNodeGroup_TaintPropertyList
	_jsii_.Get(
		j,
		"taint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) TaintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) Timeouts() AwsEksNodeGroup_TimeoutsPropertyOutputReference {
	var returns AwsEksNodeGroup_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) UpdateConfig() AwsEksNodeGroup_UpdateConfigPropertyOutputReference {
	var returns AwsEksNodeGroup_UpdateConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"updateConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) UpdateConfigInput() *AwsEksNodeGroup_UpdateConfigProperty {
	var returns *AwsEksNodeGroup_UpdateConfigProperty
	_jsii_.Get(
		j,
		"updateConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) WarmPoolConfig() AwsEksNodeGroup_WarmPoolConfigPropertyOutputReference {
	var returns AwsEksNodeGroup_WarmPoolConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"warmPoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksNodeGroup) WarmPoolConfigInput() *AwsEksNodeGroup_WarmPoolConfigProperty {
	var returns *AwsEksNodeGroup_WarmPoolConfigProperty
	_jsii_.Get(
		j,
		"warmPoolConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group aws_eks_node_group} Resource.
// Experimental.
func NewAwsEksNodeGroup(scope constructs.Construct, id *string, config *AwsEksNodeGroupConfig) AwsEksNodeGroup {
	_init_.Initialize()

	if err := validateNewAwsEksNodeGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEksNodeGroup{}

	_jsii_.Create(
		"@cdktn/aws-eks.AwsEksNodeGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group aws_eks_node_group} Resource.
// Experimental.
func NewAwsEksNodeGroup_Override(a AwsEksNodeGroup, scope constructs.Construct, id *string, config *AwsEksNodeGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.AwsEksNodeGroup",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetAmiType(val *string) {
	if err := j.validateSetAmiTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amiType",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetCapacityType(val *string) {
	if err := j.validateSetCapacityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityType",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetDiskSize(val *float64) {
	if err := j.validateSetDiskSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSize",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetForceUpdateVersion(val interface{}) {
	if err := j.validateSetForceUpdateVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceUpdateVersion",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetInstanceTypes(val *[]*string) {
	if err := j.validateSetInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypes",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetNodeGroupName(val *string) {
	if err := j.validateSetNodeGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeGroupName",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetNodeGroupNamePrefix(val *string) {
	if err := j.validateSetNodeGroupNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeGroupNamePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetNodeRoleArn(val *string) {
	if err := j.validateSetNodeRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetReleaseVersion(val *string) {
	if err := j.validateSetReleaseVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseVersion",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsEksNodeGroup)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTN code for importing a AwsEksNodeGroup resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsEksNodeGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsEksNodeGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.AwsEksNodeGroup",
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
func AwsEksNodeGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEksNodeGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.AwsEksNodeGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsEksNodeGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEksNodeGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.AwsEksNodeGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsEksNodeGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEksNodeGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-eks.AwsEksNodeGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsEksNodeGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-eks.AwsEksNodeGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsEksNodeGroup) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEksNodeGroup) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEksNodeGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEksNodeGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEksNodeGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEksNodeGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEksNodeGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEksNodeGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEksNodeGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEksNodeGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksNodeGroup) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEksNodeGroup) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsEksNodeGroup) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) PutLaunchTemplate(value *AwsEksNodeGroup_LaunchTemplateProperty) {
	if err := a.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) PutNodeRepairConfig(value *AwsEksNodeGroup_NodeRepairConfigProperty) {
	if err := a.validatePutNodeRepairConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNodeRepairConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) PutRemoteAccess(value *AwsEksNodeGroup_RemoteAccessProperty) {
	if err := a.validatePutRemoteAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRemoteAccess",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) PutScalingConfig(value *AwsEksNodeGroup_ScalingConfigProperty) {
	if err := a.validatePutScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScalingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) PutTaint(value interface{}) {
	if err := a.validatePutTaintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTaint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) PutTimeouts(value *AwsEksNodeGroup_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) PutUpdateConfig(value *AwsEksNodeGroup_UpdateConfigProperty) {
	if err := a.validatePutUpdateConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUpdateConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) PutWarmPoolConfig(value *AwsEksNodeGroup_WarmPoolConfigProperty) {
	if err := a.validatePutWarmPoolConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWarmPoolConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetAmiType() {
	_jsii_.InvokeVoid(
		a,
		"resetAmiType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetCapacityType() {
	_jsii_.InvokeVoid(
		a,
		"resetCapacityType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetDiskSize() {
	_jsii_.InvokeVoid(
		a,
		"resetDiskSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetForceUpdateVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetForceUpdateVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetInstanceTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetLabels() {
	_jsii_.InvokeVoid(
		a,
		"resetLabels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetNodeGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetNodeGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetNodeGroupNamePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetNodeGroupNamePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetNodeRepairConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetNodeRepairConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetReleaseVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetReleaseVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetRemoteAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetRemoteAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetTaint() {
	_jsii_.InvokeVoid(
		a,
		"resetTaint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetUpdateConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetUpdateConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) ResetWarmPoolConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetWarmPoolConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksNodeGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksNodeGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksNodeGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksNodeGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksNodeGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksNodeGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksNodeGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

