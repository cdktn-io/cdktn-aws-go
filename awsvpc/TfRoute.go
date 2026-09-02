package awsvpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsvpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route aws_route}.
// Experimental.
type TfRoute interface {
	cdktn.TerraformResource
	// Experimental.
	CarrierGatewayId() *string
	// Experimental.
	SetCarrierGatewayId(val *string)
	// Experimental.
	CarrierGatewayIdInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	CoreNetworkArn() *string
	// Experimental.
	SetCoreNetworkArn(val *string)
	// Experimental.
	CoreNetworkArnInput() *string
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
	DestinationIpv6CidrBlock() *string
	// Experimental.
	SetDestinationIpv6CidrBlock(val *string)
	// Experimental.
	DestinationIpv6CidrBlockInput() *string
	// Experimental.
	DestinationPrefixListId() *string
	// Experimental.
	SetDestinationPrefixListId(val *string)
	// Experimental.
	DestinationPrefixListIdInput() *string
	// Experimental.
	EgressOnlyGatewayId() *string
	// Experimental.
	SetEgressOnlyGatewayId(val *string)
	// Experimental.
	EgressOnlyGatewayIdInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	GatewayId() *string
	// Experimental.
	SetGatewayId(val *string)
	// Experimental.
	GatewayIdInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	InstanceId() *string
	// Experimental.
	InstanceOwnerId() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LocalGatewayId() *string
	// Experimental.
	SetLocalGatewayId(val *string)
	// Experimental.
	LocalGatewayIdInput() *string
	// Experimental.
	NatGatewayId() *string
	// Experimental.
	SetNatGatewayId(val *string)
	// Experimental.
	NatGatewayIdInput() *string
	// Experimental.
	NetworkInterfaceId() *string
	// Experimental.
	SetNetworkInterfaceId(val *string)
	// Experimental.
	NetworkInterfaceIdInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OdbNetworkArn() *string
	// Experimental.
	SetOdbNetworkArn(val *string)
	// Experimental.
	OdbNetworkArnInput() *string
	// Experimental.
	Origin() *string
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
	RouteTableId() *string
	// Experimental.
	SetRouteTableId(val *string)
	// Experimental.
	RouteTableIdInput() *string
	// Experimental.
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() TfRoute_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TransitGatewayId() *string
	// Experimental.
	SetTransitGatewayId(val *string)
	// Experimental.
	TransitGatewayIdInput() *string
	// Experimental.
	VpcEndpointId() *string
	// Experimental.
	SetVpcEndpointId(val *string)
	// Experimental.
	VpcEndpointIdInput() *string
	// Experimental.
	VpcPeeringConnectionId() *string
	// Experimental.
	SetVpcPeeringConnectionId(val *string)
	// Experimental.
	VpcPeeringConnectionIdInput() *string
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
	PutTimeouts(value *TfRoute_TimeoutsProperty)
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
	ResetCarrierGatewayId()
	// Experimental.
	ResetCoreNetworkArn()
	// Experimental.
	ResetDestinationCidrBlock()
	// Experimental.
	ResetDestinationIpv6CidrBlock()
	// Experimental.
	ResetDestinationPrefixListId()
	// Experimental.
	ResetEgressOnlyGatewayId()
	// Experimental.
	ResetGatewayId()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLocalGatewayId()
	// Experimental.
	ResetNatGatewayId()
	// Experimental.
	ResetNetworkInterfaceId()
	// Experimental.
	ResetOdbNetworkArn()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTransitGatewayId()
	// Experimental.
	ResetVpcEndpointId()
	// Experimental.
	ResetVpcPeeringConnectionId()
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

// The jsii proxy struct for TfRoute
type jsiiProxy_TfRoute struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfRoute) CarrierGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"carrierGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) CarrierGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"carrierGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) CoreNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) CoreNetworkArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) DestinationCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) DestinationCidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationCidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) DestinationIpv6CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationIpv6CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) DestinationIpv6CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationIpv6CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) DestinationPrefixListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPrefixListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) DestinationPrefixListIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPrefixListIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) EgressOnlyGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egressOnlyGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) EgressOnlyGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egressOnlyGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) GatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) GatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) InstanceOwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) LocalGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) LocalGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) NatGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) NatGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) NetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) NetworkInterfaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) OdbNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) OdbNetworkArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) Origin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) RouteTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeTableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) RouteTableIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeTableIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) Timeouts() TfRoute_TimeoutsPropertyOutputReference {
	var returns TfRoute_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) TransitGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) TransitGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) VpcEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) VpcEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) VpcPeeringConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcPeeringConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute) VpcPeeringConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcPeeringConnectionIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route aws_route} Resource.
// Experimental.
func NewTfRoute(scope constructs.Construct, id *string, config *TfRouteConfig) TfRoute {
	_init_.Initialize()

	if err := validateNewTfRouteParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRoute{}

	_jsii_.Create(
		"@cdktn/aws-vpc.TfRoute",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route aws_route} Resource.
// Experimental.
func NewTfRoute_Override(t TfRoute, scope constructs.Construct, id *string, config *TfRouteConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.TfRoute",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfRoute)SetCarrierGatewayId(val *string) {
	if err := j.validateSetCarrierGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"carrierGatewayId",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetCoreNetworkArn(val *string) {
	if err := j.validateSetCoreNetworkArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreNetworkArn",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetDestinationCidrBlock(val *string) {
	if err := j.validateSetDestinationCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationCidrBlock",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetDestinationIpv6CidrBlock(val *string) {
	if err := j.validateSetDestinationIpv6CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationIpv6CidrBlock",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetDestinationPrefixListId(val *string) {
	if err := j.validateSetDestinationPrefixListIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPrefixListId",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetEgressOnlyGatewayId(val *string) {
	if err := j.validateSetEgressOnlyGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressOnlyGatewayId",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetGatewayId(val *string) {
	if err := j.validateSetGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayId",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetLocalGatewayId(val *string) {
	if err := j.validateSetLocalGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localGatewayId",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetNatGatewayId(val *string) {
	if err := j.validateSetNatGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"natGatewayId",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetNetworkInterfaceId(val *string) {
	if err := j.validateSetNetworkInterfaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaceId",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetOdbNetworkArn(val *string) {
	if err := j.validateSetOdbNetworkArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"odbNetworkArn",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetRouteTableId(val *string) {
	if err := j.validateSetRouteTableIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeTableId",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetTransitGatewayId(val *string) {
	if err := j.validateSetTransitGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitGatewayId",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetVpcEndpointId(val *string) {
	if err := j.validateSetVpcEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcEndpointId",
		val,
	)
}

func (j *jsiiProxy_TfRoute)SetVpcPeeringConnectionId(val *string) {
	if err := j.validateSetVpcPeeringConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcPeeringConnectionId",
		val,
	)
}

// Generates CDKTN code for importing a TfRoute resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfRoute_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfRoute_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc.TfRoute",
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
func TfRoute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfRoute_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc.TfRoute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfRoute_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfRoute_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc.TfRoute",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfRoute_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfRoute_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc.TfRoute",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfRoute_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-vpc.TfRoute",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfRoute) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfRoute) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfRoute) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRoute) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRoute) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRoute) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRoute) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRoute) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRoute) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRoute) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRoute) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRoute) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRoute) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfRoute) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRoute) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfRoute) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfRoute) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfRoute) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfRoute) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfRoute) PutTimeouts(value *TfRoute_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRoute) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfRoute) ResetCarrierGatewayId() {
	_jsii_.InvokeVoid(
		t,
		"resetCarrierGatewayId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) ResetCoreNetworkArn() {
	_jsii_.InvokeVoid(
		t,
		"resetCoreNetworkArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) ResetDestinationCidrBlock() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationCidrBlock",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) ResetDestinationIpv6CidrBlock() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationIpv6CidrBlock",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) ResetDestinationPrefixListId() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationPrefixListId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) ResetEgressOnlyGatewayId() {
	_jsii_.InvokeVoid(
		t,
		"resetEgressOnlyGatewayId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) ResetGatewayId() {
	_jsii_.InvokeVoid(
		t,
		"resetGatewayId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) ResetLocalGatewayId() {
	_jsii_.InvokeVoid(
		t,
		"resetLocalGatewayId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) ResetNatGatewayId() {
	_jsii_.InvokeVoid(
		t,
		"resetNatGatewayId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) ResetNetworkInterfaceId() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkInterfaceId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) ResetOdbNetworkArn() {
	_jsii_.InvokeVoid(
		t,
		"resetOdbNetworkArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) ResetTransitGatewayId() {
	_jsii_.InvokeVoid(
		t,
		"resetTransitGatewayId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) ResetVpcEndpointId() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcEndpointId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) ResetVpcPeeringConnectionId() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcPeeringConnectionId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRoute) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRoute) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRoute) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRoute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRoute) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRoute) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

