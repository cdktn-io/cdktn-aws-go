package awsebs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsebs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsebs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSnapshotImport_ClientDataPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Comment() *string
	// Experimental.
	SetComment(val *string)
	// Experimental.
	CommentInput() *string
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
	InternalValue() *TfSnapshotImport_ClientDataProperty
	// Experimental.
	SetInternalValue(val *TfSnapshotImport_ClientDataProperty)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UploadEnd() *string
	// Experimental.
	SetUploadEnd(val *string)
	// Experimental.
	UploadEndInput() *string
	// Experimental.
	UploadSize() *float64
	// Experimental.
	SetUploadSize(val *float64)
	// Experimental.
	UploadSizeInput() *float64
	// Experimental.
	UploadStart() *string
	// Experimental.
	SetUploadStart(val *string)
	// Experimental.
	UploadStartInput() *string
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
	ResetComment()
	// Experimental.
	ResetUploadEnd()
	// Experimental.
	ResetUploadSize()
	// Experimental.
	ResetUploadStart()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfSnapshotImport_ClientDataPropertyOutputReference
type jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) InternalValue() *TfSnapshotImport_ClientDataProperty {
	var returns *TfSnapshotImport_ClientDataProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) UploadEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) UploadEndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) UploadSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uploadSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) UploadSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uploadSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) UploadStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) UploadStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadStartInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfSnapshotImport_ClientDataPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfSnapshotImport_ClientDataPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfSnapshotImport_ClientDataPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ebs.TfSnapshotImport.ClientDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfSnapshotImport_ClientDataPropertyOutputReference_Override(t TfSnapshotImport_ClientDataPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ebs.TfSnapshotImport.ClientDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference)SetInternalValue(val *TfSnapshotImport_ClientDataProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference)SetUploadEnd(val *string) {
	if err := j.validateSetUploadEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadEnd",
		val,
	)
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference)SetUploadSize(val *float64) {
	if err := j.validateSetUploadSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadSize",
		val,
	)
}

func (j *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference)SetUploadStart(val *string) {
	if err := j.validateSetUploadStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadStart",
		val,
	)
}

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		t,
		"resetComment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) ResetUploadEnd() {
	_jsii_.InvokeVoid(
		t,
		"resetUploadEnd",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) ResetUploadSize() {
	_jsii_.InvokeVoid(
		t,
		"resetUploadSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) ResetUploadStart() {
	_jsii_.InvokeVoid(
		t,
		"resetUploadStart",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfSnapshotImport_ClientDataPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

