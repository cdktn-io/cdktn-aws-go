package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsroute53/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsroute53/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record aws_route53_record}.
// Experimental.
type TfRecord interface {
	cdktn.TerraformResource
	// Experimental.
	Alias() TfRecord_AliasPropertyOutputReference
	// Experimental.
	AliasInput() *TfRecord_AliasProperty
	// Experimental.
	AllowOverwrite() interface{}
	// Experimental.
	SetAllowOverwrite(val interface{})
	// Experimental.
	AllowOverwriteInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CidrRoutingPolicy() TfRecord_CidrRoutingPolicyPropertyOutputReference
	// Experimental.
	CidrRoutingPolicyInput() *TfRecord_CidrRoutingPolicyProperty
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
	FailoverRoutingPolicy() TfRecord_FailoverRoutingPolicyPropertyOutputReference
	// Experimental.
	FailoverRoutingPolicyInput() *TfRecord_FailoverRoutingPolicyProperty
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqdn() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	GeolocationRoutingPolicy() TfRecord_GeolocationRoutingPolicyPropertyOutputReference
	// Experimental.
	GeolocationRoutingPolicyInput() *TfRecord_GeolocationRoutingPolicyProperty
	// Experimental.
	GeoproximityRoutingPolicy() TfRecord_GeoproximityRoutingPolicyPropertyOutputReference
	// Experimental.
	GeoproximityRoutingPolicyInput() *TfRecord_GeoproximityRoutingPolicyProperty
	// Experimental.
	HealthCheckId() *string
	// Experimental.
	SetHealthCheckId(val *string)
	// Experimental.
	HealthCheckIdInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	LatencyRoutingPolicy() TfRecord_LatencyRoutingPolicyPropertyOutputReference
	// Experimental.
	LatencyRoutingPolicyInput() *TfRecord_LatencyRoutingPolicyProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MultivalueAnswerRoutingPolicy() interface{}
	// Experimental.
	SetMultivalueAnswerRoutingPolicy(val interface{})
	// Experimental.
	MultivalueAnswerRoutingPolicyInput() interface{}
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
	Records() *[]*string
	// Experimental.
	SetRecords(val *[]*string)
	// Experimental.
	RecordsInput() *[]*string
	// Experimental.
	SetIdentifier() *string
	// Experimental.
	SetSetIdentifier(val *string)
	// Experimental.
	SetIdentifierInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() TfRecord_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Ttl() *float64
	// Experimental.
	SetTtl(val *float64)
	// Experimental.
	TtlInput() *float64
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
	// Experimental.
	WeightedRoutingPolicy() TfRecord_WeightedRoutingPolicyPropertyOutputReference
	// Experimental.
	WeightedRoutingPolicyInput() *TfRecord_WeightedRoutingPolicyProperty
	// Experimental.
	ZoneId() *string
	// Experimental.
	SetZoneId(val *string)
	// Experimental.
	ZoneIdInput() *string
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
	PutAlias(value *TfRecord_AliasProperty)
	// Experimental.
	PutCidrRoutingPolicy(value *TfRecord_CidrRoutingPolicyProperty)
	// Experimental.
	PutFailoverRoutingPolicy(value *TfRecord_FailoverRoutingPolicyProperty)
	// Experimental.
	PutGeolocationRoutingPolicy(value *TfRecord_GeolocationRoutingPolicyProperty)
	// Experimental.
	PutGeoproximityRoutingPolicy(value *TfRecord_GeoproximityRoutingPolicyProperty)
	// Experimental.
	PutLatencyRoutingPolicy(value *TfRecord_LatencyRoutingPolicyProperty)
	// Experimental.
	PutTimeouts(value *TfRecord_TimeoutsProperty)
	// Experimental.
	PutWeightedRoutingPolicy(value *TfRecord_WeightedRoutingPolicyProperty)
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
	ResetAlias()
	// Experimental.
	ResetAllowOverwrite()
	// Experimental.
	ResetCidrRoutingPolicy()
	// Experimental.
	ResetFailoverRoutingPolicy()
	// Experimental.
	ResetGeolocationRoutingPolicy()
	// Experimental.
	ResetGeoproximityRoutingPolicy()
	// Experimental.
	ResetHealthCheckId()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLatencyRoutingPolicy()
	// Experimental.
	ResetMultivalueAnswerRoutingPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRecords()
	// Experimental.
	ResetSetIdentifier()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTtl()
	// Experimental.
	ResetWeightedRoutingPolicy()
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

// The jsii proxy struct for TfRecord
type jsiiProxy_TfRecord struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfRecord) Alias() TfRecord_AliasPropertyOutputReference {
	var returns TfRecord_AliasPropertyOutputReference
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) AliasInput() *TfRecord_AliasProperty {
	var returns *TfRecord_AliasProperty
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) AllowOverwrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowOverwrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) AllowOverwriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowOverwriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) CidrRoutingPolicy() TfRecord_CidrRoutingPolicyPropertyOutputReference {
	var returns TfRecord_CidrRoutingPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"cidrRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) CidrRoutingPolicyInput() *TfRecord_CidrRoutingPolicyProperty {
	var returns *TfRecord_CidrRoutingPolicyProperty
	_jsii_.Get(
		j,
		"cidrRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) FailoverRoutingPolicy() TfRecord_FailoverRoutingPolicyPropertyOutputReference {
	var returns TfRecord_FailoverRoutingPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"failoverRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) FailoverRoutingPolicyInput() *TfRecord_FailoverRoutingPolicyProperty {
	var returns *TfRecord_FailoverRoutingPolicyProperty
	_jsii_.Get(
		j,
		"failoverRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) GeolocationRoutingPolicy() TfRecord_GeolocationRoutingPolicyPropertyOutputReference {
	var returns TfRecord_GeolocationRoutingPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"geolocationRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) GeolocationRoutingPolicyInput() *TfRecord_GeolocationRoutingPolicyProperty {
	var returns *TfRecord_GeolocationRoutingPolicyProperty
	_jsii_.Get(
		j,
		"geolocationRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) GeoproximityRoutingPolicy() TfRecord_GeoproximityRoutingPolicyPropertyOutputReference {
	var returns TfRecord_GeoproximityRoutingPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"geoproximityRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) GeoproximityRoutingPolicyInput() *TfRecord_GeoproximityRoutingPolicyProperty {
	var returns *TfRecord_GeoproximityRoutingPolicyProperty
	_jsii_.Get(
		j,
		"geoproximityRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) HealthCheckId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) HealthCheckIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) LatencyRoutingPolicy() TfRecord_LatencyRoutingPolicyPropertyOutputReference {
	var returns TfRecord_LatencyRoutingPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"latencyRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) LatencyRoutingPolicyInput() *TfRecord_LatencyRoutingPolicyProperty {
	var returns *TfRecord_LatencyRoutingPolicyProperty
	_jsii_.Get(
		j,
		"latencyRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) MultivalueAnswerRoutingPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multivalueAnswerRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) MultivalueAnswerRoutingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multivalueAnswerRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) Records() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"records",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) RecordsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) SetIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) SetIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) Timeouts() TfRecord_TimeoutsPropertyOutputReference {
	var returns TfRecord_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) WeightedRoutingPolicy() TfRecord_WeightedRoutingPolicyPropertyOutputReference {
	var returns TfRecord_WeightedRoutingPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"weightedRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) WeightedRoutingPolicyInput() *TfRecord_WeightedRoutingPolicyProperty {
	var returns *TfRecord_WeightedRoutingPolicyProperty
	_jsii_.Get(
		j,
		"weightedRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRecord) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record aws_route53_record} Resource.
// Experimental.
func NewTfRecord(scope constructs.Construct, id *string, config *TfRecordConfig) TfRecord {
	_init_.Initialize()

	if err := validateNewTfRecordParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRecord{}

	_jsii_.Create(
		"@cdktn/aws-route-53.TfRecord",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record aws_route53_record} Resource.
// Experimental.
func NewTfRecord_Override(t TfRecord, scope constructs.Construct, id *string, config *TfRecordConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53.TfRecord",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfRecord)SetAllowOverwrite(val interface{}) {
	if err := j.validateSetAllowOverwriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowOverwrite",
		val,
	)
}

func (j *jsiiProxy_TfRecord)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfRecord)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfRecord)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfRecord)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfRecord)SetHealthCheckId(val *string) {
	if err := j.validateSetHealthCheckIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckId",
		val,
	)
}

func (j *jsiiProxy_TfRecord)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfRecord)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfRecord)SetMultivalueAnswerRoutingPolicy(val interface{}) {
	if err := j.validateSetMultivalueAnswerRoutingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multivalueAnswerRoutingPolicy",
		val,
	)
}

func (j *jsiiProxy_TfRecord)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfRecord)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfRecord)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfRecord)SetRecords(val *[]*string) {
	if err := j.validateSetRecordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"records",
		val,
	)
}

func (j *jsiiProxy_TfRecord)SetSetIdentifier(val *string) {
	if err := j.validateSetSetIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setIdentifier",
		val,
	)
}

func (j *jsiiProxy_TfRecord)SetTtl(val *float64) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_TfRecord)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_TfRecord)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

// Generates CDKTN code for importing a TfRecord resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfRecord_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfRecord_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53.TfRecord",
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
func TfRecord_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfRecord_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53.TfRecord",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfRecord_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfRecord_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53.TfRecord",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfRecord_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfRecord_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53.TfRecord",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfRecord_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-route-53.TfRecord",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfRecord) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfRecord) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfRecord) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRecord) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRecord) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRecord) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRecord) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRecord) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRecord) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRecord) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRecord) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRecord) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRecord) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfRecord) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRecord) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfRecord) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfRecord) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfRecord) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfRecord) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfRecord) PutAlias(value *TfRecord_AliasProperty) {
	if err := t.validatePutAliasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAlias",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRecord) PutCidrRoutingPolicy(value *TfRecord_CidrRoutingPolicyProperty) {
	if err := t.validatePutCidrRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCidrRoutingPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRecord) PutFailoverRoutingPolicy(value *TfRecord_FailoverRoutingPolicyProperty) {
	if err := t.validatePutFailoverRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFailoverRoutingPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRecord) PutGeolocationRoutingPolicy(value *TfRecord_GeolocationRoutingPolicyProperty) {
	if err := t.validatePutGeolocationRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGeolocationRoutingPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRecord) PutGeoproximityRoutingPolicy(value *TfRecord_GeoproximityRoutingPolicyProperty) {
	if err := t.validatePutGeoproximityRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGeoproximityRoutingPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRecord) PutLatencyRoutingPolicy(value *TfRecord_LatencyRoutingPolicyProperty) {
	if err := t.validatePutLatencyRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLatencyRoutingPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRecord) PutTimeouts(value *TfRecord_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRecord) PutWeightedRoutingPolicy(value *TfRecord_WeightedRoutingPolicyProperty) {
	if err := t.validatePutWeightedRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWeightedRoutingPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRecord) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfRecord) ResetAlias() {
	_jsii_.InvokeVoid(
		t,
		"resetAlias",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecord) ResetAllowOverwrite() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowOverwrite",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecord) ResetCidrRoutingPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetCidrRoutingPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecord) ResetFailoverRoutingPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetFailoverRoutingPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecord) ResetGeolocationRoutingPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetGeolocationRoutingPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecord) ResetGeoproximityRoutingPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetGeoproximityRoutingPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecord) ResetHealthCheckId() {
	_jsii_.InvokeVoid(
		t,
		"resetHealthCheckId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecord) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecord) ResetLatencyRoutingPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetLatencyRoutingPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecord) ResetMultivalueAnswerRoutingPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetMultivalueAnswerRoutingPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecord) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecord) ResetRecords() {
	_jsii_.InvokeVoid(
		t,
		"resetRecords",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecord) ResetSetIdentifier() {
	_jsii_.InvokeVoid(
		t,
		"resetSetIdentifier",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecord) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecord) ResetTtl() {
	_jsii_.InvokeVoid(
		t,
		"resetTtl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecord) ResetWeightedRoutingPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetWeightedRoutingPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRecord) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRecord) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRecord) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRecord) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRecord) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRecord) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRecord) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

