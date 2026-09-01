package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference interface {
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
	Doc() *string
	// Experimental.
	SetDoc(val *string)
	// Experimental.
	DocInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	Id() *float64
	// Experimental.
	SetId(val *float64)
	// Experimental.
	IdInput() *float64
	// Experimental.
	InitialDefault() *string
	// Experimental.
	SetInitialDefault(val *string)
	// Experimental.
	InitialDefaultInput() *string
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
	Required() interface{}
	// Experimental.
	SetRequired(val interface{})
	// Experimental.
	RequiredInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
	// Experimental.
	WriteDefault() *string
	// Experimental.
	SetWriteDefault(val *string)
	// Experimental.
	WriteDefaultInput() *string
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
	ResetDoc()
	// Experimental.
	ResetInitialDefault()
	// Experimental.
	ResetWriteDefault()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference
type jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) Doc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"doc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) DocInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) IdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) InitialDefault() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) InitialDefaultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) Required() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) RequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) WriteDefault() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) WriteDefaultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeDefaultInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsGlueCatalogTable.OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference_Override(a AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsGlueCatalogTable.OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference)SetDoc(val *string) {
	if err := j.validateSetDocParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"doc",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference)SetId(val *float64) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference)SetInitialDefault(val *string) {
	if err := j.validateSetInitialDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialDefault",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference)SetRequired(val interface{}) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference)SetWriteDefault(val *string) {
	if err := j.validateSetWriteDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeDefault",
		val,
	)
}

func (a *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) ResetDoc() {
	_jsii_.InvokeVoid(
		a,
		"resetDoc",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) ResetInitialDefault() {
	_jsii_.InvokeVoid(
		a,
		"resetInitialDefault",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) ResetWriteDefault() {
	_jsii_.InvokeVoid(
		a,
		"resetWriteDefault",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGlueCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSchemaFieldsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

