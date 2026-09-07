package drs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/drs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/drs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template aws_drs_replication_configuration_template}.
// Experimental.
type AwsReplicationConfigurationTemplate interface {
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
	PitPolicy() AwsReplicationConfigurationTemplate_PitPolicyPropertyList
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
	Timeouts() AwsReplicationConfigurationTemplate_TimeoutsPropertyOutputReference
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
	PutTimeouts(value *AwsReplicationConfigurationTemplate_TimeoutsProperty)
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

// The jsii proxy struct for AwsReplicationConfigurationTemplate
type jsiiProxy_AwsReplicationConfigurationTemplate struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) AssociateDefaultSecurityGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateDefaultSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) AssociateDefaultSecurityGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateDefaultSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) AutoReplicateNewDisks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoReplicateNewDisks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) AutoReplicateNewDisksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoReplicateNewDisksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) BandwidthThrottling() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthThrottling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) BandwidthThrottlingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthThrottlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) CreatePublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) CreatePublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) DataPlaneRouting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPlaneRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) DataPlaneRoutingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPlaneRoutingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) DefaultLargeStagingDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLargeStagingDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) DefaultLargeStagingDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLargeStagingDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) EbsEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) EbsEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) EbsEncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) EbsEncryptionKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryptionKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) PitPolicy() AwsReplicationConfigurationTemplate_PitPolicyPropertyList {
	var returns AwsReplicationConfigurationTemplate_PitPolicyPropertyList
	_jsii_.Get(
		j,
		"pitPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) PitPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pitPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) ReplicationServerInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationServerInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) ReplicationServerInstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationServerInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) ReplicationServersSecurityGroupsIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicationServersSecurityGroupsIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) ReplicationServersSecurityGroupsIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicationServersSecurityGroupsIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) StagingAreaSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingAreaSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) StagingAreaSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingAreaSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) StagingAreaTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stagingAreaTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) StagingAreaTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stagingAreaTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) Timeouts() AwsReplicationConfigurationTemplate_TimeoutsPropertyOutputReference {
	var returns AwsReplicationConfigurationTemplate_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) UseDedicatedReplicationServer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDedicatedReplicationServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate) UseDedicatedReplicationServerInput() interface{} {
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
func NewAwsReplicationConfigurationTemplate(scope constructs.Construct, id *string, config *AwsReplicationConfigurationTemplateConfig) AwsReplicationConfigurationTemplate {
	_init_.Initialize()

	if err := validateNewAwsReplicationConfigurationTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsReplicationConfigurationTemplate{}

	_jsii_.Create(
		"@cdktn/aws-drs.AwsReplicationConfigurationTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template aws_drs_replication_configuration_template} Resource.
// Experimental.
func NewAwsReplicationConfigurationTemplate_Override(a AwsReplicationConfigurationTemplate, scope constructs.Construct, id *string, config *AwsReplicationConfigurationTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-drs.AwsReplicationConfigurationTemplate",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetAssociateDefaultSecurityGroup(val interface{}) {
	if err := j.validateSetAssociateDefaultSecurityGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associateDefaultSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetAutoReplicateNewDisks(val interface{}) {
	if err := j.validateSetAutoReplicateNewDisksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoReplicateNewDisks",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetBandwidthThrottling(val *float64) {
	if err := j.validateSetBandwidthThrottlingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthThrottling",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetCreatePublicIp(val interface{}) {
	if err := j.validateSetCreatePublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createPublicIp",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetDataPlaneRouting(val *string) {
	if err := j.validateSetDataPlaneRoutingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataPlaneRouting",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetDefaultLargeStagingDiskType(val *string) {
	if err := j.validateSetDefaultLargeStagingDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLargeStagingDiskType",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetEbsEncryption(val *string) {
	if err := j.validateSetEbsEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsEncryption",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetEbsEncryptionKeyArn(val *string) {
	if err := j.validateSetEbsEncryptionKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsEncryptionKeyArn",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetReplicationServerInstanceType(val *string) {
	if err := j.validateSetReplicationServerInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationServerInstanceType",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetReplicationServersSecurityGroupsIds(val *[]*string) {
	if err := j.validateSetReplicationServersSecurityGroupsIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationServersSecurityGroupsIds",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetStagingAreaSubnetId(val *string) {
	if err := j.validateSetStagingAreaSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingAreaSubnetId",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetStagingAreaTags(val *map[string]*string) {
	if err := j.validateSetStagingAreaTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingAreaTags",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationConfigurationTemplate)SetUseDedicatedReplicationServer(val interface{}) {
	if err := j.validateSetUseDedicatedReplicationServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDedicatedReplicationServer",
		val,
	)
}

// Generates CDKTN code for importing a AwsReplicationConfigurationTemplate resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsReplicationConfigurationTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsReplicationConfigurationTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-drs.AwsReplicationConfigurationTemplate",
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
func AwsReplicationConfigurationTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsReplicationConfigurationTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-drs.AwsReplicationConfigurationTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsReplicationConfigurationTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsReplicationConfigurationTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-drs.AwsReplicationConfigurationTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsReplicationConfigurationTemplate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsReplicationConfigurationTemplate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-drs.AwsReplicationConfigurationTemplate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsReplicationConfigurationTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-drs.AwsReplicationConfigurationTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) PutPitPolicy(value interface{}) {
	if err := a.validatePutPitPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPitPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) PutTimeouts(value *AwsReplicationConfigurationTemplate_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) ResetAutoReplicateNewDisks() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoReplicateNewDisks",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) ResetEbsEncryptionKeyArn() {
	_jsii_.InvokeVoid(
		a,
		"resetEbsEncryptionKeyArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) ResetPitPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPitPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReplicationConfigurationTemplate) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

