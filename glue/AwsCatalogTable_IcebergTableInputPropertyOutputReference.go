package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/glue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/glue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCatalogTable_IcebergTableInputPropertyOutputReference interface {
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
	InternalValue() *AwsCatalogTable_IcebergTableInputProperty
	// Experimental.
	SetInternalValue(val *AwsCatalogTable_IcebergTableInputProperty)
	// Experimental.
	Location() *string
	// Experimental.
	SetLocation(val *string)
	// Experimental.
	LocationInput() *string
	// Experimental.
	PartitionSpec() AwsCatalogTable_PartitionSpecPropertyOutputReference
	// Experimental.
	PartitionSpecInput() *AwsCatalogTable_PartitionSpecProperty
	// Experimental.
	Properties() *map[string]*string
	// Experimental.
	SetProperties(val *map[string]*string)
	// Experimental.
	PropertiesInput() *map[string]*string
	// Experimental.
	Schema() AwsCatalogTable_SchemaPropertyOutputReference
	// Experimental.
	SchemaInput() *AwsCatalogTable_SchemaProperty
	// Experimental.
	SortOrder() AwsCatalogTable_SortOrderPropertyOutputReference
	// Experimental.
	SortOrderInput() *AwsCatalogTable_SortOrderProperty
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
	PutPartitionSpec(value *AwsCatalogTable_PartitionSpecProperty)
	// Experimental.
	PutSchema(value *AwsCatalogTable_SchemaProperty)
	// Experimental.
	PutSortOrder(value *AwsCatalogTable_SortOrderProperty)
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

// The jsii proxy struct for AwsCatalogTable_IcebergTableInputPropertyOutputReference
type jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) InternalValue() *AwsCatalogTable_IcebergTableInputProperty {
	var returns *AwsCatalogTable_IcebergTableInputProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) PartitionSpec() AwsCatalogTable_PartitionSpecPropertyOutputReference {
	var returns AwsCatalogTable_PartitionSpecPropertyOutputReference
	_jsii_.Get(
		j,
		"partitionSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) PartitionSpecInput() *AwsCatalogTable_PartitionSpecProperty {
	var returns *AwsCatalogTable_PartitionSpecProperty
	_jsii_.Get(
		j,
		"partitionSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) Schema() AwsCatalogTable_SchemaPropertyOutputReference {
	var returns AwsCatalogTable_SchemaPropertyOutputReference
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) SchemaInput() *AwsCatalogTable_SchemaProperty {
	var returns *AwsCatalogTable_SchemaProperty
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) SortOrder() AwsCatalogTable_SortOrderPropertyOutputReference {
	var returns AwsCatalogTable_SortOrderPropertyOutputReference
	_jsii_.Get(
		j,
		"sortOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) SortOrderInput() *AwsCatalogTable_SortOrderProperty {
	var returns *AwsCatalogTable_SortOrderProperty
	_jsii_.Get(
		j,
		"sortOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCatalogTable_IcebergTableInputPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCatalogTable_IcebergTableInputPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCatalogTable_IcebergTableInputPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsCatalogTable.IcebergTableInputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCatalogTable_IcebergTableInputPropertyOutputReference_Override(a AwsCatalogTable_IcebergTableInputPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsCatalogTable.IcebergTableInputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference)SetInternalValue(val *AwsCatalogTable_IcebergTableInputProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) PutPartitionSpec(value *AwsCatalogTable_PartitionSpecProperty) {
	if err := a.validatePutPartitionSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPartitionSpec",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) PutSchema(value *AwsCatalogTable_SchemaProperty) {
	if err := a.validatePutSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchema",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) PutSortOrder(value *AwsCatalogTable_SortOrderProperty) {
	if err := a.validatePutSortOrderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSortOrder",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) ResetPartitionSpec() {
	_jsii_.InvokeVoid(
		a,
		"resetPartitionSpec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		a,
		"resetProperties",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) ResetSortOrder() {
	_jsii_.InvokeVoid(
		a,
		"resetSortOrder",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCatalogTable_IcebergTableInputPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

