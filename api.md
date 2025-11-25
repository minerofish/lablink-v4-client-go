# Shared Params Types

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go/shared#ResultParam">ResultParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go/shared#Result">Result</a>

# Bloodlab

Response Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BloodlabOrganizationRelation">BloodlabOrganizationRelation</a>

Methods:

- <code title="get /v4/bloodlab/examinations">client.Bloodlab.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BloodlabService.ListExaminations">ListExaminations</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BloodlabOrganizationRelation">BloodlabOrganizationRelation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v4/bloodlab/organization-relations">client.Bloodlab.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BloodlabService.ListOrganizationRelations">ListOrganizationRelations</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BloodlabOrganizationRelation">BloodlabOrganizationRelation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Orders

Response Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BloodlabOrderListNewOrdersResponse">BloodlabOrderListNewOrdersResponse</a>

Methods:

- <code title="get /v4/bloodlab/orders">client.Bloodlab.Orders.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BloodlabOrderService.ListNewOrders">ListNewOrders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BloodlabOrderListNewOrdersResponse">BloodlabOrderListNewOrdersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v4/bloodlab/orders/downloaded">client.Bloodlab.Orders.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BloodlabOrderService.MarkAsDownloaded">MarkAsDownloaded</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BloodlabOrderMarkAsDownloadedParams">BloodlabOrderMarkAsDownloadedParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v4/bloodlab/orders/state">client.Bloodlab.Orders.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BloodlabOrderService.SetState">SetState</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BloodlabOrderSetStateParams">BloodlabOrderSetStateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Procedures

Params Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#CodingSystem">CodingSystem</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#CodingSystem">CodingSystem</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#ProcedureResponse">ProcedureResponse</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#ProcedureUpdateResponse">ProcedureUpdateResponse</a>

Methods:

- <code title="post /v4/procedures">client.Procedures.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#ProcedureService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#ProcedureNewParams">ProcedureNewParams</a>) (<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#ProcedureResponse">ProcedureResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v4/procedures/{procedureId}">client.Procedures.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#ProcedureService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, procedureID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#ProcedureUpdateParams">ProcedureUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#ProcedureUpdateResponse">ProcedureUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v4/procedures">client.Procedures.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#ProcedureService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#ProcedureListParams">ProcedureListParams</a>) ([]<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#ProcedureResponse">ProcedureResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v4/procedures/{procedureId}">client.Procedures.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#ProcedureService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, procedureID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Examinations

Response Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Examination">Examination</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#ExaminationListResponse">ExaminationListResponse</a>

Methods:

- <code title="get /v4/examinations">client.Examinations.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#ExaminationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#ExaminationListParams">ExaminationListParams</a>) (<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#ExaminationListResponse">ExaminationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Laboratories

Response Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Laboratory">Laboratory</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryUserRelations">LaboratoryUserRelations</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryNewResponse">LaboratoryNewResponse</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryQueryContractsResponse">LaboratoryQueryContractsResponse</a>

Methods:

- <code title="post /v4/laboratories">client.Laboratories.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryNewParams">LaboratoryNewParams</a>) ([]<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryNewResponse">LaboratoryNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v4/laboratories">client.Laboratories.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryListParams">LaboratoryListParams</a>) ([]<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Laboratory">Laboratory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v4/laboratories">client.Laboratories.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryDeleteParams">LaboratoryDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v4/laboratories/{laboratoryId}/contracts-query">client.Laboratories.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryService.QueryContracts">QueryContracts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, laboratoryID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryQueryContractsParams">LaboratoryQueryContractsParams</a>) (<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryQueryContractsResponse">LaboratoryQueryContractsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v4/laboratories/{laboratoryId}/document">client.Laboratories.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryService.UploadDocument">UploadDocument</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, laboratoryID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryUploadDocumentParams">LaboratoryUploadDocumentParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v4/laboratories/{laboratoryId}/results">client.Laboratories.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryService.UploadResults">UploadResults</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, laboratoryID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryUploadResultsParams">LaboratoryUploadResultsParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Orders

Response Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryOrderListResponse">LaboratoryOrderListResponse</a>

Methods:

- <code title="patch /v4/laboratories/{laboratoryId}/orders">client.Laboratories.Orders.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryOrderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, laboratoryID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryOrderUpdateParams">LaboratoryOrderUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v4/laboratories/{laboratoryId}/orders">client.Laboratories.Orders.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryOrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, laboratoryID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryOrderListParams">LaboratoryOrderListParams</a>) (<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryOrderListResponse">LaboratoryOrderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Contracts

Response Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Contract">Contract</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryContractGetResponse">LaboratoryContractGetResponse</a>

Methods:

- <code title="post /v4/laboratories/{laboratoryId}/contracts">client.Laboratories.Contracts.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryContractService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, laboratoryID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryContractNewParams">LaboratoryContractNewParams</a>) (<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Contract">Contract</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v4/laboratories/{laboratoryId}/contracts">client.Laboratories.Contracts.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryContractService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, laboratoryID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryContractGetParams">LaboratoryContractGetParams</a>) (<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryContractGetResponse">LaboratoryContractGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v4/laboratories/{laboratoryId}/contracts">client.Laboratories.Contracts.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryContractService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, laboratoryID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryContractUpdateParams">LaboratoryContractUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /v4/laboratories/{laboratoryId}/contracts">client.Laboratories.Contracts.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryContractService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, laboratoryID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryContractDeleteParams">LaboratoryContractDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Query

Response Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryQueryExecuteResponse">LaboratoryQueryExecuteResponse</a>

Methods:

- <code title="post /v4/laboratories-query">client.Laboratories.Query.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryQueryService.Execute">Execute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryQueryExecuteParams">LaboratoryQueryExecuteParams</a>) (<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LaboratoryQueryExecuteResponse">LaboratoryQueryExecuteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Locations

Response Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Location">Location</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LocationListResponse">LocationListResponse</a>

Methods:

- <code title="post /v4/locations">client.Locations.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LocationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LocationNewParams">LocationNewParams</a>) ([]<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Location">Location</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v4/locations">client.Locations.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LocationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LocationListParams">LocationListParams</a>) (<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LocationListResponse">LocationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v4/locations">client.Locations.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LocationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LocationDeleteParams">LocationDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Login

Response Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#AuthToken">AuthToken</a>

Methods:

- <code title="post /v4/login">client.Login.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LoginService.Authenticate">Authenticate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#LoginAuthenticateParams">LoginAuthenticateParams</a>) (<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#AuthToken">AuthToken</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Authorize

Methods:

- <code title="get /v4/authorize">client.Authorize.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#AuthorizeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#AuthorizeGetParams">AuthorizeGetParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Token

Methods:

- <code title="post /v4/token">client.Token.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#TokenService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#TokenNewParams">TokenNewParams</a>) (<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#AuthToken">AuthToken</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Orders

Params Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BloodDonorParam">BloodDonorParam</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BloodType">BloodType</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BoneMarrowDonorParam">BoneMarrowDonorParam</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Gender">Gender</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Kell">Kell</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderStateType">OrderStateType</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderType">OrderType</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#PatientParam">PatientParam</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#PseudonymParam">PseudonymParam</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Rhesus">Rhesus</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BloodDonor">BloodDonor</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BloodType">BloodType</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#BoneMarrowDonor">BoneMarrowDonor</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Gender">Gender</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Kell">Kell</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Link">Link</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderExamination">OrderExamination</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderMetadata">OrderMetadata</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderStateType">OrderStateType</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderType">OrderType</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Patient">Patient</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Pseudonym">Pseudonym</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Rhesus">Rhesus</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderGetResponse">OrderGetResponse</a>

Methods:

- <code title="post /v4/orders">client.Orders.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderNewParams">OrderNewParams</a>) ([]<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderMetadata">OrderMetadata</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v4/orders/{orderId}">client.Orders.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderGetResponse">OrderGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v4/orders">client.Orders.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderDeleteParams">OrderDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /v4/orders/{orderId}">client.Orders.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderService.DeleteOrder">DeleteOrder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v4/orders/{orderId}/documents/{documentId}">client.Orders.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderService.GetDocument">GetDocument</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, documentID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderGetDocumentParams">OrderGetDocumentParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Tags

Methods:

- <code title="delete /v4/orders/tags">client.Orders.Tags.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderTagService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderTagDeleteParams">OrderTagDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v4/orders/tags">client.Orders.Tags.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderTagService.Tag">Tag</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderTagTagParams">OrderTagTagParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Examinations

Response Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderExaminationListResponse">OrderExaminationListResponse</a>
- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderExaminationAddResponse">OrderExaminationAddResponse</a>

Methods:

- <code title="get /v4/orders/{orderId}/examinations">client.Orders.Examinations.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderExaminationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderExaminationListResponse">OrderExaminationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v4/orders/{orderId}/examinations">client.Orders.Examinations.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderExaminationService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderExaminationAddParams">OrderExaminationAddParams</a>) ([]<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderExaminationAddResponse">OrderExaminationAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v4/orders/{orderId}/examinations">client.Orders.Examinations.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderExaminationService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderExaminationRemoveParams">OrderExaminationRemoveParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Query

Response Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderQueryNewResponse">OrderQueryNewResponse</a>

Methods:

- <code title="post /v4/orders-query">client.Orders.Query.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderQueryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderQueryNewParams">OrderQueryNewParams</a>) (<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderQueryNewResponse">OrderQueryNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v4/orders-query/{transactionId}">client.Orders.Query.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderQueryService.CommitTransaction">CommitTransaction</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrderQueryCommitTransactionParams">OrderQueryCommitTransactionParams</a>) ([]<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Organizations

Params Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrganizationRole">OrganizationRole</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Organization">Organization</a>

Methods:

- <code title="post /v4/organizations">client.Organizations.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrganizationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrganizationNewParams">OrganizationNewParams</a>) ([]<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Organization">Organization</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v4/organizations/{organizationId}">client.Organizations.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrganizationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, organizationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrganizationUpdateParams">OrganizationUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Organization">Organization</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v4/organizations">client.Organizations.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrganizationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrganizationListParams">OrganizationListParams</a>) ([]<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#Organization">Organization</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Query

Response Types:

- <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrganizationQueryNewResponse">OrganizationQueryNewResponse</a>

Methods:

- <code title="post /v4/organizations-query">client.Organizations.Query.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrganizationQueryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrganizationQueryNewParams">OrganizationQueryNewParams</a>) (<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go">lablinkv4client</a>.<a href="https://pkg.go.dev/github.com/minerofish/lablink-v4-client-go#OrganizationQueryNewResponse">OrganizationQueryNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
