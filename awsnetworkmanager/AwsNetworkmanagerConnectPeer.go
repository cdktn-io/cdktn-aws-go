package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsnetworkmanager/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsnetworkmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_connect_peer aws_networkmanager_connect_peer}.
// Experimental.
type AwsNetworkmanagerConnectPeer interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	BgpOptions() AwsNetworkmanagerConnectPeer_BgpOptionsPropertyOutputReference
	// Experimental.
	BgpOptionsInput() *AwsNetworkmanagerConnectPeer_BgpOptionsProperty
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Configuration() AwsNetworkmanagerConnectPeer_ConfigurationPropertyList
	// Experimental.
	ConnectAttachmentId() *string
	// Experimental.
	SetConnectAttachmentId(val *string)
	// Experimental.
	ConnectAttachmentIdInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConnectPeerId() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	CoreNetworkAddress() *string
	// Experimental.
	SetCoreNetworkAddress(val *string)
	// Experimental.
	CoreNetworkAddressInput() *string
	// Experimental.
	CoreNetworkId() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EdgeLocation() *string
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
	InsideCidrBlocks() *[]*string
	// Experimental.
	SetInsideCidrBlocks(val *[]*string)
	// Experimental.
	InsideCidrBlocksInput() *[]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PeerAddress() *string
	// Experimental.
	SetPeerAddress(val *string)
	// Experimental.
	PeerAddressInput() *string
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
	State() *string
	// Experimental.
	SubnetArn() *string
	// Experimental.
	SetSubnetArn(val *string)
	// Experimental.
	SubnetArnInput() *string
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
	Timeouts() AwsNetworkmanagerConnectPeer_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
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
	PutBgpOptions(value *AwsNetworkmanagerConnectPeer_BgpOptionsProperty)
	// Experimental.
	PutTimeouts(value *AwsNetworkmanagerConnectPeer_TimeoutsProperty)
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
	ResetBgpOptions()
	// Experimental.
	ResetCoreNetworkAddress()
	// Experimental.
	ResetId()
	// Experimental.
	ResetInsideCidrBlocks()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetSubnetArn()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
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

// The jsii proxy struct for AwsNetworkmanagerConnectPeer
type jsiiProxy_AwsNetworkmanagerConnectPeer struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) BgpOptions() AwsNetworkmanagerConnectPeer_BgpOptionsPropertyOutputReference {
	var returns AwsNetworkmanagerConnectPeer_BgpOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"bgpOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) BgpOptionsInput() *AwsNetworkmanagerConnectPeer_BgpOptionsProperty {
	var returns *AwsNetworkmanagerConnectPeer_BgpOptionsProperty
	_jsii_.Get(
		j,
		"bgpOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) Configuration() AwsNetworkmanagerConnectPeer_ConfigurationPropertyList {
	var returns AwsNetworkmanagerConnectPeer_ConfigurationPropertyList
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) ConnectAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) ConnectAttachmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectAttachmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) ConnectPeerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectPeerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) CoreNetworkAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) CoreNetworkAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) CoreNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) EdgeLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) InsideCidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insideCidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) InsideCidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insideCidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) PeerAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) PeerAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) SubnetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) SubnetArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) Timeouts() AwsNetworkmanagerConnectPeer_TimeoutsPropertyOutputReference {
	var returns AwsNetworkmanagerConnectPeer_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_connect_peer aws_networkmanager_connect_peer} Resource.
// Experimental.
func NewAwsNetworkmanagerConnectPeer(scope constructs.Construct, id *string, config *AwsNetworkmanagerConnectPeerConfig) AwsNetworkmanagerConnectPeer {
	_init_.Initialize()

	if err := validateNewAwsNetworkmanagerConnectPeerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsNetworkmanagerConnectPeer{}

	_jsii_.Create(
		"@cdktn/aws-network-manager.AwsNetworkmanagerConnectPeer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_connect_peer aws_networkmanager_connect_peer} Resource.
// Experimental.
func NewAwsNetworkmanagerConnectPeer_Override(a AwsNetworkmanagerConnectPeer, scope constructs.Construct, id *string, config *AwsNetworkmanagerConnectPeerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-manager.AwsNetworkmanagerConnectPeer",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer)SetConnectAttachmentId(val *string) {
	if err := j.validateSetConnectAttachmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectAttachmentId",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer)SetCoreNetworkAddress(val *string) {
	if err := j.validateSetCoreNetworkAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coreNetworkAddress",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer)SetInsideCidrBlocks(val *[]*string) {
	if err := j.validateSetInsideCidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insideCidrBlocks",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer)SetPeerAddress(val *string) {
	if err := j.validateSetPeerAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerAddress",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer)SetSubnetArn(val *string) {
	if err := j.validateSetSubnetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetArn",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkmanagerConnectPeer)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a AwsNetworkmanagerConnectPeer resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsNetworkmanagerConnectPeer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsNetworkmanagerConnectPeer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-network-manager.AwsNetworkmanagerConnectPeer",
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
func AwsNetworkmanagerConnectPeer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsNetworkmanagerConnectPeer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-network-manager.AwsNetworkmanagerConnectPeer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsNetworkmanagerConnectPeer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsNetworkmanagerConnectPeer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-network-manager.AwsNetworkmanagerConnectPeer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsNetworkmanagerConnectPeer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsNetworkmanagerConnectPeer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-network-manager.AwsNetworkmanagerConnectPeer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsNetworkmanagerConnectPeer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-network-manager.AwsNetworkmanagerConnectPeer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) PutBgpOptions(value *AwsNetworkmanagerConnectPeer_BgpOptionsProperty) {
	if err := a.validatePutBgpOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBgpOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) PutTimeouts(value *AwsNetworkmanagerConnectPeer_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) ResetBgpOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetBgpOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) ResetCoreNetworkAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetCoreNetworkAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) ResetInsideCidrBlocks() {
	_jsii_.InvokeVoid(
		a,
		"resetInsideCidrBlocks",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) ResetSubnetArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkmanagerConnectPeer) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

