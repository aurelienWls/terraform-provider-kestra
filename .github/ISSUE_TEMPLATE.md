Hi there,

Thank you for opening an issue. Please note that we try to keep the Kestra Terraform provider issue tracker reserved for bug reports and feature requests.

### Terraform Version & Kestra Terraform provider version
- Run `terraform -v` to show the version. 
- Provide the `required_providers` configuration :
```hcl
terraform {
  required_providers {
    kestra = {
      source = "kestra-io/kestra"
      version = "0.1"
    }
  }
}
```

### Affected Resource(s)
Please list the resources as a list, for example:
- kestra_flow
- kestra_template

If this issue appears to affect multiple resources, it may be an issue with Terraform's core, so please mention this.

### Terraform Configuration Files
```hcl
# Copy-paste your Terraform configurations here - for large Terraform configs,
# please use a service like Dropbox and share a link to the ZIP file. For
# security, you can also encrypt the files using our GPG public key.
```

### Debug Output
Please provider a link to a GitHub Gist containing the complete debug output: https://www.terraform.io/docs/internals/debugging.html. Please do NOT paste the debug output in the issue; just paste a link to the Gist.

### Panic Output
If Terraform produced a panic, please provide a link to a GitHub Gist containing the output of the `crash.log`.

### Expected Behavior
What should have happened?

### Actual Behavior
What actually happened?

### Steps to Reproduce
Please list the steps required to reproduce the issue, for example:
1. `terraform apply`

### Important Factoids
Are there anything atypical about your Kestra installation that we should know? For example: Running behind a reverse proxy? Running on GKE, Docker?

### References
Are there any other GitHub issues (open or closed) or Pull Requests that should be linked here? For example:
- Relate issue #1234
