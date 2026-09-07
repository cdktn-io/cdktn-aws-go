package comprehend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/comprehend/jsii"

	"github.com/cdktn-io/cdktn-aws-go/comprehend/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEntityRecognizer_InputDataConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Annotations() AwsEntityRecognizer_AnnotationsPropertyOutputReference
	// Experimental.
	AnnotationsInput() *AwsEntityRecognizer_AnnotationsProperty
	// Experimental.
	AugmentedManifests() AwsEntityRecognizer_AugmentedManifestsPropertyList
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
	Documents() AwsEntityRecognizer_DocumentsPropertyOutputReference
	// Experimental.
	DocumentsInput() *AwsEntityRecognizer_DocumentsProperty
	// Experimental.
	EntityList() AwsEntityRecognizer_EntityListPropertyOutputReference
	// Experimental.
	EntityListInput() *AwsEntityRecognizer_EntityListProperty
	// Experimental.
	EntityTypes() AwsEntityRecognizer_EntityTypesPropertyList
	// Experimental.
	EntityTypesInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsEntityRecognizer_InputDataConfigProperty
	// Experimental.
	SetInternalValue(val *AwsEntityRecognizer_InputDataConfigProperty)
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
	PutAnnotations(value *AwsEntityRecognizer_AnnotationsProperty)
	// Experimental.
	PutAugmentedManifests(value interface{})
	// Experimental.
	PutDocuments(value *AwsEntityRecognizer_DocumentsProperty)
	// Experimental.
	PutEntityList(value *AwsEntityRecognizer_EntityListProperty)
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

// The jsii proxy struct for AwsEntityRecognizer_InputDataConfigPropertyOutputReference
type jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) Annotations() AwsEntityRecognizer_AnnotationsPropertyOutputReference {
	var returns AwsEntityRecognizer_AnnotationsPropertyOutputReference
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) AnnotationsInput() *AwsEntityRecognizer_AnnotationsProperty {
	var returns *AwsEntityRecognizer_AnnotationsProperty
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) AugmentedManifests() AwsEntityRecognizer_AugmentedManifestsPropertyList {
	var returns AwsEntityRecognizer_AugmentedManifestsPropertyList
	_jsii_.Get(
		j,
		"augmentedManifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) AugmentedManifestsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"augmentedManifestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) DataFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) DataFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) Documents() AwsEntityRecognizer_DocumentsPropertyOutputReference {
	var returns AwsEntityRecognizer_DocumentsPropertyOutputReference
	_jsii_.Get(
		j,
		"documents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) DocumentsInput() *AwsEntityRecognizer_DocumentsProperty {
	var returns *AwsEntityRecognizer_DocumentsProperty
	_jsii_.Get(
		j,
		"documentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) EntityList() AwsEntityRecognizer_EntityListPropertyOutputReference {
	var returns AwsEntityRecognizer_EntityListPropertyOutputReference
	_jsii_.Get(
		j,
		"entityList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) EntityListInput() *AwsEntityRecognizer_EntityListProperty {
	var returns *AwsEntityRecognizer_EntityListProperty
	_jsii_.Get(
		j,
		"entityListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) EntityTypes() AwsEntityRecognizer_EntityTypesPropertyList {
	var returns AwsEntityRecognizer_EntityTypesPropertyList
	_jsii_.Get(
		j,
		"entityTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) EntityTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entityTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) InternalValue() *AwsEntityRecognizer_InputDataConfigProperty {
	var returns *AwsEntityRecognizer_InputDataConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEntityRecognizer_InputDataConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEntityRecognizer_InputDataConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEntityRecognizer_InputDataConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-comprehend.AwsEntityRecognizer.InputDataConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEntityRecognizer_InputDataConfigPropertyOutputReference_Override(a AwsEntityRecognizer_InputDataConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-comprehend.AwsEntityRecognizer.InputDataConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference)SetDataFormat(val *string) {
	if err := j.validateSetDataFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFormat",
		val,
	)
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference)SetInternalValue(val *AwsEntityRecognizer_InputDataConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) PutAnnotations(value *AwsEntityRecognizer_AnnotationsProperty) {
	if err := a.validatePutAnnotationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAnnotations",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) PutAugmentedManifests(value interface{}) {
	if err := a.validatePutAugmentedManifestsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAugmentedManifests",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) PutDocuments(value *AwsEntityRecognizer_DocumentsProperty) {
	if err := a.validatePutDocumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDocuments",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) PutEntityList(value *AwsEntityRecognizer_EntityListProperty) {
	if err := a.validatePutEntityListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEntityList",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) PutEntityTypes(value interface{}) {
	if err := a.validatePutEntityTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEntityTypes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) ResetAnnotations() {
	_jsii_.InvokeVoid(
		a,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) ResetAugmentedManifests() {
	_jsii_.InvokeVoid(
		a,
		"resetAugmentedManifests",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) ResetDataFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetDataFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) ResetDocuments() {
	_jsii_.InvokeVoid(
		a,
		"resetDocuments",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) ResetEntityList() {
	_jsii_.InvokeVoid(
		a,
		"resetEntityList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEntityRecognizer_InputDataConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

