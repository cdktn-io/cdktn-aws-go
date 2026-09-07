package transitgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/transitgateway/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/transitgateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry aws_ec2_transit_gateway_metering_policy_entry}.
// Experimental.
type AwsMeteringPolicyEntry interface {
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
	Timeouts() AwsMeteringPolicyEntry_TimeoutsPropertyOutputReference
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
	PutTimeouts(value *AwsMeteringPolicyEntry_TimeoutsProperty)
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

// The jsii proxy struct for AwsMeteringPolicyEntry
type jsiiProxy_AwsMeteringPolicyEntry struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) DestinationCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) DestinationCidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationCidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) DestinationPortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) DestinationPortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) DestinationTransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTransitGatewayAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) DestinationTransitGatewayAttachmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTransitGatewayAttachmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) DestinationTransitGatewayAttachmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTransitGatewayAttachmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) DestinationTransitGatewayAttachmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTransitGatewayAttachmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) MeteredAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meteredAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) MeteredAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meteredAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) PolicyRuleNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyRuleNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) PolicyRuleNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyRuleNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) SourceCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) SourceCidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) SourcePortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) SourcePortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) SourceTransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTransitGatewayAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) SourceTransitGatewayAttachmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTransitGatewayAttachmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) SourceTransitGatewayAttachmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTransitGatewayAttachmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) SourceTransitGatewayAttachmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTransitGatewayAttachmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) Timeouts() AwsMeteringPolicyEntry_TimeoutsPropertyOutputReference {
	var returns AwsMeteringPolicyEntry_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) TransitGatewayMeteringPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayMeteringPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMeteringPolicyEntry) TransitGatewayMeteringPolicyIdInput() *string {
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
func NewAwsMeteringPolicyEntry(scope constructs.Construct, id *string, config *AwsMeteringPolicyEntryConfig) AwsMeteringPolicyEntry {
	_init_.Initialize()

	if err := validateNewAwsMeteringPolicyEntryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMeteringPolicyEntry{}

	_jsii_.Create(
		"@cdktn/aws-transit-gateway.AwsMeteringPolicyEntry",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway_metering_policy_entry aws_ec2_transit_gateway_metering_policy_entry} Resource.
// Experimental.
func NewAwsMeteringPolicyEntry_Override(a AwsMeteringPolicyEntry, scope constructs.Construct, id *string, config *AwsMeteringPolicyEntryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transit-gateway.AwsMeteringPolicyEntry",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetDestinationCidrBlock(val *string) {
	if err := j.validateSetDestinationCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationCidrBlock",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetDestinationPortRange(val *string) {
	if err := j.validateSetDestinationPortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPortRange",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetDestinationTransitGatewayAttachmentId(val *string) {
	if err := j.validateSetDestinationTransitGatewayAttachmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationTransitGatewayAttachmentId",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetDestinationTransitGatewayAttachmentType(val *string) {
	if err := j.validateSetDestinationTransitGatewayAttachmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationTransitGatewayAttachmentType",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetMeteredAccount(val *string) {
	if err := j.validateSetMeteredAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"meteredAccount",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetPolicyRuleNumber(val *float64) {
	if err := j.validateSetPolicyRuleNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyRuleNumber",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetSourceCidrBlock(val *string) {
	if err := j.validateSetSourceCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCidrBlock",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetSourcePortRange(val *string) {
	if err := j.validateSetSourcePortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePortRange",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetSourceTransitGatewayAttachmentId(val *string) {
	if err := j.validateSetSourceTransitGatewayAttachmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTransitGatewayAttachmentId",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetSourceTransitGatewayAttachmentType(val *string) {
	if err := j.validateSetSourceTransitGatewayAttachmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTransitGatewayAttachmentType",
		val,
	)
}

func (j *jsiiProxy_AwsMeteringPolicyEntry)SetTransitGatewayMeteringPolicyId(val *string) {
	if err := j.validateSetTransitGatewayMeteringPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitGatewayMeteringPolicyId",
		val,
	)
}

// Generates CDKTN code for importing a AwsMeteringPolicyEntry resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsMeteringPolicyEntry_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsMeteringPolicyEntry_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-transit-gateway.AwsMeteringPolicyEntry",
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
func AwsMeteringPolicyEntry_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsMeteringPolicyEntry_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-transit-gateway.AwsMeteringPolicyEntry",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsMeteringPolicyEntry_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsMeteringPolicyEntry_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-transit-gateway.AwsMeteringPolicyEntry",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsMeteringPolicyEntry_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsMeteringPolicyEntry_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-transit-gateway.AwsMeteringPolicyEntry",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsMeteringPolicyEntry_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-transit-gateway.AwsMeteringPolicyEntry",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMeteringPolicyEntry) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMeteringPolicyEntry) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMeteringPolicyEntry) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMeteringPolicyEntry) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMeteringPolicyEntry) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMeteringPolicyEntry) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMeteringPolicyEntry) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMeteringPolicyEntry) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMeteringPolicyEntry) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMeteringPolicyEntry) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsMeteringPolicyEntry) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) PutTimeouts(value *AwsMeteringPolicyEntry_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) ResetDestinationCidrBlock() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationCidrBlock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) ResetDestinationPortRange() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationPortRange",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) ResetDestinationTransitGatewayAttachmentId() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationTransitGatewayAttachmentId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) ResetDestinationTransitGatewayAttachmentType() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationTransitGatewayAttachmentType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) ResetProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) ResetSourceCidrBlock() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceCidrBlock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) ResetSourcePortRange() {
	_jsii_.InvokeVoid(
		a,
		"resetSourcePortRange",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) ResetSourceTransitGatewayAttachmentId() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceTransitGatewayAttachmentId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) ResetSourceTransitGatewayAttachmentType() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceTransitGatewayAttachmentType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMeteringPolicyEntry) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

