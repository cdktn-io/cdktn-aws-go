package awscodegurureviewer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodegurureviewer/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodegurureviewer/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRepositoryAssociation_RepositoryPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Bitbucket() TfRepositoryAssociation_BitbucketPropertyOutputReference
	// Experimental.
	BitbucketInput() *TfRepositoryAssociation_BitbucketProperty
	// Experimental.
	Codecommit() TfRepositoryAssociation_CodecommitPropertyOutputReference
	// Experimental.
	CodecommitInput() *TfRepositoryAssociation_CodecommitProperty
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
	GithubEnterpriseServer() TfRepositoryAssociation_GithubEnterpriseServerPropertyOutputReference
	// Experimental.
	GithubEnterpriseServerInput() *TfRepositoryAssociation_GithubEnterpriseServerProperty
	// Experimental.
	InternalValue() *TfRepositoryAssociation_RepositoryProperty
	// Experimental.
	SetInternalValue(val *TfRepositoryAssociation_RepositoryProperty)
	// Experimental.
	S3Bucket() TfRepositoryAssociation_S3BucketPropertyOutputReference
	// Experimental.
	S3BucketInput() *TfRepositoryAssociation_S3BucketProperty
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
	PutBitbucket(value *TfRepositoryAssociation_BitbucketProperty)
	// Experimental.
	PutCodecommit(value *TfRepositoryAssociation_CodecommitProperty)
	// Experimental.
	PutGithubEnterpriseServer(value *TfRepositoryAssociation_GithubEnterpriseServerProperty)
	// Experimental.
	PutS3Bucket(value *TfRepositoryAssociation_S3BucketProperty)
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

// The jsii proxy struct for TfRepositoryAssociation_RepositoryPropertyOutputReference
type jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) Bitbucket() TfRepositoryAssociation_BitbucketPropertyOutputReference {
	var returns TfRepositoryAssociation_BitbucketPropertyOutputReference
	_jsii_.Get(
		j,
		"bitbucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) BitbucketInput() *TfRepositoryAssociation_BitbucketProperty {
	var returns *TfRepositoryAssociation_BitbucketProperty
	_jsii_.Get(
		j,
		"bitbucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) Codecommit() TfRepositoryAssociation_CodecommitPropertyOutputReference {
	var returns TfRepositoryAssociation_CodecommitPropertyOutputReference
	_jsii_.Get(
		j,
		"codecommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) CodecommitInput() *TfRepositoryAssociation_CodecommitProperty {
	var returns *TfRepositoryAssociation_CodecommitProperty
	_jsii_.Get(
		j,
		"codecommitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) GithubEnterpriseServer() TfRepositoryAssociation_GithubEnterpriseServerPropertyOutputReference {
	var returns TfRepositoryAssociation_GithubEnterpriseServerPropertyOutputReference
	_jsii_.Get(
		j,
		"githubEnterpriseServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) GithubEnterpriseServerInput() *TfRepositoryAssociation_GithubEnterpriseServerProperty {
	var returns *TfRepositoryAssociation_GithubEnterpriseServerProperty
	_jsii_.Get(
		j,
		"githubEnterpriseServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) InternalValue() *TfRepositoryAssociation_RepositoryProperty {
	var returns *TfRepositoryAssociation_RepositoryProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) S3Bucket() TfRepositoryAssociation_S3BucketPropertyOutputReference {
	var returns TfRepositoryAssociation_S3BucketPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) S3BucketInput() *TfRepositoryAssociation_S3BucketProperty {
	var returns *TfRepositoryAssociation_S3BucketProperty
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRepositoryAssociation_RepositoryPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfRepositoryAssociation_RepositoryPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRepositoryAssociation_RepositoryPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codeguru-reviewer.TfRepositoryAssociation.RepositoryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRepositoryAssociation_RepositoryPropertyOutputReference_Override(t TfRepositoryAssociation_RepositoryPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codeguru-reviewer.TfRepositoryAssociation.RepositoryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference)SetInternalValue(val *TfRepositoryAssociation_RepositoryProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) PutBitbucket(value *TfRepositoryAssociation_BitbucketProperty) {
	if err := t.validatePutBitbucketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBitbucket",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) PutCodecommit(value *TfRepositoryAssociation_CodecommitProperty) {
	if err := t.validatePutCodecommitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCodecommit",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) PutGithubEnterpriseServer(value *TfRepositoryAssociation_GithubEnterpriseServerProperty) {
	if err := t.validatePutGithubEnterpriseServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGithubEnterpriseServer",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) PutS3Bucket(value *TfRepositoryAssociation_S3BucketProperty) {
	if err := t.validatePutS3BucketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Bucket",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) ResetBitbucket() {
	_jsii_.InvokeVoid(
		t,
		"resetBitbucket",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) ResetCodecommit() {
	_jsii_.InvokeVoid(
		t,
		"resetCodecommit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) ResetGithubEnterpriseServer() {
	_jsii_.InvokeVoid(
		t,
		"resetGithubEnterpriseServer",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) ResetS3Bucket() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Bucket",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRepositoryAssociation_RepositoryPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

