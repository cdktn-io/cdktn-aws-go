package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGlueCatalogTable_SkewedInfoPropertyOutputReference interface {
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
	InternalValue() *AwsGlueCatalogTable_SkewedInfoProperty
	// Experimental.
	SetInternalValue(val *AwsGlueCatalogTable_SkewedInfoProperty)
	// Experimental.
	SkewedColumnNames() *[]*string
	// Experimental.
	SetSkewedColumnNames(val *[]*string)
	// Experimental.
	SkewedColumnNamesInput() *[]*string
	// Experimental.
	SkewedColumnValueLocationMaps() *map[string]*string
	// Experimental.
	SetSkewedColumnValueLocationMaps(val *map[string]*string)
	// Experimental.
	SkewedColumnValueLocationMapsInput() *map[string]*string
	// Experimental.
	SkewedColumnValues() *[]*string
	// Experimental.
	SetSkewedColumnValues(val *[]*string)
	// Experimental.
	SkewedColumnValuesInput() *[]*string
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
	ResetSkewedColumnNames()
	// Experimental.
	ResetSkewedColumnValueLocationMaps()
	// Experimental.
	ResetSkewedColumnValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsGlueCatalogTable_SkewedInfoPropertyOutputReference
type jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) InternalValue() *AwsGlueCatalogTable_SkewedInfoProperty {
	var returns *AwsGlueCatalogTable_SkewedInfoProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) SkewedColumnNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) SkewedColumnNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) SkewedColumnValueLocationMaps() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"skewedColumnValueLocationMaps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) SkewedColumnValueLocationMapsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"skewedColumnValueLocationMapsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) SkewedColumnValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) SkewedColumnValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"skewedColumnValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGlueCatalogTable_SkewedInfoPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsGlueCatalogTable_SkewedInfoPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGlueCatalogTable_SkewedInfoPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsGlueCatalogTable.SkewedInfoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGlueCatalogTable_SkewedInfoPropertyOutputReference_Override(a AwsGlueCatalogTable_SkewedInfoPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsGlueCatalogTable.SkewedInfoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference)SetInternalValue(val *AwsGlueCatalogTable_SkewedInfoProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference)SetSkewedColumnNames(val *[]*string) {
	if err := j.validateSetSkewedColumnNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skewedColumnNames",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference)SetSkewedColumnValueLocationMaps(val *map[string]*string) {
	if err := j.validateSetSkewedColumnValueLocationMapsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skewedColumnValueLocationMaps",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference)SetSkewedColumnValues(val *[]*string) {
	if err := j.validateSetSkewedColumnValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skewedColumnValues",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) ResetSkewedColumnNames() {
	_jsii_.InvokeVoid(
		a,
		"resetSkewedColumnNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) ResetSkewedColumnValueLocationMaps() {
	_jsii_.InvokeVoid(
		a,
		"resetSkewedColumnValueLocationMaps",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) ResetSkewedColumnValues() {
	_jsii_.InvokeVoid(
		a,
		"resetSkewedColumnValues",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGlueCatalogTable_SkewedInfoPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

