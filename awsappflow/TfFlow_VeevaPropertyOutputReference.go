package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFlow_VeevaPropertyOutputReference interface {
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
	DocumentType() *string
	// Experimental.
	SetDocumentType(val *string)
	// Experimental.
	DocumentTypeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	IncludeAllVersions() interface{}
	// Experimental.
	SetIncludeAllVersions(val interface{})
	// Experimental.
	IncludeAllVersionsInput() interface{}
	// Experimental.
	IncludeRenditions() interface{}
	// Experimental.
	SetIncludeRenditions(val interface{})
	// Experimental.
	IncludeRenditionsInput() interface{}
	// Experimental.
	IncludeSourceFiles() interface{}
	// Experimental.
	SetIncludeSourceFiles(val interface{})
	// Experimental.
	IncludeSourceFilesInput() interface{}
	// Experimental.
	InternalValue() *TfFlow_VeevaProperty
	// Experimental.
	SetInternalValue(val *TfFlow_VeevaProperty)
	// Experimental.
	Object() *string
	// Experimental.
	SetObject(val *string)
	// Experimental.
	ObjectInput() *string
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
	ResetDocumentType()
	// Experimental.
	ResetIncludeAllVersions()
	// Experimental.
	ResetIncludeRenditions()
	// Experimental.
	ResetIncludeSourceFiles()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFlow_VeevaPropertyOutputReference
type jsiiProxy_TfFlow_VeevaPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference) DocumentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference) DocumentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference) IncludeAllVersions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeAllVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference) IncludeAllVersionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeAllVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference) IncludeRenditions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeRenditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference) IncludeRenditionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeRenditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference) IncludeSourceFiles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSourceFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference) IncludeSourceFilesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSourceFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference) InternalValue() *TfFlow_VeevaProperty {
	var returns *TfFlow_VeevaProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference) Object() *string {
	var returns *string
	_jsii_.Get(
		j,
		"object",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference) ObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFlow_VeevaPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFlow_VeevaPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFlow_VeevaPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFlow_VeevaPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.VeevaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFlow_VeevaPropertyOutputReference_Override(t TfFlow_VeevaPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.VeevaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference)SetDocumentType(val *string) {
	if err := j.validateSetDocumentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentType",
		val,
	)
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference)SetIncludeAllVersions(val interface{}) {
	if err := j.validateSetIncludeAllVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeAllVersions",
		val,
	)
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference)SetIncludeRenditions(val interface{}) {
	if err := j.validateSetIncludeRenditionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeRenditions",
		val,
	)
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference)SetIncludeSourceFiles(val interface{}) {
	if err := j.validateSetIncludeSourceFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSourceFiles",
		val,
	)
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference)SetInternalValue(val *TfFlow_VeevaProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference)SetObject(val *string) {
	if err := j.validateSetObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"object",
		val,
	)
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFlow_VeevaPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) ResetDocumentType() {
	_jsii_.InvokeVoid(
		t,
		"resetDocumentType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) ResetIncludeAllVersions() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeAllVersions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) ResetIncludeRenditions() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeRenditions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) ResetIncludeSourceFiles() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeSourceFiles",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFlow_VeevaPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

