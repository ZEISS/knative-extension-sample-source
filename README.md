# Knative Eventing Template Source

## Install Sources CRDs for Knative

### Install from Source

Install the CRD providing the control / dataplane for the various `Sources` and `Bindings`:

```bash
# define environment variables accordingly, e.g. when using kind
export KIND_CLUSTER_NAME=example
export KO_DOCKER_REPO=kind.local

ko apply -BRf config
```

### Examples 

Install the [Knative Serving](https://knative.dev/docs/install/yaml-install/serving/install-serving-with-yaml/#prerequisites) before installing the examples.

To see examples of the `Sources` and `Bindings` in action, check out our
[examples](./examples/README.md) directory.