package vpcipam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/vpcipam/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/vpcipam/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool aws_vpc_ipam_pool}.
// Experimental.
type AwsPool interface {
	cdktn.TerraformResource
	// Experimental.
	AddressFamily() *string
	// Experimental.
	SetAddressFamily(val *string)
	// Experimental.
	AddressFamilyInput() *string
	// Experimental.
	AllocationDefaultNetmaskLength() *float64
	// Experimental.
	SetAllocationDefaultNetmaskLength(val *float64)
	// Experimental.
	AllocationDefaultNetmaskLengthInput() *float64
	// Experimental.
	AllocationMaxNetmaskLength() *float64
	// Experimental.
	SetAllocationMaxNetmaskLength(val *float64)
	// Experimental.
	AllocationMaxNetmaskLengthInput() *float64
	// Experimental.
	AllocationMinNetmaskLength() *float64
	// Experimental.
	SetAllocationMinNetmaskLength(val *float64)
	// Experimental.
	AllocationMinNetmaskLengthInput() *float64
	// Experimental.
	AllocationResourceTags() *map[string]*string
	// Experimental.
	SetAllocationResourceTags(val *map[string]*string)
	// Experimental.
	AllocationResourceTagsInput() *map[string]*string
	// Experimental.
	Arn() *string
	// Experimental.
	AutoImport() interface{}
	// Experimental.
	SetAutoImport(val interface{})
	// Experimental.
	AutoImportInput() interface{}
	// Experimental.
	AwsService() *string
	// Experimental.
	SetAwsService(val *string)
	// Experimental.
	AwsServiceInput() *string
	// Experimental.
	Cascade() interface{}
	// Experimental.
	SetCascade(val interface{})
	// Experimental.
	CascadeInput() interface{}
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
	IpamScopeId() *string
	// Experimental.
	SetIpamScopeId(val *string)
	// Experimental.
	IpamScopeIdInput() *string
	// Experimental.
	IpamScopeType() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	Locale() *string
	// Experimental.
	SetLocale(val *string)
	// Experimental.
	LocaleInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PoolDepth() *float64
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	PublicIpSource() *string
	// Experimental.
	SetPublicIpSource(val *string)
	// Experimental.
	PublicIpSourceInput() *string
	// Experimental.
	PubliclyAdvertisable() interface{}
	// Experimental.
	SetPubliclyAdvertisable(val interface{})
	// Experimental.
	PubliclyAdvertisableInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	SourceIpamPoolId() *string
	// Experimental.
	SetSourceIpamPoolId(val *string)
	// Experimental.
	SourceIpamPoolIdInput() *string
	// Experimental.
	SourceResource() AwsPool_SourceResourcePropertyOutputReference
	// Experimental.
	SourceResourceInput() *AwsPool_SourceResourceProperty
	// Experimental.
	State() *string
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
	Timeouts() AwsPool_TimeoutsPropertyOutputReference
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
	PutSourceResource(value *AwsPool_SourceResourceProperty)
	// Experimental.
	PutTimeouts(value *AwsPool_TimeoutsProperty)
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
	ResetAllocationDefaultNetmaskLength()
	// Experimental.
	ResetAllocationMaxNetmaskLength()
	// Experimental.
	ResetAllocationMinNetmaskLength()
	// Experimental.
	ResetAllocationResourceTags()
	// Experimental.
	ResetAutoImport()
	// Experimental.
	ResetAwsService()
	// Experimental.
	ResetCascade()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLocale()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPublicIpSource()
	// Experimental.
	ResetPubliclyAdvertisable()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSourceIpamPoolId()
	// Experimental.
	ResetSourceResource()
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

// The jsii proxy struct for AwsPool
type jsiiProxy_AwsPool struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsPool) AddressFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) AddressFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) AllocationDefaultNetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationDefaultNetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) AllocationDefaultNetmaskLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationDefaultNetmaskLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) AllocationMaxNetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationMaxNetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) AllocationMaxNetmaskLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationMaxNetmaskLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) AllocationMinNetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationMinNetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) AllocationMinNetmaskLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationMinNetmaskLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) AllocationResourceTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"allocationResourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) AllocationResourceTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"allocationResourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) AutoImport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoImport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) AutoImportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoImportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) AwsService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) AwsServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) Cascade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cascade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) CascadeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cascadeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) IpamScopeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamScopeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) IpamScopeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamScopeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) IpamScopeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamScopeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) Locale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) LocaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) PoolDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"poolDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) PublicIpSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) PublicIpSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) PubliclyAdvertisable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAdvertisable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) PubliclyAdvertisableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAdvertisableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) SourceIpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIpamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) SourceIpamPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIpamPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) SourceResource() AwsPool_SourceResourcePropertyOutputReference {
	var returns AwsPool_SourceResourcePropertyOutputReference
	_jsii_.Get(
		j,
		"sourceResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) SourceResourceInput() *AwsPool_SourceResourceProperty {
	var returns *AwsPool_SourceResourceProperty
	_jsii_.Get(
		j,
		"sourceResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) Timeouts() AwsPool_TimeoutsPropertyOutputReference {
	var returns AwsPool_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPool) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool aws_vpc_ipam_pool} Resource.
// Experimental.
func NewAwsPool(scope constructs.Construct, id *string, config *AwsPoolConfig) AwsPool {
	_init_.Initialize()

	if err := validateNewAwsPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPool{}

	_jsii_.Create(
		"@cdktn/aws-vpc-ipam.AwsPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_ipam_pool aws_vpc_ipam_pool} Resource.
// Experimental.
func NewAwsPool_Override(a AwsPool, scope constructs.Construct, id *string, config *AwsPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc-ipam.AwsPool",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsPool)SetAddressFamily(val *string) {
	if err := j.validateSetAddressFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressFamily",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetAllocationDefaultNetmaskLength(val *float64) {
	if err := j.validateSetAllocationDefaultNetmaskLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationDefaultNetmaskLength",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetAllocationMaxNetmaskLength(val *float64) {
	if err := j.validateSetAllocationMaxNetmaskLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationMaxNetmaskLength",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetAllocationMinNetmaskLength(val *float64) {
	if err := j.validateSetAllocationMinNetmaskLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationMinNetmaskLength",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetAllocationResourceTags(val *map[string]*string) {
	if err := j.validateSetAllocationResourceTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationResourceTags",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetAutoImport(val interface{}) {
	if err := j.validateSetAutoImportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoImport",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetAwsService(val *string) {
	if err := j.validateSetAwsServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsService",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetCascade(val interface{}) {
	if err := j.validateSetCascadeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cascade",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetIpamScopeId(val *string) {
	if err := j.validateSetIpamScopeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipamScopeId",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetLocale(val *string) {
	if err := j.validateSetLocaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locale",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetPublicIpSource(val *string) {
	if err := j.validateSetPublicIpSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpSource",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetPubliclyAdvertisable(val interface{}) {
	if err := j.validateSetPubliclyAdvertisableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAdvertisable",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetSourceIpamPoolId(val *string) {
	if err := j.validateSetSourceIpamPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIpamPoolId",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsPool)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a AwsPool resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsPool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsPool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc-ipam.AwsPool",
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
func AwsPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc-ipam.AwsPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc-ipam.AwsPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsPool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsPool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpc-ipam.AwsPool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-vpc-ipam.AwsPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsPool) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsPool) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPool) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPool) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPool) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPool) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPool) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsPool) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPool) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsPool) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsPool) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsPool) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsPool) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsPool) PutSourceResource(value *AwsPool_SourceResourceProperty) {
	if err := a.validatePutSourceResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourceResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPool) PutTimeouts(value *AwsPool_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPool) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsPool) ResetAllocationDefaultNetmaskLength() {
	_jsii_.InvokeVoid(
		a,
		"resetAllocationDefaultNetmaskLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetAllocationMaxNetmaskLength() {
	_jsii_.InvokeVoid(
		a,
		"resetAllocationMaxNetmaskLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetAllocationMinNetmaskLength() {
	_jsii_.InvokeVoid(
		a,
		"resetAllocationMinNetmaskLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetAllocationResourceTags() {
	_jsii_.InvokeVoid(
		a,
		"resetAllocationResourceTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetAutoImport() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoImport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetAwsService() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsService",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetCascade() {
	_jsii_.InvokeVoid(
		a,
		"resetCascade",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetLocale() {
	_jsii_.InvokeVoid(
		a,
		"resetLocale",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetPublicIpSource() {
	_jsii_.InvokeVoid(
		a,
		"resetPublicIpSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetPubliclyAdvertisable() {
	_jsii_.InvokeVoid(
		a,
		"resetPubliclyAdvertisable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetSourceIpamPoolId() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceIpamPoolId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetSourceResource() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPool) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

