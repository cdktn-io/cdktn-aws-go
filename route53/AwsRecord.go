package route53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/route53/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/route53/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record aws_route53_record}.
// Experimental.
type AwsRecord interface {
	cdktn.TerraformResource
	// Experimental.
	Alias() AwsRecord_AliasPropertyOutputReference
	// Experimental.
	AliasInput() *AwsRecord_AliasProperty
	// Experimental.
	AllowOverwrite() interface{}
	// Experimental.
	SetAllowOverwrite(val interface{})
	// Experimental.
	AllowOverwriteInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CidrRoutingPolicy() AwsRecord_CidrRoutingPolicyPropertyOutputReference
	// Experimental.
	CidrRoutingPolicyInput() *AwsRecord_CidrRoutingPolicyProperty
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
	FailoverRoutingPolicy() AwsRecord_FailoverRoutingPolicyPropertyOutputReference
	// Experimental.
	FailoverRoutingPolicyInput() *AwsRecord_FailoverRoutingPolicyProperty
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
	GeolocationRoutingPolicy() AwsRecord_GeolocationRoutingPolicyPropertyOutputReference
	// Experimental.
	GeolocationRoutingPolicyInput() *AwsRecord_GeolocationRoutingPolicyProperty
	// Experimental.
	GeoproximityRoutingPolicy() AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference
	// Experimental.
	GeoproximityRoutingPolicyInput() *AwsRecord_GeoproximityRoutingPolicyProperty
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
	LatencyRoutingPolicy() AwsRecord_LatencyRoutingPolicyPropertyOutputReference
	// Experimental.
	LatencyRoutingPolicyInput() *AwsRecord_LatencyRoutingPolicyProperty
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
	Timeouts() AwsRecord_TimeoutsPropertyOutputReference
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
	WeightedRoutingPolicy() AwsRecord_WeightedRoutingPolicyPropertyOutputReference
	// Experimental.
	WeightedRoutingPolicyInput() *AwsRecord_WeightedRoutingPolicyProperty
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
	PutAlias(value *AwsRecord_AliasProperty)
	// Experimental.
	PutCidrRoutingPolicy(value *AwsRecord_CidrRoutingPolicyProperty)
	// Experimental.
	PutFailoverRoutingPolicy(value *AwsRecord_FailoverRoutingPolicyProperty)
	// Experimental.
	PutGeolocationRoutingPolicy(value *AwsRecord_GeolocationRoutingPolicyProperty)
	// Experimental.
	PutGeoproximityRoutingPolicy(value *AwsRecord_GeoproximityRoutingPolicyProperty)
	// Experimental.
	PutLatencyRoutingPolicy(value *AwsRecord_LatencyRoutingPolicyProperty)
	// Experimental.
	PutTimeouts(value *AwsRecord_TimeoutsProperty)
	// Experimental.
	PutWeightedRoutingPolicy(value *AwsRecord_WeightedRoutingPolicyProperty)
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

// The jsii proxy struct for AwsRecord
type jsiiProxy_AwsRecord struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsRecord) Alias() AwsRecord_AliasPropertyOutputReference {
	var returns AwsRecord_AliasPropertyOutputReference
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) AliasInput() *AwsRecord_AliasProperty {
	var returns *AwsRecord_AliasProperty
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) AllowOverwrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowOverwrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) AllowOverwriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowOverwriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) CidrRoutingPolicy() AwsRecord_CidrRoutingPolicyPropertyOutputReference {
	var returns AwsRecord_CidrRoutingPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"cidrRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) CidrRoutingPolicyInput() *AwsRecord_CidrRoutingPolicyProperty {
	var returns *AwsRecord_CidrRoutingPolicyProperty
	_jsii_.Get(
		j,
		"cidrRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) FailoverRoutingPolicy() AwsRecord_FailoverRoutingPolicyPropertyOutputReference {
	var returns AwsRecord_FailoverRoutingPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"failoverRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) FailoverRoutingPolicyInput() *AwsRecord_FailoverRoutingPolicyProperty {
	var returns *AwsRecord_FailoverRoutingPolicyProperty
	_jsii_.Get(
		j,
		"failoverRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) GeolocationRoutingPolicy() AwsRecord_GeolocationRoutingPolicyPropertyOutputReference {
	var returns AwsRecord_GeolocationRoutingPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"geolocationRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) GeolocationRoutingPolicyInput() *AwsRecord_GeolocationRoutingPolicyProperty {
	var returns *AwsRecord_GeolocationRoutingPolicyProperty
	_jsii_.Get(
		j,
		"geolocationRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) GeoproximityRoutingPolicy() AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference {
	var returns AwsRecord_GeoproximityRoutingPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"geoproximityRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) GeoproximityRoutingPolicyInput() *AwsRecord_GeoproximityRoutingPolicyProperty {
	var returns *AwsRecord_GeoproximityRoutingPolicyProperty
	_jsii_.Get(
		j,
		"geoproximityRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) HealthCheckId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) HealthCheckIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) LatencyRoutingPolicy() AwsRecord_LatencyRoutingPolicyPropertyOutputReference {
	var returns AwsRecord_LatencyRoutingPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"latencyRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) LatencyRoutingPolicyInput() *AwsRecord_LatencyRoutingPolicyProperty {
	var returns *AwsRecord_LatencyRoutingPolicyProperty
	_jsii_.Get(
		j,
		"latencyRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) MultivalueAnswerRoutingPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multivalueAnswerRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) MultivalueAnswerRoutingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multivalueAnswerRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) Records() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"records",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) RecordsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) SetIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) SetIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) Timeouts() AwsRecord_TimeoutsPropertyOutputReference {
	var returns AwsRecord_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) WeightedRoutingPolicy() AwsRecord_WeightedRoutingPolicyPropertyOutputReference {
	var returns AwsRecord_WeightedRoutingPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"weightedRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) WeightedRoutingPolicyInput() *AwsRecord_WeightedRoutingPolicyProperty {
	var returns *AwsRecord_WeightedRoutingPolicyProperty
	_jsii_.Get(
		j,
		"weightedRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRecord) ZoneIdInput() *string {
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
func NewAwsRecord(scope constructs.Construct, id *string, config *AwsRecordConfig) AwsRecord {
	_init_.Initialize()

	if err := validateNewAwsRecordParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRecord{}

	_jsii_.Create(
		"@cdktn/aws-route-53.AwsRecord",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record aws_route53_record} Resource.
// Experimental.
func NewAwsRecord_Override(a AwsRecord, scope constructs.Construct, id *string, config *AwsRecordConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53.AwsRecord",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsRecord)SetAllowOverwrite(val interface{}) {
	if err := j.validateSetAllowOverwriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowOverwrite",
		val,
	)
}

func (j *jsiiProxy_AwsRecord)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsRecord)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsRecord)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsRecord)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsRecord)SetHealthCheckId(val *string) {
	if err := j.validateSetHealthCheckIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckId",
		val,
	)
}

func (j *jsiiProxy_AwsRecord)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsRecord)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsRecord)SetMultivalueAnswerRoutingPolicy(val interface{}) {
	if err := j.validateSetMultivalueAnswerRoutingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multivalueAnswerRoutingPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsRecord)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsRecord)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsRecord)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsRecord)SetRecords(val *[]*string) {
	if err := j.validateSetRecordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"records",
		val,
	)
}

func (j *jsiiProxy_AwsRecord)SetSetIdentifier(val *string) {
	if err := j.validateSetSetIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsRecord)SetTtl(val *float64) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_AwsRecord)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AwsRecord)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

// Generates CDKTN code for importing a AwsRecord resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsRecord_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsRecord_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53.AwsRecord",
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
func AwsRecord_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsRecord_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53.AwsRecord",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsRecord_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsRecord_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53.AwsRecord",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsRecord_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsRecord_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53.AwsRecord",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsRecord_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-route-53.AwsRecord",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsRecord) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsRecord) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsRecord) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRecord) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRecord) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRecord) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRecord) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRecord) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRecord) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRecord) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRecord) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRecord) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRecord) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsRecord) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRecord) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsRecord) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsRecord) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsRecord) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsRecord) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsRecord) PutAlias(value *AwsRecord_AliasProperty) {
	if err := a.validatePutAliasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAlias",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRecord) PutCidrRoutingPolicy(value *AwsRecord_CidrRoutingPolicyProperty) {
	if err := a.validatePutCidrRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCidrRoutingPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRecord) PutFailoverRoutingPolicy(value *AwsRecord_FailoverRoutingPolicyProperty) {
	if err := a.validatePutFailoverRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFailoverRoutingPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRecord) PutGeolocationRoutingPolicy(value *AwsRecord_GeolocationRoutingPolicyProperty) {
	if err := a.validatePutGeolocationRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGeolocationRoutingPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRecord) PutGeoproximityRoutingPolicy(value *AwsRecord_GeoproximityRoutingPolicyProperty) {
	if err := a.validatePutGeoproximityRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGeoproximityRoutingPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRecord) PutLatencyRoutingPolicy(value *AwsRecord_LatencyRoutingPolicyProperty) {
	if err := a.validatePutLatencyRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLatencyRoutingPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRecord) PutTimeouts(value *AwsRecord_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRecord) PutWeightedRoutingPolicy(value *AwsRecord_WeightedRoutingPolicyProperty) {
	if err := a.validatePutWeightedRoutingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWeightedRoutingPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRecord) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsRecord) ResetAlias() {
	_jsii_.InvokeVoid(
		a,
		"resetAlias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord) ResetAllowOverwrite() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowOverwrite",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord) ResetCidrRoutingPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetCidrRoutingPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord) ResetFailoverRoutingPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetFailoverRoutingPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord) ResetGeolocationRoutingPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetGeolocationRoutingPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord) ResetGeoproximityRoutingPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetGeoproximityRoutingPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord) ResetHealthCheckId() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheckId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord) ResetLatencyRoutingPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetLatencyRoutingPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord) ResetMultivalueAnswerRoutingPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetMultivalueAnswerRoutingPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord) ResetRecords() {
	_jsii_.InvokeVoid(
		a,
		"resetRecords",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord) ResetSetIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetSetIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord) ResetTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord) ResetWeightedRoutingPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetWeightedRoutingPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRecord) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRecord) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRecord) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRecord) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRecord) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRecord) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRecord) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

