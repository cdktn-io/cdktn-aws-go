package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution aws_cloudfront_distribution}.
// Experimental.
type AwsCloudfrontDistribution interface {
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
	CacheTagConfig() AwsCloudfrontDistribution_CacheTagConfigPropertyOutputReference
	// Experimental.
	CacheTagConfigInput() *AwsCloudfrontDistribution_CacheTagConfigProperty
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
	ConnectionFunctionAssociation() AwsCloudfrontDistribution_ConnectionFunctionAssociationPropertyOutputReference
	// Experimental.
	ConnectionFunctionAssociationInput() *AwsCloudfrontDistribution_ConnectionFunctionAssociationProperty
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
	CustomErrorResponse() AwsCloudfrontDistribution_CustomErrorResponsePropertyList
	// Experimental.
	CustomErrorResponseInput() interface{}
	// Experimental.
	DefaultCacheBehavior() AwsCloudfrontDistribution_DefaultCacheBehaviorPropertyOutputReference
	// Experimental.
	DefaultCacheBehaviorInput() *AwsCloudfrontDistribution_DefaultCacheBehaviorProperty
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
	LoggingConfig() AwsCloudfrontDistribution_LoggingConfigPropertyOutputReference
	// Experimental.
	LoggingConfigInput() *AwsCloudfrontDistribution_LoggingConfigProperty
	// Experimental.
	LoggingV1Enabled() cdktn.IResolvable
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OrderedCacheBehavior() AwsCloudfrontDistribution_OrderedCacheBehaviorPropertyList
	// Experimental.
	OrderedCacheBehaviorInput() interface{}
	// Experimental.
	Origin() AwsCloudfrontDistribution_OriginPropertyList
	// Experimental.
	OriginGroup() AwsCloudfrontDistribution_OriginGroupPropertyList
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
	Restrictions() AwsCloudfrontDistribution_RestrictionsPropertyOutputReference
	// Experimental.
	RestrictionsInput() *AwsCloudfrontDistribution_RestrictionsProperty
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
	TrustedKeyGroups() AwsCloudfrontDistribution_TrustedKeyGroupsPropertyList
	// Experimental.
	TrustedSigners() AwsCloudfrontDistribution_TrustedSignersPropertyList
	// Experimental.
	ViewerCertificate() AwsCloudfrontDistribution_ViewerCertificatePropertyOutputReference
	// Experimental.
	ViewerCertificateInput() *AwsCloudfrontDistribution_ViewerCertificateProperty
	// Experimental.
	ViewerMtlsConfig() AwsCloudfrontDistribution_ViewerMtlsConfigPropertyOutputReference
	// Experimental.
	ViewerMtlsConfigInput() *AwsCloudfrontDistribution_ViewerMtlsConfigProperty
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
	PutCacheTagConfig(value *AwsCloudfrontDistribution_CacheTagConfigProperty)
	// Experimental.
	PutConnectionFunctionAssociation(value *AwsCloudfrontDistribution_ConnectionFunctionAssociationProperty)
	// Experimental.
	PutCustomErrorResponse(value interface{})
	// Experimental.
	PutDefaultCacheBehavior(value *AwsCloudfrontDistribution_DefaultCacheBehaviorProperty)
	// Experimental.
	PutLoggingConfig(value *AwsCloudfrontDistribution_LoggingConfigProperty)
	// Experimental.
	PutOrderedCacheBehavior(value interface{})
	// Experimental.
	PutOrigin(value interface{})
	// Experimental.
	PutOriginGroup(value interface{})
	// Experimental.
	PutRestrictions(value *AwsCloudfrontDistribution_RestrictionsProperty)
	// Experimental.
	PutViewerCertificate(value *AwsCloudfrontDistribution_ViewerCertificateProperty)
	// Experimental.
	PutViewerMtlsConfig(value *AwsCloudfrontDistribution_ViewerMtlsConfigProperty)
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

// The jsii proxy struct for AwsCloudfrontDistribution
type jsiiProxy_AwsCloudfrontDistribution struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) AliasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) AnycastIpListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anycastIpListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) AnycastIpListIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anycastIpListIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) CacheTagConfig() AwsCloudfrontDistribution_CacheTagConfigPropertyOutputReference {
	var returns AwsCloudfrontDistribution_CacheTagConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"cacheTagConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) CacheTagConfigInput() *AwsCloudfrontDistribution_CacheTagConfigProperty {
	var returns *AwsCloudfrontDistribution_CacheTagConfigProperty
	_jsii_.Get(
		j,
		"cacheTagConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) CallerReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) ConnectionFunctionAssociation() AwsCloudfrontDistribution_ConnectionFunctionAssociationPropertyOutputReference {
	var returns AwsCloudfrontDistribution_ConnectionFunctionAssociationPropertyOutputReference
	_jsii_.Get(
		j,
		"connectionFunctionAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) ConnectionFunctionAssociationInput() *AwsCloudfrontDistribution_ConnectionFunctionAssociationProperty {
	var returns *AwsCloudfrontDistribution_ConnectionFunctionAssociationProperty
	_jsii_.Get(
		j,
		"connectionFunctionAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) ContinuousDeploymentPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"continuousDeploymentPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) ContinuousDeploymentPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"continuousDeploymentPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) CustomErrorResponse() AwsCloudfrontDistribution_CustomErrorResponsePropertyList {
	var returns AwsCloudfrontDistribution_CustomErrorResponsePropertyList
	_jsii_.Get(
		j,
		"customErrorResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) CustomErrorResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customErrorResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) DefaultCacheBehavior() AwsCloudfrontDistribution_DefaultCacheBehaviorPropertyOutputReference {
	var returns AwsCloudfrontDistribution_DefaultCacheBehaviorPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultCacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) DefaultCacheBehaviorInput() *AwsCloudfrontDistribution_DefaultCacheBehaviorProperty {
	var returns *AwsCloudfrontDistribution_DefaultCacheBehaviorProperty
	_jsii_.Get(
		j,
		"defaultCacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) DefaultRootObject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) DefaultRootObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) HttpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) HttpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) InProgressValidationBatches() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inProgressValidationBatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) IsIpv6Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIpv6Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) IsIpv6EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIpv6EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) LoggingConfig() AwsCloudfrontDistribution_LoggingConfigPropertyOutputReference {
	var returns AwsCloudfrontDistribution_LoggingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) LoggingConfigInput() *AwsCloudfrontDistribution_LoggingConfigProperty {
	var returns *AwsCloudfrontDistribution_LoggingConfigProperty
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) LoggingV1Enabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"loggingV1Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) OrderedCacheBehavior() AwsCloudfrontDistribution_OrderedCacheBehaviorPropertyList {
	var returns AwsCloudfrontDistribution_OrderedCacheBehaviorPropertyList
	_jsii_.Get(
		j,
		"orderedCacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) OrderedCacheBehaviorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orderedCacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Origin() AwsCloudfrontDistribution_OriginPropertyList {
	var returns AwsCloudfrontDistribution_OriginPropertyList
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) OriginGroup() AwsCloudfrontDistribution_OriginGroupPropertyList {
	var returns AwsCloudfrontDistribution_OriginGroupPropertyList
	_jsii_.Get(
		j,
		"originGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) OriginGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) OriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) PriceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) PriceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Restrictions() AwsCloudfrontDistribution_RestrictionsPropertyOutputReference {
	var returns AwsCloudfrontDistribution_RestrictionsPropertyOutputReference
	_jsii_.Get(
		j,
		"restrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) RestrictionsInput() *AwsCloudfrontDistribution_RestrictionsProperty {
	var returns *AwsCloudfrontDistribution_RestrictionsProperty
	_jsii_.Get(
		j,
		"restrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) RetainOnDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) RetainOnDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainOnDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Staging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) StagingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stagingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) TrustedKeyGroups() AwsCloudfrontDistribution_TrustedKeyGroupsPropertyList {
	var returns AwsCloudfrontDistribution_TrustedKeyGroupsPropertyList
	_jsii_.Get(
		j,
		"trustedKeyGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) TrustedSigners() AwsCloudfrontDistribution_TrustedSignersPropertyList {
	var returns AwsCloudfrontDistribution_TrustedSignersPropertyList
	_jsii_.Get(
		j,
		"trustedSigners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) ViewerCertificate() AwsCloudfrontDistribution_ViewerCertificatePropertyOutputReference {
	var returns AwsCloudfrontDistribution_ViewerCertificatePropertyOutputReference
	_jsii_.Get(
		j,
		"viewerCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) ViewerCertificateInput() *AwsCloudfrontDistribution_ViewerCertificateProperty {
	var returns *AwsCloudfrontDistribution_ViewerCertificateProperty
	_jsii_.Get(
		j,
		"viewerCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) ViewerMtlsConfig() AwsCloudfrontDistribution_ViewerMtlsConfigPropertyOutputReference {
	var returns AwsCloudfrontDistribution_ViewerMtlsConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"viewerMtlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) ViewerMtlsConfigInput() *AwsCloudfrontDistribution_ViewerMtlsConfigProperty {
	var returns *AwsCloudfrontDistribution_ViewerMtlsConfigProperty
	_jsii_.Get(
		j,
		"viewerMtlsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) WaitForDeployment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) WaitForDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) WebAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontDistribution) WebAclIdInput() *string {
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
func NewAwsCloudfrontDistribution(scope constructs.Construct, id *string, config *AwsCloudfrontDistributionConfig) AwsCloudfrontDistribution {
	_init_.Initialize()

	if err := validateNewAwsCloudfrontDistributionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCloudfrontDistribution{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsCloudfrontDistribution",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution aws_cloudfront_distribution} Resource.
// Experimental.
func NewAwsCloudfrontDistribution_Override(a AwsCloudfrontDistribution, scope constructs.Construct, id *string, config *AwsCloudfrontDistributionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsCloudfrontDistribution",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetAliases(val *[]*string) {
	if err := j.validateSetAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliases",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetAnycastIpListId(val *string) {
	if err := j.validateSetAnycastIpListIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anycastIpListId",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetContinuousDeploymentPolicyId(val *string) {
	if err := j.validateSetContinuousDeploymentPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"continuousDeploymentPolicyId",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetDefaultRootObject(val *string) {
	if err := j.validateSetDefaultRootObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRootObject",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetHttpVersion(val *string) {
	if err := j.validateSetHttpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpVersion",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetIsIpv6Enabled(val interface{}) {
	if err := j.validateSetIsIpv6EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isIpv6Enabled",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetPriceClass(val *string) {
	if err := j.validateSetPriceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priceClass",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetRetainOnDelete(val interface{}) {
	if err := j.validateSetRetainOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retainOnDelete",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetStaging(val interface{}) {
	if err := j.validateSetStagingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staging",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetWaitForDeployment(val interface{}) {
	if err := j.validateSetWaitForDeploymentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForDeployment",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontDistribution)SetWebAclId(val *string) {
	if err := j.validateSetWebAclIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webAclId",
		val,
	)
}

// Generates CDKTN code for importing a AwsCloudfrontDistribution resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsCloudfrontDistribution_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsCloudfrontDistribution_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.AwsCloudfrontDistribution",
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
func AwsCloudfrontDistribution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCloudfrontDistribution_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.AwsCloudfrontDistribution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCloudfrontDistribution_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCloudfrontDistribution_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.AwsCloudfrontDistribution",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCloudfrontDistribution_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCloudfrontDistribution_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.AwsCloudfrontDistribution",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsCloudfrontDistribution_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-cloudfront.AwsCloudfrontDistribution",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsCloudfrontDistribution) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCloudfrontDistribution) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudfrontDistribution) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCloudfrontDistribution) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCloudfrontDistribution) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCloudfrontDistribution) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCloudfrontDistribution) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCloudfrontDistribution) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCloudfrontDistribution) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCloudfrontDistribution) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudfrontDistribution) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsCloudfrontDistribution) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) PutCacheTagConfig(value *AwsCloudfrontDistribution_CacheTagConfigProperty) {
	if err := a.validatePutCacheTagConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCacheTagConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) PutConnectionFunctionAssociation(value *AwsCloudfrontDistribution_ConnectionFunctionAssociationProperty) {
	if err := a.validatePutConnectionFunctionAssociationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConnectionFunctionAssociation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) PutCustomErrorResponse(value interface{}) {
	if err := a.validatePutCustomErrorResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomErrorResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) PutDefaultCacheBehavior(value *AwsCloudfrontDistribution_DefaultCacheBehaviorProperty) {
	if err := a.validatePutDefaultCacheBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultCacheBehavior",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) PutLoggingConfig(value *AwsCloudfrontDistribution_LoggingConfigProperty) {
	if err := a.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) PutOrderedCacheBehavior(value interface{}) {
	if err := a.validatePutOrderedCacheBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOrderedCacheBehavior",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) PutOrigin(value interface{}) {
	if err := a.validatePutOriginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOrigin",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) PutOriginGroup(value interface{}) {
	if err := a.validatePutOriginGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOriginGroup",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) PutRestrictions(value *AwsCloudfrontDistribution_RestrictionsProperty) {
	if err := a.validatePutRestrictionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRestrictions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) PutViewerCertificate(value *AwsCloudfrontDistribution_ViewerCertificateProperty) {
	if err := a.validatePutViewerCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putViewerCertificate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) PutViewerMtlsConfig(value *AwsCloudfrontDistribution_ViewerMtlsConfigProperty) {
	if err := a.validatePutViewerMtlsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putViewerMtlsConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetAliases() {
	_jsii_.InvokeVoid(
		a,
		"resetAliases",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetAnycastIpListId() {
	_jsii_.InvokeVoid(
		a,
		"resetAnycastIpListId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetCacheTagConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheTagConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetComment() {
	_jsii_.InvokeVoid(
		a,
		"resetComment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetConnectionFunctionAssociation() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionFunctionAssociation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetContinuousDeploymentPolicyId() {
	_jsii_.InvokeVoid(
		a,
		"resetContinuousDeploymentPolicyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetCustomErrorResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomErrorResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetDefaultRootObject() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultRootObject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetHttpVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetIsIpv6Enabled() {
	_jsii_.InvokeVoid(
		a,
		"resetIsIpv6Enabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetOrderedCacheBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetOrderedCacheBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetOriginGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetOriginGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetPriceClass() {
	_jsii_.InvokeVoid(
		a,
		"resetPriceClass",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetRetainOnDelete() {
	_jsii_.InvokeVoid(
		a,
		"resetRetainOnDelete",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetStaging() {
	_jsii_.InvokeVoid(
		a,
		"resetStaging",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetViewerMtlsConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetViewerMtlsConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetWaitForDeployment() {
	_jsii_.InvokeVoid(
		a,
		"resetWaitForDeployment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ResetWebAclId() {
	_jsii_.InvokeVoid(
		a,
		"resetWebAclId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontDistribution) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontDistribution) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontDistribution) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontDistribution) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

