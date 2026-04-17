# checks-types

Shared Kubernetes CRD types for the [Best Practices for Kubernetes](https://github.com/redhat-best-practices-for-k8s) checks ecosystem.

This module provides the `bps.openshift.io/v1alpha1` API types used by both the [bps-operator](https://github.com/sebrandon1/bps-operator) and [checks-qe](https://github.com/redhat-best-practices-for-k8s/checks-qe), so neither project needs to import the other.

## Types

- **BestPracticeScanner** — CR that triggers a best-practice scan against a target namespace
- **BestPracticeResult** — per-check compliance result produced by a scan

## Usage

```go
import bpsv1alpha1 "github.com/redhat-best-practices-for-k8s/checks-types/api/v1alpha1"
```

Register the types with your controller-runtime scheme:

```go
import bpsv1alpha1 "github.com/redhat-best-practices-for-k8s/checks-types/api/v1alpha1"

scheme := runtime.NewScheme()
bpsv1alpha1.AddToScheme(scheme)
```
