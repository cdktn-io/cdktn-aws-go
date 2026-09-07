package storagegateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/storagegateway/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/storagegateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway aws_storagegateway_gateway}.
// Experimental.
type AwsGateway interface {
	cdktn.TerraformResource
	// Experimental.
	ActivationKey() *string
	// Experimental.
	SetActivationKey(val *string)
	// Experimental.
	ActivationKeyInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	AverageDownloadRateLimitInBitsPerSec() *float64
	// Experimental.
	SetAverageDownloadRateLimitInBitsPerSec(val *float64)
	// Experimental.
	AverageDownloadRateLimitInBitsPerSecInput() *float64
	// Experimental.
	AverageUploadRateLimitInBitsPerSec() *float64
	// Experimental.
	SetAverageUploadRateLimitInBitsPerSec(val *float64)
	// Experimental.
	AverageUploadRateLimitInBitsPerSecInput() *float64
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CloudwatchLogGroupArn() *string
	// Experimental.
	SetCloudwatchLogGroupArn(val *string)
	// Experimental.
	CloudwatchLogGroupArnInput() *string
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
	Ec2InstanceId() *string
	// Experimental.
	EndpointType() *string
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
	GatewayIpAddress() *string
	// Experimental.
	SetGatewayIpAddress(val *string)
	// Experimental.
	GatewayIpAddressInput() *string
	// Experimental.
	GatewayName() *string
	// Experimental.
	SetGatewayName(val *string)
	// Experimental.
	GatewayNameInput() *string
	// Experimental.
	GatewayNetworkInterface() AwsGateway_GatewayNetworkInterfacePropertyList
	// Experimental.
	GatewayTimezone() *string
	// Experimental.
	SetGatewayTimezone(val *string)
	// Experimental.
	GatewayTimezoneInput() *string
	// Experimental.
	GatewayType() *string
	// Experimental.
	SetGatewayType(val *string)
	// Experimental.
	GatewayTypeInput() *string
	// Experimental.
	GatewayVpcEndpoint() *string
	// Experimental.
	SetGatewayVpcEndpoint(val *string)
	// Experimental.
	GatewayVpcEndpointInput() *string
	// Experimental.
	HostEnvironment() *string
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
	MaintenanceStartTime() AwsGateway_MaintenanceStartTimePropertyOutputReference
	// Experimental.
	MaintenanceStartTimeInput() *AwsGateway_MaintenanceStartTimeProperty
	// Experimental.
	MediumChangerType() *string
	// Experimental.
	SetMediumChangerType(val *string)
	// Experimental.
	MediumChangerTypeInput() *string
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
	SmbActiveDirectorySettings() AwsGateway_SmbActiveDirectorySettingsPropertyOutputReference
	// Experimental.
	SmbActiveDirectorySettingsInput() *AwsGateway_SmbActiveDirectorySettingsProperty
	// Experimental.
	SmbFileShareVisibility() interface{}
	// Experimental.
	SetSmbFileShareVisibility(val interface{})
	// Experimental.
	SmbFileShareVisibilityInput() interface{}
	// Experimental.
	SmbGuestPassword() *string
	// Experimental.
	SetSmbGuestPassword(val *string)
	// Experimental.
	SmbGuestPasswordInput() *string
	// Experimental.
	SmbSecurityStrategy() *string
	// Experimental.
	SetSmbSecurityStrategy(val *string)
	// Experimental.
	SmbSecurityStrategyInput() *string
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
	TapeDriveType() *string
	// Experimental.
	SetTapeDriveType(val *string)
	// Experimental.
	TapeDriveTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsGateway_TimeoutsPropertyOutputReference
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
	PutMaintenanceStartTime(value *AwsGateway_MaintenanceStartTimeProperty)
	// Experimental.
	PutSmbActiveDirectorySettings(value *AwsGateway_SmbActiveDirectorySettingsProperty)
	// Experimental.
	PutTimeouts(value *AwsGateway_TimeoutsProperty)
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
	ResetActivationKey()
	// Experimental.
	ResetAverageDownloadRateLimitInBitsPerSec()
	// Experimental.
	ResetAverageUploadRateLimitInBitsPerSec()
	// Experimental.
	ResetCloudwatchLogGroupArn()
	// Experimental.
	ResetGatewayIpAddress()
	// Experimental.
	ResetGatewayType()
	// Experimental.
	ResetGatewayVpcEndpoint()
	// Experimental.
	ResetId()
	// Experimental.
	ResetMaintenanceStartTime()
	// Experimental.
	ResetMediumChangerType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSmbActiveDirectorySettings()
	// Experimental.
	ResetSmbFileShareVisibility()
	// Experimental.
	ResetSmbGuestPassword()
	// Experimental.
	ResetSmbSecurityStrategy()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTapeDriveType()
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

// The jsii proxy struct for AwsGateway
type jsiiProxy_AwsGateway struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsGateway) ActivationKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) ActivationKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) AverageDownloadRateLimitInBitsPerSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"averageDownloadRateLimitInBitsPerSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) AverageDownloadRateLimitInBitsPerSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"averageDownloadRateLimitInBitsPerSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) AverageUploadRateLimitInBitsPerSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"averageUploadRateLimitInBitsPerSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) AverageUploadRateLimitInBitsPerSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"averageUploadRateLimitInBitsPerSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) CloudwatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) CloudwatchLogGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) Ec2InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) GatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) GatewayIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) GatewayIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) GatewayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) GatewayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) GatewayNetworkInterface() AwsGateway_GatewayNetworkInterfacePropertyList {
	var returns AwsGateway_GatewayNetworkInterfacePropertyList
	_jsii_.Get(
		j,
		"gatewayNetworkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) GatewayTimezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) GatewayTimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) GatewayType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) GatewayTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) GatewayVpcEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayVpcEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) GatewayVpcEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayVpcEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) HostEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) MaintenanceStartTime() AwsGateway_MaintenanceStartTimePropertyOutputReference {
	var returns AwsGateway_MaintenanceStartTimePropertyOutputReference
	_jsii_.Get(
		j,
		"maintenanceStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) MaintenanceStartTimeInput() *AwsGateway_MaintenanceStartTimeProperty {
	var returns *AwsGateway_MaintenanceStartTimeProperty
	_jsii_.Get(
		j,
		"maintenanceStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) MediumChangerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediumChangerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) MediumChangerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediumChangerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) SmbActiveDirectorySettings() AwsGateway_SmbActiveDirectorySettingsPropertyOutputReference {
	var returns AwsGateway_SmbActiveDirectorySettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"smbActiveDirectorySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) SmbActiveDirectorySettingsInput() *AwsGateway_SmbActiveDirectorySettingsProperty {
	var returns *AwsGateway_SmbActiveDirectorySettingsProperty
	_jsii_.Get(
		j,
		"smbActiveDirectorySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) SmbFileShareVisibility() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbFileShareVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) SmbFileShareVisibilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbFileShareVisibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) SmbGuestPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smbGuestPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) SmbGuestPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smbGuestPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) SmbSecurityStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smbSecurityStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) SmbSecurityStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smbSecurityStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) TapeDriveType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tapeDriveType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) TapeDriveTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tapeDriveTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) Timeouts() AwsGateway_TimeoutsPropertyOutputReference {
	var returns AwsGateway_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGateway) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway aws_storagegateway_gateway} Resource.
// Experimental.
func NewAwsGateway(scope constructs.Construct, id *string, config *AwsGatewayConfig) AwsGateway {
	_init_.Initialize()

	if err := validateNewAwsGatewayParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGateway{}

	_jsii_.Create(
		"@cdktn/aws-storage-gateway.AwsGateway",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway aws_storagegateway_gateway} Resource.
// Experimental.
func NewAwsGateway_Override(a AwsGateway, scope constructs.Construct, id *string, config *AwsGatewayConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-storage-gateway.AwsGateway",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsGateway)SetActivationKey(val *string) {
	if err := j.validateSetActivationKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activationKey",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetAverageDownloadRateLimitInBitsPerSec(val *float64) {
	if err := j.validateSetAverageDownloadRateLimitInBitsPerSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"averageDownloadRateLimitInBitsPerSec",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetAverageUploadRateLimitInBitsPerSec(val *float64) {
	if err := j.validateSetAverageUploadRateLimitInBitsPerSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"averageUploadRateLimitInBitsPerSec",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetCloudwatchLogGroupArn(val *string) {
	if err := j.validateSetCloudwatchLogGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchLogGroupArn",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetGatewayIpAddress(val *string) {
	if err := j.validateSetGatewayIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayIpAddress",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetGatewayName(val *string) {
	if err := j.validateSetGatewayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayName",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetGatewayTimezone(val *string) {
	if err := j.validateSetGatewayTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayTimezone",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetGatewayType(val *string) {
	if err := j.validateSetGatewayTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayType",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetGatewayVpcEndpoint(val *string) {
	if err := j.validateSetGatewayVpcEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayVpcEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetMediumChangerType(val *string) {
	if err := j.validateSetMediumChangerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mediumChangerType",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetSmbFileShareVisibility(val interface{}) {
	if err := j.validateSetSmbFileShareVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smbFileShareVisibility",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetSmbGuestPassword(val *string) {
	if err := j.validateSetSmbGuestPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smbGuestPassword",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetSmbSecurityStrategy(val *string) {
	if err := j.validateSetSmbSecurityStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smbSecurityStrategy",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsGateway)SetTapeDriveType(val *string) {
	if err := j.validateSetTapeDriveTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tapeDriveType",
		val,
	)
}

// Generates CDKTN code for importing a AwsGateway resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsGateway_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsGateway_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-storage-gateway.AwsGateway",
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
func AwsGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsGateway_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-storage-gateway.AwsGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsGateway_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsGateway_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-storage-gateway.AwsGateway",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsGateway_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsGateway_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-storage-gateway.AwsGateway",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsGateway_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-storage-gateway.AwsGateway",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsGateway) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsGateway) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsGateway) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGateway) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGateway) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGateway) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGateway) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGateway) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGateway) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGateway) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGateway) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGateway) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGateway) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsGateway) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGateway) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsGateway) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsGateway) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsGateway) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsGateway) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsGateway) PutMaintenanceStartTime(value *AwsGateway_MaintenanceStartTimeProperty) {
	if err := a.validatePutMaintenanceStartTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMaintenanceStartTime",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGateway) PutSmbActiveDirectorySettings(value *AwsGateway_SmbActiveDirectorySettingsProperty) {
	if err := a.validatePutSmbActiveDirectorySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSmbActiveDirectorySettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGateway) PutTimeouts(value *AwsGateway_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGateway) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsGateway) ResetActivationKey() {
	_jsii_.InvokeVoid(
		a,
		"resetActivationKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetAverageDownloadRateLimitInBitsPerSec() {
	_jsii_.InvokeVoid(
		a,
		"resetAverageDownloadRateLimitInBitsPerSec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetAverageUploadRateLimitInBitsPerSec() {
	_jsii_.InvokeVoid(
		a,
		"resetAverageUploadRateLimitInBitsPerSec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetCloudwatchLogGroupArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchLogGroupArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetGatewayIpAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetGatewayIpAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetGatewayType() {
	_jsii_.InvokeVoid(
		a,
		"resetGatewayType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetGatewayVpcEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetGatewayVpcEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetMaintenanceStartTime() {
	_jsii_.InvokeVoid(
		a,
		"resetMaintenanceStartTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetMediumChangerType() {
	_jsii_.InvokeVoid(
		a,
		"resetMediumChangerType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetSmbActiveDirectorySettings() {
	_jsii_.InvokeVoid(
		a,
		"resetSmbActiveDirectorySettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetSmbFileShareVisibility() {
	_jsii_.InvokeVoid(
		a,
		"resetSmbFileShareVisibility",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetSmbGuestPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetSmbGuestPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetSmbSecurityStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetSmbSecurityStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetTapeDriveType() {
	_jsii_.InvokeVoid(
		a,
		"resetTapeDriveType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGateway) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGateway) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGateway) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGateway) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGateway) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGateway) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

