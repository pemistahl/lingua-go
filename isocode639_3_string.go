// Code generated by "stringer -type=IsoCode639_3"; DO NOT EDIT.

package lingua

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AFR-0]
	_ = x[ARA-1]
	_ = x[AZE-2]
	_ = x[BEL-3]
	_ = x[BEN-4]
	_ = x[BOS-5]
	_ = x[BUL-6]
	_ = x[CAT-7]
	_ = x[CES-8]
	_ = x[CYM-9]
	_ = x[DAN-10]
	_ = x[DEU-11]
	_ = x[ELL-12]
	_ = x[ENG-13]
	_ = x[EPO-14]
	_ = x[EST-15]
	_ = x[EUS-16]
	_ = x[FAS-17]
	_ = x[FIN-18]
	_ = x[FRA-19]
	_ = x[GLE-20]
	_ = x[GUJ-21]
	_ = x[HEB-22]
	_ = x[HIN-23]
	_ = x[HRV-24]
	_ = x[HUN-25]
	_ = x[HYE-26]
	_ = x[IND-27]
	_ = x[ISL-28]
	_ = x[ITA-29]
	_ = x[JPN-30]
	_ = x[KAT-31]
	_ = x[KAZ-32]
	_ = x[KOR-33]
	_ = x[LAT-34]
	_ = x[LAV-35]
	_ = x[LIT-36]
	_ = x[LUG-37]
	_ = x[MAR-38]
	_ = x[MKD-39]
	_ = x[MON-40]
	_ = x[MRI-41]
	_ = x[MSA-42]
	_ = x[NLD-43]
	_ = x[NNO-44]
	_ = x[NOB-45]
	_ = x[PAN-46]
	_ = x[POL-47]
	_ = x[POR-48]
	_ = x[RON-49]
	_ = x[RUS-50]
	_ = x[SLK-51]
	_ = x[SLV-52]
	_ = x[SNA-53]
	_ = x[SOM-54]
	_ = x[SOT-55]
	_ = x[SPA-56]
	_ = x[SQI-57]
	_ = x[SRP-58]
	_ = x[SWA-59]
	_ = x[SWE-60]
	_ = x[TAM-61]
	_ = x[TEL-62]
	_ = x[TGL-63]
	_ = x[THA-64]
	_ = x[TSN-65]
	_ = x[TSO-66]
	_ = x[TUR-67]
	_ = x[UKR-68]
	_ = x[URD-69]
	_ = x[VIE-70]
	_ = x[XHO-71]
	_ = x[YOR-72]
	_ = x[ZHO-73]
	_ = x[ZUL-74]
	_ = x[MAL-75]
	_ = x[UnknownIsoCode639_3-76]
}

const _IsoCode639_3_name = "AFRARAAZEBELBENBOSBULCATCESCYMDANDEUELLENGEPOESTEUSFASFINFRAGLEGUJHEBHINHRVHUNHYEINDISLITAJPNKATKAZKORLATLAVLITLUGMARMKDMONMRIMSANLDNNONOBPANPOLPORRONRUSSLKSLVSNASOMSOTSPASQISRPSWASWETAMTELTGLTHATSNTSOTURUKRURDVIEXHOYORZHOZULMALUnknownIsoCode639_3"

var _IsoCode639_3_index = [...]uint8{0, 3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54, 57, 60, 63, 66, 69, 72, 75, 78, 81, 84, 87, 90, 93, 96, 99, 102, 105, 108, 111, 114, 117, 120, 123, 126, 129, 132, 135, 138, 141, 144, 147, 150, 153, 156, 159, 162, 165, 168, 171, 174, 177, 180, 183, 186, 189, 192, 195, 198, 201, 204, 207, 210, 213, 216, 219, 222, 225, 228, 247}

func (i IsoCode639_3) String() string {
	if i < 0 || i >= IsoCode639_3(len(_IsoCode639_3_index)-1) {
		return "IsoCode639_3(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IsoCode639_3_name[_IsoCode639_3_index[i]:_IsoCode639_3_index[i+1]]
}
