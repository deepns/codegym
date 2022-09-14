# AZ-900 Exam Notes

## Core components

### Infrastructure

- AZ account -> Subscription -> Resource Groups -> Resources
- Activated Learn Sandbox. Hands on with Azure Cloud Shell, Bash
- [Global Infrastructure Site](https://infrastructuremap.microsoft.com/)
- Data centers grouped into Regions or Availability Zones
- Region - one or more data centers, networked together with a low-latency network. intelligent resource placement within the region for appropriate workload balancing.
  - Most resources require the region to be specified. Some are region agnostic: Azure Active Directory, Azure Traffic Manager, and Azure DNS
  - Cost varies per region due to difference in the local cost and laws.
  - Some resource and resource types may be available only in a particular region
- Availability Zone - isolation boundary for resources. physically separate data centers within a region. Some regions may not support availability zones. Used for high availability.
  - primarily for VMs, managed disks, load balancers, and SQL databases.
  - Services that support az fall under three types
    - Zonal - resource pinned to specific zone
    - Zone-redundant - Azure automatically replicates the resource to another zone within the region
    - Non-regional - resource not affected by zone-wide or region-wide outages.
  - ![availability-zone-and-regions](https://docs.microsoft.com/en-us/azure/availability-zones/media/availability-zones.png)
- Region Pairs - high availability for regions
  - Most regions are paired with another region that is at least 300km apart. For disaster recovery.
  - Most region pairings are bi-directional (active/active), while some are uni-directional (active/passive)
  - Data will reside in the same [geography](https://azure.microsoft.com/en-us/explore/global-infrastructure/geographies/#geographies) as the pair
  - Planned updates are rolled out to paired regions, one region at a time.
  - In case of outage, one region out of every pair is prioritized for quicker recovery.
- Sovereign region - separate instance of Azure, completely isolated from the main instance of Azure.
  - e.g. regions -> US DoD Central, US Gov Virginia, US Iowa - require special clearance and compliance to operate
  - China East, China North - regions in Azure in China, operated by 21ViaNet.

### Subscriptions

- Resource - fundamental unit. VMs, Disks, Databases, Networks, Blobs etc.
- Resource Groups - group of resources. Resources inherit properties, access policies from the resource groups. A resource can’t be in multiple resource groups. No nesting either. Deleting a resource group will delete the resources under the group too.
- Subscription - grouping of resource groups.
- Every Azure account required to have at least one subscription. more is optional. Subscription links to an Azure Account, which is an identity in the Azure Active Directory or in a directory that Azure AD trusts.
- Subscription Boundaries
  - Billing boundary - separate billing reports and invoices for each subscription. e.g. subscription for every dept within an org.
  - Access control boundary - apply access management policies at the subscription level.
- may have some hard limits depending on the resources

### Management Groups

- Logical container to group subscriptions, for better compliance and governance.
- Subscription inherits the policies and RBAC applied at the management group level.
- max up to 10000 mgmt. groups under a directory
- up to six level of depth
- management group can have multiple management groups under its hierarchy. Supports up to six levels deep.

![resource-hierarchy](https://docs.microsoft.com/en-us/training/wwl-azure/describe-core-architectural-components-of-azure/media/management-groups-subscriptions-dfd5a108.png)

### Resources and Resource Groups

- Resource Group - logical container for resource group
- grouping typically based on Lifecycle of the resource
- Azure Resource Manager - control plane for the resources. Manage the resources through declarative templates. Template defined in JSON file.

## Compute

### Azure VMs

- IaaS offering in Azure. Configure, update and maintain sw running in the VMs. Support for multiple OS (Windows, Linux, Oracle (Solaris?), IBM (z/OS?)
- Ideal for solutions requiring total control over the OS, custom software
- Use preconfigured image for rapid provisioning
- Grouping of VMs - Scale Sets and Availability Sets
  - Scale Sets - create and manage group of identical, load balanced VMs. Scale Set comes with load balancer. Automatic scale up or down.
  - Availability Sets - to ensure staggered updates, and fault isolation in network and power connectivity
  - VMs grouped by fault-domain (failure in network/power in one domain doesn’t affect VMs in another fault domain) and update-domain (updates performed on one group at a time).
  - No additional cost for configuring availability set. Pay only for VMs
- Use cases - dev/test, lift-and-shift of on-prem servers, extending on-prem datacenter to cloud by creating a virtual network and placing Azure VMs under that network
- Common VM resources - Size (purpose, cores, processors), Storage disks (for persistence), networking (virtual network, public IP, port configurations)
- Azure Virtual Desktop - Cloud hosted Windows desktop. Data and apps separated from the underlying hardware. Supports MFA and granular RBAC. Also supports multi session Windows 10 or 11 with the Enterprise version.

### Azure Containers

- [Azure Container Instances](https://docs.microsoft.com/en-us/azure/container-instances/container-instances-overview) - PaaS offering to run containers in the cloud. Best suited for microservice based architecture
- supports linux or windows container images from Docker hub, Azure Container Registry or any other hosted container registry
- containers can be grouped into Container Groups, exposed with a FQDN (customlabel.azureregion.azurecontainer.io), custom label provided at the time of container creation
- Container groups includes group of containers configured to run on the same host
- ACI supports hypervisor-level security (secured multi-tenancy), custom sizes (CPU, Memory)
- Native support of Azure File Shares to provide persistence to the container. File Shares can be automatically mounted
- When deployed in a virtual network, ACI instances can communicate with other resources in the same virtual network
- Some functionalities (volume mounts, multiple container per container groups, virtual network) are supported only on linux containers

![Container-groups](https://docs.microsoft.com/en-us/azure/container-instances/media/container-instances-container-groups/container-groups-example.png)

### Azure Kubernetes Service

- [AKS](https://docs.microsoft.com/en-us/azure/aks/intro-kubernetes) overview

### Azure Functions

- Event driven. Serverless compute. (Similar to AWS Lambda, GCP Cloud Functions)
- Automatic scaling. Pay only for the resource (CPU time) consumed when the function runs.
- Can be stateless (default) or stateful (aka Durable functions) where a context (with the state information) is passed to each invocation.

### Azure App Service

- PaaS offering to host web applications, background jobs, REST APIs, mobile backend.
- Supports windows and linux environments. Multiple languages supported (.NET, .NET Core, Java, Ruby, Node.js, PHP, or Python.)
- Supports automated deployments from source repo (e.g. Github)
- Type of app services - web apps, api apps, web jobs, mobile apps

### Azure Virtual Desktop

## Networking

### Azure Virtual Networking

- Virtual network and subnets to enable Azure resources connect each other, connect to Internet, connect to on premises resources
- Provides isolation and segmentation through proper subnetting.
- Enable internet communications through public IP or by keeping the resource behind a load balancer
- Secure communication between Azure resources - not just VMs, but other resources like App Service, AKS, VM Scale Sets
- Communicate with on-prem resources
  - Point-to-site - a client (from on-prem?) initiates a secured VPN connection to the Azure Virtual Network
  - Site-to-site - extend on-prem VPN to connect to Azure VPN gateway, bringing in AZ resources in the same network as the cor network
  - Express Route - dedicated private connection on the Microsoft backbone network, not through internet.
- Route traffic using routing tables on the virtual network. Supports BGP with Azure VPN gateway.
- Filter traffic using network security groups (NSG) with defined inbound and outbound rules, or dedicated Network Appliances (specialized VM) running firewall or WAN optimizations.
- Connect virtual network using virtual network peering.  Peered traffic runs on the Microsoft backbone. Able to make a global network on the Azure network using network peering between multiple virtual network that spans regions across geographies.
- UDR - User Defined Route - to fine grained control over the routing tables between subnets within virtual network or between virtual networks.

### Azure VPN

- VPN gateways deployed in dedicated subnet of the virtual network
- Used to
  - connect on-prem data center to Azure virtual network through site-to-site connection
  - Connect devices to virtual network securely through point-to-site connections
  - Network-to-network connection
- Type of VPN
  - Policy based - IP address of the tunnel is specified statically
  - Route based - modeled after IPSec tunnel. Specify which network (or virtual network) interface to use for tunneling. Preferred VPN type. Use it for connections between virtual networks, point-to-site connections, multi-site connections, coexistence with Azure Express Route Gateway (for failover reasons)
- High Availability
  - Active/Standby - default option. Azure provisions a standby VPN by default. Failover in few seconds for planned events, under 90 seconds for unplanned disruption
  - Active/Active - enabled by the support of BGP. Assign a public IP address for each instance. Create separate tunnels to each IP address.
  - Express Route Failover - in the event of physical failures on the express route connection, failover to use secured VPN traffic (which goes over the internet).
  - Zone redundant gateways - in regions supporting multiple availability zones, VPN gateways and Express Route Gateways can be deployed in zone redundant configuration. Zone redundant gateways use Standard public IP address instead of Basic IP address. (For difference between standard and basic, see here)

### Azure ExpressRoute

- Expand on-prem network into Microsoft cloud (Azure and other Microsoft services like 365, Dynamics 365) over a private connection (Express Route Circuit)
- Some benefits => connectivity to MS cloud services in all geopolitical regions, global connectivity through ExpressRoute Global Reach, built in redundancy
- Services that work with Express Route => Office 365, Dynamics 365, Azure Compute and Network services, Azure Storage, Databases
- Global connectivty - connect branch offices across the world together by connecting through ExpressRoute circuits using Global Reach. Enables on-prem applications to connect across the world without going over the internet.
- Dynamic routing - enabled by BGP support.
- Connectivity modes
  - Colocation at cloud exchange
  - Point to point ethernet connection (point to point connection between on-prem to azure)
  - Any-to-any connection (adding azure network into the existing WAN network)
  - Directly from ExpressRoute sites (connection from a peering facility. Direct dual 100Gpbs or 10Gbps connectivity)
- Security considerations => Data traffic on ExpressRoute do not go over the internet. However some essential traffic like DNS queries, CRL checks, Azure CDN are still sent over the public internet.

### Azure DNS

- Managed DNS hosting service from Azure
- DNS domains hosted on the Azure global network of name servers. Uses AnyCast networking (so query is answered by closest available server)
- Based on Azure Resource Manager. Supports RBAC, activity logs, resource locking and many other features applicable to a resource.
- Can manage DNS records for Azure services and external resources
- Supports private domains too (can use custom domain names in private virtual networks, avoiding the use of azure provided names)
- Supports Alias record sets.
- Can’t buy a domain name though. That can be done by App Service domain or third party domain name registrar.

## Storage

### Azure Storage Services

#### Basics

- Azure Storage Account - unique name space for all data stored under the service types. What services offered by Azure Storage Services?
  - Blob storage
  - Data lake Storage
  - Azure Files (for NFS/CIFS)
  - Queue Storage (for messaging services)
  - Table Storage (no-sql key-value store, different from cosmos DB?)
- Resource endpoint looks like this `https://<storage-account-name>.file.core.windows.net`. Account naming has some restrictions. 3-24 characters, only numbers and lowercase, naming must be unique across Azure.
- Storage account type determines which of the above services are supported and the redundancy levels offered.
- What are the available storage account types?
  - Standard General Purpose V2 - most common. supports all services. Supports all redundancy types.
  - Premium block blobs - for SAN workloads (for block blobs and append blobs).
  - Premium file shares - for NFS/CIFS workloads
  - Premium page blob - for VM disks (workloads in 512b pages), closer to real disks
- What redundancy levels offered?
  - Locally Redundant Storage (LRS) - synchronous replication within the local data center. 3 copies. 11 nines of durability.
  - Zone Redundant Storage (ZRS) - synchronous replication across the availability zones within a region. 3 copies. 12 nines of durability. (note: not all regions support availability zones). Recommended for applications requiring high availability, compliance in storing data within a region. data available for read and write in the event of a zone failure.
  - Geo Redundant Storage (GRS) - synchronous replication within a single location in the primary region. 3 copies. async replication to secondary region using LRS. 16 nines of durability.
  - Geo ZRS (GZRS) - sync replication to 3 availability zones in the primary region. async replication to secondary using LRS. 16 nines of durability.
  - Read-Access Geo Redundant Storage (RA-GRS) - GRS for read heavy workloads.
  - Read-Access Geo ZRS (RA-GZRS) - GZRS for read heavy workloads.
- LRS and ZRS offered in the primary region. For protection against regional failures, choose geo redundant type.
  - Data replicated to a secondary region (determined by the region pairs. can’t be changed).
  - GRS ~= LRS in two region, GZRS ~= ZRS in two regions.
  - must failover to secondary region to gain read and write access (doesn’t seem to be supporting automatic failover) similar to any DR solution.
  - RPO of < 15 minutes.

#### Storage Services (Blob, Files, Queue, Table, Disk)

- Blob storage - Azure’s object storage.
  - can hold all kinds of data. Ideal for serving images, documents, storing files for distributed access, streaming media, backup and restore, data analytics.
  - Access through HTTP/HTTPS, REST, Azure Shell, CLI or client libraries.
  - Storage tiers
    - Hot - frequently accessed data. High storage cost, low retrieval cost.
    - Cool - infrequently accessed, stored for at least 30 days. Low storage cost, high access cost and lower SLA
    - Archive - rarely accessed, stored for at least 180 days. Lowest storage cost, stored offline (meaning tapes?), highest access cost. Ideal for long term backups.
    - Only hot/cool tier can be set at the account level. Set archive option at the blob level (?)
- Azure Files
  - Managed file shares, accessible via NFS and SMB protocols.
  - Concurrent access from cloud or on-prem applications. SMB supported on windows, linux and macOS clients, NFS on Linux, macOS clients.
  - Azure File Sync caches the file shares in Windows Server. For better data affinity.
  - Create and manage file shares through cmdlets, Azure CLI, Azure Portal or Azure Storage Explorer.
  - Data in the file shares can be accessed via file system I/O calls. Can also use Azure Storage Client libraries or Azure Storage REST API.
- Queue storage - highly scalable message queue. Max individual message size 64KB. sample use case: user submits a form on the webpage. App adds that to a message queue. Message queue triggers an event which invokes a function in Azure Functions.
- Disk storage - block storage for VMs. (#doubt aka Page Blobs?)

JSON view of the storage account created in the sandbox.

```json
{
    "sku": {
        "name": "Standard_LRS",
        "tier": "Standard"
    },
    "kind": "StorageV2",
    "id": "/subscriptions/8672fcd9-4133-4ab5-a899-acba167f727c/resourceGroups/learn-f2cd4ff8-0ffd-40ae-9172-7ad1b137d617/providers/Microsoft.Storage/storageAccounts/deepazurestorage",
    "name": "deepazurestorage",
    "type": "Microsoft.Storage/storageAccounts",
    "location": "eastus",
    "tags": {},
    "properties": {
        "minimumTlsVersion": "TLS1_2",
        "allowBlobPublicAccess": true,
        "allowSharedKeyAccess": true,
        "networkAcls": {
            "bypass": "AzureServices",
            "virtualNetworkRules": [],
            "ipRules": [],
            "defaultAction": "Allow"
        },
        "supportsHttpsTrafficOnly": true,
        "encryption": {
            "services": {
                "file": {
                    "enabled": true,
                    "lastEnabledTime": "2022-09-12T15:25:25.6484294Z"
                },
                "blob": {
                    "enabled": true,
                    "lastEnabledTime": "2022-09-12T15:25:25.6484294Z"
                }
            },
            "keySource": "Microsoft.Storage"
        },
        "accessTier": "Hot",
        "provisioningState": "Succeeded",
        "creationTime": "2022-09-12T15:25:25.5233582Z",
        "primaryEndpoints": {
            "dfs": "https://deepazurestorage.dfs.core.windows.net/",
            "web": "https://deepazurestorage.z13.web.core.windows.net/",
            "blob": "https://deepazurestorage.blob.core.windows.net/",
            "queue": "https://deepazurestorage.queue.core.windows.net/",
            "table": "https://deepazurestorage.table.core.windows.net/",
            "file": "https://deepazurestorage.file.core.windows.net/"
        },
        "primaryLocation": "eastus",
        "statusOfPrimary": "available"
    }
}
```

### Data migration

Two options available to migrate on-prem data (takes different form here: raw data, VMs, databases, applications etc.) to Azure.

#### Azure Migrate - unified platform to track migration from on-prem to Azure

- Comes with many integrated tools to manage the migration
- Discovery and Assessment - to assess the readiness of on-prem servers running on VMware, Hyper-V and physical servers
- Server Migration - to do the actual migration of the physical, virtualized servers in on-prem and other public cloud.
- Data Migration Assistant - to assess SQL servers
- Database Migration Service - migrate on-prem databases to Azure VMs with SQL server, Azure SQL or SQL managed instances.
- WebApp migration assistant - to migrate on-prem web applications to Azure App Service
- Azure Databox - move large amounts of offline data to Azure.

#### Azure Data box - physical migration service to transfer large amounts of data to/from azure

- Data contained in a proprietary Data Box storage device. Max storage capacity 80TB.
- End-to-end tracking provided in the Data box service of Azure Portal.
- Common use cases
  - Transfer data larger than 40TB from the places with limited to no network connectivity
  - Moving media library from offline tapes into Azure
  - Migrating VM farms to Azure
- Data movement - one time, periodic or initial bulk transfer followed by periodic transfer.
- Export data out of Azure in to on-prem as well. Some common use cases => DR, compliance, migration of applications back into on-prem or other cloud providers.
- Data in Data Box wiped in accordance with NIST 800-88r1 standards

#### File movement (AzCopy, Azure Storage Explorer, Azure File Sync)

- AzCopy
  - cmdline tool to copy blobs/files to and from a storage account.
  - Copy & sync files between storage account. No bi-directional support.
  - Can move files back and forth between Azure and other clouds as well.
- Azure storage explorer (https://azure.microsoft.com/en-us/products/storage/storage-explorer/#overview)
  - Standalone GUI app to manage files and blobs. Uses AzCopy under the hood.
- Azure File Sync
  - Centralize file shares in Azure Files with the compatibility of a Windows Server.
  - Install File Sync in Windows Server to take advantage of this. Allows bi-directional sync.
  - Turns a Windows Server into mini CDN.

## Azure Architecture and Services

### Idenity, Access and Security

#### Azure Active Directory - directory service to sign into MS cloud, other cloud applications and possibly on-prem

- enables signing in to Microsoft cloud applications including Azure, Office 365, Dynamics and other cloud applications configured to use Microsoft authentication
- Same functionality (identity and access mgmt) provided by Active Directory running on Windows Server in on-prem deployments
- In on-prem AD deployments, Microsoft doesn't monitor the sign-ins. With Azure AD, Microsoft monitors it and provides additional services and protections (threat detection, suspicious sign-ins etc.)
- Azure AD services - *authentication, Single-Sign-On (SSO), App management, Device management with Intune*
- In hybrid cloud environment, maintaining two identity services (AD in on-prem and Azure AD in Cloud) is not recommended. Can combine that into one using **Azure AD Connect**. Sync changes between both the systems

#### Azure Active Directory Domain Services - managed domain services

- eliminates the need to deploy, manage and patch Windows Server domain controllers (DC) in the cloud. #question. What are domain controllers? - a brief one-liner explanation in the [glossary](https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-authsod/64781df1-ee20-413e-b8c5-6511c90dbc30#gt_76a05049-3531-4abd-aec8-30e19954b4bd)
- legacy apps using domain services in on-prem can be lift-and-shifted into cloud with the help of Azure AD DS
- integrates with Azure AD tenant
- An unique namespace is defined per Azure AD DS managed domain. Azure deploys two Domain Controllers in the chosen region and manages it for us.
- One way sync between Azure AD and Azure AD DS.

#### Authentication Methods

- supports multiple authentication methods - password based, SSO, MFA and passwordless (FIDO2)
- SSO is only as strong as the initial sign in
- MFA - for two factor authentication using tokens, code, authenticator apps like Google Authenticator, Microsoft Authenticator
- MFA - must satisfy two of the below three challenges
  - something user knows (e.g. password)
  - something user has (e.g. phone)
  - something user is (e.g. fingerprint, facial ID)
- passwordless - most convenient and secure. Azure offers passworless sign in through Windows Hello for Business, Authenticator App and FIDO2 security keys
- [FIDO](https://fidoalliance.org/) - **F**ast **IDentity** **O**nline standard - leverages standard sign in without username or password by using external security keys (or platform key built into the device such as phone or laptop)

#### Azure AD External Identities - secure interaction with users outside the org

- for secure collaboration with external users (such as partners, distributors, consumers) in sharing and managing the Azure resources
- users bring in their own identity such as govt issued digital ID, Google or Facebook (unmanaged social ID) IDs. Identity managed by the external provider, access managed and protected by us using Azure AD or AD B2C. What capabilities does it provide?
  - B2B Collaboration - to collaborate between external users and business using Azure
  - B2B Direct Connect - to collaborate between two Azure AD organization (e.g. sharing Teams channel between two different organization)
  - B2C - use Azure AD for identity and access management in apps published on Azure

#### Azure Conditional Access - allow or deny access to resources based on identity signals

- manage access to applications, resources based on factors like users' role, location, network, device etc.
- require access only through approved client applications. e.g. Allow access to business email only from Outlook and not any other third party email clients

#### Azure RBAC

- allow access based on the principle of least privilege
- use built-in roles (with common access rules) or define own roles (and associate set of access permissions)
- ![relationship between role and scope](https://docs.microsoft.com/en-us/training/wwl-azure/describe-azure-identity-access-security/media/role-based-access-scope-4b12a8f3.png)
- some built in roles - Reader, Owner, Contributor
- RBAC applied to a scope (*mgmt group, subscription, resource group, resource*). Each scope inherits from RBAC permissions from its parent scope
- [ARM](#azure-resource-manager---deployment-and-management-service-for-azure) enforces the RBAC permissions when processing the requests to Azure resources
- RBAC follows **allow-model**.

#### Zero Trust Model

- Adopt Zero Trust security model
- Guiding principles
  - *Verify explicitly* - authenticate and authorize at all points
  - *Least privilege access* - Just in Time (JIT), Just Enough Access (JEA)
  - *Assume breach* - minimize blaste radius and segment access
- moving from trusted secure network to a network with centralized policy enforcing authentication and authorization at all points.

#### Defense-in-Depth - strategy to slow the advance of attack aimed to access data

- Layered strategy to protect the most precious data
- Layers: Physical, Identity & Access, Perimeter (for DDoS prevention), Network, Compute, Application, Data
- Much of these are self explanatory and common sense
- ![Defense-in-Depth-layers](https://docs.microsoft.com/en-us/training/wwl-azure/describe-azure-identity-access-security/media/defense-depth-486afc12.png)

#### Defender for Cloud - monitoring tool for security posture management and threat protection

- provides the tools to **harden resources, track security posture, protect against attacks and streamline security management**, not only in Azure but also in hybrid and multi-cloud deployments
- native service in Azure. enabled by default in many Azure services
- Azure machines have the log Analystics Agent deployed by default to gather security related data
- Extend to hybrid and multi cloud with the help of [Azure Arc](#azure-arc---extend-azure-compliance-and-monitoring-to-hybrid-and-multi-cloud-environments). An overview [here](https://docs.microsoft.com/en-us/azure/azure-arc/overview)
- Native protections across many services
  - Azure PaaS - such as Azure App Service, Azure SQL, Azure Storage Accounts. can also do anomaly detection on Azure Activity Logs using Defender for Cloud Apps
  - Data services - perform classification of data in Azure SQL to find potential vulnerabilities
  - Networks - limit exposure to brute force attacks. reduce access to VM ports, Just-In-Time access, allow only authorized users, source IP address ranges
- **CSPM** - Cloud Security and Posture Management to extend into multi cloud environment. Agentless plan to assess other cloud (e.g. AWS) resources according to their security recommendations.
- Defender for Kubernetes - extend container threat detection and defenses to Amazon EKS Linux Cluster
- Defender for servers - to add threat detection and advanced features to Windows and Linux EC2 machines
- Strategy followed => **Continuously Assess + Secure + Defend**
- Defender for Cloud built on top of Azure Policy controls.. makes it easy to run on Azure scopes
- Azure Security Benchmark - MS authored, Azure specific set of guidelines to ensure security and compliance
- View the security health through the ![Security Score](https://docs.microsoft.com/en-us/training/wwl-azure/describe-azure-identity-access-security/media/defender-for-cloud-d47a71d8.png)
- Generates Alerts upon threat detection
- Detection followed by protection (suggests remediation, trigger logic app in response), includes fusion kill-chain analysis

## Fundamentals - Management and Governance

### Cost Management

- shift in costs - from CapEx heavy to OpEx heavy
- parameters affecting the costs
  - resource params - type, size, perf, region etc. can affect the costs
  - consumption - pay as-you-go consumption. Discounts with Azure Reservations for commited use
  - maintenance - make sure to clean up attached resources. e.g. VM provisioned with virtual networks, storage, db instances etc.. clean up the associated resources when deleting a VM. (tip: manage under resource groups)
  - geography - region, country, taxes etc.
  - subscription type - some subscription includes allowances
  - azure market place products

#### Pricing Calculator - estimate cost of provisioning Azure resources

- [Pricing Calculator in Azure](https://azure.microsoft.com/en-us/pricing/calculator/)
- get estimated cost of individual resources or build a solution involving multiple resources in Azure

#### TCO Calculator - compare cost of running on-prem infra vs Azure Cloud Infra

- provide the current infrastructure configuration. (compute, storage, network, bandwidth consumption etc), define workloads, assumptions for power, labor, IT costs. Azure computes the cost of that infrastructure in running in on-prem vs Azure Cloud
- [TCO on Azure](https://azure.microsoft.com/en-in/pricing/tco/calculator/)

#### Cost Management Tool - check resource costs, create alerts, budgets

- Azure service to quickly check costs, create alerts, track spending, budget spending and create automation to manage costs
- ![Sample Cost Analysis](https://docs.microsoft.com/en-us/training/wwl-azure/describe-cost-management-azure/media/cost-analysis-b52dedab.png)
- supports multiple alert types - *budget alerts, credit alerts, spending quota alerts*
- Creating budgets - defined by cost when created from Portal. can be defined by cost or consumption when using **Azure Consumption API**
- Budgets can be set based on subscription, resource group, service type. optionally configure the budget conditions to trigger automation to suspend or modify the resource as needed

#### Tracking cost of resources organized by tags

- Tags are another way to organize resources, finer grained than resource groups.
- Some ways to categorize using tags
  - Resource management - to locate and act on resources associated with specific workloads, BUs etc.
  - Cost management
  - Operations management - group resources based on operational availability and criticality. Formulate SLA.
  - Security and governance tags
  - workload optimization and automation
- Tags can be added, modified or deleted through PowerShell, Azure. CLI, ARM templates, REST API or Azure Portal

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
- manage infrastructure as code. specify the desired state of the resources in declarative templates ([ARM Templates](https://docs.microsoft.com/en-us/azure/azure-resource-manager/templates/overview). Much like managing kubernetes resources with spec files.
- dependencies can be specified in the template. Azure takes care of the ordering and invoking the right tools to deploy the resources
- ARM templates can be modular, nested, and also extended (with PowerShell or Bash scripts inline or externally sourced)
- [Sample ARM template to create ubuntu VM](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/quick-create-template?toc=%2Fazure%2Fazure-resource-manager%2Ftemplates%2Ftoc.json)
- note: prior to the introduction of ARM in 2014, there existed the classic deployment model. Resources existed independently and required external coordination in managing multiple resources. more on that [here](https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/deployment-models)

### Monitoring Tools

#### Azure Advisor - recommendations to optimize the cloud environment

- provides recommendations to improve *reliability, security, performance, operational excellence and cost reduction*.
- ![Dashboard](https://docs.microsoft.com/en-us/training/wwl-azure/describe-monitoring-tools-azure/media/azure-advisor-dashboard-baca22e2.png)
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

## Resources

- [About AZ-900](https://docs.microsoft.com/en-us/certifications/exams/az-900)
- [Understanding block blobs, append blobs, and page blobs](https://docs.microsoft.com/en-us/rest/api/storageservices/understanding-block-blobs--append-blobs--and-page-blobs)
- [Azure Data Box, Data Box Disk, Data Box Heavy](https://docs.microsoft.com/en-us/azure/databox/)
- [Cloud Adoption Framework](https://docs.microsoft.com/en-us/azure/cloud-adoption-framework/overview)
- [Network Security Groups](https://docs.microsoft.com/en-us/azure/virtual-network/network-security-groups-overview)
- [Application Security Groups](https://docs.microsoft.com/en-us/azure/virtual-network/application-security-groups)
- [Subscription and Service Limits](https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/azure-subscription-service-limits)