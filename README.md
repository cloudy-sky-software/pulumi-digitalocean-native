# Pulumi Native Provider for Digital Ocean (Preview)

[DigitalOcean](https://digitalocean.com/) helps you build your vision using simple, powerful and cost-effective cloud solutions.

:information_source: This provider uses DigitalOcean's [REST API](https://docs.digitalocean.com/reference/api/) directly.

> This provider was generated using [`pulschema`](https://github.com/cloudy-sky-software/pulschema) and [`pulumi-provider-framework`](https://github.com/cloudy-sky-software/pulumi-provider-framework).

## Why Is This Called DigitalOcean Native?

Despite the fact that all Pulumi providers created by Cloudy Sky Software being "native" Pulumi providers by default, there is a [prior Pulumi provider for DigitalOcean](https://github.com/pulumi/pulumi-digitalocean), albeit a TF-bridged one that is based on the [upstream TF provider](https://github.com/digitalocean/terraform-provider-digitalocean). So this provider had to be renamed to avoid naming conflicts, specifically in language package registries such as PyPi and Nuget where the packages are not namespaced under an organization or user unlike npm.

## Package SDKs

- Node.js: https://www.npmjs.com/package/@cloudyskysoftware/pulumi-digitalocean-native
- Python: https://pypi.org/project/pulumi_digitalocean_native/
- .NET: https://www.nuget.org/packages/Pulumi.DigitalOceanNative
- Go: `import github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn`

## Using The Provider

### API Key

You'll need an personal access token (aka API key). Follow DigitalOcean's [docs](https://docs.digitalocean.com/reference/api/create-personal-access-token/) for creating one.
Then set the token as a secret with `pulumi config set --secret digitalocean-native:apiKey`.

## Releasing A New Version

:info: Switch to the `main` branch first and get the latest `git pull origin main && git fetch`. Check what the last release tag was.

1. Regular releases should just increment the patch version unless a minor or a major (breaking changes) version bump is warranted.
1. Update the `CHANGELOG.md` with notes about what will be included in this release.
1. Commit the changelog with `git commit -am "vX.Y.Z"` or something similar and push `git push origin main`.
1. Tag the commit with the release version by running

   ```bash
   git tag vX.Y.Z
   git tag sdk/vX.Y.Z
   ```

1. Push the tags.

   ```bash
   git push --tags
   ```
