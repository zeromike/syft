package cyclonedxjson

import (
	"io"

	"github.com/CycloneDX/cyclonedx-go"
	"github.com/anchore/syft/syft/sbom"
	"github.com/zeromike/syft/syftinternal/formats/common/cyclonedxhelpers"
)

func encoder(output io.Writer, s sbom.SBOM) error {
	bom := cyclonedxhelpers.ToFormatModel(s)
	enc := cyclonedx.NewBOMEncoder(output, cyclonedx.BOMFileFormatJSON)
	enc.SetPretty(true)

	err := enc.Encode(bom)
	return err
}
