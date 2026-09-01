package awsfms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfms/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsfms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy aws_fms_policy}.
// Experimental.
type AwsFmsPolicy interface {
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
	DeleteAllPolicyResources() interface{}
	// Experimental.
	SetDeleteAllPolicyResources(val interface{})
	// Experimental.
	DeleteAllPolicyResourcesInput() interface{}
	// Experimental.
	DeleteUnusedFmManagedResources() interface{}
	// Experimental.
	SetDeleteUnusedFmManagedResources(val interface{})
	// Experimental.
	DeleteUnusedFmManagedResourcesInput() interface{}
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
	ExcludeMap() AwsFmsPolicy_ExcludeMapPropertyOutputReference
	// Experimental.
	ExcludeMapInput() *AwsFmsPolicy_ExcludeMapProperty
	// Experimental.
	ExcludeResourceTags() interface{}
	// Experimental.
	SetExcludeResourceTags(val interface{})
	// Experimental.
	ExcludeResourceTagsInput() interface{}
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
	IncludeMap() AwsFmsPolicy_IncludeMapPropertyOutputReference
	// Experimental.
	IncludeMapInput() *AwsFmsPolicy_IncludeMapProperty
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
	PolicyUpdateToken() *string
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
	RemediationEnabled() interface{}
	// Experimental.
	SetRemediationEnabled(val interface{})
	// Experimental.
	RemediationEnabledInput() interface{}
	// Experimental.
	ResourceSetIds() *[]*string
	// Experimental.
	SetResourceSetIds(val *[]*string)
	// Experimental.
	ResourceSetIdsInput() *[]*string
	// Experimental.
	ResourceTagLogicalOperator() *string
	// Experimental.
	SetResourceTagLogicalOperator(val *string)
	// Experimental.
	ResourceTagLogicalOperatorInput() *string
	// Experimental.
	ResourceTags() *map[string]*string
	// Experimental.
	SetResourceTags(val *map[string]*string)
	// Experimental.
	ResourceTagsInput() *map[string]*string
	// Experimental.
	ResourceType() *string
	// Experimental.
	SetResourceType(val *string)
	// Experimental.
	ResourceTypeInput() *string
	// Experimental.
	ResourceTypeList() *[]*string
	// Experimental.
	SetResourceTypeList(val *[]*string)
	// Experimental.
	ResourceTypeListInput() *[]*string
	// Experimental.
	SecurityServicePolicyData() AwsFmsPolicy_SecurityServicePolicyDataPropertyOutputReference
	// Experimental.
	SecurityServicePolicyDataInput() *AwsFmsPolicy_SecurityServicePolicyDataProperty
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
	// Experimental.
	PutExcludeMap(value *AwsFmsPolicy_ExcludeMapProperty)
	// Experimental.
	PutIncludeMap(value *AwsFmsPolicy_IncludeMapProperty)
	// Experimental.
	PutSecurityServicePolicyData(value *AwsFmsPolicy_SecurityServicePolicyDataProperty)
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
	ResetDeleteAllPolicyResources()
	// Experimental.
	ResetDeleteUnusedFmManagedResources()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetExcludeMap()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIncludeMap()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRemediationEnabled()
	// Experimental.
	ResetResourceSetIds()
	// Experimental.
	ResetResourceTagLogicalOperator()
	// Experimental.
	ResetResourceTags()
	// Experimental.
	ResetResourceType()
	// Experimental.
	ResetResourceTypeList()
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

// The jsii proxy struct for AwsFmsPolicy
type jsiiProxy_AwsFmsPolicy struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsFmsPolicy) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) DeleteAllPolicyResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAllPolicyResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) DeleteAllPolicyResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAllPolicyResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) DeleteUnusedFmManagedResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteUnusedFmManagedResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) DeleteUnusedFmManagedResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteUnusedFmManagedResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) ExcludeMap() AwsFmsPolicy_ExcludeMapPropertyOutputReference {
	var returns AwsFmsPolicy_ExcludeMapPropertyOutputReference
	_jsii_.Get(
		j,
		"excludeMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) ExcludeMapInput() *AwsFmsPolicy_ExcludeMapProperty {
	var returns *AwsFmsPolicy_ExcludeMapProperty
	_jsii_.Get(
		j,
		"excludeMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) ExcludeResourceTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeResourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) ExcludeResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeResourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) IncludeMap() AwsFmsPolicy_IncludeMapPropertyOutputReference {
	var returns AwsFmsPolicy_IncludeMapPropertyOutputReference
	_jsii_.Get(
		j,
		"includeMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) IncludeMapInput() *AwsFmsPolicy_IncludeMapProperty {
	var returns *AwsFmsPolicy_IncludeMapProperty
	_jsii_.Get(
		j,
		"includeMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) PolicyUpdateToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyUpdateToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) RemediationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remediationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) RemediationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remediationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) ResourceSetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceSetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) ResourceSetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceSetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) ResourceTagLogicalOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTagLogicalOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) ResourceTagLogicalOperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTagLogicalOperatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) ResourceTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) ResourceTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) ResourceTypeList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypeList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) ResourceTypeListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypeListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) SecurityServicePolicyData() AwsFmsPolicy_SecurityServicePolicyDataPropertyOutputReference {
	var returns AwsFmsPolicy_SecurityServicePolicyDataPropertyOutputReference
	_jsii_.Get(
		j,
		"securityServicePolicyData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) SecurityServicePolicyDataInput() *AwsFmsPolicy_SecurityServicePolicyDataProperty {
	var returns *AwsFmsPolicy_SecurityServicePolicyDataProperty
	_jsii_.Get(
		j,
		"securityServicePolicyDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFmsPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy aws_fms_policy} Resource.
// Experimental.
func NewAwsFmsPolicy(scope constructs.Construct, id *string, config *AwsFmsPolicyConfig) AwsFmsPolicy {
	_init_.Initialize()

	if err := validateNewAwsFmsPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFmsPolicy{}

	_jsii_.Create(
		"@cdktn/aws-fms.AwsFmsPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy aws_fms_policy} Resource.
// Experimental.
func NewAwsFmsPolicy_Override(a AwsFmsPolicy, scope constructs.Construct, id *string, config *AwsFmsPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fms.AwsFmsPolicy",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetDeleteAllPolicyResources(val interface{}) {
	if err := j.validateSetDeleteAllPolicyResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAllPolicyResources",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetDeleteUnusedFmManagedResources(val interface{}) {
	if err := j.validateSetDeleteUnusedFmManagedResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteUnusedFmManagedResources",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetExcludeResourceTags(val interface{}) {
	if err := j.validateSetExcludeResourceTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeResourceTags",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetRemediationEnabled(val interface{}) {
	if err := j.validateSetRemediationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remediationEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetResourceSetIds(val *[]*string) {
	if err := j.validateSetResourceSetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceSetIds",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetResourceTagLogicalOperator(val *string) {
	if err := j.validateSetResourceTagLogicalOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTagLogicalOperator",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetResourceTags(val *map[string]*string) {
	if err := j.validateSetResourceTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTags",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetResourceTypeList(val *[]*string) {
	if err := j.validateSetResourceTypeListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypeList",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsFmsPolicy)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a AwsFmsPolicy resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsFmsPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsFmsPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-fms.AwsFmsPolicy",
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
func AwsFmsPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsFmsPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fms.AwsFmsPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsFmsPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsFmsPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fms.AwsFmsPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsFmsPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsFmsPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fms.AwsFmsPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsFmsPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-fms.AwsFmsPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsFmsPolicy) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsFmsPolicy) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsFmsPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFmsPolicy) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFmsPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFmsPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFmsPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFmsPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFmsPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFmsPolicy) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFmsPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFmsPolicy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFmsPolicy) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsFmsPolicy) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFmsPolicy) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsFmsPolicy) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsFmsPolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsFmsPolicy) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsFmsPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsFmsPolicy) PutExcludeMap(value *AwsFmsPolicy_ExcludeMapProperty) {
	if err := a.validatePutExcludeMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExcludeMap",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFmsPolicy) PutIncludeMap(value *AwsFmsPolicy_IncludeMapProperty) {
	if err := a.validatePutIncludeMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIncludeMap",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFmsPolicy) PutSecurityServicePolicyData(value *AwsFmsPolicy_SecurityServicePolicyDataProperty) {
	if err := a.validatePutSecurityServicePolicyDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecurityServicePolicyData",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFmsPolicy) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsFmsPolicy) ResetDeleteAllPolicyResources() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteAllPolicyResources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy) ResetDeleteUnusedFmManagedResources() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteUnusedFmManagedResources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy) ResetExcludeMap() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludeMap",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy) ResetIncludeMap() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeMap",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy) ResetRemediationEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetRemediationEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy) ResetResourceSetIds() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceSetIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy) ResetResourceTagLogicalOperator() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceTagLogicalOperator",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy) ResetResourceTags() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy) ResetResourceType() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy) ResetResourceTypeList() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceTypeList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFmsPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFmsPolicy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFmsPolicy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFmsPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFmsPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFmsPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFmsPolicy) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

