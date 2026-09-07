package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/glue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/glue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCatalogTable_RepresentationsPropertyOutputReference interface {
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
	Dialect() *string
	// Experimental.
	SetDialect(val *string)
	// Experimental.
	DialectInput() *string
	// Experimental.
	DialectVersion() *string
	// Experimental.
	SetDialectVersion(val *string)
	// Experimental.
	DialectVersionInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ValidationConnection() *string
	// Experimental.
	SetValidationConnection(val *string)
	// Experimental.
	ValidationConnectionInput() *string
	// Experimental.
	ViewExpandedText() *string
	// Experimental.
	SetViewExpandedText(val *string)
	// Experimental.
	ViewExpandedTextInput() *string
	// Experimental.
	ViewOriginalText() *string
	// Experimental.
	SetViewOriginalText(val *string)
	// Experimental.
	ViewOriginalTextInput() *string
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
	ResetDialect()
	// Experimental.
	ResetDialectVersion()
	// Experimental.
	ResetValidationConnection()
	// Experimental.
	ResetViewExpandedText()
	// Experimental.
	ResetViewOriginalText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCatalogTable_RepresentationsPropertyOutputReference
type jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) Dialect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) DialectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) DialectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) DialectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) ValidationConnection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) ValidationConnectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationConnectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) ViewExpandedText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewExpandedText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) ViewExpandedTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewExpandedTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) ViewOriginalText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewOriginalText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) ViewOriginalTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewOriginalTextInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCatalogTable_RepresentationsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsCatalogTable_RepresentationsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCatalogTable_RepresentationsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsCatalogTable.RepresentationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCatalogTable_RepresentationsPropertyOutputReference_Override(a AwsCatalogTable_RepresentationsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsCatalogTable.RepresentationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference)SetDialect(val *string) {
	if err := j.validateSetDialectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dialect",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference)SetDialectVersion(val *string) {
	if err := j.validateSetDialectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dialectVersion",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference)SetValidationConnection(val *string) {
	if err := j.validateSetValidationConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validationConnection",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference)SetViewExpandedText(val *string) {
	if err := j.validateSetViewExpandedTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewExpandedText",
		val,
	)
}

func (j *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference)SetViewOriginalText(val *string) {
	if err := j.validateSetViewOriginalTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewOriginalText",
		val,
	)
}

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) ResetDialect() {
	_jsii_.InvokeVoid(
		a,
		"resetDialect",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) ResetDialectVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetDialectVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) ResetValidationConnection() {
	_jsii_.InvokeVoid(
		a,
		"resetValidationConnection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) ResetViewExpandedText() {
	_jsii_.InvokeVoid(
		a,
		"resetViewExpandedText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) ResetViewOriginalText() {
	_jsii_.InvokeVoid(
		a,
		"resetViewOriginalText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCatalogTable_RepresentationsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

