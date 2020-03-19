/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package apigatewaycontrollerapiv1

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/IBM/apigateway-go-sdk/apigatewaycontrollerapiv1"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ApiGatewayControllerApiV1`, func() {
	Describe(`GetAllEndpoints(getAllEndpointsOptions *GetAllEndpointsOptions)`, func() {
		getAllEndpointsPath := "/v2/endpoints"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAllEndpointsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["service_instance_crn"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["provider_id"]).To(Equal([]string{"testString"}))

				// TODO: Add check for shared query parameter

				// TODO: Add check for managed query parameter

				Expect(req.URL.Query()["swagger"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `[{}]`)
			}))
			It(`Invoke GetAllEndpoints successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := apigatewaycontrollerapiv1.NewApiGatewayControllerApiV1(&apigatewaycontrollerapiv1.ApiGatewayControllerApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAllEndpoints(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAllEndpointsOptions model
				getAllEndpointsOptionsModel := new(apigatewaycontrollerapiv1.GetAllEndpointsOptions)
				getAllEndpointsOptionsModel.ServiceInstanceCrn = core.StringPtr("testString")
				getAllEndpointsOptionsModel.Authorization = core.StringPtr("testString")
				getAllEndpointsOptionsModel.ProviderID = core.StringPtr("testString")
				getAllEndpointsOptionsModel.Shared = core.BoolPtr(true)
				getAllEndpointsOptionsModel.Managed = core.BoolPtr(true)
				getAllEndpointsOptionsModel.Swagger = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAllEndpoints(getAllEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateEndpoint(createEndpointOptions *CreateEndpointOptions)`, func() {
		createEndpointPath := "/v2/endpoints"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createEndpointPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"artifact_id": "ArtifactID", "crn": "Crn", "parent_crn": "ParentCrn", "service_instance_crn": "ServiceInstanceCrn", "account_id": "AccountID", "provider_id": "ProviderID", "name": "Name", "routes": ["Routes"], "managed_url": "ManagedURL", "alias_url": "AliasURL", "shared": true, "managed": false, "policies": [], "base_path": "/example"}`)
			}))
			It(`Invoke CreateEndpoint successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := apigatewaycontrollerapiv1.NewApiGatewayControllerApiV1(&apigatewaycontrollerapiv1.ApiGatewayControllerApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateEndpoint(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateEndpointOptions model
				createEndpointOptionsModel := new(apigatewaycontrollerapiv1.CreateEndpointOptions)
				createEndpointOptionsModel.Authorization = core.StringPtr("testString")
				createEndpointOptionsModel.ArtifactID = core.StringPtr("testString")
				createEndpointOptionsModel.ParentCrn = core.StringPtr("testString")
				createEndpointOptionsModel.ServiceInstanceCrn = core.StringPtr("testString")
				createEndpointOptionsModel.Name = core.StringPtr("testString")
				createEndpointOptionsModel.Routes = []string{"testString"}
				createEndpointOptionsModel.Managed = core.BoolPtr(true)
				createEndpointOptionsModel.Metadata = CreateMockMap()
				createEndpointOptionsModel.OpenApiDoc = CreateMockMap()

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateEndpoint(createEndpointOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetEndpoint(getEndpointOptions *GetEndpointOptions)`, func() {
		getEndpointPath := "/v2/endpoints/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getEndpointPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["service_instance_crn"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"artifact_id": "ArtifactID", "crn": "Crn", "parent_crn": "ParentCrn", "service_instance_crn": "ServiceInstanceCrn", "account_id": "AccountID", "provider_id": "ProviderID", "name": "Name", "routes": ["Routes"], "managed_url": "ManagedURL", "alias_url": "AliasURL", "shared": true, "managed": false, "policies": [], "base_path": "/example"}`)
			}))
			It(`Invoke GetEndpoint successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := apigatewaycontrollerapiv1.NewApiGatewayControllerApiV1(&apigatewaycontrollerapiv1.ApiGatewayControllerApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetEndpoint(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEndpointOptions model
				getEndpointOptionsModel := new(apigatewaycontrollerapiv1.GetEndpointOptions)
				getEndpointOptionsModel.ID = core.StringPtr("testString")
				getEndpointOptionsModel.ServiceInstanceCrn = core.StringPtr("testString")
				getEndpointOptionsModel.Authorization = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetEndpoint(getEndpointOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateEndpoint(updateEndpointOptions *UpdateEndpointOptions)`, func() {
		updateEndpointPath := "/v2/endpoints/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateEndpointPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["service_instance_crn"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"artifact_id": "ArtifactID", "crn": "Crn", "parent_crn": "ParentCrn", "service_instance_crn": "ServiceInstanceCrn", "account_id": "AccountID", "provider_id": "ProviderID", "name": "Name", "routes": ["Routes"], "managed_url": "ManagedURL", "alias_url": "AliasURL", "shared": true, "managed": false, "policies": [], "base_path": "/example"}`)
			}))
			It(`Invoke UpdateEndpoint successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := apigatewaycontrollerapiv1.NewApiGatewayControllerApiV1(&apigatewaycontrollerapiv1.ApiGatewayControllerApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateEndpoint(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateEndpointOptions model
				updateEndpointOptionsModel := new(apigatewaycontrollerapiv1.UpdateEndpointOptions)
				updateEndpointOptionsModel.ID = core.StringPtr("testString")
				updateEndpointOptionsModel.ServiceInstanceCrn = core.StringPtr("testString")
				updateEndpointOptionsModel.Authorization = core.StringPtr("testString")
				updateEndpointOptionsModel.NewArtifactID = core.StringPtr("testString")
				updateEndpointOptionsModel.NewParentCrn = core.StringPtr("testString")
				updateEndpointOptionsModel.NewServiceInstanceCrn = core.StringPtr("testString")
				updateEndpointOptionsModel.NewName = core.StringPtr("testString")
				updateEndpointOptionsModel.NewRoutes = []string{"testString"}
				updateEndpointOptionsModel.NewManaged = core.BoolPtr(true)
				updateEndpointOptionsModel.NewMetadata = CreateMockMap()
				updateEndpointOptionsModel.NewOpenApiDoc = CreateMockMap()

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateEndpoint(updateEndpointOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteEndpoint(deleteEndpointOptions *DeleteEndpointOptions)`, func() {
		deleteEndpointPath := "/v2/endpoints/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteEndpointPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["service_instance_crn"]).To(Equal([]string{"testString"}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteEndpoint successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := apigatewaycontrollerapiv1.NewApiGatewayControllerApiV1(&apigatewaycontrollerapiv1.ApiGatewayControllerApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteEndpoint(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteEndpointOptions model
				deleteEndpointOptionsModel := new(apigatewaycontrollerapiv1.DeleteEndpointOptions)
				deleteEndpointOptionsModel.ID = core.StringPtr("testString")
				deleteEndpointOptionsModel.ServiceInstanceCrn = core.StringPtr("testString")
				deleteEndpointOptionsModel.Authorization = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteEndpoint(deleteEndpointOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetEndpointSwagger(getEndpointSwaggerOptions *GetEndpointSwaggerOptions)`, func() {
		getEndpointSwaggerPath := "/v2/endpoints/testString/swagger"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getEndpointSwaggerPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["service_instance_crn"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["type"]).To(Equal([]string{"json"}))

				res.WriteHeader(200)
			}))
			It(`Invoke GetEndpointSwagger successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := apigatewaycontrollerapiv1.NewApiGatewayControllerApiV1(&apigatewaycontrollerapiv1.ApiGatewayControllerApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.GetEndpointSwagger(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetEndpointSwaggerOptions model
				getEndpointSwaggerOptionsModel := new(apigatewaycontrollerapiv1.GetEndpointSwaggerOptions)
				getEndpointSwaggerOptionsModel.ID = core.StringPtr("testString")
				getEndpointSwaggerOptionsModel.ServiceInstanceCrn = core.StringPtr("testString")
				getEndpointSwaggerOptionsModel.Authorization = core.StringPtr("testString")
				getEndpointSwaggerOptionsModel.Type = core.StringPtr("json")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.GetEndpointSwagger(getEndpointSwaggerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`EndpointActions(endpointActionsOptions *EndpointActionsOptions)`, func() {
		endpointActionsPath := "/v2/endpoints/testString/actions"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(endpointActionsPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["service_instance_crn"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["provider_id"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"artifact_id": "ArtifactID", "crn": "Crn", "parent_crn": "ParentCrn", "service_instance_crn": "ServiceInstanceCrn", "account_id": "AccountID", "provider_id": "ProviderID", "name": "Name", "routes": ["Routes"], "managed_url": "ManagedURL", "alias_url": "AliasURL", "shared": true, "managed": false, "policies": [], "base_path": "/example"}`)
			}))
			It(`Invoke EndpointActions successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := apigatewaycontrollerapiv1.NewApiGatewayControllerApiV1(&apigatewaycontrollerapiv1.ApiGatewayControllerApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.EndpointActions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EndpointActionsOptions model
				endpointActionsOptionsModel := new(apigatewaycontrollerapiv1.EndpointActionsOptions)
				endpointActionsOptionsModel.ID = core.StringPtr("testString")
				endpointActionsOptionsModel.ServiceInstanceCrn = core.StringPtr("testString")
				endpointActionsOptionsModel.ProviderID = core.StringPtr("testString")
				endpointActionsOptionsModel.Authorization = core.StringPtr("testString")
				endpointActionsOptionsModel.Type = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.EndpointActions(endpointActionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`EndpointSummary(endpointSummaryOptions *EndpointSummaryOptions)`, func() {
		endpointSummaryPath := "/v2/endpoints/summary"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(endpointSummaryPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["service_instance_crn"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["swagger"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `[{}]`)
			}))
			It(`Invoke EndpointSummary successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := apigatewaycontrollerapiv1.NewApiGatewayControllerApiV1(&apigatewaycontrollerapiv1.ApiGatewayControllerApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.EndpointSummary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EndpointSummaryOptions model
				endpointSummaryOptionsModel := new(apigatewaycontrollerapiv1.EndpointSummaryOptions)
				endpointSummaryOptionsModel.AccountID = core.StringPtr("testString")
				endpointSummaryOptionsModel.Authorization = core.StringPtr("testString")
				endpointSummaryOptionsModel.ServiceInstanceCrn = core.StringPtr("testString")
				endpointSummaryOptionsModel.Swagger = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.EndpointSummary(endpointSummaryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetAllSubscriptions(getAllSubscriptionsOptions *GetAllSubscriptionsOptions)`, func() {
		getAllSubscriptionsPath := "/v2/subscriptions"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAllSubscriptionsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["artifact_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `[{}]`)
			}))
			It(`Invoke GetAllSubscriptions successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := apigatewaycontrollerapiv1.NewApiGatewayControllerApiV1(&apigatewaycontrollerapiv1.ApiGatewayControllerApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAllSubscriptions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAllSubscriptionsOptions model
				getAllSubscriptionsOptionsModel := new(apigatewaycontrollerapiv1.GetAllSubscriptionsOptions)
				getAllSubscriptionsOptionsModel.ArtifactID = core.StringPtr("testString")
				getAllSubscriptionsOptionsModel.Authorization = core.StringPtr("testString")
				getAllSubscriptionsOptionsModel.Type = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAllSubscriptions(getAllSubscriptionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateSubscription(createSubscriptionOptions *CreateSubscriptionOptions)`, func() {
		createSubscriptionPath := "/v2/subscriptions"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createSubscriptionPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"client_id": "ClientID", "secret_provided": true, "artifact_id": "ArtifactID", "account_id": "AccountID", "name": "Name", "type": "Type"}`)
			}))
			It(`Invoke CreateSubscription successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := apigatewaycontrollerapiv1.NewApiGatewayControllerApiV1(&apigatewaycontrollerapiv1.ApiGatewayControllerApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateSubscription(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateSubscriptionOptions model
				createSubscriptionOptionsModel := new(apigatewaycontrollerapiv1.CreateSubscriptionOptions)
				createSubscriptionOptionsModel.Authorization = core.StringPtr("testString")
				createSubscriptionOptionsModel.ClientID = core.StringPtr("testString")
				createSubscriptionOptionsModel.ArtifactID = core.StringPtr("testString")
				createSubscriptionOptionsModel.ClientSecret = core.StringPtr("testString")
				createSubscriptionOptionsModel.AccountID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Name = core.StringPtr("testString")
				createSubscriptionOptionsModel.Type = core.StringPtr("testString")
				createSubscriptionOptionsModel.Metadata = CreateMockMap()

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateSubscription(createSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetSubscription(getSubscriptionOptions *GetSubscriptionOptions)`, func() {
		getSubscriptionPath := "/v2/subscriptions/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getSubscriptionPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["artifact_id"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"client_id": "ClientID", "secret_provided": true, "artifact_id": "ArtifactID", "account_id": "AccountID", "name": "Name", "type": "Type"}`)
			}))
			It(`Invoke GetSubscription successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := apigatewaycontrollerapiv1.NewApiGatewayControllerApiV1(&apigatewaycontrollerapiv1.ApiGatewayControllerApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetSubscription(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSubscriptionOptions model
				getSubscriptionOptionsModel := new(apigatewaycontrollerapiv1.GetSubscriptionOptions)
				getSubscriptionOptionsModel.ID = core.StringPtr("testString")
				getSubscriptionOptionsModel.ArtifactID = core.StringPtr("testString")
				getSubscriptionOptionsModel.Authorization = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetSubscription(getSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateSubscription(updateSubscriptionOptions *UpdateSubscriptionOptions)`, func() {
		updateSubscriptionPath := "/v2/subscriptions/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateSubscriptionPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["artifact_id"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"client_id": "ClientID", "secret_provided": true, "artifact_id": "ArtifactID", "account_id": "AccountID", "name": "Name", "type": "Type"}`)
			}))
			It(`Invoke UpdateSubscription successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := apigatewaycontrollerapiv1.NewApiGatewayControllerApiV1(&apigatewaycontrollerapiv1.ApiGatewayControllerApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateSubscription(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateSubscriptionOptions model
				updateSubscriptionOptionsModel := new(apigatewaycontrollerapiv1.UpdateSubscriptionOptions)
				updateSubscriptionOptionsModel.ID = core.StringPtr("testString")
				updateSubscriptionOptionsModel.ArtifactID = core.StringPtr("testString")
				updateSubscriptionOptionsModel.Authorization = core.StringPtr("testString")
				updateSubscriptionOptionsModel.NewClientID = core.StringPtr("testString")
				updateSubscriptionOptionsModel.NewClientSecret = core.StringPtr("testString")
				updateSubscriptionOptionsModel.NewArtifactID = core.StringPtr("testString")
				updateSubscriptionOptionsModel.NewAccountID = core.StringPtr("testString")
				updateSubscriptionOptionsModel.NewName = core.StringPtr("testString")
				updateSubscriptionOptionsModel.NewType = core.StringPtr("testString")
				updateSubscriptionOptionsModel.NewMetadata = CreateMockMap()

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateSubscription(updateSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteSubscription(deleteSubscriptionOptions *DeleteSubscriptionOptions)`, func() {
		deleteSubscriptionPath := "/v2/subscriptions/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteSubscriptionPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["artifact_id"]).To(Equal([]string{"testString"}))

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteSubscription successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := apigatewaycontrollerapiv1.NewApiGatewayControllerApiV1(&apigatewaycontrollerapiv1.ApiGatewayControllerApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteSubscription(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSubscriptionOptions model
				deleteSubscriptionOptionsModel := new(apigatewaycontrollerapiv1.DeleteSubscriptionOptions)
				deleteSubscriptionOptionsModel.ID = core.StringPtr("testString")
				deleteSubscriptionOptionsModel.ArtifactID = core.StringPtr("testString")
				deleteSubscriptionOptionsModel.Authorization = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteSubscription(deleteSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetSubscriptionArtifact(getSubscriptionArtifactOptions *GetSubscriptionArtifactOptions)`, func() {
		getSubscriptionArtifactPath := "/v2/subscriptions/artifact"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getSubscriptionArtifactPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["artifact_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["client_id"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"managed_url": "ManagedURL"}`)
			}))
			It(`Invoke GetSubscriptionArtifact successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := apigatewaycontrollerapiv1.NewApiGatewayControllerApiV1(&apigatewaycontrollerapiv1.ApiGatewayControllerApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetSubscriptionArtifact(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSubscriptionArtifactOptions model
				getSubscriptionArtifactOptionsModel := new(apigatewaycontrollerapiv1.GetSubscriptionArtifactOptions)
				getSubscriptionArtifactOptionsModel.ArtifactID = core.StringPtr("testString")
				getSubscriptionArtifactOptionsModel.ClientID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetSubscriptionArtifact(getSubscriptionArtifactOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`AddSubscriptionSecret(addSubscriptionSecretOptions *AddSubscriptionSecretOptions)`, func() {
		addSubscriptionSecretPath := "/v2/subscriptions/testString/secret"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addSubscriptionSecretPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["artifact_id"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"client_id": "ClientID", "secret_provided": true, "artifact_id": "ArtifactID", "account_id": "AccountID", "name": "Name", "type": "Type"}`)
			}))
			It(`Invoke AddSubscriptionSecret successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := apigatewaycontrollerapiv1.NewApiGatewayControllerApiV1(&apigatewaycontrollerapiv1.ApiGatewayControllerApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.AddSubscriptionSecret(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddSubscriptionSecretOptions model
				addSubscriptionSecretOptionsModel := new(apigatewaycontrollerapiv1.AddSubscriptionSecretOptions)
				addSubscriptionSecretOptionsModel.ID = core.StringPtr("testString")
				addSubscriptionSecretOptionsModel.ArtifactID = core.StringPtr("testString")
				addSubscriptionSecretOptionsModel.Authorization = core.StringPtr("testString")
				addSubscriptionSecretOptionsModel.ClientSecret = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.AddSubscriptionSecret(addSubscriptionSecretOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockMap() successfully`, func() {
			mockMap := CreateMockMap()
			Expect(mockMap).ToNot(BeNil())
		})
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockMap() map[string]interface{} {
	m := make(map[string]interface{})
	return m
}

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, len(mockData))
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Now())
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Now())
	return &d
}
