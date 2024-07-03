package release

import (
    "sort"

    "github.com/hashicorp/go-version"
)

func tagNameToVersion(tagName string) string {
    // if a tagName starts with `v`, remove it.
    if tagName[0] == 'v' {
        return tagName[1:]
    }

    return tagName
}

func minInt(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// toVersions converts []string to []*version.Version.
// Ignore if parse error.
func toVersions(versionsRaw []string) []*version.Version {
    var versions []*version.Version
    for _, raw := range versionsRaw {
        v, err := version.NewVersion(raw)
        if err != nil {
            continue
        }
        versions = append(versions, v)
    }
    return versions
}

// fromVersions converts []*version.Version to []string.
func fromVersions(versions []*version.Version) []string {
    versionsRaw := make([]string, len(versions))
    for i, v := range versions {
        raw := v.String()
        versionsRaw[i] = raw
    }
    return versionsRaw
}

// sortVersions sort a list of versions in semver order.
func sortVersions(versions []*version.Version) []*version.Version {
    sort.Sort(version.Collection(versions))
    return versions
}

// excludePreReleases excludes pre-releases such as alpha, beta, rc.
func excludePreReleases(versions []*version.Version) []*version.Version {
    // exclude pre-release
    var filtered []*version.Version
    for _, v := range versions {
        if len(v.Prerelease()) == 0 {
            filtered = append(filtered, v)
        }
    }

    return filtered
}
