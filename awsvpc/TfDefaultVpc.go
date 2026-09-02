package awsvpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsvpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_vpc aws_default_vpc}.
// Experimental.
type TfDefaultVpc interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	AssignGeneratedIpv6CidrBlock() interface{}
	// Experimental.
	SetAssignGeneratedIpv6CidrBlock(val interface{})
	// Experimental.
	AssignGeneratedIpv6CidrBlockInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CidrBlock() *string
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
	DefaultNetworkAclId() *string
	// Experimental.
	DefaultRouteTableId() *string
	// Experimental.
	DefaultSecurityGroupId() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DhcpOptionsId() *string
	// Experimental.
	EnableDnsHostnames() interface{}
	// Experimental.
	SetEnableDnsHostnames(val interface{})
	// Experimental.
	EnableDnsHostnamesInput() interface{}
	// Experimental.
	EnableDnsSupport() interface{}
	// Experimental.
	SetEnableDnsSupport(val interface{})
	// Experimental.
	EnableDnsSupportInput() interface{}
	// Experimental.
	EnableNetworkAddressUsageMetrics() interface{}
	// Experimental.
	SetEnableNetworkAddressUsageMetrics(val interface{})
	// Experimental.
	EnableNetworkAddressUsageMetricsInput() interface{}
	// Experimental.
	ExistingDefaultVpc() cdktn.IResolvable
	// Experimental.
	ForceDestroy() interface{}
	// Experimental.
	SetForceDestroy(val interface{})
	// Experimental.
	ForceDestroyInput() interface{}
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
	InstanceTenancy() *string
	// Experimental.
	Ipv6AssociationId() *string
	// Experimental.
	Ipv6CidrBlock() *string
	// Experimental.
	SetIpv6CidrBlock(val *string)
	// Experimental.
	Ipv6CidrBlockInput() *string
	// Experimental.
	Ipv6CidrBlockNetworkBorderGroup() *string
	// Experimental.
	SetIpv6CidrBlockNetworkBorderGroup(val *string)
	// Experimental.
	Ipv6CidrBlockNetworkBorderGroupInput() *string
	// Experimental.
	Ipv6IpamPoolId() *string
	// Experimental.
	SetIpv6IpamPoolId(val *string)
	// Experimental.
	Ipv6IpamPoolIdInput() *string
	// Experimental.
	Ipv6NetmaskLength() *float64
	// Experimental.
	SetIpv6NetmaskLength(val *float64)
	// Experimental.
	Ipv6NetmaskLengthInput() *float64
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MainRouteTableId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OwnerId() *string
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
	ResetAssignGeneratedIpv6CidrBlock()
	// Experimental.
	ResetEnableDnsHostnames()
	// Experimental.
	ResetEnableDnsSupport()
	// Experimental.
	ResetEnableNetworkAddressUsageMetrics()
	// Experimental.
	ResetForceDestroy()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIpv6CidrBlock()
	// Experimental.
	ResetIpv6CidrBlockNetworkBorderGroup()
	// Experimental.
	ResetIpv6IpamPoolId()
	// Experimental.
	ResetIpv6NetmaskLength()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
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

// The jsii proxy struct for TfDefaultVpc
type jsiiProxy_TfDefaultVpc struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfDefaultVpc) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) AssignGeneratedIpv6CidrBlock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignGeneratedIpv6CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) AssignGeneratedIpv6CidrBlockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignGeneratedIpv6CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) DefaultNetworkAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNetworkAclId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) DefaultRouteTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRouteTableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) DefaultSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) DhcpOptionsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhcpOptionsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) EnableDnsHostnames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDnsHostnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) EnableDnsHostnamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDnsHostnamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) EnableDnsSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDnsSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) EnableDnsSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDnsSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) EnableNetworkAddressUsageMetrics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkAddressUsageMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) EnableNetworkAddressUsageMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkAddressUsageMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) ExistingDefaultVpc() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"existingDefaultVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) InstanceTenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Ipv6AssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Ipv6CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Ipv6CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Ipv6CidrBlockNetworkBorderGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockNetworkBorderGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Ipv6CidrBlockNetworkBorderGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockNetworkBorderGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Ipv6IpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6IpamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Ipv6IpamPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6IpamPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Ipv6NetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6NetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Ipv6NetmaskLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6NetmaskLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) MainRouteTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainRouteTableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDefaultVpc) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_vpc aws_default_vpc} Resource.
// Experimental.
func NewTfDefaultVpc(scope constructs.Construct, id *string, config *TfDefaultVpcConfig) TfDefaultVpc {
	_init_.Initialize()

	if err := validateNewTfDefaultVpcParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDefaultVpc{}

	_jsii_.Create(
		"@cdktn/aws-vpc.TfDefaultVpc",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_vpc aws_default_vpc} Resource.
// Experimental.
func NewTfDefaultVpc_Override(t TfDefaultVpc, scope constructs.Construct, id *string, config *TfDefaultVpcConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.TfDefaultVpc",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetAssignGeneratedIpv6CidrBlock(val interface{}) {
	if err := j.validateSetAssignGeneratedIpv6CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assignGeneratedIpv6CidrBlock",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetEnableDnsHostnames(val interface{}) {
	if err := j.validateSetEnableDnsHostnamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDnsHostnames",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetEnableDnsSupport(val interface{}) {
	if err := j.validateSetEnableDnsSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDnsSupport",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetEnableNetworkAddressUsageMetrics(val interface{}) {
	if err := j.validateSetEnableNetworkAddressUsageMetricsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNetworkAddressUsageMetrics",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetIpv6CidrBlock(val *string) {
	if err := j.validateSetIpv6CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6CidrBlock",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetIpv6CidrBlockNetworkBorderGroup(val *string) {
	if err := j.validateSetIpv6CidrBlockNetworkBorderGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6CidrBlockNetworkBorderGroup",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetIpv6IpamPoolId(val *string) {
	if err := j.validateSetIpv6IpamPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6IpamPoolId",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetIpv6NetmaskLength(val *float64) {
	if err := j.validateSetIpv6NetmaskLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6NetmaskLength",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfDefaultVpc)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a TfDefaultVpc resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfDefaultVpc_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfDefaultVpc_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc.TfDefaultVpc",
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
func TfDefaultVpc_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDefaultVpc_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc.TfDefaultVpc",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfDefaultVpc_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDefaultVpc_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc.TfDefaultVpc",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfDefaultVpc_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDefaultVpc_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc.TfDefaultVpc",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfDefaultVpc_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-vpc.TfDefaultVpc",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfDefaultVpc) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfDefaultVpc) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfDefaultVpc) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDefaultVpc) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDefaultVpc) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDefaultVpc) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDefaultVpc) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDefaultVpc) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDefaultVpc) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDefaultVpc) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDefaultVpc) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDefaultVpc) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDefaultVpc) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfDefaultVpc) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDefaultVpc) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfDefaultVpc) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfDefaultVpc) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfDefaultVpc) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfDefaultVpc) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfDefaultVpc) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfDefaultVpc) ResetAssignGeneratedIpv6CidrBlock() {
	_jsii_.InvokeVoid(
		t,
		"resetAssignGeneratedIpv6CidrBlock",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDefaultVpc) ResetEnableDnsHostnames() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableDnsHostnames",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDefaultVpc) ResetEnableDnsSupport() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableDnsSupport",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDefaultVpc) ResetEnableNetworkAddressUsageMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableNetworkAddressUsageMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDefaultVpc) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		t,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDefaultVpc) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDefaultVpc) ResetIpv6CidrBlock() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv6CidrBlock",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDefaultVpc) ResetIpv6CidrBlockNetworkBorderGroup() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv6CidrBlockNetworkBorderGroup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDefaultVpc) ResetIpv6IpamPoolId() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv6IpamPoolId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDefaultVpc) ResetIpv6NetmaskLength() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv6NetmaskLength",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDefaultVpc) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDefaultVpc) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDefaultVpc) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDefaultVpc) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDefaultVpc) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDefaultVpc) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDefaultVpc) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDefaultVpc) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDefaultVpc) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDefaultVpc) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDefaultVpc) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

