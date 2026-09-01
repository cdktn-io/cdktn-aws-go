package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl aws_wafv2_web_acl}.
// Experimental.
type AwsWafv2WebAcl interface {
	cdktn.TerraformResource
	// Experimental.
	ApplicationIntegrationUrl() *string
	// Experimental.
	Arn() *string
	// Experimental.
	AssociationConfig() AwsWafv2WebAcl_AssociationConfigPropertyOutputReference
	// Experimental.
	AssociationConfigInput() *AwsWafv2WebAcl_AssociationConfigProperty
	// Experimental.
	Capacity() *float64
	// Experimental.
	CaptchaConfig() AwsWafv2WebAcl_CaptchaConfigPropertyOutputReference
	// Experimental.
	CaptchaConfigInput() *AwsWafv2WebAcl_CaptchaConfigProperty
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ChallengeConfig() AwsWafv2WebAcl_ChallengeConfigPropertyOutputReference
	// Experimental.
	ChallengeConfigInput() *AwsWafv2WebAcl_ChallengeConfigProperty
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
	CustomResponseBody() AwsWafv2WebAcl_CustomResponseBodyPropertyList
	// Experimental.
	CustomResponseBodyInput() interface{}
	// Experimental.
	DataProtectionConfig() AwsWafv2WebAcl_DataProtectionConfigPropertyOutputReference
	// Experimental.
	DataProtectionConfigInput() *AwsWafv2WebAcl_DataProtectionConfigProperty
	// Experimental.
	DefaultAction() AwsWafv2WebAcl_DefaultActionPropertyOutputReference
	// Experimental.
	DefaultActionInput() *AwsWafv2WebAcl_DefaultActionProperty
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
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
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LockToken() *string
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NamePrefix() *string
	// Experimental.
	SetNamePrefix(val *string)
	// Experimental.
	NamePrefixInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	Rule() AwsWafv2WebAcl_RulePropertyList
	// Experimental.
	RuleInput() interface{}
	// Experimental.
	RuleJson() *string
	// Experimental.
	SetRuleJson(val *string)
	// Experimental.
	RuleJsonInput() *string
	// Experimental.
	Scope() *string
	// Experimental.
	SetScope(val *string)
	// Experimental.
	ScopeInput() *string
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
	TokenDomains() *[]*string
	// Experimental.
	SetTokenDomains(val *[]*string)
	// Experimental.
	TokenDomainsInput() *[]*string
	// Experimental.
	VisibilityConfig() AwsWafv2WebAcl_VisibilityConfigPropertyOutputReference
	// Experimental.
	VisibilityConfigInput() *AwsWafv2WebAcl_VisibilityConfigProperty
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
	PutAssociationConfig(value *AwsWafv2WebAcl_AssociationConfigProperty)
	// Experimental.
	PutCaptchaConfig(value *AwsWafv2WebAcl_CaptchaConfigProperty)
	// Experimental.
	PutChallengeConfig(value *AwsWafv2WebAcl_ChallengeConfigProperty)
	// Experimental.
	PutCustomResponseBody(value interface{})
	// Experimental.
	PutDataProtectionConfig(value *AwsWafv2WebAcl_DataProtectionConfigProperty)
	// Experimental.
	PutDefaultAction(value *AwsWafv2WebAcl_DefaultActionProperty)
	// Experimental.
	PutRule(value interface{})
	// Experimental.
	PutVisibilityConfig(value *AwsWafv2WebAcl_VisibilityConfigProperty)
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
	ResetAssociationConfig()
	// Experimental.
	ResetCaptchaConfig()
	// Experimental.
	ResetChallengeConfig()
	// Experimental.
	ResetCustomResponseBody()
	// Experimental.
	ResetDataProtectionConfig()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetId()
	// Experimental.
	ResetName()
	// Experimental.
	ResetNamePrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRule()
	// Experimental.
	ResetRuleJson()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTokenDomains()
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

// The jsii proxy struct for AwsWafv2WebAcl
type jsiiProxy_AwsWafv2WebAcl struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsWafv2WebAcl) ApplicationIntegrationUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIntegrationUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) AssociationConfig() AwsWafv2WebAcl_AssociationConfigPropertyOutputReference {
	var returns AwsWafv2WebAcl_AssociationConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"associationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) AssociationConfigInput() *AwsWafv2WebAcl_AssociationConfigProperty {
	var returns *AwsWafv2WebAcl_AssociationConfigProperty
	_jsii_.Get(
		j,
		"associationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) Capacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) CaptchaConfig() AwsWafv2WebAcl_CaptchaConfigPropertyOutputReference {
	var returns AwsWafv2WebAcl_CaptchaConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"captchaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) CaptchaConfigInput() *AwsWafv2WebAcl_CaptchaConfigProperty {
	var returns *AwsWafv2WebAcl_CaptchaConfigProperty
	_jsii_.Get(
		j,
		"captchaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) ChallengeConfig() AwsWafv2WebAcl_ChallengeConfigPropertyOutputReference {
	var returns AwsWafv2WebAcl_ChallengeConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"challengeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) ChallengeConfigInput() *AwsWafv2WebAcl_ChallengeConfigProperty {
	var returns *AwsWafv2WebAcl_ChallengeConfigProperty
	_jsii_.Get(
		j,
		"challengeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) CustomResponseBody() AwsWafv2WebAcl_CustomResponseBodyPropertyList {
	var returns AwsWafv2WebAcl_CustomResponseBodyPropertyList
	_jsii_.Get(
		j,
		"customResponseBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) CustomResponseBodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customResponseBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) DataProtectionConfig() AwsWafv2WebAcl_DataProtectionConfigPropertyOutputReference {
	var returns AwsWafv2WebAcl_DataProtectionConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"dataProtectionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) DataProtectionConfigInput() *AwsWafv2WebAcl_DataProtectionConfigProperty {
	var returns *AwsWafv2WebAcl_DataProtectionConfigProperty
	_jsii_.Get(
		j,
		"dataProtectionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) DefaultAction() AwsWafv2WebAcl_DefaultActionPropertyOutputReference {
	var returns AwsWafv2WebAcl_DefaultActionPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) DefaultActionInput() *AwsWafv2WebAcl_DefaultActionProperty {
	var returns *AwsWafv2WebAcl_DefaultActionProperty
	_jsii_.Get(
		j,
		"defaultActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) LockToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lockToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) Rule() AwsWafv2WebAcl_RulePropertyList {
	var returns AwsWafv2WebAcl_RulePropertyList
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) RuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) RuleJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) RuleJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) TokenDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) TokenDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) VisibilityConfig() AwsWafv2WebAcl_VisibilityConfigPropertyOutputReference {
	var returns AwsWafv2WebAcl_VisibilityConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"visibilityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAcl) VisibilityConfigInput() *AwsWafv2WebAcl_VisibilityConfigProperty {
	var returns *AwsWafv2WebAcl_VisibilityConfigProperty
	_jsii_.Get(
		j,
		"visibilityConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl aws_wafv2_web_acl} Resource.
// Experimental.
func NewAwsWafv2WebAcl(scope constructs.Construct, id *string, config *AwsWafv2WebAclConfig) AwsWafv2WebAcl {
	_init_.Initialize()

	if err := validateNewAwsWafv2WebAclParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWafv2WebAcl{}

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWafv2WebAcl",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl aws_wafv2_web_acl} Resource.
// Experimental.
func NewAwsWafv2WebAcl_Override(a AwsWafv2WebAcl, scope constructs.Construct, id *string, config *AwsWafv2WebAclConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWafv2WebAcl",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl)SetRuleJson(val *string) {
	if err := j.validateSetRuleJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleJson",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAcl)SetTokenDomains(val *[]*string) {
	if err := j.validateSetTokenDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenDomains",
		val,
	)
}

// Generates CDKTN code for importing a AwsWafv2WebAcl resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsWafv2WebAcl_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsWafv2WebAcl_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-waf.AwsWafv2WebAcl",
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
func AwsWafv2WebAcl_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsWafv2WebAcl_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-waf.AwsWafv2WebAcl",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsWafv2WebAcl_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsWafv2WebAcl_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-waf.AwsWafv2WebAcl",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsWafv2WebAcl_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsWafv2WebAcl_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-waf.AwsWafv2WebAcl",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsWafv2WebAcl_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-waf.AwsWafv2WebAcl",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsWafv2WebAcl) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWafv2WebAcl) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWafv2WebAcl) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWafv2WebAcl) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWafv2WebAcl) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWafv2WebAcl) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWafv2WebAcl) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWafv2WebAcl) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWafv2WebAcl) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWafv2WebAcl) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafv2WebAcl) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWafv2WebAcl) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsWafv2WebAcl) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) PutAssociationConfig(value *AwsWafv2WebAcl_AssociationConfigProperty) {
	if err := a.validatePutAssociationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAssociationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) PutCaptchaConfig(value *AwsWafv2WebAcl_CaptchaConfigProperty) {
	if err := a.validatePutCaptchaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCaptchaConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) PutChallengeConfig(value *AwsWafv2WebAcl_ChallengeConfigProperty) {
	if err := a.validatePutChallengeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putChallengeConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) PutCustomResponseBody(value interface{}) {
	if err := a.validatePutCustomResponseBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomResponseBody",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) PutDataProtectionConfig(value *AwsWafv2WebAcl_DataProtectionConfigProperty) {
	if err := a.validatePutDataProtectionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataProtectionConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) PutDefaultAction(value *AwsWafv2WebAcl_DefaultActionProperty) {
	if err := a.validatePutDefaultActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) PutRule(value interface{}) {
	if err := a.validatePutRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) PutVisibilityConfig(value *AwsWafv2WebAcl_VisibilityConfigProperty) {
	if err := a.validatePutVisibilityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVisibilityConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) ResetAssociationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAssociationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) ResetCaptchaConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCaptchaConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) ResetChallengeConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetChallengeConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) ResetCustomResponseBody() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomResponseBody",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) ResetDataProtectionConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDataProtectionConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) ResetRule() {
	_jsii_.InvokeVoid(
		a,
		"resetRule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) ResetRuleJson() {
	_jsii_.InvokeVoid(
		a,
		"resetRuleJson",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) ResetTokenDomains() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenDomains",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAcl) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafv2WebAcl) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafv2WebAcl) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafv2WebAcl) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafv2WebAcl) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafv2WebAcl) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafv2WebAcl) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

