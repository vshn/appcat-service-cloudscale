# Terrajet Cloudscale Provider

`provider-jet-cloudscale` is a [Crossplane](https://crossplane.io/) provider that is built using
[Terrajet](https://github.com/crossplane/terrajet) code generation tools and exposes XRM-conformant managed resources
for the Cloudscale API.

## Installing the provider

Assuming that you have a working Crossplane installation and kubectl is configured for the K8s cluster running your
Crossplane installation, you can install provider-jet-cloudscale:

* Generate an API token with write permissions in the [Cloudscale web interface](https://control.cloudscale.ch/)
* Put the API token into examples/providerconfig/secret.yaml.tmpl
* Install the secret into K8s:
  ```console
  kubectl apply -f examples/providerconfig/secret.yaml.tmpl
  ```
* Install the provider configuration into K8s. This tells the provider which secret to use.
  ```console
  kubectl apply -f examples/providerconfig/providerconfig.yaml
  ```
* Install the provider. This is achieved by installing the "Crossplane package" with the name
  `provider-jet-cloudscale`, which in turn sets up the controller image `provider-jet-cloudscale-controller` which does
  the real work. You may need to adjust the version number of the "Crossplane package" first in
  `examples/providerconfig/install.yaml`.
  ```console
  kubectl apply -f examples/providerconfig/install.yaml
  ```
* Check if the provider is ready.
  ```console
  kubectl get Provider provider-jet-cloudscale
  ```
  It should both be 'installed' and 'healthy'. If everything is correct the controller is running as the
  'provider-jet-cloudscale-...' pod in the crossplane-system namespace and there are some cloudscale-specific CRDs
  available.

## Using the provider

Assuming the setup is complete you can now create demo user:

```console
kubectl apply -f examples/demouser.yaml
```

You should now see a "DemoUser" appear in the Cloudscale web interface.

```console
kubectl delete User demouser
```

The "DemoUser" should now disappear from the Cloudscale web interface.

## Developing and testing the provider

The following commands may be useful for developing and testing the provider.

Create a local kubernetes-in-docker test and development environment:

```console
cd kindev
make crossplane-setup
```

Clean up said development environment:

```console
cd kindev
make clean
```

Re-generate CRDs and Go code from the Terraform module:

```console
make generate
```

Re-generate the Crossplane package and operator (controller):

```console
make build
```

Run the Crossplane operator (controller) directly:

```console
make run
```

## Licensing

provider-jet-cloudscale is under the Apache 2.0 license.
