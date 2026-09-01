package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGlueJob_SourceControlDetailsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthStrategy() *string
	// Experimental.
	SetAuthStrategy(val *string)
	// Experimental.
	AuthStrategyInput() *string
	// Experimental.
	AuthToken() *string
	// Experimental.
	SetAuthToken(val *string)
	// Experimental.
	AuthTokenInput() *string
	// Experimental.
	Branch() *string
	// Experimental.
	SetBranch(val *string)
	// Experimental.
	BranchInput() *string
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
	Folder() *string
	// Experimental.
	SetFolder(val *string)
	// Experimental.
	FolderInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsGlueJob_SourceControlDetailsProperty
	// Experimental.
	SetInternalValue(val *AwsGlueJob_SourceControlDetailsProperty)
	// Experimental.
	LastCommitId() *string
	// Experimental.
	SetLastCommitId(val *string)
	// Experimental.
	LastCommitIdInput() *string
	// Experimental.
	Owner() *string
	// Experimental.
	SetOwner(val *string)
	// Experimental.
	OwnerInput() *string
	// Experimental.
	Provider() *string
	// Experimental.
	SetProvider(val *string)
	// Experimental.
	ProviderInput() *string
	// Experimental.
	Repository() *string
	// Experimental.
	SetRepository(val *string)
	// Experimental.
	RepositoryInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetAuthStrategy()
	// Experimental.
	ResetAuthToken()
	// Experimental.
	ResetBranch()
	// Experimental.
	ResetFolder()
	// Experimental.
	ResetLastCommitId()
	// Experimental.
	ResetOwner()
	// Experimental.
	ResetProvider()
	// Experimental.
	ResetRepository()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsGlueJob_SourceControlDetailsPropertyOutputReference
type jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) AuthStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) AuthStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) AuthToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) AuthTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) Branch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) BranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) Folder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) FolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) InternalValue() *AwsGlueJob_SourceControlDetailsProperty {
	var returns *AwsGlueJob_SourceControlDetailsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) LastCommitId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastCommitId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) LastCommitIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastCommitIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) Provider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) ProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) Repository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) RepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGlueJob_SourceControlDetailsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsGlueJob_SourceControlDetailsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGlueJob_SourceControlDetailsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsGlueJob.SourceControlDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGlueJob_SourceControlDetailsPropertyOutputReference_Override(a AwsGlueJob_SourceControlDetailsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsGlueJob.SourceControlDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference)SetAuthStrategy(val *string) {
	if err := j.validateSetAuthStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authStrategy",
		val,
	)
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference)SetAuthToken(val *string) {
	if err := j.validateSetAuthTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authToken",
		val,
	)
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference)SetBranch(val *string) {
	if err := j.validateSetBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branch",
		val,
	)
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference)SetFolder(val *string) {
	if err := j.validateSetFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"folder",
		val,
	)
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference)SetInternalValue(val *AwsGlueJob_SourceControlDetailsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference)SetLastCommitId(val *string) {
	if err := j.validateSetLastCommitIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastCommitId",
		val,
	)
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference)SetProvider(val *string) {
	if err := j.validateSetProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference)SetRepository(val *string) {
	if err := j.validateSetRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repository",
		val,
	)
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) ResetAuthStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) ResetAuthToken() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) ResetBranch() {
	_jsii_.InvokeVoid(
		a,
		"resetBranch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) ResetFolder() {
	_jsii_.InvokeVoid(
		a,
		"resetFolder",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) ResetLastCommitId() {
	_jsii_.InvokeVoid(
		a,
		"resetLastCommitId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) ResetOwner() {
	_jsii_.InvokeVoid(
		a,
		"resetOwner",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) ResetProvider() {
	_jsii_.InvokeVoid(
		a,
		"resetProvider",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) ResetRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGlueJob_SourceControlDetailsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

