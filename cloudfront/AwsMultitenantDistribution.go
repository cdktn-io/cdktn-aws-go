package cloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudfront/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/cloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution aws_cloudfront_multitenant_distribution}.
// Experimental.
type AwsMultitenantDistribution interface {
	cdktn.TerraformResource
	// Experimental.
	ActiveTrustedKeyGroups() AwsMultitenantDistribution_ActiveTrustedKeyGroupsPropertyList
	// Experimental.
	ActiveTrustedKeyGroupsInput() interface{}
	// Experimental.
	Arn() *string
	// Experimental.
	CacheBehavior() AwsMultitenantDistribution_CacheBehaviorPropertyList
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
	CustomErrorResponse() AwsMultitenantDistribution_CustomErrorResponsePropertyList
	// Experimental.
	CustomErrorResponseInput() interface{}
	// Experimental.
	DefaultCacheBehavior() AwsMultitenantDistribution_DefaultCacheBehaviorPropertyList
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
	Origin() AwsMultitenantDistribution_OriginPropertyList
	// Experimental.
	OriginGroup() AwsMultitenantDistribution_OriginGroupPropertyList
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
	Restrictions() AwsMultitenantDistribution_RestrictionsPropertyList
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
	TenantConfig() AwsMultitenantDistribution_TenantConfigPropertyList
	// Experimental.
	TenantConfigInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsMultitenantDistribution_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	ViewerCertificate() AwsMultitenantDistribution_ViewerCertificatePropertyList
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
	PutTimeouts(value *AwsMultitenantDistribution_TimeoutsProperty)
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

// The jsii proxy struct for AwsMultitenantDistribution
type jsiiProxy_AwsMultitenantDistribution struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsMultitenantDistribution) ActiveTrustedKeyGroups() AwsMultitenantDistribution_ActiveTrustedKeyGroupsPropertyList {
	var returns AwsMultitenantDistribution_ActiveTrustedKeyGroupsPropertyList
	_jsii_.Get(
		j,
		"activeTrustedKeyGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) ActiveTrustedKeyGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeTrustedKeyGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) CacheBehavior() AwsMultitenantDistribution_CacheBehaviorPropertyList {
	var returns AwsMultitenantDistribution_CacheBehaviorPropertyList
	_jsii_.Get(
		j,
		"cacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) CacheBehaviorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) CallerReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) ConnectionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) CustomErrorResponse() AwsMultitenantDistribution_CustomErrorResponsePropertyList {
	var returns AwsMultitenantDistribution_CustomErrorResponsePropertyList
	_jsii_.Get(
		j,
		"customErrorResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) CustomErrorResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customErrorResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) DefaultCacheBehavior() AwsMultitenantDistribution_DefaultCacheBehaviorPropertyList {
	var returns AwsMultitenantDistribution_DefaultCacheBehaviorPropertyList
	_jsii_.Get(
		j,
		"defaultCacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) DefaultCacheBehaviorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultCacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) DefaultRootObject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) DefaultRootObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) HttpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) HttpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) InProgressInvalidationBatches() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inProgressInvalidationBatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) Origin() AwsMultitenantDistribution_OriginPropertyList {
	var returns AwsMultitenantDistribution_OriginPropertyList
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) OriginGroup() AwsMultitenantDistribution_OriginGroupPropertyList {
	var returns AwsMultitenantDistribution_OriginGroupPropertyList
	_jsii_.Get(
		j,
		"originGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) OriginGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) OriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) Restrictions() AwsMultitenantDistribution_RestrictionsPropertyList {
	var returns AwsMultitenantDistribution_RestrictionsPropertyList
	_jsii_.Get(
		j,
		"restrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) RestrictionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) TenantConfig() AwsMultitenantDistribution_TenantConfigPropertyList {
	var returns AwsMultitenantDistribution_TenantConfigPropertyList
	_jsii_.Get(
		j,
		"tenantConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) TenantConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tenantConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) Timeouts() AwsMultitenantDistribution_TimeoutsPropertyOutputReference {
	var returns AwsMultitenantDistribution_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) ViewerCertificate() AwsMultitenantDistribution_ViewerCertificatePropertyList {
	var returns AwsMultitenantDistribution_ViewerCertificatePropertyList
	_jsii_.Get(
		j,
		"viewerCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) ViewerCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"viewerCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) WebAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMultitenantDistribution) WebAclIdInput() *string {
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
func NewAwsMultitenantDistribution(scope constructs.Construct, id *string, config *AwsMultitenantDistributionConfig) AwsMultitenantDistribution {
	_init_.Initialize()

	if err := validateNewAwsMultitenantDistributionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMultitenantDistribution{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsMultitenantDistribution",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution aws_cloudfront_multitenant_distribution} Resource.
// Experimental.
func NewAwsMultitenantDistribution_Override(a AwsMultitenantDistribution, scope constructs.Construct, id *string, config *AwsMultitenantDistributionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsMultitenantDistribution",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution)SetDefaultRootObject(val *string) {
	if err := j.validateSetDefaultRootObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRootObject",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution)SetHttpVersion(val *string) {
	if err := j.validateSetHttpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpVersion",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsMultitenantDistribution)SetWebAclId(val *string) {
	if err := j.validateSetWebAclIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webAclId",
		val,
	)
}

// Generates CDKTN code for importing a AwsMultitenantDistribution resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsMultitenantDistribution_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsMultitenantDistribution_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.AwsMultitenantDistribution",
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
func AwsMultitenantDistribution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsMultitenantDistribution_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.AwsMultitenantDistribution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsMultitenantDistribution_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsMultitenantDistribution_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.AwsMultitenantDistribution",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsMultitenantDistribution_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsMultitenantDistribution_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.AwsMultitenantDistribution",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsMultitenantDistribution_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-cloudfront.AwsMultitenantDistribution",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsMultitenantDistribution) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMultitenantDistribution) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMultitenantDistribution) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMultitenantDistribution) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMultitenantDistribution) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMultitenantDistribution) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMultitenantDistribution) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMultitenantDistribution) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMultitenantDistribution) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMultitenantDistribution) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMultitenantDistribution) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMultitenantDistribution) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsMultitenantDistribution) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) PutActiveTrustedKeyGroups(value interface{}) {
	if err := a.validatePutActiveTrustedKeyGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putActiveTrustedKeyGroups",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) PutCacheBehavior(value interface{}) {
	if err := a.validatePutCacheBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCacheBehavior",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) PutCustomErrorResponse(value interface{}) {
	if err := a.validatePutCustomErrorResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomErrorResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) PutDefaultCacheBehavior(value interface{}) {
	if err := a.validatePutDefaultCacheBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultCacheBehavior",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) PutOrigin(value interface{}) {
	if err := a.validatePutOriginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOrigin",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) PutOriginGroup(value interface{}) {
	if err := a.validatePutOriginGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOriginGroup",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) PutRestrictions(value interface{}) {
	if err := a.validatePutRestrictionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRestrictions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) PutTenantConfig(value interface{}) {
	if err := a.validatePutTenantConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTenantConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) PutTimeouts(value *AwsMultitenantDistribution_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) PutViewerCertificate(value interface{}) {
	if err := a.validatePutViewerCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putViewerCertificate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) ResetActiveTrustedKeyGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetActiveTrustedKeyGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) ResetCacheBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) ResetCustomErrorResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomErrorResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) ResetDefaultCacheBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultCacheBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) ResetDefaultRootObject() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultRootObject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) ResetHttpVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) ResetOrigin() {
	_jsii_.InvokeVoid(
		a,
		"resetOrigin",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) ResetOriginGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetOriginGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) ResetRestrictions() {
	_jsii_.InvokeVoid(
		a,
		"resetRestrictions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) ResetTenantConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetTenantConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) ResetViewerCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetViewerCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) ResetWebAclId() {
	_jsii_.InvokeVoid(
		a,
		"resetWebAclId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMultitenantDistribution) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMultitenantDistribution) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMultitenantDistribution) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMultitenantDistribution) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMultitenantDistribution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMultitenantDistribution) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMultitenantDistribution) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

