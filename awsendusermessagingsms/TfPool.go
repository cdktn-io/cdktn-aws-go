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
type TfPool interface {
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
	Timeouts() TfPool_TimeoutsPropertyOutputReference
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
	PutTimeouts(value *TfPool_TimeoutsProperty)
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

// The jsii proxy struct for TfPool
type jsiiProxy_TfPool struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfPool) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) DeletionProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) DeletionProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) IsoCountryCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isoCountryCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) IsoCountryCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isoCountryCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) MessageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) MessageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) OptOutListName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optOutListName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) OptOutListNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optOutListNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) OriginationIdentities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originationIdentities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) OriginationIdentitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originationIdentitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) SelfManagedOptOutsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedOptOutsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) SelfManagedOptOutsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedOptOutsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) SharedRoutesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedRoutesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) SharedRoutesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedRoutesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) Timeouts() TfPool_TimeoutsPropertyOutputReference {
	var returns TfPool_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) TwoWayChannelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) TwoWayChannelArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) TwoWayChannelRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) TwoWayChannelRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) TwoWayEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"twoWayEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPool) TwoWayEnabledInput() interface{} {
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
func NewTfPool(scope constructs.Construct, id *string, config *TfPoolConfig) TfPool {
	_init_.Initialize()

	if err := validateNewTfPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPool{}

	_jsii_.Create(
		"@cdktn/aws-end-user-messaging-sms.TfPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_pool aws_pinpointsmsvoicev2_pool} Resource.
// Experimental.
func NewTfPool_Override(t TfPool, scope constructs.Construct, id *string, config *TfPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-end-user-messaging-sms.TfPool",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfPool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetDeletionProtectionEnabled(val interface{}) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetIsoCountryCode(val *string) {
	if err := j.validateSetIsoCountryCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isoCountryCode",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetMessageType(val *string) {
	if err := j.validateSetMessageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageType",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetOptOutListName(val *string) {
	if err := j.validateSetOptOutListNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optOutListName",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetOriginationIdentities(val *[]*string) {
	if err := j.validateSetOriginationIdentitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originationIdentities",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetSelfManagedOptOutsEnabled(val interface{}) {
	if err := j.validateSetSelfManagedOptOutsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfManagedOptOutsEnabled",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetSharedRoutesEnabled(val interface{}) {
	if err := j.validateSetSharedRoutesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedRoutesEnabled",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetTwoWayChannelArn(val *string) {
	if err := j.validateSetTwoWayChannelArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayChannelArn",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetTwoWayChannelRole(val *string) {
	if err := j.validateSetTwoWayChannelRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayChannelRole",
		val,
	)
}

func (j *jsiiProxy_TfPool)SetTwoWayEnabled(val interface{}) {
	if err := j.validateSetTwoWayEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayEnabled",
		val,
	)
}

// Generates CDKTN code for importing a TfPool resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfPool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfPool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-end-user-messaging-sms.TfPool",
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
func TfPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-end-user-messaging-sms.TfPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-end-user-messaging-sms.TfPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfPool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfPool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-end-user-messaging-sms.TfPool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-end-user-messaging-sms.TfPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfPool) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfPool) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPool) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPool) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPool) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPool) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPool) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfPool) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPool) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfPool) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfPool) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfPool) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfPool) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfPool) PutTimeouts(value *TfPool_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPool) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfPool) ResetDeletionProtectionEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetDeletionProtectionEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPool) ResetIsoCountryCode() {
	_jsii_.InvokeVoid(
		t,
		"resetIsoCountryCode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPool) ResetOptOutListName() {
	_jsii_.InvokeVoid(
		t,
		"resetOptOutListName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPool) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPool) ResetSelfManagedOptOutsEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetSelfManagedOptOutsEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPool) ResetSharedRoutesEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetSharedRoutesEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPool) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPool) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPool) ResetTwoWayChannelArn() {
	_jsii_.InvokeVoid(
		t,
		"resetTwoWayChannelArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPool) ResetTwoWayChannelRole() {
	_jsii_.InvokeVoid(
		t,
		"resetTwoWayChannelRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPool) ResetTwoWayEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetTwoWayEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPool) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

