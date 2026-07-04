#!/usr/bin/env bash

set -euo pipefail

artifact_dir="${RELEASE_ARTIFACT_DIR:-${BUILD_OUTPUT_DIR:-dist}}"
tag="${RELEASE_TAG:-$(mise run release:tag)}"

GOOS="${GOOS:-$(go env GOOS)}"
GOARCH="${GOARCH:-$(go env GOARCH)}"

mapfile -t binaries < <(find "${artifact_dir}" -maxdepth 1 -type f ! -name '*.tar.gz' ! -name '*.zip' -print | sort)

case "${#binaries[@]}" in
  0)
    echo "No build artifact was found in ${artifact_dir}"
    exit 1
    ;;
  1)
    binary_path="${binaries[0]}"
    ;;
  *)
    echo "Expected exactly one build artifact in ${artifact_dir}, found ${#binaries[@]}"
    printf '  %s\n' "${binaries[@]}"
    exit 1
    ;;
esac

binary_name="${binary_path##*/}"
package_name="${binary_name%.exe}-${tag}-${GOOS}-${GOARCH}"
if [[ -n "${RELEASE_ARTIFACT_EXT:-}" ]]; then
  artifact_ext="${RELEASE_ARTIFACT_EXT}"
elif [[ "${GOOS}" == "windows" ]]; then
  artifact_ext="zip"
else
  artifact_ext="tar.gz"
fi

case "${artifact_ext}" in
  zip)
    (cd "${artifact_dir}" && zip -j "${package_name}.zip" "${binary_name}")
    ;;
  tar.gz)
    tar -C "${artifact_dir}" -czf "${artifact_dir}/${package_name}.tar.gz" "${binary_name}"
    ;;
  *)
    echo "Unsupported release artifact extension ${artifact_ext}"
    exit 1
    ;;
esac
