package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/glue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/glue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCatalogTable_SchemaReferencePropertyOutputReference interface {
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
	InternalValue() *AwsCatalogTable_SchemaReferenceProperty
	// Experimental.
	SetInternalValue(val *AwsCatalogTable_SchemaReferenceProperty)
	// Experimental.
	SchemaId() AwsCatalogTable_SchemaIdPropertyOutputReference
	// Experimental.
	SchemaIdInput() *AwsCatalogTable_SchemaIdProperty
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
	PutSchemaId(value *AwsCatalogTable_SchemaIdProperty)
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

// The jsii proxy struct for AwsCatalogTable_SchemaReferencePropertyOutputReference
type jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) InternalValue() *AwsCatalogTable_SchemaReferenceProperty {
	var returns *AwsCatalogTable_SchemaReferenceProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) SchemaId() AwsCatalogTable_SchemaIdPropertyOutputReference {
	var returns AwsCatalogTable_SchemaIdPropertyOutputReference
	_jsii_.Get(
		j,
		"schemaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) SchemaIdInput() *AwsCatalogTable_SchemaIdProperty {
	var returns *AwsCatalogTable_SchemaIdProperty
	_jsii_.Get(
		j,
		"schemaIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) SchemaVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) SchemaVersionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaVersionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) SchemaVersionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schemaVersionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) SchemaVersionNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schemaVersionNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCatalogTable_SchemaReferencePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCatalogTable_SchemaReferencePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCatalogTable_SchemaReferencePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsCatalogTable.SchemaReferencePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCatalogTable_SchemaReferencePropertyOutputReference_Override(a AwsCatalogTable_SchemaReferencePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsCatalogTable.SchemaReferencePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference)SetInternalValue(val *AwsCatalogTable_SchemaReferenceProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference)SetSchemaVersionId(val *string) {
	if err := j.validateSetSchemaVersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaVersionId",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference)SetSchemaVersionNumber(val *float64) {
	if err := j.validateSetSchemaVersionNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaVersionNumber",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) PutSchemaId(value *AwsCatalogTable_SchemaIdProperty) {
	if err := a.validatePutSchemaIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchemaId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) ResetSchemaId() {
	_jsii_.InvokeVoid(
		a,
		"resetSchemaId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) ResetSchemaVersionId() {
	_jsii_.InvokeVoid(
		a,
		"resetSchemaVersionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCatalogTable_SchemaReferencePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

