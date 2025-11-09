# Security Policy

## Supported Versions

This portfolio website is built using Jekyll and hosted on GitHub Pages. The following versions are currently supported with security updates:

| Component | Version | Supported          |
| --------- | ------- | ------------------ |
| Jekyll    | 4.4.1+  | :white_check_mark: |
| Ruby      | 2.6.0+  | :white_check_mark: |
| < 4.4.1   |         | :x:                |

### Dependencies

The following Ruby gems are used in this project and are kept up-to-date:

- `jekyll` (~> 4.4.1)
- `jekyll-sitemap`
- `jekyll-seo-tag`
- `jekyll-theme-minimal`
- `webrick` (~> 1.7)

**Note**: GitHub Pages manages the Jekyll version automatically. For local development, ensure you're using a supported version as specified in the `Gemfile`.

## Reporting a Vulnerability

We take security seriously and appreciate your help in keeping this project secure. If you discover a security vulnerability, please follow these guidelines:

### How to Report

1. **Email**: Send a detailed report to `contact@mohamedlakssirt.tech`
2. **Subject**: Use the prefix `[SECURITY]` in your email subject line
3. **Do NOT** open a public GitHub issue for security vulnerabilities

### What to Include

Please provide the following information in your report:

- **Description**: A clear description of the vulnerability
- **Steps to Reproduce**: Detailed steps to reproduce the issue
- **Impact**: Potential impact and severity assessment
- **Affected Components**: Which parts of the project are affected
- **Suggested Fix**: If you have suggestions for remediation (optional but appreciated)
- **Proof of Concept**: If applicable, include a proof of concept or exploit code

### Response Timeline

- **Acknowledgment**: Within 48 hours of receiving your report
- **Initial Assessment**: Within 7 days
- **Status Updates**: Regular updates every 7-14 days until resolution
- **Resolution**: Depends on severity and complexity, typically within 30 days for critical issues

### What to Expect

- **If Accepted**:
  - You will receive acknowledgment and regular updates
  - A fix will be developed and tested
  - You will be credited in the security advisory (unless you prefer to remain anonymous)
  - The fix will be released in a timely manner

- **If Declined**:
  - You will receive an explanation of why the issue was declined
  - If it's out of scope, we'll direct you to the appropriate party (e.g., GitHub Pages team)

## Scope

### In Scope

The following are considered in scope for security reporting:

- **Jekyll Configuration**: Vulnerabilities in `_config.yml` or Jekyll build process
- **Dependencies**: Security issues in Ruby gems listed in `Gemfile`
- **Site Content**: Malicious or compromised content in markdown files, HTML templates, or assets
- **Build Process**: Issues in the Jekyll build and deployment workflow
- **GitHub Actions**: Security issues in CI/CD workflows (`.github/workflows/jekyll.yml`)
- **Deployment Security**: Vulnerabilities in the deployment process, workflow permissions, or artifact handling

### Out of Scope

The following are **not** considered in scope:

- **GitHub Pages Infrastructure**: Issues with GitHub Pages hosting platform itself
- **Third-Party Services**: Vulnerabilities in external services (e.g., GitHub, external APIs)
- **Social Engineering**: Phishing attacks or social engineering attempts
- **Denial of Service**: DoS attacks that don't exploit a code vulnerability
- **Content Policy Violations**: Spam, harassment, or other content policy issues (report via GitHub's reporting system)
- **Dependency Vulnerabilities**: Issues in dependencies that are already publicly disclosed and have available patches

## Security Best Practices

### Dependency Management

- **Regular Updates**: Dependencies are reviewed and updated regularly
- **Security Advisories**: We monitor RubyGems security advisories and GitHub Dependabot alerts
- **Version Pinning**: Critical dependencies are pinned to specific versions in `Gemfile`

### Content Security

- **Input Validation**: All user-generated content (if any) is validated and sanitized
- **No Sensitive Data**: This repository does not contain sensitive credentials, API keys, or personal information
- **Asset Security**: All assets are reviewed before inclusion

### GitHub Pages Security

- **HTTPS**: GitHub Pages automatically provides HTTPS for all sites
- **Content Security Policy**: Consider implementing CSP headers if adding interactive features
- **Subresource Integrity**: Use SRI hashes for external scripts when applicable

### Deployment Security

This project uses GitHub Actions for automated deployment to GitHub Pages. The following security measures are in place:

- **Workflow Permissions**: The deployment workflow uses minimal required permissions:
  - `contents: read` - Read repository contents
  - `pages: write` - Deploy to GitHub Pages
  - `id-token: write` - Required for OIDC authentication

- **Deployment Process**:
  - **Automated Builds**: All deployments are built in isolated GitHub Actions runners
  - **Environment Protection**: The `github-pages` environment is protected and requires approval for deployments
  - **Concurrency Control**: Only one deployment runs at a time to prevent race conditions
  - **Build Verification**: All builds are verified before deployment

- **CI/CD Security**:
  - **Action Versioning**: GitHub Actions are pinned to specific versions (e.g., `@v5`, `@v4`)
  - **Dependabot**: Automated dependency updates for GitHub Actions workflows
  - **Workflow Validation**: All workflow changes are reviewed before merging
  - **Artifact Security**: Build artifacts are stored securely and automatically cleaned up after deployment

- **Deployment Triggers**:
  - **Automatic**: Deployments trigger on pushes to the `master` branch
  - **Manual**: Workflow can be manually triggered via `workflow_dispatch` for controlled deployments
  - **Security Updates**: Critical security updates may trigger immediate deployment after verification

- **Rollback Procedures**:
  - Previous deployments can be restored via GitHub Pages settings
  - Git history allows reverting to previous stable versions
  - Build artifacts are retained for a limited time for rollback purposes

- **Security Monitoring**:
  - GitHub Actions logs are monitored for suspicious activity
  - Failed deployments are investigated immediately
  - Unauthorized workflow modifications trigger security alerts

### Reporting Security Issues in Dependencies

If you discover a security issue in a dependency (e.g., Jekyll or a Ruby gem):

1. Check if the issue has already been reported upstream
2. If not, report it to the maintainers of that dependency first
3. If the issue affects this project specifically, report it to us as well

## Security Updates

Security updates will be:

- Released as soon as possible after verification
- Documented in the repository's commit history
- Communicated via security advisories if the issue is critical
- Applied to all supported versions
- Deployed automatically via GitHub Actions workflow after merging to the `master` branch
- Tested in isolated build environments before production deployment

## Acknowledgments

We appreciate the security research community's efforts to keep open-source projects secure. Security researchers who responsibly disclose vulnerabilities will be credited (unless they prefer to remain anonymous).

## Contact

For security-related inquiries, please contact:

- **Email**: [contact@mohamedlakssirt.tech](mailto:contact@mohamedlakssirt.tech)
- **GitHub**: [@thejokers69](https://github.com/thejokers69)

---

**Last Updated**: 2025-01-27
