# github_cli


github_cli is a command-line application that allows\
	you to fetch various data related to a target user from GitHub. \
	It provides convenient and efficient ways to retrieve information \
	such as pull requests, repositories, followers, and more. 

Usage:
  - github_cli [command] [target user]

Available Commands: 
  - completion  Generate the autocompletion script for the - specified shell
  - following   Get repos followed by the target user
  - help        Help about any command
  - prs         Get a list of pull requests for the target user
  - repos       Retrieve information about repositories owned by the target user

Flags:\
      ---config string   config file (default is $HOME/.github_cli.yaml) \
  -h, --help            help for github_cli \
  -t, --toggle          Help message for toggle 

Use "github_cli [command] --help" for more information about a command.



# Small investigations

1. What is terraform and why should we use it? 
  
  
  R: Terraform is an open-source infrastructure as code (IaC) tool developed by HashiCorp. It enables you
   to define and manage your infrastructure in a declarative manner using a simple and human-readable 
   configuration language. With Terraform, you can create, update, and destroy infrastructure resources 
   across various cloud providers and on-premises environments.

   Reasons to use:

  - Improve speed: Automation is faster than manually navigating an interface when you need to deploy
   and/or connect resources.

  - Improve reliability: If your infrastructure is large, it becomes easy to misconfigure a resource or
    provision services in the wrong order. With terraform, the resources are always provisioned and configured
    exactly as declared.

  - Prevent configuration drift: Configuration drift occurs when the configuration that provisioned your
    environment no longer matches the actual environment.

  - Support experimentation, testing, and optimization: Because Terraform makes provisioning
    new infrastructure so much faster and easier, you can make and test experimental changes without investing
    lots of time and resources; and if you like the results, you can quickly scale up the new infrastructure
    for production.

  - Open source: Terraform is backed by large communities of contributors who build plugins to the platform.
    Regardless of which cloud provider you use, it’s easy to find plugins, extensions, and professional support.
    This also means Terraform evolves quickly, with new benefits and improvements added consistently.

  - Supports a wide range of cloud providers: can be used with different cloud services provider. Most other IaC tools are
    designed to work with single cloud provider.

  - Immutable infrastructure: Most Infrastructure as Code tools create mutable infrastructure, meaning the
    infrastructure can change to accommodate changes such as new storage server. The danger with mutable infrastructure is configuration drift—as the changes pile up, the actual provisioning
    of different servers or other infrastructure elements ‘drifts’ further from the original configuration,
    making bugs or performance issues difficult to diagnose and correct. Terraform provisions immutable
    infrastructure, which means that with each change to the environment, the current configuration is replaced
    with a new one that accounts for the change, and the infrastructure is reprovisioned. Even better, previous
    configurations can be retained as versions to enable rollbacks if necessary or desired.

2. How to guarantee that no manual changes are done in cloud projects deployed in terraform? 


R:

  There is no way to guarantee the absence of manual changes, but can be minimized by:

  - Immutable Infrastructure: Design your infrastructure using immutable principles, which means that infrastructure changes are made by creating new resources rather than modifying existing ones. This approach helps minimize the risk of manual changes since any modifications require the deployment of new infrastructure.


  - Limited Access: Restrict access to the cloud infrastructure. Only authorized individuals, working with the cloud infrastructure to minimize the possibility of manual changes. Use proper identity and access management (IAM) controls provided by your cloud provider to manage user roles and permissions effectively.

  

3. How can I guarantee naming conventions using terraform?


R:

  - Use Linters: Utilize linter tools specific to Terraform, such as "TFLint." These tools
    analyze Terraform code and configurations, including resource names, and provide feedback on potential
    issues or violations. Configure the linter rules to include checks for naming conventions. 

  - CI/CD Pipeline Checks: Incorporate naming convention checks as part of your CI/CD pipeline. Include a validation
    step that examines the Terraform code for adherence to naming conventions. Fail the pipeline if any violations are
    detected, preventing the deployment of non-compliant infrastructure.

   
4. How to check if secrets are leaked in github repos? 

R:

  - Use GitHub's secret scanning feature. GitHub offers a free secret scanning feature that can be configured to scans
    repositories for leaked secrets and generate alerts. Secret scanning can scan your entire Git history on all branches
    present in your GitHub repository for secrets. 
     
  - Utilize Third-Party Tools: There are various third-party tools available that can scan GitHub repositories for potential
    leaked secrets. These tools analyze the repository's codebase, commit history, and other relevant files to identify
    potential secrets. Some popular tools include GitGuardian, TruffleHog, and Gitleaks. These tools can be integrated into
    your CI/CD pipeline or run locally against specific repositories. 
  
  - Inspect Repository Files: Manually inspect the repository's files for any potential secrets. Look for files that commonly 
    contain sensitive information, such as configuration files, environment files, or scripts. Pay attention to any files that
    contain hard-coded values or suspicious-looking data. 