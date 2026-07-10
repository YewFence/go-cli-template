#!/usr/bin/env bash

set -euo pipefail

if [[ "$#" -ne 1 ]]; then
  echo "Usage: $0 <build|publish>"
  exit 1
fi

mode="$1"
case "${mode}" in
  build | publish) ;;
  *)
    echo "Unsupported container release mode ${mode}"
    exit 1
    ;;
esac

lowercase() {
  printf '%s' "$1" | LC_ALL=C tr '[:upper:]' '[:lower:]'
}

tag="${RELEASE_TAG:-}"
version="${BUILD_VERSION:-}"
if [[ "${version}" == v* ]]; then
  echo "BUILD_VERSION must not include the v prefix"
  exit 1
fi

if [[ -n "${tag}" ]]; then
  tag_version="${tag#v}"
  if [[ -z "${tag_version}" ]]; then
    echo "RELEASE_TAG must include a version"
    exit 1
  fi
  if [[ -z "${version}" ]]; then
    version="${tag_version}"
  elif [[ "${version}" != "${tag_version}" ]]; then
    echo "BUILD_VERSION ${version} does not match RELEASE_TAG ${tag}"
    exit 1
  fi
elif [[ -n "${version}" ]]; then
  if [[ "${version}" == "dev" ]]; then
    tag="dev"
  else
    tag="v${version}"
  fi
elif [[ "${mode}" == "build" ]]; then
  version="dev"
  tag="dev"
else
  echo "RELEASE_TAG must be set when publishing"
  exit 1
fi

if [[ "${mode}" == "publish" && "${tag}" != v?* ]]; then
  echo "RELEASE_TAG must use the v prefix when publishing"
  exit 1
fi

repo="${KO_DOCKER_REPO:-}"
if [[ -z "${repo}" ]]; then
  if [[ -n "${GITHUB_REPOSITORY:-}" ]]; then
    repo="ghcr.io/${GITHUB_REPOSITORY}"
  elif [[ "${mode}" == "build" ]]; then
    repo="ghcr.io/example/your-cli-repo"
  else
    echo "KO_DOCKER_REPO or GITHUB_REPOSITORY must be set when publishing"
    exit 1
  fi
fi
repo="$(lowercase "${repo}")"

tags=("${tag}")
if [[ "${mode}" == "build" ]]; then
  platforms="${KO_PLATFORMS:-linux/amd64}"
else
  platforms="${KO_PLATFORMS:-linux/amd64,linux/arm64}"
  tags[${#tags[@]}]="${version}"
fi

if [[ "${mode}" == "publish" ]]; then
  registry="${KO_REGISTRY:-${repo%%/*}}"
  registry="$(lowercase "${registry}")"
  username="${KO_REGISTRY_USERNAME:-}"
  password="${KO_REGISTRY_PASSWORD:-}"

  if [[ -z "${username}" && -n "${password}" && "${registry}" == "ghcr.io" ]]; then
    username="${GITHUB_ACTOR:-}"
  fi
  if [[ -z "${username}" && -z "${password}" && "${registry}" == "ghcr.io" && -n "${GITHUB_TOKEN:-}" && -n "${GITHUB_ACTOR:-}" ]]; then
    username="${GITHUB_ACTOR}"
    password="${GITHUB_TOKEN}"
  fi

  if [[ -n "${username}" || -n "${password}" ]]; then
    if [[ -z "${username}" || -z "${password}" ]]; then
      echo "KO_REGISTRY_USERNAME and KO_REGISTRY_PASSWORD must be set together"
      exit 1
    fi
    printf '%s' "${password}" | ko login "${registry}" --username "${username}" --password-stdin
  fi
fi

args=(
  ./cmd/your-cli
  --bare
)
for image_tag in "${tags[@]}"; do
  args+=(--tags "${image_tag}")
done
args+=(
  --platform "${platforms}"
  --ldflags "-s -w -X main.version=${version}"
)
if [[ "${mode}" == "build" ]]; then
  args+=(--push=false)
else
  args+=(--push=true)
fi
if [[ -n "${KO_IMAGE_REFS_OUTPUT:-}" ]]; then
  args+=(--image-refs "${KO_IMAGE_REFS_OUTPUT}")
fi

KO_DOCKER_REPO="${repo}" ko build "${args[@]}"

if [[ -n "${KO_METADATA_OUTPUT:-}" ]]; then
  tags_csv="$(IFS=,; printf '%s' "${tags[*]}")"
  {
    printf 'repository=%s\n' "${repo}"
    printf 'platforms=%s\n' "${platforms}"
    printf 'tags=%s\n' "${tags_csv}"
  } >> "${KO_METADATA_OUTPUT}"
fi
