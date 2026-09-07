package codegurureviewer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/codegurureviewer/jsii"

	"github.com/cdktn-io/cdktn-aws-go/codegurureviewer/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRepositoryAssociation_RepositoryPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Bitbucket() AwsRepositoryAssociation_BitbucketPropertyOutputReference
	// Experimental.
	BitbucketInput() *AwsRepositoryAssociation_BitbucketProperty
	// Experimental.
	Codecommit() AwsRepositoryAssociation_CodecommitPropertyOutputReference
	// Experimental.
	CodecommitInput() *AwsRepositoryAssociation_CodecommitProperty
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
	Fqn() *string
	// Experimental.
	GithubEnterpriseServer() AwsRepositoryAssociation_GithubEnterpriseServerPropertyOutputReference
	// Experimental.
	GithubEnterpriseServerInput() *AwsRepositoryAssociation_GithubEnterpriseServerProperty
	// Experimental.
	InternalValue() *AwsRepositoryAssociation_RepositoryProperty
	// Experimental.
	SetInternalValue(val *AwsRepositoryAssociation_RepositoryProperty)
	// Experimental.
	S3Bucket() AwsRepositoryAssociation_S3BucketPropertyOutputReference
	// Experimental.
	S3BucketInput() *AwsRepositoryAssociation_S3BucketProperty
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
	PutBitbucket(value *AwsRepositoryAssociation_BitbucketProperty)
	// Experimental.
	PutCodecommit(value *AwsRepositoryAssociation_CodecommitProperty)
	// Experimental.
	PutGithubEnterpriseServer(value *AwsRepositoryAssociation_GithubEnterpriseServerProperty)
	// Experimental.
	PutS3Bucket(value *AwsRepositoryAssociation_S3BucketProperty)
	// Experimental.
	ResetBitbucket()
	// Experimental.
	ResetCodecommit()
	// Experimental.
	ResetGithubEnterpriseServer()
	// Experimental.
	ResetS3Bucket()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsRepositoryAssociation_RepositoryPropertyOutputReference
type jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) Bitbucket() AwsRepositoryAssociation_BitbucketPropertyOutputReference {
	var returns AwsRepositoryAssociation_BitbucketPropertyOutputReference
	_jsii_.Get(
		j,
		"bitbucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) BitbucketInput() *AwsRepositoryAssociation_BitbucketProperty {
	var returns *AwsRepositoryAssociation_BitbucketProperty
	_jsii_.Get(
		j,
		"bitbucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) Codecommit() AwsRepositoryAssociation_CodecommitPropertyOutputReference {
	var returns AwsRepositoryAssociation_CodecommitPropertyOutputReference
	_jsii_.Get(
		j,
		"codecommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) CodecommitInput() *AwsRepositoryAssociation_CodecommitProperty {
	var returns *AwsRepositoryAssociation_CodecommitProperty
	_jsii_.Get(
		j,
		"codecommitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) GithubEnterpriseServer() AwsRepositoryAssociation_GithubEnterpriseServerPropertyOutputReference {
	var returns AwsRepositoryAssociation_GithubEnterpriseServerPropertyOutputReference
	_jsii_.Get(
		j,
		"githubEnterpriseServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) GithubEnterpriseServerInput() *AwsRepositoryAssociation_GithubEnterpriseServerProperty {
	var returns *AwsRepositoryAssociation_GithubEnterpriseServerProperty
	_jsii_.Get(
		j,
		"githubEnterpriseServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) InternalValue() *AwsRepositoryAssociation_RepositoryProperty {
	var returns *AwsRepositoryAssociation_RepositoryProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) S3Bucket() AwsRepositoryAssociation_S3BucketPropertyOutputReference {
	var returns AwsRepositoryAssociation_S3BucketPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) S3BucketInput() *AwsRepositoryAssociation_S3BucketProperty {
	var returns *AwsRepositoryAssociation_S3BucketProperty
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRepositoryAssociation_RepositoryPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsRepositoryAssociation_RepositoryPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRepositoryAssociation_RepositoryPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codeguru-reviewer.AwsRepositoryAssociation.RepositoryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRepositoryAssociation_RepositoryPropertyOutputReference_Override(a AwsRepositoryAssociation_RepositoryPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codeguru-reviewer.AwsRepositoryAssociation.RepositoryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference)SetInternalValue(val *AwsRepositoryAssociation_RepositoryProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) PutBitbucket(value *AwsRepositoryAssociation_BitbucketProperty) {
	if err := a.validatePutBitbucketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBitbucket",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) PutCodecommit(value *AwsRepositoryAssociation_CodecommitProperty) {
	if err := a.validatePutCodecommitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodecommit",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) PutGithubEnterpriseServer(value *AwsRepositoryAssociation_GithubEnterpriseServerProperty) {
	if err := a.validatePutGithubEnterpriseServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGithubEnterpriseServer",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) PutS3Bucket(value *AwsRepositoryAssociation_S3BucketProperty) {
	if err := a.validatePutS3BucketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Bucket",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) ResetBitbucket() {
	_jsii_.InvokeVoid(
		a,
		"resetBitbucket",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) ResetCodecommit() {
	_jsii_.InvokeVoid(
		a,
		"resetCodecommit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) ResetGithubEnterpriseServer() {
	_jsii_.InvokeVoid(
		a,
		"resetGithubEnterpriseServer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) ResetS3Bucket() {
	_jsii_.InvokeVoid(
		a,
		"resetS3Bucket",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRepositoryAssociation_RepositoryPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

