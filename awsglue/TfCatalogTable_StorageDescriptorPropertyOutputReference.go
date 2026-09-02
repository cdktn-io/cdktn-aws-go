package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCatalogTable_StorageDescriptorPropertyOutputReference interface {
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
	Columns() TfCatalogTable_ColumnsPropertyList
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
	InternalValue() *TfCatalogTable_StorageDescriptorProperty
	// Experimental.
	SetInternalValue(val *TfCatalogTable_StorageDescriptorProperty)
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
	SchemaReference() TfCatalogTable_SchemaReferencePropertyOutputReference
	// Experimental.
	SchemaReferenceInput() *TfCatalogTable_SchemaReferenceProperty
	// Experimental.
	SerDeInfo() TfCatalogTable_SerDeInfoPropertyOutputReference
	// Experimental.
	SerDeInfoInput() *TfCatalogTable_SerDeInfoProperty
	// Experimental.
	SkewedInfo() TfCatalogTable_SkewedInfoPropertyOutputReference
	// Experimental.
	SkewedInfoInput() *TfCatalogTable_SkewedInfoProperty
	// Experimental.
	SortColumns() TfCatalogTable_SortColumnsPropertyList
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
	PutSchemaReference(value *TfCatalogTable_SchemaReferenceProperty)
	// Experimental.
	PutSerDeInfo(value *TfCatalogTable_SerDeInfoProperty)
	// Experimental.
	PutSkewedInfo(value *TfCatalogTable_SkewedInfoProperty)
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
	ResetSchemaReference()
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

// The jsii proxy struct for TfCatalogTable_StorageDescriptorPropertyOutputReference
type jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) AdditionalLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) AdditionalLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) BucketColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) BucketColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) Columns() TfCatalogTable_ColumnsPropertyList {
	var returns TfCatalogTable_ColumnsPropertyList
	_jsii_.Get(
		j,
		"columns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) Compressed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) CompressedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) InputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) InputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) InternalValue() *TfCatalogTable_StorageDescriptorProperty {
	var returns *TfCatalogTable_StorageDescriptorProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) NumberOfBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) NumberOfBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) OutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) SchemaReference() TfCatalogTable_SchemaReferencePropertyOutputReference {
	var returns TfCatalogTable_SchemaReferencePropertyOutputReference
	_jsii_.Get(
		j,
		"schemaReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) SchemaReferenceInput() *TfCatalogTable_SchemaReferenceProperty {
	var returns *TfCatalogTable_SchemaReferenceProperty
	_jsii_.Get(
		j,
		"schemaReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) SerDeInfo() TfCatalogTable_SerDeInfoPropertyOutputReference {
	var returns TfCatalogTable_SerDeInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"serDeInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) SerDeInfoInput() *TfCatalogTable_SerDeInfoProperty {
	var returns *TfCatalogTable_SerDeInfoProperty
	_jsii_.Get(
		j,
		"serDeInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) SkewedInfo() TfCatalogTable_SkewedInfoPropertyOutputReference {
	var returns TfCatalogTable_SkewedInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"skewedInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) SkewedInfoInput() *TfCatalogTable_SkewedInfoProperty {
	var returns *TfCatalogTable_SkewedInfoProperty
	_jsii_.Get(
		j,
		"skewedInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) SortColumns() TfCatalogTable_SortColumnsPropertyList {
	var returns TfCatalogTable_SortColumnsPropertyList
	_jsii_.Get(
		j,
		"sortColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) SortColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) StoredAsSubDirectories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirectories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) StoredAsSubDirectoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirectoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCatalogTable_StorageDescriptorPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCatalogTable_StorageDescriptorPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCatalogTable_StorageDescriptorPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.TfCatalogTable.StorageDescriptorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCatalogTable_StorageDescriptorPropertyOutputReference_Override(t TfCatalogTable_StorageDescriptorPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.TfCatalogTable.StorageDescriptorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference)SetAdditionalLocations(val *[]*string) {
	if err := j.validateSetAdditionalLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalLocations",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference)SetBucketColumns(val *[]*string) {
	if err := j.validateSetBucketColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketColumns",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference)SetCompressed(val interface{}) {
	if err := j.validateSetCompressedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressed",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference)SetInputFormat(val *string) {
	if err := j.validateSetInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFormat",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference)SetInternalValue(val *TfCatalogTable_StorageDescriptorProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference)SetNumberOfBuckets(val *float64) {
	if err := j.validateSetNumberOfBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfBuckets",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference)SetOutputFormat(val *string) {
	if err := j.validateSetOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference)SetStoredAsSubDirectories(val interface{}) {
	if err := j.validateSetStoredAsSubDirectoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storedAsSubDirectories",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) PutColumns(value interface{}) {
	if err := t.validatePutColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putColumns",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) PutSchemaReference(value *TfCatalogTable_SchemaReferenceProperty) {
	if err := t.validatePutSchemaReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSchemaReference",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) PutSerDeInfo(value *TfCatalogTable_SerDeInfoProperty) {
	if err := t.validatePutSerDeInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSerDeInfo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) PutSkewedInfo(value *TfCatalogTable_SkewedInfoProperty) {
	if err := t.validatePutSkewedInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSkewedInfo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) PutSortColumns(value interface{}) {
	if err := t.validatePutSortColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSortColumns",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ResetAdditionalLocations() {
	_jsii_.InvokeVoid(
		t,
		"resetAdditionalLocations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ResetBucketColumns() {
	_jsii_.InvokeVoid(
		t,
		"resetBucketColumns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ResetColumns() {
	_jsii_.InvokeVoid(
		t,
		"resetColumns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ResetCompressed() {
	_jsii_.InvokeVoid(
		t,
		"resetCompressed",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ResetInputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetInputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ResetNumberOfBuckets() {
	_jsii_.InvokeVoid(
		t,
		"resetNumberOfBuckets",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ResetOutputFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ResetSchemaReference() {
	_jsii_.InvokeVoid(
		t,
		"resetSchemaReference",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ResetSerDeInfo() {
	_jsii_.InvokeVoid(
		t,
		"resetSerDeInfo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ResetSkewedInfo() {
	_jsii_.InvokeVoid(
		t,
		"resetSkewedInfo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ResetSortColumns() {
	_jsii_.InvokeVoid(
		t,
		"resetSortColumns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ResetStoredAsSubDirectories() {
	_jsii_.InvokeVoid(
		t,
		"resetStoredAsSubDirectories",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCatalogTable_StorageDescriptorPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

