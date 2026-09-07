package codebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/codebuild/jsii"

	"github.com/cdktn-io/cdktn-aws-go/codebuild/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsProject_ArtifactsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ArtifactIdentifier() *string
	// Experimental.
	SetArtifactIdentifier(val *string)
	// Experimental.
	ArtifactIdentifierInput() *string
	// Experimental.
	BucketOwnerAccess() *string
	// Experimental.
	SetBucketOwnerAccess(val *string)
	// Experimental.
	BucketOwnerAccessInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EncryptionDisabled() interface{}
	// Experimental.
	SetEncryptionDisabled(val interface{})
	// Experimental.
	EncryptionDisabledInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsProject_ArtifactsProperty
	// Experimental.
	SetInternalValue(val *AwsProject_ArtifactsProperty)
	// Experimental.
	Location() *string
	// Experimental.
	SetLocation(val *string)
	// Experimental.
	LocationInput() *string
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NamespaceType() *string
	// Experimental.
	SetNamespaceType(val *string)
	// Experimental.
	NamespaceTypeInput() *string
	// Experimental.
	OverrideArtifactName() interface{}
	// Experimental.
	SetOverrideArtifactName(val interface{})
	// Experimental.
	OverrideArtifactNameInput() interface{}
	// Experimental.
	Packaging() *string
	// Experimental.
	SetPackaging(val *string)
	// Experimental.
	PackagingInput() *string
	// Experimental.
	Path() *string
	// Experimental.
	SetPath(val *string)
	// Experimental.
	PathInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	ResetArtifactIdentifier()
	// Experimental.
	ResetBucketOwnerAccess()
	// Experimental.
	ResetEncryptionDisabled()
	// Experimental.
	ResetLocation()
	// Experimental.
	ResetName()
	// Experimental.
	ResetNamespaceType()
	// Experimental.
	ResetOverrideArtifactName()
	// Experimental.
	ResetPackaging()
	// Experimental.
	ResetPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsProject_ArtifactsPropertyOutputReference
type jsiiProxy_AwsProject_ArtifactsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) ArtifactIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) ArtifactIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) BucketOwnerAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketOwnerAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) BucketOwnerAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketOwnerAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) EncryptionDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) EncryptionDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) InternalValue() *AwsProject_ArtifactsProperty {
	var returns *AwsProject_ArtifactsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) NamespaceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) NamespaceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) OverrideArtifactName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideArtifactName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) OverrideArtifactNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideArtifactNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) Packaging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packaging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) PackagingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packagingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsProject_ArtifactsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsProject_ArtifactsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsProject_ArtifactsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsProject_ArtifactsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codebuild.AwsProject.ArtifactsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsProject_ArtifactsPropertyOutputReference_Override(a AwsProject_ArtifactsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codebuild.AwsProject.ArtifactsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference)SetArtifactIdentifier(val *string) {
	if err := j.validateSetArtifactIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"artifactIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference)SetBucketOwnerAccess(val *string) {
	if err := j.validateSetBucketOwnerAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketOwnerAccess",
		val,
	)
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference)SetEncryptionDisabled(val interface{}) {
	if err := j.validateSetEncryptionDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionDisabled",
		val,
	)
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference)SetInternalValue(val *AwsProject_ArtifactsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference)SetNamespaceType(val *string) {
	if err := j.validateSetNamespaceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespaceType",
		val,
	)
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference)SetOverrideArtifactName(val interface{}) {
	if err := j.validateSetOverrideArtifactNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideArtifactName",
		val,
	)
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference)SetPackaging(val *string) {
	if err := j.validateSetPackagingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packaging",
		val,
	)
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) ResetArtifactIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetArtifactIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) ResetBucketOwnerAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetBucketOwnerAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) ResetEncryptionDisabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptionDisabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) ResetNamespaceType() {
	_jsii_.InvokeVoid(
		a,
		"resetNamespaceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) ResetOverrideArtifactName() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideArtifactName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) ResetPackaging() {
	_jsii_.InvokeVoid(
		a,
		"resetPackaging",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		a,
		"resetPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProject_ArtifactsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

