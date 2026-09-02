package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFileCache_DataRepositoryAssociationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AssociationId() *string
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
	DataRepositoryPath() *string
	// Experimental.
	SetDataRepositoryPath(val *string)
	// Experimental.
	DataRepositoryPathInput() *string
	// Experimental.
	DataRepositorySubdirectories() *[]*string
	// Experimental.
	SetDataRepositorySubdirectories(val *[]*string)
	// Experimental.
	DataRepositorySubdirectoriesInput() *[]*string
	// Experimental.
	FileCacheId() *string
	// Experimental.
	FileCachePath() *string
	// Experimental.
	SetFileCachePath(val *string)
	// Experimental.
	FileCachePathInput() *string
	// Experimental.
	FileSystemId() *string
	// Experimental.
	FileSystemPath() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	ImportedFileChunkSize() *float64
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Nfs() TfFileCache_NfsPropertyList
	// Experimental.
	NfsInput() interface{}
	// Experimental.
	ResourceArn() *string
	// Experimental.
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsInput() *map[string]*string
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
	PutNfs(value interface{})
	// Experimental.
	ResetDataRepositorySubdirectories()
	// Experimental.
	ResetNfs()
	// Experimental.
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFileCache_DataRepositoryAssociationPropertyOutputReference
type jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) AssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) DataRepositoryPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataRepositoryPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) DataRepositoryPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataRepositoryPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) DataRepositorySubdirectories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataRepositorySubdirectories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) DataRepositorySubdirectoriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataRepositorySubdirectoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) FileCacheId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileCacheId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) FileCachePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileCachePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) FileCachePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileCachePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) FileSystemPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) ImportedFileChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importedFileChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) Nfs() TfFileCache_NfsPropertyList {
	var returns TfFileCache_NfsPropertyList
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) NfsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFileCache_DataRepositoryAssociationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfFileCache_DataRepositoryAssociationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFileCache_DataRepositoryAssociationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fsx.TfFileCache.DataRepositoryAssociationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFileCache_DataRepositoryAssociationPropertyOutputReference_Override(t TfFileCache_DataRepositoryAssociationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.TfFileCache.DataRepositoryAssociationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference)SetDataRepositoryPath(val *string) {
	if err := j.validateSetDataRepositoryPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataRepositoryPath",
		val,
	)
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference)SetDataRepositorySubdirectories(val *[]*string) {
	if err := j.validateSetDataRepositorySubdirectoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataRepositorySubdirectories",
		val,
	)
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference)SetFileCachePath(val *string) {
	if err := j.validateSetFileCachePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileCachePath",
		val,
	)
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) PutNfs(value interface{}) {
	if err := t.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNfs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) ResetDataRepositorySubdirectories() {
	_jsii_.InvokeVoid(
		t,
		"resetDataRepositorySubdirectories",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		t,
		"resetNfs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFileCache_DataRepositoryAssociationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

