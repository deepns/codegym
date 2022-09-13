# AZ-900 Exam Notes

## Azure Architecture and Services

### Idenity, Access and Security

#### Azure Active Directory - directory service to sign into MS cloud, other cloud applications and possibly on-prem

- enables signing in to Microsoft cloud applications including Azure, Office 365, Dynamics and other cloud applications configured to use Microsoft authentication
- Same functionality (identity and access mgmt) provided by Active Directory running on Windows Server in on-prem deployments
- In on-prem AD deployments, Microsoft doesn't monitor the sign-ins. With Azure AD, Microsoft monitors it and provides additional services and protections (threat detection, suspicious sign-ins etc.)
- Azure AD services - *authentication, Single-Sign-On (SSO), App management, Device management with Intune*
- In hybrid cloud environment, maintaining two identity services (AD in on-prem and Azure AD in Cloud) is not recommended. Can combine that into one using **Azure AD Connect**. Sync changes between both the systems

#### Azure Active Directory Domain Services - managed domain services

- eliminates the need to deploy, manage and patch Windows Server domain controllers (DC) in the cloud. #question. What are domain controllers?
- legacy apps using domain services in on-prem can be lift-and-shifted into cloud with the help of Azure AD DS
- integrates with Azure AD tenant
- An unique namespace is defined per Azure AD DS managed domain. Azure deploys two Domain Controllers in the chosen region and manages it for us.
- One way sync between Azure AD and Azure AD DS.

## Fundamentals - Management and Governance

### Azure Blueprints - define repeatable settings and policies

- standardize cloud subscription and environment deployments to enforce settings and policies
- comprises of one or more artifacts. what is an artifact here? component in blueprint definition that defines the settings and the parameters to be configured
- artifact parameters can be specified at the creation time or when applying the blueprint to a particular scope. Allows the BP to be more configurable
- what other info can be specified in an artifact? - Role assignments, policy assignments, ARM (Azure Resource Manager) templates, Resource Groups
- Blueprints are version-able
- manages policies across multiple subscriptions
- orchestrates the deployment of resource templates, role assignments, policy assignments, ARM templates

### Azure Policy - enforce rules across resource configurations to maintain compliance

- create, assign and manage policies to control and audit Azure resources
- define individual policy or group of policies (called as **Initiatives**)
- can be set at each level individually or a high level. Resources inherit the policies from its parent.
- can be applied to new resources as well as existing ones.
- automatic remediation of policy non-compliance in some cases. e.g. re-applying a tag defined in the policy
- resources can be configured to be exempted from certain policies
- many [built-in policy and intiative definitions](https://docs.microsoft.com/en-us/azure/governance/policy/samples/built-in-initiatives) for Compute, Storage, Networking, Monitoring and Security Center

### Azure Policy Initiatives - group of Azure policies

- grouping of policies to track compliance for a larger goal.
- e.g. **Enable Throughput** initiative combines two policies. See [definition](https://github.com/Azure/azure-policy/blob/master/built-in-policies/policySetDefinitions/Cosmos%20DB/Cosmos_Throughput.json)

### Resource Locks - to prevent accidental changes and deletion to a resource

- prevents resource from being deleted or changed.
- types of resource locks
  - **Delete** - prevents deletion of a reosurce
  - **ReadOnly** - make the resource read-only
- applies regardless of the RBAC permissions for a resource
- resource inherits locks from its parent. e.g. locks applied at Resource Groups apply to all resources under that group
- manage resource locks from Azure Portal, PowerShell, Azure CLI or from ARM template
- If lock is accidentally removed, it can be automatically re-instated if Azure Blueprint is in use
- locks need to be removed before deleting a resource where a lock is applied
- Organize resource using tags. Some ways to categorize using tags
  - Resource management - to locate and act on resources associated with specific workloads, BUs etc.
  - Cost management
  - Operations management - group resources based on operational availability and criticality. Formulate SLA.
  - Security and governance tags
  - Manage tags
    - PowerShell, Azure CLI, ARM templates, REST API, Azure Portal
    - Azure Policy
      - to inherit tags from a resource group to a resource under that group
      - to enforce tagging rules an conventions
      - to control and audit resources

### Service Trust Portal - to view Microsoft's security, privacy and compliance practices

- [Service Trust Portal](https://servicetrust.microsoft.com/) explains MS's implementation of various controls, processes and services to protect cloud services and customer data
- Documents in Service Trust Portal available for 12 months from publishing
- [Trust Center](https://www.microsoft.com/en-us/trust-center) - explaining privacy and compliance practices and solutions

### Azure Arc - extend Azure compliance and monitoring to hybrid and multi-cloud environments

- Most common tools *Azure Portal*, *Azure Cloud Shell*, *Azure PowerShell*, *Azure CLI* to manage resources within Azure.
- Azure Arc utilizes Azure Resource Manager to extend Azure compliance and monitoring to hybrid and multi cloud environments.
- manage VMs, K8s clusters, databases (SQL servers) as if they are running in Azure

### Azure Resource Manager - deployment and management service for Azure

- forms the backbone of the management layer
- all deployment, access and management goes through ARM. All API, CLI, portal invoke the same API to ARM
- manage infrastructure as code. specify the desired state of the resources in declarative templates (**ARM Templates**). Much like managing kubernetes resources with spec files.
- dependencies can be specified in the template. Azure takes care of the ordering and invoking the right tools to deploy the resources
- ARM templates can be modular, nested, and also extended (with PowerShell or Bash scripts inline or externally sourced)

### Monitoring Tools

#### Azure Advisor - recommendations to optimize the cloud environment

- provides recommendations to improve *reliability, security, performance, operational excellence and cost reduction*.
- [Dashboard](https://docs.microsoft.com/en-us/training/wwl-azure/describe-monitoring-tools-azure/media/azure-advisor-dashboard-baca22e2.png)
- recommendations available in the Azure Portal or through Advisor APIs

#### Azure Service Health - status of deployed resources and overall Azure services

- **Azure Status** - provide status of various Azure services across the globe
- **Service Health** - focus on the service and regions used in the account. sessions are authenticated. want to know about commmunications about outages, planned maintenance and health advisories? Service Health is the place. Can subscribe to planned activities through Service Health Alerts.
- **Resource Health** - tailored view of the actual resources under use.

#### Azure Monitor - platform to collect metrics and logs, analyze and act on the results

- Collect metrics, logs, traces from resources in Azure, on-prem and multi cloud.
- [comprehensive picture](https://docs.microsoft.com/en-us/training/wwl-azure/describe-monitoring-tools-azure/media/azure-monitor-overview-614cd2fd.svg)
- logging and metric data stored in central repositories and fed to other components (visualizers, analyzers etc.)
- **Azure Log Analytics** - tool to write and run log queries on the data gather in Azure Monitor. Run simple to complex queries to filter records, visualize the results etc.
- **Azure Monitor Alerts** - set up alerts to monitor the logs or metrics, to get notified when specified threshold is crossed for a resource. (e.g. VM CPU usage exceeding 80%, disk usage exceeding certain size etc.).

#### Application Insights - to monitor web applications in Azure, on-prem and multi cloud

- add Application Insight support to an application through Application Insight SDK or Agent
- provides lot of information such as request rates, response times, page views, performance counters from VMs etc.
- can also be used to send artificial traffic to the application during periods of low-activicty to check the status
