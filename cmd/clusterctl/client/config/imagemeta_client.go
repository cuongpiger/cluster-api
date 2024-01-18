package config

import (
	"fmt"
	"sigs.k8s.io/cluster-api/util/container"
	"strings"

	"github.com/pkg/errors"
)

// ****************************************************** CONSTS *******************************************************

const (
	// CertManagerImageComponent define the name of the cert-manager component in image overrides.
	CertManagerImageComponent = "cert-manager"

	imagesConfigKey = "images"
	allImageConfig  = "all"
)

// **************************************************** INTERFACES *****************************************************

// ImageMetaClient has methods to work with image meta configurations.
type ImageMetaClient interface {
	// AlterImage alters an image name according to the current image override configurations.
	AlterImage(component, image string) (string, error)
}

func newImageMetaClient(reader Reader) *imageMetaClient {
	return &imageMetaClient{
		reader:         reader,
		imageMetaCache: map[string]*imageMeta{},
	}
}

// ****************************************************** OBJECTS ******************************************************

// _____________________________________________________________________________________________________ imageMetaClient

type imageMetaClient struct {
	reader         Reader
	imageMetaCache map[string]*imageMeta
}

func (p *imageMetaClient) AlterImage(component, imageString string) (string, error) {
	image, err := container.ImageFromString(imageString)
	if err != nil {
		return "", err
	}

	// Gets the image meta that applies to the selected component/image; if none, returns early
	meta, err := p.getImageMeta(component, image.Name)
	if err != nil {
		return "", err
	}
	if meta == nil {
		return imageString, nil
	}

	// Apply the image meta to image name
	return meta.ApplyToImage(image), nil
}

// getImageMeta returns the image meta that applies to the selected component/image.
func (p *imageMetaClient) getImageMeta(component, imageName string) (*imageMeta, error) {
	// if the image meta for the component is already known, return it
	if im, ok := p.imageMetaCache[imageMetaCacheKey(component, imageName)]; ok {
		return im, nil
	}

	// Otherwise read the image override configurations.
	var meta map[string]imageMeta
	if err := p.reader.UnmarshalKey(imagesConfigKey, &meta); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal image override configurations")
	}

	// If there are not image override configurations, return.
	if meta == nil {
		p.imageMetaCache[imageMetaCacheKey(component, imageName)] = nil
		return nil, nil
	}

	// Gets the image configuration for:
	//	- all the components,
	//	- the component (and to all its images)
	//	- the selected component/image
	//	and returns the union of all the above.
	m := &imageMeta{}
	if allMeta, ok := meta[allImageConfig]; ok {
		m.Union(&allMeta)
	}

	if componentMeta, ok := meta[component]; ok {
		m.Union(&componentMeta)
	}
	p.imageMetaCache[component] = m

	if imageNameMeta, ok := meta[imageMetaCacheKey(component, imageName)]; ok {
		m.Union(&imageNameMeta)
	}
	p.imageMetaCache[imageMetaCacheKey(component, imageName)] = m

	return m, nil
}

// ___________________________________________________________________________________________________________ imageMeta

type imageMeta struct {
	// repository sets the container registry to pull images from.
	Repository string `json:"repository,omitempty"`

	// Tag allows to specify a tag for the images.
	Tag string `json:"tag,omitempty"`
}

// ApplyToImage changes an image name applying the transformations defined in the current imageMeta.
func (i *imageMeta) ApplyToImage(image container.Image) string {
	// apply transformations
	if i.Repository != "" {
		image.Repository = strings.TrimSuffix(i.Repository, "/")
	}
	if i.Tag != "" {
		image.Tag = i.Tag
	}

	// returns the resulting image name
	return image.String()
}

// Union allows to merge two imageMeta transformation; in case both the imageMeta defines new values for the same field,
// the other transformation takes precedence on the existing one.
func (i *imageMeta) Union(other *imageMeta) {
	if other.Repository != "" {
		i.Repository = other.Repository
	}
	if other.Tag != "" {
		i.Tag = other.Tag
	}
}

// ************************************************** PRIVATE METHODS **************************************************

func imageMetaCacheKey(component, imageName string) string {
	return fmt.Sprintf("%s/%s", component, imageName)
}
