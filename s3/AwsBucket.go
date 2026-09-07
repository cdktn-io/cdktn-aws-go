package s3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/s3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/s3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket aws_s3_bucket}.
// Experimental.
type AwsBucket interface {
	cdktn.TerraformResource
	// Experimental.
	AccelerationStatus() *string
	// Experimental.
	SetAccelerationStatus(val *string)
	// Experimental.
	AccelerationStatusInput() *string
	// Experimental.
	Acl() *string
	// Experimental.
	SetAcl(val *string)
	// Experimental.
	AclInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	Bucket() *string
	// Experimental.
	SetBucket(val *string)
	// Experimental.
	BucketDomainName() *string
	// Experimental.
	BucketInput() *string
	// Experimental.
	BucketNamespace() *string
	// Experimental.
	SetBucketNamespace(val *string)
	// Experimental.
	BucketNamespaceInput() *string
	// Experimental.
	BucketPrefix() *string
	// Experimental.
	SetBucketPrefix(val *string)
	// Experimental.
	BucketPrefixInput() *string
	// Experimental.
	BucketRegion() *string
	// Experimental.
	BucketRegionalDomainName() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	CorsRule() AwsBucket_CorsRulePropertyList
	// Experimental.
	CorsRuleInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	Grant() AwsBucket_GrantPropertyList
	// Experimental.
	GrantInput() interface{}
	// Experimental.
	HostedZoneId() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LifecycleRule() AwsBucket_LifecycleRulePropertyList
	// Experimental.
	LifecycleRuleInput() interface{}
	// Experimental.
	Logging() AwsBucket_LoggingPropertyOutputReference
	// Experimental.
	LoggingInput() *AwsBucket_LoggingProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	ObjectLockConfiguration() AwsBucket_ObjectLockConfigurationPropertyOutputReference
	// Experimental.
	ObjectLockConfigurationInput() *AwsBucket_ObjectLockConfigurationProperty
	// Experimental.
	ObjectLockEnabled() interface{}
	// Experimental.
	SetObjectLockEnabled(val interface{})
	// Experimental.
	ObjectLockEnabledInput() interface{}
	// Experimental.
	Policy() *string
	// Experimental.
	SetPolicy(val *string)
	// Experimental.
	PolicyInput() *string
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
	ReplicationConfiguration() AwsBucket_ReplicationConfigurationPropertyOutputReference
	// Experimental.
	ReplicationConfigurationInput() *AwsBucket_ReplicationConfigurationProperty
	// Experimental.
	RequestPayer() *string
	// Experimental.
	SetRequestPayer(val *string)
	// Experimental.
	RequestPayerInput() *string
	// Experimental.
	ServerSideEncryptionConfiguration() AwsBucket_ServerSideEncryptionConfigurationPropertyOutputReference
	// Experimental.
	ServerSideEncryptionConfigurationInput() *AwsBucket_ServerSideEncryptionConfigurationProperty
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
	Timeouts() AwsBucket_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Versioning() AwsBucket_VersioningPropertyOutputReference
	// Experimental.
	VersioningInput() *AwsBucket_VersioningProperty
	// Experimental.
	Website() AwsBucket_WebsitePropertyOutputReference
	// Experimental.
	WebsiteDomain() *string
	// Experimental.
	WebsiteEndpoint() *string
	// Experimental.
	WebsiteInput() *AwsBucket_WebsiteProperty
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
	PutCorsRule(value interface{})
	// Experimental.
	PutGrant(value interface{})
	// Experimental.
	PutLifecycleRule(value interface{})
	// Experimental.
	PutLogging(value *AwsBucket_LoggingProperty)
	// Experimental.
	PutObjectLockConfiguration(value *AwsBucket_ObjectLockConfigurationProperty)
	// Experimental.
	PutReplicationConfiguration(value *AwsBucket_ReplicationConfigurationProperty)
	// Experimental.
	PutServerSideEncryptionConfiguration(value *AwsBucket_ServerSideEncryptionConfigurationProperty)
	// Experimental.
	PutTimeouts(value *AwsBucket_TimeoutsProperty)
	// Experimental.
	PutVersioning(value *AwsBucket_VersioningProperty)
	// Experimental.
	PutWebsite(value *AwsBucket_WebsiteProperty)
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
	ResetAccelerationStatus()
	// Experimental.
	ResetAcl()
	// Experimental.
	ResetBucket()
	// Experimental.
	ResetBucketNamespace()
	// Experimental.
	ResetBucketPrefix()
	// Experimental.
	ResetCorsRule()
	// Experimental.
	ResetForceDestroy()
	// Experimental.
	ResetGrant()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLifecycleRule()
	// Experimental.
	ResetLogging()
	// Experimental.
	ResetObjectLockConfiguration()
	// Experimental.
	ResetObjectLockEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPolicy()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetReplicationConfiguration()
	// Experimental.
	ResetRequestPayer()
	// Experimental.
	ResetServerSideEncryptionConfiguration()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetVersioning()
	// Experimental.
	ResetWebsite()
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

// The jsii proxy struct for AwsBucket
type jsiiProxy_AwsBucket struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsBucket) AccelerationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accelerationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) AccelerationStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accelerationStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Acl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) AclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) BucketDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) BucketNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) BucketNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) BucketPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) BucketPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) BucketRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) BucketRegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) CorsRule() AwsBucket_CorsRulePropertyList {
	var returns AwsBucket_CorsRulePropertyList
	_jsii_.Get(
		j,
		"corsRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) CorsRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"corsRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Grant() AwsBucket_GrantPropertyList {
	var returns AwsBucket_GrantPropertyList
	_jsii_.Get(
		j,
		"grant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) GrantInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) LifecycleRule() AwsBucket_LifecycleRulePropertyList {
	var returns AwsBucket_LifecycleRulePropertyList
	_jsii_.Get(
		j,
		"lifecycleRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) LifecycleRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecycleRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Logging() AwsBucket_LoggingPropertyOutputReference {
	var returns AwsBucket_LoggingPropertyOutputReference
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) LoggingInput() *AwsBucket_LoggingProperty {
	var returns *AwsBucket_LoggingProperty
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) ObjectLockConfiguration() AwsBucket_ObjectLockConfigurationPropertyOutputReference {
	var returns AwsBucket_ObjectLockConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"objectLockConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) ObjectLockConfigurationInput() *AwsBucket_ObjectLockConfigurationProperty {
	var returns *AwsBucket_ObjectLockConfigurationProperty
	_jsii_.Get(
		j,
		"objectLockConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) ObjectLockEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"objectLockEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) ObjectLockEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"objectLockEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) ReplicationConfiguration() AwsBucket_ReplicationConfigurationPropertyOutputReference {
	var returns AwsBucket_ReplicationConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"replicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) ReplicationConfigurationInput() *AwsBucket_ReplicationConfigurationProperty {
	var returns *AwsBucket_ReplicationConfigurationProperty
	_jsii_.Get(
		j,
		"replicationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) RequestPayer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestPayer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) RequestPayerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestPayerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) ServerSideEncryptionConfiguration() AwsBucket_ServerSideEncryptionConfigurationPropertyOutputReference {
	var returns AwsBucket_ServerSideEncryptionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"serverSideEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) ServerSideEncryptionConfigurationInput() *AwsBucket_ServerSideEncryptionConfigurationProperty {
	var returns *AwsBucket_ServerSideEncryptionConfigurationProperty
	_jsii_.Get(
		j,
		"serverSideEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Timeouts() AwsBucket_TimeoutsPropertyOutputReference {
	var returns AwsBucket_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Versioning() AwsBucket_VersioningPropertyOutputReference {
	var returns AwsBucket_VersioningPropertyOutputReference
	_jsii_.Get(
		j,
		"versioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) VersioningInput() *AwsBucket_VersioningProperty {
	var returns *AwsBucket_VersioningProperty
	_jsii_.Get(
		j,
		"versioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) Website() AwsBucket_WebsitePropertyOutputReference {
	var returns AwsBucket_WebsitePropertyOutputReference
	_jsii_.Get(
		j,
		"website",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) WebsiteDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) WebsiteEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucket) WebsiteInput() *AwsBucket_WebsiteProperty {
	var returns *AwsBucket_WebsiteProperty
	_jsii_.Get(
		j,
		"websiteInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket aws_s3_bucket} Resource.
// Experimental.
func NewAwsBucket(scope constructs.Construct, id *string, config *AwsBucketConfig) AwsBucket {
	_init_.Initialize()

	if err := validateNewAwsBucketParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBucket{}

	_jsii_.Create(
		"@cdktn/aws-s3.AwsBucket",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket aws_s3_bucket} Resource.
// Experimental.
func NewAwsBucket_Override(a AwsBucket, scope constructs.Construct, id *string, config *AwsBucketConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.AwsBucket",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsBucket)SetAccelerationStatus(val *string) {
	if err := j.validateSetAccelerationStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accelerationStatus",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetAcl(val *string) {
	if err := j.validateSetAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acl",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetBucketNamespace(val *string) {
	if err := j.validateSetBucketNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketNamespace",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetBucketPrefix(val *string) {
	if err := j.validateSetBucketPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketPrefix",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetObjectLockEnabled(val interface{}) {
	if err := j.validateSetObjectLockEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetPolicy(val *string) {
	if err := j.validateSetPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetRequestPayer(val *string) {
	if err := j.validateSetRequestPayerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestPayer",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsBucket)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a AwsBucket resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsBucket_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsBucket_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-s3.AwsBucket",
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
func AwsBucket_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsBucket_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-s3.AwsBucket",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsBucket_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsBucket_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-s3.AwsBucket",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsBucket_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsBucket_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-s3.AwsBucket",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsBucket_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-s3.AwsBucket",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsBucket) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsBucket) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsBucket) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBucket) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBucket) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBucket) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBucket) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBucket) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBucket) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBucket) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBucket) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBucket) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucket) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsBucket) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBucket) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsBucket) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsBucket) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsBucket) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsBucket) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsBucket) PutCorsRule(value interface{}) {
	if err := a.validatePutCorsRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCorsRule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucket) PutGrant(value interface{}) {
	if err := a.validatePutGrantParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGrant",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucket) PutLifecycleRule(value interface{}) {
	if err := a.validatePutLifecycleRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLifecycleRule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucket) PutLogging(value *AwsBucket_LoggingProperty) {
	if err := a.validatePutLoggingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogging",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucket) PutObjectLockConfiguration(value *AwsBucket_ObjectLockConfigurationProperty) {
	if err := a.validatePutObjectLockConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putObjectLockConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucket) PutReplicationConfiguration(value *AwsBucket_ReplicationConfigurationProperty) {
	if err := a.validatePutReplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putReplicationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucket) PutServerSideEncryptionConfiguration(value *AwsBucket_ServerSideEncryptionConfigurationProperty) {
	if err := a.validatePutServerSideEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServerSideEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucket) PutTimeouts(value *AwsBucket_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucket) PutVersioning(value *AwsBucket_VersioningProperty) {
	if err := a.validatePutVersioningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVersioning",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucket) PutWebsite(value *AwsBucket_WebsiteProperty) {
	if err := a.validatePutWebsiteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWebsite",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucket) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsBucket) ResetAccelerationStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetAccelerationStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetAcl() {
	_jsii_.InvokeVoid(
		a,
		"resetAcl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetBucket() {
	_jsii_.InvokeVoid(
		a,
		"resetBucket",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetBucketNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetBucketNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetBucketPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetBucketPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetCorsRule() {
	_jsii_.InvokeVoid(
		a,
		"resetCorsRule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		a,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetGrant() {
	_jsii_.InvokeVoid(
		a,
		"resetGrant",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetLifecycleRule() {
	_jsii_.InvokeVoid(
		a,
		"resetLifecycleRule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetLogging() {
	_jsii_.InvokeVoid(
		a,
		"resetLogging",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetObjectLockConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetObjectLockConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetObjectLockEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetObjectLockEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetReplicationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetReplicationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetRequestPayer() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestPayer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetServerSideEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetServerSideEncryptionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetVersioning() {
	_jsii_.InvokeVoid(
		a,
		"resetVersioning",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) ResetWebsite() {
	_jsii_.InvokeVoid(
		a,
		"resetWebsite",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucket) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucket) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucket) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucket) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucket) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucket) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucket) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

