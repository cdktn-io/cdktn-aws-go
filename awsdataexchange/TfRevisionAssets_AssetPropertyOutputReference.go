package awsdataexchange

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdataexchange/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdataexchange/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRevisionAssets_AssetPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Arn() *string
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
	// Experimental.
	CreatedAt() *string
	// Experimental.
	CreateS3DataAccessFromS3Bucket() TfRevisionAssets_CreateS3DataAccessFromS3BucketPropertyList
	// Experimental.
	CreateS3DataAccessFromS3BucketInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	Id() *string
	// Experimental.
	ImportAssetsFromS3() TfRevisionAssets_ImportAssetsFromS3PropertyList
	// Experimental.
	ImportAssetsFromS3Input() interface{}
	// Experimental.
	ImportAssetsFromSignedUrl() TfRevisionAssets_ImportAssetsFromSignedUrlPropertyList
	// Experimental.
	ImportAssetsFromSignedUrlInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Name() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UpdatedAt() *string
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
	PutCreateS3DataAccessFromS3Bucket(value interface{})
	// Experimental.
	PutImportAssetsFromS3(value interface{})
	// Experimental.
	PutImportAssetsFromSignedUrl(value interface{})
	// Experimental.
	ResetCreateS3DataAccessFromS3Bucket()
	// Experimental.
	ResetImportAssetsFromS3()
	// Experimental.
	ResetImportAssetsFromSignedUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfRevisionAssets_AssetPropertyOutputReference
type jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) CreateS3DataAccessFromS3Bucket() TfRevisionAssets_CreateS3DataAccessFromS3BucketPropertyList {
	var returns TfRevisionAssets_CreateS3DataAccessFromS3BucketPropertyList
	_jsii_.Get(
		j,
		"createS3DataAccessFromS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) CreateS3DataAccessFromS3BucketInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createS3DataAccessFromS3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) ImportAssetsFromS3() TfRevisionAssets_ImportAssetsFromS3PropertyList {
	var returns TfRevisionAssets_ImportAssetsFromS3PropertyList
	_jsii_.Get(
		j,
		"importAssetsFromS3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) ImportAssetsFromS3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importAssetsFromS3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) ImportAssetsFromSignedUrl() TfRevisionAssets_ImportAssetsFromSignedUrlPropertyList {
	var returns TfRevisionAssets_ImportAssetsFromSignedUrlPropertyList
	_jsii_.Get(
		j,
		"importAssetsFromSignedUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) ImportAssetsFromSignedUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importAssetsFromSignedUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRevisionAssets_AssetPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfRevisionAssets_AssetPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRevisionAssets_AssetPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-data-exchange.TfRevisionAssets.AssetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRevisionAssets_AssetPropertyOutputReference_Override(t TfRevisionAssets_AssetPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-data-exchange.TfRevisionAssets.AssetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) PutCreateS3DataAccessFromS3Bucket(value interface{}) {
	if err := t.validatePutCreateS3DataAccessFromS3BucketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCreateS3DataAccessFromS3Bucket",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) PutImportAssetsFromS3(value interface{}) {
	if err := t.validatePutImportAssetsFromS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putImportAssetsFromS3",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) PutImportAssetsFromSignedUrl(value interface{}) {
	if err := t.validatePutImportAssetsFromSignedUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putImportAssetsFromSignedUrl",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) ResetCreateS3DataAccessFromS3Bucket() {
	_jsii_.InvokeVoid(
		t,
		"resetCreateS3DataAccessFromS3Bucket",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) ResetImportAssetsFromS3() {
	_jsii_.InvokeVoid(
		t,
		"resetImportAssetsFromS3",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) ResetImportAssetsFromSignedUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetImportAssetsFromSignedUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRevisionAssets_AssetPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

