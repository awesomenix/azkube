// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package training

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.2/customvision/training"

type BaseClient = original.BaseClient
type Classifier = original.Classifier

const (
	Multiclass Classifier = original.Multiclass
	Multilabel Classifier = original.Multilabel
)

type DomainType = original.DomainType

const (
	Classification  DomainType = original.Classification
	ObjectDetection DomainType = original.ObjectDetection
)

type ExportFlavorModel = original.ExportFlavorModel

const (
	Linux   ExportFlavorModel = original.Linux
	ONNX10  ExportFlavorModel = original.ONNX10
	ONNX12  ExportFlavorModel = original.ONNX12
	Windows ExportFlavorModel = original.Windows
)

type ExportPlatformModel = original.ExportPlatformModel

const (
	CoreML     ExportPlatformModel = original.CoreML
	DockerFile ExportPlatformModel = original.DockerFile
	ONNX       ExportPlatformModel = original.ONNX
	TensorFlow ExportPlatformModel = original.TensorFlow
)

type ExportStatusModel = original.ExportStatusModel

const (
	Done      ExportStatusModel = original.Done
	Exporting ExportStatusModel = original.Exporting
	Failed    ExportStatusModel = original.Failed
)

type ImageCreateStatus = original.ImageCreateStatus

const (
	ErrorImageFormat                      ImageCreateStatus = original.ErrorImageFormat
	ErrorImageSize                        ImageCreateStatus = original.ErrorImageSize
	ErrorLimitExceed                      ImageCreateStatus = original.ErrorLimitExceed
	ErrorNegativeAndRegularTagOnSameImage ImageCreateStatus = original.ErrorNegativeAndRegularTagOnSameImage
	ErrorRegionLimitExceed                ImageCreateStatus = original.ErrorRegionLimitExceed
	ErrorSource                           ImageCreateStatus = original.ErrorSource
	ErrorStorage                          ImageCreateStatus = original.ErrorStorage
	ErrorTagLimitExceed                   ImageCreateStatus = original.ErrorTagLimitExceed
	ErrorUnknown                          ImageCreateStatus = original.ErrorUnknown
	OK                                    ImageCreateStatus = original.OK
	OKDuplicate                           ImageCreateStatus = original.OKDuplicate
)

type OrderBy = original.OrderBy

const (
	Newest    OrderBy = original.Newest
	Oldest    OrderBy = original.Oldest
	Suggested OrderBy = original.Suggested
)

type TagType = original.TagType

const (
	Negative TagType = original.Negative
	Regular  TagType = original.Regular
)

type BoundingBox = original.BoundingBox
type Domain = original.Domain
type Export = original.Export
type Image = original.Image
type ImageCreateResult = original.ImageCreateResult
type ImageCreateSummary = original.ImageCreateSummary
type ImageFileCreateBatch = original.ImageFileCreateBatch
type ImageFileCreateEntry = original.ImageFileCreateEntry
type ImageIDCreateBatch = original.ImageIDCreateBatch
type ImageIDCreateEntry = original.ImageIDCreateEntry
type ImagePerformance = original.ImagePerformance
type ImagePrediction = original.ImagePrediction
type ImageRegion = original.ImageRegion
type ImageRegionCreateBatch = original.ImageRegionCreateBatch
type ImageRegionCreateEntry = original.ImageRegionCreateEntry
type ImageRegionCreateResult = original.ImageRegionCreateResult
type ImageRegionCreateSummary = original.ImageRegionCreateSummary
type ImageRegionProposal = original.ImageRegionProposal
type ImageTag = original.ImageTag
type ImageTagCreateBatch = original.ImageTagCreateBatch
type ImageTagCreateEntry = original.ImageTagCreateEntry
type ImageTagCreateSummary = original.ImageTagCreateSummary
type ImageURL = original.ImageURL
type ImageURLCreateBatch = original.ImageURLCreateBatch
type ImageURLCreateEntry = original.ImageURLCreateEntry
type Int32 = original.Int32
type Iteration = original.Iteration
type IterationPerformance = original.IterationPerformance
type ListDomain = original.ListDomain
type ListExport = original.ListExport
type ListImage = original.ListImage
type ListImagePerformance = original.ListImagePerformance
type ListIteration = original.ListIteration
type ListProject = original.ListProject
type ListTag = original.ListTag
type Prediction = original.Prediction
type PredictionQueryResult = original.PredictionQueryResult
type PredictionQueryTag = original.PredictionQueryTag
type PredictionQueryToken = original.PredictionQueryToken
type Project = original.Project
type ProjectSettings = original.ProjectSettings
type Region = original.Region
type RegionProposal = original.RegionProposal
type StoredImagePrediction = original.StoredImagePrediction
type Tag = original.Tag
type TagPerformance = original.TagPerformance

func New(aPIKey string, endpoint string) BaseClient {
	return original.New(aPIKey, endpoint)
}
func NewWithoutDefaults(aPIKey string, endpoint string) BaseClient {
	return original.NewWithoutDefaults(aPIKey, endpoint)
}
func PossibleClassifierValues() []Classifier {
	return original.PossibleClassifierValues()
}
func PossibleDomainTypeValues() []DomainType {
	return original.PossibleDomainTypeValues()
}
func PossibleExportFlavorModelValues() []ExportFlavorModel {
	return original.PossibleExportFlavorModelValues()
}
func PossibleExportPlatformModelValues() []ExportPlatformModel {
	return original.PossibleExportPlatformModelValues()
}
func PossibleExportStatusModelValues() []ExportStatusModel {
	return original.PossibleExportStatusModelValues()
}
func PossibleImageCreateStatusValues() []ImageCreateStatus {
	return original.PossibleImageCreateStatusValues()
}
func PossibleOrderByValues() []OrderBy {
	return original.PossibleOrderByValues()
}
func PossibleTagTypeValues() []TagType {
	return original.PossibleTagTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}