# asset_bundle_hash

Tool to calculate Unity AssetBunlde hash

The hash calculation method is as follows

1. Read asset list from manifest file
1. Calculate SHA256 sums of each asset
1. Print AssetBundle SHA256 sums that calculate from Path and SHA256 sums of eash asset

# How to use

```sh
$ asset_bundle_hash AssetBundles/sample.manifest
678e851755589c2ed8905b10e0e8302a694af7e40ee86abacfa4bddfba859bb1

$ asset_bundle_hash --project-root ./UnityProjectSample ./UnityProjectSample/AssetBundles/sample.manifest
678e851755589c2ed8905b10e0e8302a694af7e40ee86abacfa4bddfba859bb1
```
