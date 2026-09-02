package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCatalogTable_SchemaReferencePropertyOutputReference interface {
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
	InternalValue() *TfCatalogTable_SchemaReferenceProperty
	// Experimental.
	SetInternalValue(val *TfCatalogTable_SchemaReferenceProperty)
	// Experimental.
	SchemaId() TfCatalogTable_SchemaIdPropertyOutputReference
	// Experimental.
	SchemaIdInput() *TfCatalogTable_SchemaIdProperty
	// Experimental.
	SchemaVersionId() *string
	// Experimental.
	SetSchemaVersionId(val *string)
	// Experimental.
	SchemaVersionIdInput() *string
	// Experimental.
	SchemaVersionNumber() *float64
	// Experimental.
	SetSchemaVersionNumber(val *float64)
	// Experimental.
	SchemaVersionNumberInput() *float64
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
	PutSchemaId(value *TfCatalogTable_SchemaIdProperty)
	// Experimental.
	ResetSchemaId()
	// Experimental.
	ResetSchemaVersionId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCatalogTable_SchemaReferencePropertyOutputReference
type jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) InternalValue() *TfCatalogTable_SchemaReferenceProperty {
	var returns *TfCatalogTable_SchemaReferenceProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) SchemaId() TfCatalogTable_SchemaIdPropertyOutputReference {
	var returns TfCatalogTable_SchemaIdPropertyOutputReference
	_jsii_.Get(
		j,
		"schemaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) SchemaIdInput() *TfCatalogTable_SchemaIdProperty {
	var returns *TfCatalogTable_SchemaIdProperty
	_jsii_.Get(
		j,
		"schemaIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) SchemaVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) SchemaVersionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaVersionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) SchemaVersionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schemaVersionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) SchemaVersionNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schemaVersionNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCatalogTable_SchemaReferencePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCatalogTable_SchemaReferencePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCatalogTable_SchemaReferencePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.TfCatalogTable.SchemaReferencePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCatalogTable_SchemaReferencePropertyOutputReference_Override(t TfCatalogTable_SchemaReferencePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.TfCatalogTable.SchemaReferencePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference)SetInternalValue(val *TfCatalogTable_SchemaReferenceProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference)SetSchemaVersionId(val *string) {
	if err := j.validateSetSchemaVersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaVersionId",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference)SetSchemaVersionNumber(val *float64) {
	if err := j.validateSetSchemaVersionNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaVersionNumber",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) PutSchemaId(value *TfCatalogTable_SchemaIdProperty) {
	if err := t.validatePutSchemaIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSchemaId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) ResetSchemaId() {
	_jsii_.InvokeVoid(
		t,
		"resetSchemaId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) ResetSchemaVersionId() {
	_jsii_.InvokeVoid(
		t,
		"resetSchemaVersionId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCatalogTable_SchemaReferencePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

