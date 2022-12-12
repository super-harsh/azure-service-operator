// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220301

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_Image_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Image_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImage_Spec_ARM, Image_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImage_Spec_ARM runs a test to see if a specific instance of Image_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImage_Spec_ARM(subject Image_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Image_Spec_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Image_Spec_ARM instances for property testing - lazily instantiated by Image_Spec_ARMGenerator()
var image_Spec_ARMGenerator gopter.Gen

// Image_Spec_ARMGenerator returns a generator of Image_Spec_ARM instances for property testing.
// We first initialize image_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Image_Spec_ARMGenerator() gopter.Gen {
	if image_Spec_ARMGenerator != nil {
		return image_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImage_Spec_ARM(generators)
	image_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Image_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImage_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForImage_Spec_ARM(generators)
	image_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Image_Spec_ARM{}), generators)

	return image_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForImage_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImage_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImage_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImage_Spec_ARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_ARMGenerator())
	gens["Properties"] = gen.PtrOf(ImageProperties_ARMGenerator())
}

func Test_ExtendedLocation_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExtendedLocation_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExtendedLocation_ARM, ExtendedLocation_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExtendedLocation_ARM runs a test to see if a specific instance of ExtendedLocation_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExtendedLocation_ARM(subject ExtendedLocation_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExtendedLocation_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ExtendedLocation_ARM instances for property testing - lazily instantiated by
// ExtendedLocation_ARMGenerator()
var extendedLocation_ARMGenerator gopter.Gen

// ExtendedLocation_ARMGenerator returns a generator of ExtendedLocation_ARM instances for property testing.
func ExtendedLocation_ARMGenerator() gopter.Gen {
	if extendedLocation_ARMGenerator != nil {
		return extendedLocation_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExtendedLocation_ARM(generators)
	extendedLocation_ARMGenerator = gen.Struct(reflect.TypeOf(ExtendedLocation_ARM{}), generators)

	return extendedLocation_ARMGenerator
}

// AddIndependentPropertyGeneratorsForExtendedLocation_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExtendedLocation_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(ExtendedLocationType_EdgeZone))
}

func Test_ImageProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageProperties_ARM, ImageProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageProperties_ARM runs a test to see if a specific instance of ImageProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageProperties_ARM(subject ImageProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageProperties_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ImageProperties_ARM instances for property testing - lazily instantiated by
// ImageProperties_ARMGenerator()
var imageProperties_ARMGenerator gopter.Gen

// ImageProperties_ARMGenerator returns a generator of ImageProperties_ARM instances for property testing.
// We first initialize imageProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageProperties_ARMGenerator() gopter.Gen {
	if imageProperties_ARMGenerator != nil {
		return imageProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageProperties_ARM(generators)
	imageProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ImageProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForImageProperties_ARM(generators)
	imageProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ImageProperties_ARM{}), generators)

	return imageProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForImageProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageProperties_ARM(gens map[string]gopter.Gen) {
	gens["HyperVGeneration"] = gen.PtrOf(gen.OneConstOf(HyperVGenerationType_V1, HyperVGenerationType_V2))
}

// AddRelatedPropertyGeneratorsForImageProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageProperties_ARM(gens map[string]gopter.Gen) {
	gens["SourceVirtualMachine"] = gen.PtrOf(SubResource_ARMGenerator())
	gens["StorageProfile"] = gen.PtrOf(ImageStorageProfile_ARMGenerator())
}

func Test_ImageStorageProfile_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageStorageProfile_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageStorageProfile_ARM, ImageStorageProfile_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageStorageProfile_ARM runs a test to see if a specific instance of ImageStorageProfile_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageStorageProfile_ARM(subject ImageStorageProfile_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageStorageProfile_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ImageStorageProfile_ARM instances for property testing - lazily instantiated by
// ImageStorageProfile_ARMGenerator()
var imageStorageProfile_ARMGenerator gopter.Gen

// ImageStorageProfile_ARMGenerator returns a generator of ImageStorageProfile_ARM instances for property testing.
// We first initialize imageStorageProfile_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageStorageProfile_ARMGenerator() gopter.Gen {
	if imageStorageProfile_ARMGenerator != nil {
		return imageStorageProfile_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageStorageProfile_ARM(generators)
	imageStorageProfile_ARMGenerator = gen.Struct(reflect.TypeOf(ImageStorageProfile_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageStorageProfile_ARM(generators)
	AddRelatedPropertyGeneratorsForImageStorageProfile_ARM(generators)
	imageStorageProfile_ARMGenerator = gen.Struct(reflect.TypeOf(ImageStorageProfile_ARM{}), generators)

	return imageStorageProfile_ARMGenerator
}

// AddIndependentPropertyGeneratorsForImageStorageProfile_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageStorageProfile_ARM(gens map[string]gopter.Gen) {
	gens["ZoneResilient"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForImageStorageProfile_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageStorageProfile_ARM(gens map[string]gopter.Gen) {
	gens["DataDisks"] = gen.SliceOf(ImageDataDisk_ARMGenerator())
	gens["OsDisk"] = gen.PtrOf(ImageOSDisk_ARMGenerator())
}

func Test_SubResource_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResource_ARM, SubResource_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResource_ARM runs a test to see if a specific instance of SubResource_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResource_ARM(subject SubResource_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SubResource_ARM instances for property testing - lazily instantiated by SubResource_ARMGenerator()
var subResource_ARMGenerator gopter.Gen

// SubResource_ARMGenerator returns a generator of SubResource_ARM instances for property testing.
func SubResource_ARMGenerator() gopter.Gen {
	if subResource_ARMGenerator != nil {
		return subResource_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubResource_ARM(generators)
	subResource_ARMGenerator = gen.Struct(reflect.TypeOf(SubResource_ARM{}), generators)

	return subResource_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSubResource_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubResource_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_ImageDataDisk_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageDataDisk_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageDataDisk_ARM, ImageDataDisk_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageDataDisk_ARM runs a test to see if a specific instance of ImageDataDisk_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageDataDisk_ARM(subject ImageDataDisk_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageDataDisk_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ImageDataDisk_ARM instances for property testing - lazily instantiated by ImageDataDisk_ARMGenerator()
var imageDataDisk_ARMGenerator gopter.Gen

// ImageDataDisk_ARMGenerator returns a generator of ImageDataDisk_ARM instances for property testing.
// We first initialize imageDataDisk_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageDataDisk_ARMGenerator() gopter.Gen {
	if imageDataDisk_ARMGenerator != nil {
		return imageDataDisk_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageDataDisk_ARM(generators)
	imageDataDisk_ARMGenerator = gen.Struct(reflect.TypeOf(ImageDataDisk_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageDataDisk_ARM(generators)
	AddRelatedPropertyGeneratorsForImageDataDisk_ARM(generators)
	imageDataDisk_ARMGenerator = gen.Struct(reflect.TypeOf(ImageDataDisk_ARM{}), generators)

	return imageDataDisk_ARMGenerator
}

// AddIndependentPropertyGeneratorsForImageDataDisk_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageDataDisk_ARM(gens map[string]gopter.Gen) {
	gens["BlobUri"] = gen.PtrOf(gen.AlphaString())
	gens["Caching"] = gen.PtrOf(gen.OneConstOf(ImageDataDisk_Caching_None, ImageDataDisk_Caching_ReadOnly, ImageDataDisk_Caching_ReadWrite))
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["Lun"] = gen.PtrOf(gen.Int())
	gens["StorageAccountType"] = gen.PtrOf(gen.OneConstOf(
		StorageAccountType_PremiumV2_LRS,
		StorageAccountType_Premium_LRS,
		StorageAccountType_Premium_ZRS,
		StorageAccountType_StandardSSD_LRS,
		StorageAccountType_StandardSSD_ZRS,
		StorageAccountType_Standard_LRS,
		StorageAccountType_UltraSSD_LRS))
}

// AddRelatedPropertyGeneratorsForImageDataDisk_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageDataDisk_ARM(gens map[string]gopter.Gen) {
	gens["DiskEncryptionSet"] = gen.PtrOf(SubResource_ARMGenerator())
	gens["ManagedDisk"] = gen.PtrOf(SubResource_ARMGenerator())
	gens["Snapshot"] = gen.PtrOf(SubResource_ARMGenerator())
}

func Test_ImageOSDisk_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageOSDisk_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageOSDisk_ARM, ImageOSDisk_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageOSDisk_ARM runs a test to see if a specific instance of ImageOSDisk_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageOSDisk_ARM(subject ImageOSDisk_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageOSDisk_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ImageOSDisk_ARM instances for property testing - lazily instantiated by ImageOSDisk_ARMGenerator()
var imageOSDisk_ARMGenerator gopter.Gen

// ImageOSDisk_ARMGenerator returns a generator of ImageOSDisk_ARM instances for property testing.
// We first initialize imageOSDisk_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageOSDisk_ARMGenerator() gopter.Gen {
	if imageOSDisk_ARMGenerator != nil {
		return imageOSDisk_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageOSDisk_ARM(generators)
	imageOSDisk_ARMGenerator = gen.Struct(reflect.TypeOf(ImageOSDisk_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageOSDisk_ARM(generators)
	AddRelatedPropertyGeneratorsForImageOSDisk_ARM(generators)
	imageOSDisk_ARMGenerator = gen.Struct(reflect.TypeOf(ImageOSDisk_ARM{}), generators)

	return imageOSDisk_ARMGenerator
}

// AddIndependentPropertyGeneratorsForImageOSDisk_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageOSDisk_ARM(gens map[string]gopter.Gen) {
	gens["BlobUri"] = gen.PtrOf(gen.AlphaString())
	gens["Caching"] = gen.PtrOf(gen.OneConstOf(ImageOSDisk_Caching_None, ImageOSDisk_Caching_ReadOnly, ImageOSDisk_Caching_ReadWrite))
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["OsState"] = gen.PtrOf(gen.OneConstOf(ImageOSDisk_OsState_Generalized, ImageOSDisk_OsState_Specialized))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(ImageOSDisk_OsType_Linux, ImageOSDisk_OsType_Windows))
	gens["StorageAccountType"] = gen.PtrOf(gen.OneConstOf(
		StorageAccountType_PremiumV2_LRS,
		StorageAccountType_Premium_LRS,
		StorageAccountType_Premium_ZRS,
		StorageAccountType_StandardSSD_LRS,
		StorageAccountType_StandardSSD_ZRS,
		StorageAccountType_Standard_LRS,
		StorageAccountType_UltraSSD_LRS))
}

// AddRelatedPropertyGeneratorsForImageOSDisk_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageOSDisk_ARM(gens map[string]gopter.Gen) {
	gens["DiskEncryptionSet"] = gen.PtrOf(SubResource_ARMGenerator())
	gens["ManagedDisk"] = gen.PtrOf(SubResource_ARMGenerator())
	gens["Snapshot"] = gen.PtrOf(SubResource_ARMGenerator())
}
