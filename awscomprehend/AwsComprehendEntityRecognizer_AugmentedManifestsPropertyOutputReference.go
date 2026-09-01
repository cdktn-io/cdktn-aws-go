package awscomprehend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscomprehend/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscomprehend/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference interface {
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

// The jsii proxy struct for AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference
type jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) AnnotationDataS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationDataS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) AnnotationDataS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationDataS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) AttributeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) AttributeNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) DocumentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) DocumentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) S3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) S3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) SourceDocumentsS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDocumentsS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) SourceDocumentsS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDocumentsS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) Split() *string {
	var returns *string
	_jsii_.Get(
		j,
		"split",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) SplitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-comprehend.AwsComprehendEntityRecognizer.AugmentedManifestsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference_Override(a AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-comprehend.AwsComprehendEntityRecognizer.AugmentedManifestsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetAnnotationDataS3Uri(val *string) {
	if err := j.validateSetAnnotationDataS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotationDataS3Uri",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetAttributeNames(val *[]*string) {
	if err := j.validateSetAttributeNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeNames",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetDocumentType(val *string) {
	if err := j.validateSetDocumentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentType",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetS3Uri(val *string) {
	if err := j.validateSetS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Uri",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetSourceDocumentsS3Uri(val *string) {
	if err := j.validateSetSourceDocumentsS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDocumentsS3Uri",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetSplit(val *string) {
	if err := j.validateSetSplitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"split",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) ResetAnnotationDataS3Uri() {
	_jsii_.InvokeVoid(
		a,
		"resetAnnotationDataS3Uri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) ResetDocumentType() {
	_jsii_.InvokeVoid(
		a,
		"resetDocumentType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) ResetSourceDocumentsS3Uri() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceDocumentsS3Uri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) ResetSplit() {
	_jsii_.InvokeVoid(
		a,
		"resetSplit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsComprehendEntityRecognizer_AugmentedManifestsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

