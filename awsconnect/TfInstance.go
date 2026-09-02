package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconnect/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsconnect/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance aws_connect_instance}.
// Experimental.
type TfInstance interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	AutoResolveBestVoicesEnabled() interface{}
	// Experimental.
	SetAutoResolveBestVoicesEnabled(val interface{})
	// Experimental.
	AutoResolveBestVoicesEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	ContactFlowLogsEnabled() interface{}
	// Experimental.
	SetContactFlowLogsEnabled(val interface{})
	// Experimental.
	ContactFlowLogsEnabledInput() interface{}
	// Experimental.
	ContactLensEnabled() interface{}
	// Experimental.
	SetContactLensEnabled(val interface{})
	// Experimental.
	ContactLensEnabledInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CreatedTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DirectoryId() *string
	// Experimental.
	SetDirectoryId(val *string)
	// Experimental.
	DirectoryIdInput() *string
	// Experimental.
	EarlyMediaEnabled() interface{}
	// Experimental.
	SetEarlyMediaEnabled(val interface{})
	// Experimental.
	EarlyMediaEnabledInput() interface{}
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
	IdentityManagementType() *string
	// Experimental.
	SetIdentityManagementType(val *string)
	// Experimental.
	IdentityManagementTypeInput() *string
	// Experimental.
	IdInput() *string
	// Experimental.
	InboundCallsEnabled() interface{}
	// Experimental.
	SetInboundCallsEnabled(val interface{})
	// Experimental.
	InboundCallsEnabledInput() interface{}
	// Experimental.
	InstanceAlias() *string
	// Experimental.
	SetInstanceAlias(val *string)
	// Experimental.
	InstanceAliasInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MultiPartyConferenceEnabled() interface{}
	// Experimental.
	SetMultiPartyConferenceEnabled(val interface{})
	// Experimental.
	MultiPartyConferenceEnabledInput() interface{}
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OutboundCallsEnabled() interface{}
	// Experimental.
	SetOutboundCallsEnabled(val interface{})
	// Experimental.
	OutboundCallsEnabledInput() interface{}
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
	ServiceRole() *string
	// Experimental.
	Status() *string
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
	Timeouts() TfInstance_TimeoutsPropertyOutputReference
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
	PutTimeouts(value *TfInstance_TimeoutsProperty)
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
	ResetAutoResolveBestVoicesEnabled()
	// Experimental.
	ResetContactFlowLogsEnabled()
	// Experimental.
	ResetContactLensEnabled()
	// Experimental.
	ResetDirectoryId()
	// Experimental.
	ResetEarlyMediaEnabled()
	// Experimental.
	ResetId()
	// Experimental.
	ResetInstanceAlias()
	// Experimental.
	ResetMultiPartyConferenceEnabled()
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

// The jsii proxy struct for TfInstance
type jsiiProxy_TfInstance struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfInstance) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) AutoResolveBestVoicesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoResolveBestVoicesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) AutoResolveBestVoicesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoResolveBestVoicesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) ContactFlowLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contactFlowLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) ContactFlowLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contactFlowLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) ContactLensEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contactLensEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) ContactLensEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contactLensEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) EarlyMediaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"earlyMediaEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) EarlyMediaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"earlyMediaEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) IdentityManagementType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityManagementType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) IdentityManagementTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityManagementTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) InboundCallsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inboundCallsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) InboundCallsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inboundCallsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) InstanceAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) InstanceAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) MultiPartyConferenceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiPartyConferenceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) MultiPartyConferenceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiPartyConferenceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) OutboundCallsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outboundCallsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) OutboundCallsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outboundCallsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) Timeouts() TfInstance_TimeoutsPropertyOutputReference {
	var returns TfInstance_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance aws_connect_instance} Resource.
// Experimental.
func NewTfInstance(scope constructs.Construct, id *string, config *TfInstanceConfig) TfInstance {
	_init_.Initialize()

	if err := validateNewTfInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfInstance{}

	_jsii_.Create(
		"@cdktn/aws-connect.TfInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance aws_connect_instance} Resource.
// Experimental.
func NewTfInstance_Override(t TfInstance, scope constructs.Construct, id *string, config *TfInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect.TfInstance",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfInstance)SetAutoResolveBestVoicesEnabled(val interface{}) {
	if err := j.validateSetAutoResolveBestVoicesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoResolveBestVoicesEnabled",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetContactFlowLogsEnabled(val interface{}) {
	if err := j.validateSetContactFlowLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactFlowLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetContactLensEnabled(val interface{}) {
	if err := j.validateSetContactLensEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactLensEnabled",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetDirectoryId(val *string) {
	if err := j.validateSetDirectoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetEarlyMediaEnabled(val interface{}) {
	if err := j.validateSetEarlyMediaEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"earlyMediaEnabled",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetIdentityManagementType(val *string) {
	if err := j.validateSetIdentityManagementTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityManagementType",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetInboundCallsEnabled(val interface{}) {
	if err := j.validateSetInboundCallsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inboundCallsEnabled",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetInstanceAlias(val *string) {
	if err := j.validateSetInstanceAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceAlias",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetMultiPartyConferenceEnabled(val interface{}) {
	if err := j.validateSetMultiPartyConferenceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiPartyConferenceEnabled",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetOutboundCallsEnabled(val interface{}) {
	if err := j.validateSetOutboundCallsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundCallsEnabled",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfInstance)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a TfInstance resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-connect.TfInstance",
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
func TfInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-connect.TfInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-connect.TfInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-connect.TfInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-connect.TfInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfInstance) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfInstance) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfInstance) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstance) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfInstance) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfInstance) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfInstance) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfInstance) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfInstance) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfInstance) PutTimeouts(value *TfInstance_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInstance) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfInstance) ResetAutoResolveBestVoicesEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoResolveBestVoicesEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetContactFlowLogsEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetContactFlowLogsEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetContactLensEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetContactLensEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetDirectoryId() {
	_jsii_.InvokeVoid(
		t,
		"resetDirectoryId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetEarlyMediaEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetEarlyMediaEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetInstanceAlias() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceAlias",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetMultiPartyConferenceEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetMultiPartyConferenceEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInstance) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

