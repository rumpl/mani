# mani

Shows the index of an image

## Compile

```shell
$ make
```

## Use

```shell
$ ./mani rumpl/loge:latest
{
  "schemaVersion": 2,
  "manifests": null
}
$ ./mani nginx:latest
{
  "schemaVersion": 2,
  "manifests": [
    {
      "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
      "digest": "sha256:8ff4598873f588ca9d2bf1be51bdb117ec8f56cdfd5a81b5bb0224a61565aa49",
      "size": 1362,
      "platform": {
        "architecture": "amd64",
        "os": "linux"
      }
    },
    {
      "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
      "digest": "sha256:092368e66229d2df31f2f6980d4abe630bc4576cc830485ce9aec1ad6ee39a7b",
      "size": 1362,
      "platform": {
        "architecture": "arm",
        "os": "linux",
        "variant": "v5"
      }
    },
    {
      "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
      "digest": "sha256:2b94f09d2d25fffb610349b7e7d26c0b1bc6a7d55795906a34e346c52c464fd1",
      "size": 1362,
      "platform": {
        "architecture": "arm",
        "os": "linux",
        "variant": "v7"
      }
    },
    {
      "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
      "digest": "sha256:565c5b948bc1162b1b8bb1699830e79345a5fe06ce7a6e97a137b9837e9422e9",
      "size": 1362,
      "platform": {
        "architecture": "arm64",
        "os": "linux",
        "variant": "v8"
      }
    },
    {
      "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
      "digest": "sha256:4ce8bb5bca50e30a1b523de188d7538f5f0e17693dda87462c434aa7174d5b57",
      "size": 1362,
      "platform": {
        "architecture": "386",
        "os": "linux"
      }
    },
    {
      "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
      "digest": "sha256:230cf2b89e1f15b73a321a8b7637583cc0c6bbf0f948ef291d1ed844126c4635",
      "size": 1362,
      "platform": {
        "architecture": "mips64le",
        "os": "linux"
      }
    },
    {
      "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
      "digest": "sha256:3d3726baabd4a64466564b8068c1af39ca69f003bd0662f0c30e6253666290ab",
      "size": 1362,
      "platform": {
        "architecture": "ppc64le",
        "os": "linux"
      }
    },
    {
      "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
      "digest": "sha256:c2cb1aebb3e93bed04cee202f88e8c1932169beae7a164cd2b5a5d8b356b1c35",
      "size": 1362,
      "platform": {
        "architecture": "s390x",
        "os": "linux"
      }
    }
  ]
}
```
