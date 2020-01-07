package heatmap

import (
	"encoding/base64"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Create PNG", func() {
	It("should create png", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		data := &Heatmap{
			Annotations: types.Annotations{
				FontSize: 16,
				List: map[string]types.Annotation{
					"a": {Text: "a", Position: types.AnnotationPosition{X: 0.5, Y: 0.25}},
					"b": {Text: "b", Position: types.AnnotationPosition{X: 0.5, Y: 0.75}},
				},
			},
			Markers: types.Markers{
				Color: "#000000",
				List: map[string]types.Marker{
					"a": types.Marker{Height: 1, Width: 2, X: 0.25, Y: 0.5},
				},
			},
			Settings: types.Settings{
				AbundanceCap: 4,
				FillColor:    "blue",
				MinAbundance: 0,
			},
		}
		matrices := &types.Matrices{
			Abundance: [][]float64{
				{1, 1, 2, 2},
				{2, 2, 2, 3},
				{1, 1, 4, 1},
				{4, 2, 2, 3},
			},
		}
		settings := Settings{
			DownsampleThreshold: 4,
		}

		expected := "iVBORw0KGgoAAAANSUhEUgAAAFAAAABQCAIAAAABc2X6AAABF0lEQVR4nOyawa3CMBBEky8X8kt" +
			"JJ4hOTCWITlIKnYQDh1hwW6Mdafa9Uy7W6rHxRjbT+v1YFPSHpOzypymrA2F3EHYHYXcQdgdhdxB2B2F3EHYHYX" +
			"fWZdPcaU3xvIWXlutwE9Tc1/M5/f1SCI+Se/aeUnc4nXThj5amy0uHlqLV6R3ejtNzfM5CPbTSp3S57zDC7iDsDs" +
			"LuIOwOwu60fokv7lfZyTYcL5s9PIQLq36scq80wu4g7A7C7iDsDsLuIOwOwu4g7E62sPBW6I24fJz/Hls3d6cV+D" +
			"v7J7GdiWCaIgEwpjtK5LTGJqfntBRTWpru5LOUQK0k3ndUKxdFME26jdnD7iDsDsLuIOwOwu6UE34FAAD//7j1MVg/" +
			"tk3lAAAAAElFTkSuQmCC"

		createPNG(data, matrices, settings)
		pngContent, _ := afero.ReadFile(fs.Instance, "png/heatmap.png")
		Expect(base64.StdEncoding.EncodeToString((pngContent))).To(Equal(expected))
	})

	It("should downsample and create png", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		data := &Heatmap{
			Settings: types.Settings{
				AbundanceCap: 4,
				FillColor:    "blue",
				MinAbundance: 0,
			},
		}
		matrices := &types.Matrices{
			Abundance: [][]float64{
				{1, 1, 2, 2},
				{2, 2, 2, 3},
				{1, 1, 4, 1},
				{4, 2, 2, 3},
			},
		}
		settings := Settings{
			DownsampleThreshold: 2,
		}

		expected := "iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAIAAAADnC86AAAAU0lEQVR4nOzWsQmAYAyE0Siu" +
			"JTiRI9rYuYUTOILucMLf5F3/8SBVlnV/K9153XE7x+XPgcFgMBgMBo/fVFv+c9VzxGm/U4PBYDAYDG4MfwEA" +
			"AP//OmgHVhr0NIsAAAAASUVORK5CYII="

		createPNG(data, matrices, settings)
		pngContent, _ := afero.ReadFile(fs.Instance, "png/heatmap.png")
		Expect(base64.StdEncoding.EncodeToString((pngContent))).To(Equal(expected))
	})
})
