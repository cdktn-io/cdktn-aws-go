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
type TfReplicationConfigurationTemplate interface {
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
	PitPolicy() TfReplicationConfigurationTemplate_PitPolicyPropertyList
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
	Timeouts() TfReplicationConfigurationTemplate_TimeoutsPropertyOutputReference
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
	PutTimeouts(value *TfReplicationConfigurationTemplate_TimeoutsProperty)
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

// The jsii proxy struct for TfReplicationConfigurationTemplate
type jsiiProxy_TfReplicationConfigurationTemplate struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) AssociateDefaultSecurityGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateDefaultSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) AssociateDefaultSecurityGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associateDefaultSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) AutoReplicateNewDisks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoReplicateNewDisks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) AutoReplicateNewDisksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoReplicateNewDisksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) BandwidthThrottling() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthThrottling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) BandwidthThrottlingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthThrottlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) CreatePublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) CreatePublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) DataPlaneRouting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPlaneRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) DataPlaneRoutingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPlaneRoutingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) DefaultLargeStagingDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLargeStagingDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) DefaultLargeStagingDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLargeStagingDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) EbsEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) EbsEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) EbsEncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) EbsEncryptionKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsEncryptionKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) PitPolicy() TfReplicationConfigurationTemplate_PitPolicyPropertyList {
	var returns TfReplicationConfigurationTemplate_PitPolicyPropertyList
	_jsii_.Get(
		j,
		"pitPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) PitPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pitPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) ReplicationServerInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationServerInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) ReplicationServerInstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationServerInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) ReplicationServersSecurityGroupsIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicationServersSecurityGroupsIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) ReplicationServersSecurityGroupsIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicationServersSecurityGroupsIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) StagingAreaSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingAreaSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) StagingAreaSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingAreaSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) StagingAreaTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stagingAreaTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) StagingAreaTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stagingAreaTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) Timeouts() TfReplicationConfigurationTemplate_TimeoutsPropertyOutputReference {
	var returns TfReplicationConfigurationTemplate_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) UseDedicatedReplicationServer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDedicatedReplicationServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate) UseDedicatedReplicationServerInput() interface{} {
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
func NewTfReplicationConfigurationTemplate(scope constructs.Construct, id *string, config *TfReplicationConfigurationTemplateConfig) TfReplicationConfigurationTemplate {
	_init_.Initialize()

	if err := validateNewTfReplicationConfigurationTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfReplicationConfigurationTemplate{}

	_jsii_.Create(
		"@cdktn/aws-drs.TfReplicationConfigurationTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template aws_drs_replication_configuration_template} Resource.
// Experimental.
func NewTfReplicationConfigurationTemplate_Override(t TfReplicationConfigurationTemplate, scope constructs.Construct, id *string, config *TfReplicationConfigurationTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-drs.TfReplicationConfigurationTemplate",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetAssociateDefaultSecurityGroup(val interface{}) {
	if err := j.validateSetAssociateDefaultSecurityGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associateDefaultSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetAutoReplicateNewDisks(val interface{}) {
	if err := j.validateSetAutoReplicateNewDisksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoReplicateNewDisks",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetBandwidthThrottling(val *float64) {
	if err := j.validateSetBandwidthThrottlingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidthThrottling",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetCreatePublicIp(val interface{}) {
	if err := j.validateSetCreatePublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createPublicIp",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetDataPlaneRouting(val *string) {
	if err := j.validateSetDataPlaneRoutingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataPlaneRouting",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetDefaultLargeStagingDiskType(val *string) {
	if err := j.validateSetDefaultLargeStagingDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLargeStagingDiskType",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetEbsEncryption(val *string) {
	if err := j.validateSetEbsEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsEncryption",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetEbsEncryptionKeyArn(val *string) {
	if err := j.validateSetEbsEncryptionKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsEncryptionKeyArn",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetReplicationServerInstanceType(val *string) {
	if err := j.validateSetReplicationServerInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationServerInstanceType",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetReplicationServersSecurityGroupsIds(val *[]*string) {
	if err := j.validateSetReplicationServersSecurityGroupsIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationServersSecurityGroupsIds",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetStagingAreaSubnetId(val *string) {
	if err := j.validateSetStagingAreaSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingAreaSubnetId",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetStagingAreaTags(val *map[string]*string) {
	if err := j.validateSetStagingAreaTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingAreaTags",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfigurationTemplate)SetUseDedicatedReplicationServer(val interface{}) {
	if err := j.validateSetUseDedicatedReplicationServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDedicatedReplicationServer",
		val,
	)
}

// Generates CDKTN code for importing a TfReplicationConfigurationTemplate resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfReplicationConfigurationTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfReplicationConfigurationTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-drs.TfReplicationConfigurationTemplate",
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
func TfReplicationConfigurationTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfReplicationConfigurationTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-drs.TfReplicationConfigurationTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfReplicationConfigurationTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfReplicationConfigurationTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-drs.TfReplicationConfigurationTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfReplicationConfigurationTemplate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfReplicationConfigurationTemplate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-drs.TfReplicationConfigurationTemplate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfReplicationConfigurationTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-drs.TfReplicationConfigurationTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfReplicationConfigurationTemplate) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReplicationConfigurationTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfReplicationConfigurationTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfReplicationConfigurationTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfReplicationConfigurationTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfReplicationConfigurationTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfReplicationConfigurationTemplate) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfReplicationConfigurationTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfReplicationConfigurationTemplate) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReplicationConfigurationTemplate) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfReplicationConfigurationTemplate) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) PutPitPolicy(value interface{}) {
	if err := t.validatePutPitPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPitPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) PutTimeouts(value *TfReplicationConfigurationTemplate_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) ResetAutoReplicateNewDisks() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoReplicateNewDisks",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) ResetEbsEncryptionKeyArn() {
	_jsii_.InvokeVoid(
		t,
		"resetEbsEncryptionKeyArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) ResetPitPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetPitPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationConfigurationTemplate) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

