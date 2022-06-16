package text

import (
	"flag"
	"testing"

	"github.com/zeromike/syft/syftinternal/formats/common/testutils"
)

var updateTextEncoderGoldenFiles = flag.Bool("update-text", false, "update the *.golden files for text encoder")

func TestTextDirectoryEncoder(t *testing.T) {
	testutils.AssertEncoderAgainstGoldenSnapshot(t,
		Format(),
		testutils.DirectoryInput(t),
		*updateTextEncoderGoldenFiles,
	)
}

func TestTextImageEncoder(t *testing.T) {
	testImage := "image-simple"
	testutils.AssertEncoderAgainstGoldenImageSnapshot(t,
		Format(),
		testutils.ImageInput(t, testImage, testutils.FromSnapshot()),
		testImage,
		*updateTextEncoderGoldenFiles,
	)
}
