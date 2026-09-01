package awsendusermessagingsms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsendusermessagingsms/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsendusermessagingsms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_pool aws_pinpointsmsvoicev2_pool}.
// Experimental.
type AwsPinpointsmsvoicev2Pool interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
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
	DeletionProtectionEnabled() interface{}
	// Experimental.
	SetDeletionProtectionEnabled(val interface{})
	// Experimental.
	DeletionProtectionEnabledInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	IsoCountryCode() *string
	// Experimental.
	SetIsoCountryCode(val *string)
	// Experimental.
	IsoCountryCodeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MessageType() *string
	// Experimental.
	SetMessageType(val *string)
	// Experimental.
	MessageTypeInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OptOutListName() *string
	// Experimental.
	SetOptOutListName(val *string)
	// Experimental.
	OptOutListNameInput() *string
	// Experimental.
	OriginationIdentities() *[]*string
	// Experimental.
	SetOriginationIdentities(val *[]*string)
	// Experimental.
	OriginationIdentitiesInput() *[]*string
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
	SelfManagedOptOutsEnabled() interface{}
	// Experimental.
	SetSelfManagedOptOutsEnabled(val interface{})
	// Experimental.
	SelfManagedOptOutsEnabledInput() interface{}
	// Experimental.
	SharedRoutesEnabled() interface{}
	// Experimental.
	SetSharedRoutesEnabled(val interface{})
	// Experimental.
	SharedRoutesEnabledInput() interface{}
	// Experimental.
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsAll() cdktn.StringMap
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsPinpointsmsvoicev2Pool_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TwoWayChannelArn() *string
	// Experimental.
	SetTwoWayChannelArn(val *string)
	// Experimental.
	TwoWayChannelArnInput() *string
	// Experimental.
	TwoWayChannelRole() *string
	// Experimental.
	SetTwoWayChannelRole(val *string)
	// Experimental.
	TwoWayChannelRoleInput() *string
	// Experimental.
	TwoWayEnabled() interface{}
	// Experimental.
	SetTwoWayEnabled(val interface{})
	// Experimental.
	TwoWayEnabledInput() interface{}
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
	PutTimeouts(value *AwsPinpointsmsvoicev2Pool_TimeoutsProperty)
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
	ResetDeletionProtectionEnabled()
	// Experimental.
	ResetIsoCountryCode()
	// Experimental.
	ResetOptOutListName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSelfManagedOptOutsEnabled()
	// Experimental.
	ResetSharedRoutesEnabled()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTwoWayChannelArn()
	// Experimental.
	ResetTwoWayChannelRole()
	// Experimental.
	ResetTwoWayEnabled()
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

// The jsii proxy struct for AwsPinpointsmsvoicev2Pool
type jsiiProxy_AwsPinpointsmsvoicev2Pool struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) DeletionProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) DeletionProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) IsoCountryCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isoCountryCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) IsoCountryCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isoCountryCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) MessageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) MessageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) OptOutListName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optOutListName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) OptOutListNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optOutListNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) OriginationIdentities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originationIdentities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) OriginationIdentitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originationIdentitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) SelfManagedOptOutsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedOptOutsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) SelfManagedOptOutsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedOptOutsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) SharedRoutesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedRoutesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) SharedRoutesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedRoutesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) Timeouts() AwsPinpointsmsvoicev2Pool_TimeoutsPropertyOutputReference {
	var returns AwsPinpointsmsvoicev2Pool_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) TwoWayChannelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) TwoWayChannelArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) TwoWayChannelRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) TwoWayChannelRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) TwoWayEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"twoWayEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool) TwoWayEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"twoWayEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_pool aws_pinpointsmsvoicev2_pool} Resource.
// Experimental.
func NewAwsPinpointsmsvoicev2Pool(scope constructs.Construct, id *string, config *AwsPinpointsmsvoicev2PoolConfig) AwsPinpointsmsvoicev2Pool {
	_init_.Initialize()

	if err := validateNewAwsPinpointsmsvoicev2PoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPinpointsmsvoicev2Pool{}

	_jsii_.Create(
		"@cdktn/aws-end-user-messaging-sms.AwsPinpointsmsvoicev2Pool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_pool aws_pinpointsmsvoicev2_pool} Resource.
// Experimental.
func NewAwsPinpointsmsvoicev2Pool_Override(a AwsPinpointsmsvoicev2Pool, scope constructs.Construct, id *string, config *AwsPinpointsmsvoicev2PoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-end-user-messaging-sms.AwsPinpointsmsvoicev2Pool",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetDeletionProtectionEnabled(val interface{}) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetIsoCountryCode(val *string) {
	if err := j.validateSetIsoCountryCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isoCountryCode",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetMessageType(val *string) {
	if err := j.validateSetMessageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageType",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetOptOutListName(val *string) {
	if err := j.validateSetOptOutListNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optOutListName",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetOriginationIdentities(val *[]*string) {
	if err := j.validateSetOriginationIdentitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originationIdentities",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetSelfManagedOptOutsEnabled(val interface{}) {
	if err := j.validateSetSelfManagedOptOutsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfManagedOptOutsEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetSharedRoutesEnabled(val interface{}) {
	if err := j.validateSetSharedRoutesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedRoutesEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetTwoWayChannelArn(val *string) {
	if err := j.validateSetTwoWayChannelArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayChannelArn",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetTwoWayChannelRole(val *string) {
	if err := j.validateSetTwoWayChannelRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayChannelRole",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointsmsvoicev2Pool)SetTwoWayEnabled(val interface{}) {
	if err := j.validateSetTwoWayEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayEnabled",
		val,
	)
}

// Generates CDKTN code for importing a AwsPinpointsmsvoicev2Pool resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsPinpointsmsvoicev2Pool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsPinpointsmsvoicev2Pool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-end-user-messaging-sms.AwsPinpointsmsvoicev2Pool",
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
func AwsPinpointsmsvoicev2Pool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsPinpointsmsvoicev2Pool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-end-user-messaging-sms.AwsPinpointsmsvoicev2Pool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsPinpointsmsvoicev2Pool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsPinpointsmsvoicev2Pool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-end-user-messaging-sms.AwsPinpointsmsvoicev2Pool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsPinpointsmsvoicev2Pool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsPinpointsmsvoicev2Pool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-end-user-messaging-sms.AwsPinpointsmsvoicev2Pool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsPinpointsmsvoicev2Pool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-end-user-messaging-sms.AwsPinpointsmsvoicev2Pool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) PutTimeouts(value *AwsPinpointsmsvoicev2Pool_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) ResetDeletionProtectionEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDeletionProtectionEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) ResetIsoCountryCode() {
	_jsii_.InvokeVoid(
		a,
		"resetIsoCountryCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) ResetOptOutListName() {
	_jsii_.InvokeVoid(
		a,
		"resetOptOutListName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) ResetSelfManagedOptOutsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetSelfManagedOptOutsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) ResetSharedRoutesEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedRoutesEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) ResetTwoWayChannelArn() {
	_jsii_.InvokeVoid(
		a,
		"resetTwoWayChannelArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) ResetTwoWayChannelRole() {
	_jsii_.InvokeVoid(
		a,
		"resetTwoWayChannelRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) ResetTwoWayEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTwoWayEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPinpointsmsvoicev2Pool) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

