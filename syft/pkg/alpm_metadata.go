package pkg

import (
	"sort"
	"time"

	"github.com/anchore/packageurl-go"
	"github.com/anchore/syft/syft/file"
	"github.com/anchore/syft/syft/linux"
	"github.com/scylladb/go-set/strset"
)

const AlpmDBGlob = "/var/lib/pacman/local/**/desc"

type AlpmMetadata struct {
	BasePackage  string           `mapstructure:"base" json:"basepackage"`
	Package      string           `mapstructure:"name" json:"package"`
	Version      string           `mapstructure:"version" json:"version"`
	Description  string           `mapstructure:"desc" json:"description"`
	Architecture string           `mapstructure:"arch" json:"architecture"`
	Size         int              `mapstructure:"size" json:"size" cyclonedx:"size"`
	Packager     string           `mapstructure:"packager" json:"packager"`
	License      string           `mapstructure:"license" json:"license"`
	Url          string           `mapstructure:"url" json:"url"`
	Validation   string           `mapstructure:"validation" json:"validation"`
	Reason       int              `mapstructure:"reason" json:"reason"`
	Files        []AlpmFileRecord `mapstructure:"files" json:"files"`
	Backup       []AlpmFileRecord `mapstructure:"backup" json:"backup"`
}

// TODO: Implement mtree support
type AlpmFileRecord struct {
	Path         string       `mapstruture:"path" json:"path"`
	Type         string       `mapstructure:"type" json:"type,omitempty"`
	Uid          string       `mapstructure:"uid" json:"uid,omitempty"`
	Gid          string       `mapstructure:"gid" json:"gid,omitempty"`
	Time         time.Time    `mapstructure:"time" json:"time,omitempty"`
	Size         string       `mapstructure:"size" json:"size,omitempty"`
	Md5digest    string       `mapstructure:"md5digest" json:"md5digest,omitempty"`
	Sha256digest string       `mapstructure:"sha256digest" json:"sh256digest,omitempty"`
	Link         string       `mapstructure:"link" json:"link,omitempty"`
	Digest       *file.Digest `json:"digest,omitempty"`
}

// PackageURL returns the PURL for the specific Arch Linux package (see https://github.com/package-url/purl-spec)
func (m AlpmMetadata) PackageURL(distro *linux.Release) string {
	qualifiers := map[string]string{
		PURLQualifierArch: m.Architecture,
	}

	if m.BasePackage != "" {
		qualifiers[PURLQualifierUpstream] = m.BasePackage
	}

	return packageurl.NewPackageURL(
		"alpm",
		distro.ID,
		m.Package,
		m.Version,
		purlQualifiers(
			qualifiers,
			distro,
		),
		"",
	).ToString()
}

func (m AlpmMetadata) OwnedFiles() (result []string) {
	s := strset.New()
	for _, f := range m.Files {
		if f.Path != "" {
			s.Add(f.Path)
		}
	}
	result = s.List()
	sort.Strings(result)
	return result
}
