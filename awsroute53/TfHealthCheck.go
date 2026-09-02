package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsroute53/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsroute53/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check aws_route53_health_check}.
// Experimental.
type TfHealthCheck interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ChildHealthchecks() *[]*string
	// Experimental.
	SetChildHealthchecks(val *[]*string)
	// Experimental.
	ChildHealthchecksInput() *[]*string
	// Experimental.
	ChildHealthThreshold() *float64
	// Experimental.
	SetChildHealthThreshold(val *float64)
	// Experimental.
	ChildHealthThresholdInput() *float64
	// Experimental.
	CloudwatchAlarmName() *string
	// Experimental.
	SetCloudwatchAlarmName(val *string)
	// Experimental.
	CloudwatchAlarmNameInput() *string
	// Experimental.
	CloudwatchAlarmRegion() *string
	// Experimental.
	SetCloudwatchAlarmRegion(val *string)
	// Experimental.
	CloudwatchAlarmRegionInput() *string
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
	Disabled() interface{}
	// Experimental.
	SetDisabled(val interface{})
	// Experimental.
	DisabledInput() interface{}
	// Experimental.
	EnableSni() interface{}
	// Experimental.
	SetEnableSni(val interface{})
	// Experimental.
	EnableSniInput() interface{}
	// Experimental.
	FailureThreshold() *float64
	// Experimental.
	SetFailureThreshold(val *float64)
	// Experimental.
	FailureThresholdInput() *float64
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqdn() *string
	// Experimental.
	SetFqdn(val *string)
	// Experimental.
	FqdnInput() *string
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
	InsufficientDataHealthStatus() *string
	// Experimental.
	SetInsufficientDataHealthStatus(val *string)
	// Experimental.
	InsufficientDataHealthStatusInput() *string
	// Experimental.
	InvertHealthcheck() interface{}
	// Experimental.
	SetInvertHealthcheck(val interface{})
	// Experimental.
	InvertHealthcheckInput() interface{}
	// Experimental.
	IpAddress() *string
	// Experimental.
	SetIpAddress(val *string)
	// Experimental.
	IpAddressInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MeasureLatency() interface{}
	// Experimental.
	SetMeasureLatency(val interface{})
	// Experimental.
	MeasureLatencyInput() interface{}
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Port() *float64
	// Experimental.
	SetPort(val *float64)
	// Experimental.
	PortInput() *float64
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
	ReferenceName() *string
	// Experimental.
	SetReferenceName(val *string)
	// Experimental.
	ReferenceNameInput() *string
	// Experimental.
	Regions() *[]*string
	// Experimental.
	SetRegions(val *[]*string)
	// Experimental.
	RegionsInput() *[]*string
	// Experimental.
	RequestInterval() *float64
	// Experimental.
	SetRequestInterval(val *float64)
	// Experimental.
	RequestIntervalInput() *float64
	// Experimental.
	ResourcePath() *string
	// Experimental.
	SetResourcePath(val *string)
	// Experimental.
	ResourcePathInput() *string
	// Experimental.
	RoutingControlArn() *string
	// Experimental.
	SetRoutingControlArn(val *string)
	// Experimental.
	RoutingControlArnInput() *string
	// Experimental.
	SearchString() *string
	// Experimental.
	SetSearchString(val *string)
	// Experimental.
	SearchStringInput() *string
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
	Triggers() *map[string]*string
	// Experimental.
	SetTriggers(val *map[string]*string)
	// Experimental.
	TriggersInput() *map[string]*string
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	ResetChildHealthchecks()
	// Experimental.
	ResetChildHealthThreshold()
	// Experimental.
	ResetCloudwatchAlarmName()
	// Experimental.
	ResetCloudwatchAlarmRegion()
	// Experimental.
	ResetDisabled()
	// Experimental.
	ResetEnableSni()
	// Experimental.
	ResetFailureThreshold()
	// Experimental.
	ResetFqdn()
	// Experimental.
	ResetId()
	// Experimental.
	ResetInsufficientDataHealthStatus()
	// Experimental.
	ResetInvertHealthcheck()
	// Experimental.
	ResetIpAddress()
	// Experimental.
	ResetMeasureLatency()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPort()
	// Experimental.
	ResetReferenceName()
	// Experimental.
	ResetRegions()
	// Experimental.
	ResetRequestInterval()
	// Experimental.
	ResetResourcePath()
	// Experimental.
	ResetRoutingControlArn()
	// Experimental.
	ResetSearchString()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTriggers()
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

// The jsii proxy struct for TfHealthCheck
type jsiiProxy_TfHealthCheck struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfHealthCheck) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) ChildHealthchecks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"childHealthchecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) ChildHealthchecksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"childHealthchecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) ChildHealthThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"childHealthThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) ChildHealthThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"childHealthThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) CloudwatchAlarmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchAlarmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) CloudwatchAlarmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchAlarmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) CloudwatchAlarmRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchAlarmRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) CloudwatchAlarmRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchAlarmRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) EnableSni() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSni",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) EnableSniInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSniInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) FailureThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) FailureThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) FqdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) InsufficientDataHealthStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insufficientDataHealthStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) InsufficientDataHealthStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insufficientDataHealthStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) InvertHealthcheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invertHealthcheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) InvertHealthcheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invertHealthcheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) MeasureLatency() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"measureLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) MeasureLatencyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"measureLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) ReferenceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) ReferenceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) RequestInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) RequestIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) ResourcePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) ResourcePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourcePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) RoutingControlArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingControlArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) RoutingControlArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingControlArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) SearchString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) SearchStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) Triggers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) TriggersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHealthCheck) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check aws_route53_health_check} Resource.
// Experimental.
func NewTfHealthCheck(scope constructs.Construct, id *string, config *TfHealthCheckConfig) TfHealthCheck {
	_init_.Initialize()

	if err := validateNewTfHealthCheckParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfHealthCheck{}

	_jsii_.Create(
		"@cdktn/aws-route-53.TfHealthCheck",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_health_check aws_route53_health_check} Resource.
// Experimental.
func NewTfHealthCheck_Override(t TfHealthCheck, scope constructs.Construct, id *string, config *TfHealthCheckConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53.TfHealthCheck",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetChildHealthchecks(val *[]*string) {
	if err := j.validateSetChildHealthchecksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"childHealthchecks",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetChildHealthThreshold(val *float64) {
	if err := j.validateSetChildHealthThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"childHealthThreshold",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetCloudwatchAlarmName(val *string) {
	if err := j.validateSetCloudwatchAlarmNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchAlarmName",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetCloudwatchAlarmRegion(val *string) {
	if err := j.validateSetCloudwatchAlarmRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchAlarmRegion",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetEnableSni(val interface{}) {
	if err := j.validateSetEnableSniParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSni",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetFailureThreshold(val *float64) {
	if err := j.validateSetFailureThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureThreshold",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetFqdn(val *string) {
	if err := j.validateSetFqdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fqdn",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetInsufficientDataHealthStatus(val *string) {
	if err := j.validateSetInsufficientDataHealthStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insufficientDataHealthStatus",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetInvertHealthcheck(val interface{}) {
	if err := j.validateSetInvertHealthcheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invertHealthcheck",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetMeasureLatency(val interface{}) {
	if err := j.validateSetMeasureLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measureLatency",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetReferenceName(val *string) {
	if err := j.validateSetReferenceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referenceName",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetRequestInterval(val *float64) {
	if err := j.validateSetRequestIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestInterval",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetResourcePath(val *string) {
	if err := j.validateSetResourcePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePath",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetRoutingControlArn(val *string) {
	if err := j.validateSetRoutingControlArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingControlArn",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetSearchString(val *string) {
	if err := j.validateSetSearchStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchString",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetTriggers(val *map[string]*string) {
	if err := j.validateSetTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggers",
		val,
	)
}

func (j *jsiiProxy_TfHealthCheck)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTN code for importing a TfHealthCheck resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfHealthCheck_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfHealthCheck_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53.TfHealthCheck",
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
func TfHealthCheck_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfHealthCheck_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53.TfHealthCheck",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfHealthCheck_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfHealthCheck_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53.TfHealthCheck",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfHealthCheck_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfHealthCheck_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53.TfHealthCheck",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfHealthCheck_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-route-53.TfHealthCheck",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfHealthCheck) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfHealthCheck) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfHealthCheck) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfHealthCheck) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHealthCheck) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfHealthCheck) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfHealthCheck) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfHealthCheck) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfHealthCheck) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfHealthCheck) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfHealthCheck) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfHealthCheck) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHealthCheck) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfHealthCheck) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHealthCheck) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfHealthCheck) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfHealthCheck) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfHealthCheck) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfHealthCheck) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfHealthCheck) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetChildHealthchecks() {
	_jsii_.InvokeVoid(
		t,
		"resetChildHealthchecks",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetChildHealthThreshold() {
	_jsii_.InvokeVoid(
		t,
		"resetChildHealthThreshold",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetCloudwatchAlarmName() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchAlarmName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetCloudwatchAlarmRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchAlarmRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetDisabled() {
	_jsii_.InvokeVoid(
		t,
		"resetDisabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetEnableSni() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableSni",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetFailureThreshold() {
	_jsii_.InvokeVoid(
		t,
		"resetFailureThreshold",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetFqdn() {
	_jsii_.InvokeVoid(
		t,
		"resetFqdn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetInsufficientDataHealthStatus() {
	_jsii_.InvokeVoid(
		t,
		"resetInsufficientDataHealthStatus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetInvertHealthcheck() {
	_jsii_.InvokeVoid(
		t,
		"resetInvertHealthcheck",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetIpAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetMeasureLatency() {
	_jsii_.InvokeVoid(
		t,
		"resetMeasureLatency",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetPort() {
	_jsii_.InvokeVoid(
		t,
		"resetPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetReferenceName() {
	_jsii_.InvokeVoid(
		t,
		"resetReferenceName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetRegions() {
	_jsii_.InvokeVoid(
		t,
		"resetRegions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetRequestInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetRequestInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetResourcePath() {
	_jsii_.InvokeVoid(
		t,
		"resetResourcePath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetRoutingControlArn() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingControlArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetSearchString() {
	_jsii_.InvokeVoid(
		t,
		"resetSearchString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) ResetTriggers() {
	_jsii_.InvokeVoid(
		t,
		"resetTriggers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHealthCheck) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHealthCheck) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHealthCheck) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHealthCheck) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHealthCheck) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHealthCheck) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHealthCheck) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

