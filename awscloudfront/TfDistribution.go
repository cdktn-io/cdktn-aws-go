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
type TfDistribution interface {
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
	CacheTagConfig() TfDistribution_CacheTagConfigPropertyOutputReference
	// Experimental.
	CacheTagConfigInput() *TfDistribution_CacheTagConfigProperty
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
	ConnectionFunctionAssociation() TfDistribution_ConnectionFunctionAssociationPropertyOutputReference
	// Experimental.
	ConnectionFunctionAssociationInput() *TfDistribution_ConnectionFunctionAssociationProperty
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
	CustomErrorResponse() TfDistribution_CustomErrorResponsePropertyList
	// Experimental.
	CustomErrorResponseInput() interface{}
	// Experimental.
	DefaultCacheBehavior() TfDistribution_DefaultCacheBehaviorPropertyOutputReference
	// Experimental.
	DefaultCacheBehaviorInput() *TfDistribution_DefaultCacheBehaviorProperty
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
	LoggingConfig() TfDistribution_LoggingConfigPropertyOutputReference
	// Experimental.
	LoggingConfigInput() *TfDistribution_LoggingConfigProperty
	// Experimental.
	LoggingV1Enabled() cdktn.IResolvable
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OrderedCacheBehavior() TfDistribution_OrderedCacheBehaviorPropertyList
	// Experimental.
	OrderedCacheBehaviorInput() interface{}
	// Experimental.
	Origin() TfDistribution_OriginPropertyList
	// Experimental.
	OriginGroup() TfDistribution_OriginGroupPropertyList
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
	Restrictions() TfDistribution_RestrictionsPropertyOutputReference
	// Experimental.
	RestrictionsInput() *TfDistribution_RestrictionsProperty
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
	TrustedKeyGroups() TfDistribution_TrustedKeyGroupsPropertyList
	// Experimental.
	TrustedSigners() TfDistribution_TrustedSignersPropertyList
	// Experimental.
	ViewerCertificate() TfDistribution_ViewerCertificatePropertyOutputReference
	// Experimental.
	ViewerCertificateInput() *TfDistribution_ViewerCertificateProperty
	// Experimental.
	ViewerMtlsConfig() TfDistribution_ViewerMtlsConfigPropertyOutputReference
	// Experimental.
	ViewerMtlsConfigInput() *TfDistribution_ViewerMtlsConfigProperty
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
	PutCacheTagConfig(value *TfDistribution_CacheTagConfigProperty)
	// Experimental.
	PutConnectionFunctionAssociation(value *TfDistribution_ConnectionFunctionAssociationProperty)
	// Experimental.
	PutCustomErrorResponse(value interface{})
	// Experimental.
	PutDefaultCacheBehavior(value *TfDistribution_DefaultCacheBehaviorProperty)
	// Experimental.
	PutLoggingConfig(value *TfDistribution_LoggingConfigProperty)
	// Experimental.
	PutOrderedCacheBehavior(value interface{})
	// Experimental.
	PutOrigin(value interface{})
	// Experimental.
	PutOriginGroup(value interface{})
	// Experimental.
	PutRestrictions(value *TfDistribution_RestrictionsProperty)
	// Experimental.
	PutViewerCertificate(value *TfDistribution_ViewerCertificateProperty)
	// Experimental.
	PutViewerMtlsConfig(value *TfDistribution_ViewerMtlsConfigProperty)
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

// The jsii proxy struct for TfDistribution
type jsiiProxy_TfDistribution struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfDistribution) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) AliasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) AnycastIpListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anycastIpListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) AnycastIpListIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anycastIpListIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) CacheTagConfig() TfDistribution_CacheTagConfigPropertyOutputReference {
	var returns TfDistribution_CacheTagConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"cacheTagConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) CacheTagConfigInput() *TfDistribution_CacheTagConfigProperty {
	var returns *TfDistribution_CacheTagConfigProperty
	_jsii_.Get(
		j,
		"cacheTagConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) CallerReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) ConnectionFunctionAssociation() TfDistribution_ConnectionFunctionAssociationPropertyOutputReference {
	var returns TfDistribution_ConnectionFunctionAssociationPropertyOutputReference
	_jsii_.Get(
		j,
		"connectionFunctionAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) ConnectionFunctionAssociationInput() *TfDistribution_ConnectionFunctionAssociationProperty {
	var returns *TfDistribution_ConnectionFunctionAssociationProperty
	_jsii_.Get(
		j,
		"connectionFunctionAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) ContinuousDeploymentPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"continuousDeploymentPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) ContinuousDeploymentPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"continuousDeploymentPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) CustomErrorResponse() TfDistribution_CustomErrorResponsePropertyList {
	var returns TfDistribution_CustomErrorResponsePropertyList
	_jsii_.Get(
		j,
		"customErrorResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) CustomErrorResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customErrorResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) DefaultCacheBehavior() TfDistribution_DefaultCacheBehaviorPropertyOutputReference {
	var returns TfDistribution_DefaultCacheBehaviorPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultCacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) DefaultCacheBehaviorInput() *TfDistribution_DefaultCacheBehaviorProperty {
	var returns *TfDistribution_DefaultCacheBehaviorProperty
	_jsii_.Get(
		j,
		"defaultCacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) DefaultRootObject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) DefaultRootObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) HttpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) HttpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) InProgressValidationBatches() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inProgressValidationBatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) IsIpv6Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIpv6Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) IsIpv6EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIpv6EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) LoggingConfig() TfDistribution_LoggingConfigPropertyOutputReference {
	var returns TfDistribution_LoggingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) LoggingConfigInput() *TfDistribution_LoggingConfigProperty {
	var returns *TfDistribution_LoggingConfigProperty
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) LoggingV1Enabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"loggingV1Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) OrderedCacheBehavior() TfDistribution_OrderedCacheBehaviorPropertyList {
	var returns TfDistribution_OrderedCacheBehaviorPropertyList
	_jsii_.Get(
		j,
		"orderedCacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) OrderedCacheBehaviorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orderedCacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) Origin() TfDistribution_OriginPropertyList {
	var returns TfDistribution_OriginPropertyList
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) OriginGroup() TfDistribution_OriginGroupPropertyList {
	var returns TfDistribution_OriginGroupPropertyList
	_jsii_.Get(
		j,
		"originGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) OriginGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) OriginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) PriceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) PriceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) Restrictions() TfDistribution_RestrictionsPropertyOutputReference {
	var returns TfDistribution_RestrictionsPropertyOutputReference
	_jsii_.Get(
		j,
		"restrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) RestrictionsInput() *TfDistribution_RestrictionsProperty {
	var returns *TfDistribution_RestrictionsProperty
	_jsii_.Get(
		j,
		"restrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) RetainOnDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) RetainOnDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainOnDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) Staging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) StagingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stagingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) TrustedKeyGroups() TfDistribution_TrustedKeyGroupsPropertyList {
	var returns TfDistribution_TrustedKeyGroupsPropertyList
	_jsii_.Get(
		j,
		"trustedKeyGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) TrustedSigners() TfDistribution_TrustedSignersPropertyList {
	var returns TfDistribution_TrustedSignersPropertyList
	_jsii_.Get(
		j,
		"trustedSigners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) ViewerCertificate() TfDistribution_ViewerCertificatePropertyOutputReference {
	var returns TfDistribution_ViewerCertificatePropertyOutputReference
	_jsii_.Get(
		j,
		"viewerCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) ViewerCertificateInput() *TfDistribution_ViewerCertificateProperty {
	var returns *TfDistribution_ViewerCertificateProperty
	_jsii_.Get(
		j,
		"viewerCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) ViewerMtlsConfig() TfDistribution_ViewerMtlsConfigPropertyOutputReference {
	var returns TfDistribution_ViewerMtlsConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"viewerMtlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) ViewerMtlsConfigInput() *TfDistribution_ViewerMtlsConfigProperty {
	var returns *TfDistribution_ViewerMtlsConfigProperty
	_jsii_.Get(
		j,
		"viewerMtlsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) WaitForDeployment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) WaitForDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) WebAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution) WebAclIdInput() *string {
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
func NewTfDistribution(scope constructs.Construct, id *string, config *TfDistributionConfig) TfDistribution {
	_init_.Initialize()

	if err := validateNewTfDistributionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDistribution{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfDistribution",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution aws_cloudfront_distribution} Resource.
// Experimental.
func NewTfDistribution_Override(t TfDistribution, scope constructs.Construct, id *string, config *TfDistributionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfDistribution",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfDistribution)SetAliases(val *[]*string) {
	if err := j.validateSetAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliases",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetAnycastIpListId(val *string) {
	if err := j.validateSetAnycastIpListIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anycastIpListId",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetContinuousDeploymentPolicyId(val *string) {
	if err := j.validateSetContinuousDeploymentPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"continuousDeploymentPolicyId",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetDefaultRootObject(val *string) {
	if err := j.validateSetDefaultRootObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRootObject",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetHttpVersion(val *string) {
	if err := j.validateSetHttpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpVersion",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetIsIpv6Enabled(val interface{}) {
	if err := j.validateSetIsIpv6EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isIpv6Enabled",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetPriceClass(val *string) {
	if err := j.validateSetPriceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priceClass",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetRetainOnDelete(val interface{}) {
	if err := j.validateSetRetainOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retainOnDelete",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetStaging(val interface{}) {
	if err := j.validateSetStagingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staging",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetWaitForDeployment(val interface{}) {
	if err := j.validateSetWaitForDeploymentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForDeployment",
		val,
	)
}

func (j *jsiiProxy_TfDistribution)SetWebAclId(val *string) {
	if err := j.validateSetWebAclIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webAclId",
		val,
	)
}

// Generates CDKTN code for importing a TfDistribution resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfDistribution_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfDistribution_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.TfDistribution",
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
func TfDistribution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDistribution_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.TfDistribution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfDistribution_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDistribution_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.TfDistribution",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfDistribution_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDistribution_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudfront.TfDistribution",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfDistribution_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-cloudfront.TfDistribution",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfDistribution) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfDistribution) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfDistribution) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDistribution) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistribution) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDistribution) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDistribution) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDistribution) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDistribution) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDistribution) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDistribution) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDistribution) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfDistribution) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistribution) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfDistribution) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfDistribution) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfDistribution) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfDistribution) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfDistribution) PutCacheTagConfig(value *TfDistribution_CacheTagConfigProperty) {
	if err := t.validatePutCacheTagConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCacheTagConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution) PutConnectionFunctionAssociation(value *TfDistribution_ConnectionFunctionAssociationProperty) {
	if err := t.validatePutConnectionFunctionAssociationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConnectionFunctionAssociation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution) PutCustomErrorResponse(value interface{}) {
	if err := t.validatePutCustomErrorResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomErrorResponse",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution) PutDefaultCacheBehavior(value *TfDistribution_DefaultCacheBehaviorProperty) {
	if err := t.validatePutDefaultCacheBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDefaultCacheBehavior",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution) PutLoggingConfig(value *TfDistribution_LoggingConfigProperty) {
	if err := t.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution) PutOrderedCacheBehavior(value interface{}) {
	if err := t.validatePutOrderedCacheBehaviorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOrderedCacheBehavior",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution) PutOrigin(value interface{}) {
	if err := t.validatePutOriginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOrigin",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution) PutOriginGroup(value interface{}) {
	if err := t.validatePutOriginGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOriginGroup",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution) PutRestrictions(value *TfDistribution_RestrictionsProperty) {
	if err := t.validatePutRestrictionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRestrictions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution) PutViewerCertificate(value *TfDistribution_ViewerCertificateProperty) {
	if err := t.validatePutViewerCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putViewerCertificate",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution) PutViewerMtlsConfig(value *TfDistribution_ViewerMtlsConfigProperty) {
	if err := t.validatePutViewerMtlsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putViewerMtlsConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfDistribution) ResetAliases() {
	_jsii_.InvokeVoid(
		t,
		"resetAliases",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetAnycastIpListId() {
	_jsii_.InvokeVoid(
		t,
		"resetAnycastIpListId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetCacheTagConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCacheTagConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetComment() {
	_jsii_.InvokeVoid(
		t,
		"resetComment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetConnectionFunctionAssociation() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectionFunctionAssociation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetContinuousDeploymentPolicyId() {
	_jsii_.InvokeVoid(
		t,
		"resetContinuousDeploymentPolicyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetCustomErrorResponse() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomErrorResponse",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetDefaultRootObject() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultRootObject",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetHttpVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetIsIpv6Enabled() {
	_jsii_.InvokeVoid(
		t,
		"resetIsIpv6Enabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetOrderedCacheBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetOrderedCacheBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetOriginGroup() {
	_jsii_.InvokeVoid(
		t,
		"resetOriginGroup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetPriceClass() {
	_jsii_.InvokeVoid(
		t,
		"resetPriceClass",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetRetainOnDelete() {
	_jsii_.InvokeVoid(
		t,
		"resetRetainOnDelete",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetStaging() {
	_jsii_.InvokeVoid(
		t,
		"resetStaging",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetViewerMtlsConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetViewerMtlsConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetWaitForDeployment() {
	_jsii_.InvokeVoid(
		t,
		"resetWaitForDeployment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) ResetWebAclId() {
	_jsii_.InvokeVoid(
		t,
		"resetWebAclId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

