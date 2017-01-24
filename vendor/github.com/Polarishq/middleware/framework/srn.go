package framework

import uuid "github.com/satori/go.uuid"
import "fmt"
import "strings"

// SRN represents Splunk Resource Names for identifying all resources across the Polaris product line
type SRN struct {
	serviceName  string
	resourceName string
	uuid         string
}

// NewSRN creates a new SRN object
func NewSRN(serviceName string, resourceName string) (*SRN, error) {
	if !(validateSRNSegment(&serviceName) && validateSRNSegment(&resourceName)) {
		return nil, fmt.Errorf("%s & %s can be only composed of lowercase alphabets and dashes", serviceName, resourceName)
	}
	srn := &SRN{
		serviceName:  serviceName,
		resourceName: resourceName,
		uuid:         uuid.NewV4().String(),
	}
	return srn, nil
}

// NewSRNFromString creates a new SRN object from an existing string representation
func NewSRNFromString(srn *string) (*SRN, error) {
	components := strings.Split(*srn, ":")
	if len(components) != 4 && components[0] != "srn" {
		return nil, fmt.Errorf("%s is not in valid SRN format srn:service-name:resource-name:guid", *srn)
	}
	_, err := uuid.FromString(components[3])
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid UUID. Error: %v", components[3], err)
	}
	if !(validateSRNSegment(&components[1]) && validateSRNSegment(&components[2])) {
		return nil, fmt.Errorf("%s & %s can be only composed of lowercase alphabets and dashes", components[1], components[2])
	}

	newSRN := &SRN{
		serviceName:  components[1],
		resourceName: components[2],
		uuid:         components[3],
	}

	return newSRN, nil
}

// ToString returns a string representation of the SRN
func (p *SRN) ToString() *string {
	srn := fmt.Sprintf("srn:%s:%s:%s", p.serviceName, p.resourceName, p.uuid)
	return &srn
}

func validateSRNSegment(component *string) bool {
	var isAlpha, isDash bool
	runeChecker := func(r rune) bool {
		isAlpha = r >= 'a' && r <= 'z'
		isDash = r == '-'
		return !(isAlpha || isDash)
	}
	if strings.IndexFunc(*component, runeChecker) != -1 {
		return false
	}
	return true
}
