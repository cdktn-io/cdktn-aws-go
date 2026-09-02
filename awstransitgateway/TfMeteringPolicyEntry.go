package awstransitgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstransitgateway/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awstransitgateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry aws_ec2_transit_gateway_metering_policy_entry}.
// Experimental.
type TfMeteringPolicyEntry interface {
	cdktn.TerraformResource
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
	DestinationCidrBlock() *string
	// Experimental.
	SetDestinationCidrBlock(val *string)
	// Experimental.
	DestinationCidrBlockInput() *string
	// Experimental.
	DestinationPortRange() *string
	// Experimental.
	SetDestinationPortRange(val *string)
	// Experimental.
	DestinationPortRangeInput() *string
	// Experimental.
	DestinationTransitGatewayAttachmentId() *string
	// Experimental.
	SetDestinationTransitGatewayAttachmentId(val *string)
	// Experimental.
	DestinationTransitGatewayAttachmentIdInput() *string
	// Experimental.
	DestinationTransitGatewayAttachmentType() *string
	// Experimental.
	SetDestinationTransitGatewayAttachmentType(val *string)
	// Experimental.
	DestinationTransitGatewayAttachmentTypeInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MeteredAccount() *string
	// Experimental.
	SetMeteredAccount(val *string)
	// Experimental.
	MeteredAccountInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PolicyRuleNumber() *float64
	// Experimental.
	SetPolicyRuleNumber(val *float64)
	// Experimental.
	PolicyRuleNumberInput() *float64
	// Experimental.
	Protocol() *string
	// Experimental.
	SetProtocol(val *string)
	// Experimental.
	ProtocolInput() *string
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
	SourceCidrBlock() *string
	// Experimental.
	SetSourceCidrBlock(val *string)
	// Experimental.
	SourceCidrBlockInput() *string
	// Experimental.
	SourcePortRange() *string
	// Experimental.
	SetSourcePortRange(val *string)
	// Experimental.
	SourcePortRangeInput() *string
	// Experimental.
	SourceTransitGatewayAttachmentId() *string
	// Experimental.
	SetSourceTransitGatewayAttachmentId(val *string)
	// Experimental.
	SourceTransitGatewayAttachmentIdInput() *string
	// Experimental.
	SourceTransitGatewayAttachmentType() *string
	// Experimental.
	SetSourceTransitGatewayAttachmentType(val *string)
	// Experimental.
	SourceTransitGatewayAttachmentTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() TfMeteringPolicyEntry_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TransitGatewayMeteringPolicyId() *string
	// Experimental.
	SetTransitGatewayMeteringPolicyId(val *string)
	// Experimental.
	TransitGatewayMeteringPolicyIdInput() *string
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
	PutTimeouts(value *TfMeteringPolicyEntry_TimeoutsProperty)
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
	ResetDestinationCidrBlock()
	// Experimental.
	ResetDestinationPortRange()
	// Experimental.
	ResetDestinationTransitGatewayAttachmentId()
	// Experimental.
	ResetDestinationTransitGatewayAttachmentType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetProtocol()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSourceCidrBlock()
	// Experimental.
	ResetSourcePortRange()
	// Experimental.
	ResetSourceTransitGatewayAttachmentId()
	// Experimental.
	ResetSourceTransitGatewayAttachmentType()
	// Experimental.
	ResetTimeouts()
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

// The jsii proxy struct for TfMeteringPolicyEntry
type jsiiProxy_TfMeteringPolicyEntry struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfMeteringPolicyEntry) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) DestinationCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) DestinationCidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationCidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) DestinationPortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) DestinationPortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) DestinationTransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTransitGatewayAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) DestinationTransitGatewayAttachmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTransitGatewayAttachmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) DestinationTransitGatewayAttachmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTransitGatewayAttachmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) DestinationTransitGatewayAttachmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTransitGatewayAttachmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) MeteredAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meteredAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) MeteredAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meteredAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) PolicyRuleNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyRuleNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) PolicyRuleNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyRuleNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) SourceCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) SourceCidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) SourcePortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) SourcePortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) SourceTransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTransitGatewayAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) SourceTransitGatewayAttachmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTransitGatewayAttachmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) SourceTransitGatewayAttachmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTransitGatewayAttachmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) SourceTransitGatewayAttachmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTransitGatewayAttachmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) Timeouts() TfMeteringPolicyEntry_TimeoutsPropertyOutputReference {
	var returns TfMeteringPolicyEntry_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) TransitGatewayMeteringPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayMeteringPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMeteringPolicyEntry) TransitGatewayMeteringPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayMeteringPolicyIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry aws_ec2_transit_gateway_metering_policy_entry} Resource.
// Experimental.
func NewTfMeteringPolicyEntry(scope constructs.Construct, id *string, config *TfMeteringPolicyEntryConfig) TfMeteringPolicyEntry {
	_init_.Initialize()

	if err := validateNewTfMeteringPolicyEntryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMeteringPolicyEntry{}

	_jsii_.Create(
		"@cdktn/aws-transit-gateway.TfMeteringPolicyEntry",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry aws_ec2_transit_gateway_metering_policy_entry} Resource.
// Experimental.
func NewTfMeteringPolicyEntry_Override(t TfMeteringPolicyEntry, scope constructs.Construct, id *string, config *TfMeteringPolicyEntryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transit-gateway.TfMeteringPolicyEntry",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetDestinationCidrBlock(val *string) {
	if err := j.validateSetDestinationCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationCidrBlock",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetDestinationPortRange(val *string) {
	if err := j.validateSetDestinationPortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPortRange",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetDestinationTransitGatewayAttachmentId(val *string) {
	if err := j.validateSetDestinationTransitGatewayAttachmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationTransitGatewayAttachmentId",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetDestinationTransitGatewayAttachmentType(val *string) {
	if err := j.validateSetDestinationTransitGatewayAttachmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationTransitGatewayAttachmentType",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetMeteredAccount(val *string) {
	if err := j.validateSetMeteredAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"meteredAccount",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetPolicyRuleNumber(val *float64) {
	if err := j.validateSetPolicyRuleNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyRuleNumber",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetSourceCidrBlock(val *string) {
	if err := j.validateSetSourceCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCidrBlock",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetSourcePortRange(val *string) {
	if err := j.validateSetSourcePortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePortRange",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetSourceTransitGatewayAttachmentId(val *string) {
	if err := j.validateSetSourceTransitGatewayAttachmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTransitGatewayAttachmentId",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetSourceTransitGatewayAttachmentType(val *string) {
	if err := j.validateSetSourceTransitGatewayAttachmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTransitGatewayAttachmentType",
		val,
	)
}

func (j *jsiiProxy_TfMeteringPolicyEntry)SetTransitGatewayMeteringPolicyId(val *string) {
	if err := j.validateSetTransitGatewayMeteringPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitGatewayMeteringPolicyId",
		val,
	)
}

// Generates CDKTN code for importing a TfMeteringPolicyEntry resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfMeteringPolicyEntry_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfMeteringPolicyEntry_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-transit-gateway.TfMeteringPolicyEntry",
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
func TfMeteringPolicyEntry_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfMeteringPolicyEntry_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-transit-gateway.TfMeteringPolicyEntry",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfMeteringPolicyEntry_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfMeteringPolicyEntry_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-transit-gateway.TfMeteringPolicyEntry",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfMeteringPolicyEntry_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfMeteringPolicyEntry_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-transit-gateway.TfMeteringPolicyEntry",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfMeteringPolicyEntry_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-transit-gateway.TfMeteringPolicyEntry",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfMeteringPolicyEntry) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMeteringPolicyEntry) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMeteringPolicyEntry) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMeteringPolicyEntry) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMeteringPolicyEntry) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMeteringPolicyEntry) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMeteringPolicyEntry) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMeteringPolicyEntry) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMeteringPolicyEntry) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMeteringPolicyEntry) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMeteringPolicyEntry) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMeteringPolicyEntry) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfMeteringPolicyEntry) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) PutTimeouts(value *TfMeteringPolicyEntry_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) ResetDestinationCidrBlock() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationCidrBlock",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) ResetDestinationPortRange() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationPortRange",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) ResetDestinationTransitGatewayAttachmentId() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationTransitGatewayAttachmentId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) ResetDestinationTransitGatewayAttachmentType() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationTransitGatewayAttachmentType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) ResetProtocol() {
	_jsii_.InvokeVoid(
		t,
		"resetProtocol",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) ResetSourceCidrBlock() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceCidrBlock",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) ResetSourcePortRange() {
	_jsii_.InvokeVoid(
		t,
		"resetSourcePortRange",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) ResetSourceTransitGatewayAttachmentId() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceTransitGatewayAttachmentId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) ResetSourceTransitGatewayAttachmentType() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceTransitGatewayAttachmentType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMeteringPolicyEntry) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMeteringPolicyEntry) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMeteringPolicyEntry) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMeteringPolicyEntry) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMeteringPolicyEntry) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMeteringPolicyEntry) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMeteringPolicyEntry) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

