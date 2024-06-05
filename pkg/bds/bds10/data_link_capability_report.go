package bds10

import (
	"errors"
	"fmt"

	"github.com/twuillemin/modes/pkg/bds/bds10/fields"
	"github.com/twuillemin/modes/pkg/bds/register"
)

// DataLinkCapabilityReport is a message at the format BDS 1,0
//
// Specified in Doc 9871 / D.2.4.1
type DataLinkCapabilityReport struct {
	ContinuationFlag                     fields.ContinuationFlag
	OverlayCommandCapability             fields.OverlayCommandCapability
	ACASStatus                           fields.ACASStatus
	ModeSSubnetworkVersion               fields.ModeSSubnetworkVersion
	TransponderEnhancedProtocolIndicator fields.TransponderEnhancedProtocolIndicator
	ModeSSpecificServicesCapability      fields.ModeSSpecificServicesCapability
	UplinkELMThroughputCapability        fields.UplinkELMThroughputCapability
	DownlinkELMThroughputCapability      fields.DownlinkELMThroughputCapability
	AircraftIdentificationCapability     fields.AircraftIdentificationCapability
	SquitterCapabilitySubfield           fields.SquitterCapabilitySubfield
	SurveillanceIdentifierCode           fields.SurveillanceIdentifierCode
	CommonUsageGICBCapability            fields.CommonUsageGICBCapability
	ACASHybridSurveillanceCapability     fields.ACASHybridSurveillanceCapability
	ACASGenerationCapability             fields.ACASGenerationCapability
	ACASApplicableDocument               fields.ACASApplicableDocument
	DTESubAddressStatuses                fields.DTESubAddressStatuses
}

// GetRegister returns the Register the message
func (message DataLinkCapabilityReport) GetRegister() register.Register {
	return register.BDS10
}

// CheckCoherency checks that the data of the message are somehow coherent, such as for example: no Reserved values, etc.
func (message DataLinkCapabilityReport) CheckCoherency() error {
	// If no ACAS, all ACAS info should be zeroed
	if message.ACASStatus == fields.ACASFailedOrStandBy {
		if message.ACASHybridSurveillanceCapability != fields.ACASHybridSurveillanceNotCapable {
			return errors.New("field ACASStatus is FailedOrStandBy but ACASHybridSurveillanceCapability is not HybridSurveillanceNotCapable")
		}
		if message.ACASGenerationCapability != fields.ACASGenerationNotCapable {
			return errors.New("field ACASStatus is FailedOrStandBy but ACASGenerationCapability is not ACASGenerationNotCapable")
		}
		if message.ACASApplicableDocument != fields.ACASApplicableDocument185 {
			return errors.New("field ACASStatus is FailedOrStandBy but ACASApplicableDocument is not ACASApplicableDocument185")
		}
	} else {
		if message.ACASApplicableDocument > 2 {
			return errors.New("field ACASApplicableDocument is a Reserved value")
		}
	}

	if message.ModeSSubnetworkVersion > 5 {
		return errors.New("field ModeSSubnetworkVersion is a Reserved value")
	}

	if message.UplinkELMThroughputCapability > 6 {
		return errors.New("field UplinkELMThroughputCapability is a Reserved value")
	}

	if message.DownlinkELMThroughputCapability > 6 {
		return errors.New("field DownlinkELMThroughputCapability is a Reserved value")
	}

	return nil
}

// ToString returns a basic, but readable, representation of the message
func (message DataLinkCapabilityReport) ToString() string {
	return fmt.Sprintf(""+
		"Message:                                 %v\n"+
		"Continuation Flag:                       %v\n"+
		"ACAS Status                              %v\n"+
		"ModeS Subnetwork Version                 %v\n"+
		"Transponder Enhanced Protocol Indicator  %v\n"+
		"ModeS Specific Services Capability       %v\n"+
		"Uplink ELM Throughput Capability         %v\n"+
		"Downlink ELM Throughput Capability       %v\n"+
		"Aircraft Identification Capability       %v\n"+
		"Squitter Capability Subfield             %v\n"+
		"Surveillance Identifier Code             %v\n"+
		"Common Usage GICB Capability             %v\n"+
		"ACAS Hybrid Surveillance Capability      %v\n"+
		"ACAS Generation Capability               %v\n"+
		"ACAS Applicable Document                 %v\n"+
		"DTE SubAddress Statuses                  %v\n",
		message.GetRegister().ToString(),
		message.OverlayCommandCapability.ToString(),
		message.ACASStatus.ToString(),
		message.ModeSSubnetworkVersion.ToString(),
		message.TransponderEnhancedProtocolIndicator.ToString(),
		message.ModeSSpecificServicesCapability.ToString(),
		message.UplinkELMThroughputCapability.ToString(),
		message.DownlinkELMThroughputCapability.ToString(),
		message.AircraftIdentificationCapability.ToString(),
		message.SquitterCapabilitySubfield.ToString(),
		message.SurveillanceIdentifierCode.ToString(),
		message.CommonUsageGICBCapability.ToString(),
		message.ACASHybridSurveillanceCapability.ToString(),
		message.ACASGenerationCapability.ToString(),
		message.ACASApplicableDocument.ToString(),
		message.DTESubAddressStatuses.ToString())
}

// ReadDataLinkCapabilityReport reads a message as a DataLinkCapabilityReport
func ReadDataLinkCapabilityReport(data []byte) (*DataLinkCapabilityReport, error) {

	if len(data) != 7 {
		return nil, errors.New("the data for Comm-B DataLinkCapabilityReport message must be 7 bytes long")
	}

	// First byte is simply the BDS format 0001 0000
	if data[0] != 0x10 {
		return nil, errors.New("the first byte of data is not 0x10")
	}

	// Bits 10 to 14 are reserved and must be 0
	if data[1]&0x7C != 0 {
		return nil, errors.New("the bits 10 to 14 are reserved and must be 0")
	}

	return &DataLinkCapabilityReport{
		ContinuationFlag:                     fields.ReadContinuationFlag(data),
		OverlayCommandCapability:             fields.ReadOverlayCommandCapability(data),
		ACASStatus:                           fields.ReadACASStatus(data),
		ModeSSubnetworkVersion:               fields.ReadModeSSubnetworkVersion(data),
		TransponderEnhancedProtocolIndicator: fields.ReadTransponderEnhancedProtocolIndicator(data),
		ModeSSpecificServicesCapability:      fields.ReadModeSSpecificServicesCapability(data),
		UplinkELMThroughputCapability:        fields.ReadUplinkELMThroughputCapability(data),
		DownlinkELMThroughputCapability:      fields.ReadDownlinkELMThroughputCapability(data),
		AircraftIdentificationCapability:     fields.ReadAircraftIdentificationCapability(data),
		SquitterCapabilitySubfield:           fields.ReadSquitterCapabilitySubfield(data),
		SurveillanceIdentifierCode:           fields.ReadSurveillanceIdentifierCode(data),
		CommonUsageGICBCapability:            fields.ReadCommonUsageGICBCapability(data),
		ACASHybridSurveillanceCapability:     fields.ReadACASHybridSurveillanceCapability(data),
		ACASGenerationCapability:             fields.ReadACASGenerationCapability(data),
		ACASApplicableDocument:               fields.ReadACASApplicableDocument(data),
		DTESubAddressStatuses:                fields.ReadDTESubAddressStatuses(data),
	}, nil
}
