package awscomprehend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscomprehend/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscomprehend/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEntityRecognizer_InputDataConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Annotations() TfEntityRecognizer_AnnotationsPropertyOutputReference
	// Experimental.
	AnnotationsInput() *TfEntityRecognizer_AnnotationsProperty
	// Experimental.
	AugmentedManifests() TfEntityRecognizer_AugmentedManifestsPropertyList
	// Experimental.
	AugmentedManifestsInput() interface{}
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
	DataFormat() *string
	// Experimental.
	SetDataFormat(val *string)
	// Experimental.
	DataFormatInput() *string
	// Experimental.
	Documents() TfEntityRecognizer_DocumentsPropertyOutputReference
	// Experimental.
	DocumentsInput() *TfEntityRecognizer_DocumentsProperty
	// Experimental.
	EntityList() TfEntityRecognizer_EntityListPropertyOutputReference
	// Experimental.
	EntityListInput() *TfEntityRecognizer_EntityListProperty
	// Experimental.
	EntityTypes() TfEntityRecognizer_EntityTypesPropertyList
	// Experimental.
	EntityTypesInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfEntityRecognizer_InputDataConfigProperty
	// Experimental.
	SetInternalValue(val *TfEntityRecognizer_InputDataConfigProperty)
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
	PutAnnotations(value *TfEntityRecognizer_AnnotationsProperty)
	// Experimental.
	PutAugmentedManifests(value interface{})
	// Experimental.
	PutDocuments(value *TfEntityRecognizer_DocumentsProperty)
	// Experimental.
	PutEntityList(value *TfEntityRecognizer_EntityListProperty)
	// Experimental.
	PutEntityTypes(value interface{})
	// Experimental.
	ResetAnnotations()
	// Experimental.
	ResetAugmentedManifests()
	// Experimental.
	ResetDataFormat()
	// Experimental.
	ResetDocuments()
	// Experimental.
	ResetEntityList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEntityRecognizer_InputDataConfigPropertyOutputReference
type jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) Annotations() TfEntityRecognizer_AnnotationsPropertyOutputReference {
	var returns TfEntityRecognizer_AnnotationsPropertyOutputReference
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) AnnotationsInput() *TfEntityRecognizer_AnnotationsProperty {
	var returns *TfEntityRecognizer_AnnotationsProperty
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) AugmentedManifests() TfEntityRecognizer_AugmentedManifestsPropertyList {
	var returns TfEntityRecognizer_AugmentedManifestsPropertyList
	_jsii_.Get(
		j,
		"augmentedManifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) AugmentedManifestsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"augmentedManifestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) DataFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) DataFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) Documents() TfEntityRecognizer_DocumentsPropertyOutputReference {
	var returns TfEntityRecognizer_DocumentsPropertyOutputReference
	_jsii_.Get(
		j,
		"documents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) DocumentsInput() *TfEntityRecognizer_DocumentsProperty {
	var returns *TfEntityRecognizer_DocumentsProperty
	_jsii_.Get(
		j,
		"documentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) EntityList() TfEntityRecognizer_EntityListPropertyOutputReference {
	var returns TfEntityRecognizer_EntityListPropertyOutputReference
	_jsii_.Get(
		j,
		"entityList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) EntityListInput() *TfEntityRecognizer_EntityListProperty {
	var returns *TfEntityRecognizer_EntityListProperty
	_jsii_.Get(
		j,
		"entityListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) EntityTypes() TfEntityRecognizer_EntityTypesPropertyList {
	var returns TfEntityRecognizer_EntityTypesPropertyList
	_jsii_.Get(
		j,
		"entityTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) EntityTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entityTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) InternalValue() *TfEntityRecognizer_InputDataConfigProperty {
	var returns *TfEntityRecognizer_InputDataConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEntityRecognizer_InputDataConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEntityRecognizer_InputDataConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEntityRecognizer_InputDataConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-comprehend.TfEntityRecognizer.InputDataConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEntityRecognizer_InputDataConfigPropertyOutputReference_Override(t TfEntityRecognizer_InputDataConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-comprehend.TfEntityRecognizer.InputDataConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference)SetDataFormat(val *string) {
	if err := j.validateSetDataFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFormat",
		val,
	)
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference)SetInternalValue(val *TfEntityRecognizer_InputDataConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) PutAnnotations(value *TfEntityRecognizer_AnnotationsProperty) {
	if err := t.validatePutAnnotationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAnnotations",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) PutAugmentedManifests(value interface{}) {
	if err := t.validatePutAugmentedManifestsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAugmentedManifests",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) PutDocuments(value *TfEntityRecognizer_DocumentsProperty) {
	if err := t.validatePutDocumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDocuments",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) PutEntityList(value *TfEntityRecognizer_EntityListProperty) {
	if err := t.validatePutEntityListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEntityList",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) PutEntityTypes(value interface{}) {
	if err := t.validatePutEntityTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEntityTypes",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) ResetAnnotations() {
	_jsii_.InvokeVoid(
		t,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) ResetAugmentedManifests() {
	_jsii_.InvokeVoid(
		t,
		"resetAugmentedManifests",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) ResetDataFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetDataFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) ResetDocuments() {
	_jsii_.InvokeVoid(
		t,
		"resetDocuments",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) ResetEntityList() {
	_jsii_.InvokeVoid(
		t,
		"resetEntityList",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEntityRecognizer_InputDataConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

