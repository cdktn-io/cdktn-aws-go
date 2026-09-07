package dynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/dynamodb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/dynamodb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTableExport_IncrementalExportSpecificationPropertyOutputReference interface {
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
	InternalValue() *AwsTableExport_IncrementalExportSpecificationProperty
	// Experimental.
	SetInternalValue(val *AwsTableExport_IncrementalExportSpecificationProperty)
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

// The jsii proxy struct for AwsTableExport_IncrementalExportSpecificationPropertyOutputReference
type jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) ExportFromTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportFromTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) ExportFromTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportFromTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) ExportToTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) ExportToTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) ExportViewType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportViewType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) ExportViewTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportViewTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) InternalValue() *AwsTableExport_IncrementalExportSpecificationProperty {
	var returns *AwsTableExport_IncrementalExportSpecificationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTableExport_IncrementalExportSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsTableExport_IncrementalExportSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTableExport_IncrementalExportSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dynamodb.AwsTableExport.IncrementalExportSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTableExport_IncrementalExportSpecificationPropertyOutputReference_Override(a AwsTableExport_IncrementalExportSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dynamodb.AwsTableExport.IncrementalExportSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference)SetExportFromTime(val *string) {
	if err := j.validateSetExportFromTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportFromTime",
		val,
	)
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference)SetExportToTime(val *string) {
	if err := j.validateSetExportToTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToTime",
		val,
	)
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference)SetExportViewType(val *string) {
	if err := j.validateSetExportViewTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportViewType",
		val,
	)
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference)SetInternalValue(val *AwsTableExport_IncrementalExportSpecificationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) ResetExportFromTime() {
	_jsii_.InvokeVoid(
		a,
		"resetExportFromTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) ResetExportToTime() {
	_jsii_.InvokeVoid(
		a,
		"resetExportToTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) ResetExportViewType() {
	_jsii_.InvokeVoid(
		a,
		"resetExportViewType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTableExport_IncrementalExportSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

