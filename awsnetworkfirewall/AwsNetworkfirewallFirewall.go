package awsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall aws_networkfirewall_firewall}.
// Experimental.
type AwsNetworkfirewallFirewall interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	AvailabilityZoneChangeProtection() interface{}
	// Experimental.
	SetAvailabilityZoneChangeProtection(val interface{})
	// Experimental.
	AvailabilityZoneChangeProtectionInput() interface{}
	// Experimental.
	AvailabilityZoneMapping() AwsNetworkfirewallFirewall_AvailabilityZoneMappingPropertyList
	// Experimental.
	AvailabilityZoneMappingInput() interface{}
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
	DeleteProtection() interface{}
	// Experimental.
	SetDeleteProtection(val interface{})
	// Experimental.
	DeleteProtectionInput() interface{}
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
	EnabledAnalysisTypes() *[]*string
	// Experimental.
	SetEnabledAnalysisTypes(val *[]*string)
	// Experimental.
	EnabledAnalysisTypesInput() *[]*string
	// Experimental.
	EncryptionConfiguration() AwsNetworkfirewallFirewall_EncryptionConfigurationPropertyOutputReference
	// Experimental.
	EncryptionConfigurationInput() *AwsNetworkfirewallFirewall_EncryptionConfigurationProperty
	// Experimental.
	FirewallPolicyArn() *string
	// Experimental.
	SetFirewallPolicyArn(val *string)
	// Experimental.
	FirewallPolicyArnInput() *string
	// Experimental.
	FirewallPolicyChangeProtection() interface{}
	// Experimental.
	SetFirewallPolicyChangeProtection(val interface{})
	// Experimental.
	FirewallPolicyChangeProtectionInput() interface{}
	// Experimental.
	FirewallStatus() AwsNetworkfirewallFirewall_FirewallStatusPropertyList
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
	SubnetChangeProtection() interface{}
	// Experimental.
	SetSubnetChangeProtection(val interface{})
	// Experimental.
	SubnetChangeProtectionInput() interface{}
	// Experimental.
	SubnetMapping() AwsNetworkfirewallFirewall_SubnetMappingPropertyList
	// Experimental.
	SubnetMappingInput() interface{}
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
	Timeouts() AwsNetworkfirewallFirewall_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TransitGatewayId() *string
	// Experimental.
	SetTransitGatewayId(val *string)
	// Experimental.
	TransitGatewayIdInput() *string
	// Experimental.
	TransitGatewayOwnerAccountId() *string
	// Experimental.
	UpdateToken() *string
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
	PutAvailabilityZoneMapping(value interface{})
	// Experimental.
	PutEncryptionConfiguration(value *AwsNetworkfirewallFirewall_EncryptionConfigurationProperty)
	// Experimental.
	PutSubnetMapping(value interface{})
	// Experimental.
	PutTimeouts(value *AwsNetworkfirewallFirewall_TimeoutsProperty)
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
	ResetAvailabilityZoneChangeProtection()
	// Experimental.
	ResetAvailabilityZoneMapping()
	// Experimental.
	ResetDeleteProtection()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetEnabledAnalysisTypes()
	// Experimental.
	ResetEncryptionConfiguration()
	// Experimental.
	ResetFirewallPolicyChangeProtection()
	// Experimental.
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSubnetChangeProtection()
	// Experimental.
	ResetSubnetMapping()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTransitGatewayId()
	// Experimental.
	ResetVpcId()
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

// The jsii proxy struct for AwsNetworkfirewallFirewall
type jsiiProxy_AwsNetworkfirewallFirewall struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) AvailabilityZoneChangeProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availabilityZoneChangeProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) AvailabilityZoneChangeProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availabilityZoneChangeProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) AvailabilityZoneMapping() AwsNetworkfirewallFirewall_AvailabilityZoneMappingPropertyList {
	var returns AwsNetworkfirewallFirewall_AvailabilityZoneMappingPropertyList
	_jsii_.Get(
		j,
		"availabilityZoneMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) AvailabilityZoneMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availabilityZoneMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) DeleteProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) DeleteProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) EnabledAnalysisTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledAnalysisTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) EnabledAnalysisTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledAnalysisTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) EncryptionConfiguration() AwsNetworkfirewallFirewall_EncryptionConfigurationPropertyOutputReference {
	var returns AwsNetworkfirewallFirewall_EncryptionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) EncryptionConfigurationInput() *AwsNetworkfirewallFirewall_EncryptionConfigurationProperty {
	var returns *AwsNetworkfirewallFirewall_EncryptionConfigurationProperty
	_jsii_.Get(
		j,
		"encryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) FirewallPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallPolicyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) FirewallPolicyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallPolicyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) FirewallPolicyChangeProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewallPolicyChangeProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) FirewallPolicyChangeProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewallPolicyChangeProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) FirewallStatus() AwsNetworkfirewallFirewall_FirewallStatusPropertyList {
	var returns AwsNetworkfirewallFirewall_FirewallStatusPropertyList
	_jsii_.Get(
		j,
		"firewallStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) SubnetChangeProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetChangeProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) SubnetChangeProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetChangeProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) SubnetMapping() AwsNetworkfirewallFirewall_SubnetMappingPropertyList {
	var returns AwsNetworkfirewallFirewall_SubnetMappingPropertyList
	_jsii_.Get(
		j,
		"subnetMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) SubnetMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) Timeouts() AwsNetworkfirewallFirewall_TimeoutsPropertyOutputReference {
	var returns AwsNetworkfirewallFirewall_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) TransitGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) TransitGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) TransitGatewayOwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayOwnerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) UpdateToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall aws_networkfirewall_firewall} Resource.
// Experimental.
func NewAwsNetworkfirewallFirewall(scope constructs.Construct, id *string, config *AwsNetworkfirewallFirewallConfig) AwsNetworkfirewallFirewall {
	_init_.Initialize()

	if err := validateNewAwsNetworkfirewallFirewallParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsNetworkfirewallFirewall{}

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsNetworkfirewallFirewall",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_firewall aws_networkfirewall_firewall} Resource.
// Experimental.
func NewAwsNetworkfirewallFirewall_Override(a AwsNetworkfirewallFirewall, scope constructs.Construct, id *string, config *AwsNetworkfirewallFirewallConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsNetworkfirewallFirewall",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetAvailabilityZoneChangeProtection(val interface{}) {
	if err := j.validateSetAvailabilityZoneChangeProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneChangeProtection",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetDeleteProtection(val interface{}) {
	if err := j.validateSetDeleteProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteProtection",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetEnabledAnalysisTypes(val *[]*string) {
	if err := j.validateSetEnabledAnalysisTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledAnalysisTypes",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetFirewallPolicyArn(val *string) {
	if err := j.validateSetFirewallPolicyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallPolicyArn",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetFirewallPolicyChangeProtection(val interface{}) {
	if err := j.validateSetFirewallPolicyChangeProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallPolicyChangeProtection",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetSubnetChangeProtection(val interface{}) {
	if err := j.validateSetSubnetChangeProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetChangeProtection",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetTransitGatewayId(val *string) {
	if err := j.validateSetTransitGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitGatewayId",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewall)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTN code for importing a AwsNetworkfirewallFirewall resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsNetworkfirewallFirewall_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsNetworkfirewallFirewall_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-network-firewall.AwsNetworkfirewallFirewall",
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
func AwsNetworkfirewallFirewall_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsNetworkfirewallFirewall_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-network-firewall.AwsNetworkfirewallFirewall",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsNetworkfirewallFirewall_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsNetworkfirewallFirewall_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-network-firewall.AwsNetworkfirewallFirewall",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsNetworkfirewallFirewall_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsNetworkfirewallFirewall_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-network-firewall.AwsNetworkfirewallFirewall",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsNetworkfirewallFirewall_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-network-firewall.AwsNetworkfirewallFirewall",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewall) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewall) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewall) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewall) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewall) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewall) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewall) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewall) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewall) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewall) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewall) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) PutAvailabilityZoneMapping(value interface{}) {
	if err := a.validatePutAvailabilityZoneMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAvailabilityZoneMapping",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) PutEncryptionConfiguration(value *AwsNetworkfirewallFirewall_EncryptionConfigurationProperty) {
	if err := a.validatePutEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) PutSubnetMapping(value interface{}) {
	if err := a.validatePutSubnetMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSubnetMapping",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) PutTimeouts(value *AwsNetworkfirewallFirewall_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ResetAvailabilityZoneChangeProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZoneChangeProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ResetAvailabilityZoneMapping() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZoneMapping",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ResetDeleteProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ResetEnabledAnalysisTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabledAnalysisTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ResetEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ResetFirewallPolicyChangeProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetFirewallPolicyChangeProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ResetSubnetChangeProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetChangeProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ResetSubnetMapping() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetMapping",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ResetTransitGatewayId() {
	_jsii_.InvokeVoid(
		a,
		"resetTransitGatewayId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ResetVpcId() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkfirewallFirewall) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

