# pulumi-convert
This is an attempt to use the `pulumi convert --from terraform --language go --out <output_path>`
## Requirements

- pulumi cli
- go

## Trying it out

Clone the repo

Create a local pulumi stack
```
pulumi login --local

pulumi stack init

pulumi config set apiToken <CAST_API_KEY> --secret

pulumi preview

pulumi up
```
Doesn't work yet, still need to fix alot of things.

