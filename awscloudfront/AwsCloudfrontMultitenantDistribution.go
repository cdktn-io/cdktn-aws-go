package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution aws_cloudfront_multitenant_distribution}.
// Experimental.
type AwsCloudfrontMultitenantDistribution interface {
	cdktn.TerraformResource
	// Experimental.
	ActiveTrustedKeyGroups() AwsCloudfrontMultitenantDistribution_ActiveTrustedKeyGroupsPropertyList
	// Experimental.
	ActiveTrustedKeyGroupsInput() interface{}
	// Experimental.
	Arn() *string
	// Experimental.
	CacheBehavior() AwsCloudfrontMultitenantDistribution_CacheBehaviorPropertyList
	// Experimental.
	CacheBehaviorInput() interface{}
	// Experimental.
	CallerReference() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Comment() *string
	// Experimental.
	SetComment(val *string)
	// Experimental.
	CommentInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConnectionMode() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CustomErrorResponse() AwsCloudfrontMultitenantDistribution_CustomErrorResponsePropertyList
	// Experimental.
	CustomErrorResponseInput() interface{}
	// Experimental.
	DefaultCacheBehavior() AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyList
	// Experimental.
	DefaultCacheBehaviorInput() interface{}
	// Experimental.
	DefaultRootObject() *string
	// Experimental.
	SetDefaultRootObject(val *string)
	// Experimental.
	DefaultRootObjectInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DomainName() *string
	// Experimental.
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	Etag() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HttpVersion() *string
	// Experimental.
	SetHttpVersion(val *string)
	// Experimental.
	HttpVersionInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	InProgressInvalidationBatches() *float64
	// Experimental.
	LastModifiedTime() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Origin() AwsCloudfrontMultitenantDistribution_OriginPropertyList
	// Experimental.
	OriginGroup() AwsCloudfrontMultitenantDistribution_OriginGroupPropertyList
	// Experimental.
	OriginGroupInput() interface{}
	// Experimental.
	OriginInput() interface{}
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
	Restrictions() AwsCloudfrontMultitenantDistribution_RestrictionsPropertyList
	// Experimental.
	RestrictionsInput() interface{}
	// Experimental.
	Status() *string
	// Experimental.
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsAll() cdktn.StringMap
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	TenantConfig() AwsCloudfrontMultitenantDistribution_TenantConfigPropertyList
	// Experimental.
	TenantConfigInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsCloudfrontMultitenantDistribution_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	ViewerCertificate() AwsCloudfrontMultitenantDistribution_ViewerCertificatePropertyList
	// Experimental.
	ViewerCertificateInput() interface{}
	// Experimental.
	WebAclId() *string
	// Experimental.
	SetWebAclId(val *string)
	// Experimental.
	WebAclIdInput() *string
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
	PutActiveTrustedKeyGroups(value interface{})
	// Experimental.
	PutCacheBehavior(value interface{})
	// Experimental.
	PutCustomErrorResponse(value interface{})
	// Experimental.
	PutDefaultCacheBehavior(value interface{})
	// Experimental.
	PutOrigin(value interface{})
	// Experimental.
	PutOriginGroup(value interface{})
	// Experimental.
	PutRestrictions(value interface{})
	// Experimental.
	PutTenantConfig(value interface{})
	// Experimental.
	PutTimeouts(value *AwsCloudfrontMultitenantDistribution_TimeoutsProperty)
	// Experimental.
	PutViewerCertificate(value interface{})
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
	ResetActiveTrustedKeyGroups()
	// Experimental.
	ResetCacheBehavior()
	// Experimental.
	ResetCustomErrorResponse()
	// Experimental.
	ResetDefaultCacheBehavior()
	// Experimental.
	ResetDefaultRootObject()
	// Experimental.
	ResetHttpVersion()
	// Experimental.
	ResetOrigin()
	// Experimental.
	ResetOriginGroup()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRestrictions()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTenantConfig()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetViewerCertificate()
	// Experimental.
	ResetWebAclId()
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

// The jsii proxy struct for AwsCloudfrontMultitenantDistribution
type jsiiProxy_AwsCloudfrontMultitenantDistribution struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) ActiveTrustedKeyGroups() AwsCloudfrontMultitenantDistribution_ActiveTrustedKeyGroupsPropertyList {
	var returns AwsCloudfrontMultitenantDistribution_ActiveTrustedKeyGroupsPropertyList
	_jsii_.Get(
		j,
		"activeTrustedKeyGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) ActiveTrustedKeyGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeTrustedKeyGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) CacheBehavior() AwsCloudfrontMultitenantDistribution_CacheBehaviorPropertyList {
	var returns AwsCloudfrontMultitenantDistribution_CacheBehaviorPropertyList
	_jsii_.Get(
		j,
		"cacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) CacheBehaviorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) CallerReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) ConnectionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) CustomErrorResponse() AwsCloudfrontMultitenantDistribution_CustomErrorResponsePropertyList {
	var returns AwsCloudfrontMultitenantDistribution_CustomErrorResponsePropertyList
	_jsii_.Get(
		j,
		"customErrorResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) CustomErrorResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customErrorResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) DefaultCacheBehavior() AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyList {
	var returns AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyList
	_jsii_.Get(
		j,
		"defaultCacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) DefaultCacheBehaviorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultCacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) DefaultRootObject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) DefaultRootObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) HttpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) HttpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) InProgressInvalidationBatches() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inProgressInvalidationBatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) Origin() AwsCloudfrontMultitenantDistribution_OriginPropertyList {
	var returns AwsCloudfrontMultitenantDistribution_OriginPropertyList
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) OriginGroup() AwsCloudfrontMultitenantDistribution_OriginGroupPropertyList {
	var returns AwsCloudfrontMultitenantDistribution_OriginGroupPropertyList
	_jsii_.Get(
		j,
		"originGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) OriginGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) OriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) Restrictions() AwsCloudfrontMultitenantDistribution_RestrictionsPropertyList {
	var returns AwsCloudfrontMultitenantDistribution_RestrictionsPropertyList
	_jsii_.Get(
		j,
		"restrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) RestrictionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) TenantConfig() AwsCloudfrontMultitenantDistribution_TenantConfigPropertyList {
	var returns AwsCloudfrontMultitenantDistribution_TenantConfigPropertyList
	_jsii_.Get(
		j,
		"tenantConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) TenantConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tenantConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) Timeouts() AwsCloudfrontMultitenantDistribution_TimeoutsPropertyOutputReference {
	var returns AwsCloudfrontMultitenantDistribution_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) ViewerCertificate() AwsCloudfrontMultitenantDistribution_ViewerCertificatePropertyList {
	var returns AwsCloudfrontMultitenantDistribution_ViewerCertificatePropertyList
	_jsii_.Get(
		j,
		"viewerCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) ViewerCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"viewerCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) WebAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution) WebAclIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution aws_cloudfront_multitenant_distribution} Resource.
// Experimental.
func NewAwsCloudfrontMultitenantDistribution(scope constructs.Construct, id *string, config *AwsCloudfrontMultitenantDistributionConfig) AwsCloudfrontMultitenantDistribution {
	_init_.Initialize()

	if err := validateNewAwsCloudfrontMultitenantDistributionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCloudfrontMultitenantDistribution{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsCloudfrontMultitenantDistribution",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution aws_cloudfront_multitenant_distribution} Resource.
// Experimental.
func NewAwsCloudfrontMultitenantDistribution_Override(a AwsCloudfrontMultitenantDistribution, scope constructs.Construct, id *string, config *AwsCloudfrontMultitenantDistributionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsCloudfrontMultitenantDistribution",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution)SetDefaultRootObject(val *string) {
	if err := j.validateSetDefaultRootObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRootObject",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution)SetHttpVersion(val *string) {
	if err := j.validateSetHttpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpVersion",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution)SetWebAclId(val *string) {
	if err := j.validateSetWebAclIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webAclId",
		val,
	)
}

// Generates CDKTN code for importing a AwsCloudfrontMultitenantDistribution resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsCloudfrontMultitenantDistribution_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsCloudfrontMultitenantDistribution_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.AwsCloudfrontMultitenantDistribution",
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
func AwsCloudfrontMultitenantDistribution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCloudfrontMultitenantDistribution_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.AwsCloudfrontMultitenantDistribution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCloudfrontMultitenantDistribution_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCloudfrontMultitenantDistribution_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.AwsCloudfrontMultitenantDistribution",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCloudfrontMultitenantDistribution_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCloudfrontMultitenantDistribution_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.AwsCloudfrontMultitenantDistribution",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsCloudfrontMultitenantDistribution_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-cloudfront.AwsCloudfrontMultitenantDistribution",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) PutActiveTrustedKeyGroups(value interface{}) {
	if err := a.validatePutActiveTrustedKeyGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putActiveTrustedKeyGroups",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) PutCacheBehavior(value interface{}) {
	if err := a.validatePutCacheBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCacheBehavior",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) PutCustomErrorResponse(value interface{}) {
	if err := a.validatePutCustomErrorResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomErrorResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) PutDefaultCacheBehavior(value interface{}) {
	if err := a.validatePutDefaultCacheBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultCacheBehavior",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) PutOrigin(value interface{}) {
	if err := a.validatePutOriginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOrigin",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) PutOriginGroup(value interface{}) {
	if err := a.validatePutOriginGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOriginGroup",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) PutRestrictions(value interface{}) {
	if err := a.validatePutRestrictionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRestrictions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) PutTenantConfig(value interface{}) {
	if err := a.validatePutTenantConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTenantConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) PutTimeouts(value *AwsCloudfrontMultitenantDistribution_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) PutViewerCertificate(value interface{}) {
	if err := a.validatePutViewerCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putViewerCertificate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ResetActiveTrustedKeyGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetActiveTrustedKeyGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ResetCacheBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ResetCustomErrorResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomErrorResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ResetDefaultCacheBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultCacheBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ResetDefaultRootObject() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultRootObject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ResetHttpVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ResetOrigin() {
	_jsii_.InvokeVoid(
		a,
		"resetOrigin",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ResetOriginGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetOriginGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ResetRestrictions() {
	_jsii_.InvokeVoid(
		a,
		"resetRestrictions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ResetTenantConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetTenantConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ResetViewerCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetViewerCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ResetWebAclId() {
	_jsii_.InvokeVoid(
		a,
		"resetWebAclId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

