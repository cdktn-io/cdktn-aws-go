package awsendusermessaging

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsendusermessaging/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsendusermessaging/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEmailTemplate_EmailTemplatePropertyOutputReference interface {
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
	DefaultSubstitutions() *string
	// Experimental.
	SetDefaultSubstitutions(val *string)
	// Experimental.
	DefaultSubstitutionsInput() *string
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	Header() TfEmailTemplate_HeaderPropertyList
	// Experimental.
	HeaderInput() interface{}
	// Experimental.
	HtmlPart() *string
	// Experimental.
	SetHtmlPart(val *string)
	// Experimental.
	HtmlPartInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	RecommenderId() *string
	// Experimental.
	SetRecommenderId(val *string)
	// Experimental.
	RecommenderIdInput() *string
	// Experimental.
	Subject() *string
	// Experimental.
	SetSubject(val *string)
	// Experimental.
	SubjectInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TextPart() *string
	// Experimental.
	SetTextPart(val *string)
	// Experimental.
	TextPartInput() *string
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
	PutHeader(value interface{})
	// Experimental.
	ResetDefaultSubstitutions()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetHeader()
	// Experimental.
	ResetHtmlPart()
	// Experimental.
	ResetRecommenderId()
	// Experimental.
	ResetSubject()
	// Experimental.
	ResetTextPart()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEmailTemplate_EmailTemplatePropertyOutputReference
type jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) DefaultSubstitutions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSubstitutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) DefaultSubstitutionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSubstitutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) Header() TfEmailTemplate_HeaderPropertyList {
	var returns TfEmailTemplate_HeaderPropertyList
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) HtmlPart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlPart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) HtmlPartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlPartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) RecommenderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommenderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) RecommenderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommenderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) SubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) TextPart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textPart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) TextPartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textPartInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEmailTemplate_EmailTemplatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfEmailTemplate_EmailTemplatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEmailTemplate_EmailTemplatePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-end-user-messaging.TfEmailTemplate.EmailTemplatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEmailTemplate_EmailTemplatePropertyOutputReference_Override(t TfEmailTemplate_EmailTemplatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-end-user-messaging.TfEmailTemplate.EmailTemplatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference)SetDefaultSubstitutions(val *string) {
	if err := j.validateSetDefaultSubstitutionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSubstitutions",
		val,
	)
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference)SetHtmlPart(val *string) {
	if err := j.validateSetHtmlPartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"htmlPart",
		val,
	)
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference)SetRecommenderId(val *string) {
	if err := j.validateSetRecommenderIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recommenderId",
		val,
	)
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference)SetSubject(val *string) {
	if err := j.validateSetSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subject",
		val,
	)
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference)SetTextPart(val *string) {
	if err := j.validateSetTextPartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textPart",
		val,
	)
}

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) PutHeader(value interface{}) {
	if err := t.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHeader",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) ResetDefaultSubstitutions() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultSubstitutions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) ResetHtmlPart() {
	_jsii_.InvokeVoid(
		t,
		"resetHtmlPart",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) ResetRecommenderId() {
	_jsii_.InvokeVoid(
		t,
		"resetRecommenderId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) ResetSubject() {
	_jsii_.InvokeVoid(
		t,
		"resetSubject",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) ResetTextPart() {
	_jsii_.InvokeVoid(
		t,
		"resetTextPart",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEmailTemplate_EmailTemplatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

