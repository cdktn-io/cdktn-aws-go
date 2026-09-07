package cloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudfront/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/cloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution aws_cloudfront_distribution}.
// Experimental.
type AwsDistribution interface {
	cdktn.TerraformResource
	// Experimental.
	Aliases() *[]*string
	// Experimental.
	SetAliases(val *[]*string)
	// Experimental.
	AliasesInput() *[]*string
	// Experimental.
	AnycastIpListId() *string
	// Experimental.
	SetAnycastIpListId(val *string)
	// Experimental.
	AnycastIpListIdInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	CacheTagConfig() AwsDistribution_CacheTagConfigPropertyOutputReference
	// Experimental.
	CacheTagConfigInput() *AwsDistribution_CacheTagConfigProperty
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
	ConnectionFunctionAssociation() AwsDistribution_ConnectionFunctionAssociationPropertyOutputReference
	// Experimental.
	ConnectionFunctionAssociationInput() *AwsDistribution_ConnectionFunctionAssociationProperty
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	ContinuousDeploymentPolicyId() *string
	// Experimental.
	SetContinuousDeploymentPolicyId(val *string)
	// Experimental.
	ContinuousDeploymentPolicyIdInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CustomErrorResponse() AwsDistribution_CustomErrorResponsePropertyList
	// Experimental.
	CustomErrorResponseInput() interface{}
	// Experimental.
	DefaultCacheBehavior() AwsDistribution_DefaultCacheBehaviorPropertyOutputReference
	// Experimental.
	DefaultCacheBehaviorInput() *AwsDistribution_DefaultCacheBehaviorProperty
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
	HostedZoneId() *string
	// Experimental.
	HttpVersion() *string
	// Experimental.
	SetHttpVersion(val *string)
	// Experimental.
	HttpVersionInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	InProgressValidationBatches() *float64
	// Experimental.
	IsIpv6Enabled() interface{}
	// Experimental.
	SetIsIpv6Enabled(val interface{})
	// Experimental.
	IsIpv6EnabledInput() interface{}
	// Experimental.
	LastModifiedTime() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LoggingConfig() AwsDistribution_LoggingConfigPropertyOutputReference
	// Experimental.
	LoggingConfigInput() *AwsDistribution_LoggingConfigProperty
	// Experimental.
	LoggingV1Enabled() cdktn.IResolvable
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OrderedCacheBehavior() AwsDistribution_OrderedCacheBehaviorPropertyList
	// Experimental.
	OrderedCacheBehaviorInput() interface{}
	// Experimental.
	Origin() AwsDistribution_OriginPropertyList
	// Experimental.
	OriginGroup() AwsDistribution_OriginGroupPropertyList
	// Experimental.
	OriginGroupInput() interface{}
	// Experimental.
	OriginInput() interface{}
	// Experimental.
	PriceClass() *string
	// Experimental.
	SetPriceClass(val *string)
	// Experimental.
	PriceClassInput() *string
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
	Restrictions() AwsDistribution_RestrictionsPropertyOutputReference
	// Experimental.
	RestrictionsInput() *AwsDistribution_RestrictionsProperty
	// Experimental.
	RetainOnDelete() interface{}
	// Experimental.
	SetRetainOnDelete(val interface{})
	// Experimental.
	RetainOnDeleteInput() interface{}
	// Experimental.
	Staging() interface{}
	// Experimental.
	SetStaging(val interface{})
	// Experimental.
	StagingInput() interface{}
	// Experimental.
	Status() *string
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
	TrustedKeyGroups() AwsDistribution_TrustedKeyGroupsPropertyList
	// Experimental.
	TrustedSigners() AwsDistribution_TrustedSignersPropertyList
	// Experimental.
	ViewerCertificate() AwsDistribution_ViewerCertificatePropertyOutputReference
	// Experimental.
	ViewerCertificateInput() *AwsDistribution_ViewerCertificateProperty
	// Experimental.
	ViewerMtlsConfig() AwsDistribution_ViewerMtlsConfigPropertyOutputReference
	// Experimental.
	ViewerMtlsConfigInput() *AwsDistribution_ViewerMtlsConfigProperty
	// Experimental.
	WaitForDeployment() interface{}
	// Experimental.
	SetWaitForDeployment(val interface{})
	// Experimental.
	WaitForDeploymentInput() interface{}
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
	PutCacheTagConfig(value *AwsDistribution_CacheTagConfigProperty)
	// Experimental.
	PutConnectionFunctionAssociation(value *AwsDistribution_ConnectionFunctionAssociationProperty)
	// Experimental.
	PutCustomErrorResponse(value interface{})
	// Experimental.
	PutDefaultCacheBehavior(value *AwsDistribution_DefaultCacheBehaviorProperty)
	// Experimental.
	PutLoggingConfig(value *AwsDistribution_LoggingConfigProperty)
	// Experimental.
	PutOrderedCacheBehavior(value interface{})
	// Experimental.
	PutOrigin(value interface{})
	// Experimental.
	PutOriginGroup(value interface{})
	// Experimental.
	PutRestrictions(value *AwsDistribution_RestrictionsProperty)
	// Experimental.
	PutViewerCertificate(value *AwsDistribution_ViewerCertificateProperty)
	// Experimental.
	PutViewerMtlsConfig(value *AwsDistribution_ViewerMtlsConfigProperty)
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
	ResetAliases()
	// Experimental.
	ResetAnycastIpListId()
	// Experimental.
	ResetCacheTagConfig()
	// Experimental.
	ResetComment()
	// Experimental.
	ResetConnectionFunctionAssociation()
	// Experimental.
	ResetContinuousDeploymentPolicyId()
	// Experimental.
	ResetCustomErrorResponse()
	// Experimental.
	ResetDefaultRootObject()
	// Experimental.
	ResetHttpVersion()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIsIpv6Enabled()
	// Experimental.
	ResetLoggingConfig()
	// Experimental.
	ResetOrderedCacheBehavior()
	// Experimental.
	ResetOriginGroup()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPriceClass()
	// Experimental.
	ResetRetainOnDelete()
	// Experimental.
	ResetStaging()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetViewerMtlsConfig()
	// Experimental.
	ResetWaitForDeployment()
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

// The jsii proxy struct for AwsDistribution
type jsiiProxy_AwsDistribution struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsDistribution) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) AliasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) AnycastIpListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anycastIpListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) AnycastIpListIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anycastIpListIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) CacheTagConfig() AwsDistribution_CacheTagConfigPropertyOutputReference {
	var returns AwsDistribution_CacheTagConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"cacheTagConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) CacheTagConfigInput() *AwsDistribution_CacheTagConfigProperty {
	var returns *AwsDistribution_CacheTagConfigProperty
	_jsii_.Get(
		j,
		"cacheTagConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) CallerReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) ConnectionFunctionAssociation() AwsDistribution_ConnectionFunctionAssociationPropertyOutputReference {
	var returns AwsDistribution_ConnectionFunctionAssociationPropertyOutputReference
	_jsii_.Get(
		j,
		"connectionFunctionAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) ConnectionFunctionAssociationInput() *AwsDistribution_ConnectionFunctionAssociationProperty {
	var returns *AwsDistribution_ConnectionFunctionAssociationProperty
	_jsii_.Get(
		j,
		"connectionFunctionAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) ContinuousDeploymentPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"continuousDeploymentPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) ContinuousDeploymentPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"continuousDeploymentPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) CustomErrorResponse() AwsDistribution_CustomErrorResponsePropertyList {
	var returns AwsDistribution_CustomErrorResponsePropertyList
	_jsii_.Get(
		j,
		"customErrorResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) CustomErrorResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customErrorResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) DefaultCacheBehavior() AwsDistribution_DefaultCacheBehaviorPropertyOutputReference {
	var returns AwsDistribution_DefaultCacheBehaviorPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultCacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) DefaultCacheBehaviorInput() *AwsDistribution_DefaultCacheBehaviorProperty {
	var returns *AwsDistribution_DefaultCacheBehaviorProperty
	_jsii_.Get(
		j,
		"defaultCacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) DefaultRootObject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) DefaultRootObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) HttpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) HttpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) InProgressValidationBatches() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inProgressValidationBatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) IsIpv6Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIpv6Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) IsIpv6EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIpv6EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) LoggingConfig() AwsDistribution_LoggingConfigPropertyOutputReference {
	var returns AwsDistribution_LoggingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) LoggingConfigInput() *AwsDistribution_LoggingConfigProperty {
	var returns *AwsDistribution_LoggingConfigProperty
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) LoggingV1Enabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"loggingV1Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) OrderedCacheBehavior() AwsDistribution_OrderedCacheBehaviorPropertyList {
	var returns AwsDistribution_OrderedCacheBehaviorPropertyList
	_jsii_.Get(
		j,
		"orderedCacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) OrderedCacheBehaviorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orderedCacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) Origin() AwsDistribution_OriginPropertyList {
	var returns AwsDistribution_OriginPropertyList
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) OriginGroup() AwsDistribution_OriginGroupPropertyList {
	var returns AwsDistribution_OriginGroupPropertyList
	_jsii_.Get(
		j,
		"originGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) OriginGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) OriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) PriceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) PriceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) Restrictions() AwsDistribution_RestrictionsPropertyOutputReference {
	var returns AwsDistribution_RestrictionsPropertyOutputReference
	_jsii_.Get(
		j,
		"restrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) RestrictionsInput() *AwsDistribution_RestrictionsProperty {
	var returns *AwsDistribution_RestrictionsProperty
	_jsii_.Get(
		j,
		"restrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) RetainOnDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) RetainOnDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainOnDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) Staging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) StagingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stagingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) TrustedKeyGroups() AwsDistribution_TrustedKeyGroupsPropertyList {
	var returns AwsDistribution_TrustedKeyGroupsPropertyList
	_jsii_.Get(
		j,
		"trustedKeyGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) TrustedSigners() AwsDistribution_TrustedSignersPropertyList {
	var returns AwsDistribution_TrustedSignersPropertyList
	_jsii_.Get(
		j,
		"trustedSigners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) ViewerCertificate() AwsDistribution_ViewerCertificatePropertyOutputReference {
	var returns AwsDistribution_ViewerCertificatePropertyOutputReference
	_jsii_.Get(
		j,
		"viewerCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) ViewerCertificateInput() *AwsDistribution_ViewerCertificateProperty {
	var returns *AwsDistribution_ViewerCertificateProperty
	_jsii_.Get(
		j,
		"viewerCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) ViewerMtlsConfig() AwsDistribution_ViewerMtlsConfigPropertyOutputReference {
	var returns AwsDistribution_ViewerMtlsConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"viewerMtlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) ViewerMtlsConfigInput() *AwsDistribution_ViewerMtlsConfigProperty {
	var returns *AwsDistribution_ViewerMtlsConfigProperty
	_jsii_.Get(
		j,
		"viewerMtlsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) WaitForDeployment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) WaitForDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) WebAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution) WebAclIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution aws_cloudfront_distribution} Resource.
// Experimental.
func NewAwsDistribution(scope constructs.Construct, id *string, config *AwsDistributionConfig) AwsDistribution {
	_init_.Initialize()

	if err := validateNewAwsDistributionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDistribution{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsDistribution",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution aws_cloudfront_distribution} Resource.
// Experimental.
func NewAwsDistribution_Override(a AwsDistribution, scope constructs.Construct, id *string, config *AwsDistributionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsDistribution",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsDistribution)SetAliases(val *[]*string) {
	if err := j.validateSetAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliases",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetAnycastIpListId(val *string) {
	if err := j.validateSetAnycastIpListIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anycastIpListId",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetContinuousDeploymentPolicyId(val *string) {
	if err := j.validateSetContinuousDeploymentPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"continuousDeploymentPolicyId",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetDefaultRootObject(val *string) {
	if err := j.validateSetDefaultRootObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRootObject",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetHttpVersion(val *string) {
	if err := j.validateSetHttpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpVersion",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetIsIpv6Enabled(val interface{}) {
	if err := j.validateSetIsIpv6EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isIpv6Enabled",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetPriceClass(val *string) {
	if err := j.validateSetPriceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priceClass",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetRetainOnDelete(val interface{}) {
	if err := j.validateSetRetainOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retainOnDelete",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetStaging(val interface{}) {
	if err := j.validateSetStagingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staging",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetWaitForDeployment(val interface{}) {
	if err := j.validateSetWaitForDeploymentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForDeployment",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution)SetWebAclId(val *string) {
	if err := j.validateSetWebAclIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webAclId",
		val,
	)
}

// Generates CDKTN code for importing a AwsDistribution resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsDistribution_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsDistribution_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.AwsDistribution",
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
func AwsDistribution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDistribution_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.AwsDistribution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsDistribution_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDistribution_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.AwsDistribution",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsDistribution_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDistribution_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.AwsDistribution",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsDistribution_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-cloudfront.AwsDistribution",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsDistribution) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsDistribution) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsDistribution) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDistribution) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDistribution) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDistribution) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDistribution) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDistribution) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDistribution) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDistribution) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDistribution) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDistribution) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDistribution) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsDistribution) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDistribution) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsDistribution) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsDistribution) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsDistribution) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsDistribution) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsDistribution) PutCacheTagConfig(value *AwsDistribution_CacheTagConfigProperty) {
	if err := a.validatePutCacheTagConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCacheTagConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution) PutConnectionFunctionAssociation(value *AwsDistribution_ConnectionFunctionAssociationProperty) {
	if err := a.validatePutConnectionFunctionAssociationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConnectionFunctionAssociation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution) PutCustomErrorResponse(value interface{}) {
	if err := a.validatePutCustomErrorResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomErrorResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution) PutDefaultCacheBehavior(value *AwsDistribution_DefaultCacheBehaviorProperty) {
	if err := a.validatePutDefaultCacheBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultCacheBehavior",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution) PutLoggingConfig(value *AwsDistribution_LoggingConfigProperty) {
	if err := a.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution) PutOrderedCacheBehavior(value interface{}) {
	if err := a.validatePutOrderedCacheBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOrderedCacheBehavior",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution) PutOrigin(value interface{}) {
	if err := a.validatePutOriginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOrigin",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution) PutOriginGroup(value interface{}) {
	if err := a.validatePutOriginGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOriginGroup",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution) PutRestrictions(value *AwsDistribution_RestrictionsProperty) {
	if err := a.validatePutRestrictionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRestrictions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution) PutViewerCertificate(value *AwsDistribution_ViewerCertificateProperty) {
	if err := a.validatePutViewerCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putViewerCertificate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution) PutViewerMtlsConfig(value *AwsDistribution_ViewerMtlsConfigProperty) {
	if err := a.validatePutViewerMtlsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putViewerMtlsConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsDistribution) ResetAliases() {
	_jsii_.InvokeVoid(
		a,
		"resetAliases",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetAnycastIpListId() {
	_jsii_.InvokeVoid(
		a,
		"resetAnycastIpListId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetCacheTagConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheTagConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetComment() {
	_jsii_.InvokeVoid(
		a,
		"resetComment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetConnectionFunctionAssociation() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionFunctionAssociation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetContinuousDeploymentPolicyId() {
	_jsii_.InvokeVoid(
		a,
		"resetContinuousDeploymentPolicyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetCustomErrorResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomErrorResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetDefaultRootObject() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultRootObject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetHttpVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetIsIpv6Enabled() {
	_jsii_.InvokeVoid(
		a,
		"resetIsIpv6Enabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetOrderedCacheBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetOrderedCacheBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetOriginGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetOriginGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetPriceClass() {
	_jsii_.InvokeVoid(
		a,
		"resetPriceClass",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetRetainOnDelete() {
	_jsii_.InvokeVoid(
		a,
		"resetRetainOnDelete",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetStaging() {
	_jsii_.InvokeVoid(
		a,
		"resetStaging",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetViewerMtlsConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetViewerMtlsConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetWaitForDeployment() {
	_jsii_.InvokeVoid(
		a,
		"resetWaitForDeployment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) ResetWebAclId() {
	_jsii_.InvokeVoid(
		a,
		"resetWebAclId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDistribution) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDistribution) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDistribution) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDistribution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDistribution) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDistribution) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

