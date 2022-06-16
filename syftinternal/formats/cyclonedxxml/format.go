package cyclonedxxml

import (
	"github.com/CycloneDX/cyclonedx-go"
	"github.com/anchore/syft/syft/sbom"
	"github.com/zeromike/syft/syftinternal/formats/common/cyclonedxhelpers"
)

const ID sbom.FormatID = "cyclonedx-1-xml"

func Format() sbom.Format {
	return sbom.NewFormat(
		ID,
		encoder,
		cyclonedxhelpers.GetDecoder(cyclonedx.BOMFileFormatXML),
		cyclonedxhelpers.GetValidator(cyclonedx.BOMFileFormatXML),
	)
}
