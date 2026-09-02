package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() *TfProject_ServiceCatalogProvisioningDetailsProperty
	// Experimental.
	SetInternalValue(val *TfProject_ServiceCatalogProvisioningDetailsProperty)
	// Experimental.
	PathId() *string
	// Experimental.
	SetPathId(val *string)
	// Experimental.
	PathIdInput() *string
	// Experimental.
	ProductId() *string
	// Experimental.
	SetProductId(val *string)
	// Experimental.
	ProductIdInput() *string
	// Experimental.
	ProvisioningArtifactId() *string
	// Experimental.
	SetProvisioningArtifactId(val *string)
	// Experimental.
	ProvisioningArtifactIdInput() *string
	// Experimental.
	ProvisioningParameter() TfProject_ProvisioningParameterPropertyList
	// Experimental.
	ProvisioningParameterInput() interface{}
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
	PutProvisioningParameter(value interface{})
	// Experimental.
	ResetPathId()
	// Experimental.
	ResetProvisioningArtifactId()
	// Experimental.
	ResetProvisioningParameter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference
type jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) InternalValue() *TfProject_ServiceCatalogProvisioningDetailsProperty {
	var returns *TfProject_ServiceCatalogProvisioningDetailsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) PathId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) PathIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) ProductIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) ProvisioningArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningArtifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) ProvisioningArtifactIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningArtifactIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) ProvisioningParameter() TfProject_ProvisioningParameterPropertyList {
	var returns TfProject_ProvisioningParameterPropertyList
	_jsii_.Get(
		j,
		"provisioningParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) ProvisioningParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisioningParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfProject_ServiceCatalogProvisioningDetailsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfProject.ServiceCatalogProvisioningDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference_Override(t TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfProject.ServiceCatalogProvisioningDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference)SetInternalValue(val *TfProject_ServiceCatalogProvisioningDetailsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference)SetPathId(val *string) {
	if err := j.validateSetPathIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathId",
		val,
	)
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference)SetProductId(val *string) {
	if err := j.validateSetProductIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productId",
		val,
	)
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference)SetProvisioningArtifactId(val *string) {
	if err := j.validateSetProvisioningArtifactIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningArtifactId",
		val,
	)
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) PutProvisioningParameter(value interface{}) {
	if err := t.validatePutProvisioningParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProvisioningParameter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) ResetPathId() {
	_jsii_.InvokeVoid(
		t,
		"resetPathId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) ResetProvisioningArtifactId() {
	_jsii_.InvokeVoid(
		t,
		"resetProvisioningArtifactId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) ResetProvisioningParameter() {
	_jsii_.InvokeVoid(
		t,
		"resetProvisioningParameter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfProject_ServiceCatalogProvisioningDetailsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

