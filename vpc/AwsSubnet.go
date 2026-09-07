package vpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/vpc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/vpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet aws_subnet}.
// Experimental.
type AwsSubnet interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	AssignIpv6AddressOnCreation() interface{}
	// Experimental.
	SetAssignIpv6AddressOnCreation(val interface{})
	// Experimental.
	AssignIpv6AddressOnCreationInput() interface{}
	// Experimental.
	AvailabilityZone() *string
	// Experimental.
	SetAvailabilityZone(val *string)
	// Experimental.
	AvailabilityZoneId() *string
	// Experimental.
	SetAvailabilityZoneId(val *string)
	// Experimental.
	AvailabilityZoneIdInput() *string
	// Experimental.
	AvailabilityZoneInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CidrBlock() *string
	// Experimental.
	SetCidrBlock(val *string)
	// Experimental.
	CidrBlockInput() *string
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
	CustomerOwnedIpv4Pool() *string
	// Experimental.
	SetCustomerOwnedIpv4Pool(val *string)
	// Experimental.
	CustomerOwnedIpv4PoolInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EnableDns64() interface{}
	// Experimental.
	SetEnableDns64(val interface{})
	// Experimental.
	EnableDns64Input() interface{}
	// Experimental.
	EnableLniAtDeviceIndex() *float64
	// Experimental.
	SetEnableLniAtDeviceIndex(val *float64)
	// Experimental.
	EnableLniAtDeviceIndexInput() *float64
	// Experimental.
	EnableResourceNameDnsAaaaRecordOnLaunch() interface{}
	// Experimental.
	SetEnableResourceNameDnsAaaaRecordOnLaunch(val interface{})
	// Experimental.
	EnableResourceNameDnsAaaaRecordOnLaunchInput() interface{}
	// Experimental.
	EnableResourceNameDnsARecordOnLaunch() interface{}
	// Experimental.
	SetEnableResourceNameDnsARecordOnLaunch(val interface{})
	// Experimental.
	EnableResourceNameDnsARecordOnLaunchInput() interface{}
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
	Ipv4IpamPoolId() *string
	// Experimental.
	SetIpv4IpamPoolId(val *string)
	// Experimental.
	Ipv4IpamPoolIdInput() *string
	// Experimental.
	Ipv4NetmaskLength() *float64
	// Experimental.
	SetIpv4NetmaskLength(val *float64)
	// Experimental.
	Ipv4NetmaskLengthInput() *float64
	// Experimental.
	Ipv6CidrBlock() *string
	// Experimental.
	SetIpv6CidrBlock(val *string)
	// Experimental.
	Ipv6CidrBlockAssociationId() *string
	// Experimental.
	Ipv6CidrBlockInput() *string
	// Experimental.
	Ipv6IpamPoolId() *string
	// Experimental.
	SetIpv6IpamPoolId(val *string)
	// Experimental.
	Ipv6IpamPoolIdInput() *string
	// Experimental.
	Ipv6Native() interface{}
	// Experimental.
	SetIpv6Native(val interface{})
	// Experimental.
	Ipv6NativeInput() interface{}
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
	MapCustomerOwnedIpOnLaunch() interface{}
	// Experimental.
	SetMapCustomerOwnedIpOnLaunch(val interface{})
	// Experimental.
	MapCustomerOwnedIpOnLaunchInput() interface{}
	// Experimental.
	MapPublicIpOnLaunch() interface{}
	// Experimental.
	SetMapPublicIpOnLaunch(val interface{})
	// Experimental.
	MapPublicIpOnLaunchInput() interface{}
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OutpostArn() *string
	// Experimental.
	SetOutpostArn(val *string)
	// Experimental.
	OutpostArnInput() *string
	// Experimental.
	OwnerId() *string
	// Experimental.
	PrivateDnsHostnameTypeOnLaunch() *string
	// Experimental.
	SetPrivateDnsHostnameTypeOnLaunch(val *string)
	// Experimental.
	PrivateDnsHostnameTypeOnLaunchInput() *string
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
	// Experimental.
	Timeouts() AwsSubnet_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	VpcId() *string
	// Experimental.
	SetVpcId(val *string)
	// Experimental.
	VpcIdInput() *string
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
	PutTimeouts(value *AwsSubnet_TimeoutsProperty)
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
	ResetAssignIpv6AddressOnCreation()
	// Experimental.
	ResetAvailabilityZone()
	// Experimental.
	ResetAvailabilityZoneId()
	// Experimental.
	ResetCidrBlock()
	// Experimental.
	ResetCustomerOwnedIpv4Pool()
	// Experimental.
	ResetEnableDns64()
	// Experimental.
	ResetEnableLniAtDeviceIndex()
	// Experimental.
	ResetEnableResourceNameDnsAaaaRecordOnLaunch()
	// Experimental.
	ResetEnableResourceNameDnsARecordOnLaunch()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIpv4IpamPoolId()
	// Experimental.
	ResetIpv4NetmaskLength()
	// Experimental.
	ResetIpv6CidrBlock()
	// Experimental.
	ResetIpv6IpamPoolId()
	// Experimental.
	ResetIpv6Native()
	// Experimental.
	ResetIpv6NetmaskLength()
	// Experimental.
	ResetMapCustomerOwnedIpOnLaunch()
	// Experimental.
	ResetMapPublicIpOnLaunch()
	// Experimental.
	ResetOutpostArn()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPrivateDnsHostnameTypeOnLaunch()
	// Experimental.
	ResetRegion()
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

// The jsii proxy struct for AwsSubnet
type jsiiProxy_AwsSubnet struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsSubnet) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) AssignIpv6AddressOnCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignIpv6AddressOnCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) AssignIpv6AddressOnCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assignIpv6AddressOnCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) AvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) AvailabilityZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) CustomerOwnedIpv4Pool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerOwnedIpv4Pool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) CustomerOwnedIpv4PoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerOwnedIpv4PoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) EnableDns64() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDns64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) EnableDns64Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDns64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) EnableLniAtDeviceIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enableLniAtDeviceIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) EnableLniAtDeviceIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enableLniAtDeviceIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) EnableResourceNameDnsAaaaRecordOnLaunch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsAaaaRecordOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) EnableResourceNameDnsAaaaRecordOnLaunchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsAaaaRecordOnLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) EnableResourceNameDnsARecordOnLaunch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsARecordOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) EnableResourceNameDnsARecordOnLaunchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsARecordOnLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Ipv4IpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4IpamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Ipv4IpamPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4IpamPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Ipv4NetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4NetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Ipv4NetmaskLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv4NetmaskLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Ipv6CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Ipv6CidrBlockAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Ipv6CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Ipv6IpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6IpamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Ipv6IpamPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6IpamPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Ipv6Native() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6Native",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Ipv6NativeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6NativeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Ipv6NetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6NetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Ipv6NetmaskLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ipv6NetmaskLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) MapCustomerOwnedIpOnLaunch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapCustomerOwnedIpOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) MapCustomerOwnedIpOnLaunchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapCustomerOwnedIpOnLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) MapPublicIpOnLaunch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapPublicIpOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) MapPublicIpOnLaunchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mapPublicIpOnLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) OutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) OutpostArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) PrivateDnsHostnameTypeOnLaunch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsHostnameTypeOnLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) PrivateDnsHostnameTypeOnLaunchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsHostnameTypeOnLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) Timeouts() AwsSubnet_TimeoutsPropertyOutputReference {
	var returns AwsSubnet_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSubnet) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet aws_subnet} Resource.
// Experimental.
func NewAwsSubnet(scope constructs.Construct, id *string, config *AwsSubnetConfig) AwsSubnet {
	_init_.Initialize()

	if err := validateNewAwsSubnetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSubnet{}

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsSubnet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet aws_subnet} Resource.
// Experimental.
func NewAwsSubnet_Override(a AwsSubnet, scope constructs.Construct, id *string, config *AwsSubnetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsSubnet",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsSubnet)SetAssignIpv6AddressOnCreation(val interface{}) {
	if err := j.validateSetAssignIpv6AddressOnCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assignIpv6AddressOnCreation",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetAvailabilityZoneId(val *string) {
	if err := j.validateSetAvailabilityZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneId",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetCidrBlock(val *string) {
	if err := j.validateSetCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrBlock",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetCustomerOwnedIpv4Pool(val *string) {
	if err := j.validateSetCustomerOwnedIpv4PoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerOwnedIpv4Pool",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetEnableDns64(val interface{}) {
	if err := j.validateSetEnableDns64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDns64",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetEnableLniAtDeviceIndex(val *float64) {
	if err := j.validateSetEnableLniAtDeviceIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLniAtDeviceIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetEnableResourceNameDnsAaaaRecordOnLaunch(val interface{}) {
	if err := j.validateSetEnableResourceNameDnsAaaaRecordOnLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourceNameDnsAaaaRecordOnLaunch",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetEnableResourceNameDnsARecordOnLaunch(val interface{}) {
	if err := j.validateSetEnableResourceNameDnsARecordOnLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourceNameDnsARecordOnLaunch",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetIpv4IpamPoolId(val *string) {
	if err := j.validateSetIpv4IpamPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4IpamPoolId",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetIpv4NetmaskLength(val *float64) {
	if err := j.validateSetIpv4NetmaskLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4NetmaskLength",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetIpv6CidrBlock(val *string) {
	if err := j.validateSetIpv6CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6CidrBlock",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetIpv6IpamPoolId(val *string) {
	if err := j.validateSetIpv6IpamPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6IpamPoolId",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetIpv6Native(val interface{}) {
	if err := j.validateSetIpv6NativeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Native",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetIpv6NetmaskLength(val *float64) {
	if err := j.validateSetIpv6NetmaskLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6NetmaskLength",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetMapCustomerOwnedIpOnLaunch(val interface{}) {
	if err := j.validateSetMapCustomerOwnedIpOnLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapCustomerOwnedIpOnLaunch",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetMapPublicIpOnLaunch(val interface{}) {
	if err := j.validateSetMapPublicIpOnLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mapPublicIpOnLaunch",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetOutpostArn(val *string) {
	if err := j.validateSetOutpostArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outpostArn",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetPrivateDnsHostnameTypeOnLaunch(val *string) {
	if err := j.validateSetPrivateDnsHostnameTypeOnLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateDnsHostnameTypeOnLaunch",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsSubnet)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTN code for importing a AwsSubnet resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsSubnet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsSubnet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc.AwsSubnet",
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
func AwsSubnet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSubnet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc.AwsSubnet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsSubnet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSubnet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc.AwsSubnet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsSubnet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSubnet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc.AwsSubnet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsSubnet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-vpc.AwsSubnet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsSubnet) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsSubnet) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsSubnet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSubnet) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSubnet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSubnet) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSubnet) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSubnet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSubnet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSubnet) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSubnet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSubnet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSubnet) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsSubnet) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSubnet) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsSubnet) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsSubnet) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsSubnet) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsSubnet) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsSubnet) PutTimeouts(value *AwsSubnet_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSubnet) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsSubnet) ResetAssignIpv6AddressOnCreation() {
	_jsii_.InvokeVoid(
		a,
		"resetAssignIpv6AddressOnCreation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetAvailabilityZoneId() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZoneId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetCidrBlock() {
	_jsii_.InvokeVoid(
		a,
		"resetCidrBlock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetCustomerOwnedIpv4Pool() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerOwnedIpv4Pool",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetEnableDns64() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableDns64",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetEnableLniAtDeviceIndex() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableLniAtDeviceIndex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetEnableResourceNameDnsAaaaRecordOnLaunch() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableResourceNameDnsAaaaRecordOnLaunch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetEnableResourceNameDnsARecordOnLaunch() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableResourceNameDnsARecordOnLaunch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetIpv4IpamPoolId() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv4IpamPoolId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetIpv4NetmaskLength() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv4NetmaskLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetIpv6CidrBlock() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv6CidrBlock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetIpv6IpamPoolId() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv6IpamPoolId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetIpv6Native() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv6Native",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetIpv6NetmaskLength() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv6NetmaskLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetMapCustomerOwnedIpOnLaunch() {
	_jsii_.InvokeVoid(
		a,
		"resetMapCustomerOwnedIpOnLaunch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetMapPublicIpOnLaunch() {
	_jsii_.InvokeVoid(
		a,
		"resetMapPublicIpOnLaunch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetOutpostArn() {
	_jsii_.InvokeVoid(
		a,
		"resetOutpostArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetPrivateDnsHostnameTypeOnLaunch() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateDnsHostnameTypeOnLaunch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSubnet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSubnet) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSubnet) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSubnet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSubnet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSubnet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSubnet) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

