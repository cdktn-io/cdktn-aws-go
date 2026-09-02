package awstransitgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstransitgateway/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awstransitgateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway aws_ec2_transit_gateway}.
// Experimental.
type TfTransitGateway interface {
	cdktn.TerraformResource
	// Experimental.
	AmazonSideAsn() *float64
	// Experimental.
	SetAmazonSideAsn(val *float64)
	// Experimental.
	AmazonSideAsnInput() *float64
	// Experimental.
	Arn() *string
	// Experimental.
	AssociationDefaultRouteTableId() *string
	// Experimental.
	AutoAcceptSharedAttachments() *string
	// Experimental.
	SetAutoAcceptSharedAttachments(val *string)
	// Experimental.
	AutoAcceptSharedAttachmentsInput() *string
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
	DefaultRouteTableAssociation() *string
	// Experimental.
	SetDefaultRouteTableAssociation(val *string)
	// Experimental.
	DefaultRouteTableAssociationInput() *string
	// Experimental.
	DefaultRouteTablePropagation() *string
	// Experimental.
	SetDefaultRouteTablePropagation(val *string)
	// Experimental.
	DefaultRouteTablePropagationInput() *string
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
	DnsSupport() *string
	// Experimental.
	SetDnsSupport(val *string)
	// Experimental.
	DnsSupportInput() *string
	// Experimental.
	EncryptionSupport() *string
	// Experimental.
	SetEncryptionSupport(val *string)
	// Experimental.
	EncryptionSupportInput() *string
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
	MulticastSupport() *string
	// Experimental.
	SetMulticastSupport(val *string)
	// Experimental.
	MulticastSupportInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OwnerId() *string
	// Experimental.
	PropagationDefaultRouteTableId() *string
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
	SecurityGroupReferencingSupport() *string
	// Experimental.
	SetSecurityGroupReferencingSupport(val *string)
	// Experimental.
	SecurityGroupReferencingSupportInput() *string
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
	Timeouts() TfTransitGateway_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TransitGatewayCidrBlocks() *[]*string
	// Experimental.
	SetTransitGatewayCidrBlocks(val *[]*string)
	// Experimental.
	TransitGatewayCidrBlocksInput() *[]*string
	// Experimental.
	VpnEcmpSupport() *string
	// Experimental.
	SetVpnEcmpSupport(val *string)
	// Experimental.
	VpnEcmpSupportInput() *string
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
	PutTimeouts(value *TfTransitGateway_TimeoutsProperty)
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
	ResetAmazonSideAsn()
	// Experimental.
	ResetAutoAcceptSharedAttachments()
	// Experimental.
	ResetDefaultRouteTableAssociation()
	// Experimental.
	ResetDefaultRouteTablePropagation()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetDnsSupport()
	// Experimental.
	ResetEncryptionSupport()
	// Experimental.
	ResetId()
	// Experimental.
	ResetMulticastSupport()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSecurityGroupReferencingSupport()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTransitGatewayCidrBlocks()
	// Experimental.
	ResetVpnEcmpSupport()
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

// The jsii proxy struct for TfTransitGateway
type jsiiProxy_TfTransitGateway struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfTransitGateway) AmazonSideAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"amazonSideAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) AmazonSideAsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"amazonSideAsnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) AssociationDefaultRouteTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationDefaultRouteTableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) AutoAcceptSharedAttachments() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoAcceptSharedAttachments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) AutoAcceptSharedAttachmentsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoAcceptSharedAttachmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) DefaultRouteTableAssociation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRouteTableAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) DefaultRouteTableAssociationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRouteTableAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) DefaultRouteTablePropagation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRouteTablePropagation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) DefaultRouteTablePropagationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRouteTablePropagationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) DnsSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) DnsSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) EncryptionSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) EncryptionSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) MulticastSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multicastSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) MulticastSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multicastSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) PropagationDefaultRouteTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagationDefaultRouteTableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) SecurityGroupReferencingSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupReferencingSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) SecurityGroupReferencingSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupReferencingSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) Timeouts() TfTransitGateway_TimeoutsPropertyOutputReference {
	var returns TfTransitGateway_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) TransitGatewayCidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transitGatewayCidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) TransitGatewayCidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transitGatewayCidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) VpnEcmpSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnEcmpSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransitGateway) VpnEcmpSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnEcmpSupportInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway aws_ec2_transit_gateway} Resource.
// Experimental.
func NewTfTransitGateway(scope constructs.Construct, id *string, config *TfTransitGatewayConfig) TfTransitGateway {
	_init_.Initialize()

	if err := validateNewTfTransitGatewayParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTransitGateway{}

	_jsii_.Create(
		"@cdktn/aws-transit-gateway.TfTransitGateway",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_transit_gateway aws_ec2_transit_gateway} Resource.
// Experimental.
func NewTfTransitGateway_Override(t TfTransitGateway, scope constructs.Construct, id *string, config *TfTransitGatewayConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transit-gateway.TfTransitGateway",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetAmazonSideAsn(val *float64) {
	if err := j.validateSetAmazonSideAsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amazonSideAsn",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetAutoAcceptSharedAttachments(val *string) {
	if err := j.validateSetAutoAcceptSharedAttachmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoAcceptSharedAttachments",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetDefaultRouteTableAssociation(val *string) {
	if err := j.validateSetDefaultRouteTableAssociationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRouteTableAssociation",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetDefaultRouteTablePropagation(val *string) {
	if err := j.validateSetDefaultRouteTablePropagationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRouteTablePropagation",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetDnsSupport(val *string) {
	if err := j.validateSetDnsSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsSupport",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetEncryptionSupport(val *string) {
	if err := j.validateSetEncryptionSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionSupport",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetMulticastSupport(val *string) {
	if err := j.validateSetMulticastSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multicastSupport",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetSecurityGroupReferencingSupport(val *string) {
	if err := j.validateSetSecurityGroupReferencingSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupReferencingSupport",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetTransitGatewayCidrBlocks(val *[]*string) {
	if err := j.validateSetTransitGatewayCidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitGatewayCidrBlocks",
		val,
	)
}

func (j *jsiiProxy_TfTransitGateway)SetVpnEcmpSupport(val *string) {
	if err := j.validateSetVpnEcmpSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnEcmpSupport",
		val,
	)
}

// Generates CDKTN code for importing a TfTransitGateway resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfTransitGateway_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfTransitGateway_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-transit-gateway.TfTransitGateway",
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
func TfTransitGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTransitGateway_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-transit-gateway.TfTransitGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfTransitGateway_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTransitGateway_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-transit-gateway.TfTransitGateway",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfTransitGateway_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTransitGateway_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-transit-gateway.TfTransitGateway",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfTransitGateway_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-transit-gateway.TfTransitGateway",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfTransitGateway) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfTransitGateway) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfTransitGateway) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTransitGateway) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTransitGateway) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTransitGateway) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTransitGateway) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTransitGateway) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTransitGateway) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTransitGateway) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTransitGateway) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTransitGateway) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTransitGateway) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfTransitGateway) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTransitGateway) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfTransitGateway) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfTransitGateway) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfTransitGateway) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfTransitGateway) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfTransitGateway) PutTimeouts(value *TfTransitGateway_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransitGateway) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfTransitGateway) ResetAmazonSideAsn() {
	_jsii_.InvokeVoid(
		t,
		"resetAmazonSideAsn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransitGateway) ResetAutoAcceptSharedAttachments() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoAcceptSharedAttachments",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransitGateway) ResetDefaultRouteTableAssociation() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultRouteTableAssociation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransitGateway) ResetDefaultRouteTablePropagation() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultRouteTablePropagation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransitGateway) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransitGateway) ResetDnsSupport() {
	_jsii_.InvokeVoid(
		t,
		"resetDnsSupport",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransitGateway) ResetEncryptionSupport() {
	_jsii_.InvokeVoid(
		t,
		"resetEncryptionSupport",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransitGateway) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransitGateway) ResetMulticastSupport() {
	_jsii_.InvokeVoid(
		t,
		"resetMulticastSupport",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransitGateway) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransitGateway) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransitGateway) ResetSecurityGroupReferencingSupport() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroupReferencingSupport",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransitGateway) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransitGateway) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransitGateway) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransitGateway) ResetTransitGatewayCidrBlocks() {
	_jsii_.InvokeVoid(
		t,
		"resetTransitGatewayCidrBlocks",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransitGateway) ResetVpnEcmpSupport() {
	_jsii_.InvokeVoid(
		t,
		"resetVpnEcmpSupport",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransitGateway) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTransitGateway) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTransitGateway) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTransitGateway) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTransitGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTransitGateway) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTransitGateway) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

