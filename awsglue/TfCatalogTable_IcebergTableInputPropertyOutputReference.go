package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCatalogTable_IcebergTableInputPropertyOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *TfCatalogTable_IcebergTableInputProperty
	// Experimental.
	SetInternalValue(val *TfCatalogTable_IcebergTableInputProperty)
	// Experimental.
	Location() *string
	// Experimental.
	SetLocation(val *string)
	// Experimental.
	LocationInput() *string
	// Experimental.
	PartitionSpec() TfCatalogTable_PartitionSpecPropertyOutputReference
	// Experimental.
	PartitionSpecInput() *TfCatalogTable_PartitionSpecProperty
	// Experimental.
	Properties() *map[string]*string
	// Experimental.
	SetProperties(val *map[string]*string)
	// Experimental.
	PropertiesInput() *map[string]*string
	// Experimental.
	Schema() TfCatalogTable_SchemaPropertyOutputReference
	// Experimental.
	SchemaInput() *TfCatalogTable_SchemaProperty
	// Experimental.
	SortOrder() TfCatalogTable_SortOrderPropertyOutputReference
	// Experimental.
	SortOrderInput() *TfCatalogTable_SortOrderProperty
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
	PutPartitionSpec(value *TfCatalogTable_PartitionSpecProperty)
	// Experimental.
	PutSchema(value *TfCatalogTable_SchemaProperty)
	// Experimental.
	PutSortOrder(value *TfCatalogTable_SortOrderProperty)
	// Experimental.
	ResetPartitionSpec()
	// Experimental.
	ResetProperties()
	// Experimental.
	ResetSortOrder()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCatalogTable_IcebergTableInputPropertyOutputReference
type jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) InternalValue() *TfCatalogTable_IcebergTableInputProperty {
	var returns *TfCatalogTable_IcebergTableInputProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) PartitionSpec() TfCatalogTable_PartitionSpecPropertyOutputReference {
	var returns TfCatalogTable_PartitionSpecPropertyOutputReference
	_jsii_.Get(
		j,
		"partitionSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) PartitionSpecInput() *TfCatalogTable_PartitionSpecProperty {
	var returns *TfCatalogTable_PartitionSpecProperty
	_jsii_.Get(
		j,
		"partitionSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) Schema() TfCatalogTable_SchemaPropertyOutputReference {
	var returns TfCatalogTable_SchemaPropertyOutputReference
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) SchemaInput() *TfCatalogTable_SchemaProperty {
	var returns *TfCatalogTable_SchemaProperty
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) SortOrder() TfCatalogTable_SortOrderPropertyOutputReference {
	var returns TfCatalogTable_SortOrderPropertyOutputReference
	_jsii_.Get(
		j,
		"sortOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) SortOrderInput() *TfCatalogTable_SortOrderProperty {
	var returns *TfCatalogTable_SortOrderProperty
	_jsii_.Get(
		j,
		"sortOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCatalogTable_IcebergTableInputPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCatalogTable_IcebergTableInputPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCatalogTable_IcebergTableInputPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.TfCatalogTable.IcebergTableInputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCatalogTable_IcebergTableInputPropertyOutputReference_Override(t TfCatalogTable_IcebergTableInputPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.TfCatalogTable.IcebergTableInputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference)SetInternalValue(val *TfCatalogTable_IcebergTableInputProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) PutPartitionSpec(value *TfCatalogTable_PartitionSpecProperty) {
	if err := t.validatePutPartitionSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPartitionSpec",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) PutSchema(value *TfCatalogTable_SchemaProperty) {
	if err := t.validatePutSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSchema",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) PutSortOrder(value *TfCatalogTable_SortOrderProperty) {
	if err := t.validatePutSortOrderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSortOrder",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) ResetPartitionSpec() {
	_jsii_.InvokeVoid(
		t,
		"resetPartitionSpec",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		t,
		"resetProperties",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) ResetSortOrder() {
	_jsii_.InvokeVoid(
		t,
		"resetSortOrder",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCatalogTable_IcebergTableInputPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

