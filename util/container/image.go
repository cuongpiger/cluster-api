package container

import (
	"fmt"
	"github.com/distribution/reference"
	"path"
)

// ****************************************************** OBJECTS ******************************************************

// _______________________________________________________________________________________________________________ Image

// Image type represents the container image details.
type Image struct {
	Repository string
	Name       string
	Tag        string
	Digest     string
}

func (i Image) String() string {
	// repo/name [ ":" tag ] [ "@" digest ]
	ref := fmt.Sprintf("%s/%s", i.Repository, i.Name)
	if i.Tag != "" {
		ref = fmt.Sprintf("%s:%s", ref, i.Tag)
	}
	if i.Digest != "" {
		ref = fmt.Sprintf("%s@%s", ref, i.Digest)
	}
	return ref
}

// ImageFromString parses a docker image string into three parts: repo, tag and digest.
func ImageFromString(image string) (Image, error) {
	named, err := reference.ParseNamed(image)
	if err != nil {
		return Image{}, fmt.Errorf("couldn't parse image name: %v", err)
	}

	var repo, tag, digest string
	_, nameOnly := path.Split(reference.Path(named))
	if nameOnly != "" {
		// split out the part of the name after the last /
		lenOfCompleteName := len(named.Name())
		repo = named.Name()[:lenOfCompleteName-len(nameOnly)-1]
	}

	tagged, ok := named.(reference.Tagged)
	if ok {
		tag = tagged.Tag()
	}

	digested, ok := named.(reference.Digested)
	if ok {
		digest = digested.Digest().String()
	}

	return Image{Repository: repo, Name: nameOnly, Tag: tag, Digest: digest}, nil
}
