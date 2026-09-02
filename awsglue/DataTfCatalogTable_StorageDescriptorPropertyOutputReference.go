package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfCatalogTable_StorageDescriptorPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdditionalLocations() *[]*string
	// Experimental.
	BucketColumns() *[]*string
	// Experimental.
	Columns() DataTfCatalogTable_ColumnsPropertyList
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
	Compressed() cdktn.IResolvable
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
	InternalValue() *DataTfCatalogTable_StorageDescriptorProperty
	// Experimental.
	SetInternalValue(val *DataTfCatalogTable_StorageDescriptorProperty)
	// Experimental.
	Location() *string
	// Experimental.
	NumberOfBuckets() *float64
	// Experimental.
	OutputFormat() *string
	// Experimental.
	Parameters() cdktn.StringMap
	// Experimental.
	SchemaReference() DataTfCatalogTable_SchemaReferencePropertyList
	// Experimental.
	SerDeInfo() DataTfCatalogTable_SerDeInfoPropertyList
	// Experimental.
	SkewedInfo() DataTfCatalogTable_SkewedInfoPropertyList
	// Experimental.
	SortColumns() DataTfCatalogTable_SortColumnsPropertyList
	// Experimental.
	StoredAsSubDirectories() cdktn.IResolvable
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataTfCatalogTable_StorageDescriptorPropertyOutputReference
type jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) AdditionalLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) BucketColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"bucketColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) Columns() DataTfCatalogTable_ColumnsPropertyList {
	var returns DataTfCatalogTable_ColumnsPropertyList
	_jsii_.Get(
		j,
		"columns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) Compressed() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"compressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) InputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) InternalValue() *DataTfCatalogTable_StorageDescriptorProperty {
	var returns *DataTfCatalogTable_StorageDescriptorProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) NumberOfBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) Parameters() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) SchemaReference() DataTfCatalogTable_SchemaReferencePropertyList {
	var returns DataTfCatalogTable_SchemaReferencePropertyList
	_jsii_.Get(
		j,
		"schemaReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) SerDeInfo() DataTfCatalogTable_SerDeInfoPropertyList {
	var returns DataTfCatalogTable_SerDeInfoPropertyList
	_jsii_.Get(
		j,
		"serDeInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) SkewedInfo() DataTfCatalogTable_SkewedInfoPropertyList {
	var returns DataTfCatalogTable_SkewedInfoPropertyList
	_jsii_.Get(
		j,
		"skewedInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) SortColumns() DataTfCatalogTable_SortColumnsPropertyList {
	var returns DataTfCatalogTable_SortColumnsPropertyList
	_jsii_.Get(
		j,
		"sortColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) StoredAsSubDirectories() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"storedAsSubDirectories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfCatalogTable_StorageDescriptorPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataTfCatalogTable_StorageDescriptorPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfCatalogTable_StorageDescriptorPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.DataTfCatalogTable.StorageDescriptorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfCatalogTable_StorageDescriptorPropertyOutputReference_Override(d DataTfCatalogTable_StorageDescriptorPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.DataTfCatalogTable.StorageDescriptorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference)SetInternalValue(val *DataTfCatalogTable_StorageDescriptorProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCatalogTable_StorageDescriptorPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

