package awsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecr/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecr/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference interface {
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
	// Experimental.
	CountNumber() *float64
	// Experimental.
	SetCountNumber(val *float64)
	// Experimental.
	CountNumberInput() *float64
	// Experimental.
	CountType() *string
	// Experimental.
	SetCountType(val *string)
	// Experimental.
	CountTypeInput() *string
	// Experimental.
	CountUnit() *string
	// Experimental.
	SetCountUnit(val *string)
	// Experimental.
	CountUnitInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	StorageClass() *string
	// Experimental.
	SetStorageClass(val *string)
	// Experimental.
	StorageClassInput() *string
	// Experimental.
	TagPatternList() *[]*string
	// Experimental.
	SetTagPatternList(val *[]*string)
	// Experimental.
	TagPatternListInput() *[]*string
	// Experimental.
	TagPrefixList() *[]*string
	// Experimental.
	SetTagPrefixList(val *[]*string)
	// Experimental.
	TagPrefixListInput() *[]*string
	// Experimental.
	TagStatus() *string
	// Experimental.
	SetTagStatus(val *string)
	// Experimental.
	TagStatusInput() *string
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
	ResetCountUnit()
	// Experimental.
	ResetStorageClass()
	// Experimental.
	ResetTagPatternList()
	// Experimental.
	ResetTagPrefixList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference
type jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) CountNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) CountNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) CountType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) CountTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) CountUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) CountUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) TagPatternList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagPatternList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) TagPatternListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagPatternListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) TagPrefixList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagPrefixList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) TagPrefixListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagPrefixListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) TagStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) TagStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecr.DataAwsEcrLifecyclePolicyDocument.SelectionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference_Override(d DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecr.DataAwsEcrLifecyclePolicyDocument.SelectionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference)SetCountNumber(val *float64) {
	if err := j.validateSetCountNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countNumber",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference)SetCountType(val *string) {
	if err := j.validateSetCountTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countType",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference)SetCountUnit(val *string) {
	if err := j.validateSetCountUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countUnit",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference)SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference)SetTagPatternList(val *[]*string) {
	if err := j.validateSetTagPatternListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagPatternList",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference)SetTagPrefixList(val *[]*string) {
	if err := j.validateSetTagPrefixListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagPrefixList",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference)SetTagStatus(val *string) {
	if err := j.validateSetTagStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagStatus",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) ResetCountUnit() {
	_jsii_.InvokeVoid(
		d,
		"resetCountUnit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) ResetStorageClass() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageClass",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) ResetTagPatternList() {
	_jsii_.InvokeVoid(
		d,
		"resetTagPatternList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) ResetTagPrefixList() {
	_jsii_.InvokeVoid(
		d,
		"resetTagPrefixList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEcrLifecyclePolicyDocument_SelectionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

