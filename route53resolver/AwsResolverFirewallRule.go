package route53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/route53resolver/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/route53resolver/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule aws_route53_resolver_firewall_rule}.
// Experimental.
type AwsResolverFirewallRule interface {
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

// The jsii proxy struct for AwsResolverFirewallRule
type jsiiProxy_AwsResolverFirewallRule struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsResolverFirewallRule) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) BlockOverrideDnsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDnsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) BlockOverrideDnsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDnsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) BlockOverrideDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) BlockOverrideDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) BlockOverrideTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockOverrideTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) BlockOverrideTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockOverrideTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) BlockResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) BlockResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) ConfidenceThreshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidenceThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) ConfidenceThresholdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidenceThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) DnsThreatProtection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsThreatProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) DnsThreatProtectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsThreatProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) FirewallDomainListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) FirewallDomainListIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) FirewallDomainRedirectionAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainRedirectionAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) FirewallDomainRedirectionActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainRedirectionActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) FirewallRuleGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) FirewallRuleGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) FirewallThreatProtectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallThreatProtectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) QType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) QTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResolverFirewallRule) TerraformResourceType() *string {
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
func NewAwsResolverFirewallRule(scope constructs.Construct, id *string, config *AwsResolverFirewallRuleConfig) AwsResolverFirewallRule {
	_init_.Initialize()

	if err := validateNewAwsResolverFirewallRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsResolverFirewallRule{}

	_jsii_.Create(
		"@cdktn/aws-route-53-resolver.AwsResolverFirewallRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_firewall_rule aws_route53_resolver_firewall_rule} Resource.
// Experimental.
func NewAwsResolverFirewallRule_Override(a AwsResolverFirewallRule, scope constructs.Construct, id *string, config *AwsResolverFirewallRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53-resolver.AwsResolverFirewallRule",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetBlockOverrideDnsType(val *string) {
	if err := j.validateSetBlockOverrideDnsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockOverrideDnsType",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetBlockOverrideDomain(val *string) {
	if err := j.validateSetBlockOverrideDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockOverrideDomain",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetBlockOverrideTtl(val *float64) {
	if err := j.validateSetBlockOverrideTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockOverrideTtl",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetBlockResponse(val *string) {
	if err := j.validateSetBlockResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockResponse",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetConfidenceThreshold(val *string) {
	if err := j.validateSetConfidenceThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidenceThreshold",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetDnsThreatProtection(val *string) {
	if err := j.validateSetDnsThreatProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsThreatProtection",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetFirewallDomainListId(val *string) {
	if err := j.validateSetFirewallDomainListIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallDomainListId",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetFirewallDomainRedirectionAction(val *string) {
	if err := j.validateSetFirewallDomainRedirectionActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallDomainRedirectionAction",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetFirewallRuleGroupId(val *string) {
	if err := j.validateSetFirewallRuleGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallRuleGroupId",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetQType(val *string) {
	if err := j.validateSetQTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qType",
		val,
	)
}

func (j *jsiiProxy_AwsResolverFirewallRule)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTN code for importing a AwsResolverFirewallRule resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsResolverFirewallRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsResolverFirewallRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53-resolver.AwsResolverFirewallRule",
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
func AwsResolverFirewallRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsResolverFirewallRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53-resolver.AwsResolverFirewallRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsResolverFirewallRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsResolverFirewallRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53-resolver.AwsResolverFirewallRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsResolverFirewallRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsResolverFirewallRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53-resolver.AwsResolverFirewallRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsResolverFirewallRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-route-53-resolver.AwsResolverFirewallRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsResolverFirewallRule) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsResolverFirewallRule) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResolverFirewallRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsResolverFirewallRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsResolverFirewallRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsResolverFirewallRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsResolverFirewallRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsResolverFirewallRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsResolverFirewallRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsResolverFirewallRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResolverFirewallRule) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResolverFirewallRule) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsResolverFirewallRule) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) ResetBlockOverrideDnsType() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockOverrideDnsType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) ResetBlockOverrideDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockOverrideDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) ResetBlockOverrideTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockOverrideTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) ResetBlockResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetBlockResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) ResetConfidenceThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetConfidenceThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) ResetDnsThreatProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetDnsThreatProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) ResetFirewallDomainListId() {
	_jsii_.InvokeVoid(
		a,
		"resetFirewallDomainListId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) ResetFirewallDomainRedirectionAction() {
	_jsii_.InvokeVoid(
		a,
		"resetFirewallDomainRedirectionAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) ResetQType() {
	_jsii_.InvokeVoid(
		a,
		"resetQType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResolverFirewallRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResolverFirewallRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResolverFirewallRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResolverFirewallRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResolverFirewallRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResolverFirewallRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResolverFirewallRule) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

