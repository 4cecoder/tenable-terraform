# tenable-terraform

## Overview
This document outlines the development roadmap for the `tenable-terraform` custom provider. This provider interfaces with Tenable's API (referenced here: [Tenable API Documentation](https://developer.tenable.com/reference/navigate)) to manage resources like managed credentials, scanner groups, and vulnerability scans.

## Development Roadmap

| Task                                                | Status | Expected Due Date | Notes |
|-----------------------------------------------------|--------|-------------------|-------|
| Develop scaffolding for local development           |        |                   | Basic scaffolding and documentation for local development of the Tenable provider. Includes necessary packages, environment settings, docker setup, etc. |
| Development of internal API for "managed credentials" |        |                   | Should have all CRUD(L) capabilities offered |
| Development of internal API for "ScannerGroups"     |        |                   | Should have all CRUD(L) capabilities offered |
| Development of internal API for "Vulnerability Scan" |        |                   | Should have all CRUD(L) capabilities offered |
| Development of internal API for "web scans"         |        |                   | Should have all CRUD(L) capabilities offered |
| Development of test cases for internal API functions and resources |        |                   |       |
| Development of Terraform provider for "managed credentials" |        |                   | Should have all CRUD(L) capabilities offered |
| Development of Terraform provider for "ScannerGroups" |        |                   | Should have all CRUD(L) capabilities offered |
| Development of Terraform provider for "vulnerability scan" |        |                   | Should have all CRUD(L) capabilities offered |
| Development of Terraform provider for "web scans"   |        |                   | Should have all CRUD(L) capabilities offered |
| Development of test cases for terraform provider and its resources |        |                   |       |
| Generated docs for Terraform provider               |        |                   |       |

## Additional Information
- **Scaffolding**: Details the setup process for a local development environment.
- **Internal API Development**: Involves creating APIs that interact with Tenable's services.
- **Terraform Provider Development**: Focuses on developing the Terraform provider to manage Tenable resources.
- **Testing and Documentation**: Ensures all developments are properly tested and documented.

For further details and updates, refer to the [Tenable API Documentation](https://developer.tenable.com/reference/navigate).
