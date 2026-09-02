package awsvpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsvpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_traffic_mirror_filter_rule aws_ec2_traffic_mirror_filter_rule}.
// Experimental.
type TfEc2TrafficMirrorFilterRule interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
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
	DestinationCidrBlock() *string
	// Experimental.
	SetDestinationCidrBlock(val *string)
	// Experimental.
	DestinationCidrBlockInput() *string
	// Experimental.
	DestinationPortRange() TfEc2TrafficMirrorFilterRule_DestinationPortRangePropertyOutputReference
	// Experimental.
	DestinationPortRangeInput() *TfEc2TrafficMirrorFilterRule_DestinationPortRangeProperty
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
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Protocol() *float64
	// Experimental.
	SetProtocol(val *float64)
	// Experimental.
	ProtocolInput() *float64
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
	RuleAction() *string
	// Experimental.
	SetRuleAction(val *string)
	// Experimental.
	RuleActionInput() *string
	// Experimental.
	RuleNumber() *float64
	// Experimental.
	SetRuleNumber(val *float64)
	// Experimental.
	RuleNumberInput() *float64
	// Experimental.
	SourceCidrBlock() *string
	// Experimental.
	SetSourceCidrBlock(val *string)
	// Experimental.
	SourceCidrBlockInput() *string
	// Experimental.
	SourcePortRange() TfEc2TrafficMirrorFilterRule_SourcePortRangePropertyOutputReference
	// Experimental.
	SourcePortRangeInput() *TfEc2TrafficMirrorFilterRule_SourcePortRangeProperty
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	TrafficDirection() *string
	// Experimental.
	SetTrafficDirection(val *string)
	// Experimental.
	TrafficDirectionInput() *string
	// Experimental.
	TrafficMirrorFilterId() *string
	// Experimental.
	SetTrafficMirrorFilterId(val *string)
	// Experimental.
	TrafficMirrorFilterIdInput() *string
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
	PutDestinationPortRange(value *TfEc2TrafficMirrorFilterRule_DestinationPortRangeProperty)
	// Experimental.
	PutSourcePortRange(value *TfEc2TrafficMirrorFilterRule_SourcePortRangeProperty)
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
	ResetDescription()
	// Experimental.
	ResetDestinationPortRange()
	// Experimental.
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetProtocol()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSourcePortRange()
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

// The jsii proxy struct for TfEc2TrafficMirrorFilterRule
type jsiiProxy_TfEc2TrafficMirrorFilterRule struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) DestinationCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) DestinationCidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationCidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) DestinationPortRange() TfEc2TrafficMirrorFilterRule_DestinationPortRangePropertyOutputReference {
	var returns TfEc2TrafficMirrorFilterRule_DestinationPortRangePropertyOutputReference
	_jsii_.Get(
		j,
		"destinationPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) DestinationPortRangeInput() *TfEc2TrafficMirrorFilterRule_DestinationPortRangeProperty {
	var returns *TfEc2TrafficMirrorFilterRule_DestinationPortRangeProperty
	_jsii_.Get(
		j,
		"destinationPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) Protocol() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) ProtocolInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) RuleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) RuleActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) RuleNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) RuleNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) SourceCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) SourceCidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) SourcePortRange() TfEc2TrafficMirrorFilterRule_SourcePortRangePropertyOutputReference {
	var returns TfEc2TrafficMirrorFilterRule_SourcePortRangePropertyOutputReference
	_jsii_.Get(
		j,
		"sourcePortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) SourcePortRangeInput() *TfEc2TrafficMirrorFilterRule_SourcePortRangeProperty {
	var returns *TfEc2TrafficMirrorFilterRule_SourcePortRangeProperty
	_jsii_.Get(
		j,
		"sourcePortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) TrafficDirection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) TrafficDirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) TrafficMirrorFilterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficMirrorFilterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule) TrafficMirrorFilterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficMirrorFilterIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_traffic_mirror_filter_rule aws_ec2_traffic_mirror_filter_rule} Resource.
// Experimental.
func NewTfEc2TrafficMirrorFilterRule(scope constructs.Construct, id *string, config *TfEc2TrafficMirrorFilterRuleConfig) TfEc2TrafficMirrorFilterRule {
	_init_.Initialize()

	if err := validateNewTfEc2TrafficMirrorFilterRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEc2TrafficMirrorFilterRule{}

	_jsii_.Create(
		"@cdktn/aws-vpc.TfEc2TrafficMirrorFilterRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_traffic_mirror_filter_rule aws_ec2_traffic_mirror_filter_rule} Resource.
// Experimental.
func NewTfEc2TrafficMirrorFilterRule_Override(t TfEc2TrafficMirrorFilterRule, scope constructs.Construct, id *string, config *TfEc2TrafficMirrorFilterRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.TfEc2TrafficMirrorFilterRule",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule)SetDestinationCidrBlock(val *string) {
	if err := j.validateSetDestinationCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationCidrBlock",
		val,
	)
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule)SetProtocol(val *float64) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule)SetRuleAction(val *string) {
	if err := j.validateSetRuleActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleAction",
		val,
	)
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule)SetRuleNumber(val *float64) {
	if err := j.validateSetRuleNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleNumber",
		val,
	)
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule)SetSourceCidrBlock(val *string) {
	if err := j.validateSetSourceCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCidrBlock",
		val,
	)
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule)SetTrafficDirection(val *string) {
	if err := j.validateSetTrafficDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficDirection",
		val,
	)
}

func (j *jsiiProxy_TfEc2TrafficMirrorFilterRule)SetTrafficMirrorFilterId(val *string) {
	if err := j.validateSetTrafficMirrorFilterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficMirrorFilterId",
		val,
	)
}

// Generates CDKTN code for importing a TfEc2TrafficMirrorFilterRule resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfEc2TrafficMirrorFilterRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfEc2TrafficMirrorFilterRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc.TfEc2TrafficMirrorFilterRule",
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
func TfEc2TrafficMirrorFilterRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfEc2TrafficMirrorFilterRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc.TfEc2TrafficMirrorFilterRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfEc2TrafficMirrorFilterRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfEc2TrafficMirrorFilterRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc.TfEc2TrafficMirrorFilterRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfEc2TrafficMirrorFilterRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfEc2TrafficMirrorFilterRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc.TfEc2TrafficMirrorFilterRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfEc2TrafficMirrorFilterRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-vpc.TfEc2TrafficMirrorFilterRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) PutDestinationPortRange(value *TfEc2TrafficMirrorFilterRule_DestinationPortRangeProperty) {
	if err := t.validatePutDestinationPortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDestinationPortRange",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) PutSourcePortRange(value *TfEc2TrafficMirrorFilterRule_SourcePortRangeProperty) {
	if err := t.validatePutSourcePortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSourcePortRange",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) ResetDestinationPortRange() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationPortRange",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) ResetProtocol() {
	_jsii_.InvokeVoid(
		t,
		"resetProtocol",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) ResetSourcePortRange() {
	_jsii_.InvokeVoid(
		t,
		"resetSourcePortRange",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEc2TrafficMirrorFilterRule) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

