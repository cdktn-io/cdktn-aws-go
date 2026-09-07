//go:build no_runtime_type_checking

package s3

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsBuckets_BucketsPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsBuckets_BucketsPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsBuckets_BucketsPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsBuckets_BucketsPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsBuckets_BucketsPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsBuckets_BucketsPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsBuckets_BucketsPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

