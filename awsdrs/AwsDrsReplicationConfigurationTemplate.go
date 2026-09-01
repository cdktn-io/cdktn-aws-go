package awsdrs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdrs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsdrs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template aws_drs_replication_configuration_template}.
// Experimental.
type AwsDrsReplicationConfigurationTemplate interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	AssociateDefaultSecurityGroup() interface{}
	// Experimental.
	SetAssociateDefaultSecurityGroup(val interface{})
	// Experimental.
	AssociateDefaultSecurityGroupInput() interface{}
	// Experimental.
	AutoReplicateNewDisks() interface{}
	// Experimental.
	SetAutoReplicateNewDisks(val interface{})
	// Experimental.
	AutoReplicateNewDisksInput() interface{}
	// Experimental.
	BandwidthThrottling() *float64
	// Experimental.
	SetBandwidthThrottling(val *float64)
	// Experimental.
	BandwidthThrottlingInput() *float64
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
	CreatePublicIp() interface{}
	// Experimental.
	SetCreatePublicIp(val interface{})
	// Experimental.
	CreatePublicIpInput() interface{}
	// Experimental.
	DataPlaneRouting() *string
	// Experimental.
	SetDataPlaneRouting(val *string)
	// Experimental.
	DataPlaneRoutingInput() *string
	// Experimental.
	DefaultLargeStagingDiskType() *string
	// Experimental.
	SetDefaultLargeStagingDiskType(val *string)
	// Experimental.
	DefaultLargeStagingDiskTypeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EbsEncryption() *string
	// Experimental.
	SetEbsEncryption(val *string)
	// Experimental.
	EbsEncryptionInput() *string
	// Experimental.
	EbsEncryptionKeyArn() *string
	// Experimental.
	SetEbsEncryptionKeyArn(val *string)
	// Experimental.
	EbsEncryptionKeyArnInput() *string
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
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PitPolicy() AwsDrsReplicationConfigurationTemplate_PitPolicyPropertyList
	// Experimental.
	PitPolicyInput() interface{}
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
	ReplicationServerInstanceType() *string
	// Experimental.
	SetReplicationServerInstanceType(val *string)
	// Experimental.
	ReplicationServerInstanceTypeInput() *string
	// Experimental.
	ReplicationServersSecurityGroupsIds() *[]*string
	// Experimental.
	SetReplicationServersSecurityGroupsIds(val *[]*string)
	// Experimental.
	ReplicationServersSecurityGroupsIdsInput() *[]*string
	// Experimental.
	StagingAreaSubnetId() *string
	// Experimental.
	SetStagingAreaSubnetId(val *string)
	// Experimental.
	StagingAreaSubnetIdInput() *string
	// Experimental.
	StagingAreaTags() *map[string]*string
	// Experimental.
	SetStagingAreaTags(val *map[string]*string)
	// Experimental.
	StagingAreaTagsInput() *map[string]*string
	// Experimental.
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsAll() cdktn.StringMap
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsDrsReplicationConfigurationTemplate_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	UseDedicatedReplicationServer() interface{}
	// Experimental.
	SetUseDedicatedReplicationServer(val interface{})
	// Experimental.
	UseDedicatedReplicationServerInput() interface{}
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
	PutPitPolicy(value interface{})
	// Experimental.
	PutTimeouts(value *AwsDrsReplicationConfigurationTemplate_TimeoutsProperty)
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
	ResetAutoReplicateNewDisks()
	// Experimental.
	ResetEbsEncryptionKeyArn()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPitPolicy()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTimeouts()
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

// The jsii proxy struct for AwsDrsReplicationConfigurationTemplate
type jsiiProxy_AwsDrsReplicationConfigurationTemplate struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) AssociateDefaultSecurityGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateDefaultSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) AssociateDefaultSecurityGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateDefaultSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) AutoReplicateNewDisks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoReplicateNewDisks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) AutoReplicateNewDisksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoReplicateNewDisksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) BandwidthThrottling() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthThrottling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) BandwidthThrottlingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthThrottlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) CreatePublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) CreatePublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) DataPlaneRouting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPlaneRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) DataPlaneRoutingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPlaneRoutingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) DefaultLargeStagingDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLargeStagingDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) DefaultLargeStagingDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLargeStagingDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) EbsEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) EbsEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) EbsEncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) EbsEncryptionKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryptionKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) PitPolicy() AwsDrsReplicationConfigurationTemplate_PitPolicyPropertyList {
	var returns AwsDrsReplicationConfigurationTemplate_PitPolicyPropertyList
	_jsii_.Get(
		j,
		"pitPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) PitPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pitPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ReplicationServerInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationServerInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ReplicationServerInstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationServerInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ReplicationServersSecurityGroupsIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicationServersSecurityGroupsIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ReplicationServersSecurityGroupsIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicationServersSecurityGroupsIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) StagingAreaSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingAreaSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) StagingAreaSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingAreaSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) StagingAreaTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stagingAreaTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) StagingAreaTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stagingAreaTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) Timeouts() AwsDrsReplicationConfigurationTemplate_TimeoutsPropertyOutputReference {
	var returns AwsDrsReplicationConfigurationTemplate_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) UseDedicatedReplicationServer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDedicatedReplicationServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate) UseDedicatedReplicationServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDedicatedReplicationServerInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template aws_drs_replication_configuration_template} Resource.
// Experimental.
func NewAwsDrsReplicationConfigurationTemplate(scope constructs.Construct, id *string, config *AwsDrsReplicationConfigurationTemplateConfig) AwsDrsReplicationConfigurationTemplate {
	_init_.Initialize()

	if err := validateNewAwsDrsReplicationConfigurationTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDrsReplicationConfigurationTemplate{}

	_jsii_.Create(
		"@cdktn/aws-drs.AwsDrsReplicationConfigurationTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template aws_drs_replication_configuration_template} Resource.
// Experimental.
func NewAwsDrsReplicationConfigurationTemplate_Override(a AwsDrsReplicationConfigurationTemplate, scope constructs.Construct, id *string, config *AwsDrsReplicationConfigurationTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-drs.AwsDrsReplicationConfigurationTemplate",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetAssociateDefaultSecurityGroup(val interface{}) {
	if err := j.validateSetAssociateDefaultSecurityGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associateDefaultSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetAutoReplicateNewDisks(val interface{}) {
	if err := j.validateSetAutoReplicateNewDisksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoReplicateNewDisks",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetBandwidthThrottling(val *float64) {
	if err := j.validateSetBandwidthThrottlingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthThrottling",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetCreatePublicIp(val interface{}) {
	if err := j.validateSetCreatePublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createPublicIp",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetDataPlaneRouting(val *string) {
	if err := j.validateSetDataPlaneRoutingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataPlaneRouting",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetDefaultLargeStagingDiskType(val *string) {
	if err := j.validateSetDefaultLargeStagingDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLargeStagingDiskType",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetEbsEncryption(val *string) {
	if err := j.validateSetEbsEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsEncryption",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetEbsEncryptionKeyArn(val *string) {
	if err := j.validateSetEbsEncryptionKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsEncryptionKeyArn",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetReplicationServerInstanceType(val *string) {
	if err := j.validateSetReplicationServerInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationServerInstanceType",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetReplicationServersSecurityGroupsIds(val *[]*string) {
	if err := j.validateSetReplicationServersSecurityGroupsIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationServersSecurityGroupsIds",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetStagingAreaSubnetId(val *string) {
	if err := j.validateSetStagingAreaSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingAreaSubnetId",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetStagingAreaTags(val *map[string]*string) {
	if err := j.validateSetStagingAreaTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingAreaTags",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsDrsReplicationConfigurationTemplate)SetUseDedicatedReplicationServer(val interface{}) {
	if err := j.validateSetUseDedicatedReplicationServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDedicatedReplicationServer",
		val,
	)
}

// Generates CDKTN code for importing a AwsDrsReplicationConfigurationTemplate resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsDrsReplicationConfigurationTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsDrsReplicationConfigurationTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-drs.AwsDrsReplicationConfigurationTemplate",
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
func AwsDrsReplicationConfigurationTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDrsReplicationConfigurationTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-drs.AwsDrsReplicationConfigurationTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsDrsReplicationConfigurationTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDrsReplicationConfigurationTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-drs.AwsDrsReplicationConfigurationTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsDrsReplicationConfigurationTemplate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDrsReplicationConfigurationTemplate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-drs.AwsDrsReplicationConfigurationTemplate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsDrsReplicationConfigurationTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-drs.AwsDrsReplicationConfigurationTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) PutPitPolicy(value interface{}) {
	if err := a.validatePutPitPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPitPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) PutTimeouts(value *AwsDrsReplicationConfigurationTemplate_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ResetAutoReplicateNewDisks() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoReplicateNewDisks",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ResetEbsEncryptionKeyArn() {
	_jsii_.InvokeVoid(
		a,
		"resetEbsEncryptionKeyArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ResetPitPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPitPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDrsReplicationConfigurationTemplate) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

