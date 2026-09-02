package awscomprehend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscomprehend/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscomprehend/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEntityRecognizer_AugmentedManifestsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AnnotationDataS3Uri() *string
	// Experimental.
	SetAnnotationDataS3Uri(val *string)
	// Experimental.
	AnnotationDataS3UriInput() *string
	// Experimental.
	AttributeNames() *[]*string
	// Experimental.
	SetAttributeNames(val *[]*string)
	// Experimental.
	AttributeNamesInput() *[]*string
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	S3Uri() *string
	// Experimental.
	SetS3Uri(val *string)
	// Experimental.
	S3UriInput() *string
	// Experimental.
	SourceDocumentsS3Uri() *string
	// Experimental.
	SetSourceDocumentsS3Uri(val *string)
	// Experimental.
	SourceDocumentsS3UriInput() *string
	// Experimental.
	Split() *string
	// Experimental.
	SetSplit(val *string)
	// Experimental.
	SplitInput() *string
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
	ResetAnnotationDataS3Uri()
	// Experimental.
	ResetDocumentType()
	// Experimental.
	ResetSourceDocumentsS3Uri()
	// Experimental.
	ResetSplit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEntityRecognizer_AugmentedManifestsPropertyOutputReference
type jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) AnnotationDataS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationDataS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) AnnotationDataS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationDataS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) AttributeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) AttributeNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) DocumentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) DocumentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) S3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) S3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) SourceDocumentsS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDocumentsS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) SourceDocumentsS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDocumentsS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) Split() *string {
	var returns *string
	_jsii_.Get(
		j,
		"split",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) SplitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEntityRecognizer_AugmentedManifestsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfEntityRecognizer_AugmentedManifestsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEntityRecognizer_AugmentedManifestsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-comprehend.TfEntityRecognizer.AugmentedManifestsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEntityRecognizer_AugmentedManifestsPropertyOutputReference_Override(t TfEntityRecognizer_AugmentedManifestsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-comprehend.TfEntityRecognizer.AugmentedManifestsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetAnnotationDataS3Uri(val *string) {
	if err := j.validateSetAnnotationDataS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotationDataS3Uri",
		val,
	)
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetAttributeNames(val *[]*string) {
	if err := j.validateSetAttributeNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeNames",
		val,
	)
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetDocumentType(val *string) {
	if err := j.validateSetDocumentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentType",
		val,
	)
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetS3Uri(val *string) {
	if err := j.validateSetS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Uri",
		val,
	)
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetSourceDocumentsS3Uri(val *string) {
	if err := j.validateSetSourceDocumentsS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDocumentsS3Uri",
		val,
	)
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetSplit(val *string) {
	if err := j.validateSetSplitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"split",
		val,
	)
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) ResetAnnotationDataS3Uri() {
	_jsii_.InvokeVoid(
		t,
		"resetAnnotationDataS3Uri",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) ResetDocumentType() {
	_jsii_.InvokeVoid(
		t,
		"resetDocumentType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) ResetSourceDocumentsS3Uri() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceDocumentsS3Uri",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) ResetSplit() {
	_jsii_.InvokeVoid(
		t,
		"resetSplit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEntityRecognizer_AugmentedManifestsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

