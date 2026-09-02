package awscloudwatchlogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudwatchlogs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudwatchlogs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTransformer_TransformerConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AddKeys() TfTransformer_AddKeysPropertyList
	// Experimental.
	AddKeysInput() interface{}
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
	CopyValue() TfTransformer_CopyValuePropertyList
	// Experimental.
	CopyValueInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Csv() TfTransformer_CsvPropertyList
	// Experimental.
	CsvInput() interface{}
	// Experimental.
	DateTimeConverter() TfTransformer_DateTimeConverterPropertyList
	// Experimental.
	DateTimeConverterInput() interface{}
	// Experimental.
	DeleteKeys() TfTransformer_DeleteKeysPropertyList
	// Experimental.
	DeleteKeysInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	Grok() TfTransformer_GrokPropertyList
	// Experimental.
	GrokInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	ListToMap() TfTransformer_ListToMapPropertyList
	// Experimental.
	ListToMapInput() interface{}
	// Experimental.
	LowerCaseString() TfTransformer_LowerCaseStringPropertyList
	// Experimental.
	LowerCaseStringInput() interface{}
	// Experimental.
	MoveKeys() TfTransformer_MoveKeysPropertyList
	// Experimental.
	MoveKeysInput() interface{}
	// Experimental.
	ParseCloudfront() TfTransformer_ParseCloudfrontPropertyList
	// Experimental.
	ParseCloudfrontInput() interface{}
	// Experimental.
	ParseJson() TfTransformer_ParseJsonPropertyList
	// Experimental.
	ParseJsonInput() interface{}
	// Experimental.
	ParseKeyValue() TfTransformer_ParseKeyValuePropertyList
	// Experimental.
	ParseKeyValueInput() interface{}
	// Experimental.
	ParsePostgres() TfTransformer_ParsePostgresPropertyList
	// Experimental.
	ParsePostgresInput() interface{}
	// Experimental.
	ParseRoute53() TfTransformer_ParseRoute53PropertyList
	// Experimental.
	ParseRoute53Input() interface{}
	// Experimental.
	ParseToOcsf() TfTransformer_ParseToOcsfPropertyList
	// Experimental.
	ParseToOcsfInput() interface{}
	// Experimental.
	ParseVpc() TfTransformer_ParseVpcPropertyList
	// Experimental.
	ParseVpcInput() interface{}
	// Experimental.
	ParseWaf() TfTransformer_ParseWafPropertyList
	// Experimental.
	ParseWafInput() interface{}
	// Experimental.
	RenameKeys() TfTransformer_RenameKeysPropertyList
	// Experimental.
	RenameKeysInput() interface{}
	// Experimental.
	SplitString() TfTransformer_SplitStringPropertyList
	// Experimental.
	SplitStringInput() interface{}
	// Experimental.
	SubstituteString() TfTransformer_SubstituteStringPropertyList
	// Experimental.
	SubstituteStringInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TrimString() TfTransformer_TrimStringPropertyList
	// Experimental.
	TrimStringInput() interface{}
	// Experimental.
	TypeConverter() TfTransformer_TypeConverterPropertyList
	// Experimental.
	TypeConverterInput() interface{}
	// Experimental.
	UpperCaseString() TfTransformer_UpperCaseStringPropertyList
	// Experimental.
	UpperCaseStringInput() interface{}
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
	PutAddKeys(value interface{})
	// Experimental.
	PutCopyValue(value interface{})
	// Experimental.
	PutCsv(value interface{})
	// Experimental.
	PutDateTimeConverter(value interface{})
	// Experimental.
	PutDeleteKeys(value interface{})
	// Experimental.
	PutGrok(value interface{})
	// Experimental.
	PutListToMap(value interface{})
	// Experimental.
	PutLowerCaseString(value interface{})
	// Experimental.
	PutMoveKeys(value interface{})
	// Experimental.
	PutParseCloudfront(value interface{})
	// Experimental.
	PutParseJson(value interface{})
	// Experimental.
	PutParseKeyValue(value interface{})
	// Experimental.
	PutParsePostgres(value interface{})
	// Experimental.
	PutParseRoute53(value interface{})
	// Experimental.
	PutParseToOcsf(value interface{})
	// Experimental.
	PutParseVpc(value interface{})
	// Experimental.
	PutParseWaf(value interface{})
	// Experimental.
	PutRenameKeys(value interface{})
	// Experimental.
	PutSplitString(value interface{})
	// Experimental.
	PutSubstituteString(value interface{})
	// Experimental.
	PutTrimString(value interface{})
	// Experimental.
	PutTypeConverter(value interface{})
	// Experimental.
	PutUpperCaseString(value interface{})
	// Experimental.
	ResetAddKeys()
	// Experimental.
	ResetCopyValue()
	// Experimental.
	ResetCsv()
	// Experimental.
	ResetDateTimeConverter()
	// Experimental.
	ResetDeleteKeys()
	// Experimental.
	ResetGrok()
	// Experimental.
	ResetListToMap()
	// Experimental.
	ResetLowerCaseString()
	// Experimental.
	ResetMoveKeys()
	// Experimental.
	ResetParseCloudfront()
	// Experimental.
	ResetParseJson()
	// Experimental.
	ResetParseKeyValue()
	// Experimental.
	ResetParsePostgres()
	// Experimental.
	ResetParseRoute53()
	// Experimental.
	ResetParseToOcsf()
	// Experimental.
	ResetParseVpc()
	// Experimental.
	ResetParseWaf()
	// Experimental.
	ResetRenameKeys()
	// Experimental.
	ResetSplitString()
	// Experimental.
	ResetSubstituteString()
	// Experimental.
	ResetTrimString()
	// Experimental.
	ResetTypeConverter()
	// Experimental.
	ResetUpperCaseString()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTransformer_TransformerConfigPropertyOutputReference
type jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) AddKeys() TfTransformer_AddKeysPropertyList {
	var returns TfTransformer_AddKeysPropertyList
	_jsii_.Get(
		j,
		"addKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) AddKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) CopyValue() TfTransformer_CopyValuePropertyList {
	var returns TfTransformer_CopyValuePropertyList
	_jsii_.Get(
		j,
		"copyValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) CopyValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) Csv() TfTransformer_CsvPropertyList {
	var returns TfTransformer_CsvPropertyList
	_jsii_.Get(
		j,
		"csv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) CsvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"csvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) DateTimeConverter() TfTransformer_DateTimeConverterPropertyList {
	var returns TfTransformer_DateTimeConverterPropertyList
	_jsii_.Get(
		j,
		"dateTimeConverter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) DateTimeConverterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dateTimeConverterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) DeleteKeys() TfTransformer_DeleteKeysPropertyList {
	var returns TfTransformer_DeleteKeysPropertyList
	_jsii_.Get(
		j,
		"deleteKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) DeleteKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) Grok() TfTransformer_GrokPropertyList {
	var returns TfTransformer_GrokPropertyList
	_jsii_.Get(
		j,
		"grok",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) GrokInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grokInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ListToMap() TfTransformer_ListToMapPropertyList {
	var returns TfTransformer_ListToMapPropertyList
	_jsii_.Get(
		j,
		"listToMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ListToMapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"listToMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) LowerCaseString() TfTransformer_LowerCaseStringPropertyList {
	var returns TfTransformer_LowerCaseStringPropertyList
	_jsii_.Get(
		j,
		"lowerCaseString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) LowerCaseStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lowerCaseStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) MoveKeys() TfTransformer_MoveKeysPropertyList {
	var returns TfTransformer_MoveKeysPropertyList
	_jsii_.Get(
		j,
		"moveKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) MoveKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"moveKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ParseCloudfront() TfTransformer_ParseCloudfrontPropertyList {
	var returns TfTransformer_ParseCloudfrontPropertyList
	_jsii_.Get(
		j,
		"parseCloudfront",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ParseCloudfrontInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseCloudfrontInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ParseJson() TfTransformer_ParseJsonPropertyList {
	var returns TfTransformer_ParseJsonPropertyList
	_jsii_.Get(
		j,
		"parseJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ParseJsonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ParseKeyValue() TfTransformer_ParseKeyValuePropertyList {
	var returns TfTransformer_ParseKeyValuePropertyList
	_jsii_.Get(
		j,
		"parseKeyValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ParseKeyValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseKeyValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ParsePostgres() TfTransformer_ParsePostgresPropertyList {
	var returns TfTransformer_ParsePostgresPropertyList
	_jsii_.Get(
		j,
		"parsePostgres",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ParsePostgresInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parsePostgresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ParseRoute53() TfTransformer_ParseRoute53PropertyList {
	var returns TfTransformer_ParseRoute53PropertyList
	_jsii_.Get(
		j,
		"parseRoute53",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ParseRoute53Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseRoute53Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ParseToOcsf() TfTransformer_ParseToOcsfPropertyList {
	var returns TfTransformer_ParseToOcsfPropertyList
	_jsii_.Get(
		j,
		"parseToOcsf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ParseToOcsfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseToOcsfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ParseVpc() TfTransformer_ParseVpcPropertyList {
	var returns TfTransformer_ParseVpcPropertyList
	_jsii_.Get(
		j,
		"parseVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ParseVpcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseVpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ParseWaf() TfTransformer_ParseWafPropertyList {
	var returns TfTransformer_ParseWafPropertyList
	_jsii_.Get(
		j,
		"parseWaf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ParseWafInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseWafInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) RenameKeys() TfTransformer_RenameKeysPropertyList {
	var returns TfTransformer_RenameKeysPropertyList
	_jsii_.Get(
		j,
		"renameKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) RenameKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renameKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) SplitString() TfTransformer_SplitStringPropertyList {
	var returns TfTransformer_SplitStringPropertyList
	_jsii_.Get(
		j,
		"splitString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) SplitStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splitStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) SubstituteString() TfTransformer_SubstituteStringPropertyList {
	var returns TfTransformer_SubstituteStringPropertyList
	_jsii_.Get(
		j,
		"substituteString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) SubstituteStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"substituteStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) TrimString() TfTransformer_TrimStringPropertyList {
	var returns TfTransformer_TrimStringPropertyList
	_jsii_.Get(
		j,
		"trimString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) TrimStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trimStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) TypeConverter() TfTransformer_TypeConverterPropertyList {
	var returns TfTransformer_TypeConverterPropertyList
	_jsii_.Get(
		j,
		"typeConverter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) TypeConverterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"typeConverterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) UpperCaseString() TfTransformer_UpperCaseStringPropertyList {
	var returns TfTransformer_UpperCaseStringPropertyList
	_jsii_.Get(
		j,
		"upperCaseString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) UpperCaseStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upperCaseStringInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTransformer_TransformerConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfTransformer_TransformerConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTransformer_TransformerConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-logs.TfTransformer.TransformerConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTransformer_TransformerConfigPropertyOutputReference_Override(t TfTransformer_TransformerConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-logs.TfTransformer.TransformerConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutAddKeys(value interface{}) {
	if err := t.validatePutAddKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAddKeys",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutCopyValue(value interface{}) {
	if err := t.validatePutCopyValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCopyValue",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutCsv(value interface{}) {
	if err := t.validatePutCsvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCsv",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutDateTimeConverter(value interface{}) {
	if err := t.validatePutDateTimeConverterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDateTimeConverter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutDeleteKeys(value interface{}) {
	if err := t.validatePutDeleteKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeleteKeys",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutGrok(value interface{}) {
	if err := t.validatePutGrokParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGrok",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutListToMap(value interface{}) {
	if err := t.validatePutListToMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putListToMap",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutLowerCaseString(value interface{}) {
	if err := t.validatePutLowerCaseStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLowerCaseString",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutMoveKeys(value interface{}) {
	if err := t.validatePutMoveKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMoveKeys",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutParseCloudfront(value interface{}) {
	if err := t.validatePutParseCloudfrontParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putParseCloudfront",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutParseJson(value interface{}) {
	if err := t.validatePutParseJsonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putParseJson",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutParseKeyValue(value interface{}) {
	if err := t.validatePutParseKeyValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putParseKeyValue",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutParsePostgres(value interface{}) {
	if err := t.validatePutParsePostgresParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putParsePostgres",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutParseRoute53(value interface{}) {
	if err := t.validatePutParseRoute53Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putParseRoute53",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutParseToOcsf(value interface{}) {
	if err := t.validatePutParseToOcsfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putParseToOcsf",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutParseVpc(value interface{}) {
	if err := t.validatePutParseVpcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putParseVpc",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutParseWaf(value interface{}) {
	if err := t.validatePutParseWafParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putParseWaf",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutRenameKeys(value interface{}) {
	if err := t.validatePutRenameKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRenameKeys",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutSplitString(value interface{}) {
	if err := t.validatePutSplitStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSplitString",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutSubstituteString(value interface{}) {
	if err := t.validatePutSubstituteStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSubstituteString",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutTrimString(value interface{}) {
	if err := t.validatePutTrimStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTrimString",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutTypeConverter(value interface{}) {
	if err := t.validatePutTypeConverterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTypeConverter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) PutUpperCaseString(value interface{}) {
	if err := t.validatePutUpperCaseStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUpperCaseString",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetAddKeys() {
	_jsii_.InvokeVoid(
		t,
		"resetAddKeys",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetCopyValue() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetCsv() {
	_jsii_.InvokeVoid(
		t,
		"resetCsv",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetDateTimeConverter() {
	_jsii_.InvokeVoid(
		t,
		"resetDateTimeConverter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetDeleteKeys() {
	_jsii_.InvokeVoid(
		t,
		"resetDeleteKeys",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetGrok() {
	_jsii_.InvokeVoid(
		t,
		"resetGrok",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetListToMap() {
	_jsii_.InvokeVoid(
		t,
		"resetListToMap",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetLowerCaseString() {
	_jsii_.InvokeVoid(
		t,
		"resetLowerCaseString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetMoveKeys() {
	_jsii_.InvokeVoid(
		t,
		"resetMoveKeys",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetParseCloudfront() {
	_jsii_.InvokeVoid(
		t,
		"resetParseCloudfront",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetParseJson() {
	_jsii_.InvokeVoid(
		t,
		"resetParseJson",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetParseKeyValue() {
	_jsii_.InvokeVoid(
		t,
		"resetParseKeyValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetParsePostgres() {
	_jsii_.InvokeVoid(
		t,
		"resetParsePostgres",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetParseRoute53() {
	_jsii_.InvokeVoid(
		t,
		"resetParseRoute53",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetParseToOcsf() {
	_jsii_.InvokeVoid(
		t,
		"resetParseToOcsf",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetParseVpc() {
	_jsii_.InvokeVoid(
		t,
		"resetParseVpc",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetParseWaf() {
	_jsii_.InvokeVoid(
		t,
		"resetParseWaf",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetRenameKeys() {
	_jsii_.InvokeVoid(
		t,
		"resetRenameKeys",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetSplitString() {
	_jsii_.InvokeVoid(
		t,
		"resetSplitString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetSubstituteString() {
	_jsii_.InvokeVoid(
		t,
		"resetSubstituteString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetTrimString() {
	_jsii_.InvokeVoid(
		t,
		"resetTrimString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetTypeConverter() {
	_jsii_.InvokeVoid(
		t,
		"resetTypeConverter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ResetUpperCaseString() {
	_jsii_.InvokeVoid(
		t,
		"resetUpperCaseString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTransformer_TransformerConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

