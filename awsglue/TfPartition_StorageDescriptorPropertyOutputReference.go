package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPartition_StorageDescriptorPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdditionalLocations() *[]*string
	// Experimental.
	SetAdditionalLocations(val *[]*string)
	// Experimental.
	AdditionalLocationsInput() *[]*string
	// Experimental.
	BucketColumns() *[]*string
	// Experimental.
	SetBucketColumns(val *[]*string)
	// Experimental.
	BucketColumnsInput() *[]*string
	// Experimental.
	Columns() TfPartition_ColumnsPropertyList
	// Experimental.
	ColumnsInput() interface{}
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
	Compressed() interface{}
	// Experimental.
	SetCompressed(val interface{})
	// Experimental.
	CompressedInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InputFormat() *string
	// Experimental.
	SetInputFormat(val *string)
	// Experimental.
	InputFormatInput() *string
	// Experimental.
	InternalValue() *TfPartition_StorageDescriptorProperty
	// Experimental.
	SetInternalValue(val *TfPartition_StorageDescriptorProperty)
	// Experimental.
	Location() *string
	// Experimental.
	SetLocation(val *string)
	// Experimental.
	LocationInput() *string
	// Experimental.
	NumberOfBuckets() *float64
	// Experimental.
	SetNumberOfBuckets(val *float64)
	// Experimental.
	NumberOfBucketsInput() *float64
	// Experimental.
	OutputFormat() *string
	// Experimental.
	SetOutputFormat(val *string)
	// Experimental.
	OutputFormatInput() *string
	// Experimental.
	Parameters() *map[string]*string
	// Experimental.
	SetParameters(val *map[string]*string)
	// Experimental.
	ParametersInput() *map[string]*string
	// Experimental.
	SerDeInfo() TfPartition_SerDeInfoPropertyOutputReference
	// Experimental.
	SerDeInfoInput() *TfPartition_SerDeInfoProperty
	// Experimental.
	SkewedInfo() TfPartition_SkewedInfoPropertyOutputReference
	// Experimental.
	SkewedInfoInput() *TfPartition_SkewedInfoProperty
	// Experimental.
	SortColumns() TfPartition_SortColumnsPropertyList
	// Experimental.
	SortColumnsInput() interface{}
	// Experimental.
	StoredAsSubDirectories() interface{}
	// Experimental.
	SetStoredAsSubDirectories(val interface{})
	// Experimental.
	StoredAsSubDirectoriesInput() interface{}
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
	PutColumns(value interface{})
	// Experimental.
	PutSerDeInfo(value *TfPartition_SerDeInfoProperty)
	// Experimental.
	PutSkewedInfo(value *TfPartition_SkewedInfoProperty)
	// Experimental.
	PutSortColumns(value interface{})
	// Experimental.
	ResetAdditionalLocations()
	// Experimental.
	ResetBucketColumns()
	// Experimental.
	ResetColumns()
	// Experimental.
	ResetCompressed()
	// Experimental.
	ResetInputFormat()
	// Experimental.
	ResetLocation()
	// Experimental.
	ResetNumberOfBuckets()
	// Experimental.
	ResetOutputFormat()
	// Experimental.
	ResetParameters()
	// Experimental.
	ResetSerDeInfo()
	// Experimental.
	ResetSkewedInfo()
	// Experimental.
	ResetSortColumns()
	// Experimental.
	ResetStoredAsSubDirectories()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPartition_StorageDescriptorPropertyOutputReference
type jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) AdditionalLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) AdditionalLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) BucketColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) BucketColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) Columns() TfPartition_ColumnsPropertyList {
	var returns TfPartition_ColumnsPropertyList
	_jsii_.Get(
		j,
		"columns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) Compressed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) CompressedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) InputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) InputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) InternalValue() *TfPartition_StorageDescriptorProperty {
	var returns *TfPartition_StorageDescriptorProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) NumberOfBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) NumberOfBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) OutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) SerDeInfo() TfPartition_SerDeInfoPropertyOutputReference {
	var returns TfPartition_SerDeInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"serDeInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) SerDeInfoInput() *TfPartition_SerDeInfoProperty {
	var returns *TfPartition_SerDeInfoProperty
	_jsii_.Get(
		j,
		"serDeInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) SkewedInfo() TfPartition_SkewedInfoPropertyOutputReference {
	var returns TfPartition_SkewedInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"skewedInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) SkewedInfoInput() *TfPartition_SkewedInfoProperty {
	var returns *TfPartition_SkewedInfoProperty
	_jsii_.Get(
		j,
		"skewedInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) SortColumns() TfPartition_SortColumnsPropertyList {
	var returns TfPartition_SortColumnsPropertyList
	_jsii_.Get(
		j,
		"sortColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) SortColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) StoredAsSubDirectories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirectories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) StoredAsSubDirectoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirectoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPartition_StorageDescriptorPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPartition_StorageDescriptorPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPartition_StorageDescriptorPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.TfPartition.StorageDescriptorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPartition_StorageDescriptorPropertyOutputReference_Override(t TfPartition_StorageDescriptorPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.TfPartition.StorageDescriptorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference)SetAdditionalLocations(val *[]*string) {
	if err := j.validateSetAdditionalLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalLocations",
		val,
	)
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference)SetBucketColumns(val *[]*string) {
	if err := j.validateSetBucketColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketColumns",
		val,
	)
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference)SetCompressed(val interface{}) {
	if err := j.validateSetCompressedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressed",
		val,
	)
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference)SetInputFormat(val *string) {
	if err := j.validateSetInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFormat",
		val,
	)
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference)SetInternalValue(val *TfPartition_StorageDescriptorProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference)SetNumberOfBuckets(val *float64) {
	if err := j.validateSetNumberOfBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfBuckets",
		val,
	)
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference)SetOutputFormat(val *string) {
	if err := j.validateSetOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference)SetStoredAsSubDirectories(val interface{}) {
	if err := j.validateSetStoredAsSubDirectoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storedAsSubDirectories",
		val,
	)
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) PutColumns(value interface{}) {
	if err := t.validatePutColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putColumns",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) PutSerDeInfo(value *TfPartition_SerDeInfoProperty) {
	if err := t.validatePutSerDeInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSerDeInfo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) PutSkewedInfo(value *TfPartition_SkewedInfoProperty) {
	if err := t.validatePutSkewedInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSkewedInfo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) PutSortColumns(value interface{}) {
	if err := t.validatePutSortColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSortColumns",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ResetAdditionalLocations() {
	_jsii_.InvokeVoid(
		t,
		"resetAdditionalLocations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ResetBucketColumns() {
	_jsii_.InvokeVoid(
		t,
		"resetBucketColumns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ResetColumns() {
	_jsii_.InvokeVoid(
		t,
		"resetColumns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ResetCompressed() {
	_jsii_.InvokeVoid(
		t,
		"resetCompressed",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ResetInputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetInputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ResetNumberOfBuckets() {
	_jsii_.InvokeVoid(
		t,
		"resetNumberOfBuckets",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ResetOutputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ResetSerDeInfo() {
	_jsii_.InvokeVoid(
		t,
		"resetSerDeInfo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ResetSkewedInfo() {
	_jsii_.InvokeVoid(
		t,
		"resetSkewedInfo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ResetSortColumns() {
	_jsii_.InvokeVoid(
		t,
		"resetSortColumns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ResetStoredAsSubDirectories() {
	_jsii_.InvokeVoid(
		t,
		"resetStoredAsSubDirectories",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPartition_StorageDescriptorPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

