// cdk-terrain bindings for the Route 53 Domains group of terraform-provider-aws 6.62.0
package awsroute53domains

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsroute53domains/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsroute53domains/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain aws_route53domains_registered_domain}.
// Experimental.
type AwsRoute53DomainsRegisteredDomain interface {
	cdktn.TerraformResource
	// Experimental.
	AbuseContactEmail() *string
	// Experimental.
	AbuseContactPhone() *string
	// Experimental.
	AdminContact() AwsRoute53DomainsRegisteredDomain_AdminContactPropertyOutputReference
	// Experimental.
	AdminContactInput() *AwsRoute53DomainsRegisteredDomain_AdminContactProperty
	// Experimental.
	AdminPrivacy() interface{}
	// Experimental.
	SetAdminPrivacy(val interface{})
	// Experimental.
	AdminPrivacyInput() interface{}
	// Experimental.
	AutoRenew() interface{}
	// Experimental.
	SetAutoRenew(val interface{})
	// Experimental.
	AutoRenewInput() interface{}
	// Experimental.
	BillingContact() AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference
	// Experimental.
	BillingContactInput() *AwsRoute53DomainsRegisteredDomain_BillingContactProperty
	// Experimental.
	BillingPrivacy() interface{}
	// Experimental.
	SetBillingPrivacy(val interface{})
	// Experimental.
	BillingPrivacyInput() interface{}
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
	CreationDate() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DomainName() *string
	// Experimental.
	SetDomainName(val *string)
	// Experimental.
	DomainNameInput() *string
	// Experimental.
	ExpirationDate() *string
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
	NameServer() AwsRoute53DomainsRegisteredDomain_NameServerPropertyList
	// Experimental.
	NameServerInput() interface{}
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
	RegistrantContact() AwsRoute53DomainsRegisteredDomain_RegistrantContactPropertyOutputReference
	// Experimental.
	RegistrantContactInput() *AwsRoute53DomainsRegisteredDomain_RegistrantContactProperty
	// Experimental.
	RegistrantPrivacy() interface{}
	// Experimental.
	SetRegistrantPrivacy(val interface{})
	// Experimental.
	RegistrantPrivacyInput() interface{}
	// Experimental.
	RegistrarName() *string
	// Experimental.
	RegistrarUrl() *string
	// Experimental.
	Reseller() *string
	// Experimental.
	StatusList() *[]*string
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
	TechContact() AwsRoute53DomainsRegisteredDomain_TechContactPropertyOutputReference
	// Experimental.
	TechContactInput() *AwsRoute53DomainsRegisteredDomain_TechContactProperty
	// Experimental.
	TechPrivacy() interface{}
	// Experimental.
	SetTechPrivacy(val interface{})
	// Experimental.
	TechPrivacyInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsRoute53DomainsRegisteredDomain_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TransferLock() interface{}
	// Experimental.
	SetTransferLock(val interface{})
	// Experimental.
	TransferLockInput() interface{}
	// Experimental.
	UpdatedDate() *string
	// Experimental.
	WhoisServer() *string
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
	PutAdminContact(value *AwsRoute53DomainsRegisteredDomain_AdminContactProperty)
	// Experimental.
	PutBillingContact(value *AwsRoute53DomainsRegisteredDomain_BillingContactProperty)
	// Experimental.
	PutNameServer(value interface{})
	// Experimental.
	PutRegistrantContact(value *AwsRoute53DomainsRegisteredDomain_RegistrantContactProperty)
	// Experimental.
	PutTechContact(value *AwsRoute53DomainsRegisteredDomain_TechContactProperty)
	// Experimental.
	PutTimeouts(value *AwsRoute53DomainsRegisteredDomain_TimeoutsProperty)
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
	ResetAdminContact()
	// Experimental.
	ResetAdminPrivacy()
	// Experimental.
	ResetAutoRenew()
	// Experimental.
	ResetBillingContact()
	// Experimental.
	ResetBillingPrivacy()
	// Experimental.
	ResetId()
	// Experimental.
	ResetNameServer()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegistrantContact()
	// Experimental.
	ResetRegistrantPrivacy()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTechContact()
	// Experimental.
	ResetTechPrivacy()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTransferLock()
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

// The jsii proxy struct for AwsRoute53DomainsRegisteredDomain
type jsiiProxy_AwsRoute53DomainsRegisteredDomain struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) AbuseContactEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"abuseContactEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) AbuseContactPhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"abuseContactPhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) AdminContact() AwsRoute53DomainsRegisteredDomain_AdminContactPropertyOutputReference {
	var returns AwsRoute53DomainsRegisteredDomain_AdminContactPropertyOutputReference
	_jsii_.Get(
		j,
		"adminContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) AdminContactInput() *AwsRoute53DomainsRegisteredDomain_AdminContactProperty {
	var returns *AwsRoute53DomainsRegisteredDomain_AdminContactProperty
	_jsii_.Get(
		j,
		"adminContactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) AdminPrivacy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminPrivacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) AdminPrivacyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminPrivacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) AutoRenew() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRenew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) AutoRenewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRenewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) BillingContact() AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference {
	var returns AwsRoute53DomainsRegisteredDomain_BillingContactPropertyOutputReference
	_jsii_.Get(
		j,
		"billingContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) BillingContactInput() *AwsRoute53DomainsRegisteredDomain_BillingContactProperty {
	var returns *AwsRoute53DomainsRegisteredDomain_BillingContactProperty
	_jsii_.Get(
		j,
		"billingContactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) BillingPrivacy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"billingPrivacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) BillingPrivacyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"billingPrivacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) CreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ExpirationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) NameServer() AwsRoute53DomainsRegisteredDomain_NameServerPropertyList {
	var returns AwsRoute53DomainsRegisteredDomain_NameServerPropertyList
	_jsii_.Get(
		j,
		"nameServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) NameServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nameServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) RegistrantContact() AwsRoute53DomainsRegisteredDomain_RegistrantContactPropertyOutputReference {
	var returns AwsRoute53DomainsRegisteredDomain_RegistrantContactPropertyOutputReference
	_jsii_.Get(
		j,
		"registrantContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) RegistrantContactInput() *AwsRoute53DomainsRegisteredDomain_RegistrantContactProperty {
	var returns *AwsRoute53DomainsRegisteredDomain_RegistrantContactProperty
	_jsii_.Get(
		j,
		"registrantContactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) RegistrantPrivacy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registrantPrivacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) RegistrantPrivacyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registrantPrivacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) RegistrarName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registrarName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) RegistrarUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registrarUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) Reseller() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reseller",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) StatusList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statusList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) TechContact() AwsRoute53DomainsRegisteredDomain_TechContactPropertyOutputReference {
	var returns AwsRoute53DomainsRegisteredDomain_TechContactPropertyOutputReference
	_jsii_.Get(
		j,
		"techContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) TechContactInput() *AwsRoute53DomainsRegisteredDomain_TechContactProperty {
	var returns *AwsRoute53DomainsRegisteredDomain_TechContactProperty
	_jsii_.Get(
		j,
		"techContactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) TechPrivacy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"techPrivacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) TechPrivacyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"techPrivacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) Timeouts() AwsRoute53DomainsRegisteredDomain_TimeoutsPropertyOutputReference {
	var returns AwsRoute53DomainsRegisteredDomain_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) TransferLock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transferLock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) TransferLockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transferLockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) UpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain) WhoisServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoisServer",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain aws_route53domains_registered_domain} Resource.
// Experimental.
func NewAwsRoute53DomainsRegisteredDomain(scope constructs.Construct, id *string, config *AwsRoute53DomainsRegisteredDomainConfig) AwsRoute53DomainsRegisteredDomain {
	_init_.Initialize()

	if err := validateNewAwsRoute53DomainsRegisteredDomainParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRoute53DomainsRegisteredDomain{}

	_jsii_.Create(
		"@cdktn/aws-route-53-domains.AwsRoute53DomainsRegisteredDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain aws_route53domains_registered_domain} Resource.
// Experimental.
func NewAwsRoute53DomainsRegisteredDomain_Override(a AwsRoute53DomainsRegisteredDomain, scope constructs.Construct, id *string, config *AwsRoute53DomainsRegisteredDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53-domains.AwsRoute53DomainsRegisteredDomain",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain)SetAdminPrivacy(val interface{}) {
	if err := j.validateSetAdminPrivacyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPrivacy",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain)SetAutoRenew(val interface{}) {
	if err := j.validateSetAutoRenewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRenew",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain)SetBillingPrivacy(val interface{}) {
	if err := j.validateSetBillingPrivacyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingPrivacy",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain)SetRegistrantPrivacy(val interface{}) {
	if err := j.validateSetRegistrantPrivacyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registrantPrivacy",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain)SetTechPrivacy(val interface{}) {
	if err := j.validateSetTechPrivacyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"techPrivacy",
		val,
	)
}

func (j *jsiiProxy_AwsRoute53DomainsRegisteredDomain)SetTransferLock(val interface{}) {
	if err := j.validateSetTransferLockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transferLock",
		val,
	)
}

// Generates CDKTN code for importing a AwsRoute53DomainsRegisteredDomain resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsRoute53DomainsRegisteredDomain_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsRoute53DomainsRegisteredDomain_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53-domains.AwsRoute53DomainsRegisteredDomain",
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
func AwsRoute53DomainsRegisteredDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsRoute53DomainsRegisteredDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53-domains.AwsRoute53DomainsRegisteredDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsRoute53DomainsRegisteredDomain_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsRoute53DomainsRegisteredDomain_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53-domains.AwsRoute53DomainsRegisteredDomain",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsRoute53DomainsRegisteredDomain_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsRoute53DomainsRegisteredDomain_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-route-53-domains.AwsRoute53DomainsRegisteredDomain",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsRoute53DomainsRegisteredDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-route-53-domains.AwsRoute53DomainsRegisteredDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) PutAdminContact(value *AwsRoute53DomainsRegisteredDomain_AdminContactProperty) {
	if err := a.validatePutAdminContactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdminContact",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) PutBillingContact(value *AwsRoute53DomainsRegisteredDomain_BillingContactProperty) {
	if err := a.validatePutBillingContactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBillingContact",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) PutNameServer(value interface{}) {
	if err := a.validatePutNameServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNameServer",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) PutRegistrantContact(value *AwsRoute53DomainsRegisteredDomain_RegistrantContactProperty) {
	if err := a.validatePutRegistrantContactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRegistrantContact",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) PutTechContact(value *AwsRoute53DomainsRegisteredDomain_TechContactProperty) {
	if err := a.validatePutTechContactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTechContact",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) PutTimeouts(value *AwsRoute53DomainsRegisteredDomain_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ResetAdminContact() {
	_jsii_.InvokeVoid(
		a,
		"resetAdminContact",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ResetAdminPrivacy() {
	_jsii_.InvokeVoid(
		a,
		"resetAdminPrivacy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ResetAutoRenew() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoRenew",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ResetBillingContact() {
	_jsii_.InvokeVoid(
		a,
		"resetBillingContact",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ResetBillingPrivacy() {
	_jsii_.InvokeVoid(
		a,
		"resetBillingPrivacy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ResetNameServer() {
	_jsii_.InvokeVoid(
		a,
		"resetNameServer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ResetRegistrantContact() {
	_jsii_.InvokeVoid(
		a,
		"resetRegistrantContact",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ResetRegistrantPrivacy() {
	_jsii_.InvokeVoid(
		a,
		"resetRegistrantPrivacy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ResetTechContact() {
	_jsii_.InvokeVoid(
		a,
		"resetTechContact",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ResetTechPrivacy() {
	_jsii_.InvokeVoid(
		a,
		"resetTechPrivacy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ResetTransferLock() {
	_jsii_.InvokeVoid(
		a,
		"resetTransferLock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute53DomainsRegisteredDomain) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

