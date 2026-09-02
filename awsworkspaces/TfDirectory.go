package awsworkspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsworkspaces/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsworkspaces/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory aws_workspaces_directory}.
// Experimental.
type TfDirectory interface {
	cdktn.TerraformResource
	// Experimental.
	ActiveDirectoryConfig() TfDirectory_ActiveDirectoryConfigPropertyOutputReference
	// Experimental.
	ActiveDirectoryConfigInput() *TfDirectory_ActiveDirectoryConfigProperty
	// Experimental.
	Alias() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CertificateBasedAuthProperties() TfDirectory_CertificateBasedAuthPropertiesPropertyOutputReference
	// Experimental.
	CertificateBasedAuthPropertiesInput() *TfDirectory_CertificateBasedAuthPropertiesProperty
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
	CustomerUserName() *string
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
	DirectoryName() *string
	// Experimental.
	DirectoryType() *string
	// Experimental.
	DnsIpAddresses() *[]*string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	IamRoleId() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	IpGroupIds() *[]*string
	// Experimental.
	SetIpGroupIds(val *[]*string)
	// Experimental.
	IpGroupIdsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
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
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	RegistrationCode() *string
	// Experimental.
	SamlProperties() TfDirectory_SamlPropertiesPropertyOutputReference
	// Experimental.
	SamlPropertiesInput() *TfDirectory_SamlPropertiesProperty
	// Experimental.
	SelfServicePermissions() TfDirectory_SelfServicePermissionsPropertyOutputReference
	// Experimental.
	SelfServicePermissionsInput() *TfDirectory_SelfServicePermissionsProperty
	// Experimental.
	SubnetIds() *[]*string
	// Experimental.
	SetSubnetIds(val *[]*string)
	// Experimental.
	SubnetIdsInput() *[]*string
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
	Tenancy() *string
	// Experimental.
	SetTenancy(val *string)
	// Experimental.
	TenancyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	UserIdentityType() *string
	// Experimental.
	SetUserIdentityType(val *string)
	// Experimental.
	UserIdentityTypeInput() *string
	// Experimental.
	WorkspaceAccessProperties() TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference
	// Experimental.
	WorkspaceAccessPropertiesInput() *TfDirectory_WorkspaceAccessPropertiesProperty
	// Experimental.
	WorkspaceCreationProperties() TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference
	// Experimental.
	WorkspaceCreationPropertiesInput() *TfDirectory_WorkspaceCreationPropertiesProperty
	// Experimental.
	WorkspaceDirectoryDescription() *string
	// Experimental.
	SetWorkspaceDirectoryDescription(val *string)
	// Experimental.
	WorkspaceDirectoryDescriptionInput() *string
	// Experimental.
	WorkspaceDirectoryName() *string
	// Experimental.
	SetWorkspaceDirectoryName(val *string)
	// Experimental.
	WorkspaceDirectoryNameInput() *string
	// Experimental.
	WorkspaceSecurityGroupId() *string
	// Experimental.
	WorkspaceType() *string
	// Experimental.
	SetWorkspaceType(val *string)
	// Experimental.
	WorkspaceTypeInput() *string
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
	PutActiveDirectoryConfig(value *TfDirectory_ActiveDirectoryConfigProperty)
	// Experimental.
	PutCertificateBasedAuthProperties(value *TfDirectory_CertificateBasedAuthPropertiesProperty)
	// Experimental.
	PutSamlProperties(value *TfDirectory_SamlPropertiesProperty)
	// Experimental.
	PutSelfServicePermissions(value *TfDirectory_SelfServicePermissionsProperty)
	// Experimental.
	PutWorkspaceAccessProperties(value *TfDirectory_WorkspaceAccessPropertiesProperty)
	// Experimental.
	PutWorkspaceCreationProperties(value *TfDirectory_WorkspaceCreationPropertiesProperty)
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
	ResetActiveDirectoryConfig()
	// Experimental.
	ResetCertificateBasedAuthProperties()
	// Experimental.
	ResetDirectoryId()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIpGroupIds()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSamlProperties()
	// Experimental.
	ResetSelfServicePermissions()
	// Experimental.
	ResetSubnetIds()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTenancy()
	// Experimental.
	ResetUserIdentityType()
	// Experimental.
	ResetWorkspaceAccessProperties()
	// Experimental.
	ResetWorkspaceCreationProperties()
	// Experimental.
	ResetWorkspaceDirectoryDescription()
	// Experimental.
	ResetWorkspaceDirectoryName()
	// Experimental.
	ResetWorkspaceType()
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

// The jsii proxy struct for TfDirectory
type jsiiProxy_TfDirectory struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfDirectory) ActiveDirectoryConfig() TfDirectory_ActiveDirectoryConfigPropertyOutputReference {
	var returns TfDirectory_ActiveDirectoryConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"activeDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) ActiveDirectoryConfigInput() *TfDirectory_ActiveDirectoryConfigProperty {
	var returns *TfDirectory_ActiveDirectoryConfigProperty
	_jsii_.Get(
		j,
		"activeDirectoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) CertificateBasedAuthProperties() TfDirectory_CertificateBasedAuthPropertiesPropertyOutputReference {
	var returns TfDirectory_CertificateBasedAuthPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"certificateBasedAuthProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) CertificateBasedAuthPropertiesInput() *TfDirectory_CertificateBasedAuthPropertiesProperty {
	var returns *TfDirectory_CertificateBasedAuthPropertiesProperty
	_jsii_.Get(
		j,
		"certificateBasedAuthPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) CustomerUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) DirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) DirectoryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) DnsIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) IamRoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) IpGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) IpGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) RegistrationCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registrationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) SamlProperties() TfDirectory_SamlPropertiesPropertyOutputReference {
	var returns TfDirectory_SamlPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"samlProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) SamlPropertiesInput() *TfDirectory_SamlPropertiesProperty {
	var returns *TfDirectory_SamlPropertiesProperty
	_jsii_.Get(
		j,
		"samlPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) SelfServicePermissions() TfDirectory_SelfServicePermissionsPropertyOutputReference {
	var returns TfDirectory_SelfServicePermissionsPropertyOutputReference
	_jsii_.Get(
		j,
		"selfServicePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) SelfServicePermissionsInput() *TfDirectory_SelfServicePermissionsProperty {
	var returns *TfDirectory_SelfServicePermissionsProperty
	_jsii_.Get(
		j,
		"selfServicePermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) Tenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) TenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) UserIdentityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdentityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) UserIdentityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdentityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) WorkspaceAccessProperties() TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference {
	var returns TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"workspaceAccessProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) WorkspaceAccessPropertiesInput() *TfDirectory_WorkspaceAccessPropertiesProperty {
	var returns *TfDirectory_WorkspaceAccessPropertiesProperty
	_jsii_.Get(
		j,
		"workspaceAccessPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) WorkspaceCreationProperties() TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference {
	var returns TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"workspaceCreationProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) WorkspaceCreationPropertiesInput() *TfDirectory_WorkspaceCreationPropertiesProperty {
	var returns *TfDirectory_WorkspaceCreationPropertiesProperty
	_jsii_.Get(
		j,
		"workspaceCreationPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) WorkspaceDirectoryDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceDirectoryDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) WorkspaceDirectoryDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceDirectoryDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) WorkspaceDirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceDirectoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) WorkspaceDirectoryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceDirectoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) WorkspaceSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) WorkspaceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory) WorkspaceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory aws_workspaces_directory} Resource.
// Experimental.
func NewTfDirectory(scope constructs.Construct, id *string, config *TfDirectoryConfig) TfDirectory {
	_init_.Initialize()

	if err := validateNewTfDirectoryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDirectory{}

	_jsii_.Create(
		"@cdktn/aws-workspaces.TfDirectory",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory aws_workspaces_directory} Resource.
// Experimental.
func NewTfDirectory_Override(t TfDirectory, scope constructs.Construct, id *string, config *TfDirectoryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-workspaces.TfDirectory",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfDirectory)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetDirectoryId(val *string) {
	if err := j.validateSetDirectoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetIpGroupIds(val *[]*string) {
	if err := j.validateSetIpGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipGroupIds",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetTenancy(val *string) {
	if err := j.validateSetTenancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenancy",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetUserIdentityType(val *string) {
	if err := j.validateSetUserIdentityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userIdentityType",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetWorkspaceDirectoryDescription(val *string) {
	if err := j.validateSetWorkspaceDirectoryDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceDirectoryDescription",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetWorkspaceDirectoryName(val *string) {
	if err := j.validateSetWorkspaceDirectoryNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceDirectoryName",
		val,
	)
}

func (j *jsiiProxy_TfDirectory)SetWorkspaceType(val *string) {
	if err := j.validateSetWorkspaceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceType",
		val,
	)
}

// Generates CDKTN code for importing a TfDirectory resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfDirectory_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfDirectory_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-workspaces.TfDirectory",
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
func TfDirectory_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDirectory_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-workspaces.TfDirectory",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfDirectory_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDirectory_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-workspaces.TfDirectory",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfDirectory_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDirectory_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-workspaces.TfDirectory",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfDirectory_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-workspaces.TfDirectory",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfDirectory) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfDirectory) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfDirectory) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDirectory) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDirectory) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDirectory) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDirectory) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDirectory) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDirectory) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDirectory) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDirectory) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDirectory) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDirectory) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfDirectory) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDirectory) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfDirectory) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfDirectory) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfDirectory) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfDirectory) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfDirectory) PutActiveDirectoryConfig(value *TfDirectory_ActiveDirectoryConfigProperty) {
	if err := t.validatePutActiveDirectoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putActiveDirectoryConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDirectory) PutCertificateBasedAuthProperties(value *TfDirectory_CertificateBasedAuthPropertiesProperty) {
	if err := t.validatePutCertificateBasedAuthPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCertificateBasedAuthProperties",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDirectory) PutSamlProperties(value *TfDirectory_SamlPropertiesProperty) {
	if err := t.validatePutSamlPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSamlProperties",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDirectory) PutSelfServicePermissions(value *TfDirectory_SelfServicePermissionsProperty) {
	if err := t.validatePutSelfServicePermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSelfServicePermissions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDirectory) PutWorkspaceAccessProperties(value *TfDirectory_WorkspaceAccessPropertiesProperty) {
	if err := t.validatePutWorkspaceAccessPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWorkspaceAccessProperties",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDirectory) PutWorkspaceCreationProperties(value *TfDirectory_WorkspaceCreationPropertiesProperty) {
	if err := t.validatePutWorkspaceCreationPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWorkspaceCreationProperties",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDirectory) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfDirectory) ResetActiveDirectoryConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetActiveDirectoryConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetCertificateBasedAuthProperties() {
	_jsii_.InvokeVoid(
		t,
		"resetCertificateBasedAuthProperties",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetDirectoryId() {
	_jsii_.InvokeVoid(
		t,
		"resetDirectoryId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetIpGroupIds() {
	_jsii_.InvokeVoid(
		t,
		"resetIpGroupIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetSamlProperties() {
	_jsii_.InvokeVoid(
		t,
		"resetSamlProperties",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetSelfServicePermissions() {
	_jsii_.InvokeVoid(
		t,
		"resetSelfServicePermissions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		t,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetTenancy() {
	_jsii_.InvokeVoid(
		t,
		"resetTenancy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetUserIdentityType() {
	_jsii_.InvokeVoid(
		t,
		"resetUserIdentityType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetWorkspaceAccessProperties() {
	_jsii_.InvokeVoid(
		t,
		"resetWorkspaceAccessProperties",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetWorkspaceCreationProperties() {
	_jsii_.InvokeVoid(
		t,
		"resetWorkspaceCreationProperties",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetWorkspaceDirectoryDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetWorkspaceDirectoryDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetWorkspaceDirectoryName() {
	_jsii_.InvokeVoid(
		t,
		"resetWorkspaceDirectoryName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) ResetWorkspaceType() {
	_jsii_.InvokeVoid(
		t,
		"resetWorkspaceType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDirectory) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDirectory) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDirectory) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDirectory) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDirectory) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDirectory) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

