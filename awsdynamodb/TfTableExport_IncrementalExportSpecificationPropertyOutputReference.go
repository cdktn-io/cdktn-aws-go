package awsdynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdynamodb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdynamodb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTableExport_IncrementalExportSpecificationPropertyOutputReference interface {
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
	ExportFromTime() *string
	// Experimental.
	SetExportFromTime(val *string)
	// Experimental.
	ExportFromTimeInput() *string
	// Experimental.
	ExportToTime() *string
	// Experimental.
	SetExportToTime(val *string)
	// Experimental.
	ExportToTimeInput() *string
	// Experimental.
	ExportViewType() *string
	// Experimental.
	SetExportViewType(val *string)
	// Experimental.
	ExportViewTypeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfTableExport_IncrementalExportSpecificationProperty
	// Experimental.
	SetInternalValue(val *TfTableExport_IncrementalExportSpecificationProperty)
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
	ResetExportFromTime()
	// Experimental.
	ResetExportToTime()
	// Experimental.
	ResetExportViewType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTableExport_IncrementalExportSpecificationPropertyOutputReference
type jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) ExportFromTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportFromTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) ExportFromTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportFromTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) ExportToTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) ExportToTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) ExportViewType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportViewType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) ExportViewTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportViewTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) InternalValue() *TfTableExport_IncrementalExportSpecificationProperty {
	var returns *TfTableExport_IncrementalExportSpecificationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTableExport_IncrementalExportSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTableExport_IncrementalExportSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTableExport_IncrementalExportSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dynamodb.TfTableExport.IncrementalExportSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTableExport_IncrementalExportSpecificationPropertyOutputReference_Override(t TfTableExport_IncrementalExportSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dynamodb.TfTableExport.IncrementalExportSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference)SetExportFromTime(val *string) {
	if err := j.validateSetExportFromTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportFromTime",
		val,
	)
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference)SetExportToTime(val *string) {
	if err := j.validateSetExportToTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToTime",
		val,
	)
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference)SetExportViewType(val *string) {
	if err := j.validateSetExportViewTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportViewType",
		val,
	)
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference)SetInternalValue(val *TfTableExport_IncrementalExportSpecificationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) ResetExportFromTime() {
	_jsii_.InvokeVoid(
		t,
		"resetExportFromTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) ResetExportToTime() {
	_jsii_.InvokeVoid(
		t,
		"resetExportToTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) ResetExportViewType() {
	_jsii_.InvokeVoid(
		t,
		"resetExportViewType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTableExport_IncrementalExportSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

