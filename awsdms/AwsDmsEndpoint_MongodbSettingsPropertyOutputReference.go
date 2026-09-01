package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDmsEndpoint_MongodbSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthMechanism() *string
	// Experimental.
	SetAuthMechanism(val *string)
	// Experimental.
	AuthMechanismInput() *string
	// Experimental.
	AuthSource() *string
	// Experimental.
	SetAuthSource(val *string)
	// Experimental.
	AuthSourceInput() *string
	// Experimental.
	AuthType() *string
	// Experimental.
	SetAuthType(val *string)
	// Experimental.
	AuthTypeInput() *string
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
	DocsToInvestigate() *string
	// Experimental.
	SetDocsToInvestigate(val *string)
	// Experimental.
	DocsToInvestigateInput() *string
	// Experimental.
	ExtractDocId() *string
	// Experimental.
	SetExtractDocId(val *string)
	// Experimental.
	ExtractDocIdInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDmsEndpoint_MongodbSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsDmsEndpoint_MongodbSettingsProperty)
	// Experimental.
	NestingLevel() *string
	// Experimental.
	SetNestingLevel(val *string)
	// Experimental.
	NestingLevelInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UseUpdateLookup() interface{}
	// Experimental.
	SetUseUpdateLookup(val interface{})
	// Experimental.
	UseUpdateLookupInput() interface{}
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
	ResetAuthMechanism()
	// Experimental.
	ResetAuthSource()
	// Experimental.
	ResetAuthType()
	// Experimental.
	ResetDocsToInvestigate()
	// Experimental.
	ResetExtractDocId()
	// Experimental.
	ResetNestingLevel()
	// Experimental.
	ResetUseUpdateLookup()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDmsEndpoint_MongodbSettingsPropertyOutputReference
type jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) AuthMechanism() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMechanism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) AuthMechanismInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMechanismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) AuthSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) AuthSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) DocsToInvestigate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docsToInvestigate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) DocsToInvestigateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docsToInvestigateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) ExtractDocId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extractDocId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) ExtractDocIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extractDocIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) InternalValue() *AwsDmsEndpoint_MongodbSettingsProperty {
	var returns *AwsDmsEndpoint_MongodbSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) NestingLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nestingLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) NestingLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nestingLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) UseUpdateLookup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useUpdateLookup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) UseUpdateLookupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useUpdateLookupInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDmsEndpoint_MongodbSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDmsEndpoint_MongodbSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDmsEndpoint_MongodbSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.AwsDmsEndpoint.MongodbSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDmsEndpoint_MongodbSettingsPropertyOutputReference_Override(a AwsDmsEndpoint_MongodbSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.AwsDmsEndpoint.MongodbSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference)SetAuthMechanism(val *string) {
	if err := j.validateSetAuthMechanismParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authMechanism",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference)SetAuthSource(val *string) {
	if err := j.validateSetAuthSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authSource",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference)SetDocsToInvestigate(val *string) {
	if err := j.validateSetDocsToInvestigateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"docsToInvestigate",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference)SetExtractDocId(val *string) {
	if err := j.validateSetExtractDocIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extractDocId",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference)SetInternalValue(val *AwsDmsEndpoint_MongodbSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference)SetNestingLevel(val *string) {
	if err := j.validateSetNestingLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nestingLevel",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference)SetUseUpdateLookup(val interface{}) {
	if err := j.validateSetUseUpdateLookupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useUpdateLookup",
		val,
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) ResetAuthMechanism() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthMechanism",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) ResetAuthSource() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) ResetAuthType() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) ResetDocsToInvestigate() {
	_jsii_.InvokeVoid(
		a,
		"resetDocsToInvestigate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) ResetExtractDocId() {
	_jsii_.InvokeVoid(
		a,
		"resetExtractDocId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) ResetNestingLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetNestingLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) ResetUseUpdateLookup() {
	_jsii_.InvokeVoid(
		a,
		"resetUseUpdateLookup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDmsEndpoint_MongodbSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

