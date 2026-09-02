package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference interface {
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
	FieldId() *float64
	// Experimental.
	SetFieldId(val *float64)
	// Experimental.
	FieldIdInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	SourceId() *float64
	// Experimental.
	SetSourceId(val *float64)
	// Experimental.
	SourceIdInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Transform() *string
	// Experimental.
	SetTransform(val *string)
	// Experimental.
	TransformInput() *string
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
	ResetFieldId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference
type jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) FieldId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fieldId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) FieldIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fieldIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) SourceId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) SourceIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) Transform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) TransformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.TfCatalogTable.OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference_Override(t TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.TfCatalogTable.OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference)SetFieldId(val *float64) {
	if err := j.validateSetFieldIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldId",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference)SetSourceId(val *float64) {
	if err := j.validateSetSourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceId",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference)SetTransform(val *string) {
	if err := j.validateSetTransformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transform",
		val,
	)
}

func (t *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) ResetFieldId() {
	_jsii_.InvokeVoid(
		t,
		"resetFieldId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

