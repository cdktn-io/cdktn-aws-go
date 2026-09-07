package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/glue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/glue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPartition_StorageDescriptorPropertyOutputReference interface {
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
	Columns() AwsPartition_ColumnsPropertyList
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
	InternalValue() *AwsPartition_StorageDescriptorProperty
	// Experimental.
	SetInternalValue(val *AwsPartition_StorageDescriptorProperty)
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
	SerDeInfo() AwsPartition_SerDeInfoPropertyOutputReference
	// Experimental.
	SerDeInfoInput() *AwsPartition_SerDeInfoProperty
	// Experimental.
	SkewedInfo() AwsPartition_SkewedInfoPropertyOutputReference
	// Experimental.
	SkewedInfoInput() *AwsPartition_SkewedInfoProperty
	// Experimental.
	SortColumns() AwsPartition_SortColumnsPropertyList
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
	PutSerDeInfo(value *AwsPartition_SerDeInfoProperty)
	// Experimental.
	PutSkewedInfo(value *AwsPartition_SkewedInfoProperty)
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

// The jsii proxy struct for AwsPartition_StorageDescriptorPropertyOutputReference
type jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) AdditionalLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) AdditionalLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) BucketColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) BucketColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) Columns() AwsPartition_ColumnsPropertyList {
	var returns AwsPartition_ColumnsPropertyList
	_jsii_.Get(
		j,
		"columns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) Compressed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) CompressedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) InputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) InputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) InternalValue() *AwsPartition_StorageDescriptorProperty {
	var returns *AwsPartition_StorageDescriptorProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) NumberOfBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) NumberOfBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) OutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) SerDeInfo() AwsPartition_SerDeInfoPropertyOutputReference {
	var returns AwsPartition_SerDeInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"serDeInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) SerDeInfoInput() *AwsPartition_SerDeInfoProperty {
	var returns *AwsPartition_SerDeInfoProperty
	_jsii_.Get(
		j,
		"serDeInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) SkewedInfo() AwsPartition_SkewedInfoPropertyOutputReference {
	var returns AwsPartition_SkewedInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"skewedInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) SkewedInfoInput() *AwsPartition_SkewedInfoProperty {
	var returns *AwsPartition_SkewedInfoProperty
	_jsii_.Get(
		j,
		"skewedInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) SortColumns() AwsPartition_SortColumnsPropertyList {
	var returns AwsPartition_SortColumnsPropertyList
	_jsii_.Get(
		j,
		"sortColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) SortColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) StoredAsSubDirectories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirectories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) StoredAsSubDirectoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storedAsSubDirectoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPartition_StorageDescriptorPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsPartition_StorageDescriptorPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPartition_StorageDescriptorPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsPartition.StorageDescriptorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPartition_StorageDescriptorPropertyOutputReference_Override(a AwsPartition_StorageDescriptorPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsPartition.StorageDescriptorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference)SetAdditionalLocations(val *[]*string) {
	if err := j.validateSetAdditionalLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalLocations",
		val,
	)
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference)SetBucketColumns(val *[]*string) {
	if err := j.validateSetBucketColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketColumns",
		val,
	)
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference)SetCompressed(val interface{}) {
	if err := j.validateSetCompressedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressed",
		val,
	)
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference)SetInputFormat(val *string) {
	if err := j.validateSetInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFormat",
		val,
	)
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference)SetInternalValue(val *AwsPartition_StorageDescriptorProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference)SetNumberOfBuckets(val *float64) {
	if err := j.validateSetNumberOfBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfBuckets",
		val,
	)
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference)SetOutputFormat(val *string) {
	if err := j.validateSetOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference)SetStoredAsSubDirectories(val interface{}) {
	if err := j.validateSetStoredAsSubDirectoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storedAsSubDirectories",
		val,
	)
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) PutColumns(value interface{}) {
	if err := a.validatePutColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putColumns",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) PutSerDeInfo(value *AwsPartition_SerDeInfoProperty) {
	if err := a.validatePutSerDeInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSerDeInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) PutSkewedInfo(value *AwsPartition_SkewedInfoProperty) {
	if err := a.validatePutSkewedInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSkewedInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) PutSortColumns(value interface{}) {
	if err := a.validatePutSortColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSortColumns",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ResetAdditionalLocations() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalLocations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ResetBucketColumns() {
	_jsii_.InvokeVoid(
		a,
		"resetBucketColumns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ResetColumns() {
	_jsii_.InvokeVoid(
		a,
		"resetColumns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ResetCompressed() {
	_jsii_.InvokeVoid(
		a,
		"resetCompressed",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ResetInputFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetInputFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ResetNumberOfBuckets() {
	_jsii_.InvokeVoid(
		a,
		"resetNumberOfBuckets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ResetOutputFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ResetSerDeInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetSerDeInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ResetSkewedInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetSkewedInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ResetSortColumns() {
	_jsii_.InvokeVoid(
		a,
		"resetSortColumns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ResetStoredAsSubDirectories() {
	_jsii_.InvokeVoid(
		a,
		"resetStoredAsSubDirectories",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPartition_StorageDescriptorPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

