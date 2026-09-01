package awsendusermessaging

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsendusermessaging/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsendusermessaging/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference interface {
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
	Header() AwsPinpointEmailTemplate_HeaderPropertyList
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

// The jsii proxy struct for AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference
type jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) DefaultSubstitutions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSubstitutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) DefaultSubstitutionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSubstitutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) Header() AwsPinpointEmailTemplate_HeaderPropertyList {
	var returns AwsPinpointEmailTemplate_HeaderPropertyList
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) HtmlPart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlPart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) HtmlPartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlPartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) RecommenderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommenderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) RecommenderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommenderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) SubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) TextPart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textPart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) TextPartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textPartInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPinpointEmailTemplate_EmailTemplatePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-end-user-messaging.AwsPinpointEmailTemplate.EmailTemplatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference_Override(a AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-end-user-messaging.AwsPinpointEmailTemplate.EmailTemplatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference)SetDefaultSubstitutions(val *string) {
	if err := j.validateSetDefaultSubstitutionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSubstitutions",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference)SetHtmlPart(val *string) {
	if err := j.validateSetHtmlPartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"htmlPart",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference)SetRecommenderId(val *string) {
	if err := j.validateSetRecommenderIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recommenderId",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference)SetSubject(val *string) {
	if err := j.validateSetSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subject",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference)SetTextPart(val *string) {
	if err := j.validateSetTextPartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textPart",
		val,
	)
}

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) PutHeader(value interface{}) {
	if err := a.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHeader",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) ResetDefaultSubstitutions() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultSubstitutions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) ResetHtmlPart() {
	_jsii_.InvokeVoid(
		a,
		"resetHtmlPart",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) ResetRecommenderId() {
	_jsii_.InvokeVoid(
		a,
		"resetRecommenderId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) ResetSubject() {
	_jsii_.InvokeVoid(
		a,
		"resetSubject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) ResetTextPart() {
	_jsii_.InvokeVoid(
		a,
		"resetTextPart",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPinpointEmailTemplate_EmailTemplatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

