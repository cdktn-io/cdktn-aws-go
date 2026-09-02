package awsworkmail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsworkmail/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsworkmail/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user aws_workmail_user}.
// Experimental.
type TfUser interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	City() *string
	// Experimental.
	SetCity(val *string)
	// Experimental.
	CityInput() *string
	// Experimental.
	Company() *string
	// Experimental.
	SetCompany(val *string)
	// Experimental.
	CompanyInput() *string
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
	Country() *string
	// Experimental.
	SetCountry(val *string)
	// Experimental.
	CountryInput() *string
	// Experimental.
	Department() *string
	// Experimental.
	SetDepartment(val *string)
	// Experimental.
	DepartmentInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DisabledDate() *string
	// Experimental.
	DisplayName() *string
	// Experimental.
	SetDisplayName(val *string)
	// Experimental.
	DisplayNameInput() *string
	// Experimental.
	Email() *string
	// Experimental.
	SetEmail(val *string)
	// Experimental.
	EmailInput() *string
	// Experimental.
	EnabledDate() *string
	// Experimental.
	FirstName() *string
	// Experimental.
	SetFirstName(val *string)
	// Experimental.
	FirstNameInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HiddenFromGlobalAddressList() interface{}
	// Experimental.
	SetHiddenFromGlobalAddressList(val interface{})
	// Experimental.
	HiddenFromGlobalAddressListInput() interface{}
	// Experimental.
	IdentityProviderIdentityStoreId() *string
	// Experimental.
	IdentityProviderUserId() *string
	// Experimental.
	SetIdentityProviderUserId(val *string)
	// Experimental.
	IdentityProviderUserIdInput() *string
	// Experimental.
	Initials() *string
	// Experimental.
	SetInitials(val *string)
	// Experimental.
	InitialsInput() *string
	// Experimental.
	JobTitle() *string
	// Experimental.
	SetJobTitle(val *string)
	// Experimental.
	JobTitleInput() *string
	// Experimental.
	LastName() *string
	// Experimental.
	SetLastName(val *string)
	// Experimental.
	LastNameInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MailboxDeprovisionedDate() *string
	// Experimental.
	MailboxProvisionedDate() *string
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
	Office() *string
	// Experimental.
	SetOffice(val *string)
	// Experimental.
	OfficeInput() *string
	// Experimental.
	OrganizationId() *string
	// Experimental.
	SetOrganizationId(val *string)
	// Experimental.
	OrganizationIdInput() *string
	// Experimental.
	Password() *string
	// Experimental.
	SetPassword(val *string)
	// Experimental.
	PasswordInput() *string
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
	State() *string
	// Experimental.
	Street() *string
	// Experimental.
	SetStreet(val *string)
	// Experimental.
	StreetInput() *string
	// Experimental.
	Telephone() *string
	// Experimental.
	SetTelephone(val *string)
	// Experimental.
	TelephoneInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	UserId() *string
	// Experimental.
	UserRole() *string
	// Experimental.
	SetUserRole(val *string)
	// Experimental.
	UserRoleInput() *string
	// Experimental.
	ZipCode() *string
	// Experimental.
	SetZipCode(val *string)
	// Experimental.
	ZipCodeInput() *string
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
	ResetCity()
	// Experimental.
	ResetCompany()
	// Experimental.
	ResetCountry()
	// Experimental.
	ResetDepartment()
	// Experimental.
	ResetFirstName()
	// Experimental.
	ResetHiddenFromGlobalAddressList()
	// Experimental.
	ResetIdentityProviderUserId()
	// Experimental.
	ResetInitials()
	// Experimental.
	ResetJobTitle()
	// Experimental.
	ResetLastName()
	// Experimental.
	ResetOffice()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPassword()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetStreet()
	// Experimental.
	ResetTelephone()
	// Experimental.
	ResetUserRole()
	// Experimental.
	ResetZipCode()
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

// The jsii proxy struct for TfUser
type jsiiProxy_TfUser struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfUser) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) CityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Company() *string {
	var returns *string
	_jsii_.Get(
		j,
		"company",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) CompanyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"companyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Country() *string {
	var returns *string
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) CountryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Department() *string {
	var returns *string
	_jsii_.Get(
		j,
		"department",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) DepartmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"departmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) DisabledDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disabledDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) EnabledDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enabledDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) FirstNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) HiddenFromGlobalAddressList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hiddenFromGlobalAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) HiddenFromGlobalAddressListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hiddenFromGlobalAddressListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) IdentityProviderIdentityStoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderIdentityStoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) IdentityProviderUserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderUserId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) IdentityProviderUserIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderUserIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Initials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) InitialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) JobTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) JobTitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) LastNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) MailboxDeprovisionedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailboxDeprovisionedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) MailboxProvisionedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mailboxProvisionedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Office() *string {
	var returns *string
	_jsii_.Get(
		j,
		"office",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) OfficeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"officeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) OrganizationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) OrganizationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Street() *string {
	var returns *string
	_jsii_.Get(
		j,
		"street",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) StreetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) Telephone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"telephone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) TelephoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"telephoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) UserRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) UserRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) ZipCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser) ZipCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipCodeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user aws_workmail_user} Resource.
// Experimental.
func NewTfUser(scope constructs.Construct, id *string, config *TfUserConfig) TfUser {
	_init_.Initialize()

	if err := validateNewTfUserParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfUser{}

	_jsii_.Create(
		"@cdktn/aws-workmail.TfUser",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user aws_workmail_user} Resource.
// Experimental.
func NewTfUser_Override(t TfUser, scope constructs.Construct, id *string, config *TfUserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-workmail.TfUser",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfUser)SetCity(val *string) {
	if err := j.validateSetCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"city",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetCompany(val *string) {
	if err := j.validateSetCompanyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"company",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetCountry(val *string) {
	if err := j.validateSetCountryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"country",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetDepartment(val *string) {
	if err := j.validateSetDepartmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"department",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetFirstName(val *string) {
	if err := j.validateSetFirstNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstName",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetHiddenFromGlobalAddressList(val interface{}) {
	if err := j.validateSetHiddenFromGlobalAddressListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hiddenFromGlobalAddressList",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetIdentityProviderUserId(val *string) {
	if err := j.validateSetIdentityProviderUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProviderUserId",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetInitials(val *string) {
	if err := j.validateSetInitialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initials",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetJobTitle(val *string) {
	if err := j.validateSetJobTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobTitle",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetLastName(val *string) {
	if err := j.validateSetLastNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastName",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetOffice(val *string) {
	if err := j.validateSetOfficeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"office",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetOrganizationId(val *string) {
	if err := j.validateSetOrganizationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationId",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetStreet(val *string) {
	if err := j.validateSetStreetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"street",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetTelephone(val *string) {
	if err := j.validateSetTelephoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"telephone",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetUserRole(val *string) {
	if err := j.validateSetUserRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userRole",
		val,
	)
}

func (j *jsiiProxy_TfUser)SetZipCode(val *string) {
	if err := j.validateSetZipCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zipCode",
		val,
	)
}

// Generates CDKTN code for importing a TfUser resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfUser_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfUser_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-workmail.TfUser",
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
func TfUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-workmail.TfUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfUser_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfUser_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-workmail.TfUser",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfUser_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfUser_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-workmail.TfUser",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfUser_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-workmail.TfUser",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfUser) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfUser) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfUser) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfUser) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUser) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfUser) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfUser) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfUser) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfUser) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfUser) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfUser) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfUser) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUser) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfUser) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUser) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfUser) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfUser) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfUser) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfUser) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfUser) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfUser) ResetCity() {
	_jsii_.InvokeVoid(
		t,
		"resetCity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) ResetCompany() {
	_jsii_.InvokeVoid(
		t,
		"resetCompany",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) ResetCountry() {
	_jsii_.InvokeVoid(
		t,
		"resetCountry",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) ResetDepartment() {
	_jsii_.InvokeVoid(
		t,
		"resetDepartment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) ResetFirstName() {
	_jsii_.InvokeVoid(
		t,
		"resetFirstName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) ResetHiddenFromGlobalAddressList() {
	_jsii_.InvokeVoid(
		t,
		"resetHiddenFromGlobalAddressList",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) ResetIdentityProviderUserId() {
	_jsii_.InvokeVoid(
		t,
		"resetIdentityProviderUserId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) ResetInitials() {
	_jsii_.InvokeVoid(
		t,
		"resetInitials",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) ResetJobTitle() {
	_jsii_.InvokeVoid(
		t,
		"resetJobTitle",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) ResetLastName() {
	_jsii_.InvokeVoid(
		t,
		"resetLastName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) ResetOffice() {
	_jsii_.InvokeVoid(
		t,
		"resetOffice",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) ResetPassword() {
	_jsii_.InvokeVoid(
		t,
		"resetPassword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) ResetStreet() {
	_jsii_.InvokeVoid(
		t,
		"resetStreet",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) ResetTelephone() {
	_jsii_.InvokeVoid(
		t,
		"resetTelephone",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) ResetUserRole() {
	_jsii_.InvokeVoid(
		t,
		"resetUserRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) ResetZipCode() {
	_jsii_.InvokeVoid(
		t,
		"resetZipCode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUser) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUser) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUser) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUser) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUser) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUser) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

