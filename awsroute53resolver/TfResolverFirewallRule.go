package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsroute53resolver/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsroute53resolver/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule aws_route53_resolver_firewall_rule}.
// Experimental.
type TfResolverFirewallRule interface {
	cdktn.TerraformResource
	// Experimental.
	Action() *string
	// Experimental.
	SetAction(val *string)
	// Experimental.
	ActionInput() *string
	// Experimental.
	BlockOverrideDnsType() *string
	// Experimental.
	SetBlockOverrideDnsType(val *string)
	// Experimental.
	BlockOverrideDnsTypeInput() *string
	// Experimental.
	BlockOverrideDomain() *string
	// Experimental.
	SetBlockOverrideDomain(val *string)
	// Experimental.
	BlockOverrideDomainInput() *string
	// Experimental.
	BlockOverrideTtl() *float64
	// Experimental.
	SetBlockOverrideTtl(val *float64)
	// Experimental.
	BlockOverrideTtlInput() *float64
	// Experimental.
	BlockResponse() *string
	// Experimental.
	SetBlockResponse(val *string)
	// Experimental.
	BlockResponseInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConfidenceThreshold() *string
	// Experimental.
	SetConfidenceThreshold(val *string)
	// Experimental.
	ConfidenceThresholdInput() *string
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
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DnsThreatProtection() *string
	// Experimental.
	SetDnsThreatProtection(val *string)
	// Experimental.
	DnsThreatProtectionInput() *string
	// Experimental.
	FirewallDomainListId() *string
	// Experimental.
	SetFirewallDomainListId(val *string)
	// Experimental.
	FirewallDomainListIdInput() *string
	// Experimental.
	FirewallDomainRedirectionAction() *string
	// Experimental.
	SetFirewallDomainRedirectionAction(val *string)
	// Experimental.
	FirewallDomainRedirectionActionInput() *string
	// Experimental.
	FirewallRuleGroupId() *string
	// Experimental.
	SetFirewallRuleGroupId(val *string)
	// Experimental.
	FirewallRuleGroupIdInput() *string
	// Experimental.
	FirewallThreatProtectionId() *string
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
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Priority() *float64
	// Experimental.
	SetPriority(val *float64)
	// Experimental.
	PriorityInput() *float64
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	QType() *string
	// Experimental.
	SetQType(val *string)
	// Experimental.
	QTypeInput() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetBlockOverrideDnsType()
	// Experimental.
	ResetBlockOverrideDomain()
	// Experimental.
	ResetBlockOverrideTtl()
	// Experimental.
	ResetBlockResponse()
	// Experimental.
	ResetConfidenceThreshold()
	// Experimental.
	ResetDnsThreatProtection()
	// Experimental.
	ResetFirewallDomainListId()
	// Experimental.
	ResetFirewallDomainRedirectionAction()
	// Experimental.
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetQType()
	// Experimental.
	ResetRegion()
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

// The jsii proxy struct for TfResolverFirewallRule
type jsiiProxy_TfResolverFirewallRule struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfResolverFirewallRule) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) BlockOverrideDnsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDnsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) BlockOverrideDnsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDnsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) BlockOverrideDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) BlockOverrideDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) BlockOverrideTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockOverrideTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) BlockOverrideTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockOverrideTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) BlockResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) BlockResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) ConfidenceThreshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidenceThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) ConfidenceThresholdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidenceThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) DnsThreatProtection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsThreatProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) DnsThreatProtectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsThreatProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) FirewallDomainListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) FirewallDomainListIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) FirewallDomainRedirectionAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainRedirectionAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) FirewallDomainRedirectionActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainRedirectionActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) FirewallRuleGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) FirewallRuleGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) FirewallThreatProtectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallThreatProtectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) QType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) QTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfResolverFirewallRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule aws_route53_resolver_firewall_rule} Resource.
// Experimental.
func NewTfResolverFirewallRule(scope constructs.Construct, id *string, config *TfResolverFirewallRuleConfig) TfResolverFirewallRule {
	_init_.Initialize()

	if err := validateNewTfResolverFirewallRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfResolverFirewallRule{}

	_jsii_.Create(
		"@cdktn/aws-route-53-resolver.TfResolverFirewallRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule aws_route53_resolver_firewall_rule} Resource.
// Experimental.
func NewTfResolverFirewallRule_Override(t TfResolverFirewallRule, scope constructs.Construct, id *string, config *TfResolverFirewallRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53-resolver.TfResolverFirewallRule",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetBlockOverrideDnsType(val *string) {
	if err := j.validateSetBlockOverrideDnsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockOverrideDnsType",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetBlockOverrideDomain(val *string) {
	if err := j.validateSetBlockOverrideDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockOverrideDomain",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetBlockOverrideTtl(val *float64) {
	if err := j.validateSetBlockOverrideTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockOverrideTtl",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetBlockResponse(val *string) {
	if err := j.validateSetBlockResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockResponse",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetConfidenceThreshold(val *string) {
	if err := j.validateSetConfidenceThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidenceThreshold",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetDnsThreatProtection(val *string) {
	if err := j.validateSetDnsThreatProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsThreatProtection",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetFirewallDomainListId(val *string) {
	if err := j.validateSetFirewallDomainListIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallDomainListId",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetFirewallDomainRedirectionAction(val *string) {
	if err := j.validateSetFirewallDomainRedirectionActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallDomainRedirectionAction",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetFirewallRuleGroupId(val *string) {
	if err := j.validateSetFirewallRuleGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallRuleGroupId",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetQType(val *string) {
	if err := j.validateSetQTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qType",
		val,
	)
}

func (j *jsiiProxy_TfResolverFirewallRule)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTN code for importing a TfResolverFirewallRule resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfResolverFirewallRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfResolverFirewallRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53-resolver.TfResolverFirewallRule",
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
func TfResolverFirewallRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfResolverFirewallRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53-resolver.TfResolverFirewallRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfResolverFirewallRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfResolverFirewallRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53-resolver.TfResolverFirewallRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfResolverFirewallRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfResolverFirewallRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53-resolver.TfResolverFirewallRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfResolverFirewallRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-route-53-resolver.TfResolverFirewallRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfResolverFirewallRule) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfResolverFirewallRule) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfResolverFirewallRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfResolverFirewallRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfResolverFirewallRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfResolverFirewallRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfResolverFirewallRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfResolverFirewallRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfResolverFirewallRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfResolverFirewallRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResolverFirewallRule) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfResolverFirewallRule) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfResolverFirewallRule) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) ResetBlockOverrideDnsType() {
	_jsii_.InvokeVoid(
		t,
		"resetBlockOverrideDnsType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) ResetBlockOverrideDomain() {
	_jsii_.InvokeVoid(
		t,
		"resetBlockOverrideDomain",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) ResetBlockOverrideTtl() {
	_jsii_.InvokeVoid(
		t,
		"resetBlockOverrideTtl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) ResetBlockResponse() {
	_jsii_.InvokeVoid(
		t,
		"resetBlockResponse",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) ResetConfidenceThreshold() {
	_jsii_.InvokeVoid(
		t,
		"resetConfidenceThreshold",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) ResetDnsThreatProtection() {
	_jsii_.InvokeVoid(
		t,
		"resetDnsThreatProtection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) ResetFirewallDomainListId() {
	_jsii_.InvokeVoid(
		t,
		"resetFirewallDomainListId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) ResetFirewallDomainRedirectionAction() {
	_jsii_.InvokeVoid(
		t,
		"resetFirewallDomainRedirectionAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) ResetQType() {
	_jsii_.InvokeVoid(
		t,
		"resetQType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfResolverFirewallRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResolverFirewallRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResolverFirewallRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResolverFirewallRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResolverFirewallRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResolverFirewallRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfResolverFirewallRule) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

